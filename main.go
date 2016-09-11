package main

import (
	"compress/bzip2"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/Luzifer/gobuilder/autoupdate"
	"github.com/Luzifer/rconfig"
	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
)

const (
	autoUpdateRepo  = "github.com/Luzifer/ed-fast-travel"
	autoUpdateLabel = "master"
	edsmDumpURL     = "http://assets.luzifer.io/systemsWithCoordinates.json.bz2"
)

var (
	cfg = struct {
		Color          bool   `flag:"color" vardefault:"color" description:"Use color for output"`
		EDSMDumpPath   string `flag:"data-path" default:"~/.local/share/ed-fast-travel" description:"Path to store EDSM data"`
		SelfUpdate     bool   `flag:"self-update" default:"false" description:"Update the tool to the latest version"`
		Silent         bool   `flag:"silent,s" default:"false" description:"Suppress every message except the flight plan"`
		UpdateData     bool   `flag:"update" default:"false" description:"Fetch latest dump from EDSM"`
		VersionAndExit bool   `flag:"version" default:"false" description:"Prints current version and exits"`
	}{}

	version = "dev"
)

func init() {
	rconfig.SetVariableDefaults(defaultSettings)
	if err := rconfig.Parse(&cfg); err != nil {
		log.Fatalf("Unable to parse commandline options: %s", err)
	}

	color.NoColor = !cfg.Color

	if cfg.VersionAndExit {
		fmt.Printf("ed-fast-travel %s\n", version)
		os.Exit(0)
	}

	if cfg.SelfUpdate {
		selfUpdate()
		os.Exit(1)
	}

	var err error
	cfg.EDSMDumpPath, err = homedir.Expand(cfg.EDSMDumpPath)
	if err != nil {
		log.Fatalf("Could not expand data-path: %s", err)
	}
}

func verboseLog(format string, args ...interface{}) {
	if !cfg.Silent {
		log.Printf(format, args...)
	}
}

func main() {
	checkUpdates()

	if _, err := os.Stat(path.Join(cfg.EDSMDumpPath, "dump.json")); err != nil || cfg.UpdateData {
		if err := refreshEDSMData(); err != nil {
			log.Fatalf("Unable to refresh EDSM data: %s", err)
		}
	}

	if len(rconfig.Args()) != 4 {
		fmt.Println("Usage: ed-fast-travel <start system> <target system> <distance between nav points>\n")
		fmt.Println("Example: ed-fast-travel 'Sol' 'Dryooe Prou GL-Y d369' 500\n  This will calculate stops on your route from Sol to Dryooe Prou GL-Y d369 every 500Ly")
		os.Exit(1)
	}

	// Parse input
	startSystem := rconfig.Args()[1]
	targetSystem := rconfig.Args()[2]
	stopDistance, err := strconv.ParseFloat(rconfig.Args()[3], 64)
	if err != nil {
		log.Fatalf("Please use a valid number in format 0.00 as distance")
	}

	// Load database
	verboseLog("Loading database...")
	starSystems, err := loadStarSystems()
	if err != nil {
		log.Fatalf("Could not load star systems from dump: %s", err)
	}

	// Search system
	verboseLog("Searching your start / destination system...")
	start := getSystemByName(starSystems, startSystem)
	target := getSystemByName(starSystems, targetSystem)

	if start == nil {
		log.Fatalf("Could not find system '%s' in EDSM database.", startSystem)
	}
	if target == nil {
		log.Fatalf("Could not find system '%s' in EDSM database.", targetSystem)
	}

	linearDistance := start.Coords.DistanceLy(target.Coords)

	verboseLog("Found start system '%s' at coordinates (%.5f / %.5f / %.5f)",
		start.Name, start.Coords.X, start.Coords.Y, start.Coords.Z)
	verboseLog("Found destination system '%s' at coordinates (%.5f / %.5f / %.5f)",
		target.Name, target.Coords.X, target.Coords.Y, target.Coords.Z)
	verboseLog("Linear distance between that systems is %.2f Ly",
		linearDistance)

	// Calculate steps
	totalFlight := 0.0
	stopNo := 1
	for start.Coords.DistanceLy(target.Coords) > 0 {
		stop := getSystemByNearestCoordinate(starSystems, start.Coords.PartVectorTarget(target.Coords, stopDistance), *start)
		dist := start.Coords.DistanceLy(stop.Coords)
		totalFlight += dist

		fmt.Printf("%4d: '%s' %s with %s distance (total: %s)\n",
			stopNo,
			color.GreenString(stop.Name), stop.Coords,
			color.YellowString(fmt.Sprintf("%.2f Ly", dist)), color.YellowString(fmt.Sprintf("%.2f Ly", totalFlight)),
		)

		start = stop
		stopNo += 1
	}

	verboseLog("Calculation shows an overhead of %.2f Ly in comparison to linear distance.", totalFlight-linearDistance)
}

func checkUpdates() {
	if hasUpdate, err := autoupdate.New(autoUpdateRepo, autoUpdateLabel).HasUpdate(); err != nil {
		log.Printf("Could not look for updates: %s", err)
	} else {
		if hasUpdate {
			log.Printf("An update to ed-fast-travel is available. Run %s --self-update to update.", rconfig.Args()[0])
		}
	}
}

func selfUpdate() {
	if version == "dev" {
		return
	}

	if err := autoupdate.New(autoUpdateRepo, autoUpdateLabel).SingleRun(); err != nil {
		log.Printf("Update failed: %s", err)
	} else {
		log.Printf("Update successful.")
	}
}

func getSystemByNearestCoordinate(systems []starSystem, coords starCoordinate, skipSystem starSystem) *starSystem {
	dist := math.MaxFloat64
	var storedSystem *starSystem

	for i := range systems {
		if d := systems[i].Coords.DistanceLy(coords); d < dist && systems[i].ID != skipSystem.ID {
			dist = d
			storedSystem = &systems[i]
		}
	}

	return storedSystem
}

func getSystemByName(systems []starSystem, name string) *starSystem {
	for i := range systems {
		if strings.ToLower(systems[i].Name) == strings.ToLower(name) {
			return &systems[i]
		}
	}

	return nil
}

func loadStarSystems() ([]starSystem, error) {
	starSystems := []starSystem{}
	dump, err := os.Open(path.Join(cfg.EDSMDumpPath, "dump.json"))
	if err != nil {
		return nil, err
	}
	defer dump.Close()

	return starSystems, json.NewDecoder(dump).Decode(&starSystems)
}

func refreshEDSMData() error {
	verboseLog("No local EDSM dump found or update forced, fetching dump...")
	if err := os.MkdirAll(cfg.EDSMDumpPath, 0755); err != nil {
		return err
	}

	if _, err := os.Stat(path.Join(cfg.EDSMDumpPath, "etag.txt")); err == nil {
		d, err := ioutil.ReadFile(path.Join(cfg.EDSMDumpPath, "etag.txt"))
		if err != nil {
			return err
		}
		eTag := string(d)

		resp, err := http.Head(edsmDumpURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.Header.Get("ETag") == eTag {
			return nil
		}
	}

	dump, err := os.Create(path.Join(cfg.EDSMDumpPath, "dump.json"))
	if err != nil {
		return err
	}
	defer dump.Close()

	resp, err := http.Get(edsmDumpURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := ioutil.WriteFile(path.Join(cfg.EDSMDumpPath, "etag.txt"), []byte(resp.Header.Get("ETag")), 0644); err != nil {
		return err
	}

	bzr := bzip2.NewReader(resp.Body)

	_, err = io.Copy(dump, bzr)
	return err
}
