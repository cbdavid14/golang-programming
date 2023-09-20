package main

import (
	"fmt"
	"strconv"
)

var a int

type money int

var b money

func main() {
	fmt.Println("Why go?")

	fmt.Println("Aspectos mas conocidos: Concurrencia integrada y recoleccion de basura")

	x := 45 + 9
	fmt.Println(x)
	foo(10)

	a = 109
	b = 800
	fmt.Println(b, a)
	fmt.Printf("%T\n", b)

	var t interface{}
	t = b
	switch i := t.(type) {
	default:
		fmt.Print("nil type", i)
	case money:
		fmt.Println("money type")
	case int:
		fmt.Println("Integer type")
	}

	z := strconv.Itoa(int(b))
	fmt.Println(z)
	fmt.Printf("%T", z)

}

func foo(number int) {
	name := fmt.Sprintf("numero %d de la casa", number)
	fmt.Println(name)
}
