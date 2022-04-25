package ecode

const (
	SUCCESS       = 200
	ErrBadRequest = 400
	FAIL          = 500

	ErrUsernameUsed            = 1001
	ErrUserNotFound            = 1002
	ErrVerificationCodeWrong   = 1003
	ErrVerificationCodeExpired = 1004
	ErrWrongPassword           = 1005
	ErrTokenGenerateFail       = 1006
	ErrTokenNotFound           = 1007
	ErrInvalidToken            = 1008
	ErrTokenExpired            = 1009
	ErrAccessDenied            = 1010

	ErrGetPageTitleFailed = 2001

	ErrRedisKeyGenFailed = 3001
	ErrCacheInsertFailed = 3002
	ErrCacheGetFailed    = 3003
	ErrCacheDeleteFailed = 3004
	ErrGetRedisKeyFailed = 3005

	ErrUploadPictureFailed = 4001
)

var codeMsg = map[int]string{
	SUCCESS:       "OK",
	FAIL:          "FAIL",
	ErrBadRequest: "请求格式错误",

	ErrUsernameUsed:            "用户名已存在",
	ErrUserNotFound:            "用户不存在",
	ErrVerificationCodeWrong:   "验证码错误",
	ErrVerificationCodeExpired: "验证码过期",
	ErrWrongPassword:           "密码错误",
	ErrTokenGenerateFail:       "获取token失败",
	ErrTokenNotFound:           "token未携带",
	ErrInvalidToken:            "token无效",
	ErrTokenExpired:            "token已过期",
	ErrAccessDenied:            "用户未登录",

	ErrGetPageTitleFailed: "自动获取网页标题失败",

	ErrRedisKeyGenFailed: "redis key生成失败",
	ErrCacheInsertFailed: "缓存插入数据失败",
	ErrCacheGetFailed:    "获取缓存数据失败",
	ErrCacheDeleteFailed: "缓存数据删除失败",
	ErrGetRedisKeyFailed: "获取缓存key失败",

	ErrUploadPictureFailed: "背景图上传失败",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
