package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/avalonbits/adventcode2021/lib"
)

func main() {
	// Read all numbers from input.
	all := make([]int, 0, 2048)
	err := lib.ForLine("./prob1.input", func(line string) {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		all = append(all, value)
	})
	if err != nil {
		panic(err)
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
