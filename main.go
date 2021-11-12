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

	foods := [3]string{"kimchi", "sam", "chicken"}

	for _, food := range foods {
		fmt.Println(food)
	}

	people := []string{"Antuni", "JJM", "Beast"}

	for i := 0; i < len(people); i++ {
		fmt.Println(people[i])
	}
	more := append(people, "JK")

	println(more[0], more[1], more[2], more[3])
	println(len(more))
	println(len(people))
}
