package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to day 5")
	fmt.Println("Challenge one =", ChallengeOne("input.txt"))
	fmt.Println("Challenge two =", ChallengeTwo("input.txt"))
}

func ChallengeOne(filename string) uint64 {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	return findLowestLocation(s)
}

func ChallengeTwo(filename string) uint64 {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	return findLowestLocationPartTwo(s)
}

type mapType int

const (
	seedsToSoil mapType = iota
	soilToFertilizer
	fertilizerToWater
	waterToLight
	lightToTemperature
	temeratureToHumity
	humidityToLocation
	totalMapTypes
)

type rangeMap []rangeTuple
type seedPair struct {
	start  uint64
	length uint64
}

type almanac struct {
	seedPairs []seedPair
	seeds     []uint64
	maps      [totalMapTypes]rangeMap
}

type rangeTuple struct {
	source      uint64
	destination uint64
	length      uint64
}

func findLowestLocationPartTwo(s *bufio.Scanner) uint64 {
	var mapIndex uint64
	extractedSeeds := false
	var almanac almanac
	mapBuffer := []string{}
	for s.Scan() {
		line := s.Text()
		if !extractedSeeds {
			almanac.seedPairs = extractSeedPairs(line)
			extractedSeeds = true
			s.Scan() //skip the next empty line
		} else {
			if !strings.Contains(line, ":") && line != "" {
				mapBuffer = append(mapBuffer, line)
			} else if line == "" {
				almanac.populateMap(mapIndex, mapBuffer)
				mapIndex++
				mapBuffer = []string{}
			}
		}
	}
	almanac.populateMap(mapIndex, mapBuffer)
	return almanac.findLowestLocationFromSeedPairs()
}

func findLowestLocation(s *bufio.Scanner) uint64 {
	var mapIndex uint64
	extractedSeeds := false
	var almanac almanac
	mapBuffer := []string{}
	for s.Scan() {
		line := s.Text()
		if !extractedSeeds {
			almanac.seeds = extractSeeds(line)
			extractedSeeds = true
			s.Scan() //skip the next empty line
		} else {
			if !strings.Contains(line, ":") && line != "" {
				mapBuffer = append(mapBuffer, line)
			} else if line == "" {
				almanac.populateMap(mapIndex, mapBuffer)
				mapIndex++
				mapBuffer = []string{}
			}
		}
	}
	almanac.populateMap(mapIndex, mapBuffer)
	return almanac.findLowestLocationFromSeed()
}

var seedPairsMatch = regexp.MustCompile(`\d+ \d+`)
var digitMatch = regexp.MustCompile(`\d+`)

func extractSeedPairs(line string) []seedPair {
	seedPairs := []seedPair{}
	for _, match := range seedPairsMatch.FindAllString(line, -1) {
		splits := strings.Split(match, " ")
		if len(splits) != 2 {
			panic("expected 2 matches")
		}
		pair := seedPair{}
		var err error
		pair.start, err = strconv.ParseUint(splits[0], 10, 32)
		if err != nil {
			panic(err)
		}
		pair.length, err = strconv.ParseUint(splits[1], 10, 32)
		if err != nil {
			panic(err)
		}
		seedPairs = append(seedPairs, pair)
	}
	return seedPairs
}

func extractSeeds(line string) []uint64 {
	seeds := []uint64{}
	for _, match := range digitMatch.FindAllString(line, -1) {
		seed, err := strconv.ParseUint(match, 10, 32)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, seed)
	}
	return seeds
}

func (a almanac) findLowestLocationFromSeedPairs() uint64 {
	var lowest uint64
	found := false
	for _, seedPair := range a.seedPairs {
		for i := uint64(0); i < seedPair.length; i++ {
			seed := seedPair.start + i
			val := a.traverseMaps(seed)
			if !found || val < lowest {
				lowest = val
				found = true
			}
		}
	}
	return lowest
}

func (a almanac) findLowestLocationFromSeed() uint64 {
	var lowest uint64
	found := false
	for _, seed := range a.seeds {
		val := a.traverseMaps(seed)
		if !found || val < lowest {
			lowest = val
			found = true
		}
	}
	return lowest
}

func (a *almanac) populateMap(mapIndex uint64, buffer []string) {
	allRanges := []rangeTuple{}
	for _, line := range buffer {
		rangeMap := rangeTuple{}
		matches := strings.Split(line, " ")
		if len(matches) != 3 {
			panic("expected 3 values from map line")
		}
		var err error
		rangeMap.destination, err = strconv.ParseUint(matches[0], 10, 32)
		if err != nil {
			panic(err)
		}
		rangeMap.source, err = strconv.ParseUint(matches[1], 10, 32)
		if err != nil {
			panic(err)
		}
		rangeMap.length, err = strconv.ParseUint(matches[2], 10, 32)
		if err != nil {
			panic(err)
		}
		allRanges = append(allRanges, rangeMap)
	}
	a.maps[mapIndex] = allRanges
}

func (a almanac) traverseMaps(seed uint64) uint64 {
	for i := 0; i < int(totalMapTypes); i++ {
		seed = a.maps[i].mapSeed(seed)
	}
	return seed
}

func (r rangeMap) mapSeed(seed uint64) (dest uint64) {
	dest = seed
	for _, tuple := range r {
		if val, ok := tuple.inMap(seed); ok {
			dest = val
			return
		}
	}
	return
}

func (r rangeTuple) inMap(seed uint64) (dest uint64, found bool) {
	if seed >= r.source && seed < r.source+r.length {
		dest = r.destination + (seed - r.source)
		found = true
	}
	return
}
