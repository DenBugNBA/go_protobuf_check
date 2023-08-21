package main

import (
	"fmt"
	"go_protobuf_check/model"
)

func main() {
	person := model.Person{Name: "Jack"}
	fmt.Println(person.GetId()) // 0

	person.Id = 10
	fmt.Println(person.GetId()) // 10

	name := "Jack"

	personOptional := model.PersonOptional{Name: &name}
	fmt.Println(personOptional.GetId()) //0
}
