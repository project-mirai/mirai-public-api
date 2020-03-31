package server

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func PluginListPage(ctx *fasthttp.RequestCtx) {
	//paths, err := walkDir(GetMiraiApiServer().Config["basicRepoPath"])//+ "/" + GetMiraiApiServer().Config["pluginListPath"])

	res := ResponseInfo{
		Success: true,
		Info:    "success",
		Result:  GetService().PluginBasicJsonInfo,
	}

	resp, _ := json.Marshal(res)
	fmt.Fprintf(ctx, string(resp))
}
