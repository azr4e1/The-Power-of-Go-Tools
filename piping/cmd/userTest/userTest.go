package main

import (
	"piping"
)

func main() {
	piping.FromFile("../../testdata/log.txt").Column(1).Sort(false).Freq().Sort(true).First(10).Stdout()
}
