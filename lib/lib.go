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

// Package lib provides common functions used by all problems.
package lib

import (
	"bufio"
	"constraints"
	"os"
	"strconv"
)

// ForLine will read the contents of fname and call fn for each line.
func ForLine(fname string, fn func(line string)) {
	// Open input file.
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Setup scanner for line reading.
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// Process each line.
	for scanner.Scan() {
		fn(scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

// ToInt converts value to an integer and panics in case the conversion fails.
func ToInt(value string) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return v
}

// Numeric is a type constraint that accepts all go number types.
type Numeric interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// Sum returns the sum of all numbers in nums.
func Sum[N Numeric](nums []N) N {
	var value N
	for _, n := range nums {
		value += n
	}
	return value
}
