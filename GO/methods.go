package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perm() int {
	return 2 * (r.height + r.width)
}

func main() {
	r := rect{width: 45, height: 30}
	fmt.Println("area", r.area())
	fmt.Println("area", r.perm())

	/*
		Go automatically handles conversion between
		values and pointers for method calls.
		You may want to use a pointer receiver type
		to avoid copying on method calls or to allow
		the method to mutate the receiving struct.
	*/
	b := &r
	fmt.Println("area", b.area())
	fmt.Println("area", b.perm())
}
