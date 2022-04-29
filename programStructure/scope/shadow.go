package main

import "fmt"

func f() {}

var x = "x"

func main() {
	f := "f"
	fmt.Println(f) // "f"; local var f shadows package-level func f
	fmt.Println(x) // "x"; package-level var
	// fmt.Println(h) // compile error: undefined: h
}
