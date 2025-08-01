package utils

import (
	"regexp"
	"strings"
)

// ParsePhoneNumber parses and validates Iraqi phone numbers, returning them with the 964 country code
func ParsePhoneNumber(phoneNumber string) string {
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.TrimSpace(phoneNumber)

	if phoneNumber == "" {
		return ""
	}

	if strings.HasPrefix(phoneNumber, "+964") {
		phoneNumber = phoneNumber[4:]
	} else if strings.HasPrefix(phoneNumber, "00964") {
		phoneNumber = phoneNumber[5:]
	} else if strings.HasPrefix(phoneNumber, "964") {
		phoneNumber = phoneNumber[3:]
	}

	phoneNumber = strings.TrimPrefix(phoneNumber, "0")

	iraqiMobileRegex := regexp.MustCompile(`^7[5789][0-9]{8}$`)
	if !iraqiMobileRegex.MatchString(phoneNumber) {
		return ""
	}

	return "964" + phoneNumber
}

// ParseIntPhoneNumber parses international phone numbers by removing the '+' prefix and normalizing the format
// This function does not enforce any specific country code and preserves the original country code
func ParseIntPhoneNumber(phoneNumber string) string {
	// Remove all spaces and trim
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "(", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, ")", "")
	phoneNumber = strings.TrimSpace(phoneNumber)

	if phoneNumber == "" {
		return ""
	}

	// Remove the '+' prefix if present
	phoneNumber = strings.TrimPrefix(phoneNumber, "+")
	phoneNumber = strings.TrimPrefix(phoneNumber, "00")

	// Basic validation - ensure it contains only digits and has reasonable length
	digitRegex := regexp.MustCompile(`^[0-9]+$`)
	if !digitRegex.MatchString(phoneNumber) {
		return ""
	}

	// Reasonable length check (typically 7-15 digits for international numbers)
	if len(phoneNumber) < 7 || len(phoneNumber) > 15 {
		return ""
	}

	return phoneNumber
}
