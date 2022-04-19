package dao

import (
	"bookmark/db"
	"bookmark/ecode"
	"bookmark/model"
	"errors"
	"github.com/jinzhu/gorm"
)


//CheckUsernameExist 查询用户名是否存在
func CheckUsernameExist(username string) bool {
	var count int
	db.DB.Model(&model.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

//AddUser 添加新用户
func AddUser(user *model.User) int {
	if err := db.DB.Create(user).Error; err != nil {
		return ecode.FAIL
	}
	return ecode.SUCCESS
}


//GetUserById 根据id查询用户
func GetUserById(id int) (int, *model.User) {
	user := model.User{}
	err := db.DB.Model(model.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ecode.ErrUserNotFound, nil
		} else {
			return ecode.FAIL, nil
		}
	}
	return ecode.SUCCESS, &user
}

//GetUserByUsername 根据username查询用户
func GetUserByUsername(username string) (int, *model.User) {
	user := model.User{}
	err := db.DB.Model(model.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ecode.ErrUserNotFound, nil
		} else {
			return ecode.FAIL, nil
		}
	}
	return ecode.SUCCESS, &user
}