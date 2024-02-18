package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	data, err := os.Open("../testdata/log.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer data.Close()

	// get ips
	ips := []string{}
	ipsCount := map[string]int{}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		ip := strings.Fields(line)[0]
		ips = append(ips, ip)
	}
	for _, ip := range ips {
		ipsCount[ip]++
	}

	keys := make([]string, len(ipsCount))
	i := 0
	for el := range ipsCount {
		keys[i] = el
		i++
	}

	slices.SortFunc(keys, func(a, b string) int {
		switch {
		case ipsCount[a] > ipsCount[b]:
			return -1
		case ipsCount[a] < ipsCount[b]:
			return 1
		default:
			return 0
		}
	})

	for _, ip := range keys[:10] {
		fmt.Println(ipsCount[ip], ip)
	}
}
