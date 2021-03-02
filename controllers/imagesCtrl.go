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

func DockerImages(query map[string][]string) ([]map[string]interface{}, int, error) {
	var (
		ctx context.Context
		cancel context.CancelFunc
		cli *client.Client
		err error
		imageList *types.ImageListOptions
		images []types.ImageSummary
		image types.ImageSummary
		imagesResp []map[string]interface{}
		imageMap map[string]interface{}
		skip int
		limit int
		verifiedSkip int
		verifiedLimit int
	)
	//  不带参数，返回全部
	if len(query) == 0 {
		imageList = &types.ImageListOptions{
			All: true,
		}
		// not worked
		skip = 0
		limit = 10
	} else {
		/*参考源码
		https://github.com/moby/moby/blob/master/integration-cli/docker_api_images_test.go
		*/
		// 带参数默认返回10
		imageFilters := filters.NewArgs()
		for search := range query{
			if search == "skip" ||  search == "limit" {
				continue
			}else {
				if search == "image" && len(query["image"]) != 0 {
					imageFilters.Add("reference", query[search][0])
				}
			}
		}
		imageList = &types.ImageListOptions{
			All: false,
			Filters: imageFilters,
		}
	}
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cli = tools.GClient

	// 查看所有镜像
	if images, err = cli.ImageList(ctx, *imageList); err != nil {
		fmt.Println("images list err", err)
		return imagesResp, len(imagesResp), err
	}
	for _, image = range images {
		imageMap = structs.Map(&image)
		imagesResp = append(imagesResp, imageMap)
		imagesResp = append(imagesResp, imageMap)
	}
	if _, ok := query["skip"]; ok {
		skip, _ = strconv.Atoi(query["skip"][0])
		limit, _ =  strconv.Atoi(query["limit"][0])
	} else {
		skip = 0
		limit = 10
	}
	// bug ,查询容器会返回两个重复的,貌似是开源库问题
	verifiedSkip, verifiedLimit = tools.LimitVerify(skip, limit, len(imagesResp))
	return imagesResp[verifiedSkip: verifiedLimit], len(imagesResp), nil
}
