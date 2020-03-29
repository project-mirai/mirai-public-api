package server

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func PluginListPage(ctx *fasthttp.RequestCtx) {
	paths, err := walkDir(Config["basicRepoPath"] + "/" + Config["pluginListPath"])
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
