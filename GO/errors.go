package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with this")
	}
	return arg + 10, nil
}

/*
It’s possible to use custom types as errors by
implementing the Error() method on them.
*/
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// we use &argError syntax to build a new struct, supplying values for the two fields arg and prob.
		return -1, &argError{arg, "can't work with this 2 "}
	}
	return arg + 10, nil
}

func main() {
	for _, i := range []int{7, 42} {
		// Note that the use of an inline error
		// check on the if line is a common idiom in Go code.
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	// go type assertion
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
