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
	fmt.Println("Welcome to day 4")
	fmt.Println("Challenge one =", ChallengeOne("input.txt"))
	fmt.Println("Challenge two =", ChallengeTwo("input.txt"))
}

func ChallengeOne(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	return sumWinningPoints(s)
}

func ChallengeTwo(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	return sumScratchCards(s)
}

func sumScratchCards(s *bufio.Scanner) int {
	memo := make(winningSet, 1)
	reprocessQueue := []int{}
	index := 0
	for s.Scan() {
		line := s.Text()
		var scratchCardsWon int
		memo, scratchCardsWon = processScratchCards(line, memo)
		for i := 0; i < scratchCardsWon; i++ {
			reprocessQueue = append(reprocessQueue, i+index+1)
		}
		index++
	}
	index += processExtraCards(reprocessQueue, memo)
	return index
}

type winningSet []*int

func grow(slice winningSet, newCap int) winningSet {
	if cap(slice) >= newCap {
		return slice
	}
	newSlice := make(winningSet, newCap)
	copy(newSlice, slice)
	return newSlice
}

func processExtraCards(queue []int, memo winningSet) int {
	index := 0
	for index != len(queue) {
		gameID := queue[index]
		cardsWon := *memo[gameID]
		for i := 0; i < cardsWon; i++ {
			queue = append(queue, gameID+i+1)
		}
		index++
	}
	return index
}

func processScratchCards(line string, memo winningSet) (winningSet, int) {
	gameID := extractGameID(line)
	if len(memo) <= gameID {
		memo = grow(memo, len(memo)+50)
	}
	if memo[gameID] == nil {
		winningNumbers, obtainedNumbers := extractScratchCardNumbers(line)
		_, count := getWinningNumbers(winningNumbers, obtainedNumbers)
		memo[gameID] = &count
	}
	return memo, *memo[gameID]
}

func sumWinningPoints(s *bufio.Scanner) int {
	total := 0
	for s.Scan() {
		line := s.Text()
		total += sumGamePoints(line)
	}
	return total
}

var reDigitMatch = regexp.MustCompile(`(?m)\d+`)

func sumGamePoints(line string) int {
	winningNumbers, obtainedNumbers := extractScratchCardNumbers(line)
	total := 0
	_, winCount := getWinningNumbers(winningNumbers, obtainedNumbers)
	for i := 0; i < winCount; i++ {
		if total == 0 {
			total++
		} else {
			total *= 2
		}
	}
	return total
}

func extractGameID(line string) int {
	matches := strings.Split(line, "|")
	if len(matches) != 2 {
		panic("no | separated found")
	}
	left := matches[0]
	gameID, err := strconv.Atoi(reDigitMatch.FindString(left[:strings.Index(left, ":")]))
	if err != nil {
		panic(err)
	}
	return gameID - 1
}

func extractScratchCardNumbers(line string) (want, got map[int]struct{}) {
	matches := strings.Split(line, "|")
	if len(matches) != 2 {
		panic("no | separated found")
	}
	left := matches[0]
	left = left[strings.Index(left, ":"):]
	right := matches[1]
	want = numbersToMap(left)
	got = numbersToMap(right)
	return
}

func getWinningNumbers(want, got map[int]struct{}) ([]int, int) {
	winningNums := []int{}
	for num := range got {
		if _, ok := want[num]; ok {
			winningNums = append(winningNums, num)
		}
	}
	return winningNums, len(winningNums)
}

func numbersToMap(line string) map[int]struct{} {
	matches := reDigitMatch.FindAllString(line, -1)
	res := map[int]struct{}{}
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			panic(err)
		}
		res[num] = struct{}{}
	}
	return res
}
