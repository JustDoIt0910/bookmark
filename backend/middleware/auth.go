package middleware

import (
	"bookmark/ecode"
	"bookmark/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": ecode.ErrTokenNotFound,
				"msg": ecode.GetErrMsg(ecode.ErrTokenNotFound),
			})
			c.Abort()
			return
		}
		tokenField := strings.SplitN(tokenHeader, " ", 2)
		if len(tokenField) != 2 || tokenField[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"code": ecode.ErrInvalidToken,
				"msg": ecode.GetErrMsg(ecode.ErrInvalidToken),
			})
			c.Abort()
			return
		}
		token := tokenField[1]
		code, myClaims := utils.ParseToken(token)
		if code != ecode.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg": ecode.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > myClaims.ExpiresAt {
			code = ecode.ErrTokenExpired
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg": ecode.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", myClaims.Username)
		c.Set("userID", myClaims.UserID)
		c.Next()
	}
}
