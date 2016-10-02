package main

import (
	"bufio"
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path"

	"github.com/cheggaaa/pb"
)

func generateGOBDatabase() error {
	tmp := newStarSystemDatabase()

	if err := readSystems(tmp); err != nil {
		return err
	}

	if err := readBodies(tmp); err != nil {
		return err
	}

	log.Printf("Creating files...")
	if err := os.MkdirAll(cfg.EDSMDumpPath, 0755); err != nil {
		return err
	}

	fp, err := os.Create(path.Join(cfg.EDSMDumpPath, readableDumpName))
	if err != nil {
		return err
	}
	defer fp.Close()

	fpzip, err := os.Create(path.Join(cfg.EDSMDumpPath, compressedDumpName))
	if err != nil {
		return err
	}
	defer fpzip.Close()

	zw := gzip.NewWriter(fpzip)

	log.Printf("Writing own database...")
	if err := gob.NewEncoder(fp).Encode(tmp); err != nil {
		return err
	}

	log.Printf("Writing gzipped database...")
	if err := gob.NewEncoder(zw).Encode(tmp); err != nil {
		return err
	}

	if err := zw.Flush(); err != nil {
		return err
	}
	if err := zw.Close(); err != nil {
		return err
	}

	return nil
}

func readBodies(tmp *starSystemDatabase) error {
	log.Printf("Retrieving bodies dump...")
	bodiesDump, err := http.Get(originalBodyDumpURL)
	if err != nil {
		return err
	}
	defer bodiesDump.Body.Close()

	bodiesFile, err := os.Create(path.Join(cfg.EDSMDumpPath, "bodies.jsonl"))
	if err != nil {
		log.Fatalf("Unable to create dump file")
	}
	defer func() {
		bodiesFile.Close()
		os.Remove(path.Join(cfg.EDSMDumpPath, "bodies.jsonl"))
	}()
	io.Copy(bodiesFile, bodiesDump.Body)
	bodiesFile.Seek(0, 0)

	bodiesScanner := bufio.NewScanner(bodiesFile)

	bar := pb.StartNew(0)
	for bodiesScanner.Scan() {
		bar.Increment()
		b := body{}
		if err := json.Unmarshal(bodiesScanner.Bytes(), &b); err != nil {
			return err
		}
		if !b.IsMainStar {
			continue
		}
		tmp.GetSystemByID(b.SystemID).Scoopable = b.IsScoopable()
	}
	bar.Finish()

	return nil
}

func readSystems(tmp *starSystemDatabase) error {
	log.Printf("Retrieving system dump...")
	systemsDump, err := http.Get(originalSystemDumpURL)
	if err != nil {
		return err
	}
	defer systemsDump.Body.Close()

	systemsFile, err := os.Create(path.Join(cfg.EDSMDumpPath, "systems.jsonl"))
	if err != nil {
		log.Fatalf("Unable to create dump file")
	}
	defer func() {
		systemsFile.Close()
		os.Remove(path.Join(cfg.EDSMDumpPath, "systems.jsonl"))
	}()
	io.Copy(systemsFile, systemsDump.Body)
	systemsFile.Seek(0, 0)

	log.Printf("Transforming input data...")
	var minX, minY, minZ float64 = math.MaxFloat64, math.MaxFloat64, math.MaxFloat64
	var maxX, maxY, maxZ float64 = -1 * math.MaxFloat64, -1 * math.MaxFloat64, -1 * math.MaxFloat64

	systemsScanner := bufio.NewScanner(systemsFile)

	bar := pb.StartNew(0)
	for systemsScanner.Scan() {
		system, err := starSystemFromEDDBData(systemsScanner.Bytes())
		if err != nil {
			log.Fatalf("ERR while parsing system: %s", err)
		}

		if system.Coords.X < minX {
			minX = system.Coords.X
		}
		if system.Coords.Y < minY {
			minY = system.Coords.Y
		}
		if system.Coords.Z < minZ {
			minZ = system.Coords.Z
		}
		if system.Coords.X > maxX {
			maxX = system.Coords.X
		}
		if system.Coords.Y > maxY {
			maxY = system.Coords.Y
		}
		if system.Coords.Z > maxZ {
			maxZ = system.Coords.Z
		}

		if err := tmp.AddSystem(system); err != nil {
			log.Fatalf("ERR while adding system %#v: %s", system, err)
		}
		bar.Increment()
	}
	bar.FinishPrint("Done")
	tmp.Min, tmp.Max = starCoordinate{minX, minY, minY}, starCoordinate{maxX, maxY, maxZ}

	return nil
}
