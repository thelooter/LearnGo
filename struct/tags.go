package main

import "reflect"

type animal struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(animal{})
	field, _ := t.FieldByName("Name")
	println(field.Tag)
}
