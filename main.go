package main

import "fmt"

func main() {
	a := 1234
	b := a
	c := &a
	d := *&a

	fmt.Println(a, "\n", b, "\n", c, "\n", d, "\n", *c)
	a = 4321
	fmt.Println(a, "\n", b, "\n", c, "\n", d, "\n", *c)
}
