package main

import (
	"github.com/mamoe/mirai-public-api/server"
)

func main() {
	server := server.MiraiApiServer{}
	server.LoadAll()
}
