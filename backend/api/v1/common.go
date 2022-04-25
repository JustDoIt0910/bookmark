package v1

import (
	"bookmark/ecode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUserID(c *gin.Context) uint {
	uidRaw, ok := c.Get("userID")
	//如果是未登录状态
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": ecode.ErrAccessDenied,
			"msg": ecode.GetErrMsg(ecode.ErrAccessDenied),
		})
		return 0
	}
	uid, _ := uidRaw.(uint)
	return uid
}
