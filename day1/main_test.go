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

func TestCalibrateValues(t *testing.T) {
	testVal := []string{"1abc2\n",
		"pqr3stu8vwx\n",
		"a1b2c3d4e5f\n",
		"treb7uchet\n"}
	mockFile := mockFile{data: testVal}
	s := bufio.NewScanner(&mockFile)
	res := CalibrateValues(s)
	if res != 142 {
		t.Logf("Result %d != 142", res)
		t.Fail()
	}
}

func TestCalibrateValuesWithStringNumbers(t *testing.T) {
	testVal := []string{"two1nine\n",
		"eightwothree\n",
		"abcone2threexyz\n",
		"xtwone3four\n",
		"4nineeightseven2\n",
		"zoneight234\n",
		"7pqrstsixteen\n"}
	mockFile := mockFile{data: testVal}
	s := bufio.NewScanner(&mockFile)
	res := CalibrateValuesWithStringNumbers(s)
	if res != 281 {
		t.Logf("Result %d != 281", res)
		t.Fail()
	}
}
