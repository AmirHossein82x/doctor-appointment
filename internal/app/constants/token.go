package constants

import "time"

type TokenType string

const (
    // Token types
    TokenTypeAccess  TokenType = "access"
    TokenTypeRefresh TokenType = "refresh"

    // Token lifetimes
    AccessTokenLifetime  = 15 * time.Minute   // Access token valid for 15 minutes
    RefreshTokenLifetime = 7 * 24 * time.Hour // Refresh token valid for 7 days
)