package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"time"
)

type TimeOnly struct {
	sql.NullTime
}

func NewTimeOnly(t time.Time) TimeOnly {
	return TimeOnly{sql.NullTime{Time: t, Valid: true}}
}

func NewTimeOnlyFromPtr(t *time.Time) TimeOnly {
	if t == nil {
		return NewInvalidatedTimeOnly()
	}
	return NewTimeOnly(*t)
}

func NewInvalidatedTimeOnly() TimeOnly {
	return TimeOnly{}
}

func (t TimeOnly) IsNull() bool {
	return !t.Valid
}

func (t TimeOnly) Ptr() *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

// Marshal and unmarshal
func (t *TimeOnly) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time.Format("15:04:05"))
}

func (t *TimeOnly) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t.Valid = false
		return nil
	}

	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	parsedTime, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		return err
	}

	t.Valid = true
	t.Time = parsedTime
	return nil
}

func (t TimeOnly) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time.Format("15:04:05"), nil
	// return t.Time, nil
}

func (t *TimeOnly) Scan(value interface{}) error {
	var b sql.NullString
	if err := b.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*t = TimeOnly{
			sql.NullTime{
				Valid: false,
			},
		}
	} else {
		tm, err := time.Parse("15:04:05", b.String)
		if err != nil {
			*t = TimeOnly{sql.NullTime{Valid: false}}
		} else {
			*t = TimeOnly{
				sql.NullTime{
					Valid: true,
					Time:  tm,
				},
			}
		}
	}
	return nil
}
