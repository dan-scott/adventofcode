package inputs

import (
	"bufio"
	"strconv"
)

type IntScanner struct {
	scanner *bufio.Scanner
}

func ScanInts(input string) *IntScanner {
	scanner := ScanWords(input)
	return &IntScanner{scanner: scanner}
}

func (i *IntScanner) Scan() bool {
	return i.scanner.Scan()
}

func (i *IntScanner) Int() int {
	v, _ := strconv.ParseInt(i.scanner.Text(), 10, 64)
	return int(v)
}

func (i *IntScanner) NextInt() (int, bool) {
	hasNext := i.Scan()
	return i.Int(), hasNext
}
