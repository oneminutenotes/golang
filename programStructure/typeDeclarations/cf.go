// type general form: type name underlying-type
package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// %v is used to format current type values
// %s is used to format string values
// %g is used to format float values

func main() {

	// underlying type defines intrinsic operation
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	// fmt.Printf("%g\n", boilingF-FreezingC) // compile error: type mismatch

	// Comparison operators like == and < can be used
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	// fmt.Println(c == f) // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // "true"!

	// type's methods defines new behaviour of type
	c = FToC(212.0)
	fmt.Println(c.String()) // "100°C"; call String method
	fmt.Printf("%v\n", c)   // "100°C"; call String method
	fmt.Printf("%s\n", c)   // "100°C"; call String method
	fmt.Println(c)          // "100°C"; call String method
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String

}
