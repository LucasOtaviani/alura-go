package main

import "fmt"

func main() {
	name := "Lucas"
	version := 1.17

	fmt.Println("Hello, Mr.", name)
	fmt.Println("This program is running on version", version)

	fmt.Println("1. Iniciar Monitoramento")
	fmt.Println("2. Exibir Logs")
	fmt.Println("0. Sair do Programa")

	var command int

	fmt.Scan(&command)

	if command == 1 {
		fmt.Println("Yay!")
	} else if command == 2 {
		fmt.Println("Wow!")
	} else if command == 0 {
		fmt.Println("Dying! :)")
	} else {
		fmt.Println("Sorry, I'm a robot and I don't know this expression :(")
	}
}
