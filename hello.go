package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitorings = 3
const delay = 5

func main() {
	showIntroduction()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			getLogs()
		case 0:
			fmt.Println("Dying! :)")
			os.Exit(0)
		default:
			fmt.Println("Sorry, I'm a robot and I don't know this expression :(")
			os.Exit(-1)
		}
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

func startMonitoring() {
	websites := readWebsiteFile()

	for i := 0; i < monitorings; i++ {
		for _, website := range websites {
			testSite(website)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func getLogs() {

}

func testSite(website string) {
	response, error := http.Get(website)

	if error != nil {
		fmt.Println(error)
		return
	}

	if response.StatusCode == 200 {
		fmt.Println("Website:", website, "was loaded successfully")
	} else {
		fmt.Println("Something went wrong. Status code:", response.StatusCode)
	}
}

func readWebsiteFile() []string {
	var sites []string

	file, error := os.Open("websites.txt")

	if error != nil {
		fmt.Println(error)
	}

	reader := bufio.NewReader(file)

	for {
		line, error := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if error == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}
