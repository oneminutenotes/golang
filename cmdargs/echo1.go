package main

import (
  "fmt"
  "os"
)

// os.Args
// os.Args[0] - commands itself
// os.Args[1:n-1] - n: length of os.args

func main(){
  var s, sep string
  for i:=1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
}
