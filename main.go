package main

import (
	"fmt"

	"rsc.io/quote"
)

func plus(a, b int, name string) (int, string) {
	return a + b, name
}

func sum(a ...int) int {
	var total int
	for _, el := range a {
		total += el
	}
	return total
}
func main() {
	fmt.Println(quote.Go())
	result, name := plus(1, 2, "Ant")
	fmt.Println(result)
	fmt.Println(name)

	sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 1100)
	fmt.Println(sum)
}
