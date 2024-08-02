package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	welcomeIntro()
	for {
		menuOption()
		option := readMenuOption()

		switch option {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println(">>>Logging")
		case 0:
			fmt.Println(">>>Ending Program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid Option: ", option)
			os.Exit(-1)
		}
	}
}

func welcomeIntro() {
	//variables
	name := "Poncio Costa"
	versionSystem := 1.2
	//
	fmt.Println("Hello, ", name)
	fmt.Print(">> Version program: ", versionSystem, " <<")
}

func menuOption() {
	fmt.Println("\n<<< MENU OPTIONS >>>")
	fmt.Println("1- Start Monitoring Program")
	fmt.Println("2- Show logs")
	fmt.Println("0- End Program")
}

func readMenuOption() int {
	var optionSelected int
	fmt.Scan(&optionSelected)
	fmt.Println("Menu option selected: ", optionSelected)
	return optionSelected
}

func startMonitoring() {
	fmt.Println(">>>Monitoring...")
	website := "https://www.alura.com.br/"
	response, _ := http.Get(website)

	if response.StatusCode == 200 {
		fmt.Println("Website:", website, "was reached! ")
	} else {
		fmt.Println("Website:", website, "with problems. Status Code:", response.StatusCode)
	}
}
