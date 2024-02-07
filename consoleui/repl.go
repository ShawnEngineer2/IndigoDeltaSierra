package consoleui

import (
	"bufio"
	"fmt"
	"indigodeltasierra/appconstants"
	"os"
)

//This package provides a REPL loop for the menu system of this simlator

// Constants used throughout this module

func RunInteractive() {
	//Call the main menu routine and wait for user input
	mainMenu()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		inputChar, _, _ := reader.ReadRune()

		switch inputChar {
		case '1':
			//startSimulation()
			fmt.Println("")
			os.Exit(appconstants.EXIT_CLEAN)

		case '2':
			//generateTestEvents()
			fmt.Println("")
			os.Exit(appconstants.EXIT_CLEAN)

		case '9':
			fmt.Println("")
			fmt.Println("Program Exit - Good-bye")
			fmt.Println("")
			os.Exit(appconstants.EXIT_CLEAN)
		}
	}

}
