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

func TestGetSumOfPowers(t *testing.T) {
	input := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\n",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\n",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green}\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 2286
	got := GetSumOfPowers(s)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestGetValidGameTotal(t *testing.T) {
	input := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\n",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\n",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green}\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 8
	availableCubes := gamebag{"red": 12, "blue": 14, "green": 13}
	got := GetValidGameTotal(s, availableCubes)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestVerifyValidGame(t *testing.T) {
	tests := []struct {
		input          gamebag
		availableCubes gamebag
		expected       bool
	}{{
		input:          gamebag{"red": 12, "blue": 14, "green": 13},
		availableCubes: gamebag{"red": 12, "blue": 14, "green": 13},
		expected:       true,
	}, {
		input:          gamebag{"red": 13, "blue": 14, "green": 13},
		availableCubes: gamebag{"red": 12, "blue": 14, "green": 13},
		expected:       false,
	},
	}
	for _, test := range tests {
		got := verifyValidGame(test.input, test.availableCubes)
		if got != test.expected {
			t.Logf("Expected %t, got %t", test.expected, got)
			t.Fail()
		}
	}
}

func TestColourCounts(t *testing.T) {
	input := "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	expected := gamebag{"red": 4, "blue": 6, "green": 2}
	res := extractCounts(input)
	for key, val := range res {
		if expected[key] != val {
			t.Logf("Expected key %s with value %d, got %d", key, expected[key], val)
			t.Fail()
		}
	}
}
