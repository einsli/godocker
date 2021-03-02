package server

import (
	_ "dockerui/routers"
	"dockerui/tools"
	"log"
	"net/http"
	"strconv"
)

func Run() {
	var (
		portUsed bool
		errMsg string
	)
	tools.InitDockerClient()
	port := tools.GetServerPort()
	portUsed = tools.PortInUse(port)
	if portUsed {
		errMsg = "Port "+ strconv.Itoa(port) + " Is Already In Used!"
		log.Fatalf(errMsg)
	}
	portStr := strconv.Itoa(port)
	var runPort = "0.0.0.0:" +
		portStr
	tools.PrintProcessStart(port)
	server := http.Server{
		Addr: runPort,
	}
	server.ListenAndServe()
	defer tools.GClient.Close()
}
