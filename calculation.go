package main

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"strings"
)

const numberOfBucketLayers = 6

func init() {
	gob.Register(starSystemDatabase{})
}

type traceResult struct {
	FlightDistance      float64     `json:"flight_distance"`
	Progress            float64     `json:"progress"`
	Requested           bool        `json:"requested"`
	StarSystem          *starSystem `json:"star_system"`
	TotalFlightDistance float64     `json:"total_flight_distance"`
}

type coordinateBucket struct {
	Bounds  [2]starCoordinate
	Buckets []*coordinateBucket
	Systems []int64
}

func (c *coordinateBucket) CreateSubBuckets(depth int) {
	if depth == 0 {
		return
	}

	minC := c.Bounds[0]
	maxC := c.Bounds[1]

	midX := c.Bounds[0].X + (c.Bounds[1].X-c.Bounds[0].X)/2
	midY := c.Bounds[0].Y + (c.Bounds[1].Y-c.Bounds[0].Y)/2
	midZ := c.Bounds[0].Z + (c.Bounds[1].Z-c.Bounds[0].Z)/2

	c.Buckets = []*coordinateBucket{
		{Bounds: [2]starCoordinate{{minC.X, minC.Y, minC.Z}, {midX, midY, midZ}}},
		{Bounds: [2]starCoordinate{{midX, minC.Y, minC.Z}, {maxC.X, midY, midZ}}},
		{Bounds: [2]starCoordinate{{minC.X, midY, minC.Z}, {midX, maxC.Y, midZ}}},
		{Bounds: [2]starCoordinate{{midX, midY, minC.Z}, {maxC.X, maxC.Y, midZ}}},

		{Bounds: [2]starCoordinate{{minC.X, minC.Y, midZ}, {midX, midY, maxC.Z}}},
		{Bounds: [2]starCoordinate{{midX, minC.Y, midZ}, {maxC.X, midY, maxC.Z}}},
		{Bounds: [2]starCoordinate{{minC.X, midY, midZ}, {midX, maxC.Y, maxC.Z}}},
		{Bounds: [2]starCoordinate{{midX, midY, midZ}, {maxC.X, maxC.Y, maxC.Z}}},
	}

	for i := range c.Buckets {
		c.Buckets[i].CreateSubBuckets(depth - 1)
	}
}

func (c coordinateBucket) ContainsCoordinate(in starCoordinate) bool {
	return c.Bounds[0].X <= in.X &&
		c.Bounds[0].Y <= in.Y &&
		c.Bounds[0].Z <= in.Z &&
		c.Bounds[1].X >= in.X &&
		c.Bounds[1].Y >= in.Y &&
		c.Bounds[1].Z >= in.Z
}

func (c *coordinateBucket) Add(in starCoordinate, id int64) bool {
	if !c.ContainsCoordinate(in) {
		return false
	}

	if len(c.Buckets) > 0 {
		for _, b := range c.Buckets {
			if ok := b.Add(in, id); ok {
				return true
			}
		}
	}

	c.Systems = append(c.Systems, id)
	return true
}

func (c coordinateBucket) GetByCoordinate(in starCoordinate) *coordinateBucket {
	if !c.ContainsCoordinate(in) {
		return nil
	}

	for _, b := range c.Buckets {
		if ret := b.GetByCoordinate(in); ret != nil {
			return ret
		}
	}

	return &c
}

type starSystemDatabase struct {
	Systems          map[int64]*starSystem
	NameDB           map[string]int64
	CoordinateBucket *coordinateBucket
}

func loadStarSystems() (*starSystemDatabase, error) {
	starSystems := newStarSystemDatabase()
	dump, err := os.Open(path.Join(cfg.EDSMDumpPath, readableDumpName))
	if err != nil {
		return starSystems, err
	}
	defer dump.Close()

	return starSystems, gob.NewDecoder(dump).Decode(&starSystems)
}

func newStarSystemDatabase() *starSystemDatabase {
	return &starSystemDatabase{
		Systems: make(map[int64]*starSystem),
		NameDB:  make(map[string]int64),
	}
}

func (systems *starSystemDatabase) GenerateCoordinateBuckets(min, max starCoordinate) error {

	startBucket := &coordinateBucket{
		Bounds:  [2]starCoordinate{min, max},
		Buckets: []*coordinateBucket{},
		Systems: []int64{},
	}

	startBucket.CreateSubBuckets(numberOfBucketLayers - 1)
	systems.CoordinateBucket = startBucket

	return nil
}

func (systems *starSystemDatabase) AddSystem(s *starSystem) error {
	systems.Systems[s.ID] = s
	systems.NameDB[strings.ToLower(s.Name)] = s.ID
	if !systems.CoordinateBucket.Add(s.Coords, s.ID) {
		return errors.New("Could not put coordinate into bucket: No bucket found.")
	}
	return nil
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
	doneChan := make(chan struct{})
	defer close(doneChan)

	starSystemsLock.RLock()
	defer starSystemsLock.RUnlock()

	keepRunning := true

	go func() {
		defer close(rChan)
		defer close(eChan)

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

	if route, err := cache.GetRoute(*a, *b, int64(stopDistance)); err == nil {
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

	if err := cache.StoreRoute(*a, *b, int64(stopDistance), cachedRoute); err != nil {
		log.Printf("Could not cache route: %s", err)
	}

	doneChan <- struct{}{}
}

func (systems starSystemDatabase) GetSystemByNearestCoordinate(coords starCoordinate, skipSystem starSystem) *starSystem {
	dist := math.MaxFloat64
	var storedSystem *starSystem

	bucket := systems.CoordinateBucket.GetByCoordinate(coords)
	if bucket == nil {
		return nil
	}

	for _, i := range bucket.Systems {
		if d := systems.Systems[i].Coords.DistanceLy(coords); d < dist && systems.Systems[i].ID != skipSystem.ID {
			dist = d
			storedSystem = systems.Systems[i]
		}
	}

	return storedSystem
}

func (systems starSystemDatabase) GetSystemByName(name string) *starSystem {
	if systemID, ok := systems.NameDB[strings.ToLower(name)]; ok {
		return systems.Systems[systemID]
	}

	return nil
}

func (systems starSystemDatabase) GetSystemByID(id int64) *starSystem {
	return systems.Systems[id]
}
