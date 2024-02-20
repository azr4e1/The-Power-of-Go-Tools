package main

import (
	"os"
	"piping"
)

func main() {
	// This is what we want to write
	piping.WalkFiles(os.DirFS(".")).Grep("\\.go$").Concat().Lines().Stdout()
}
