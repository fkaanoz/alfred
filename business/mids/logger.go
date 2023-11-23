package mids

import (
	"context"
	"github.com/fkaanoz/alfred/foundation/server"
	"go.uber.org/zap"
	"net/http"
)

func Logger(logger *zap.SugaredLogger) server.Middleware {

	m := func(handler server.Handler) server.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			logger.Infow("REQUEST", "status", "started", "trackID", server.GetRequestTrackID(ctx))
			handler(ctx, w, r)
			logger.Infow("REQUEST", "status", "ended", "trackID", server.GetRequestTrackID(ctx))
		}

		return h
	}
	return m
}
