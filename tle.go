package tle

import (
	"errors"
//	"fmt"
	"math"
	"strconv"
	"strings"
)


type TLE struct {
	SatelliteNumber         uint64
	Classification          string
	InternationalDesignator string
	EpochYear               uint64
	EpochDay                float64
	FirstTimeDerivative     string
	SecondTimeDerivative    string
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


func (tle *TLE) parseLine1(line1 string) (*TLE, error) {
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
	firstTimeDerivative                 := line1[33:43]
	secondTimeDerivative                := line1[44:52]
	bstarBuf                            := line1[53:61]
//	fmt.Println(bstarBuf)
//	fmt.Println(strings.Trim(bstarBuf[0:6], " "))
//	fmt.Println(bstarBuf[6:8])
	bstarBase, _                        := strconv.ParseFloat("0." + strings.Trim(bstarBuf[0:6], " "), 64)
	bstarExp, _                         := strconv.ParseFloat(bstarBuf[6:8], 64)
//	fmt.Printf("%f\n", bstarBase)
//	fmt.Printf("%f\n", bstarExp)
	bstar                               := bstarBase * math.Pow(10.0, bstarExp)
	ephemerisType, _                    := strconv.ParseUint(line1[62:63], 10, 64)
	elementNumber, _                    := strconv.ParseUint(strings.Trim(line1[64:68], " "), 10, 64)
//	checksum, _                         := strconv.Atoi(line1[68:69])

	tle.SatelliteNumber         = satelliteNumber
	tle.Classification          = classification
	tle.InternationalDesignator = strings.Trim(internationalDesignatorYear +
		internationalDesignatorLaunchNumber +
		internationalDesignatorPiece, " ")
	tle.EpochYear               = epochYear
	tle.EpochDay                = epocheDay
	tle.FirstTimeDerivative     = firstTimeDerivative
	tle.SecondTimeDerivative    = secondTimeDerivative
	tle.BSTAR                   = bstar
	tle.EphemerisType           = ephemerisType
	tle.ElementNumber           = elementNumber

	return tle, nil
}


func (tle *TLE) parseLine2(line2 string) (*TLE, error) {
	line2 = strings.TrimRight(line2, "\r\n")
	if len(line2) != 69 && line2[0:1] != "2" {
		return tle, errors.New("Invalid line-2 for TLE.")
	}

	inclination, _       := strconv.ParseFloat(line2[8:16], 64)
	rightAscension, _    := strconv.ParseFloat(line2[17:25], 64)
	eccentricity, _      := strconv.ParseFloat(line2[26:33], 64)
	argumentOfPerigee, _ := strconv.ParseFloat(line2[34:42], 64)
	meanAnomaly, _       := strconv.ParseFloat(line2[43:51], 64)
	meanMotion, _        := strconv.ParseFloat(line2[52:63], 64)
	revolutionNumber, _  := strconv.ParseUint(line2[63:68], 10, 64)
//	checksum, _          := strconv.Atoi(lien2[68:69])

	tle.Inclination       = inclination
	tle.RightAscension    = rightAscension
	tle.Eccentricity      = eccentricity
	tle.ArgumentOfPerigee = argumentOfPerigee
	tle.MeanAnomaly       = meanAnomaly
	tle.MeanMotion        = meanMotion
	tle.RevolutionNumber  = revolutionNumber

	return tle, nil
}
