package docker

import (
	"fmt"

	"github.com/docker/docker/client"
)

type ServiceInfo struct {
	Name     string
	Image    string
	Replicas string
	Mode     string
}

func GetSwarmServices() ([]ServiceInfo, error) {
	cli, err := client.NewClientWithOpts(clientt.FromEnv)
	if err != nil {
		return nil, err
	}

	var result []ServiceInfo

	for _, svc := range services {
		name := svc.Spec.Name
		image := svc.Spec.TaskTemplate.ContainerSpec.Image
		mode := "replicated"
		replicas := "?"

		if svc.Spec.Mode.Global != nil {
			mode = "global"
			replicas = "-"
		} else if svc.Spec.Mode.Replicated != nil {
			replicas = fmt.Sprintf("%d", *svc.Spec.Mode.Replicated.Replicas)
		}

		result = append(result, ServiceInfo{
			Name:     name,
			Image:    image,
			Replicas: replicas,
			Mode:     mode,
		})
	}

	return result, nil
}
