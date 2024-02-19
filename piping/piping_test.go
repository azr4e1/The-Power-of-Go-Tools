package piping_test

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
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

func TestStdoutMethodTreatsErrorsSafely(t *testing.T) {
	t.Parallel()
	out := new(bytes.Buffer)
	want := ""
	p := piping.FromString("This is the content of the pipeline.")
	p.Output = out
	p.Error = errors.New("")
	p.Stdout()
	got := out.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestFromFile_ReadsCorrectlyFromAFile(t *testing.T) {
	t.Parallel()
	out := new(bytes.Buffer)
	filename := "testdata/test.txt"
	fsys := os.DirFS(".")
	wantByte, err := fs.ReadFile(fsys, filename)
	want := string(wantByte)
	if err != nil {
		t.Fatal(err)
	}
	p := piping.FromFile(filename)
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

func TestFromFile_HandlesErrorsCorrectly(t *testing.T) {
	t.Parallel()
	filename := "bogus"
	p := piping.FromFile(filename)

	if p.Error == nil {
		t.Fatal("want error opening non-existent file, got nil")
	}
}
