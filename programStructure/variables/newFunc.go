// Expression: pointerName := new(T)
package main

import (
	"fmt"
)

func main() {
	// p, of type *int, points to an unnamed int
	p := new(int)
	fmt.Printf("p: %v,%T \n", p, p)
	fmt.Printf("*p: %v,%T \n", *p, *p)

	// update variable
	*p = 2
	fmt.Printf("*p: %v,%T \n", *p, *p)

	// 'new()' and '&' are fundamentally same
	fmt.Println("new():", f(), ", &:", newInt())

	// each call to new returns new address
	// exception are struct{}, [0]int
	m := new(int)
	n := new(int)
	fmt.Println(m == n) // "false"

	// redefine new
	fmt.Println(delta(1, 2))

}

func delta(old, new int) int {
	// new is predeclared func, not keyword
	// here built-in new func is unavailable
	return new - old
}

func f() *int {
	return new(int)
}
func newInt() *int {
	var dummy int
	return &dummy
}
