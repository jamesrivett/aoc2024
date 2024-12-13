package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputString := ""
	buf := make([]byte, 1024)

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

	for {
		bytesRead, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if bytesRead == 0 {
			break
		}

		bufString := string(buf[:bytesRead])
		inputString = (inputString + bufString)
	}

	mulChecker := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	matches := mulChecker.FindAll([]byte(inputString), -1)
	matchesStrings := []string{}
	for i := 0; i < len(matches); i++ {
		matchesStrings = append(matchesStrings, string(matches[i]))
	}
	fmt.Print(matchesStrings)
	flatMatches := []byte{}
	for _, j := range matches {
		flatMatches = append(flatMatches, j...)
	}

	// fmt.Println(string(flatMatches))

	leftGrabber := regexp.MustCompile(`[0-9]+,`)
	rightGrabber := regexp.MustCompile(`[0-9]+\)`)
	leftBytes := leftGrabber.FindAll(flatMatches, -1)
	rightBytes := rightGrabber.FindAll(flatMatches, -1)
	leftStrings := []string{}
	rightStrings := []string{}
	leftNums := []int{}
	rightNums := []int{}
	sum := 0

	for i := 0; i < len(leftBytes); i++ {
		leftStrings = append(leftStrings, string(leftBytes[i])[:len(string(leftBytes[i]))-1])
	}
	for i := 0; i < len(rightBytes); i++ {
		rightStrings = append(rightStrings, string(rightBytes[i])[:len(string(rightBytes[i]))-1])
	}
	for i := range leftStrings {
		newInt, err := strconv.Atoi(leftStrings[i])
		if err != nil {
			panic(err)
		}
		leftNums = append(leftNums, newInt)
	}
	for i := range rightStrings {
		newInt, err := strconv.Atoi(rightStrings[i])
		if err != nil {
			panic(err)
		}
		rightNums = append(rightNums, newInt)
	}

	for i := range leftNums {
		sum = sum + leftNums[i]*rightNums[i]
		// fmt.Printf("%d * %d = %d \n", leftNums[i], rightNums[i], leftNums[i]*rightNums[i])
	}
	fmt.Println(len(leftNums))
	fmt.Print(sum)
}
