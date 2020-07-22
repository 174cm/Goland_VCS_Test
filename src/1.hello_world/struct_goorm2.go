package main

import "fmt"

type person2 struct {
	name, contact string
	age           int
}

func addAgeRef(a *person2) {
	a.age += 4
}

func addAgeVal(a person2) {
	a.age += 4
}

func main() {
	var p1 = new(person2)
	var p2 = person2{}

	fmt.Println(p1, p2)

	p1.age = 25
	p2.age = 25

	addAgeRef(p1)
	addAgeVal(p2)

	fmt.Println(p1.age)
	fmt.Println(p2.age)

	addAgeRef(&p2)
	fmt.Println(p2.age)
}
