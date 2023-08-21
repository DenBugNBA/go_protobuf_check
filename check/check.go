package main

import (
	"fmt"
	"go_protobuf_check/model"
)

func main() {
	person := model.Person{Name: "Jack"}
	fmt.Println(person.GetId())

	name := "Jack"
	personOptional := model.PersonOptional{Name: &name}
	fmt.Println(personOptional.GetId())
}
