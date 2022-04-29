package main

import "fmt"

func main() {
	demo1()
	demo2()
}

func demo2() {
	fmt.Println("\ndemo2()")
	x := "hello!"
	for _, x := range x {
		fmt.Printf("%c\t", x)
		x := x + 'A' - 'a'
		fmt.Printf("%c\n", x) // "HELLO" (one letter per iteration)
	}
}

func demo1() {
	fmt.Println("\ndemo1()")
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			fmt.Printf("%c\t", x)
			x := x + 'A' - 'a'
			fmt.Printf("%c\n", x) // "HELLO" (one letter per iteration)
		}
	}
}
