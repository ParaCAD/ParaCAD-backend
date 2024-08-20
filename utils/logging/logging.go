package logging

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type CtxKey string

const ErrorKey CtxKey = "error"
const StatusKey CtxKey = "status"

func Middleware(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		timeStart := time.Now()
		h(w, r, p)
		duration := time.Since(timeStart)

		logAttrs := []slog.Attr{
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Duration("duration", duration),
		}
		logLevel := slog.LevelInfo
		status := http.StatusOK
		msg := "Ok"

		err, isErr := r.Context().Value(ErrorKey).(string)
		if isErr {
			logAttrs = append(logAttrs, slog.String("error", err))
			logLevel = slog.LevelError
			status, _ = r.Context().Value(StatusKey).(int)
			msg = "Error"
		}

		logAttrs = append(logAttrs, slog.Int("status", status))

		slog.LogAttrs(r.Context(), logLevel, msg, logAttrs...)
	}
}

func AddErrorToRequestCtx(r *http.Request, err error, status int) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, ErrorKey, err.Error())
	ctx = context.WithValue(ctx, StatusKey, status)
	*r = *r.WithContext(ctx)
}

func Init(level slog.Level) {
	slog.SetLogLoggerLevel(level)
}
