// general form : var name type = expression
package main

import (
	"fmt"
	"os"
)

func main() {

	// general form
	var c int = 1
	fmt.Printf("c: %v,%T \n", c, c)

	// expression ommited, default value is zero of type
	var s string
	fmt.Printf("s: %v,%T \n", s, s)

	// type ommited, type is determined by expression
	var o = "oneMinuteNotes"
	fmt.Printf("o: %v,%T \n", o, o)

	// initialize set of variable of same type at once
	var i, j, k int // int, int, int
	fmt.Printf("i:%v,%T, j:%v,%T, k:%v,%T \n", i, i, j, j, k, k)

	// initialze set of variable of different type at once
	var bo, fl, st = true, 3.3, "one" // bool, float64, string
	fmt.Printf("bo:%v,%T, fl:%v,%T, st:%v,%T \n", bo, bo, fl, fl, st, st)

	// initialize set by calling function with multiple return values
	var f, err = os.Open("notes.txt") // os.Open returns file,error
	if err != nil {
		fmt.Println("error while opening file")
	}
	f.Close()
}
