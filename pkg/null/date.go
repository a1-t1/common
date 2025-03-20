package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/a1-t1/common/pkg/timeutils"
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
	if string(data) == "null" {
		t.Valid = false
		return nil
	}
	// since incoming (data) is a string "2006-01-02", we need to remove the quotes and parse it to a time.Time
	str := string(data)
	str = strings.Trim(str, "\"")
	time, err := timeutils.ParseDate(str)
	if err != nil {
		t.Valid = false
		return nil
	}
	t.Valid = true
	t.Time = time
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
