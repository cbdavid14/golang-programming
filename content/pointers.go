package main

import "fmt"

type student struct {
	name    string
	age     int
	isAdult bool
	nota    int
}

// with pointer allow update field isAdult
// without pointer not allow update field isAdult
func (s *student) getAgeStudent() {
	if s.age > 18 {
		s.isAdult = true
	}
	fmt.Println("He student", s.name, "have", s.age, "and is", s.isAdult)
}

func main() {

	x := 100
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	st := student{
		name: "david",
		age:  20,
	}
	st.getAgeStudent()
	fmt.Println(st)

	updateNote(&st)
	fmt.Println(st)

}

func updateNote(s *student) {
	s.nota = 20
}
