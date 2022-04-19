package v1

import (
	"bookmark/ecode"
	"bookmark/model"
	"bookmark/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateBookmark(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		return
	}
	var bookmark model.Bookmark
	_ = c.ShouldBindJSON(&bookmark)
	bookmark.Uid = uid
	code := service.CreateBookmark(&bookmark)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": bookmark,
		"msg": ecode.GetErrMsg(code),
	})
}

func GetBookmarks(c *gin.Context)  {
	uid := getUserID(c)
	if uid == 0 {
		return
	}
	cids := c.Param("cid")
	var cid = 0
	if len(cids) > 1 {
		cid, _ = strconv.Atoi(cids[1:])
	}
	code, bookmarks := service.GetBookmarks(uid, uint(cid))
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": bookmarks,
		"msg": ecode.GetErrMsg(code),
	})
}

func UpdateBookmark(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		return
	}
	bidStr := c.Param("bid")
	var code = ecode.SUCCESS
	//单个更改
	if len(bidStr) > 1 {
		bid, _ := strconv.Atoi(bidStr[1:])
		var bookmark model.Bookmark
		_ = c.ShouldBindJSON(&bookmark)
		bookmark.ID = uint(bid)
		code, _ = service.UpdateBookmark(uid, &bookmark)
	} else {
		//批量更改
		var batchUpdate service.BookmarkBatchUpdate
		_ = c.ShouldBindJSON(&batchUpdate)
		code = service.BatchUpdateBookmarks(uid, &batchUpdate)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": ecode.GetErrMsg(code),
	})
}

func RemoveBookmark(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		return
	}
	var remove service.BookmarkBatchRemove
	_ = c.ShouldBindJSON(&remove)
	code := service.BatchRemoveBookmarks(uid, &remove)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": ecode.GetErrMsg(code),
	})
}

// GetRecycleBin 获取回收站中的数据
func GetRecycleBin(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		return
	}
	code, bookmarks := service.GetRecycleBin(uid)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": bookmarks,
		"msg": ecode.GetErrMsg(code),
	})
}

// RestoreRecycleBin 批量还原回收站数据
func RestoreRecycleBin(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		return
	}
	var (
		code = ecode.SUCCESS
		restore service.RestoreStruct
	)
	err := c.ShouldBindJSON(&restore)
	if err != nil {
		code = ecode.ErrBadRequest
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg": ecode.GetErrMsg(code),
		})
		return
	}
	code = service.RestoreRecycleBin(uid, &restore)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": ecode.GetErrMsg(code),
	})
}

// ClearRecycleBin 清空回收站
func ClearRecycleBin(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		return
	}
	code := service.ClearRecycleBin(uid)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": ecode.GetErrMsg(code),
	})
}