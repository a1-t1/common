package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

type Int struct {
	sql.NullInt64
}

func NewInt(i int64) Int {
	return Int{sql.NullInt64{Int64: i, Valid: true}}
}

func NewIntFromPtr(i *int64) Int {
	if i == nil {
		return NewInvalidatedInt()
	}
	return NewInt(*i)
}

func (i *Int) ValueOr(v int64) int64 {
	if !i.Valid {
		return v
	}
	return i.Int64
}

func NewInvalidatedInt() Int {
	return Int{}
}

func (i Int) IsNull() bool {
	return !i.Valid
}

func (i Int) Ptr() *int64 {
	if !i.Valid {
		return nil
	}
	return &i.Int64
}

// Marshal and unmarshal
func (i *Int) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.Int64)
}

func (i *Int) UnmarshalJSON(data []byte) error {
	var n *int64
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	if n != nil {
		i.Valid = true
		i.Int64 = *n
	} else {
		i.Valid = false
	}
	return nil
}

func (i Int) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.Int64, nil
}

func (i *Int) Scan(value interface{}) error {
	var num sql.NullInt64
	if err := num.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*i = Int{
			sql.NullInt64{
				Valid: false,
			},
		}
	} else {
		*i = Int{
			sql.NullInt64{
				Valid: true,
				Int64: num.Int64,
			},
		}
	}
	return nil
}
