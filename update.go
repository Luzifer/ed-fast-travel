package main

import (
	"log"

	"github.com/Luzifer/go_helpers/github"
	"github.com/Luzifer/rconfig"
)

func checkUpdates() {
	if version == "dev" {
		return
	}

	updater, err := github.NewUpdater(autoUpdateRepo, version)
	if err != nil {
		log.Printf("Could not initialize update engine: %s", err)
		return
	}

	if hasUpdate, err := updater.HasUpdate(false); err != nil {
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

	updater, err := github.NewUpdater(autoUpdateRepo, version)
	if err != nil {
		log.Printf("Could not initialize update engine: %s", err)
		return
	}

	if err := updater.Apply(); err != nil {
		log.Printf("Update failed: %s", err)
	} else {
		log.Printf("Update successful.")
	}
}
