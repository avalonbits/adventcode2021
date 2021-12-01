package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// Open input file.
	f, err := os.Open("./prob1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Setup scanner for line reading.
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// For each line, convert to int and compare with last seen value.
	last := math.MaxInt
	increaseCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if last < value {
			increaseCount++
		}
		last = value
	}
	fmt.Println(increaseCount)
}
