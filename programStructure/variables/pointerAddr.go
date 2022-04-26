// pointers can store variable address & can update variable value
package main

import (
	"fmt"
)

func main() {

	// declare variable
	var x int = 1
	fmt.Printf("x: %v,%T \n", x, x)

	// declare pointer
	// '&' operator may only be applied to addresable values
	p := &x // p of type *int, points to x
	fmt.Printf("p: %v,%T \n", p, p)

	// yield value of variable pointed by pointer
	fmt.Printf("*p: %v,%T \n", *p, *p)

	// update value by pointer
	*p = 2 // equivalent to x = 2

	// print x
	fmt.Printf("x: %v,%T \n", x, x)

	// declare pointer with its zero value
	var q *int
	fmt.Printf("q: %v,%T \n", q, q)

	// pointers are comparable
	var y int
	fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"

	// its safe for function to return pointer
	p = newInt()
	fmt.Printf("p: %v,%T \n", p, p)

	// each call to newInt() returns different value
	fmt.Println(newInt() == newInt())  // "false"	

	// pass pointer as function argument
	c := 1
	incr(&c) // side effect: c is now 2
	fmt.Println("c:",incr(&c)) // "3" (and c is 3)
}

func incr(p *int) int {
	*p++ // increments what p points to; p remains same
	return *p // here *p is alias of c
}

func f() *int {
	return new(int)
}

func newInt() *int {
	var dummy int
	return &dummy
}
