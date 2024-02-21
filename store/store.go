package store

import (
	"encoding/gob"
	"errors"
	"io/fs"
	"os"
)

type store struct {
	path   string
	memory map[string]string
}

func OpenStore(path string) (*store, error) {

	s := &store{path: path, memory: map[string]string{}}
	f, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return s, nil
	}

	if err != nil {
		return nil, err
	}

	defer f.Close()
	err = gob.NewDecoder(f).Decode(&s.memory)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *store) Get(key string) (string, bool) {
	v, ok := s.memory[key]

	return v, ok
}

func (s *store) Set(key, value string) {
	s.memory[key] = value
}

func (s *store) Save() error {
	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer f.Close()

	return gob.NewEncoder(f).Encode(s.memory)
}
