package v1

import (
	"bookmark/ecode"
	"bookmark/utils/oss"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadBgPictures(c *gin.Context) {
	code := ecode.SUCCESS
	ok := oss.Upload()
	if !ok {
		code = ecode.ErrUploadPictureFailed
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ecode.GetErrMsg(code),
	})
}
