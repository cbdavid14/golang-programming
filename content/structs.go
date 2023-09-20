package main

import "fmt"

type person struct {
	name    string
	surname string
}

type agentSecret struct {
	person
	lpm bool
}

func main() {
	//Tipo de dato compuesto

	per := person{
		name:    "david",
		surname: "consa",
	}
	fmt.Printf("%v\n", per)

	as := agentSecret{
		person: per,
		lpm:    true,
	}
	fmt.Println(as)
	fmt.Println(as.person.name)
	fmt.Println(as.name)

	//structs anonimos
	p1 := struct {
		name    string
		surname string
		age     int
	}{
		name:    "david",
		surname: "consa",
		age:     32,
	}
	fmt.Println(p1)
}
