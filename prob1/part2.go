/*
 * Copyright (C) 2021  Igor Cananea <icc@avalonbits.com>
 * Author: Igor Cananea <icc@avalonbits.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"math"

	"github.com/avalonbits/adventcode2021/lib"
)

func main() {
	// Read all numbers from input.
	all := make([]int, 0, 2048)
	lib.ForLine("./prob1.input", func(line string) {
		all = append(all, lib.ToInt(line))
	})

	// For each 3-numbers, do the sum and compare with last seen.
	increasedCount := 0
	last := math.MaxInt
	for window := 0; window < len(all)-2; window++ {
		// Make sure we are within bounds when doing the sums.
		value := lib.Sum(all[window : window+3])

		if last < value {
			increasedCount++
		}
		last = value
	}
	fmt.Println(increasedCount)
}
