package v1

import (
	"bookmark/dao"
	"bookmark/ecode"
	"bookmark/model"
	"bookmark/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateCategory(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		return
	}
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	category.Uid = uid
	code := service.CreateCategory(&category)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": category,
		"msg": ecode.GetErrMsg(code),
	})
}

func GetCategories(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		return
	}
	pids := c.Param("pid")
	var pid = 0
	if len(pids) > 1 {
		pid, _ = strconv.Atoi(pids[1:])
	}
	code, categories := service.GetCategories(uid, uint(pid))
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": categories,
		"msg": ecode.GetErrMsg(code),
	})
}

func UpdateCategory(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		return
	}
	var (
		id = 0
		code = ecode.SUCCESS
	)
	idStr := c.Param("id")
	//更新单个
	if len(idStr) > 1 {
		id, _ = strconv.Atoi(idStr[1:])
		var update dao.CategoryUpdateStruct
		err := c.ShouldBindJSON(&update)
		if err != nil {
			code = ecode.ErrBadRequest
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg": ecode.GetErrMsg(code),
			})
			return
		}
		code = service.UpdateCategory(uid, uint(id), &update)
	} else {
		//批量更新
		var batchUpdate service.CateBatchUpdate
		err := c.ShouldBindJSON(&batchUpdate)
		if err != nil {
			code = ecode.ErrBadRequest
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg": ecode.GetErrMsg(code),
			})
			return
		}
		code = service.BatchUpdateCategory(uid, &batchUpdate)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": ecode.GetErrMsg(code),
	})
}

func RemoveCategory(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		return
	}
	var (
		id = 0
		code = ecode.SUCCESS
	)
	idStr := c.Param("id")
	var remove service.CateBatchRemove
	err := c.ShouldBindJSON(&remove)
	fmt.Println(remove)
	if err != nil {
		code = ecode.ErrBadRequest
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg": ecode.GetErrMsg(code),
		})
		return
	}
	//删除单个
	if len(idStr) > 1 {
		id, _ = strconv.Atoi(idStr[1:])
		code = service.RemoveCategory(uid, uint(id), remove.RetainContent)
	} else {
		//批量删除
		code = service.BatchRemoveCategory(uid, &remove)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": ecode.GetErrMsg(code),
	})
}