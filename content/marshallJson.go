package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	//Marshall
	type colorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := colorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"red", "rubi", "Maroon"},
	}
	marshal, err := json.Marshal(group)
	if err != nil {
		fmt.Println(marshal)
	}
	os.Stdout.Write(marshal)

	//UnMarshall
	var jsonBlob = []byte(`{"ID":100,"Name":"Reds-100","Colors":["red","rubi","Maroon"]}`)
	var colors colorGroup
	err = json.Unmarshal(jsonBlob, &colors)

	if err == nil {
		fmt.Println("Error al realizar el unMarshall")
	}
	fmt.Printf("%+v\n", colors)

	//Ej 2
	s := `[{"Nombre":"James","Apellido":"Bond","Edad":32},{"Nombre":"Miss","Apellido":"Moneypenny","Edad":27}]`
	bs := []byte(s)
	var personas []persona
	err = json.Unmarshal(bs, &personas)
	if err != nil {
		fmt.Println(err)
	}
}
