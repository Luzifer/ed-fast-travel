package main

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go assets/

import (
	"compress/bzip2"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"

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
		Color                  bool          `flag:"color" vardefault:"color" description:"Use color for output"`
		DisableSoftwareControl bool          `flag:"disable-software-control" default:"false" description:"Do not let web-users control update / shutdown"`
		EDSMDumpPath           string        `flag:"data-path" default:"~/.local/share/ed-fast-travel" description:"Path to store EDSM data"`
		Listen                 string        `flag:"listen" default:":3000" description:"IP/Port to listen on when starting in web mode"`
		SelfUpdate             bool          `flag:"self-update" default:"false" description:"Update the tool to the latest version"`
		Silent                 bool          `flag:"silent,s" default:"false" description:"Suppress every message except the flight plan"`
		UpdateData             bool          `flag:"update" default:"false" description:"Fetch latest dump from EDSM"`
		VersionAndExit         bool          `flag:"version" default:"false" description:"Prints current version and exits"`
		WebRouteTimeout        time.Duration `flag:"web-route-timeout" default:"10m" description:"Timout for route calculations requested via web interface"`
		WebRouteStopMin        float64       `flag:"web-route-stop-min" default:"100" description:"Min distance between stops"`
	}{}

	version = "dev"

	starSystems starSystemDatabase
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

func printHelp() {
	fmt.Println("Usage for web interface:  ed-fast-travel")
	fmt.Println("Usage for CLI calculation: ed-fast-travel <start system> <target system> <distance between nav points>\n")
	fmt.Println("Example: ed-fast-travel 'Sol' 'Dryooe Prou GL-Y d369' 500\n  This will calculate stops on your route from Sol to Dryooe Prou GL-Y d369 every 500Ly")
}

func main() {
	checkUpdates()

	if _, err := os.Stat(path.Join(cfg.EDSMDumpPath, "dump.json")); err != nil || cfg.UpdateData {
		if err := refreshEDSMData(); err != nil {
			log.Fatalf("Unable to refresh EDSM data: %s", err)
		}
	}

	// Load database
	verboseLog("Loading database...")
	var err error
	starSystems, err = loadStarSystems()
	if err != nil {
		log.Fatalf("Could not load star systems from dump: %s", err)
	}

	switch len(rconfig.Args()) {
	case 4:
		doCLICalculation()
	case 1:
		startWebService()
	default:
		printHelp()
		os.Exit(1)
	}
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
