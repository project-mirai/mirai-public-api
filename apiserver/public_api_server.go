package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/mamoe/mirai-public-api/logger"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
)

const (
	BASE_DIR        = "mirai-plugins"
	PLUGIN_LIST_DIR = "plugins"
)

type MiraiApiServer struct {
	Logger *logger.MiraiLogger
	Router *fasthttprouter.Router
}

func (this *MiraiApiServer) LoadAll() {
	//logger
	this.Logger = logger.NewLogger("MiraiAPIServer")
	this.Logger.Log("Loading Mirai API Server")
	//router
	this.Logger.Log("Loading fasthttp router...")
	this.Router = fasthttprouter.New()
	this.Router.GET("/getPluginList", PluginListPage)
	log.Fatal(fasthttp.ListenAndServe(":8181", this.Router.Handler))
}

func PluginListPage(ctx *fasthttp.RequestCtx) {
	paths, err := walkDir(BASE_DIR + "/" + PLUGIN_LIST_DIR)
	res := ResponseInfo{
		Success: err == nil,
		Info:    "success",
		Result:  paths,
	}
	if err != nil {
		res.Result = err
	}

	resp, err := json.Marshal(res)
	fmt.Fprintf(ctx, string(resp))
}

func walkDir(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, file := range files {
		if file.IsDir() {
			result = append(result, file.Name())
		}
	}
	return result, nil
}
