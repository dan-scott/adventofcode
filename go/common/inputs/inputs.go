package inputs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type CloseFn = func()

func Scanner(year, day uint) (*bufio.Scanner, CloseFn) {
	rootDir := ".."
	if d, ok := os.LookupEnv("ADVENT_OF_CODE_ROOT"); ok {
		rootDir = d
	} else {
		wd, _ := os.Getwd()
		rootDir = filepath.Join(wd, "../")
	}
	df := filepath.Join(rootDir, "inputs", fmt.Sprint(year), fmt.Sprintf("%d.txt", day))
	f, err := os.Open(df)
	if err != nil {
		panic(fmt.Sprintf("Error reading lines from input file %s: %v", df, err))
	}
	return bufio.NewScanner(f), func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}
}

type stringReader struct {
	cursor int
	str    string
}

func (s *stringReader) Read(p []byte) (n int, err error) {
	ct := copy(p, s.str[s.cursor:])
	s.cursor += ct
	return ct, nil
}

func StrScanner(input string) *bufio.Scanner {
	reader := stringReader{cursor: 0, str: input}
	return bufio.NewScanner(&reader)
}

func ScanWords(input string) *bufio.Scanner {
	scanner := StrScanner(input)
	scanner.Split(bufio.ScanWords)
	return scanner
}

func LinesAsString(year, day uint) []string {
	lines := make([]string, 0)
	s, closeScanner := Scanner(year, day)
	defer closeScanner()
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func LinesAsInt(year, day uint) []int {
	lines := LinesAsString(year, day)
	intLs := make([]int, len(lines))
	for i, s := range lines {
		if val, err := strconv.ParseInt(s, 10, 64); err != nil {
			panic(fmt.Sprintf("Error parsing input string %d '%s' to int: %v", i, s, err))
		} else {
			intLs[i] = int(val)
		}
	}
	return intLs
}
