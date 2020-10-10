package tle

import (
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

	if tle.SatelliteNumber          != 25544 ||
		tle.Classification          != "U" ||
		tle.InternationalDesignator != "98067A" ||
		tle.EpochYear               != 20 ||
		tle.EpochDay                != 284.12598781 ||
		tle.FirstTimeDerivative     != ".00000554" ||
		tle.SecondTimeDerivative    != "00000-0" ||
		tle.BSTAR                   != 0.18038 * 10 ** -4 ||
		tle.EphemerisType           != 0 ||
		tle.ElementNumber           != 999 {
			t.Fatal("failed test")
	}
}


func TestParseLine2Success(t *testing.T) {
	tle := NewTLE()
	reuslt, err := tle.parseLine2(line2)
	if err != nil {
		f.Fatalf("failed test %#v", err)
	}

	if Inclination        != 51.6443 ||
		RightAscension    != 139.7571 ||
		Eccentricity      != 0001313 ||
		ArgumentOfPerigee != 18.8517 ||
		MeanAnomaly       != 125.1639 ||
		MeanMotion        != 15.49294162 ||
		RevolutionNumber  != 24983 {
			t.Fatal("failed test")
	}

}