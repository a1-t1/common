package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

type Bool struct {
	sql.NullBool
}

func NewBool(b bool) Bool {
	return Bool{sql.NullBool{Bool: b, Valid: true}}
}

func NewBoolFromPtr(b *bool) Bool {
	if b == nil {
		return NewInvalidatedBool()
	}
	return NewBool(*b)
}

func (b *Bool) ValueOr(v bool) bool {
	if !b.Valid {
		return v
	}
	return b.Bool
}

func NewInvalidatedBool() Bool {
	return Bool{
		sql.NullBool{}}
}

func (b *Bool) IsNull() bool {
	return !b.Valid
}

func (b *Bool) Ptr() *bool {
	if !b.Valid {
		return nil
	}
	return &b.Bool
}

// Marshal and unmarshal

func (b *Bool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.Bool)
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	var n *bool
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	if n == nil {
		*b = NewInvalidatedBool()
		return nil
	}
	*b = NewBool(*n)
	return nil
}

func (b *Bool) Scan(value interface{}) error {
	var boolean sql.NullBool
	if err := boolean.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*b = Bool{
			sql.NullBool{
				Valid: false,
			},
		}
	} else {
		*b = Bool{
			sql.NullBool{
				Valid: true,
				Bool:  boolean.Bool,
			},
		}
	}
	return nil
}

func (b Bool) Value() (driver.Value, error) {
	if !b.Valid {
		return nil, nil
	}
	return b.Bool, nil
}
