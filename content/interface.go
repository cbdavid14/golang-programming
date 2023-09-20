package main

import "fmt"

type personAgent struct {
	name string
	age  int
}
type agentSecretX struct {
	personAgent
	lpm bool
}

type human interface {
	introduceOneSelf()
}

func (s agentSecretX) introduceOneSelf() {
	fmt.Println("Hi iam", s.name, "have", s.age, "iam agent secret")
}

func (s personAgent) introduceOneSelf() {
	fmt.Println("Hi iam", s.name, "have", s.age, "iam person")
}

func bar(h human) {
	fmt.Println("Fue pasado a la funcion bar", h)
	h.introduceOneSelf()
}

func main() {

	as1 := agentSecretX{
		personAgent: personAgent{
			name: "Juan",
			age:  100,
		},
		lpm: true,
	}
	per := personAgent{
		name: "david",
		age:  32,
	}
	bar(per)
	bar(as1)
}
