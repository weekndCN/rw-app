package docker

import (
	"context"

	"github.com/docker/docker/client"
)

// Stop stop a container with a specified container ID
func Stop(id string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	if err := cli.ContainerStop(ctx, id, nil); err != nil {
		return err
	}

	return nil

}
