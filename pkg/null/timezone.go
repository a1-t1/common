package null

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"time"
)

type DateTimeZone struct {
	Valid bool
	Time  time.Time
}

func NewDateTimeZone(t time.Time) DateTimeZone {
	return DateTimeZone{Valid: true, Time: t}
}

func NewDateTimeZoneFromString(dt string) (DateTimeZone, error) {
	t, err := time.Parse(time.RFC3339, dt)
	if err != nil {
		return NewInvalidatedDateTimeZone(), err
	}
	return NewDateTimeZone(t), nil
}

func NewInvalidatedDateTimeZone() DateTimeZone {
	return DateTimeZone{Valid: false}
}

func (dtz DateTimeZone) IsNull() bool {
	return !dtz.Valid
}

func (dtz DateTimeZone) Ptr() *time.Time {
	if !dtz.Valid {
		return nil
	}
	return &dtz.Time
}

// Marshal and unmarshal
func (dtz DateTimeZone) MarshalJSON() ([]byte, error) {
	if !dtz.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(dtz.Time.Format(time.DateTime))
}

func (dtz *DateTimeZone) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		dtz.Valid = false
		return nil
	}

	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return err
	}

	dtz.Time = t
	dtz.Valid = true
	return nil
}

func (dtz DateTimeZone) Value() (driver.Value, error) {
	if !dtz.Valid {
		return nil, nil
	}
	return dtz.Time.Format(time.RFC3339), nil
}

func (dtz *DateTimeZone) Scan(value interface{}) error {
	if value == nil {
		dtz.Valid = false
		return nil
	}

	switch v := value.(type) {
	case string:
		t, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return err
		}
		dtz.Time = t
		dtz.Valid = true
	case []byte:
		t, err := time.Parse(time.RFC3339, string(v))
		if err != nil {
			return err
		}
		dtz.Time = t
		dtz.Valid = true
	case time.Time:
		dtz.Time = v
		dtz.Valid = true
	default:
		return &time.ParseError{
			Layout: time.RFC3339,
			Value:  reflect.TypeOf(value).String(),
		}
	}
	return nil
}
