package hellopeople_test

import (
	"bytes"
	"hellopeople"
	"testing"
)

func TestGetNameFromStandardInput(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	want := "Lorenzo\n"

	buf.WriteString(want)
	got, err := hellopeople.ReadFrom(buf)
	if err != nil {
		t.Fatalf("expected nil error")
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestPrintTo_TOStandardOutput(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	name := "Lorenzo"
	want := "Hello, Lorenzo\n"

	hellopeople.PrintTo(buf, name)
	got := buf.String()

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
