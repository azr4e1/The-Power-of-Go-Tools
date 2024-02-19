package piping_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"piping"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestColumnSelectsTheCorrectColumn(t *testing.T) {
	t.Parallel()
	data := "1 2 3\n1 2 3\n1 2 3\n"
	want := "2\n2\n2\n"

	p := piping.FromString(data)

	got, err := p.Column(2).String()
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestColumnProducesNothingWhenPipeErrorSet(t *testing.T) {
	t.Parallel()
	input := "1 2 3\n1 2 3\n1 2 3\n"

	p := piping.FromString(input)
	p.Error = errors.New("oh no")

	data, err := io.ReadAll(p.Column(2).Data)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) > 0 {
		t.Errorf("want no output from Column after error, but got %q", data)
	}
}

func TestColumnSetsErrorAndProducesNothingGivenInvalidArg(t *testing.T) {
	t.Parallel()
	p := piping.FromString("1 2 3\n1 2 3\n1 2 3\n")
	res := p.Column(-1)
	if res.Error == nil {
		fmt.Println(p)
		t.Error("want error on non-positive Column, got nil")
	}
	data, err := io.ReadAll(p.Column(-1).Data)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) > 0 {
		t.Errorf("want no output from Column with invalid col, but got %q", data)
	}
	res = p.Column(5)
	if res.Error == nil {
		fmt.Println(p)
		t.Error("want error on non-positive Column, got nil")
	}
	data, err = io.ReadAll(p.Column(5).Data)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) > 0 {
		t.Errorf("want no output from Column with invalid col, but got %q", data)
	}
}

func TestStringReturnsPipeContents(t *testing.T) {
	t.Parallel()
	want := "Hello, world\n"
	p := piping.FromString(want)
	got, err := p.String()
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStringReturnsErrorWhenPipeErrorSet(t *testing.T) {
	t.Parallel()
	p := piping.FromString("Hello, world\n")
	p.Error = errors.New("oh no")
	_, err := p.String()
	if err == nil {
		t.Error("want error from String when pipeline has error, but got nil")
	}
}
