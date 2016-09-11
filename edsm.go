package main

import (
	"fmt"
	"math"
)

// {
//   "name": "Synuefai TQ-S a113-5",
//   "coords": {
//     "x": -348.5625,
//     "y": -41.9375,
//     "z": 154.65625
//   },
//   "id": "30990",
//   "date": "2015-05-12 15:29:33"
// }

type starSystem struct {
	Name   string         `json:"name"`
	Coords starCoordinate `json:"coords"`
	ID     int64          `json:"id,string"`
	Date   string         `json:"date"` // TODO: Use date parser
}

type starCoordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func (s starCoordinate) DistanceLy(in starCoordinate) float64 {
	return math.Sqrt(
		math.Pow(math.Abs(s.X-in.X), 2) +
			math.Pow(math.Abs(s.Y-in.Y), 2) +
			math.Pow(math.Abs(s.Z-in.Z), 2))
}

func (s starCoordinate) PartVectorTarget(destination starCoordinate, partLength float64) starCoordinate {
	dist := s.DistanceLy(destination)
	fraction := partLength / dist

	// If the max length is bigger than the distance, advice to jump directly
	if fraction >= 1 {
		return destination
	}

	return starCoordinate{
		X: s.directedPartDistanceAdd(s.X, destination.X, fraction),
		Y: s.directedPartDistanceAdd(s.Y, destination.Y, fraction),
		Z: s.directedPartDistanceAdd(s.Z, destination.Z, fraction),
	}
}

func (s starCoordinate) directedPartDistanceAdd(a, b, fraction float64) float64 {
	xSign := 1.0
	if a > b {
		xSign = -1.0
	}
	return a + math.Abs(a-b)*fraction*xSign
}

func (s starCoordinate) String() string {
	return fmt.Sprintf("(%.5f / %.5f / %.5f)", s.X, s.Y, s.Z)
}
