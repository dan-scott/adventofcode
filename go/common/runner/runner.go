package runner

import (
	"bufio"
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/loov/hrtime"
)

type Day interface {
	Year() uint
	Day() uint
	Part1(input *bufio.Scanner) string
	Part2(input *bufio.Scanner) string
}

func Run(days []Day) {
	allStart := hrtime.Now()

	for i, d := range days {

		fmt.Printf("Loading Day %d....", i+1)
		startLoad := hrtime.Now()
		scanner1, closeScanner1 := inputs.Scanner(d.Year(), d.Day())
		scanner2, closeScanner2 := inputs.Scanner(d.Year(), d.Day())
		fmt.Printf("\t  loaded in %15s \n", hrtime.Since(startLoad).String())

		fmt.Printf("\tSolving part 1...")
		startP1 := hrtime.Now()
		res := d.Part1(scanner1)
		fmt.Printf(" solved in %15s", hrtime.Since(startP1).String())
		fmt.Printf("%20s\n", res)

		closeScanner1()

		fmt.Printf("\tSolving part 2...")
		startP2 := hrtime.Now()
		res = d.Part2(scanner2)
		fmt.Printf(" solved in %15s", hrtime.Since(startP2).String())
		fmt.Printf("%20s\n\n", res)

		closeScanner2()
	}

	fmt.Printf("\nSolutions completed in %s\n", hrtime.Since(allStart).String())
}
