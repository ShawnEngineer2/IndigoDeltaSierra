package consoleui

import (
	"fmt"
)

// This package provides text-based console menus for this simulator
func mainMenu() {

	//Prints the menu to the screen
	fmt.Println("")
	fmt.Println(cyan.Style("==============================="))
	fmt.Println(cyan.Style("IOT EVENT SIMULATOR"))
	fmt.Println(cyan.Style("==============================="))
	fmt.Println("")
	fmt.Println(white.Style("This application simulates random sensor events from a number of virtual IOT devices."))
	fmt.Println(white.Style("Please choose an option from the menu below:"))
	fmt.Println("")
	fmt.Println(green.Style("1 : Begin Simulation"))
	fmt.Println(green.Style("2 : Generate Example Events"))
	fmt.Println(green.Style("3 : Test Number Generation"))
	fmt.Println(green.Style("9 : Exit"))
	fmt.Println("")
}
