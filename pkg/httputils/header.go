package httputils

import (
	"context"
	"errors"
)

func GetClaims[T any](header string) (*T, error) {
	return nil, nil
}

type HeaderType string

const (
	MainHeader HeaderType = "main_header"
)

func HeaderFromContext[T any](ctx context.Context) (*T, error) {
	if header, ok := ctx.Value(MainHeader).(*T); ok {
		return header, nil
	}
	return nil, errors.New("header not found")
}
