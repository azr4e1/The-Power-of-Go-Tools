package writer

import (
	"os"
)

func WriteToFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0644)
	return err
}
