package main

import (
	"github.com/mamoe/mirai-public-api/apiserver"
)

func main() {
	server := apiserver.MiraiApiServer{}
	server.LoadAll()
}
