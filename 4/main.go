package main

import "fmt"

func main() {
	count := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if north(&input, i, j) {
				count++
			}
			if south(&input, i, j) {
				count++
			}
			if east(&input, i, j) {
				count++
			}
			if west(&input, i, j) {
				count++
			}
			if northEast(&input, i, j) {
				count++
			}
			if southEast(&input, i, j) {
				count++
			}
			if southWest(&input, i, j) {
				count++
			}
			if northWest(&input, i, j) {
				count++
			}
		}
	}
	fmt.Print(count)
}
