package controllers

import (
	"context"
	"dockerui/tools"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/fatih/structs"
	"strconv"
	"time"
)

func DockerContainer(query map[string][]string) ([]map[string]interface{}, int, error) {
	// TODO 官方提供的容器查询没有起作用，需要自行实现查询
	var (
		cli *client.Client
		containerListopt *types.ContainerListOptions
		containers []types.Container
		containerMap map[string]interface{}
		containerQuit bool
		ctx context.Context
		cancel context.CancelFunc
		containerResp []map[string]interface{}
		err error
		skip int
		limit int
		verifiedSkip int
		verifiedLimit int
	)
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cli = tools.GClient
	// 查询所有容器
	if len(query) == 0 {
		limit = 10
		containerListopt = &types.ContainerListOptions{
			All: true,
		}
	}else {
		/*
			参考源码
			目前源码有bug,没有对Quiet字段进行处理，所以无法查询退出的容器,源码里该字段已经废弃
		    提交的issue
		    https://github.com/moby/moby/issues/42092
			https://github.com/moby/moby/blob/master/integration-cli/docker_api_containers_test.go
		*/
		if _, ok := query["quiet"]; ok {
			containerQuit = true
		}
		// 根据容器名查询， 开源代码没处理，需要自己实现
		containerFilters := filters.NewArgs()
		if _, ok := query["container"]; ok {
			containerFilters.Add("name", query["container"][0])
		}
		// Limit 需要自己实现分页，不用源码里面的limit
		containerListopt = &types.ContainerListOptions{
			All: false,
			Quiet: containerQuit,
			//Limit: limit,
			Filters: containerFilters,
		}
	}
	if containers, err = cli.ContainerList(ctx, *containerListopt); err !=nil {
		fmt.Println("Get containers error")
		return containerResp, len(containerResp), err
	}else {
		for _, container := range containers {
			containerMap = structs.Map(&container)
			containerResp = append(containerResp, containerMap)
		}
	}
	if _, ok := query["skip"]; ok {
		skip, _ = strconv.Atoi(query["skip"][0])
		limit, _ =  strconv.Atoi(query["limit"][0])
	} else {
		skip = 0
		limit = 10
	}
	verifiedSkip, verifiedLimit = tools.LimitVerify(skip, limit, len(containerResp))
	return containerResp[verifiedSkip:verifiedLimit], len(containerResp), nil
}
