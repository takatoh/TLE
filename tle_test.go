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

	if result.SatelliteNumber != 25544 { t.Fatal("failed test: SatelliteNumber") }
	if result.Classification != "U" { t.Fatal("failed test: Classification") }
	if result.InternationalDesignator != "98067A" { t.Fatal("failed test: InternationalDesignator") }
	if result.EpochYear != 20 { t.Fatal("failed test: EpochYear") }
	if result.EpochDay != 284.12598781 { t.Fatal("failed test: EpochDay") }
	if result.FirstTimeDerivative != " .00000554" { t.Fatal("failed test: FirstTimeDerivative") }
	if result.SecondTimeDerivative != " 00000-0" { t.Fatal("failed test: SecondTimeDerivative") }
	if result.BSTAR != 0.18038 * math.Pow(10.0, -4.0) { t.Fatal("failed test: BSTAR") }
	if result.EphemerisType != 0 { t.Fatal("failed test: EphemerisType") }
	if result.ElementNumber != 999 { t.Fatal("failed test: ElementNumber") }
}


func TestParseLine2Success(t *testing.T) {
	tle := NewTLE()
	result, err := tle.parseLine2(line2)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if result.Inclination != 51.6443 { t.Fatal("failed test: Inclination") }
	if result.RightAscension != 139.7571 { t.Fatal("failed test: RightAscension") }
	if result.Eccentricity != 0001313 { t.Fatal("failed test: Eccentricity") }
	if result.ArgumentOfPerigee != 18.8517 { t.Fatal("failed test: ArgumentOfPerigee") }
	if result.MeanAnomaly != 125.1639 { t.Fatal("failed test: MeanAnomaly") }
	if result.MeanMotion != 15.49294162 { t.Fatal("failed test: MeanMotion") }
	if result.RevolutionNumber != 24983 { t.Fatal("failed test: RevolutionNumber") }
}
