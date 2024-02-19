package main

import "piping"

func main() {
	piping.FromString("hello, world\n").Stdout()
}
