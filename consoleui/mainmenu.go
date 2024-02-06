package consoleui

import (
	"fmt"
)

// This package provides text-based console menus for this simulator
func mainMenu() {
	//Prints the menu to the screen
	fmt.Println("")
	fmt.Println("===============================")
	fmt.Println("IOT EVENT SIMULATOR")
	fmt.Println("===============================")
	fmt.Println("")
	fmt.Println("This application simulates random sensor events from a number of virtual IOT devices.")
	fmt.Println("Please choose an option from the menu below:")
	fmt.Println("")
	fmt.Println("1 : Begin Simulation")
	fmt.Println("2 : Generate Example Events")
	fmt.Println("9 : Exit")
	fmt.Println("")
}
