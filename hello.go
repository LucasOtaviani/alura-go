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
	
	// fmt.Scanf("%d", &command)

	fmt.Scan(&command)
	fmt.Println(command)
}
