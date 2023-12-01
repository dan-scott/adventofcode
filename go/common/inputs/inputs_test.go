package inputs

import (
	"strings"
	"testing"
)

func TestScannerFromString(t *testing.T) {
	inputStr := `a
bed
c
d`
	expected := []string{"a", "bed", "c", "d"}

	scanner := StrScanner(inputStr)

	for i, line := range expected {
		if !scanner.Scan() {
			t.Fatalf("expected scanner to have more tokens, but failed at %d", i)
		}
		next := scanner.Text()
		if strings.Compare(line, next) != 0 {
			t.Fatalf("expected token %d to equal %s but got %s", i, line, next)
		}
	}
}
