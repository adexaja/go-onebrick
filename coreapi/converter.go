package coreapi

import (
	"encoding/json"
	"strconv"
)

type Amount float64

func (a Amount) String() string {
	amount := strconv.FormatFloat(float64(a), 'f', -1, 64)
	return string(amount)
}

func (a Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

func (a *Amount) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	number, _ := strconv.ParseFloat(s, 64)
	*a = Amount(number)
	return nil
}

func (a Amount) ToFloat() float64 {
	return float64(a)
}
