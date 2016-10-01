package main

import (
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
	"path"

	"github.com/cheggaaa/pb"
)

func generateGOBDatabase() error {
	log.Printf("Retrieving original dump...")
	edsmDump, err := http.Get(originalDumpURL)
	if err != nil {
		return err
	}
	defer edsmDump.Body.Close()

	log.Printf("Reading dump...")
	originData := []*starSystem{}
	if err := json.NewDecoder(edsmDump.Body).Decode(&originData); err != nil {
		return err
	}

	log.Printf("Searching for galaxy bounds...")
	var minX, minY, minZ float64 = math.MaxFloat64, math.MaxFloat64, math.MaxFloat64
	var maxX, maxY, maxZ float64 = -1 * math.MaxFloat64, -1 * math.MaxFloat64, -1 * math.MaxFloat64

	for _, s := range originData {
		if s.Coords.X < minX {
			minX = s.Coords.X
		}
		if s.Coords.Y < minY {
			minY = s.Coords.Y
		}
		if s.Coords.Z < minZ {
			minZ = s.Coords.Z
		}
		if s.Coords.X > maxX {
			maxX = s.Coords.X
		}
		if s.Coords.Y > maxY {
			maxY = s.Coords.Y
		}
		if s.Coords.Z > maxZ {
			maxZ = s.Coords.Z
		}
	}

	log.Printf("Transforming input data...")
	tmp := newStarSystemDatabase()
	tmp.GenerateCoordinateBuckets(starCoordinate{minX, minY, minY}, starCoordinate{maxX, maxY, maxZ})

	bar := pb.StartNew(len(originData))
	for _, system := range originData {
		if err := tmp.AddSystem(system); err != nil {
			log.Fatalf("ERR while adding system %#v: %s", system, err)
		}
		bar.Increment()
	}
	bar.FinishPrint("Done")

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
