package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

var (
	_ json.Marshaler   = (*String)(nil)
	_ json.Unmarshaler = (*String)(nil)
	_ sql.Scanner      = (*String)(nil)
	_ driver.Valuer    = (*String)(nil)
)

// String is a nullable string.
type String struct {
	sql.NullString
}

func NewString(s string) String {
	if s == "" {
		return NewInvalidatedString()
	}
	return String{sql.NullString{String: s, Valid: true}}
}

func NewInvalidatedString() String {
	return String{}
}

func (s *String) ValueOr(v string) string {
	if !s.Valid {
		return v
	}
	return s.String
}

func (s String) Ptr() *string {
	if !s.Valid {
		return nil
	}
	return &s.String
}

func NewStringFromPtr(s *string) String {
	if s == nil {
		return NewInvalidatedString()
	}
	return NewString(*s)
}

func (s String) IsNull() bool {
	return !s.Valid
}

// Marshal and unmarshal
func (s *String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *String) UnmarshalJSON(data []byte) error {
	var str *string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	if str != nil {
		s.Valid = true
		s.String = *str
	} else {
		s.Valid = false
	}
	return nil
}

func (s String) Value() (driver.Value, error) {
	if s.Valid {
		return s.String, nil
	}
	return nil, nil
}

func (s *String) Scan(value interface{}) error {
	var str sql.NullString
	if err := str.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*s = String{
			sql.NullString{
				Valid: false,
			},
		}
	} else {
		*s = String{
			sql.NullString{
				Valid:  true,
				String: str.String,
			},
		}
	}
	return nil
}
