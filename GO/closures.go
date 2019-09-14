/*
Go supports anonymous functions, which can form closures.
Anonymous functions are useful
when you want to define a function inline without
having to name it.*/

package main

import "fmt"

func inSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	nextInt := inSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := inSeq()
	fmt.Println(newInts())

	val := fact(7)
	fmt.Println(val)
}
