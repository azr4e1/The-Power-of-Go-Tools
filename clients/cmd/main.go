package main

import (
	"clients"
	"fmt"
	"io"
	"net/http"
	"os"
)

const Usage = `Usage: weather LOCATION

Example: weather London,UK`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(Usage)
		os.Exit(0)
	}

	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHERMAP_API_KEY.")
		os.Exit(1)
	}

	location := os.Args[1]
	URL := clients.FormatURL(clients.BaseURL, location, key)
	r, err := http.Get(URL)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if r.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "status code %d\n", r.StatusCode)
		os.Exit(1)
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	weather, err := clients.ParseResponse(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(weather)
}
