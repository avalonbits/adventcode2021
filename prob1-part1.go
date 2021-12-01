package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/avalonbits/adventcode2021/lib"
)

func main() {
	last := math.MaxInt
	increaseCount := 0

	err := lib.ForLine("./prob1.input", func(line string) {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if last < value {
			increaseCount++
		}
		last = value
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(increaseCount)
}
