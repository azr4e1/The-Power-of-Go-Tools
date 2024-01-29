package main

import (
	"errors"
	"fmt"
	"hellopeople"
	"os"
)

func main() {
	input := os.Stdin
	output := os.Stdout

	fmt.Print("What is your name?: ")
	name, err := hellopeople.ReadFrom(input)
	if err != nil {
		errors.New("An error occured")
	}
	hellopeople.PrintTo(output, name)
}
