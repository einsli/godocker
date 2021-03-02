package model

import "net/http"

type Response struct {
	Code int `json:"Code"`
	Msg string `json:"Msg"`
	Data interface{}`json:"Data"`
}

/*
组合返回的数据格式函数
*/
func ResponseData(code int, msg string , data interface{}) *Response {
	return &Response{
		Code: code,
		Msg: msg,
		Data: data,
	}
}

/*
用于将相应数据返回给页面
*/

func WriteResponse(w http.ResponseWriter, statusCode int, bodyData []byte) {
	w.WriteHeader(statusCode)
	w.Write(bodyData)
}


type ImagesResp struct {
	Images []map[string]interface{} `json:"Images"`
	Len int `json:"Len"`
}

type ContainerResp struct {
	Containers []map[string]interface{} `json:"Containers"`
	Len int `json:"Len"`
}

type NetworkResp struct {
	Networks []map[string]interface{} `json:"Networks"`
	Len int `json:"Len"`
}