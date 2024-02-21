package battery2

import (
	"bytes"
	"encoding/json"
)

type Battery struct {
	Name             string
	ID               int
	ChargePercent    int
	TimeToFullCharge string
	Present          bool
}

func (b Battery) ToJSON() (string, error) {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(b)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
