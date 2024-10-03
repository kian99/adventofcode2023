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

func (m *mockFile) Reset() {
	m.lineCount = 0
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

func TestBigExpansionSumPathLength(t *testing.T) {
	input := []string{
		"...#......\n",
		".......#..\n",
		"#.........\n",
		"..........\n",
		"......#...\n",
		".#........\n",
		".........#\n",
		"..........\n",
		".......#..\n",
		"#...#.....\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 1030
	got := sumShortestPaths(s, 10)
	if expected != got {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
	mockFile.Reset()
	s = bufio.NewScanner(&mockFile)
	expected = 8410
	got = sumShortestPaths(s, 100)
	if expected != got {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestSumPathLength(t *testing.T) {
	input := []string{
		"...#......\n",
		".......#..\n",
		"#.........\n",
		"..........\n",
		"......#...\n",
		".#........\n",
		".........#\n",
		"..........\n",
		".......#..\n",
		"#...#.....\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expected := 374
	got := sumShortestPaths(s, 2)
	if expected != got {
		t.Logf("Expected %d, got %d", expected, got)
		t.Fail()
	}
}

func TestFillEmpty(t *testing.T) {
	input := []string{
		"...#......\n",
		".......#..\n",
		"#.........\n",
		"..........\n",
		"......#...\n",
		".#........\n",
		".........#\n",
		"..........\n",
		".......#..\n",
		"#...#.....\n",
	}
	mockFile := mockFile{data: input}
	s := bufio.NewScanner(&mockFile)
	expectedEmptyRows := map[int]struct{}{3: {}, 7: {}}
	expectedEmptyCols := map[int]struct{}{2: {}, 5: {}, 8: {}}
	u := newUniverse(2)
	for s.Scan() {
		line := s.Text()
		u.image = append(u.image, line)
	}
	u.fillEmpty()
	if len(expectedEmptyCols) != len(u.emptyColumns) || len(expectedEmptyRows) != len(u.emptyRows) {
		t.Logf("Incorrect lengths")
		t.Fail()
	}
	for col := range expectedEmptyCols {
		if _, ok := u.emptyColumns[col]; !ok {
			t.Logf("Expected %d", col)
			t.Fail()
		}
	}
	for row := range expectedEmptyRows {
		if _, ok := u.emptyRows[row]; !ok {
			t.Logf("Expected %d", row)
			t.Fail()
		}
	}

}
