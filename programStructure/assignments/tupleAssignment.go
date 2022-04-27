// Form: v1,v2,... = e1,e2,... //v=variable,e=expression
// expression are evaluated before variables update
package main

import (
	"fmt"
	"os"
)

func main() {
	// variable swap
	x := 1
	y := 2
	fmt.Println("x:", x, "| y:", y)
	x, y = y, x
	fmt.Println("x:", x, "| y:", y)

	// array elements
	c := [3]int{1, 2, 3}
	i, j := 0, 1 // declaration
	fmt.Println("c:", c)
	c[i], c[j] = c[j], c[i] // tuple assignment
	fmt.Println("c:", c)

	// function call
	gcd(30, 45)
	fib(7)

	// function call that returns two values
	f, err := os.Open("foo.txt") // declaration
	if err != nil {
		fmt.Println(err)
	}
	f.Close()
	f, err = os.Open("bar.txt") // tuple assignment
	if err != nil {
		fmt.Println(err)
	}
	f.Close()

	// other places which we can use
	// v, ok = m[key] // map lookup
	// v, ok = x.(T)  // type assertion
	// v, ok = <-ch   // channel receive

	// we can assign unwanted result to blank identifier
	// _, err = io.Copy(dst, src) // discard byte count
	// _, ok = x.(T)              // check type but discard result

}

func gcd(x, y int) int {
	fmt.Println("gcd of", x, y)
	for y != 0 {
		fmt.Println("x:", x, "| y:", y)
		x, y = y, x%y
	}
	fmt.Println("x:", x, "| y:", y)
	return x
}

func fib(n int) int {
	fmt.Println("fibonacci of", n)
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println("x:", x, "| y:", y)
		x, y = y, x+y
	}
	return x
}
