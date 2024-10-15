package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = 40
	var b interface{} = "cat"
	var c interface{} = make(chan int)

	fmt.Println(GetType(a))
	fmt.Println(GetType(b))
	fmt.Println(GetType(c))
}

// Для описанных типов
func GetType2(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	}
	return "unknown"
}

// Для любых типов
func GetType(v interface{}) string {
	t := reflect.TypeOf(v)
	return t.String()
}
