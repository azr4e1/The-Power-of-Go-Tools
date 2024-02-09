package battery

import (
	"errors"
	"strconv"
	"strings"
)

type Status struct {
	ChargePercent int
}

func ParseACPIOutput(status string) (Status, error) {
	splitStatus := strings.Split(status, " ")
	if len(splitStatus) < 4 {
		return Status{}, errors.New("problem in parsing acpi output: string too short")
	}
	percentageBit := splitStatus[3]
	i, err := strconv.Atoi(strings.Replace(percentageBit, "%,", "", -1))
	if err != nil {
		return Status{}, errors.New("couldn't convert percentage to integer")
	}

	statusStruct := Status{
		ChargePercent: i,
	}
	return statusStruct, nil
}
