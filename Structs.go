package main

import "fmt"

type person struct {
	name string
	age int
	height int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{name: "Bob", age: 20, height: 128})
	fmt.Println(person{name: "Alice", age: 30, height: 158})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon Snow"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	fmt.Println(s.age)

	sp := &s
	sp.age = 51
	fmt.Println(sp.age)
}