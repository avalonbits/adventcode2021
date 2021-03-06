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
	"strconv"

	"github.com/avalonbits/adventcode2021/lib"
	"github.com/avalonbits/adventcode2021/lib/collections"
)

const width = 12

func main() {
	alphabet := collections.NewSet[byte]().Insert('0', '1')
	trie := collections.NewTrie(alphabet)

	lib.ForLine("./input.txt", func(line string) {
		trie.Add([]byte(line)...)
	})

	most := []byte{}
	trie.Walk(func(nodes []collections.TNode[byte]) collections.TNode[byte] {
		var zero, one collections.TNode[byte]
		if ret := getZeroOne(nodes, &zero, &one); ret != nil {
			most = append(most, ret.V)
			return *ret
		}

		if one.Count >= zero.Count {
			most = append(most, one.V)
			return one
		} else {
			most = append(most, zero.V)
			return zero
		}
	})

	least := []byte{}
	trie.Walk(func(nodes []collections.TNode[byte]) collections.TNode[byte] {
		var zero, one collections.TNode[byte]
		if ret := getZeroOne(nodes, &zero, &one); ret != nil {
			least = append(least, ret.V)
			return *ret
		}

		if zero.Count <= one.Count {
			least = append(least, '0')
			return zero
		} else {
			least = append(least, '1')
			return one
		}
	})

	m, _ := strconv.ParseInt(string(most), 2, 64)
	l, _ := strconv.ParseInt(string(least), 2, 64)
	fmt.Println(string(most), m, string(least), l, m*l)
}

func getZeroOne(nodes []collections.TNode[byte], zero, one *collections.TNode[byte]) *collections.TNode[byte] {
	if len(nodes) == 1 {
		return &nodes[0]
	}
	for i := range nodes {
		if nodes[i].V == '0' {
			*zero = nodes[i]
		} else if nodes[i].V == '1' {
			*one = nodes[i]
		}
	}
	return nil
}
