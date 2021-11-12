package main

import (
	"fmt"

	"github.com/anTuni/NomadCoin/person"
)

func main() {
	man := person.Person{}
	fmt.Println(man)
	man.SetPersonal("Mentis", 33)
	fmt.Println(man)
	man.SetPersonal("Mentissss", 33)
	fmt.Println(man)
}
