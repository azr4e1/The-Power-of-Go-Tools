package howlong

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func TimeIt(command ...string) (int64, error) {
	currentTime := time.Now()
	_, err := exec.Command(command[0], command[1:]...).CombinedOutput()
	if err != nil {
		return 0, err
	}
	finalTime := time.Now()

	duration := finalTime.Sub(currentTime).Milliseconds()

	return duration, nil
}

func Main() int {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [command]\n", os.Args[0])
		fmt.Println("Measure time taken by command")
	}
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return 0
	}
	duration, err := TimeIt(args...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "an error occurred")
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Printf("This command took: %dms\n", duration)
	return 0
}
