//go:build integration

package battery_test

import (
	"battery"
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestGetACPIOutput_CapturesCmdOutput(t *testing.T) {
	t.Parallel()
	data, err := exec.Command("/usr/bin/acpi").CombinedOutput()
	if err != nil {
		t.Skipf("unable to run 'acpi' command: %v", err)
	}
	if !bytes.Contains(data, []byte("Battery")) {
		t.Skip("no battery fitted")
	}
	text, err := battery.GetACPIOutput()
	fmt.Println("text is:", text)
	if err != nil {
		t.Fatal(err)
	}
	status, err := battery.ParseACPIOutput(text)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Charge: %d%%", status.ChargePercent)
}
