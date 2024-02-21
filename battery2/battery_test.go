package battery2_test

import (
	"battery2"
	"io"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToJSON_EncodesDataCorrectlyToString(t *testing.T) {
	t.Parallel()

	batt := battery2.Battery{
		Name:             "InternalBattery-0",
		ID:               10813539,
		ChargePercent:    100,
		TimeToFullCharge: "0:00",
		Present:          true,
	}

	goldenFile, err := os.Open("testdata/batteryToJsonGoldenFile.json")
	if err != nil {
		t.Fatal(err)
	}
	wantByte, err := io.ReadAll(goldenFile)
	if err != nil {
		t.Fatal(err)
	}
	want := string(wantByte)
	got := batt.ToJSON()

	if !cmp.Equal(want, got) {
		t.Error(want, got)
	}
}
