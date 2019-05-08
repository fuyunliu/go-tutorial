package main

import "fmt"

func main() {
	a := 3
	b := 3.14
	c := 0.867 + 0.5i
	fmt.Printf("%v is of type %T\n", a, a)
	fmt.Printf("%v is of type %T\n", b, b)
	fmt.Printf("%v is of type %T\n", c, c)
}
