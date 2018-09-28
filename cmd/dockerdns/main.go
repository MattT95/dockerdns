package main

import (
	"dockerdns/pkg/eventsub"
	"fmt"

	"github.com/docker/docker/client"

	"golang.org/x/net/context"
)

const (
	DockerVersion = "1.38"
)

func main() {

	fmt.Println("Creating client with Docker SDK Client Version: " + DockerVersion)

	cli := createDockerClient(DockerVersion)
	ctx := context.Background()

	eventHandler := eventsub.LoggingEventHandler{Client: cli, Context: &ctx}

	eventsub.SubscribeToEvents(cli, ctx, eventHandler)
}

func createDockerClient(dockerVersion string) *client.Client {

	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	client.WithVersion(dockerVersion)(cli)

	return cli
}
