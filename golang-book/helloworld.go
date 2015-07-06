package main
// This is a "package declaration", these are a requirement for every program written in Go. Packages are Go's way of organizing and reusing code. Needs to be called main for this to be an executable.

import "fmt"
// import is used to include code from other packages to use with the current program. In this case fmt (short for format) implements formatting for input and output.

func main() {
// The above is a function declaration and probably makes sense based on the other languages we've used. This declaration can include a return type as well.
// The function name of "main" is important because it is the function that runs by default when it is run.
  var name string = "Mike"
  // Variable declaration looks pretty familiar although you define the type as well.
  fmt.Println("Hello, my name is " + name)
  // accessing a "Println" method from the fmt library to print out "Hello World" at the console.
}