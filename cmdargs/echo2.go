package main

import (
  "fmt"
  "os"
)

// os.Args
// os.Args[0] - commands itself
// os.Args[1:]

func main(){
  s := ""
  sep := ""

  for _, val := range os.Args[1:]{
    s += sep + val
    sep = " "
  }

  fmt.Println(s)
}
