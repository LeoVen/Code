package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	id      uint64
	name    string `about:"Full Name"`
	age     int    `about:"Current age or when person died"`
	isAlive bool   `about:"If the individual is alive or not"`
	cache   bool   `about:""`
}

func main() {
	p1 := Person{name: "My Name", age: 92, isAlive: false}

	t := reflect.TypeOf(p1)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		about, ok := f.Tag.Lookup("about")
		if !ok {
			about = "[Missing]"
		} else if about == "" {
			about = "[Empty]"
		}
		fmt.Printf("Field name: %s\n", f.Name)
		fmt.Printf("Field metadata: %s\n", about)
	}
}
