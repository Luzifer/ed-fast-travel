package main

import (
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"os"
	"path"
)

func generateGOBDatabase() error {
	edsmDump, err := http.Get(originalDumpURL)
	if err != nil {
		return err
	}
	defer edsmDump.Body.Close()

	tmp := starSystemDatabase{}
	if err := json.NewDecoder(edsmDump.Body).Decode(&tmp); err != nil {
		return err
	}

	if err := os.MkdirAll(cfg.EDSMDumpPath, 0755); err != nil {
		return err
	}

	fp, err := os.Create(path.Join(cfg.EDSMDumpPath, "dump.bin"))
	if err != nil {
		return err
	}
	defer fp.Close()

	fpzip, err := os.Create(path.Join(cfg.EDSMDumpPath, "systemsWithCoordinates.bin.gz"))
	if err != nil {
		return err
	}
	defer fpzip.Close()

	zw := gzip.NewWriter(fpzip)

	if err := gob.NewEncoder(fp).Encode(tmp); err != nil {
		return err
	}
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
