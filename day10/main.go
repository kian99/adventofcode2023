package main

import (
	"bufio"
	"fmt"
	"os"
)

const printMap = true

var startCharType = inLoopFromTop

func main() {
	fmt.Println("Welcome to day 10")
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
	return countSteps(s)
}

func ChallengeTwo(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	return countEnclosed(s)
}

type tile rune

const (
	ground     tile = '.'
	start      tile = 'S'
	vertical   tile = '|'
	horizontal tile = '-'
	NEBend     tile = 'L'
	NWBend     tile = 'J'
	SWBend     tile = '7'
	SEBend     tile = 'F'

	inLoopVertical   tile = 'V'
	inLoopHorizontal tile = 'H'
	inLoopFromTop    tile = 'T'
	inLoopFromBottom tile = 'B'
)

var pipeMap = map[tile]map[direction]coord{}

// maps current tile to valid next tiles going in the provided direction
var pipeFixtures = map[tile]map[direction][]tile{}

func init() {
	pipeMap[vertical] = map[direction]coord{up: {x: 0, y: -2}, down: {x: 0, y: 2}}
	pipeMap[horizontal] = map[direction]coord{right: {x: 2, y: 0}, left: {x: -2, y: 0}}
	pipeMap[NEBend] = map[direction]coord{left: {x: -1, y: -1}, down: {x: 1, y: 1}}
	pipeMap[NWBend] = map[direction]coord{right: {x: 1, y: -1}, down: {x: -1, y: 1}}
	pipeMap[SWBend] = map[direction]coord{right: {x: 1, y: 1}, up: {x: -1, y: -1}}
	pipeMap[SEBend] = map[direction]coord{up: {x: 1, y: -1}, left: {x: -1, y: 1}}

	upSet := []tile{SWBend, SEBend, vertical}
	downSet := []tile{NWBend, NEBend, vertical}
	rightSet := []tile{NWBend, SWBend, horizontal}
	leftSet := []tile{NEBend, SEBend, horizontal}
	pipeFixtures[vertical] = map[direction][]tile{up: upSet, down: downSet}
	pipeFixtures[horizontal] = map[direction][]tile{right: rightSet, left: leftSet}
	pipeFixtures[NEBend] = map[direction][]tile{up: upSet, right: rightSet}
	pipeFixtures[NWBend] = map[direction][]tile{up: upSet, left: leftSet}
	pipeFixtures[SWBend] = map[direction][]tile{down: downSet, left: leftSet}
	pipeFixtures[SEBend] = map[direction][]tile{down: downSet, right: rightSet}
}

type coord struct {
	x int
	y int
}

func (c coord) inBounds(g grid) bool {
	gLen := len(g)
	if gLen == 0 {
		return false
	}
	if c.y >= 0 && c.y < gLen {
		if c.x >= 0 && c.x < len(g[0]) {
			return true
		}
	}
	return false
}

func (t tile) getPipeOutput(c coord, d direction) coord {
	if t == ground || t == start {
		panic("tile is not pipe")
	}
	shift := pipeMap[t][d]
	c.x = c.x + shift.x
	c.y = c.y + shift.y
	return c
}

func walkPipe(pipe tile, curr tile, c coord, d direction) (coord, bool) {
	if pipe == start || pipe == ground {
		return coord{}, false
	}
	if pipesCanConnect(curr, pipe, d) {
		return pipe.getPipeOutput(c, d), true
	}
	return coord{}, false
}

func pipesCanConnect(curr, next tile, d direction) bool {
	checker := func(validPipes []tile) bool {
		for _, pipe := range validPipes {
			if next == pipe {
				return true
			}
		}
		return false
	}
	if curr == start {
		res := false
		for _, val := range pipeFixtures {
			res = checker(val[d])
			if res {
				return true
			}
		}
		return false
	}
	return checker(pipeFixtures[curr][d])
}

type direction int

const (
	up direction = iota
	right
	down
	left
	total
)

type grid [][]tile

func countSteps(s *bufio.Scanner) int {
	grid := grid{}
	for s.Scan() {
		line := []tile{}
		for _, val := range s.Text() {
			line = append(line, tile(val))
		}
		grid = append(grid, line)
	}
	grid.printGrid(nil)
	return len(grid.buildLoop()) / 2
}

func countEnclosed(s *bufio.Scanner) int {
	grid := grid{}
	for s.Scan() {
		line := []tile{}
		for _, val := range s.Text() {
			line = append(line, tile(val))
		}
		grid = append(grid, line)
	}
	grid.printGrid(nil)
	return grid.countEnclosedTiles()
}

func (g grid) printGrid(current *coord) {
	if !printMap {
		return
	}
	for i, x := range g {
		for j, y := range x {
			if current != nil && current.y == i && current.x == j {
				fmt.Printf("X")
			} else {
				fmt.Printf("%s", string(y))
			}
		}
		fmt.Printf("\n")
	}
}

func (g grid) buildLoop() []coord {
	start, ok := g.getStartPositon()
	if !ok {
		panic("no start position found")
	}
	loop := []coord{}
	next := start
	pipe := start
	fmt.Println("start=", start)
	loop = append(loop, start)
	for {
		// g.printGrid(&next)
		next, pipe = g.step(next, pipe)
		loop = append(loop, pipe)
		// fmt.Println("walked through", pipe)
		if next == start {
			fmt.Println("completed loop")
			break
		}
		loop = append(loop, next)
		// fmt.Println("landed on", next)
	}
	return loop
}

func (g grid) step(curr, prevPipe coord) (out, pipe coord) {
	for i := 0; i < int(total); i++ {
		next := curr
		d := direction(i)
		switch d {
		case up:
			next.y = next.y - 1
		case right:
			next.x = next.x + 1
		case down:
			next.y = next.y + 1
		case left:
			next.x = next.x - 1
		}
		if !next.inBounds(g) || next == prevPipe {
			continue
		}
		pipe := g[next.y][next.x]
		currTile := g[curr.y][curr.x]
		if out, ok := walkPipe(pipe, currTile, curr, d); ok {
			return out, next
		}
	}
	panic("no path found")
}

func (g grid) getStartPositon() (coord, bool) {
	for y, row := range g {
		for x, tile := range row {
			if tile == start {
				return coord{x: x, y: y}, true
			}
		}
	}
	return coord{}, false
}

func (g grid) countEnclosedTiles() int {
	loop := g.buildLoop()
	g.fillMap(loop)
	count := 0
	for y, line := range g {
		for x := range line {
			c := coord{x, y}
			if g[y][x] == inLoopVertical || g[y][x] == inLoopHorizontal || g[y][x] == inLoopFromTop || g[y][x] == inLoopFromBottom {
				continue
			}
			if g.isEnclosed(c) {
				count++
				g[y][x] = 'I'
			}
		}
	}
	g.printGrid(nil)
	return count
}

func (g *grid) fillMap(loop []coord) {
	for _, coord := range loop {
		tile := (*g)[coord.y][coord.x]
		if tile == vertical {
			(*g)[coord.y][coord.x] = inLoopVertical
		} else if tile == horizontal {
			(*g)[coord.y][coord.x] = inLoopHorizontal
		} else if tile == NEBend || tile == NWBend {
			(*g)[coord.y][coord.x] = inLoopFromTop
		} else if tile == SEBend || tile == SWBend {
			(*g)[coord.y][coord.x] = inLoopFromBottom
		} else if tile == start {
			// This must be changed to match the input
			(*g)[coord.y][coord.x] = startCharType
		}
	}
}

func (g grid) isEnclosed(c coord) bool {
	intersectCount := 0
	changeDirectionFromBottom := 0
	changeDirectionFromTop := 0
	for x := 0; x < c.x; x++ {
		if g[c.y][x] == inLoopVertical {
			intersectCount++
		}
		if g[c.y][x] == inLoopFromBottom {
			changeDirectionFromBottom++
		} else if g[c.y][x] == inLoopFromTop {
			changeDirectionFromTop++
		}
	}
	for {
		if changeDirectionFromBottom > 0 && changeDirectionFromTop > 0 {
			intersectCount++
			changeDirectionFromBottom--
			changeDirectionFromTop--
		} else {
			break
		}
	}
	intersectCount += changeDirectionFromBottom + changeDirectionFromTop
	return intersectCount%2 != 0
}
