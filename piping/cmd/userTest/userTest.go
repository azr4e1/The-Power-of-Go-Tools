package main

import (
	"fmt"
	"piping"
)

func main() {
	piping.FromString("hello, world\n").Stdout()
	fmt.Println("1111" < "b")
}
