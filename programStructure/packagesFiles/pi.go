package main

import (
	"fmt"
	pc "mymodule/popcount"
)

var a = b + c // a initialized third, to 3
var b = f()   // b initialized second, to 2, by calling f
var c = 1     // c initialized first, to 1
var d int

func f() int { return c + 1 }

// init() func can't be called or referenced.
// they automatically executed when program starts
// in order in which they are declared
func init() {
	d = 4
}

func init() {
	d = 3
}

func main() {
	fmt.Printf("a: %v, b:%v, c:%v, d:%v \n", a, b, c, d)

	fmt.Println(pc.PopCount(8))
}
