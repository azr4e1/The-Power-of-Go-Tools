package writer_test

import (
	"os"
	"testing"
	"writer"

	"github.com/google/go-cmp/cmp"
)

func TestWriteToFile_WritesGivenDataToFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/write_test.txt"
	want := []byte{1, 2, 3}
	err := writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestWriteToFile_GivesTheAppropriatePermissions(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/perm_test.txt"
	err := os.WriteFile(path, []byte{}, 0o644)
	if err != nil {
		t.Fatal(err)
	}
	err = writer.WriteToFile(path, []byte{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	perm := stat.Mode().Perm()
	if perm != 0o600 {
		t.Errorf("want file mode 0o600, got 0o%o", perm)
	}
}

func TestWriteToFile_ReturnsErrorForUnwritableFile(t *testing.T) {
	t.Parallel()
	path := "bogusdir/write_text.txt"
	err := writer.WriteToFile(path, []byte{})
	if err == nil {
		t.Fatal("want error when file is not writable")
	}
}

func TestWriteToFile_ClobbersGivenDataToFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/write_test.txt"
	prev_data := []byte{3, 2, 1}
	err := writer.WriteToFile(path, prev_data)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{1, 2, 3}
	err = writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestBuildByteSlice_CreatesASliceOfBytesWithSpecifiedLength(t *testing.T) {
	t.Parallel()
	want := []byte{0, 0, 0, 0, 0}
	got, err := writer.BuildByteSlice(5)
	if err != nil {
		t.Fatal("not expected error")
	}
	if len(got) != len(want) {
		t.Errorf("got lenght %d, expected length %d", len(got), len(want))
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestBuildByteSlice_Fails(t *testing.T) {
	t.Parallel()
	size := -10
	_, err := writer.BuildByteSlice(size)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
