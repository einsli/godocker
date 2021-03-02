package controllers

import (
	"context"
	"dockerui/tools"
	_ "fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/fatih/structs"
	"strconv"
	"time"
)

func DockerNetworks(query map[string][]string) ([]map[string]interface{}, int, error) {
	var (
		cli *client.Client
		ctx context.Context
		cancel context.CancelFunc
		err error
		networkOpt *types.NetworkListOptions
		networks []types.NetworkResource
		networkMap map[string]interface{}
		networkResp []map[string]interface{}
		skip int
		limit int
		verifiedSkip int
		verifiedLimit int
	)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cli = tools.GClient

	/* 参考源码
	https://github.com/moby/moby/blob/master/integration-cli/docker_api_network_test.go
	*/
	networkFilters := filters.NewArgs()
	if len(query) == 0 {
		networkFilters.Add("name", "")
	}else {
		if _, ok := query["network"]; ok {
			networkFilters.Add("name", query["network"][0])
		}
	}
	networkOpt = &types.NetworkListOptions{
		Filters: networkFilters,
	}
	if networks, err = cli.NetworkList(ctx, *networkOpt); err != nil {
		return networkResp, len(networkResp), err
	} else {
		for _, network := range networks {
			networkMap = structs.Map(&network)
			networkResp = append(networkResp, networkMap)
		}
	}
	if _, ok := query["skip"]; ok {
		skip, _ = strconv.Atoi(query["skip"][0])
		limit, _ =  strconv.Atoi(query["limit"][0])
	} else {
		skip = 0
		limit = 10
	}
	verifiedSkip, verifiedLimit = tools.LimitVerify(skip, limit, len(networkResp))
	return networkResp[verifiedSkip:verifiedLimit], len(networkResp), nil
}
