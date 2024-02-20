package main

import (
	"flag"
	"fmt"
	"indigodeltasierra/consoleui"
	"indigodeltasierra/simulator"
)

func main() {

	//Define the "silent" flag. If the "silent" flag was not sent in on the command line, show a menu and process based on the input
	silentPtr := flag.Bool("silent", false, "When specified, causes the simulator to start without displaying the user menu")

	flag.Parse()

	if *silentPtr {
		fmt.Print("Starting IOT event simulation ...")
		simulator.StartSimulation()
	} else {
		//Run interactively
		consoleui.RunInteractive()
	}
}
