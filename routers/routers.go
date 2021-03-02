package routers

import(
	_ "fmt"
	"net/http"
)

func init()  {
	http.HandleFunc("/images/all", GetDockerImages)
	http.HandleFunc("/images/search", GetDockerImages)
	http.HandleFunc("/container/all", GetContainers)
	http.HandleFunc("/container/search", GetContainers)
	http.HandleFunc("/networks/all", GetNetworks)
	http.HandleFunc("/networks/search", GetNetworks)
}
