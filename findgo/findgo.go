package findgo

import (
	"io/fs"
	"path/filepath"
)

func Files(fsys fs.FS) []string {
	var files []string

	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if filepath.Ext(p) == ".go" {
			files = append(files, p)
		}
		return nil
	})
	return files
}
