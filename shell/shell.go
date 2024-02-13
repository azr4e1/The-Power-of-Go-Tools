package shell

import (
	"errors"
	"io"
	"os/exec"
	"strings"
)

type Session struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
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

// func Main() int {
// 	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
// 	session.Run()
// 	return 0
// }
