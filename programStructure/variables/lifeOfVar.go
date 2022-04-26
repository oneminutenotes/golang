// lifetime of variable is interval of time
// during which it exists as the program executes
package main

import (
	"fmt"
)

// package level; lifetime is entire execution of program
var global *int

func main() {

	// local level; dynamic lifetime
	// 'i' created each time loop begins
	for i := 0; i < 3; i++ {
		x := i * i // 'x' created on each iteration
		fmt.Printf("square of %d: %d \n", i, x)
		// after end of iteration 'x' gets recycled/reclaimed
	} // after loop ends 'i' gets recycled/reclaimed

	// heap vs stack allocation of variables
	f()
	g()

}

func f() {
	// here x is heap allocated as x exists even
	// after function returns i.e x escapes from f()
	var x int
	x = 1
	global = &x
}

func g() {
	// here y is stack allocated cause y becomes unreachable
	// after function returns i.e y does not escape from g()
	y := new(int)
	*y = 1
}
