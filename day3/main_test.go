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

func TestGearRatioTotal(t *testing.T) {
	input := []string{
		"467..114..\n",
		"...*......\n",
		"..35..633.\n",
		"......#...\n",
		"617*......\n",
		".....+.58.\n",
		"..592.....\n",
		"......755.\n",
		"...$.*....\n",
		".664.598..,\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 467835
	got := sumParts(s, sumGearRatios)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestSumEnginePartNumbers(t *testing.T) {
	input := []string{"467..114..\n",
		"...*......\n",
		"..35..633.\n",
		"......#...\n",
		"617*......\n",
		".....+.58.\n",
		"..592.....\n",
		"......755.\n",
		"...$.*....\n",
		".664.598..,\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 4361
	got := sumParts(s, sumValidNumbersInRow)
	if got != expected {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}
func TestIsVerticallyAdjacent(t *testing.T) {
	tests := []struct {
		about    string
		adjLine  string
		line     string
		index    int
		length   int
		expected bool
	}{{
		about:    "Returns true when adjacent",
		adjLine:  "...*......",
		line:     "..35..633.",
		index:    2,
		length:   2,
		expected: true,
	}, {
		about:    "Returns true when adjacent at beginning of line",
		adjLine:  "*.........",
		line:     "35....633.",
		index:    0,
		length:   2,
		expected: true,
	}, {
		about:    "Returns false when not adjacent (marker before)",
		adjLine:  "*.........",
		line:     "..35..633.",
		index:    2,
		length:   2,
		expected: false,
	}, {
		about:    "Returns false when not adjacent (marker after)",
		adjLine:  ".....*",
		line:     ".162..",
		index:    1,
		length:   3,
		expected: false,
	}, {
		about:    "Returns true when adjacent at end of line",
		adjLine:  ".........*",
		line:     "..35...633",
		index:    7,
		length:   3,
		expected: true,
	}, {
		about:    "Returns true when adjacent and number is length 1",
		adjLine:  "./.......*",
		line:     "..3....633",
		index:    2,
		length:   1,
		expected: true,
	}}
	for _, test := range tests {
		t.Run(test.about, func(t *testing.T) {
			lines := newLines(test.adjLine, test.line, "")
			p := newPartFinder(lines, reSpecialChar, test.index, test.length)
			_, got := p.isVerticallyAdjacent(true)
			if got != test.expected {
				t.Logf("Expected %t, got %t", test.expected, got)
				t.Fail()
			}
		})

	}
}

func TestIsVerticallyAdjacentLocations(t *testing.T) {
	tests := []struct {
		about   string
		adjLine string
		line    string
		index   int
		length  int
		locs    [][]int
	}{{
		about:   "Returns absolute indices",
		adjLine: "...3......",
		line:    "..*.....",
		index:   2,
		length:  1,
		locs:    [][]int{{3, 4}},
	}, {
		about:   "Returns absolute indices (two matches)",
		adjLine: "12.45.....",
		line:    "..*.....",
		index:   2,
		length:  1,
		locs:    [][]int{{1, 2}, {3, 4}},
	}, {
		about:   "Returns absolute indices (long match)",
		adjLine: ".2345.....",
		line:    "..*.....",
		index:   2,
		length:  1,
		locs:    [][]int{{1, 4}},
	}}
	for _, test := range tests {
		t.Run(test.about, func(t *testing.T) {
			lines := newLines(test.adjLine, test.line, "")
			p := newPartFinder(lines, reDigitMatch, test.index, test.length)
			locs, _ := p.isVerticallyAdjacent(true)
			for i, loc := range locs {
				for j, index := range loc {
					if test.locs[i][j] != index {
						t.Logf("Expected %d, got %d at position [%d][%d]", test.locs[i][j], index, i, j)
						t.Fail()
					}
				}
			}
		})
	}
}

func TestGetExpandedInt(t *testing.T) {
	tests := []struct {
		about string
		line  string
		index int
		want  int
	}{{
		about: "index first num",
		line:  "..123.....",
		index: 2,
		want:  123,
	}, {
		about: "index middle num",
		line:  "..123.....",
		index: 3,
		want:  123,
	}, {
		about: "start of string",
		line:  "123.....",
		index: 1,
		want:  123,
	}, {
		about: "end of string",
		line:  "..123",
		index: 2,
		want:  123,
	}}
	for _, test := range tests {
		t.Run(test.about, func(t *testing.T) {
			got := getExpandedInt(test.line, test.index)
			if test.want != got {
				t.Logf("Expected %d, got %d", test.want, got)
				t.Fail()
			}
		})
	}
}

func TestIsHorizontallyAdjacent(t *testing.T) {
	tests := []struct {
		about    string
		line     string
		index    int
		length   int
		expected bool
	}{{
		about:    "Returns true when adjacent in front",
		line:     ".*35..633.",
		index:    2,
		length:   2,
		expected: true,
	}, {
		about:    "Returns true when adjacent behind",
		line:     ".35*..633.",
		index:    1,
		length:   2,
		expected: true,
	}, {
		about:    "Returns true when starting the line and adjacent behind",
		line:     "35*...633.",
		index:    0,
		length:   2,
		expected: true,
	}, {
		about:    "Returns true when ending the line and adjacent in front",
		line:     ".......*33",
		index:    8,
		length:   2,
		expected: true,
	}, {
		about:    "Returns false when not adjacent",
		line:     "........33",
		index:    8,
		length:   2,
		expected: false,
	}, {
		about:    "Returns false when not adjacent",
		line:     "*.334.%..",
		index:    2,
		length:   3,
		expected: false,
	}}
	for _, test := range tests {
		t.Run(test.about, func(t *testing.T) {
			lines := newLines("", test.line, "")
			p := newPartFinder(lines, reSpecialChar, test.index, test.length)
			_, got := p.isAdjacent()
			if got != test.expected {
				t.Logf("Expected %t, got %t", test.expected, got)
				t.Fail()
			}
		})
	}
}
