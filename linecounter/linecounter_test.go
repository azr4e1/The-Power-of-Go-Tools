package linecounter_test

import (
	"bytes"
	"github.com/rogpeppe/go-internal/testscript"
	"linecounter"
	"os"
	"testing"
)

func TestLineCounter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	lines := "The quick brown fox\njumps over the\n lazy dog\n"
	buf.WriteString(lines)
	want := 3
	counter, err := linecounter.NewCounter(linecounter.WithInput(buf))
	if err != nil {
		t.Fatal(err)
	}
	got := counter.Count()

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
	got := counter.Count()

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
	got := counter.Count()
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
		"linecounter": linecounter.Main,
	}))
}
