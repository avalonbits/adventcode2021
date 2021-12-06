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

	"github.com/avalonbits/adventcode2021/lib"
)

func main() {
	const width = 12
	bitCount := make([]int, width)
	lineCount := 0
	lib.ForLine("./input.txt", func(line string) {
		lineCount++
		for i := 0; i < width; i++ {
			bitCount[width-i-1] += int(line[i] - '0')
		}
	})

	gamaRate := 0
	for i := 0; i < width; i++ {
		isOne := (lineCount - bitCount[i]) < bitCount[i]
		if isOne {
			gamaRate += (1 << i)
		}
	}
	epsilonRate := ^gamaRate & 0x00000FFF
	fmt.Println(gamaRate * epsilonRate)
}
