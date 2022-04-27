// assignability: implicit assignment
package main

import (
	"fmt"
)

func main() {
	// implicit assignment to slice
	medals := []string{"gold", "silver", "bronze"}
	fmt.Println("medals", medals)
	// explicit assignment to slice
	prize := make([]string, 3)
	prize[0] = "gold"
	prize[1] = "silver"
	prize[2] = "bronze"
	fmt.Println("prize", prize)
	// Note: maps/channels elements has such implicit assignment

	// function has implicit assignment
	fmt.Println(sum(2, 3))

}

func sum(a, b int) int {
	return a + b
}
