package main

import (
	"fmt"
	"testing"
)

type starDistanceTestCase struct {
	CoordA, CoordB   starCoordinate
	ExpectedDistance float64
}

func floatEqual2(a, b float64) bool {
	return fmt.Sprintf("%.2f", a) == fmt.Sprintf("%.2f", b)
}

func TestStarDistances(t *testing.T) {
	cases := []starDistanceTestCase{
		{
			CoordA:           starCoordinate{X: -348.5625, Y: -41.9375, Z: 154.65625},
			CoordB:           starCoordinate{X: 0, Y: 0, Z: 0},
			ExpectedDistance: 383.6314976335526,
		},
		{
			CoordA:           starCoordinate{X: 999.875, Y: 676.75, Z: 738.28125},
			CoordB:           starCoordinate{X: 0, Y: 0, Z: 0},
			ExpectedDistance: 1415.20309,
		},
		{
			CoordA:           starCoordinate{X: 999.875, Y: 676.75, Z: 738.28125},
			CoordB:           starCoordinate{X: -9931.21875, Y: -1155.15625, Z: 20790.59375},
			ExpectedDistance: 22911.57,
		},
	}

	for _, tc := range cases {
		if dist := tc.CoordA.DistanceLy(tc.CoordB); !floatEqual2(dist, tc.ExpectedDistance) {
			t.Errorf("Distance between %#v and %#v was expected to be %.2f but is %.2f",
				tc.CoordA,
				tc.CoordB,
				tc.ExpectedDistance,
				dist)
		}

		if dist := tc.CoordB.DistanceLy(tc.CoordA); !floatEqual2(dist, tc.ExpectedDistance) {
			t.Errorf("Distance between %#v and %#v was expected to be %.2f but is %.2f",
				tc.CoordB,
				tc.CoordA,
				tc.ExpectedDistance,
				dist)
		}
	}
}

func TestPartVectorTarget(t *testing.T) {
	source := starCoordinate{}
	dest := starCoordinate{X: 100, Y: 100, Z: 100}

	part := source.PartVectorTarget(dest, 58.88972745734183)
	if !floatEqual2(part.X, 34) || !floatEqual2(part.Y, 34) || !floatEqual2(part.Z, 34) {
		t.Errorf("Did not get expected coordinates: %#v (expected {X: 34, Y: 34, Z: 34})", part)
	}

	dest = starCoordinate{X: -100, Y: 100, Z: -100}
	part = source.PartVectorTarget(dest, 58.88972745734183)
	if !floatEqual2(part.X, -34) || !floatEqual2(part.Y, 34) || !floatEqual2(part.Z, -34) {
		t.Errorf("Did not get expected coordinates: %#v (expected {X: -34, Y: 34, Z: -34})", part)
	}

	part = source.PartVectorTarget(dest, 169.74097914175)
	if !floatEqual2(part.X, -98) || !floatEqual2(part.Y, 98) || !floatEqual2(part.Z, -98) {
		t.Errorf("Did not get expected coordinates: %#v (expected {X: -98, Y: 98, Z: -98})", part)
	}
}
