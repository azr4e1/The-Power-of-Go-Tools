package store

import (
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
)

type store struct {
	path   string
	memory map[string]string
}

type CmdlineReturn int

const (
	Ok CmdlineReturn = iota
	TooManyArgumentsError
	FlagsCombinationError
	OpenStoreError
	KeyNotFoundError
	CouldNotSaveError
)

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

func (s *store) Drop() {
	s.memory = map[string]string{}
}

func Main() CmdlineReturn {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-get | -set | -to | -drop | -create] [path]\n", os.Args[0])
		fmt.Println("Set and retrieve keys from store")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}
	set := flag.String("set", "", "Set key to value. Must be used with -to")
	to := flag.String("to", "", "Set key to value. Must be used with -set")
	get := flag.String("get", "", "Retrieve value of key.")
	drop := flag.Bool("drop", false, "Drop all the key-value pairs of the store")
	create := flag.Bool("create", false, "Create a store on the given path")
	flag.Parse()
	cmdArgs := flag.Args()
	if len(cmdArgs) < 1 || len(cmdArgs) > 1 {
		fmt.Fprintln(os.Stderr, "you can provide only one store path argument")
		return TooManyArgumentsError
	}
	path := cmdArgs[0]

	setCond, getCond, toCond := *set != "", *get != "", *to != ""

	mutuallyExclusive := func(args ...string) bool {
		flags := map[string]bool{
			"set":    setCond,
			"to":     toCond,
			"get":    getCond,
			"drop":   *drop,
			"create": *create,
		}
		var argsVal bool = true
		for _, el := range args {
			val, ok := flags[el]
			if !ok {
				return false
			}
			argsVal = val && argsVal
			flags[el] = false
		}

		var allVal bool = false
		for _, value := range flags {
			allVal = allVal || value
		}

		if !argsVal || allVal {
			return false
		}

		return true
	}

	store, err := OpenStore(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return OpenStoreError
	}

	switch {
	case mutuallyExclusive("set", "to"):
		store.Set(*set, *to)
	case mutuallyExclusive("get"):
		value, ok := store.Get(*get)
		if !ok {
			fmt.Fprintf(os.Stderr, "%q key not found\n", *get)
			return KeyNotFoundError
		}
		fmt.Println(value)
	case mutuallyExclusive("drop"):
		store.Drop()
	case mutuallyExclusive("create"):

		fmt.Println("ciao")
		_ = true
	default:
		fmt.Fprintln(os.Stderr, "You cannot use multiple flags together except for 'set' and 'to'")
		return FlagsCombinationError
	}

	err = store.Save()
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not save the store")
		return CouldNotSaveError
	}

	return Ok
}
