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
	"strings"

	"github.com/avalonbits/adventcode2021/lib"
)

func main() {
	horz := 0
	depth := 0
	lib.ForLine("./input.txt", func(line string) {
		entry := strings.Split(line, " ")
		cmd := entry[0]
		value := lib.ToInt(entry[1])

		switch cmd {
		case "forward":
			horz += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	})
	fmt.Println(horz * depth)
}
