package battery_test

import (
	"battery"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestGetBattery_TestBehaviorFailing(t *testing.T) {
	t.Parallel()
	_, err := battery.ParseACPIOutput("")
	if err == nil {
		t.Error("expected error, got nil")
	}
}
