package shell

import (
	"bufio"
	"errors"
	"flag"
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
	transcript io.Writer
	dryRun     bool
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

func SetTranscript(transcript io.Writer) option {
	return func(s *session) error {
		if transcript == nil {
			return errors.New("Transcript buffer is nil")
		}
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
		transcript: io.Discard,
		dryRun:     false,
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
	fmt.Fprintf(s.transcript, "> ")
	input := bufio.NewScanner(s.stdin)
	for input.Scan() {
		line := input.Text()
		fmt.Fprintln(s.transcript, line)

		cmd, err := CmdFromString(line)
		if err != nil {
			fmt.Fprintf(s.stdout, "> ")
			fmt.Fprintf(s.transcript, "> ")
			continue
		}
		if s.dryRun {
			fmt.Fprintf(s.stdout, "command '%s' would have been executed\n> ", line)
			fmt.Fprintf(s.transcript, "command '%s' would have been executed\n> ", line)
			continue
		}
		data, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.stdout, "error:", err)
			fmt.Fprintln(s.transcript, "error:", err)
		}
		fmt.Fprintf(s.stdout, "%s> ", data)
		fmt.Fprintf(s.transcript, "%s> ", data)
	}
	fmt.Fprintln(s.stdout, "\nBe seeing you!")
	fmt.Fprintln(s.transcript, "\nBe seeing you!")
}

func Main() int {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-transcript]\n", os.Args[0])
		fmt.Println("Launch shell")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}
	transcript := flag.Bool("transcript", false, "record the transcript of this session and save it to file 'transcript.txt' in the current folder.")
	flag.Parse()
	args := make([]option, 0, 1)

	if *transcript {
		file, err := os.Create("transcript.txt")
		defer file.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		args = append(args, SetTranscript(file))
	}

	session := NewSession(args...)
	// session := NewSession()
	session.Run()
	return 0
}
