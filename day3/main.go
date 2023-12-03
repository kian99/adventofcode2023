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
	fmt.Println("Challenge two =", ChallengeTwo("input.txt"))
}

func ChallengeOne(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	return sumParts(scanner, sumValidNumbersInRow)
}

func ChallengeTwo(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	return sumParts(scanner, sumGearRatios)
}

type lines [3]string

func newLines(topLine, midLine, bottomLine string) lines {
	return lines{topLine, midLine, bottomLine}
}

func shiftRows(lines *lines, incomingLine string) {
	lines[0] = lines[1]
	lines[1] = lines[2]
	lines[2] = incomingLine
}

func sumParts(scanner *bufio.Scanner, f func(lines) int) int {
	total := 0
	lines := lines{}
	// Buffer the first line in before we start the loop.
	if !scanner.Scan() {
		return 0
	}
	lines[2] = scanner.Text()
	for scanner.Scan() {
		line := scanner.Text()
		shiftRows(&lines, line)
		total += f(lines)
	}
	// Buffer an empty line in so that we process the last line.
	shiftRows(&lines, "")
	total += f(lines)
	return total
}

func sumGearRatios(lines lines) int {
	total := 0
	for _, loc := range reGearMatch.FindAllStringIndex(lines[1], -1) {
		index := loc[0]
		if loc[0]+1 != loc[1] {
			panic("invalid * string found")
		}
		length := 1
		if ratio, ok := getGearRatio(lines, index, length); ok {
			total += ratio
		}
	}
	return total
}

func sumValidNumbersInRow(lines lines) int {
	total := 0
	for _, loc := range reNumberMatch.FindAllStringIndex(lines[1], -1) {
		index := loc[0]
		length := loc[1] - loc[0]
		match := lines[1][loc[0]:loc[1]]
		if isSymbolAdjacent(lines, index, length) {
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(fmt.Sprintf("unable to convert number %s", match))
			}
			total += num
		}
	}
	return total
}

type partFinder struct {
	lines   lines
	matcher *regexp.Regexp
	// index and length represent the index and length of the part number we are searching for.
	index  int
	length int
}

func newPartFinder(lines lines, match *regexp.Regexp, index, length int) partFinder {
	return partFinder{
		lines:   lines,
		matcher: match,
		index:   index,
		length:  length,
	}
}

type gearFinder struct {
	line string
	loc  [][]int
}

func (p partFinder) NewGearFinder(linePos int, loc [][]int) gearFinder {
	return gearFinder{line: p.lines[linePos], loc: loc}
}

func getGearRatio(lines lines, index, length int) (int, bool) {
	p := newPartFinder(lines, reDigitMatch, index, length)
	locAbove, _ := p.isVerticallyAdjacent(top)
	g0 := p.NewGearFinder(0, locAbove)
	locSides, _ := p.isAdjacent()
	g1 := p.NewGearFinder(1, locSides)
	locBelow, _ := p.isVerticallyAdjacent(bottom)
	g2 := p.NewGearFinder(2, locBelow)
	count := len(locAbove) + len(locSides) + len(locBelow)
	if count != 2 {
		return 0, false
	}
	total := 1
	gearFinders := []gearFinder{g0, g1, g2}
	for _, gf := range gearFinders {
		for _, gearPart := range gf.getAdjacencyList() {
			total *= gearPart
		}
	}
	return total, true
}

func isSymbolAdjacent(lines lines, index, length int) bool {
	p := newPartFinder(lines, reSpecialChar, index, length)
	_, adjacentAbove := p.isVerticallyAdjacent(top)
	_, adjacent := p.isAdjacent()
	_, adjacentBelow := p.isVerticallyAdjacent(bottom)
	return adjacentAbove || adjacent || adjacentBelow
}

func (g gearFinder) getAdjacencyList() []int {
	adjList := []int{}
	for _, indices := range g.loc {
		if len(indices) != 2 {
			panic("expected two indices")
		}
		adjList = append(adjList, getExpandedInt(g.line, indices[0]))
	}
	return adjList
}

func getExpandedInt(line string, i int) int {
	start := i
	end := i + 1
	// look-ahead
	for end < len(line) && line[end] != '.' && line[end] != '*' {
		end++
	}
	// look-behind
	for start > 0 && line[start-1] != '.' && line[start-1] != '*' {
		start--
	}
	res, err := strconv.Atoi(line[start:end])
	if err != nil {
		panic(err)
	}
	return res
}

// isVerticallyAdjacent determines whether the substring string located at
// p.Index-1 with length p.Length+1 has any matches against the regex
// matcher specified in p.matcher. The param top determines whether the substring
// searched comes from p.lines[0] (top) or p.lines[2] (bottom).
// The returned params are a slice of positions of sucessive matches shifted to
// represent their index in the original line rather than the substring
// and a boolean to indicate whether any matches were found.
func (p partFinder) isVerticallyAdjacent(top bool) ([][]int, bool) {
	adjLine := p.lines[2]
	if top {
		adjLine = p.lines[0]
	}
	if adjLine == "" {
		return nil, false
	}
	return getMatches(adjLine, p.index, p.length, p.matcher)
}

func (p partFinder) isAdjacent() ([][]int, bool) {
	return getMatches(p.lines[1], p.index, p.length, p.matcher)
}

func getMatches(line string, index, length int, matcher *regexp.Regexp) ([][]int, bool) {
	start := index - 1
	end := index + length + 1 // +1 to get the diagonal
	if start < 0 {
		start = 0
	}
	if end > len(line) {
		end = len(line)
	}
	checkSubStr := line[start:end]
	locs := matcher.FindAllStringIndex(checkSubStr, -1)
	setAbsoluteIndices(locs, start)
	return locs, len(locs) > 0
}

func setAbsoluteIndices(locs [][]int, start int) {
	for _, loc := range locs {
		for i := range loc {
			loc[i] = loc[i] + start
		}
	}
}
