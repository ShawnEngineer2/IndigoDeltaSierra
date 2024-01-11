package main

import (
	"fmt"
	"indigodeltasierra/SvcClient"
)

func main() {
	fmt.Println("Hello World")
	SvcClient.GetRandomNumbers(16, 1, 6)
}
