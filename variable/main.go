package main

// you can define variable with var or :=
import (
	"fmt"
	"reflect"
)

var name string
var age float64 = 5.2
var lastName string

var (
	age1             int = 122
	name1, lastName1 string
)

func main() {
	// var family string
	name = "amir"
	lastName = "rahmati"
	location := "Tehran"
	print(age1)
	fmt.Print(reflect.TypeOf(age), reflect.TypeOf(location))
}
