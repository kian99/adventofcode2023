package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Welcome to day 3")
	fmt.Println("Challenge one =", ChallengeOne("input.txt"))
}

func ChallengeOne(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	return sumEnginePartNumbers(scanner)
}

var reNumberMatch = regexp.MustCompile(`\d+`)

func shiftRows(lines []string, incomingLine string) {
	lines[0] = lines[1]
	lines[1] = lines[2]
	lines[2] = incomingLine
}

func sumEnginePartNumbers(scanner *bufio.Scanner) int {
	total := 0
	lines := make([]string, 3)
	// Buffer the first line in before we start the loop.
	if !scanner.Scan() {
		return 0
	}
	lines[2] = scanner.Text()
	for scanner.Scan() {
		line := scanner.Text()
		shiftRows(lines, line)
		total += sumValidNumbersInRow(lines)
	}
	// Buffer an empty line in so that we process the last line.
	shiftRows(lines, "")
	total += sumValidNumbersInRow(lines)
	return total
}

func sumValidNumbersInRow(lines []string) int {
	total := 0
	for _, loc := range reNumberMatch.FindAllStringIndex(lines[1], -1) {
		index := loc[0]
		length := loc[1] - loc[0]
		match := lines[1][loc[0]:loc[1]]
		if isSymbolAdjacent(lines[0], lines[1], lines[2], index, length) {
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(fmt.Sprintf("unable to convert number %s", match))
			}
			total += num
		}
	}
	return total
}

func isSymbolAdjacent(lineAbove, line, lineBelow string, index, length int) bool {
	return isVerticallyAdjacent(lineAbove, line, index, length) || isAdjacent(line, index, length) || isVerticallyAdjacent(lineBelow, line, index, length)
}

var reSpecialChar = regexp.MustCompile(`[^\d.]`)

func isVerticallyAdjacent(adjLine, line string, index, length int) bool {
	if adjLine == "" {
		return false
	}
	start := index - 1
	end := index + length + 1 // +1 to get the diagonal
	if start < 0 {
		start = 0
	}
	if end > len(adjLine) {
		end = len(line)
	}
	checkSubStr := adjLine[start:end]
	return len(reSpecialChar.FindStringIndex(checkSubStr)) > 0
}

func isAdjacent(line string, index, length int) bool {
	start := index - 1
	end := index + length + 1 // +1 to get the next char
	if start < 0 {
		start = 0
	}
	if end > len(line) {
		end = len(line)
	}
	checkSubStr := line[start:end]
	return len(reSpecialChar.FindStringIndex(checkSubStr)) > 0
}
