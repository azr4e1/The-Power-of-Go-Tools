package findgo_test

import (
	"archive/zip"
	"findgo"
	"os"
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"
)

func TestFilesCorrectlyListFilesInTree(t *testing.T) {
	t.Parallel()
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}
	fsys := os.DirFS("testdata/tree")
	got := findgo.Files(fsys)
	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}

func TestFilesCorrectlyListFilesInMapFS(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}
	got := findgo.Files(fsys)
	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}

func BenchmarkFilesOnDisk(b *testing.B) {
	fsys := os.DirFS("testdata/tree")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findgo.Files(fsys)
	}
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findgo.Files(fsys)
	}
}

func TestFilesCorrectlyListsFilesInZIPArchive(t *testing.T) {
	t.Parallel()
	fsys, err := zip.OpenReader("testdata/files.zip")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"tree/file.go",
		"tree/subfolder/subfolder.go",
		"tree/subfolder2/another.go",
		"tree/subfolder2/file.go",
	}
	got := findgo.Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
