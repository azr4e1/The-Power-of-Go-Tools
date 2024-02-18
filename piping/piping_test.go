package piping_test

import (
	"bytes"
	"piping"
	"testing"
)

func TestStdoutMethodOutputsToBuffer(t *testing.T) {
	t.Parallel()
	out := new(bytes.Buffer)
	want := "This is the content of the pipeline."
	p := piping.FromString(want)
	p.Output = out
	p.Stdout()
	if p.Error != nil {
		t.Fatal(p.Error)
	}
	got := out.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
