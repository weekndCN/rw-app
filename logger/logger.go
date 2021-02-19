package logger

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type loggerKey struct{}

var log = logrus.NewEntry(logrus.StandardLogger())

// WithContext return a new context with loggger
func WithContext(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// FromContext if no logger available return the default logger
func FromContext(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey{})
	if logger == nil {
		return log
	}

	return logger.(*logrus.Entry)
}

// FromRequest if no logger available return the default logger
func FromRequest(r *http.Request) *logrus.Entry {
	return FromContext(r.Context())
}
