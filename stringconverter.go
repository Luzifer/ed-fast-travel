package main

import (
	"log"
	"strconv"
)

func mustFloat64(in string) float64 {
	f, err := strconv.ParseFloat(in, 64)
	if err != nil {
		log.Fatalf("Input %q is not a valid float64: %s", in, err)
	}
	return f
}

func mustBool(in string) bool {
	f, err := strconv.ParseBool(in)
	if err != nil {
		log.Fatalf("Input %q is not a valid bool: %s", in, err)
	}
	return f
}

func mustInt64(in string) int64 {
	f, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		log.Fatalf("Input %q is not a valid int64: %s", in, err)
	}
	return f
}
