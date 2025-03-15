package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"time"
)

type Date struct {
	sql.NullTime
}

func NewDate(t time.Time) Date {
	return Date{sql.NullTime{Time: t, Valid: true}}
}

func NewDateFromPtr(t *time.Time) Date {
	if t == nil {
		return NewInvalidatedDate()
	}
	return NewDate(*t)
}

func NewInvalidatedDate() Date {
	return Date{}
}

func (t Date) IsNull() bool {
	return !t.Valid
}

func (t Date) Ptr() *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

// Marshal and unmarshal
func (t *Date) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time.Format("2006-01-02"))
}

func (t *Date) UnmarshalJSON(data []byte) error {
	var tm *sql.NullTime
	if err := json.Unmarshal(data, &tm); err != nil {
		return err
	}
	if tm != nil {
		t.Valid = true
		t.Time = tm.Time
	} else {
		t.Valid = false
	}
	return nil
}

func (t Date) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Date) Scan(value interface{}) error {
	var time sql.NullTime
	if err := time.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*t = Date{
			sql.NullTime{
				Valid: false,
			},
		}
	} else {
		*t = Date{
			sql.NullTime{
				Valid: true,
				Time:  time.Time,
			},
		}
	}
	return nil
}
