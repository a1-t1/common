package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"time"
)

type Time struct {
	sql.NullTime
}

func NewTime(t time.Time) Time {
	return Time{sql.NullTime{Time: t, Valid: true}}
}

func NewTimeFromPtr(t *time.Time) Time {
	if t == nil {
		return NewInvalidatedTime()
	}
	return NewTime(*t)
}

func NewInvalidatedTime() Time {
	return Time{}
}

func (t Time) IsNull() bool {
	return !t.Valid
}

func (t Time) Ptr() *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

// Marshal and unmarshal
func (t Time) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time.Format("2006-01-02 15:04:05"))
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t.Valid = false
		return nil
	}

	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return err
	}

	t.Time = parsedTime
	t.Valid = true
	return nil
}

func (t Time) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Time) Scan(value interface{}) error {
	var time sql.NullTime
	if err := time.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*t = Time{
			sql.NullTime{
				Valid: false,
			},
		}
	} else {
		*t = Time{
			sql.NullTime{
				Valid: true,
				Time:  time.Time,
			},
		}
	}
	return nil
}
