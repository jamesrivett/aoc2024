package main

import (
	"fmt"
	"slices"
	"sort"
)

func Abs(x int) int {
	if x < 0 {
		return (0 - x)
	}
	return x
}

var safeReports = 0

func main() {
	// for i := 0; i < 10; i++ {
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		sorted := slices.Clone[[]int](input)
		sort.IntSlice.Sort(sorted)
		reversed := slices.Clone[[]int](sorted)
		slices.Reverse[[]int](reversed)
		safe := true

		if !slices.Equal[[]int](input, sorted) && !slices.Equal[[]int](input, reversed) {
			safe = false
		}

		for j := 1; j < len(sorted); j++ {
			if Abs(sorted[j]-sorted[j-1]) > 3 ||
				Abs(sorted[j]-sorted[j-1]) == 0 {
				safe = false
			}
		}

		if safe == true {
			fmt.Println("Input: ", input)
			fmt.Println("Sorted: ", sorted)
			fmt.Println("Reversed: ", reversed)
			fmt.Print("\n")
			safeReports++
		}
	}
	fmt.Print(safeReports, "\n")
}
