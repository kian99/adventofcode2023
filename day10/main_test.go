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

func TestCountEnclosedMedium(t *testing.T) {
	input := []string{
		"...........\n",
		".S-------7.\n",
		".|F-----7|.\n",
		".||.....||.\n",
		".||.....||.\n",
		".|L-7.F-J|.\n",
		".|..|.|..|.\n",
		".L--J.L--J.\n",
		"...........\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	startCharType = inLoopFromBottom
	expected := 4
	got := countEnclosed(s)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestCountEnclosedEasy(t *testing.T) {
	input := []string{
		".F----7F7F7F7F-7....\n",
		".|F--7||||||||FJ....\n",
		".||.FJ||||||||L7....\n",
		"FJL7L7LJLJ||LJ.L-7..\n",
		"L--J.L7...LJS7F-7L7.\n",
		"....F-J..F7FJ|L7L7L7\n",
		"....L7.F7||L7|.L7L7|\n",
		".....|FJLJ|FJ|F7|.LJ\n",
		"....FJL-7.||.||||...\n",
		"....L---J.LJ.LJLJ...\n",
	}
	mockFile := mockFile{data: input}
	startCharType = inLoopFromBottom
	s := bufio.NewScanner(&mockFile)
	expected := 8
	got := countEnclosed(s)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestCountSteps(t *testing.T) {
	input := []string{
		"7-F7-\n",
		".FJ|7\n",
		"SJLL7\n",
		"|F--J\n",
		"LJ.LJ\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 8
	got := countSteps(s)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}
