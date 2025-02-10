package coreapi

import (
	"encoding/json"
	"time"
)

type BrickDateTime time.Time

const timeFormat = "2006-01-02T15:04:05.999999"

// Convert time.Time to CustomTime
func NewBrickDateTime(t time.Time) BrickDateTime {
	return BrickDateTime(t)
}

// Convert CustomTime to time.Time
func (ct BrickDateTime) Time() time.Time {
	return time.Time(ct)
}

func (ct *BrickDateTime) UnmarshalJSON(b []byte) error {
	var timeStr string
	if err := json.Unmarshal(b, &timeStr); err != nil {
		return err
	}

	t, err := time.Parse(timeFormat, timeStr)
	if err != nil {
		return err
	}

	*ct = BrickDateTime(t)
	return nil
}

// MarshalJSON implements json.Marshaler
func (ct *BrickDateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(*ct)
	return json.Marshal(t.Format(timeFormat))
}

// String implements Stringer interface
func (ct *BrickDateTime) String() string {
	return time.Time(*ct).Format(timeFormat)
}
