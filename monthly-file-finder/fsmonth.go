package fsmont

import (
	"io/fs"
	"time"
)

func MonthFiles(fsys fs.FS) []string {
	var files []string

	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		info, err := d.Info()
		if err != nil {
			return err
		}
		if info.ModTime().AddDate(0, 0, 30).Compare(time.Now()) < 0 {
			return nil
		}
		files = append(files, p)
		return nil
	})
	return []string{}
}
