package main

import (
	"bufio"
	"fmt"
	"os"
)

const galaxy = '#'

func main() {
	fmt.Println("Welcome to day 11")
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
	return sumShortestPaths(s, 2)
}

func ChallengeTwo(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	return sumShortestPaths(s, 1000000)
}

type universe struct {
	image           []string
	emptyRows       map[int]struct{}
	emptyColumns    map[int]struct{}
	expansionFactor int
}

type coord struct {
	x int
	y int
}

func sumShortestPaths(s *bufio.Scanner, expansionFactor int) int {
	u := newUniverse(expansionFactor)
	for s.Scan() {
		line := s.Text()
		u.image = append(u.image, line)
	}
	u.fillEmpty()
	return u.sumPathLengths()
}

func newUniverse(expansionFactor int) universe {
	u := universe{}
	u.emptyColumns = map[int]struct{}{}
	u.emptyRows = map[int]struct{}{}
	u.expansionFactor = expansionFactor
	return u
}

func (u universe) getGalaxyCoords() []coord {
	rows := len(u.image)
	columns := len(u.image[0])
	galaxies := []coord{}
	for j := 0; j < rows; j++ {
		for i := 0; i < columns; i++ {
			if u.image[j][i] == galaxy {
				galaxies = append(galaxies, coord{x: i, y: j})
			}
		}
	}
	return galaxies
}

func (u universe) getUniversePathLength(start, end coord) int {
	xDiff := end.x - start.x
	if xDiff < 0 {
		xDiff *= -1
	}
	for i := min(start.x, end.x); i < max(end.x, start.x); i++ {
		if _, ok := u.emptyColumns[i]; ok {
			xDiff += u.expansionFactor - 1
		}
	}
	yDiff := end.y - start.y
	if yDiff < 0 {
		yDiff *= -1
	}
	for i := min(end.y, start.y); i < max(start.y, end.y); i++ {
		if _, ok := u.emptyRows[i]; ok {
			yDiff += u.expansionFactor - 1
		}
	}
	return xDiff + yDiff
}

func (u universe) sumPathLengths() int {
	galaxies := u.getGalaxyCoords()
	pathLengths := make([]int, len(galaxies))
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pathLengths[i] += u.getUniversePathLength(galaxies[i], galaxies[j])
		}
	}
	total := 0
	for _, val := range pathLengths {
		total += val
	}
	return total
}

func (u *universe) fillEmpty() {
	rows := len(u.image)
	columns := len(u.image[0])
	for j := 0; j < rows; j++ {
		emptyRow := true
		for i := 0; i < columns; i++ {
			if u.image[j][i] == galaxy {
				emptyRow = false
			}
		}
		if emptyRow {
			u.emptyRows[j] = struct{}{}
		}
	}
	for j := 0; j < columns; j++ {
		emptyCol := true
		for i := 0; i < rows; i++ {
			if u.image[i][j] == galaxy {
				emptyCol = false
			}
		}
		if emptyCol {
			u.emptyColumns[j] = struct{}{}
		}
	}
}
