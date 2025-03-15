package utils

import (
	"regexp"
	"strings"
)

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
