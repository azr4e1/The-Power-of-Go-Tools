package main

import (
	"bufio"
	"fmt"
	"os"
	"shell"
)

func main() {
	fmt.Print("> ")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		cmd, err := shell.CmdFromString(input.Text())
		if err != nil {
			continue
		}
		data, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%s", data)
		fmt.Print("\n> ")
	}
	fmt.Println("\nBe seeing you!")
}
