package server

import (
	"context"
	"errors"
)

type contextKey struct {
	name string
}

var ctxKey = contextKey{name: "custom-serv"}

type Values struct {
	TrackID    string
	StatusCode int
}

func GetRequestTrackID(ctx context.Context) string {
	vals, ok := ctx.Value(ctxKey).(Values)
	if !ok {
		return ""
	}
	return vals.TrackID
}

func GetRequestStatusCode(ctx context.Context) int {
	vals, ok := ctx.Value(ctxKey).(Values)
	if !ok {
		return 0
	}
	return vals.StatusCode
}

func SetRequestStatusCode(ctx context.Context, statusCode int) (context.Context, error) {
	vals, ok := ctx.Value(ctxKey).(Values)
	if !ok {
		return nil, errors.New("set status code err")
	}

	vals.StatusCode = statusCode

	return context.WithValue(ctx, ctxKey, vals), nil
}
