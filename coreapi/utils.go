package coreapi

import (
	"encoding/json"
	"time"
)

type BrickDateTime time.Time

func (ct *BrickDateTime) UnmarshalJSON(b []byte) error {
	var timeStr string
	if err := json.Unmarshal(b, &timeStr); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02T15:04:05.999999", timeStr)
	if err != nil {
		return err
	}

	*ct = BrickDateTime(t)
	return nil
}
