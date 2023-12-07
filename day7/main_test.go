package main

import (
	"bufio"
	"io"
	"testing"
)

type mockFile struct {
	data      []string
	lineCount int
}

func (m *mockFile) Read(p []byte) (n int, err error) {
	if m.lineCount < len(m.data) {
		n = copy(p, m.data[m.lineCount])
	} else {
		return 0, io.EOF
	}
	m.lineCount++
	return n, nil
}

func TestGetTotalWinningsWithJokers(t *testing.T) {
	input := []string{
		"32T3K 765\n",
		"T55J5 684\n",
		"KK677 28\n",
		"KTJJT 220\n",
		"QQQJA 483\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 5905
	got := getTotalWinningsWithJokers(s)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestGetTotalWinnings(t *testing.T) {
	input := []string{
		"32T3K 765\n",
		"T55J5 684\n",
		"KK677 28\n",
		"KTJJT 220\n",
		"QQQJA 483\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 6440
	got := getTotalWinnings(s)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestSortHandsWithJokers(t *testing.T) {
	hands := []hand{{
		hand: "KKJ12",
	}, {
		hand: "KK245",
	}, {
		hand: "J1345",
	},
	}
	sortHands(hands, true, cardStrengthWithJoker)
	order := []string{"J1345", "KKJ12", "KK245"}
	for i, handWant := range order {
		if handWant != hands[i].hand {
			t.Logf("Expected %s, got %s", handWant, hands[i].hand)
			t.Fail()
		}
	}

}
