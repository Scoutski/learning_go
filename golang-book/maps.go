package main

import "fmt"

func main() {
  x := make(map[string]int)
  // The data types for the different key value pairs within a map must be set when it is defined.
  x["key"] = 10
  // additional key value pairs can be added once a map has been defined.
  fmt.Println(x)
  delete(x, "key")
  // key value pairs can be deleted with the built in delete function.
  fmt.Println(x)
}