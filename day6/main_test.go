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

func TestSingleWin(t *testing.T) {
	input := []string{
		"Time:      7  15   30\n",
		"Distance:  9  40  200\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := uint64(71503)
	got := findSingleWin(s)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestFindLowestLocationPartTwo(t *testing.T) {
	input := []string{
		"Time:      7  15   30\n",
		"Distance:  9  40  200\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := uint64(288)
	got := findWinsMultiplied(s)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestMaximumHoldTime(t *testing.T) {
	tests := []struct {
		raceTime uint64
		distance uint64
		want     uint64
		ok       bool
	}{{
		raceTime: 7,
		distance: 9,
		want:     5,
		ok:       true,
	}, {
		raceTime: 5,
		distance: 10,
		want:     5,
		ok:       false,
	},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got, ok := determineMaximumTime(test.raceTime, test.distance)
			if ok != test.ok {
				t.Logf("Not ok")
				t.Fail()
			}
			if ok && (got != test.want) {
				t.Logf("Want %d, got %d", test.want, got)
				t.Fail()
			}
		})
	}
}

func TestMinimumHoldTime(t *testing.T) {
	tests := []struct {
		raceTime uint64
		distance uint64
		want     uint64
		ok       bool
	}{{
		raceTime: 7,
		distance: 9,
		want:     2,
		ok:       true,
	}, {
		raceTime: 5,
		distance: 10,
		want:     5,
		ok:       false,
	},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got, ok := determineMinimumHoldTime(test.raceTime, test.distance)
			if ok != test.ok {
				t.Logf("Not ok")
				t.Fail()
			}
			if ok && (got != test.want) {
				t.Logf("Want %d, got %d", test.want, got)
				t.Fail()
			}
		})
	}

}
