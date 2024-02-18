package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Session struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	DryRun bool
}

func CmdFromString(cmdline string) (*exec.Cmd, error) {
	splitBySpace := strings.Fields(cmdline)
	if len(splitBySpace) < 1 {
		return nil, errors.New("empty string")
	}
	command, arguments := splitBySpace[0], splitBySpace[1:]
	cmd := exec.Command(command, arguments...)

	return cmd, nil
}

func NewSession(stdin io.Reader, stdout, stderr io.Writer) *Session {
	session := Session{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}

	return &session
}

func (s *Session) Run() {
	fmt.Fprintf(s.Stdout, "> ")
	input := bufio.NewScanner(s.Stdin)
	for input.Scan() {
		line := input.Text()
		cmd, err := CmdFromString(line)
		if err != nil {
			fmt.Fprintf(s.Stdout, "> ")
			continue
		}
		if s.DryRun {
			fmt.Fprintf(s.Stdout, "command '%s' would have been executed\n> ", line)
			continue
		}
		data, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.Stdout, "error:", err)
		}
		fmt.Fprintf(s.Stdout, "%s> ", data)
	}
	fmt.Fprintln(s.Stdout, "\nBe seeing you!")
}

func Main() int {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	session.Run()
	return 0
}
