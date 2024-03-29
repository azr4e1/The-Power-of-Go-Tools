package hello_test

import (
	"bytes"
	"testing"

	"github.com/azr4e1/hello"
)

func TestPrintTo_PrintsHelloMessageToGivenWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	printer := hello.NewPrinter()
	printer.Output = buf
	printer.Print()
	want := "Hello, world\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
