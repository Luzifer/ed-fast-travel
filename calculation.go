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

	"github.com/cheggaaa/pb"
)

const (
	traceTypeFlightStop  = "flight_stop"
	traceTypeProgress    = "trace_progress"
	numberOfBucketLayers = 6
)

func init() {
	gob.Register(starSystemDatabase{})
}

type traceResult struct {
	TraceType           string      `json:"trace_type"`
	FlightDistance      float64     `json:"flight_distance"`
	Progress            float64     `json:"progress"`
	Requested           bool        `json:"requested"`
	StarSystem          *starSystem `json:"star_system"`
	TotalFlightDistance float64     `json:"total_flight_distance"`
}

type coordinateBucket struct {
	Parent  *coordinateBucket
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
		{Parent: c, Bounds: [2]starCoordinate{{minC.X, minC.Y, minC.Z}, {midX, midY, midZ}}},
		{Parent: c, Bounds: [2]starCoordinate{{midX, minC.Y, minC.Z}, {maxC.X, midY, midZ}}},
		{Parent: c, Bounds: [2]starCoordinate{{minC.X, midY, minC.Z}, {midX, maxC.Y, midZ}}},
		{Parent: c, Bounds: [2]starCoordinate{{midX, midY, minC.Z}, {maxC.X, maxC.Y, midZ}}},

		{Parent: c, Bounds: [2]starCoordinate{{minC.X, minC.Y, midZ}, {midX, midY, maxC.Z}}},
		{Parent: c, Bounds: [2]starCoordinate{{midX, minC.Y, midZ}, {maxC.X, midY, maxC.Z}}},
		{Parent: c, Bounds: [2]starCoordinate{{minC.X, midY, midZ}, {midX, maxC.Y, maxC.Z}}},
		{Parent: c, Bounds: [2]starCoordinate{{midX, midY, midZ}, {maxC.X, maxC.Y, maxC.Z}}},
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

func (c coordinateBucket) GetSystems() []int64 {
	out := c.Systems

	for _, b := range c.Buckets {
		out = append(out, b.GetSystems()...)
	}

	return out
}

func (c coordinateBucket) InSphere(center starCoordinate, radius float64) bool {
	min, max := c.Bounds[0], c.Bounds[1]

	return c.ContainsCoordinate(center) ||
		center.DistanceLy(min) <= radius ||
		center.DistanceLy(starCoordinate{min.X, min.Y, max.Z}) <= radius ||
		center.DistanceLy(starCoordinate{min.X, max.Y, max.Z}) <= radius ||
		center.DistanceLy(starCoordinate{max.X, min.Y, max.Z}) <= radius ||
		center.DistanceLy(starCoordinate{max.X, max.Y, min.Z}) <= radius ||
		center.DistanceLy(starCoordinate{max.X, min.Y, min.Z}) <= radius ||
		center.DistanceLy(starCoordinate{min.X, max.Y, min.Z}) <= radius ||
		center.DistanceLy(max) <= radius
}

func (c coordinateBucket) GetFilledBucketsBySphere(center starCoordinate, radius float64) []*coordinateBucket {
	out := []*coordinateBucket{}

	if c.InSphere(center, radius) && len(c.Systems) > 0 {
		out = append(out, &c)
	}

	for _, b := range c.Buckets {
		out = append(out, b.GetFilledBucketsBySphere(center, radius)...)
	}

	return out
}

type starSystemDatabase struct {
	Min, Max         starCoordinate
	Systems          map[int64]*starSystem
	NameDB           map[string]int64
	coordinateBucket *coordinateBucket
}

func loadStarSystems() (*starSystemDatabase, error) {
	starSystems := newStarSystemDatabase()
	dump, err := os.Open(path.Join(cfg.EDSMDumpPath, readableDumpName))
	if err != nil {
		return starSystems, err
	}
	defer dump.Close()

	if err := gob.NewDecoder(dump).Decode(&starSystems); err != nil {
		return nil, err
	}

	starSystems.GenerateCoordinateBuckets(starSystems.Min, starSystems.Max)

	bar := pb.StartNew(len(starSystems.Systems))
	for id, sys := range starSystems.Systems {
		starSystems.coordinateBucket.Add(sys.Coords, id)
		bar.Increment()
	}
	bar.FinishPrint("Database loaded.")

	return starSystems, nil
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
	systems.coordinateBucket = startBucket

	return nil
}

func (systems *starSystemDatabase) AddSystem(s *starSystem) error {
	systems.Systems[s.ID] = s
	systems.NameDB[strings.ToLower(s.Name)] = s.ID
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

	forbiddenSystemIDs := []int64{}
	flightPlan := []*starSystem{}

	oldStop := a
	for oldStop.Coords.DistanceLy(b.Coords) > 0 && keepRunning {
		forbiddenSystemIDs = append(forbiddenSystemIDs, oldStop.ID)
		idealTargetCoordinate := oldStop.Coords.PartVectorTarget(b.Coords, stopDistance)

		sphericMatchSystems := bucketListToSystemList(systems.coordinateBucket.GetFilledBucketsBySphere(idealTargetCoordinate, stopDistance*.1))
		sphericMatchSystems = sphericMatchSystems.
			Filter(func(s *starSystem) bool { return s.Coords.DistanceLy(oldStop.Coords) <= stopDistance }).
			Filter(func(s *starSystem) bool {
				for _, id := range forbiddenSystemIDs {
					if s.ID == id {
						return false
					}
				}
				return true
			})

		stop := sphericMatchSystems.GetNearestSystem(idealTargetCoordinate)

		if stop == nil {
			if len(flightPlan) <= 1 {
				eChan <- errors.New("Was not able to calculate first stop, canceling calculation now.")
				return
			}

			flightPlan = flightPlan[0 : len(flightPlan)-1]
			oldStop = flightPlan[len(flightPlan)-1]
			continue
		}

		flightPlan = append(flightPlan, stop)
		oldStop = stop

		rChan <- traceResult{
			TraceType: traceTypeProgress,
			Progress:  (a.Coords.DistanceLy(b.Coords) - stop.Coords.DistanceLy(b.Coords)) / a.Coords.DistanceLy(b.Coords),
		}
	}

	cachedRoute := []traceResult{}

	t := traceResult{
		TraceType:  traceTypeFlightStop,
		StarSystem: a,
	}
	rChan <- t
	cachedRoute = append(cachedRoute, t)

	totalFlight := 0.0
	oldStop = a
	for _, s := range flightPlan {
		totalFlight += oldStop.Coords.DistanceLy(s.Coords)
		t = traceResult{
			TraceType:           traceTypeFlightStop,
			FlightDistance:      oldStop.Coords.DistanceLy(s.Coords),
			Progress:            (a.Coords.DistanceLy(b.Coords) - s.Coords.DistanceLy(b.Coords)) / a.Coords.DistanceLy(b.Coords),
			StarSystem:          s,
			TotalFlightDistance: totalFlight,
		}

		rChan <- t
		cachedRoute = append(cachedRoute, t)
		oldStop = s
	}

	if err := cache.StoreRoute(*a, *b, int64(stopDistance), cachedRoute); err != nil {
		log.Printf("Could not cache route: %s", err)
	}

	doneChan <- struct{}{}
}

func (systems starSystemDatabase) GetSystemByNearestCoordinate(coords starCoordinate, skipSystem starSystem) *starSystem {
	var storedSystem *starSystem

	bucket := systems.coordinateBucket.GetByCoordinate(coords)
	if bucket == nil {
		log.Printf("Could not find bucket for coordinate %#v (outer bounds: %#v / %#v)", coords, systems.Min, systems.Max)
		return nil
	}

	for storedSystem == nil {
		dist := math.MaxFloat64

		for _, i := range bucket.GetSystems() {
			if d := systems.Systems[i].Coords.DistanceLy(coords); d < dist && systems.Systems[i].ID != skipSystem.ID {
				dist = d
				storedSystem = systems.Systems[i]
			}
		}

		bucket = bucket.Parent
		if bucket == nil {
			log.Printf("Parent was nil.")
			break
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

type systemList []*starSystem

func (s systemList) Filter(f func(system *starSystem) bool) systemList {
	out := systemList{}

	for _, sys := range s {
		if f(sys) {
			out = append(out, sys)
		}
	}

	return out
}

func (s systemList) GetNearestSystem(coords starCoordinate) *starSystem {
	lowestDist := math.MaxFloat64
	var result *starSystem

	for i := range s {
		sys := s[i]
		if dist := sys.Coords.DistanceLy(coords); dist < lowestDist {
			result = sys
			lowestDist = dist
		}
	}

	return result
}

func bucketListToSystemList(c []*coordinateBucket) systemList {
	out := systemList{}

	for _, cb := range c {
		for _, id := range cb.GetSystems() {
			out = append(out, starSystems.GetSystemByID(id))
		}
	}

	return out
}
