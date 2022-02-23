package main

import "fmt"
import "reflect"

func main() {
	name := "Lucas"
	age := 20
	version := 1.17

	/* Works calling 'var' */
	// var name string = "Lucas"
	// var age int = 20
	// var version float32 = 1.17

	/* Works without a type */
	// var name = "Lucas"
	// var age = 20
	// var version = 1.17

	fmt.Println("Hello World!")
	fmt.Println("Hello, Mr.", name, "your age is", age)
	fmt.Println("This program is running on version", version)

	fmt.Println("Type of name variable is:", reflect.TypeOf(name))
	fmt.Println("Type of name variable is:", reflect.TypeOf(age))
	fmt.Println("Type of name variable is:", reflect.TypeOf(version))
}
