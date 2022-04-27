// assignment general form: variable = expression
package main

import (
	"fmt"
)

type Person struct {
	name string
}

func main() {
	// named variable
	var x int
	fmt.Println("x: ", x)
	x = 1
	fmt.Println("x: ", x)

	// indirect variable
	p := new(bool) // p is pointer of type *bool
	fmt.Println("*p: ", *p, "| p:", p)
	*p = true
	fmt.Println("*p: ", *p, "| p:", p)

	// struct field
	var person Person
	fmt.Println("person: ", person)
	person.name = "bob"
	fmt.Println("person: ", person)

	// array or slice or map element
	c := [3]int{1, 2, 3}
	scale := 5
	fmt.Println("c: ", c)
	c[x] = c[x] * scale
	fmt.Println("c: ", c)
	// each arithmetic & binary operator has an
	// assignment operator
	c[x] *= scale
	fmt.Println("c: ", c)

	// numeric variables can be incremented/decremented
	// with "++" & "--" statements
	v := 1
	fmt.Println("v: ", v)
	v++ // same as v = v + 1; v becomes 2
	fmt.Println("v: ", v)
	v-- // same as v = v - 1; v becomes 1 again
	fmt.Println("v: ", v)

}
