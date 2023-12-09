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

func TestExtrapolateBackwards(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15\n",
		"1 3 6 10 15 21\n",
		"10 13 16 21 30 45\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 2
	got := sumExtrapolatedValues(s, backwards)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestExtrapolateForwards(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15\n",
		"1 3 6 10 15 21\n",
		"10 13 16 21 30 45\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 114
	got := sumExtrapolatedValues(s, forwards)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}
