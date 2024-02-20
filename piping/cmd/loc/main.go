package main

import (
	"fmt"
	"os"
	"piping"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}
	root := args[1]
	// This is what we want to write
	// piping.WalkFiles(os.DirFS(root)).Grep("\\.py$").Stdout()
	count, err := piping.WalkFiles(os.DirFS(root)).Grep("\\.go$").Concat(os.DirFS(root)).Grep("[^\\n]").Lines()
	if err != nil {
		panic(err)
	}

	fmt.Printf("You've written %d lines of Go in this project. Nice work!\n", count)
}
