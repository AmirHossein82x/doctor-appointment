package validator

import "regexp"

// ValidatePhoneNumber validates the phone number format.
func ValidateIranianPhoneNumber(phoneNumber string) bool {
	// Regex pattern for Iranian mobile numbers
	phoneRegex := `^09(0[1-3]|1[0-9]|3[0-9]|9[0-9])[0-9]{7}$`
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phoneNumber)
}
