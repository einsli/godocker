package tools

import (
	"fmt"
	"github.com/docker/docker/client"
)

var (
	GClient *client.Client
)

func dockerClient() (*client.Client, error){
	var (
		cli *client.Client
		err error
	)
	if cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation()); err != nil {
		fmt.Println("New client error: #{err}")
		return cli, err
	}
	return cli, err
}

func InitDockerClient() {
	var (
		err error
	)
	if GClient ,err = dockerClient(); err != nil {
		panic(err.Error())
	}
}
