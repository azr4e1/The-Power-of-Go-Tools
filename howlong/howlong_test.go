package howlong_test

import (
	"testing"
	"time"

	"howlong"
)

func TestRunReportsCorrectElapsedTimeForCommand(t *testing.T) {
	t.Parallel()
	target := 100 * time.Millisecond
	elapsed, err := howlong.TimeIt("sleep", "0.1")
	if err != nil {
		t.Fatal(err)
	}
	epsilon := target - time.Duration(elapsed)
	if epsilon.Abs() > 300*time.Millisecond {
		t.Fatalf("want %s, got %s (not close enough)", target, time.Duration(elapsed))
	}
}
