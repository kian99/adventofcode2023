package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Welcome to day 9")
	fmt.Println("Challenge one =", ChallengeOne("input.txt"))
	fmt.Println("Challenge two =", ChallengeTwo("input.txt"))
}

const forwards = true
const backwards = false

func ChallengeOne(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	return sumExtrapolatedValues(s, forwards)
}

func ChallengeTwo(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	return sumExtrapolatedValues(s, backwards)
}

type history []int

func sumExtrapolatedValues(s *bufio.Scanner, direction bool) int {
	historicalReadings := []history{}
	for s.Scan() {
		historicalReadings = append(historicalReadings, makeHistory(s.Text()))
	}
	sum := 0
	for _, readings := range historicalReadings {
		sum += readings.getPredictedValue(direction)
	}
	return sum
}

var re = regexp.MustCompile(`[-\d]+`)

func makeHistory(line string) (res history) {
	readings := re.FindAllString(line, -1)
	for _, reading := range readings {
		num, err := strconv.Atoi(reading)
		if err != nil {
			panic(err)
		}
		res = append(res, num)
	}
	return
}

func (h history) onlyZeros() bool {
	for _, val := range h {
		if val != 0 {
			return false
		}
	}
	return true
}

func (h history) getPredictedValue(direction bool) int {
	if h.onlyZeros() {
		return 0
	}
	diffHist := h.getDifferences()
	if direction {
		lastItem := h[len(h)-1]
		return diffHist.getPredictedValue(direction) + lastItem
	} else {
		firstItem := h[0]
		return firstItem - diffHist.getPredictedValue(direction)
	}

}

func (h history) getDifferences() (res history) {
	for i := 0; i < len(h)-1; i++ {
		res = append(res, h[i+1]-h[i])
	}
	return
}
