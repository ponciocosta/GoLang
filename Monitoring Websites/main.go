package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitorRepeat = 5
const monitorDelayTime = 5

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
	fmt.Println("")
	return optionSelected
}

func startMonitoring() {
	websites := []string{"https://httpbin.org/status/200", "https://httpbin.org/status/400", "https://httpbin.org/status/500"}

	for i := 0; i < monitorRepeat; i++ {
		fmt.Println(">>>Monitoring...")
		for index, website := range websites {
			fmt.Println("Testing website", index, ":", website)
			runWebsiteTest(website)
		}
		time.Sleep(monitorDelayTime * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func runWebsiteTest(route string) {
	response, _ := http.Get(route)
	if response.StatusCode == 200 {
		fmt.Println("Website:", route, "was reached! Status Code:", response.StatusCode)
	} else {
		fmt.Println("Website:", route, "with problems. Status Code:", response.StatusCode)
	}
}
