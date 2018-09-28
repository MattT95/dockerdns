package containerinfo

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type ContainerInfo struct {
	ID           string
	Name         string
	HostName     string
	PortMappings nat.PortMap
}

func (cinfo ContainerInfo) String() string {
	return fmt.Sprintf("[%s - %s - %s - %s]",
		cinfo.ID,
		cinfo.Name,
		cinfo.HostName,
		cinfo.PortMappings)
}

func GetContainerInfo(client *client.Client, ctx *context.Context, containerId string) ContainerInfo {

	inspection, err := client.ContainerInspect(*ctx, containerId)

	if err != nil {
		panic(err)
	}

	containerInfo := ContainerInfo{
		ID:           containerId,
		Name:         inspection.Name,
		HostName:     inspection.Config.Hostname,
		PortMappings: inspection.HostConfig.PortBindings}

	return containerInfo
}
