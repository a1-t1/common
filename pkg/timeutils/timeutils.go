package timeutils

import (
	"fmt"
	"time"
)

// ParseDate parses a string YYYY-MM-DD to time.Time
func ParseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

// ParseTime parses a string HH:ii:ss to time.Time
func ParseTime(t string) (time.Time, error) {
	return time.Parse("15:04", t)
}

// Parse parses a string YYYY-MM-DD HH:ii:ss to time.Time
func Parse(date string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", date)
}

// ParseDatePtr parses a string YYYY-MM-DD to time.Time
func ParseDatePtr(date string) (*time.Time, error) {
	t, err := ParseDate(date)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// Since returns the time since the given time
func Since(t time.Time) string {
	now := Now()
	diff := now.Sub(t)
	// decide the unit of the time
	if diff < 1*time.Minute {
		return "just now"
	}
	if diff < 1*time.Hour {
		return fmt.Sprintf("%d minutes", diff/time.Minute)
	}
	if diff < 1*time.Hour*24 {
		return fmt.Sprintf("%d hours", diff/time.Hour)
	}
	if diff < 1*time.Hour*24*30 {
		return fmt.Sprintf("%d days", diff/(time.Hour*24))
	}
	if diff < 1*time.Hour*24*365 {
		return fmt.Sprintf("%d months", diff/(time.Hour*24*30))
	}
	return fmt.Sprintf("%d years", diff/(time.Hour*24*365))
}

// ParseDateFromPtr parses a string YYYY-MM-DD to time.Time
func ParseDateFromPtr(date *string) (*time.Time, error) {
	if date == nil {
		return nil, nil
	}
	return ParseDatePtr(*date)
}

func Now() time.Time {
	return time.Now().UTC()
}

func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// IsTimeWithinRange checks if a given time is within a specified range
func IsTimeWithinRange(t, start, end time.Time) bool {
	return (t.Equal(start) || t.After(start)) && (t.Equal(end) || t.Before(end))
}

// TimeOfDay returns a time.Time with only hour and minute set
func TimeOfDay(t time.Time) time.Time {
	return time.Date(2000, 1, 1, t.Hour(), t.Minute(), 0, 0, time.UTC)
}

// ParseTimeOfDay parses a string HH:MM and returns a time.Time with only hour and minute set
func ParseTimeOfDay(t string) (time.Time, error) {
	parsed, err := time.Parse("15:04:00", t)
	if err != nil {
		return time.Time{}, err
	}
	return TimeOfDay(parsed), nil
}

// StartOfDayFormatted returns the start of the day in the format YYYY-MM-DD 00:00:00
func StartOfDayFormatted(t time.Time) string {
	return t.Format("2006-01-02 00:00:00")
}

// EndOfDayFormatted returns the end of the day in the format YYYY-MM-DD 23:59:59
func EndOfDayFormatted(t time.Time) string {
	return t.Format("2006-01-02 23:59:59")
}

// ParseRFC3339 parses a string in RFC3339 format to time.Time
func ParseRFC3339(t string) (time.Time, error) {
	return time.Parse(time.RFC3339, t)
}
