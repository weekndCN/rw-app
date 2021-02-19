package logger

import (
	"context"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestContext(t *testing.T) {
	// generate an entry
	entry := logrus.NewEntry(logrus.StandardLogger())
	// generate context with entry
	ctx := WithContext(context.Background(), entry)
	// get entry from context
	got := FromContext(ctx)
	if got != entry {
		t.Error("Expected Logger from context")
	}
}

func TestNullConext(t *testing.T) {
	got := FromContext(context.Background())

	if got != log {
		t.Error("Expected default Logger from context")
	}
}

func TestRequest(t *testing.T) {
	entry := logrus.NewEntry(logrus.StandardLogger())

	ctx := WithContext(context.Background(), entry)
	req := new(http.Request)
	req = req.WithContext(ctx)
	got := FromRequest(req)

	if got != entry {
		t.Error("Expected logger from http.Request")
	}
}
