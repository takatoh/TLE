package tle

import (
	"errors"
	"strconv"
	"strings"
)


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


func NewTLE() *TLE {
	p := new(TLE)
	return p
}


func (tle *TLE) parseLine1(line1 strung) (*TLE, err) {
	line1 = strings.TrimRight(line1, "\r\n")
	if len(line1) != 69 && line1[0:1] != "1" {
		return tle, errors.New("Invalid line-1 for TLE.")
	}

	satelliteNumber, _                  := strconv.ParseUint(line1[2:7], 10, 64)
	classification                      := line1[7:8]
	internationalDesignatorYear         := line1[9:11]
	internationalDesignatorLaunchNumber := line1[11:14]
	internationalDesignatorPiece        := line1[14:17]
	epochYear, _                        := strconv.ParseUint(line1[18:20], 10, 64)
	epocheDay, _                        := strconv.ParseFloat(line1[20:32], 64)
	firstTimeDerivative, _              := strconv.ParseFloat(line1[33:43], 64)
	secondTimeDerivative, _             := strconv.ParseFloat(line1[44:52], 64)
	bstar, _                            := strconv.ParseFloat(line1[53:61], 64)
	ephemerisType, _                    := strconv.ParseUint(line1[62:63], 10, 64)
	elementNumber, _                    := strconv.ParseUint(line1[64:68], 10, 64)
//	checksum, _                         := strconv.Atoi(line1[68:69])

	tle.SatelliteNumber         = satelliteNumber
	tle.Classification          = classification
	tle.InternationalDesignator = internationalDesignatorYear +
		internationalDesignatorLaunchNumber +
		internationalDesignatorPiece
	tle.EpochYear               = epochYear
	tle.EpochDay                = epocheDay
	tle.FirstTimeDerivative     = firstTimeDerivative
	tle.SecondTimeDerivative    = secondTimeDerivative
	tle.BSTAR                   = bstar
	tle.EphemerisType           = ephemerisType
	tle.ElementNumber           = elementNumber

	return tle, nil
}
