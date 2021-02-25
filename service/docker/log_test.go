package docker

import (
	"context"
	"io"
	"os"
	"testing"
)

func TestTail(t *testing.T) {
	var ctx context.Context
	reader, err := Tail(ctx, "28cb50796bce")
	if err != nil {
		t.Error(err)
	}

	io.Copy(os.Stdout, reader)
}
