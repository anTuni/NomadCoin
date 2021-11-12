package main

import "fmt"

type person struct {
	name      string
	age       int
	KoreanAge int
}

func (p person) introduce() {
	fmt.Printf("hello My name is %s and I'm %d specially %d in kirea", p.name, p.age, p.KoreanAge)
}

func main() {
	man := person{name: "Ant", KoreanAge: 24, age: 22}
	man.introduce()
}
