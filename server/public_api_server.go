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
}

func RunMiraiApiServer() {
	miraiApiServer.Init()
}

func GetMiraiApiServer() *MiraiApiServer {
	return miraiApiServer
}
