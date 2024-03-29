package main

import "fmt"

type person struct {
	name string
	age  int
}

func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 234
	return &p
}

func main() {

	fmt.Println(person{"bob", 30})

	fmt.Println(person{name: "alice", age: 45})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(NewPerson("Lahiru"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
}
