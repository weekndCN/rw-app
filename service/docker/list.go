package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// CustomContainer return to web client format
type CustomContainer struct {
	ID     string       `json:"ID"`
	Names  []string     `json:"names"`
	Image  string       `json:"image"`
	Ports  []types.Port `json:"ports"`
	State  string       `json:"state"`
	Status string       `json:"status"`
}

// List list all container
func List() ([]*CustomContainer, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		// list all container
		All: true,
	})
	if err != nil {
		return nil, err
	}

	var cs []*CustomContainer

	for _, container := range containers {
		c := &CustomContainer{
			ID:     container.ID,
			Names:  container.Names,
			Image:  container.Image,
			Ports:  container.Ports,
			State:  container.State,
			Status: container.Status,
		}
		cs = append(cs, c)
	}

	return cs, nil
}
