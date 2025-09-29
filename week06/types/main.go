package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var name string
	//var id int

	//name = "Kim Inha"
	//id = 1000

	name := "Kim Inha"
	id := 1000
	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(id, reflect.TypeOf(id))
}
