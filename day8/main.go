package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Welcome to day 8")
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
	return getSteps(s, false)
}

func ChallengeTwo(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	return getSteps(s, true)
}

type choices struct {
	left  string
	right string
}

type network struct {
	in      string
	choices map[string]*choices
}

type ghost struct {
	n         network
	positions []string
}

func getSteps(s *bufio.Scanner, isGhost bool) int {
	net := network{}
	net.choices = map[string]*choices{}
	s.Scan()
	net.in = s.Text()
	s.Scan()
	for s.Scan() {
		net.addRoute(s.Text())
	}
	if isGhost {
		ghost := ghost{n: net}
		return ghost.ghostSteps()
	}
	_, count := net.lengthToDestination("AAA", regexp.MustCompile(`ZZZ`))
	return count
}

var re = regexp.MustCompile(`[A-Z1-9]+`)
var endsInZMatch = regexp.MustCompile(`Z$`)

func (n *network) addRoute(line string) {
	match := re.FindAllString(line, -1)
	if len(match) != 3 {
		panic("invalid line")
	}
	n.choices[match[0]] = &choices{left: match[1], right: match[2]}
}

func (n network) lengthToDestination(start string, match *regexp.Regexp) (string, int) {
	count := 0
	pos := start
	for {
		for _, char := range n.in {
			if char == 'L' {
				pos = n.choices[pos].left
			} else {
				pos = n.choices[pos].right
			}
			count++
			if match.FindString(pos) != "" {
				return pos, count
			}
		}
	}
}

func (g ghost) ghostSteps() int {
	for pos := range g.n.choices {
		if pos[2] == 'A' {
			g.positions = append(g.positions, pos)
		}
	}
	initialSteps := []int{}
	for i := range g.positions {
		dest, initalStep := g.n.lengthToDestination(g.positions[i], endsInZMatch)
		fmt.Println("cycle count for", g.positions[i], "ends in", dest, "after", initalStep, "steps")
		initialSteps = append(initialSteps, initalStep)
	}
	if len(initialSteps) < 2 {
		panic("need 2 cycles")
	}
	return LCM(initialSteps[0], initialSteps[1], initialSteps[2:]...)
}

// From https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
