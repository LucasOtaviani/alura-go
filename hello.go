package main

import "fmt"
import "os"

func main() {
	showIntroduction()
	showMenu()

	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Yay!")
	case 2:
		fmt.Println("Wow!")
	case 0:
		fmt.Println("Dying! :)")
		os.Exit(0)
	default:
		fmt.Println("Sorry, I'm a robot and I don't know this expression :(")
		os.Exit(-1)
	}
}

func showIntroduction() {
	name := "Lucas"
	version := 1.17

	fmt.Println("Hello, Mr.", name)
	fmt.Println("This program is running on version", version)
}

func showMenu() {
	fmt.Println("1. Iniciar Monitoramento")
	fmt.Println("2. Exibir Logs")
	fmt.Println("0. Sair do Programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}
