// short form : name := expression
package main

import (
	"fmt"
	"os"
)

type Point struct {}

func main() {
	// short form
	c := 1
	fmt.Printf("c: %v,%T \n", c, c)

	// use var where type needed is different from value
	var b float64 = 100
	fmt.Printf("b: %v,%T \n", b, b)

	// use var where initial value is not important
	// or if you want to initialize value at later time
	var names []string
	fmt.Printf("names: %v,%T \n", names, names)
	var err error
	fmt.Printf("err: %v,%T \n", err, err)
	var p Point
	fmt.Printf("p: %v,%T \n", p, p)

	// ':=' is declaration
	// initialize multiple values
	i, j := 0, 1

	// '=' is assignment
	// this is tuple assignment not initialization
	i, j = j, i
	
	// initialize set by calling function with multiple return values
	f, err := os.Open("notes.txt") // os.Open returns file,error
	if err != nil {
		fmt.Println("error while opening file")
	}
	f.Close()

	// := does not necessarily declare all variables on left
	// i.e if variable exist it will act as an assignment
	out, err := os.Create("outfile.txt")
	out.Close()

	// fix '=' if both variables are declared
	f, err = os.Create("outfile.txt")
	f.Close()


}