package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"strings"
)

type traceResult struct {
	FlightDistance      float64     `json:"flight_distance"`
	Progress            float64     `json:"progress"`
	Requested           bool        `json:"requested"`
	StarSystem          *starSystem `json:"star_system"`
	TotalFlightDistance float64     `json:"total_flight_distance"`
}

type starSystemDatabase []starSystem

func loadStarSystems() (starSystemDatabase, error) {
	starSystems := starSystemDatabase{}
	dump, err := os.Open(path.Join(cfg.EDSMDumpPath, "dump.json"))
	if err != nil {
		return nil, err
	}
	defer dump.Close()

	return starSystems, json.NewDecoder(dump).Decode(&starSystems)
}

func (systems starSystemDatabase) CalculateRoute(ctx context.Context, a, b *starSystem, stopDistance float64) (<-chan traceResult, <-chan error) {
	out := make(chan traceResult, 100)
	err := make(chan error, 10)

	if a == nil || b == nil {
		close(out)
		err <- errors.New("Systems must not be nil")
	} else {
		go systems.startRouteTracer(ctx, out, err, a, b, stopDistance)
	}

	return out, err
}

func (systems starSystemDatabase) startRouteTracer(ctx context.Context, rChan chan traceResult, eChan chan error, a, b *starSystem, stopDistance float64) {
	defer close(rChan)
	defer close(eChan)

	doneChan := make(chan struct{})
	defer close(doneChan)

	keepRunning := true

	go func() {
		for {
			select {
			case <-doneChan:
				return
			case <-ctx.Done():
				keepRunning = false
				eChan <- fmt.Errorf("Calculation deadline of %s exceeded.", cfg.WebRouteTimeout)
			}
		}
	}()

	if route, err := cache.GetRoute(*a, *b); err == nil {
		for _, tr := range route {
			rChan <- tr

			if !keepRunning {
				break
			}
		}
		return
	}

	cachedRoute := []traceResult{}
	oldStop := a

	t := traceResult{
		StarSystem: a,
		Requested:  true,
	}
	rChan <- t
	cachedRoute = append(cachedRoute, t)

	totalFlight := 0.0
	for oldStop.Coords.DistanceLy(b.Coords) > 0 && keepRunning {
		stop := systems.GetSystemByNearestCoordinate(oldStop.Coords.PartVectorTarget(b.Coords, stopDistance), *oldStop)
		dist := oldStop.Coords.DistanceLy(stop.Coords)
		totalFlight += dist

		isRequested := stop.ID == b.ID

		t = traceResult{
			FlightDistance:      dist,
			Progress:            (a.Coords.DistanceLy(b.Coords) - stop.Coords.DistanceLy(b.Coords)) / a.Coords.DistanceLy(b.Coords),
			Requested:           isRequested,
			StarSystem:          stop,
			TotalFlightDistance: totalFlight,
		}

		rChan <- t
		cachedRoute = append(cachedRoute, t)

		oldStop = stop
	}

	if err := cache.StoreRoute(*a, *b, cachedRoute); err != nil {
		log.Printf("Could not cache route: %s", err)
	}

	doneChan <- struct{}{}
}

func (systems starSystemDatabase) GetSystemByNearestCoordinate(coords starCoordinate, skipSystem starSystem) *starSystem {
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

func (systems starSystemDatabase) GetSystemByName(name string) *starSystem {
	for i := range systems {
		if strings.ToLower(systems[i].Name) == strings.ToLower(name) {
			return &systems[i]
		}
	}

	return nil
}

func (systems starSystemDatabase) GetSystemByID(id int64) *starSystem {
	for i := range systems {
		if systems[i].ID == id {
			return &systems[i]
		}
	}

	return nil
}
