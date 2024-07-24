package main

import (
	"fmt"
	"reflect"
)

func isInt(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Int
}

func main() {
	number := 10
	text := "Hello"

	fmt.Println(isInt(number))
	fmt.Println(isInt(text))
}
