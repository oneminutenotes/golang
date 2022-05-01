// basic data type: integers (signed & unsigned)
// int8, int16, int32, int64, uint8, uint16, uint32, uint64
// int, uint: 32 or 64 bits (depends on compiler)
// rune == int32, rune: indicates value is unicode code point
// byte == uint8, byte: indicates value is piece of raw data
// uintptr : can hold all bits of pointer value
// eg:byte: 1111 1111;
package main

import (
	"fmt"
)

func main() {
	// % : remainder operator
	fmt.Println("% rule: ", -5%3, -5%-3, 5%-3) // sign of dividend
	// fmt.Println(5.0 % 3.0) // error: works only on integers

	// unsignedInt range(0 to 255)
	var u uint8 = 255
	fmt.Println("uint8:", u, u+1, u*u) // "255 0 1" overflow

	// signedInt range(-128 to 127)
	var i int8 = 127
	fmt.Println("int8:", i, i+1, i*i) // "127 -128 1" overflow

	// comparison: ==, !=, <, <=, >, >=,
	var p int32 = 14
	var q int32 = 31
	fmt.Println("compare:", p == q)

	// unary operator: +, -
	fmt.Println("unary:", +3, -3)

	// bitwise binary operator: &, |, ^, &^, <<, >>
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("x: %08b\n", x)       // "00100010", the set {1, 5}
	fmt.Printf("y: %08b\n", y)       // "00000110", the set {1, 2}
	fmt.Printf("x&y: %08b\n", x&y)   // "00000010", the intersection {1}
	fmt.Printf("x|y: %08b\n", x|y)   // "00100110", the union {1, 2, 5}
	fmt.Printf("x^y: %08b\n", x^y)   // "00100100", the symmetric difference {2, 5}
	fmt.Printf("x&^y: %08b\n", x&^y) // "00100000", the difference {5}
	fmt.Printf("members: {")
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Printf("%d,", i) // "1", "5"
		}
	}
	fmt.Printf("}\n")
	// << == multiplication of 2^n, >> == division by 2^n (unsigned)
	fmt.Printf("lShift x: %08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("rShift x: %08b\n", x>>1) // "00010001", the set {0, 4}

	// len() returns signed int
	browser := []string{"safari", "chrome", "brave"}
	for i := len(browser) - 1; i >= 0; i-- {
		fmt.Printf("%v ", browser[i]) // "brave", "chrome", "safari"
	}

	// explicit conversion for different types
	var apples int32 = 1
	var oranges int16 = 2
	// var compote int = apples + oranges // compile error
	var compote = int(apples) + int(oranges)
	fmt.Println("\nconversion:", compote)

	// lose precision when convert from float to int
	f := 3.141 // a float64
	j := int(f)
	fmt.Println("losePrecision:", f, j) // "3.141 3"

	// operand is out of range of target type
	fl := 1e100                        // a float64
	rs := int(f)                       // result is implementation-dependent
	fmt.Println("outOfRange:", fl, rs) // 1e+100 1

	// intergers can be written as number, octa, hex forms
	m := 0666
	fmt.Printf("octa: %d %[1]o %#[1]o\n", m) // "438 666 0666"
	n := int64(0xdeadbeef)
	fmt.Printf("hex: %d %[1]x %#[1]x %#[1]X\n", n)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

	// rune literals
	ascii := 'a'
	unicode := 'D'
	newline := '\n'
	fmt.Printf("rune: %d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("rune: %d %[1]c %[1]q\n", unicode) // "68 D 'D'"
	fmt.Printf("rune: %d %[1]q\n", newline)       // "10 '\n'"
}
