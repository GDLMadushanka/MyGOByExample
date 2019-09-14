package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plus3(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 8
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plus3(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	a, b := vals()
	_, c := vals()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
