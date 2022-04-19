package utils

import (
	"bookmark/config"
	"bookmark/ecode"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	jwt.StandardClaims
	Username string  `json:"username"`
	UserID   uint    `json:"uid"`
}

func GenerateToken(username string, uid uint, validDuration int) (int, string) {
	key := []byte(config.GlobalConfig.GetString("jwt.secretKey"))
	ExpireTime := time.Now().Add(time.Duration(validDuration) * time.Hour)
	SetClaim := MyClaims{
		jwt.StandardClaims{
			ExpiresAt: ExpireTime.Unix(),
			Issuer: "bookmark",
		},
		username,
		uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaim)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return ecode.ErrTokenGenerateFail, ""
	}
	return ecode.SUCCESS, signedToken
}

func ParseToken(token string) (int, *MyClaims) {
	key := []byte(config.GlobalConfig.GetString("jwt.secretKey"))
	tokenClaims, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if tokenClaims != nil {
		if myClaims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return ecode.SUCCESS, myClaims
		} else {
			return ecode.ErrInvalidToken, nil
		}
	}
	return ecode.ErrInvalidToken, nil
}