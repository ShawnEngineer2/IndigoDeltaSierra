package consoleui

import (
	"bufio"
	"fmt"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/simulator"
	"indigodeltasierra/testutil"
	"os"
)

//This package provides a REPL loop for the menu system of this simlator

// Constants used throughout this module

func RunInteractive() {
	//Call the main menu routine and wait for user input

	reader := bufio.NewReader(os.Stdin)

	for {
		mainMenu()
		fmt.Print(cyan.Style("-> "))
		inputChar, _, _ := reader.ReadRune()

		switch inputChar {
		case '1':
			//startSimulation()
			fmt.Println("")
			simulator.StartSimulation()
			pressEnterToContinue()

		case '2':
			//generateTestEvents()
			fmt.Println("")
			testutil.GenerateTestEvents()
			pressEnterToContinue()

		case '9':
			goByeBye()
		}
	}

}

func pressEnterToContinue() {
	//This is just a cheesy way to avoid rewriting stuff over and over
	fmt.Println("")
	fmt.Println(cyan.Style("Press ENTER to continue ..."))
	fmt.Scanln()
}

func goByeBye() {
	//This prints the graceful shutdown message and exits the program
	fmt.Println("")
	fmt.Println(cyan.Style("Program Exit - Good-bye"))
	fmt.Println("")
	os.Exit(appconstants.EXIT_CLEAN)
}
