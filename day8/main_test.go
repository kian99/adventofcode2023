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

func TestGhostSteps(t *testing.T) {
	input := []string{
		"LR\n",
		"\n",
		"11A = (11B, XXX)\n",
		"11B = (XXX, 11Z)\n",
		"11Z = (11B, XXX)\n",
		"22A = (22B, XXX)\n",
		"22B = (22C, 22C)\n",
		"22C = (22Z, 22Z)\n",
		"22Z = (22B, 22B)\n",
		"XXX = (XXX, XXX)\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 6
	got := getSteps(s, true)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestGetSteps2(t *testing.T) {
	input := []string{
		"RL\n",
		"\n",
		"AAA = (BBB, CCC)\n",
		"BBB = (DDD, EEE)\n",
		"CCC = (ZZZ, GGG)\n",
		"DDD = (DDD, DDD)\n",
		"EEE = (EEE, EEE)\n",
		"GGG = (GGG, GGG)\n",
		"ZZZ = (ZZZ, ZZZ)\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 2
	got := getSteps(s, false)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestGetSteps(t *testing.T) {
	input := []string{
		"LLR\n",
		"\n",
		"AAA = (BBB, BBB)\n",
		"BBB = (AAA, ZZZ)\n",
		"ZZZ = (ZZZ, ZZZ)\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 6
	got := getSteps(s, false)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}
