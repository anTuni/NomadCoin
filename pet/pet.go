package pet

import "fmt"

type Pet struct {
	name string
	age  int
	root []string
}

func (i *Pet) SetName(name string, age int, root []string) {
	i.age = age
	i.name = name
	i.root = root

	for _, el := range root {
		fmt.Print(el)
	}
}
