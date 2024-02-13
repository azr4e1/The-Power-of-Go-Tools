package shell_test

// behavior:
// - it should read a line of input and interpret it as a command line
// - it should execute that command line and print the output, along with any errors
// - if the input line is empty, the shell should do nothing
// - if the input line consists only of the end-of-file character (EOF), the shell should exit

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"shell"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCmdFromString_ReturnsCmdValue(t *testing.T) {
	t.Parallel()
	want := exec.Command("sleep", "10")
	var command string = "sleep 10"
	got, err := shell.CmdFromString(command)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(*want, *got, cmpopts.IgnoreUnexported(exec.Cmd{})) {
		t.Fatal(cmp.Diff(*want, *got))
	}
}

func TestCmdFromString_ErrorsOnEmptyOutput(t *testing.T) {
	t.Parallel()
	_, err := shell.CmdFromString("")
	if err == nil {
		t.Fatal("want error on empty input, got nil")
	}
}

func TestNewSession_CreatesANewSession(t *testing.T) {
	t.Parallel()
	want := shell.Session{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	got := *shell.NewSession(os.Stdin, os.Stdout, os.Stderr)
	if want != got {
		t.Errorf("want %#v, got %#v", want, got)
	}
}

func TestRun_CreatesANewInteractiveSessionAndOutputsToStdout(t *testing.T) {
	t.Parallel()
	input := new(bytes.Buffer)
	output := new(bytes.Buffer)

	inputText := "echo test\n\n"
	want := "> hello\n> > \nBe seeing you!\n"

	input.WriteString(inputText)

	session := shell.NewSession(input, output, io.Discard)
	session.Run()

	got := output.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
