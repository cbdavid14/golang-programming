package main

import "fmt"

type personX struct {
	name string
	age  int
}

func (p personX) getAge() {
	fmt.Println("Hi, mi name is", p.name, p.age)
}

func main() {
	sum := sum(1, 2, 3, 4, 5)
	fmt.Println(sum)

	per := personX{
		name: "David",
		age:  32,
	}
	per.getAge()
}

func sum(num ...int) int {
	fmt.Println(num)

	suma := 0
	for i := range num {
		suma += i
	}
	return suma
}
