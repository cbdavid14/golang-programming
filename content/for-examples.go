package main

import "fmt"

func main() {
	// Array
	var x [6]int
	x[0] = 1
	x[1] = 1
	x[2] = 1
	x[3] = 1
	x[4] = 1
	x[5] = 1

	for i, value := range x {
		fmt.Println(i, "---", value)
	}

	fmt.Print(x)
	fmt.Print(len(x))
	fmt.Printf("%T \n", x)

	//Slice -> Puntero a un array, longitud y capacidad
	sliceInt := []int{1, 2, 3, 4, 5}
	for i, value := range sliceInt {
		fmt.Println(i, "---", value)
	}

	// Add elements in slice
	sliceInt = append(sliceInt, 10)
	fmt.Println(sliceInt[1])
	fmt.Print(len(sliceInt))
	fmt.Printf("%T\n", sliceInt)

	sliceInt = append(sliceInt, 10)
	fmt.Println(sliceInt[1])
	fmt.Print(len(sliceInt))
	fmt.Printf("%T\n", sliceInt)

	sliceX := x[:]
	sliceInt = append(sliceInt, sliceX...)
	fmt.Printf("lista de %v\n", sliceInt)

	//Copy slice
	source := []string{"a", "b", "c"}
	destine := make([]string, len(source))
	copy(destine, source)
	fmt.Printf("lista source %v\n", source)
	fmt.Printf("lista destine %v\n", destine)

	destine = append(destine, "q", "t")
	fmt.Printf("lista destine %v\n", destine)

	//Slice multidimensionales
	student1 := []string{"david", "briceira", "felipe"}
	fmt.Println(student1)
	students2 := []string{"david2", "briceira2", "felipe2"}
	fmt.Println(students2)

	multi := [][]string{students2, student1}
	fmt.Println(multi)

	//Maps
	mapStudent := map[string]int{
		"Zolia":    20,
		"david":    32,
		"briceira": 31,
	}
	fmt.Println(mapStudent)
	fmt.Println(mapStudent["david"])
	v, ok := mapStudent["davi"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := mapStudent["briceira"]; ok {
		fmt.Println(v)
	}

	delete(mapStudent, "david")

	mapStudent["todo"] = 100
	for k, v := range mapStudent {
		fmt.Println(k, "---", v)
	}
}
