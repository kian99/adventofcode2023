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

func TestFindLowestLocationPartTwo(t *testing.T) {
	input := []string{
		"seeds: 79 14 55 13\n",
		"\n",
		"seed-to-soil map:\n",
		"50 98 2\n",
		"52 50 48\n",
		"\n",
		"soil-to-fertilizer map:\n",
		"0 15 37\n",
		"37 52 2\n",
		"39 0 15\n",
		"\n",
		"fertilizer-to-water map:\n",
		"49 53 8\n",
		"0 11 42\n",
		"42 0 7\n",
		"57 7 4\n",
		"\n",
		"water-to-light map:\n",
		"88 18 7\n",
		"18 25 70\n",
		"\n",
		"light-to-temperature map:\n",
		"45 77 23\n",
		"81 45 19\n",
		"68 64 13\n",
		"\n",
		"temperature-to-humidity map:\n",
		"0 69 1\n",
		"1 0 69\n",
		"\n",
		"humidity-to-location map:\n",
		"60 56 37\n",
		"56 93 4\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := uint64(46)
	got := findLowestLocationPartTwo(s)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestFindLowestLocation(t *testing.T) {
	input := []string{
		"seeds: 79 14 55 13\n",
		"\n",
		"seed-to-soil map:\n",
		"50 98 2\n",
		"52 50 48\n",
		"\n",
		"soil-to-fertilizer map:\n",
		"0 15 37\n",
		"37 52 2\n",
		"39 0 15\n",
		"\n",
		"fertilizer-to-water map:\n",
		"49 53 8\n",
		"0 11 42\n",
		"42 0 7\n",
		"57 7 4\n",
		"\n",
		"water-to-light map:\n",
		"88 18 7\n",
		"18 25 70\n",
		"\n",
		"light-to-temperature map:\n",
		"45 77 23\n",
		"81 45 19\n",
		"68 64 13\n",
		"\n",
		"temperature-to-humidity map:\n",
		"0 69 1\n",
		"1 0 69\n",
		"\n",
		"humidity-to-location map:\n",
		"60 56 37\n",
		"56 93 4\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := uint64(35)
	got := findLowestLocation(s)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestInMap(t *testing.T) {
	r := rangeTuple{destination: 52, source: 50, length: 48}
	want := uint64(55)
	got, ok := r.inMap(53)
	if !ok {
		t.Fail()
	}
	if got != want {
		t.Logf("Want %d, got %d", want, got)
		t.Fail()
	}
}
