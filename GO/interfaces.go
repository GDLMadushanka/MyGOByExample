package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radious float64
}

/*
To implement an interface in Go, we just need to
implement all the methods in the interface.
Here we implement geometry on rects.
*/

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radious * c.radious
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radious
}

/*
If a variable has an interface type,
then we can call methods that are in the named interface.
Here’s a generic measure function taking advantage
of this to work on any geometry.
*/
func messure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radious: 7}

	messure(r)
	messure(c)
}
