package battery2

import (
	// "bytes"
	"encoding/json"
)

type Battery struct {
	Name             string
	ID               int
	ChargePercent    int
	TimeToFullCharge string
	Present          bool
}

func (b Battery) ToJSON() string {
	// buf := new(bytes.Buffer)
	// encoder := json.NewEncoder(buf)
	// err := encoder.Encode(b)
	// if err != nil {
	// 	return "", err
	// }

	// return buf.String(), nil
	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(data)
}
