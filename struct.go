package main

import "fmt"

type person struct {
	name string
	age  int
}

func newperson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}
func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{"Alice", 30})
	fmt.Println(person{"Ram", 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isgood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}
