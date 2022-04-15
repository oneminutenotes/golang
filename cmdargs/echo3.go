package main

import (
  "fmt"
  "os"
  "strings"
)

// os.Args
// os.Args[0] - commands itself
// os.Args[1:]

func main(){
  fmt.Println(strings.Join(os.Args[1:], " "))
}
