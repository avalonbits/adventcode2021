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
	"fmt"
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

type Set[T comparable] struct {
	set map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		set: map[T]struct{}{},
	}
}

func (s *Set[T]) Add(v T) *Set[T] {
	s.set[v] = struct{}{}
	return s
}

func (s *Set[T]) Insert(vs ...T) *Set[T] {
	for _, v := range vs {
		s.Add(v)
	}
	return s
}

func (s *Set[t]) Size() int {
	return len(s.set)
}

func (s *Set[T]) In(v T) bool {
	_, ok := s.set[v]
	return ok
}

func (s *Set[T]) Values() []T {
	values := make([]T, len(s.set))
	for key := range s.set {
		values = append(values, key)
	}
	return values
}

type Trie[T comparable] struct {
	alphabet *Set[T]
	count    int

	nodes map[T]*Trie[T]
}

func (t *Trie[T]) Count() int {
	return t.count
}

func NewTrie[T comparable](alphabet *Set[T]) *Trie[T] {
	return &Trie[T]{
		alphabet: alphabet,
		nodes:    map[T]*Trie[T]{},
	}
}

func (t *Trie[T]) Add(value ...T) {
	if len(value) == 0 {
		return
	}
	t.count++
	curr := t
	alphabet := t.alphabet
	for _, v := range value {
		if !alphabet.In(v) {
			panic(fmt.Sprintf("Invalid value '%v' not in alphabet %v", v, *t.alphabet))
		}
		next, ok := curr.nodes[v]
		if !ok {
			next = NewTrie(curr.alphabet)
			curr.nodes[v] = next
		}
		next.count++
		curr = next
	}
}

func (t *Trie[T]) Values() [][]T {
	all := make([][]T, t.count)
	start := 0
	for v, next := range t.nodes {
		values := next.allValues(v, all[start:start+next.count])
		for i := range values {
			all[start+i] = values[i]
		}
		start += next.count
	}
	return all
}

func (t *Trie[T]) allValues(v T, values [][]T) [][]T {
	if len(values) == 0 {
		return values
	}
	for i := range values {
		values[i] = append(values[i], v)
	}
	start := 0
	for v, next := range t.nodes {
		more := next.allValues(v, values[start:start+next.count])
		for i := range more {
			values[i+start] = more[i]
		}
		start += next.count
	}
	return values
}

func (t *Trie[T]) Longest(chain ...T) [][]T {
	curr := t
	prefix := []T{}
	for _, v := range chain {
		next, ok := curr.nodes[v]
		if !ok {
			break
		}
		prefix = append(prefix, v)
		curr = next
	}
	fmt.Println(prefix)
	values := make([][]T, curr.count)
	for i := range values {
		values[i] = append(values[i], prefix[:len(prefix)-1]...)
	}
	return curr.allValues(prefix[len(prefix)-1], values)
}

func (t *Trie[T]) Walk(fn func(nodes map[T]*Trie[T]) *Trie[T]) {
	curr := t
	count := 0
	for len(curr.nodes) != 0 {
		count++
		fmt.Println(count)
		curr = fn(curr.nodes)
	}
}
