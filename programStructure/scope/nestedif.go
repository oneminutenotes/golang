package main

import (
	"fmt"
	"os"
)

func main() {
	// demo1()
	// demo2()
	demo3()
}

func demo3() error {
	fmt.Println("demo3()")
	if f, err := os.Open("temp.txt"); err != nil { 
		// compile error: unused: f
		return err
	} else {
		// f and err are visible here too
		f.Close()
	}
	return nil
}

func demo2() error {
	fmt.Println("demo2()")
	if f, err := os.Open("temp.txt"); err != nil { // compile error: unused: f
		f.Close()
		return err
	}
	// f.Close() // compile error: undefined f
	return nil
}

func f() int      { return 1 }
func g(x int) int { return x + 1 }
func demo1() {
	fmt.Println("demo1()")
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	// fmt.Println(x, y) // compile error: x and y are not visible here
}
