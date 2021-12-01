package inputs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func LinesAsString(year, day uint) []string {
	wd, _ := os.Getwd()
	df := filepath.Join(wd, "inputs", fmt.Sprint(year), fmt.Sprintf("%d.txt", day))
	f, err := os.Open(df)
	if err != nil {
		panic(fmt.Sprintf("Error reading lines from input file %s: %v", df, err))
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	s := bufio.NewScanner(f)
	lines := make([]string, 0)

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
