package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

type ServiceInfo struct {
	Name     string
	Image    string
	Replicas string
	Mode     string
	Status   string
	Color    string
}

func GetSwarmServices() ([]ServiceInfo, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	cli.NegotiateAPIVersion(context.Background())

	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		return nil, err
	}

	tasks, err := cli.TaskList(context.Background(), types.TaskListOptions{})
	if err != nil {
		return nil, err
	}

	taskMap := map[string][]swarm.Task{}
	for _, t := range tasks {
		taskMap[t.ServiceID] = append(taskMap[t.ServiceID], t)
	}

	var result []ServiceInfo

	for _, svc := range services {
		name := svc.Spec.Name
		image := svc.Spec.TaskTemplate.ContainerSpec.Image
		mode := "replicated"
		status := "unknown"
		color := "\033[0m"

		running := 0
		total := len(taskMap[svc.ID])
		for _, t := range taskMap[svc.ID] {
			if t.Status.State == swarm.TaskStateRunning {
				running++
			}
		}

		switch {
		case running == total && total > 0:
			status = "Running"
			color = "\033[32m" // green
		case running == 0 && total > 0:
			status = "Failed"
			color = "\033[31m" // red
		default:
			status = "Partial"
			color = "\033[33m" // yellow
		}

		result = append(result, ServiceInfo{
			Name:     name,
			Image:    image,
			Replicas: fmt.Sprintf("%d/%d", running, total),
			Mode:     mode,
			Status:   status,
			Color:    color,
		})
	}

	return result, nil
}
