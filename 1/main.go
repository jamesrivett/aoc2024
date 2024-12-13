package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	// READING IN THE FILE
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	buffer := make([]byte, 1024)
	numbers := make([]string, 1, 5000)
	currentNumber := 0
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if bytesRead == 0 {
			break
		}
		for currentByte := 0; currentByte < bytesRead; currentByte++ {
			if string(buffer[currentByte]) == " " || string(buffer[currentByte]) == "\n" || string(buffer[currentByte]) == "\r" {
				if numbers[currentNumber] != "" {
					numbers = numbers[:len(numbers)+1]
					currentNumber++
					continue
				}
				continue
			}
			numbers[currentNumber] += string(buffer[currentByte])
		}
	}

	// SPLITTING AND CONVERTING THE ARRAYS
	evens := make([]int, 1, 1001)
	odds := make([]int, 1, 1001)
	for i := 0; i < len(numbers); i++ {
		if i%2 == 0 {
			converted, err := strconv.Atoi(numbers[i])
			if err != nil {
				panic(err)
			}
			evens[len(evens)-1] = converted
			evens = evens[:len(evens)+1]
		} else {
			converted, err := strconv.Atoi(numbers[i])
			if err != nil {
				panic(err)
			}
			odds[len(odds)-1] = converted
			odds = odds[:len(odds)+1]
		}
	}

	sort.Ints(evens)
	sort.Ints(odds)

	// ACTUALLY SOLVING THE PROBLEM
	totalDistance := 0
	for i := 0; i < len(evens); i++ {
		totalDistance += int(math.Abs(float64(evens[i] - odds[i])))
	}
	fmt.Printf("total distance: %d", totalDistance)
}
