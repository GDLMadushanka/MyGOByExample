package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, a := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	fmt.Println(a)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
}
