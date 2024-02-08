package fsmonth_test

import (
	"fsmonth"
	"os"
	"testing"
	"testing/fstest"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestMonthFiles_CorrectlyListFilesInTree(t *testing.T) {
	t.Parallel()
	want := []string{
		"recent_dir/even_more_recent_file",
		"recent_dir/recent_dir/barely_one_month_old",
		"recent_file",
	}
	fsys := os.DirFS("testdata")
	got := fsmonth.MonthFiles(fsys)
	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}

func TestMonthFiles_CorrectlyListFilesInMapFS(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"recent_file":                                &fstest.MapFile{ModTime: time.Now()},
		"recent_dir/even_more_recent_file":           &fstest.MapFile{ModTime: time.Now()},
		"recent_dir/old_file":                        &fstest.MapFile{ModTime: time.Date(2019, 01, 33, 1, 1, 1, 1, &time.Location{})},
		"recent_dir/recent_dir/older_file":           &fstest.MapFile{ModTime: time.Date(2004, 1, 1, 1, 1, 1, 1, &time.Location{})},
		"recent_dir/recent_dir/barely_one_month_old": &fstest.MapFile{ModTime: time.Now()},
	}
	want := []string{
		"recent_dir/even_more_recent_file",
		"recent_dir/recent_dir/barely_one_month_old",
		"recent_file",
	}
	got := fsmonth.MonthFiles(fsys)
	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}
