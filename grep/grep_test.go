package grep_test

import (
	"bytes"
	"grep"
	"testing"
)

func TestStreamingReaderToFindWord(t *testing.T) {
	t.Parallel()
	word := "hello"
	input := bytes.NewBufferString("hello how are you?\nhello, I'm good thanks you?\nYes I'm fine thank you.\nWhat did you eat yesterday?")
	finder, _ := grep.NewFinder(grep.WithInput(input))
	want := []string{"hello how are you?", "hello, I'm good thanks you?"}
	got := finder.Find(word)
	for index, el := range want {
		if el != got[index] {
			t.Errorf("want %q, got %q", el, got[index])
		}
	}
}
