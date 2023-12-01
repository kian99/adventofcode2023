package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Welcome to day 1")
	fmt.Println("Challenge 1 =", challengeOne())
	fmt.Println("Challenge 2 =", challengeTwo())
}

func challengeOne() int {
	f, err := os.Open("./challengeOne.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	return CalibrateValues(scanner)
}

func CalibrateValues(scanner *bufio.Scanner) int {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		var firstNumber rune
		var secondNumber rune
		for _, char := range line {
			if ok := unicode.IsDigit(char); ok {
				if firstNumber == 0 {
					firstNumber = char
				}
				secondNumber = char
			}
		}
		combinedChars := string(firstNumber) + string(secondNumber)
		numResult, err := strconv.Atoi(combinedChars)
		check(err)
		total += numResult
	}
	return total
}

func challengeTwo() int {
	f, err := os.Open("./challengeTwo.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	return CalibrateValuesWithStringNumbers(scanner)
}

func CalibrateValuesWithStringNumbers(scanner *bufio.Scanner) int {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		var firstNumber rune
		var secondNumber rune
		for i := range line {
			if num, ok := isDigitOrSpellsDigit(i, line); ok {
				if firstNumber == 0 {
					firstNumber = num
				}
				secondNumber = num
			}
		}
		combinedChars := string(firstNumber) + string(secondNumber)
		numResult, err := strconv.Atoi(combinedChars)
		check(err)
		total += numResult
	}
	return total
}

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func isDigitOrSpellsDigit(i int, line string) (rune, bool) {
	if ok := unicode.IsDigit(rune(line[i])); ok {
		return rune(line[i]), true
	}
	for index, number := range numbers {
		end := i + len(number)
		if end < i || end > len(line) {
			continue
		}
		temp := line[i:end]
		if temp == number {
			stringDigit := fmt.Sprintf("%d", index+1)[0]
			return rune(stringDigit), true
		}
	}
	return 0, false

}
