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

type session struct {
	stdin      io.Reader
	stdout     io.Writer
	stderr     io.Writer
	dryRun     bool
	transcript bool
}

type option func(*session) error

func SetStdin(stdin io.Reader) option {
	return func(s *session) error {
		if stdin == nil {
			return errors.New("Stdin is nil")
		}
		s.stdin = stdin
		return nil
	}
}

func SetStdout(stdout io.Writer) option {
	return func(s *session) error {
		if stdout == nil {
			return errors.New("Stdout is nil")
		}
		s.stdout = stdout
		return nil
	}
}

func SetStderr(stderr io.Writer) option {
	return func(s *session) error {
		if stderr == nil {
			return errors.New("Stderr is nil")
		}
		s.stderr = stderr
		return nil
	}
}

func SetDryRun(dryrun bool) option {
	return func(s *session) error {
		s.dryRun = dryrun
		return nil
	}
}

func SetTranscript(transcript bool) option {
	return func(s *session) error {
		s.transcript = transcript
		return nil
	}
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

func NewSession(opt ...option) *session {
	session := session{
		stdin:      os.Stdin,
		stdout:     os.Stdout,
		stderr:     os.Stderr,
		dryRun:     false,
		transcript: false,
	}

	ptr := &session

	for _, o := range opt {
		err := o(ptr)
		if err != nil {
			continue
		}
	}

	return ptr
}

func (s *session) Run() {
	fmt.Fprintf(s.stdout, "> ")
	input := bufio.NewScanner(s.stdin)
	for input.Scan() {
		line := input.Text()
		cmd, err := CmdFromString(line)
		if err != nil {
			fmt.Fprintf(s.stdout, "> ")
			continue
		}
		if s.dryRun {
			fmt.Fprintf(s.stdout, "command '%s' would have been executed\n> ", line)
			continue
		}
		data, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.stdout, "error:", err)
		}
		fmt.Fprintf(s.stdout, "%s> ", data)
	}
	fmt.Fprintln(s.stdout, "\nBe seeing you!")
}

func Main() int {
	session := NewSession()
	session.Run()
	return 0
}
