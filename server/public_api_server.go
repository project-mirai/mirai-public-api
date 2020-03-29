package server

import (
	"bufio"
	"github.com/buaazp/fasthttprouter"
	"github.com/mamoe/mirai-public-api/logger"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"os"
	"strings"
)

const (
	BaseDir       = "mirai-plugins"
	PluginListDir = "plugins"
)

var Config map[string]string

type MiraiApiServer struct {
	ConfigPath string
	Logger     *logger.MiraiLogger
	Router     *fasthttprouter.Router
}

func (this *MiraiApiServer) LoadAll() {
	//logger
	this.Logger = logger.NewLogger("MiraiAPIServer")
	this.Logger.Terminal.Color("g")
	this.Logger.Log("Loading Mirai API Server...")
	//router
	this.Logger.Log("Loading fasthttp router...")
	this.Router = fasthttprouter.New()
	this.Router.GET("/getPluginList", PluginListPage)
	//config
	this.Logger.Log("Loading config")
	this.initConfig()
	//start
	this.Logger.Log("Starting fasthttp Server")
	log.Fatal(fasthttp.ListenAndServe(Config["listenHTTP"], this.Router.Handler))
}

func (this *MiraiApiServer) initConfig() {
	if this.ConfigPath == "" || IsExist(this.ConfigPath) {
		this.ConfigPath = "app.conf"
		this.Logger.Log("Using default config file : app.conf")
	}
	config := make(map[string]string)
	f, err := os.Open(this.ConfigPath)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
		this.Logger.Log("    " + key + "=" + value)
	}
	Config = config
}
