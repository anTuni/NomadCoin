package person

type Person struct {
	name      string
	age       int
	koreanAge int
}

func (man *Person) SetPersonal(name string, age int) {
	man.name = name
	man.age = age
	man.koreanAge = age + 1
}
