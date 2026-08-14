package sharedmodels

import (
	"encoding/json"
	"strconv"
)

type Scalar string

func (value *Scalar) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*value = ""
		return nil
	}
	var text string
	if err := json.Unmarshal(data, &text); err == nil {
		*value = Scalar(text)
		return nil
	}
	var number json.Number
	if err := json.Unmarshal(data, &number); err != nil {
		return err
	}
	*value = Scalar(number.String())
	return nil
}

func (value Scalar) String() string {
	return string(value)
}

func (value Scalar) Int64() (int64, error) {
	return strconv.ParseInt(value.String(), 10, 64)
}
