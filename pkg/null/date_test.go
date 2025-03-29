package null

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestUnmarshal(t *testing.T) {
	s := `"2021-01-01"`
	date := Date{}
	err := json.Unmarshal([]byte(s), &date)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(date)
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Date
		wantErr bool
	}{
		{
			name:    "valid date",
			input:   `"2021-01-01"`,
			want:    NewDate(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			wantErr: false,
		},
		{
			name:    "null value",
			input:   "null",
			want:    NewInvalidatedDate(),
			wantErr: false,
		},
		{
			name:    "invalid format - wrong separator",
			input:   `"2021/01/01"`,
			want:    NewInvalidatedDate(),
			wantErr: false,
		},
		{
			name:    "invalid format - incomplete",
			input:   `"2021-01"`,
			want:    NewInvalidatedDate(),
			wantErr: false,
		},
		{
			name:    "invalid date - month out of range",
			input:   `"2021-13-01"`,
			want:    NewInvalidatedDate(),
			wantErr: false,
		},
		{
			name:    "invalid date - day out of range",
			input:   `"2021-01-32"`,
			want:    NewInvalidatedDate(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Date
			err := json.Unmarshal([]byte(tt.input), &got)

			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Valid != tt.want.Valid {
				t.Errorf("UnmarshalJSON() validity = %v, want %v", got.Valid, tt.want.Valid)
			}

			if got.Valid && !got.Time.Equal(tt.want.Time) {
				t.Errorf("UnmarshalJSON() time = %v, want %v", got.Time, tt.want.Time)
			}
		})
	}
}
