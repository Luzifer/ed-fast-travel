package main

import (
	"log"

	"github.com/Luzifer/gobuilder/autoupdate"
	"github.com/Luzifer/rconfig"
)

func checkUpdates() {
	if version == "dev" {
		return
	}

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
