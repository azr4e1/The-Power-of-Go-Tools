package linecounter_test

import (
	"bytes"
	"linecounter"
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
