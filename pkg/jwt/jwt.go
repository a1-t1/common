// Package jwt provides JWT functionality.
package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

// ErrTokenExpired is returned when a token is expired.
var ErrTokenExpired = errors.New("token is expired")

// Verify checks the signature of a given JWT token using a given secret.
func Verify(jwtSecret string, tokenString string) (jwt.MapClaims, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("token is empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		var validationErr *jwt.ValidationError
		if errors.As(err, &validationErr) {
			if validationErr.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			}
		}

		return nil, fmt.Errorf("could not parse JWT token: %w - %s", err, tokenString)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, errors.New("invalid JWT token claims")
}

// Sign signs a given set of claims using a given secret.
func Sign(jwtSecret string, claims jwt.MapClaims, expirationTimestamp int64) (string, error) {
	claims["exp"] = expirationTimestamp

	// Create the JWT claims, which includes the userId, externalId and expiry time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	return tokenString, err
}
