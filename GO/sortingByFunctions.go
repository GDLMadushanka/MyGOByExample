/*
Sometimes we’ll want to sort a collection by something
other than its natural order. For example, suppose
we wanted to sort strings by their length instead
of alphabetically
*/
package main

import (
	"fmt"
	"sort"
)

//In order to sort by a custom function in Go,
//we need a corresponding type. Here we’ve created a
//byLength type that is just an alias for the builtin
//[]string type.
type byLength []string

//We implement sort.Interface - Len, Less,
//and Swap - on our type so we can use the
//sort package’s generic Sort function
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func main() {
	fruits := []string{"banana", "pinapple", "apple"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
