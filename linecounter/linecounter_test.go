package linecounter_test

import (
	"bytes"
	"github.com/rogpeppe/go-internal/testscript"
	"linecounter"
	"os"
	"testing"
)

func TestLinesCountsLinesInInput(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	lines := "The quick brown fox\njumps over the\n lazy dog\n"
	buf.WriteString(lines)
	want := 3
	counter, err := linecounter.NewCounter(linecounter.WithInput(buf))
	if err != nil {
		t.Fatal(err)
	}
	got := counter.Lines()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestWithInputFromArgs_SetsInputToGivenPath(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	want := 3
	counter, err := linecounter.NewCounter(linecounter.WithInputFromArgs(args))
	if err != nil {
		t.Fatal(err)
	}
	got := counter.Lines()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestWithInputFromArgs_SetsInputToGivenPathMultipleFiles(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt", "testdata/five_lines.txt"}
	want := 8
	counter, err := linecounter.NewCounter(linecounter.WithInputFromArgs(args))
	if err != nil {
		t.Fatal(err)
	}
	got := counter.Lines()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgs_NoArgsProvided(t *testing.T) {
	t.Parallel()
	args := []string{}
	want := 0
	counter, err := linecounter.NewCounter(linecounter.WithInputFromArgs(args))
	if err != nil {
		t.Errorf("not expected error")
	}
	got := counter.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBytesCounter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	bytes := "The quick brown fox\njumps over the\n lazy dog\n"
	buf.WriteString(bytes)
	want := len(buf.Bytes())
	counter, err := linecounter.NewCounter(linecounter.WithInput(buf))
	if err != nil {
		t.Fatal(err)
	}
	got := counter.Bytes()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"count": linecounter.Main,
	}))
}

func TestWordsCountsWordsInput(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	lines := "The quick brown fox\njumps over the\n lazy   dog\n"
	buf.WriteString(lines)
	want := 9
	counter, err := linecounter.NewCounter(linecounter.WithInput(buf))
	if err != nil {
		t.Fatal(err)
	}
	got := counter.Words()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
