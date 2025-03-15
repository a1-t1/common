package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultService_parsePhoneNumber(t *testing.T) {
	tests := []struct {
		name        string
		phoneNumber string
		want        string
	}{
		{
			name:        "valid number with +964 prefix",
			phoneNumber: "+9647712345678",
			want:        "9647712345678",
		},
		{
			name:        "valid number with 00964 prefix",
			phoneNumber: "009647712345678",
			want:        "9647712345678",
		},
		{
			name:        "valid number with 964 prefix",
			phoneNumber: "9647712345678",
			want:        "9647712345678",
		},
		{
			name:        "valid number with leading 0",
			phoneNumber: "07712345678",
			want:        "9647712345678",
		},
		{
			name:        "valid number without prefix",
			phoneNumber: "7712345678",
			want:        "9647712345678",
		},
		{
			name:        "valid number with spaces",
			phoneNumber: " +964 771 234 5678 ",
			want:        "9647712345678",
		},
		{
			name:        "invalid number format",
			phoneNumber: "1234567890",
			want:        "",
		},
		{
			name:        "invalid Iraqi mobile prefix",
			phoneNumber: "9647612345678", // starts with 76 instead of 77,78,79,75
			want:        "",
		},
		{
			name:        "empty string",
			phoneNumber: "",
			want:        "",
		},
		{
			name:        "valid number starting with 75",
			phoneNumber: "9647512345678",
			want:        "9647512345678",
		},
		{
			name:        "valid number starting with 78",
			phoneNumber: "9647812345678",
			want:        "9647812345678",
		},
		{
			name:        "valid number starting with 79",
			phoneNumber: "9647912345678",
			want:        "9647912345678",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParsePhoneNumber(tt.phoneNumber)
			assert.Equal(t, tt.want, got)
		})
	}
}
