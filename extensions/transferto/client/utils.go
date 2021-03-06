package client

import (
	"encoding/json"
	"strings"

	"github.com/shopspring/decimal"
)

// CSVStrings is a list of strings which can be automatically unmarshalled from a CSV list
type CSVStrings []string

// UnmarshalJSON unmarshals this list from a CSV string
func (l *CSVStrings) UnmarshalJSON(data []byte) error {
	var asString string
	if err := json.Unmarshal(data, &asString); err != nil {
		return err
	}
	*l = strings.Split(asString, ",")
	return nil
}

// CSVDecimals is a list of decimals which can be automatically unmarshalled from a CSV list
type CSVDecimals []decimal.Decimal

// UnmarshalJSON unmarshals this list from a CSV string
func (l *CSVDecimals) UnmarshalJSON(data []byte) error {
	var asStrings CSVStrings
	if err := json.Unmarshal(data, &asStrings); err != nil {
		return err
	}

	vals := make([]decimal.Decimal, len(asStrings))
	var err error
	for v := range asStrings {
		vals[v], err = decimal.NewFromString(asStrings[v])
		if err != nil {
			return err
		}
	}

	*l = vals
	return nil
}
