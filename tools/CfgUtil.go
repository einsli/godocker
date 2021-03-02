package tools

import (
	"fmt"
	_ "fmt"
	gcfg "gopkg.in/gcfg.v1"
	"log"
	"os/exec"
)


func GetServerPort() int {
	serverConfig := struct {
		Server struct {
			PORT int
		}
	}{}
	err := gcfg.ReadFileInto(&serverConfig, "./config/config.ini")
	if err != nil {
		log.Fatalf("config error: %v", err)
		return 8080
	}else {
		serverPort := serverConfig.Server.PORT
		return serverPort
	}
}

func PortInUse(port int) bool {
	checkStatement := fmt.Sprintf("lsof -i:%d ", port)
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if len(output) > 0 {
		return true
	}
	return false
}