package eventsub

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func SubscribeToEvents(cli *client.Client, ctx context.Context, eventHandler EventHandler) {
	eventOptions := types.EventsOptions{}

	eventChannel, errChannel := cli.Events(ctx, eventOptions)

	for {
		select {
		case event := <-eventChannel:
			eventHandler.handle(event)
		case err := <-errChannel:
			if err != nil {
				panic(err)
			}
		}
	}
}
