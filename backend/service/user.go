package service

import (
	"bookmark/cache"
	"bookmark/config"
	"bookmark/dao"
	"bookmark/ecode"
	"bookmark/model"
	"bookmark/utils"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

// CodeExpireTime 验证码有效时间5分钟
const CodeExpireTime = 5


type Userinfo struct {
	Username string  `json:"username"`
	Email    string	  `json:"email"`
}

func GetVerificationCode(info *Userinfo) int {
	//检查用户名是否被占用
	if dao.CheckUsernameExist(info.Username) {
		return ecode.ErrUsernameUsed
	}
	code := utils.NewCode()
	cache.Redis.SetString(info.Username, code, CodeExpireTime * time.Minute)
	utils.EmailSender.Send(
			config.GlobalConfig.GetString("email.sender"),
			info.Email,
			"bookmark 邮箱验证",
			code,
		)
	return ecode.SUCCESS
}

func Register(code string, user *model.User) int {
	cachedCode, _ := cache.Redis.GetString(user.Username)
	//验证码过期
	if cachedCode == "" {
		return ecode.ErrVerificationCodeExpired
	}
	//验证码不正确
	if strings.Compare(code, cachedCode) != 0 {
		return ecode.ErrVerificationCodeWrong
	}
	//密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return ecode.FAIL
	}
	user.Password = string(hashedPassword)
	return dao.AddUser(user)
}

func Login(username string, password string) (int, string) {
	code, user := dao.GetUserByUsername(username)
	if code != ecode.SUCCESS {
		return code, ""
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return ecode.ErrWrongPassword, ""
	}
	code, token := utils.GenerateToken(username, user.ID,  10)
	if code != ecode.SUCCESS {
		return code, ""
	}
	return code, token
}