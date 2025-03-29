package validator

import (
    "errors"
    "unicode"
)

// ValidatePassword checks if a password meets the required criteria
func ValidatePassword(password string) error {
    var hasMinLen, hasUpper, hasLower, hasNumber, hasSpecial bool
    const minPasswordLength = 8

    if len(password) >= minPasswordLength {
        hasMinLen = true
    }

    for _, char := range password {
        switch {
        case unicode.IsUpper(char):
            hasUpper = true
        case unicode.IsLower(char):
            hasLower = true
        case unicode.IsDigit(char):
            hasNumber = true
        case unicode.IsPunct(char) || unicode.IsSymbol(char):
            hasSpecial = true
        }
    }

    if !hasMinLen {
        return errors.New("password must be at least 8 characters long")
    }
    if !hasUpper {
        return errors.New("password must contain at least one uppercase letter")
    }
    if !hasLower {
        return errors.New("password must contain at least one lowercase letter")
    }
    if !hasNumber {
        return errors.New("password must contain at least one number")
    }
    if !hasSpecial {
        return errors.New("password must contain at least one special character")
    }

    return nil
}