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
	return findWinsMultiplied(s)
}

func ChallengeTwo(filename string) uint64 {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	return findSingleWin(s)
}

func findWinsMultiplied(s *bufio.Scanner) uint64 {
	var lines [2]string
	index := 0
	for s.Scan() {
		lines[index] = s.Text()
		index++
	}
	times, distances := extractTimeAndDistance(lines[0], lines[1])
	return multiplyWaysToWin(times, distances)
}

func findSingleWin(s *bufio.Scanner) uint64 {
	var lines [2]string
	index := 0
	for s.Scan() {
		lines[index] = s.Text()
		index++
	}
	times, distances := extractTimeAndDistanceWithBadKerning(lines[0], lines[1])
	return findBigWin(times, distances)
}

func findBigWin(time, distance uint64) uint64 {
	total := uint64(1)
	minHoldTime, ok := determineMinimumHoldTime(time, distance)
	if !ok {
		panic("no min hold time")
	}
	maxHoldTime, ok := determineMaximumTime(time, distance)
	if !ok {
		panic("no max hold time")
	}
	waysToWin := maxHoldTime - minHoldTime + 1
	total *= waysToWin
	return total
}

func multiplyWaysToWin(time, distance []uint64) uint64 {
	total := uint64(1)
	for i := 0; i < len(time); i++ {
		minHoldTime, ok := determineMinimumHoldTime(time[i], distance[i])
		if !ok {
			panic("no min hold time")
		}
		maxHoldTime, ok := determineMaximumTime(time[i], distance[i])
		if !ok {
			panic("no max hold time")
		}
		waysToWin := maxHoldTime - minHoldTime + 1
		total *= waysToWin
	}
	return total
}

func determineMinimumHoldTime(rt, d uint64) (uint64, bool) {
	for hold := uint64(1); hold < rt; hold++ {
		if hold*(rt-hold) > d {
			return hold, true
		}
	}
	return 0, false
}

func determineMaximumTime(rt, d uint64) (uint64, bool) {
	for hold := rt; hold >= 1; hold-- {
		if hold*(rt-hold) > d {
			return hold, true
		}
	}
	return 0, false
}

var digitMatch = regexp.MustCompile(`\d+`)

func extractTimeAndDistance(line1, line2 string) (time, distance []uint64) {
	times := digitMatch.FindAllString(strings.Split(line1, ":")[1], -1)
	distances := digitMatch.FindAllString(strings.Split(line2, ":")[1], -1)
	for _, match := range times {
		t, _ := strconv.Atoi(match)
		time = append(time, uint64(t))
	}
	for _, match := range distances {
		d, _ := strconv.Atoi(match)
		distance = append(distance, uint64(d))
	}
	return
}

func extractTimeAndDistanceWithBadKerning(line1, line2 string) (time, distance uint64) {
	times := digitMatch.FindAllString(strings.Split(line1, ":")[1], -1)
	distances := digitMatch.FindAllString(strings.Split(line2, ":")[1], -1)
	var totalTime string
	for _, match := range times {
		totalTime += match
	}
	time, _ = strconv.ParseUint(totalTime, 10, 64)
	var totalDistance string
	for _, match := range distances {
		totalDistance += match
	}
	distance, _ = strconv.ParseUint(totalDistance, 10, 64)
	return
}
