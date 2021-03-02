package tools

import (
	"log"
	"strconv"
)


func PrintProcessStart(port int) {
	runningMsg := "http server Running on http://0.0.0.0:" + strconv.Itoa(port)
	log.Println(runningMsg)
}
