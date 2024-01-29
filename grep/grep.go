package grep

import "io"

type finder struct {
	input  io.Reader
	output io.Writer
}

type option func(*finder) error

func NewFinder(opt ...option) finder {

}
