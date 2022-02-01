package main

import (
	"fmt"
	"reflect"
)

type Metadata struct {
	FieldName string            `name:"FieldName" type:"string" about:"The mapped field name" json:"field_name,omitempty"`
	Values    map[string]string `name:"Values" type:"map[string]string" about:"A map keyed by the tag name, referencing its text value" json:"values,omitempty"`
}

func getMetadata(keys []string, value reflect.Type) []Metadata {
	result := []Metadata{}
	t := reflect.TypeOf(value)

	i := 0

	for ; i < t.NumField(); i++ {
		f := t.Field(i)
		meta := Metadata{FieldName: f.Name, Values: make(map[string]string)}
		result = append(result, meta)

		for _, key := range keys {
			value, ok := f.Tag.Lookup(key)

			if ok {
				meta.Values[key] = value
			}
		}
	}

	return result
}

func main() {
	metas := getMetadata([]string{"about", "name", "type"}, reflect.TypeOf(Metadata{}))

	println(len(metas), " Results")

	for _, meta := range metas {
		fmt.Println("------")
		for k, v := range meta.Values {
			fmt.Printf("%s: %s\n", k, v)
		}
	}
}
