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

// Defining nested maps is a bit harder because they need to be defined when they are created. The nested maps also need to be defined when they are made also. The nested maps need to be declared as such while the outer map is being defined as well as being defined when created as a nested map.
// elements := map[string]map[string]string{
  // "H": map[string]string{
  //   "name":"Hydrogen",
  //   "state":"gas",
  // },
  // "He": map[string]string{
  //   "name":"Helium",
  //   "state":"gas",
  // }
// }