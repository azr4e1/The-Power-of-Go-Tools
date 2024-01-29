package main

import (
	"fmt"
	"linecounter"
)

func main() {
	count := linecounter.Main()
	fmt.Println("---------------\nNumber of lines:", count)
}
