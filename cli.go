package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Luzifer/rconfig"
	"github.com/cheggaaa/pb"
	"github.com/fatih/color"
)

func doCLICalculation() {
	// Parse input
	startSystem := rconfig.Args()[1]
	targetSystem := rconfig.Args()[2]
	stopDistance, err := strconv.ParseFloat(rconfig.Args()[3], 64)
	if err != nil {
		log.Fatalf("Please use a valid number in format 0.00 as distance")
	}

	// Search system
	verboseLog("Searching your start / destination system...")
	start := starSystems.GetSystemByName(startSystem)
	target := starSystems.GetSystemByName(targetSystem)

	if start == nil {
		log.Fatalf("Could not find system '%s' in EDSM database.", startSystem)
	}
	if target == nil {
		log.Fatalf("Could not find system '%s' in EDSM database.", targetSystem)
	}

	linearDistance := start.Coords.DistanceLy(target.Coords)

	verboseLog("Found start system '%s' at coordinates (%.5f / %.5f / %.5f)",
		start.Name, start.Coords.X, start.Coords.Y, start.Coords.Z)
	verboseLog("Found destination system '%s' at coordinates (%.5f / %.5f / %.5f)",
		target.Name, target.Coords.X, target.Coords.Y, target.Coords.Z)
	verboseLog("Linear distance between that systems is %.2f Ly",
		linearDistance)

	// Calculate steps
	stopNo := 1
	totalFlight := 0.0
	rChan, eChan := starSystems.CalculateRoute(context.Background(), start, target, stopDistance)

	go func(eChan <-chan error) {
		for {
			select {
			case err := <-eChan:
				if err != nil {
					log.Fatalf("An error ocurred while plotting the route: %s", err)
				}
			}
		}
	}(eChan)

	bar := pb.New(10000)
	bar.ShowCounters = false
	if !cfg.Silent {
		bar.Start()
	}

	for stop := range rChan {
		if stop.TraceType == traceTypeProgress {
			bar.Set64(int64(stop.Progress * 10000))
			continue
		} else {
			bar.Finish()
		}

		fmt.Printf("%4d: '%s' %s with %s distance (total: %s)\n",
			stopNo,
			color.GreenString(stop.StarSystem.Name), stop.StarSystem.Coords,
			color.YellowString(fmt.Sprintf("%.2f Ly", stop.FlightDistance)), color.YellowString(fmt.Sprintf("%.2f Ly", stop.TotalFlightDistance)),
		)

		stopNo += 1
		totalFlight = stop.TotalFlightDistance

		if stop.Progress == 1 {
			break
		}
	}

	<-time.After(500 * time.Millisecond)

	verboseLog("Calculation shows an overhead of %.2f Ly in comparison to linear distance.", totalFlight-linearDistance)
}
