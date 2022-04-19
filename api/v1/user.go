package v1

import (
	"bookmark/ecode"
	"bookmark/model"
	"bookmark/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVerificationCode(c *gin.Context) {
	var info service.Userinfo
	_ = c.ShouldBindJSON(&info)
	code := service.GetVerificationCode(&info)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": info,
		"msg": ecode.GetErrMsg(code),
	})
}

func Register(c *gin.Context) {
	code := ecode.SUCCESS
	vCode := c.Param("code")
	var user model.User
	_ = c.ShouldBindJSON(&user)
	if vCode == "" {
		code = ecode.ErrBadRequest
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg": ecode.GetErrMsg(code),
		})
		return
	}
	code = service.Register(vCode, &user)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": user,
		"msg": ecode.GetErrMsg(code),
	})
}

func Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	fmt.Println(user)
	code, token := service.Login(user.Username, user.Password)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"token": token,
		"msg": ecode.GetErrMsg(code),
	})
}