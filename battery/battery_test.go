//go:build integration

package battery_test

import (
	"battery"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

func TestGetBattery_GetsTheRightBatteryValueDischarging(t *testing.T) {
	t.Parallel()
	batteryStatus, err := os.ReadFile("testdata/battery_discharging.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := battery.Status{
		ChargePercent: 88,
	}
	got, err := battery.ParseACPIOutput(string(batteryStatus))
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
