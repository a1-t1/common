package null

import (
	"database/sql/driver"
	"encoding/json"
	"testing"
	"time"
)

func TestNewTime(t *testing.T) {
	now := time.Now().UTC()
	time := NewTime(now)

	if !time.Valid {
		t.Error("expected Valid to be true")
	}

	if !time.Time.Equal(now) {
		t.Errorf("expected Time to be %v, got %v", now, time.Time)
	}
}

func TestNewTimeFromPtr(t *testing.T) {
	tests := []struct {
		name      string
		input     *time.Time
		wantValid bool
	}{
		{
			name:      "valid time",
			input:     func() *time.Time { t := time.Now().UTC(); return &t }(),
			wantValid: true,
		},
		{
			name:      "nil pointer",
			input:     nil,
			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTimeFromPtr(tt.input)

			if got.Valid != tt.wantValid {
				t.Errorf("NewTimeFromPtr() validity = %v, want %v", got.Valid, tt.wantValid)
			}

			if tt.input != nil && !got.Time.Equal(*tt.input) {
				t.Errorf("NewTimeFromPtr() time = %v, want %v", got.Time, *tt.input)
			}
		})
	}
}

func TestNewInvalidatedTime(t *testing.T) {
	time := NewInvalidatedTime()

	if time.Valid {
		t.Error("expected Valid to be false")
	}
}

func TestTime_IsNull(t *testing.T) {
	tests := []struct {
		name string
		time Time
		want bool
	}{
		{
			name: "valid time",
			time: NewTime(time.Now().UTC()),
			want: false,
		},
		{
			name: "invalid time",
			time: NewInvalidatedTime(),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.time.IsNull(); got != tt.want {
				t.Errorf("Time.IsNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_Ptr(t *testing.T) {
	now := time.Now().UTC()
	tests := []struct {
		name string
		time Time
		want *time.Time
	}{
		{
			name: "valid time",
			time: NewTime(now),
			want: &now,
		},
		{
			name: "invalid time",
			time: NewInvalidatedTime(),
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.time.Ptr()

			if tt.want == nil && got != nil {
				t.Errorf("Time.Ptr() = %v, want nil", *got)
				return
			}

			if tt.want != nil && got == nil {
				t.Errorf("Time.Ptr() = nil, want %v", *tt.want)
				return
			}

			if tt.want != nil && got != nil && !got.Equal(*tt.want) {
				t.Errorf("Time.Ptr() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	timeStr := "2023-01-02 15:04:05"
	parsedTime, _ := time.Parse("2006-01-02 15:04:05", timeStr)

	tests := []struct {
		name    string
		time    Time
		want    string
		wantErr bool
	}{
		{
			name:    "valid time",
			time:    NewTime(parsedTime),
			want:    `"2023-01-02 15:04:05"`,
			wantErr: false,
		},
		{
			name:    "invalid time",
			time:    NewInvalidatedTime(),
			want:    "null",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.time.MarshalJSON()

			if (err != nil) != tt.wantErr {
				t.Errorf("Time.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if string(got) != tt.want {
				t.Errorf("Time.MarshalJSON() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	timeStr := "2023-01-02 15:04:05"
	parsedTime, _ := time.Parse("2006-01-02 15:04:05", timeStr)

	tests := []struct {
		name    string
		input   string
		want    Time
		wantErr bool
	}{
		{
			name:    "valid time",
			input:   `"2023-01-02 15:04:05"`,
			want:    NewTime(parsedTime),
			wantErr: false,
		},
		{
			name:    "null value",
			input:   "null",
			want:    NewInvalidatedTime(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Time
			err := json.Unmarshal([]byte(tt.input), &got)

			if (err != nil) != tt.wantErr {
				t.Errorf("Time.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Valid != tt.want.Valid {
				t.Errorf("Time.UnmarshalJSON() validity = %v, want %v", got.Valid, tt.want.Valid)
			}

			if got.Valid && !got.Time.Equal(tt.want.Time) {
				t.Errorf("Time.UnmarshalJSON() time = %v, want %v", got.Time, tt.want.Time)
			}
		})
	}
}

func TestTime_Value(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		name    string
		time    Time
		want    driver.Value
		wantErr bool
	}{
		{
			name:    "valid time",
			time:    NewTime(now),
			want:    now,
			wantErr: false,
		},
		{
			name:    "invalid time",
			time:    NewInvalidatedTime(),
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.time.Value()

			if (err != nil) != tt.wantErr {
				t.Errorf("Time.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want == nil && got != nil {
				t.Errorf("Time.Value() = %v, want nil", got)
				return
			}

			if tt.want != nil {
				gotTime, ok := got.(time.Time)
				if !ok {
					t.Errorf("Time.Value() returned value of type %T, want time.Time", got)
					return
				}

				if !gotTime.Equal(tt.want.(time.Time)) {
					t.Errorf("Time.Value() = %v, want %v", gotTime, tt.want)
				}
			}
		})
	}
}

func TestTime_Scan(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		name    string
		value   interface{}
		want    Time
		wantErr bool
	}{
		{
			name:    "valid time",
			value:   now,
			want:    NewTime(now),
			wantErr: false,
		},
		{
			name:    "nil value",
			value:   nil,
			want:    NewInvalidatedTime(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Time
			err := got.Scan(tt.value)

			if (err != nil) != tt.wantErr {
				t.Errorf("Time.Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Valid != tt.want.Valid {
				t.Errorf("Time.Scan() validity = %v, want %v", got.Valid, tt.want.Valid)
			}

			if got.Valid && tt.want.Valid && !got.Time.Equal(tt.want.Time) {
				t.Errorf("Time.Scan() time = %v, want %v", got.Time, tt.want.Time)
			}
		})
	}
}

// Test struct serialization
func TestTimeJSONSerialization(t *testing.T) {
	type TestStruct struct {
		TimeField Time `json:"time_field"`
	}

	timeStr := "2023-01-02 15:04:05"
	parsedTime, _ := time.Parse("2006-01-02 15:04:05", timeStr)

	tests := []struct {
		name    string
		input   TestStruct
		want    string
		wantErr bool
	}{
		{
			name:    "valid time",
			input:   TestStruct{TimeField: NewTime(parsedTime)},
			want:    `{"time_field":"2023-01-02 15:04:05"}`,
			wantErr: false,
		},
		{
			name:    "null time",
			input:   TestStruct{TimeField: NewInvalidatedTime()},
			want:    `{"time_field":null}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("json.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if string(got) != tt.want {
				t.Errorf("json.Marshal() = %v, want %v", string(got), tt.want)
			}

			// Test unmarshaling back
			var unmarshaled TestStruct
			err = json.Unmarshal(got, &unmarshaled)

			if err != nil {
				t.Errorf("json.Unmarshal() error = %v", err)
				return
			}

			if unmarshaled.TimeField.Valid != tt.input.TimeField.Valid {
				t.Errorf("Unmarshaled validity = %v, want %v",
					unmarshaled.TimeField.Valid, tt.input.TimeField.Valid)
			}

			if unmarshaled.TimeField.Valid && !unmarshaled.TimeField.Time.Equal(tt.input.TimeField.Time) {
				t.Errorf("Unmarshaled time = %v, want %v",
					unmarshaled.TimeField.Time, tt.input.TimeField.Time)
			}
		})
	}
}
