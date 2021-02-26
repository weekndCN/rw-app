package docker

import (
	"bytes"
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type logReadCloser struct {
	readerCloser io.ReadCloser
	tty          bool
	lastHeader   []byte
	buffer       bytes.Buffer
}

// Tail tail container log
func Tail(ctx context.Context, id string) (io.ReadCloser, error) {
	// if you want to a timeout,using context.WithTimeout insteaded
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer defer cancel()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	options := types.ContainerLogsOptions{Follow: true, ShowStdout: true}
	// Replace this ID with a container that really exists
	out, err := cli.ContainerLogs(ctx, id, options)
	if err != nil {
		return nil, err
	}

	return out, nil
}
