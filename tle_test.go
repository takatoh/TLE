package tle

import (
	"math"
	"testing"
)

const (
	line1 = "1 25544U 98067A   20284.12598781  .00000554  00000-0  18038-4 0  9995\n"
	line2 = "2 25544  51.6443 139.7571 0001313  18.8517 125.1639 15.49294162249832\n"
)


func TestParseLine1Success(t *testing.T) {
	tle := NewTLE()
	result, err := tle.parseLine1(line1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if result.SatelliteNumber          != 25544 ||
		result.Classification          != "U" ||
		result.InternationalDesignator != "98067A" ||
		result.EpochYear               != 20 ||
		result.EpochDay                != 284.12598781 ||
		result.FirstTimeDerivative     != ".00000554" ||
		result.SecondTimeDerivative    != "00000-0" ||
		result.BSTAR                   != 0.18038 * math.Pow(10.0, -4.0) ||
		result.EphemerisType           != 0 ||
		result.ElementNumber           != 999 {
			t.Fatal("failed test")
	}
}


func TestParseLine2Success(t *testing.T) {
	tle := NewTLE()
	result, err := tle.parseLine2(line2)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if result.Inclination        != 51.6443 ||
		result.RightAscension    != 139.7571 ||
		result.Eccentricity      != 0001313 ||
		result.ArgumentOfPerigee != 18.8517 ||
		result.MeanAnomaly       != 125.1639 ||
		result.MeanMotion        != 15.49294162 ||
		result.RevolutionNumber  != 24983 {
			t.Fatal("failed test")
	}

}