package server

import (
	"bookmark/cache"
	"bookmark/config"
	"bookmark/db"
	"bookmark/server/cron"
	"bookmark/utils"
	"github.com/gin-gonic/gin"
)

var BookmarkServer *Server

type Server struct {
	Router        *gin.Engine
}

func InitServer()  {
	BookmarkServer = &Server{}
	db.InitDB()
	BookmarkServer.InitRouter()
	err := cache.InitRedisClient()
	if err != nil {
		panic(err)
	}
	err = utils.InitEmailSender(10)
	if err != nil {
		panic(err)
	}
	err = cron.Init()
	if err != nil {
		panic(err)
	}
}

func (server *Server) Run()  {
	_ = server.Router.Run(":" + config.GlobalConfig.GetString("server.port"))
}