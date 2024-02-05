package writer

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func WriteToFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0o600)
	if err != nil {
		return err
	}
	err = os.Chmod(path, 0o600)
	return err
}

func BuildByteSlice(size int) ([]byte, error) {
	if size < 0 {
		return []byte{}, errors.New("size must be positive")
	}
	seq := []byte{}
	for i := 0; i < size; i++ {
		seq = append(seq, 0x00)
	}
	return seq, nil
}

func Main() int {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-size] [files...]\n",
			os.Args[0])
		fmt.Println("Creates a file with specified name and size")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}
	size := flag.Int("size", 0, "Create file of specified size")
	flag.Parse()

	content, err := BuildByteSlice(*size)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	args := flag.Args()
	if len(args) < 1 && *size != 0 {
		fmt.Fprintln(os.Stderr, "you need to specify path to write to")
		return 3
	} else if len(args) < 1 {
		return 0
	}

	err = WriteToFile(flag.Args()[0], content)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}

	return 0
}
