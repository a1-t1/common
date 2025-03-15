package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

type Float struct {
	sql.NullFloat64
}

func NewFloat64(i float64) Float {
	return Float{sql.NullFloat64{Float64: i, Valid: true}}
}

func NewFloat64FromPtr(i *float64) Float {
	if i == nil {
		return NewInvalidatedFloat()
	}
	return NewFloat64(*i)
}

func NewInvalidatedFloat() Float {
	return Float{}
}

func (i Float) IsNull() bool {
	return !i.Valid
}

func (i Float) Ptr() *float64 {
	if !i.Valid {
		return nil
	}
	return &i.Float64
}

func (i Float) ValueOr(v float64) float64 {
	if !i.Valid {
		return v
	}
	return i.Float64
}

// Marshal and unmarshal

func (i *Float) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.Float64)
}

func (i *Float) UnmarshalJSON(data []byte) error {
	var n *float64
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	if n != nil {
		i.Valid = true
		i.Float64 = *n
	} else {
		i.Valid = false
	}
	return nil
}

// Value implements the driver Valuer interface.
func (i Float) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.Float64, nil
}

// Scan implements the Scanner interface.
func (i *Float) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*i = Float{
			sql.NullFloat64{
				Valid: false,
			},
		}
	} else {
		*i = Float{
			sql.NullFloat64{
				Valid:   true,
				Float64: f.Float64,
			},
		}
	}
	return nil
}
