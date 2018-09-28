package eventsub

import (
	"context"
	"dockerdns/pkg/containerinfo"
	"fmt"

	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
)

type EventHandler interface {
	handle(events.Message)
}

type LoggingEventHandler struct {
	Client  *client.Client
	Context *context.Context
}

func (eventHandler LoggingEventHandler) handle(event events.Message) {
	if event.Type == events.ContainerEventType {
		containerId := event.Actor.ID

		switch event.Action {
		case "start":
			containerInfo := containerinfo.GetContainerInfo(eventHandler.Client, eventHandler.Context, containerId)

			fmt.Println("Container Info: " + fmt.Sprint(containerInfo))
			fmt.Println()
			break

		default:
			fmt.Printf("Docker Action: %s - Container: %s", event.Action, containerId)
			break
		}

		fmt.Println()
	}
}
