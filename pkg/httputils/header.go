package httputils

import (
	"context"
)

type HeaderContent struct {
	UserUUID string `json:"uuid"`
	Role     string `json:"role"`
}

func GetClaims(header string) (*HeaderContent, error) {
	return nil, nil
}

type HeaderType string

const (
	MainHeader HeaderType = "main_header"
)

func HeaderFromContext(ctx context.Context) *HeaderContent {
	if header, ok := ctx.Value(MainHeader).(*HeaderContent); ok {
		return header
	}
	return nil
}
