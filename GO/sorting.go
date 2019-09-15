package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "b", "a"}
	// Note that sorting is in-place, so it changes
	// the given slice and doesn’t return a new one.
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{5, 4, 3}
	sort.Ints(ints)
	fmt.Println(ints)

	// We can also use sort to check if a slice is already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted", s)
}
