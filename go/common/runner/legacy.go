package runner

import (
	"fmt"
	"github.com/loov/hrtime"
)

type LegacyDay interface {
	Open()
	Close()
	Part1() string
	Part2() string
}

func RunLegacy(days []LegacyDay) {
	allStart := hrtime.Now()

	for i, d := range days {

		fmt.Printf("Loading Day %d....", i+1)
		startLoad := hrtime.Now()
		d.Open()
		fmt.Printf("\t  loaded in %15s \n", hrtime.Since(startLoad).String())

		fmt.Printf("\tSolving part 1...")
		startP1 := hrtime.Now()
		res := d.Part1()
		fmt.Printf(" solved in %15s", hrtime.Since(startP1).String())
		fmt.Printf("%20s\n", res)

		fmt.Printf("\tSolving part 2...")
		startP2 := hrtime.Now()
		res = d.Part2()
		fmt.Printf(" solved in %15s", hrtime.Since(startP2).String())
		fmt.Printf("%20s\n\n", res)

		d.Close()
	}

	fmt.Printf("\nSolutions completed in %s\n", hrtime.Since(allStart).String())
}
