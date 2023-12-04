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

func TestSumScratchCards(t *testing.T) {
	input := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\n",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\n",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\n",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\n",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 30
	got := sumScratchCards(s)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestProcessScratchCards(t *testing.T) {
	input := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\n",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\n",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\n",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\n",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 13
	got := sumWinningPoints(s)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}
