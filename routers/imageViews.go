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


func GetDockerImages(w http.ResponseWriter, r *http.Request) {
	/*
	该函数用于查看docker 镜像
	*/

	var (
		query map[string][]string
		imageRes []map[string]interface{}
		err error
		resp *model.Response
		respData []byte
		respLen int
	)
	query = r.URL.Query()
	imageRes = make([]map[string]interface{},10)
	if r.Method != "GET" {
		resp = model.ResponseData(-1, "Method Not Allowed!", nil)
		methodErr := structs.Map(&resp)
		respData, _ = json.Marshal(methodErr)
		model.WriteResponse(w, 405, respData)
	} else {
		if imageRes, respLen, err = controllers.DockerImages(query); err != nil {
			resp = model.ResponseData(-1, err.Error(), model.ImagesResp{
				Images: imageRes,
				Len: 0,
			})
		}else {
			resp = model.ResponseData(0, "Get Images SuccessFully!", model.ImagesResp{
				Images: imageRes,
				Len: respLen,
			})
		}
		imagesMap := structs.Map(&resp)
		respData, _ = json.Marshal(imagesMap)
		model.WriteResponse(w, 200, respData)
	}

	defer r.Body.Close()
}
