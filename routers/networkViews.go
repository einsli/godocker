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


func GetNetworks(w http.ResponseWriter, r *http.Request) {
	/*
		该函数用于查看docker网络
	*/
	var (
		query map[string][]string
		networkRes []map[string]interface{}
		err error
		resp *model.Response
		respData []byte
		respLen int
	)

	networkRes = make([]map[string]interface{},10)

	query = r.URL.Query()
	if r.Method != "GET" {
		resp = model.ResponseData(-1, "Method Not Allowed!", nil)
		methodErr := structs.Map(&resp)
		respData, _ = json.Marshal(methodErr)
		model.WriteResponse(w, 405, respData)
	}else {
		// 对容器进行查看
		if networkRes, respLen, err = controllers.DockerNetworks(query); err != nil {
			resp = model.ResponseData(-1, err.Error(), model.NetworkResp{
				Networks: networkRes,
				Len: 0,
			})
		}else {
			resp = model.ResponseData(0, "Get Network Successfully!", model.ContainerResp{
				Containers: networkRes,
				Len: respLen,
			})
		}
		networkMap := structs.Map(&resp)
		respData, _ = json.Marshal(networkMap)
		model.WriteResponse(w, 200, respData)
	}
	defer r.Body.Close()
}