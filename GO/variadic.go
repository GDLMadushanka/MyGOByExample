package main

import "fmt"

// variadic : Accepting a variable number of arguments
// fmt.Println is a common variadic function
func sum(nums ...int) int {
	fmt.Println("nums : ", nums)
	total := 0
	for _, number := range nums {
		total += number
	}
	return total
}

func main() {
	c := sum(1, 2, 3, 5, 8)
	fmt.Println(c)

	nums := []int{2, 2, 2}
	fmt.Println(sum(nums...)) // calling using a slice
}
