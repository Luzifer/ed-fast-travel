package main

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"os"
	"path"
	"strings"
	"time"
)

type traceResult struct {
	Progress            float64
	FlightDistance      float64
	TotalFlightDistance float64
	StarSystem          *starSystem
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

	deadline, deadlineExists := ctx.Deadline()
	oldStop := a

	totalFlight := 0.0
	for a.Coords.DistanceLy(b.Coords) > 0 && (!deadlineExists || deadline.After(time.Now())) {
		stop := systems.GetSystemByNearestCoordinate(oldStop.Coords.PartVectorTarget(b.Coords, stopDistance), *oldStop)
		dist := oldStop.Coords.DistanceLy(stop.Coords)
		totalFlight += dist

		rChan <- traceResult{
			Progress:            (a.Coords.DistanceLy(b.Coords) - stop.Coords.DistanceLy(b.Coords)) / a.Coords.DistanceLy(b.Coords),
			FlightDistance:      dist,
			TotalFlightDistance: totalFlight,
			StarSystem:          stop,
		}

		oldStop = stop
	}
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
