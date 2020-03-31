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

func PluginDetailedInfoPage(ctx *fasthttp.RequestCtx) {
	if value, ok := GetService().PluginInfoMap[string(ctx.QueryArgs().Peek("name"))]; ok {
		fmt.Fprintf(ctx, value.JsonFile)
		return
	}
	fmt.Fprint(ctx, "err:not found")
	ctx.Response.SetStatusCode(404)
}
