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

	// optional fields
	personOptional := model.PersonOptional{Name: &name}
	fmt.Println(personOptional.GetId()) //0

	// oneof
	some := model.SomeValue{SomeField: &model.SomeValue_Str{Str: "hello"}}
	fmt.Println(some.GetStr()) // hello
}
