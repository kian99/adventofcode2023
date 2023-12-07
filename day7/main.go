package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to day 5")
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
	return getTotalWinnings(s)
}

var cards = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
var cardsWithJoker = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

var cardStrength = map[rune]int{}
var cardStrengthWithJoker = map[rune]int{}

func init() {
	for i := 0; i < len(cards); i++ {
		cardStrengthWithJoker[cardsWithJoker[i]] = len(cards) - i
		cardStrength[cards[i]] = len(cards) - i
	}
}

type hand struct {
	hand     string
	bid      int
	handType int
}

func ChallengeTwo(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	return getTotalWinningsWithJokers(s)
}

const (
	highCard = iota
	onePair
	twoPair
	threeKind
	fullHouse
	fourKind
	fiveKind
)

func getTotalWinnings(s *bufio.Scanner) int {
	hands := []hand{}
	for s.Scan() {
		line := s.Text()
		hands = append(hands, makeHand(line))
	}
	const useJokers = false
	return scoreHands(hands, useJokers, cardStrength)
}

func getTotalWinningsWithJokers(s *bufio.Scanner) int {
	hands := []hand{}
	for s.Scan() {
		line := s.Text()
		hands = append(hands, makeHand(line))
	}
	const useJokers = true
	return scoreHands(hands, useJokers, cardStrengthWithJoker)
}

func makeHand(line string) hand {
	m := strings.Split(line, " ")
	h := hand{}
	h.hand = m[0]
	h.bid, _ = strconv.Atoi(m[1])
	return h
}

func scoreHands(hands []hand, useJokers bool, cardStrengths map[rune]int) int {
	for i := range hands {
		if useJokers {
			hands[i].computeHandReplacingJokers()
		} else {
			hands[i].computeHandType()
		}
	}
	sortHands(hands, useJokers, cardStrengths)
	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}
	return total
}

func sortHands(hands []hand, useJokers bool, cardStrengths map[rune]int) {
	f := func(h1, h2 hand) int {
		return sortHand(h1, h2, cardStrengths)
	}
	slices.SortFunc(hands, f)
}

func sortHand(h1, h2 hand, cardStrengths map[rune]int) int {
	if h1.handType < h2.handType {
		return -1
	} else if h1.handType > h2.handType {
		return 1
	}
	for i := range h1.hand {
		if cardStrengths[rune(h1.hand[i])] < cardStrengths[rune(h2.hand[i])] {
			return -1
		} else if cardStrengths[rune(h1.hand[i])] > cardStrengths[rune(h2.hand[i])] {
			return 1
		}
	}
	return 0

}

func (h *hand) computeHandReplacingJokers() {
	piles, _, largestPileCard := h.computePiles()
	secondLargest := 0
	secondLargestPileCard := rune(h.hand[0])
	for card, pile := range piles {
		if pile > secondLargest && card != largestPileCard {
			secondLargest = pile
			secondLargestPileCard = card
		}
	}
	numPiles := len(piles)
	originalHand := h.hand
	if numPiles != 1 {
		if largestPileCard == 'J' {
			h.hand = strings.ReplaceAll(h.hand, "J", string(secondLargestPileCard))
		} else {
			h.hand = strings.ReplaceAll(h.hand, "J", string(largestPileCard))
		}
	}
	h.computeHandType()
	h.hand = originalHand
}

func (h hand) computePiles() (piles map[rune]int, largestPile int, largestPileCard rune) {
	piles = map[rune]int{}
	for _, card := range h.hand {
		piles[card]++
	}
	largestPile = 1
	largestPileCard = rune(h.hand[0])
	for card, pile := range piles {
		if pile > largestPile {
			largestPile = pile
			largestPileCard = card
		}
	}
	return
}

func (h *hand) computeHandType() {
	piles, largestPile, _ := h.computePiles()
	numPiles := len(piles)
	switch numPiles {
	case 1:
		h.handType = fiveKind
	case 2:
		if largestPile == 4 {
			h.handType = fourKind
		} else {
			h.handType = fullHouse
		}
	case 3:
		if largestPile == 3 {
			h.handType = threeKind
		} else {
			h.handType = twoPair
		}
	case 4:
		h.handType = onePair
	case 5:
		h.handType = highCard
	default:
		panic("invalid hand")
	}

}
