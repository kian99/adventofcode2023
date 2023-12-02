package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type gamebag map[string]int

func main() {
	fmt.Println("Welcome to day 2")
	availableCubes := gamebag{"red": 12, "blue": 14, "green": 13}
	fmt.Println("Challenge 1 =", challengeOne("input.txt", availableCubes))
	fmt.Println("Challenge 2 =", challengeTwo("input.txt"))
}

func challengeOne(filename string, availableCubes gamebag) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	return GetValidGameTotal(scanner, availableCubes)
}

func challengeTwo(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	return GetSumOfPowers(scanner)
}

func GetSumOfPowers(scanner *bufio.Scanner) int {
	maxCountsPerGame := getMaxGameCounts(scanner)
	sumCubePower := 0
	for _, counts := range maxCountsPerGame {
		cubePower := 0
		for _, val := range counts {
			if cubePower == 0 {
				cubePower = val
			} else {
				cubePower *= val
			}
		}
		sumCubePower += cubePower
	}
	return sumCubePower

}

func GetValidGameTotal(scanner *bufio.Scanner, availableCubes gamebag) int {
	maxCountsPerGame := getMaxGameCounts(scanner)
	validGames := []int{}
	for i, game := range maxCountsPerGame {
		validGame := verifyValidGame(game, availableCubes)
		if validGame {
			validGames = append(validGames, i+1)
		}
	}
	var total int
	for _, val := range validGames {
		total += val
	}
	return total
}

func getMaxGameCounts(scanner *bufio.Scanner) []gamebag {
	maxCountsPerGame := []gamebag{}
	for scanner.Scan() {
		line := scanner.Text()
		// Sample line `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`
		colonIndex := strings.Index(line, ":")
		countSection := line[colonIndex+1:]
		counts := extractCounts(countSection)
		maxCountsPerGame = append(maxCountsPerGame, counts)
	}
	return maxCountsPerGame
}

func verifyValidGame(game gamebag, availableCubes gamebag) bool {
	validGame := true
	for colour, availableCount := range availableCubes {
		val, ok := game[colour]
		if !ok {
			panic("invalid colour found")
		}
		if val > availableCount {
			validGame = false
			break
		}
	}
	return validGame
}

var colours = []string{"red", "green", "blue"}

func extractCounts(line string) gamebag {
	colourCounts := gamebag{}
	reveals := strings.Split(line, ";")
	for _, reveal := range reveals {
		counts := strings.Split(reveal, ",")
		for _, count := range counts {
			var colourRevealed string
			var colourCount int
			for _, colour := range colours {
				if strings.Contains(count, colour) {
					colourRevealed = colour
					spaceIndex := strings.LastIndex(count, " ")
					if spaceIndex == -1 {
						panic("failed to find space in count")
					}
					var err error
					cleanCount := strings.ReplaceAll(count[0:spaceIndex], " ", "")
					colourCount, err = strconv.Atoi(cleanCount)
					if err != nil {
						panic(err)
					}
					break
				}
			}
			val, ok := colourCounts[colourRevealed]
			if !ok || colourCount > val {
				colourCounts[colourRevealed] = colourCount
			}
		}
	}
	return colourCounts
}
