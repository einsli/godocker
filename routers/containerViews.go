package routers

import (
	"dockerui/controllers"
	"dockerui/model"
	_ "dockerui/tools"
	"encoding/json"
	_ "encoding/json"
	_ "fmt"
	"github.com/fatih/structs"
	"net/http"
)

func GetContainers(w http.ResponseWriter, r *http.Request) {
	/*
		该函数用于查看docker容器
	*/
	var (
		query map[string][]string
		containerRes []map[string]interface{}
		err error
		resp *model.Response
		respData []byte
		respLen int
	)

	containerRes = make([]map[string]interface{},10)

	query = r.URL.Query()
	if r.Method != "GET" {
		resp = model.ResponseData(-1, "Method Not Allowed!", nil)
		methodErr := structs.Map(&resp)
		respData, _ = json.Marshal(methodErr)
		model.WriteResponse(w, 405, respData)
	}else {
		// 对容器进行查看
		if containerRes, respLen, err = controllers.DockerContainer(query); err != nil {
			resp = model.ResponseData(-1, err.Error(), model.ContainerResp{
				Containers: containerRes,
				Len: 0,
			})
		}else {
			resp = model.ResponseData(0, "Get Containers Successfully!", model.ContainerResp{
				Containers: containerRes,
				Len: respLen,
			})
		}
		containerMap := structs.Map(&resp)
		respData, _ = json.Marshal(containerMap)
		model.WriteResponse(w, 200, respData)
	}
	defer r.Body.Close()
}
