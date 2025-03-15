// Package env provides a set of functions to retrieve environment variables.
package env

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
)

// Keys is a list of keys that were retrieved from environment.
var Keys = []string{}

// GetString retrieves an environment variable and parses it as string.
func GetString(key, fallback string) string {
	Keys = append(Keys, key)
	if v, ok := os.LookupEnv(key); ok {
		return strings.TrimSpace(v)
	}
	return fallback
}

// GetInt retrieves an environment variable and parses it as int.
func GetInt(key string, fallback int) int {
	Keys = append(Keys, key)
	if v, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i
		}
	}
	return fallback
}

// GetBool retrieves an environment variable and parses it as bool.
//
// Note, only `true` is a valid true value.
func GetBool(key string, fallback bool) bool {
	Keys = append(Keys, key)
	if v, ok := os.LookupEnv(key); ok {
		return v == "true"
	}
	return fallback
}

// GetDuration retrieves an environment variable and parses it as duration.
func GetDuration(key string, fallback time.Duration) time.Duration {
	Keys = append(Keys, key)
	if v, ok := os.LookupEnv(key); ok {
		d, err := time.ParseDuration(v)
		if err == nil {
			return d
		}
	}
	return fallback
}

// GetUTCISODate retrieves an environment variable and parses it as ISO-8601
// date.
func GetUTCISODate(key string, fallback string) time.Time {
	return getTime(key, "2006-01-02", fallback)
}

// GetISOTime retrieves an environment variable and parses it as ISO-8601
// time.
func GetISOTime(key string, fallback string) time.Time {
	return getTime(key, time.RFC3339, fallback)
}

func getTime(key string, format string, fallback string) time.Time {
	Keys = append(Keys, key)
	if v, ok := os.LookupEnv(key); ok {
		t, err := time.Parse(format, v)
		if err == nil {
			return t
		}
	}

	t, err := time.Parse(format, fallback)
	if err != nil {
		panic(fmt.Sprintf("Invalid time fallback: %s", fallback))
	}
	return t
}

// GetSemVer retrieves an environment variable and parses it as semantic
// version.
func GetSemVer(key string, fallback string) *version.Version {
	Keys = append(Keys, key)
	if v, ok := os.LookupEnv(key); ok {
		vv, err := version.NewVersion(v)
		if err == nil {
			return vv
		}
	}

	v, err := version.NewVersion(fallback)
	if err != nil {
		panic(fmt.Sprintf("Invalid version fallback: %s", fallback))
	}
	return v
}

// GetURL retrieves an environment variable and parses it as URL.
//
// Note: you can use "*" as valid parsable, yet it result in simply "/*", so use with
// care.
func GetURL(key string, fallback string) *url.URL {
	Keys = append(Keys, key)
	if v, ok := os.LookupEnv(key); ok {
		u, err := url.Parse(v)
		if err == nil {
			return u
		}
	}

	v, err := url.Parse(fallback)
	if err != nil {
		panic(fmt.Sprintf("Invalid url fallback: %s", fallback))
	}
	return v
}
