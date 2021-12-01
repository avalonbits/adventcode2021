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

	// Read all numbers from input.
	all := make([]int, 0, 10240)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		all = append(all, value)
	}

	// For each 3-numbers, do the sum and compare with last seen.
	increasedCount := 0
	last := math.MaxInt
	for window := 0; window < len(all); window++ {
		// Make sure we are within bounds when doing the sums.
		value := all[window]
		if window < len(all)-1 {
			value += all[window+1]
		}
		if window < len(all)-2 {
			value += all[window+2]
		}

		if last < value {
			increasedCount++
		}
		last = value
	}
	fmt.Println(increasedCount)
}
