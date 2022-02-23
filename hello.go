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

	switch command {
	case 1:
		fmt.Println("Yay!")
	case 2:
		fmt.Println("Wow!")
	case 0:
		fmt.Println("Dying! :)")
	default:
		fmt.Println("Sorry, I'm a robot and I don't know this expression :(")
	}
}
