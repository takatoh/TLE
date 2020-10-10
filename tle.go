package tle

type TLE struct {
	SatelliteNumber         uint64
	Classification          string
	InternationalDesignator string
	EpochYear               uint64
	EpochDay                float64
	FirstTimeDerivative     float64
	SecondTimeDerivative    float64
	BSTAR                   float64
	EphemerisType           uint64
	ElementNumber           uint64
	Inclination             float64    // degrees
	RightAscension          float64    // degrees
	Eccentricity            float64
	ArgumentOfPerigee       float64    // degrees
	MeanAnomaly             float64    // degrees
	MeanMotion              float64
	RevolutionNumber        uint64
}
