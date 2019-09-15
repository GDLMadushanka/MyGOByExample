/*
Defer is used to ensure that a function call is
performed later in a program’s execution,
usually for purposes of cleanup
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "hello world")
}

func closeFile(f *os.File) {
	fmt.Println("Closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error : %v\n", err)
		os.Exit(1)
	}
}
