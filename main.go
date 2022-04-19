package main

import (
	"bookmark/config"
	"bookmark/logger"
	"bookmark/server"
)

func main()  {
	logger.InitLogger()
	config.InitConfig()
	server.InitServer()
	server.BookmarkServer.Run()
}