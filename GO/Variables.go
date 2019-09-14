package main

import "fmt"

func main()  {
	var a string = "lahiru"
	fmt.Println(a)
	var b,c int = 3,4
	fmt.Println(b,c)
	var d float32 = 4.5234
	fmt.Println("%.2f",d)
	f := fmt.Sprintf("%2.3f",d)
	fmt.Println(f)
}