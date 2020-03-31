package server

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/mamoe/mirai-public-api/logger"
)

var miraiApiServer = &MiraiApiServer{}

type MiraiApiServer struct {
	ConfigPath string
	Config     map[string]string
	Logger     *logger.MiraiLogger
	Router     *fasthttprouter.Router
	Service    *Service
}

func RunMiraiApiServer() {
	miraiApiServer.Init()
}

func GetService() *Service {
	return GetMiraiApiServer().Service
}

func (this *MiraiApiServer) ConstructPluginPath() string {
	return this.Config["basicRepoPath"]
}

func GetMiraiApiServer() *MiraiApiServer {
	return miraiApiServer
}
