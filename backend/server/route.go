package server

import (
	v1 "bookmark/api/v1"
	"bookmark/middleware"
	"github.com/gin-gonic/gin"
)

func (server *Server) InitRouter() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Cors())

	visitor := r.Group("/api/v1")
	{
		visitor.GET("/ping", v1.Pong)
		visitor.POST("/verificationCode", v1.GetVerificationCode)
		visitor.POST("/register/:code", v1.Register)
		visitor.POST("/login", v1.Login)
	}

	member := r.Group("/api/v1")
	member.Use(middleware.JwtMiddleware())
	{
		member.POST("/category", v1.CreateCategory)
		//获取顶层标签(不带pid参数)或者获取子标签(pid为父标签id)
		member.GET("/categories/*pid", v1.GetCategories)
		//修改一个标签(传id)或者批量移动标签(请求体中传id数组)
		member.PUT("/category/*id", v1.UpdateCategory)
		//删除一个标签(传id)或者批量删除标签(请求体中传id数组)
		member.DELETE("/category/*id", v1.RemoveCategory)

		member.POST("/bookmark", v1.CreateBookmark)
		//获取未分类书签(不带cid参数)或者获取某标签下的书签
		member.GET("/bookmark/*cid", v1.GetBookmarks)
		member.PUT("/bookmark/*bid", v1.UpdateBookmark)
		//没有设计删除单个书签的操作，请求体中传id数组，只有单个元素的时候即为删除单个书签
		member.DELETE("/bookmark", v1.RemoveBookmark)

		//获取回收站中的书签
		member.GET("/recycleBin", v1.GetRecycleBin)
		//批量还原回收站
		member.POST("/recycleBin", v1.RestoreRecycleBin)
		//清空回收站
		member.DELETE("/recycleBin", v1.ClearRecycleBin)

		member.GET("/category_bg_upload", v1.UploadBgPictures)
	}
	server.Router = r
}
