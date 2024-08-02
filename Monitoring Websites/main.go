package main

import (
	"fmt"
	"reflect"
)

func main() {
	//variables
	var name string = "Poncio Costa"
	age := 24
	versionSystem := 1.0
	//
	fmt.Println("Hello World!!	>>> Welcome , Mr", name, "<<<	#Age:", age, " #Programn version:", versionSystem)
	fmt.Print("var name typeOf: ", reflect.TypeOf(name))
	fmt.Print(" >> var age typeOf: ", reflect.TypeOf(age))
	fmt.Print(" >> var version typeOf: ", reflect.TypeOf(versionSystem))

	fmt.Println("\n\n1- Start Monitoring Program")
	fmt.Println("2- Show logs")
	fmt.Println("0- End Program")

	var menu int
	//fmt.Scanf("%d", &menu)
	fmt.Scan(&menu)
	fmt.Println("Menu option selected: ", menu)

	//
	if menu == 1 {
		println("## Start Monitoring... ##")
	} else if menu == 2 {
		println("## Show logs... ##")
	} else if menu == 0 {
		println("## End program... ##")
	} else {
		fmt.Println("Menu option not found!")
	}
}
