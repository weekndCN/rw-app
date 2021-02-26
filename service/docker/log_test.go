package docker

import (
	"context"
	"io"
	"os"
	"testing"
)

func TestTail(t *testing.T) {
	ctx := context.Background()
	reader, err := Tail(ctx, "640e686a352d")
	if err != nil {
		t.Error(err)
	}

	io.Copy(os.Stdout, reader)
}
