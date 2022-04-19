package dao

import (
	"bookmark/db"
	"bookmark/ecode"
	"bookmark/model"
	"github.com/jinzhu/gorm"
)

// CreateCategory 创建一个标签
func CreateCategory(cate *model.Category) int {
	err := db.DB.Create(cate).Error
	if err != nil {
		return ecode.FAIL
	}
	return ecode.SUCCESS
}

// GetCategories 获取标签(pid传0表示获取一级标签,否则表示获取此id下的子标签)
func GetCategories(uid, pid uint, fullInfo bool) (int, []*model.Category) {
	var categories []*model.Category
	var err error
	if fullInfo {
		err = db.DB.Model(&model.Category{}).Where("uid = ? and pid = ?", uid, pid).Find(&categories).Error
	} else {
		err = db.DB.Model(&model.Category{}).Select("id").Where("uid = ? and pid = ?", uid, pid).Find(&categories).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return ecode.FAIL, nil
	}
	return ecode.SUCCESS, categories
}

type CategoryUpdateStruct struct {
	Name string  `json:"name"`
	Pid  int     `json:"pid"`
}

// UpdateCategory 更新一个标签
func UpdateCategory(uid uint, id uint, updateStruct *CategoryUpdateStruct) int {
	var cate model.Category
	var err error
	err = db.DB.Model(&model.Category{}).Where("id = ?", id).First(&cate).Error
	//防止非法更改其他用户的数据
	if err != nil || cate.Uid != uid {
		return ecode.FAIL
	}
	//pid小于0表示不更改此字段
	if updateStruct.Pid >= 0 {
		cate.Pid = uint(updateStruct.Pid)
	}
	//name同理
	if updateStruct.Name != "" {
		cate.Name = updateStruct.Name
	}
	err = db.DB.Model(&model.Category{}).Where("id = ?", id).Updates(map[string]interface{}{
		"pid": cate.Pid,
		"name": cate.Name,
	}).Error
	if err != nil {
		return ecode.FAIL
	}
	return ecode.SUCCESS
}

// UpdateCategories 批量更新标签(只能批量移动，不可批量重命名)
func UpdateCategories(uid uint, ids []uint, pid uint) int {
	err := db.DB.Model(&model.Category{}).Where("uid = ?", uid).Where(ids).Update("pid", pid).Error
	if err != nil {
		return ecode.FAIL
	}
	return ecode.SUCCESS
}

// RemoveOneCategory 删除一个标签(不删除其下的子标签)
func RemoveOneCategory(id uint) int {
	var err error
	err = db.DB.Model(&model.Category{}).Where("id = ?", id).Delete(&model.Category{}).Error
	if err != nil {
		return ecode.FAIL
	}
	return ecode.SUCCESS
}

// RemoveCategoryAndContent 删除一个标签及其内容
func RemoveCategoryAndContent(uid, id uint) int {
	//先删除当前标签
	err := db.DB.Model(&model.Category{}).Where("id = ?", id).Delete(&model.Category{}).Error
	if err != nil {
		return ecode.FAIL
	}
	//再删除当前标签下的书签
	code, bookmarks := GetBookmarks(uid, id)
	if code != ecode.SUCCESS {
		return code
	}
	var ids []uint
	for _, bookmark := range bookmarks {
		ids = append(ids, bookmark.ID)
	}
	code = RemoveBookmarks(uid, ids)
	if code != ecode.SUCCESS {
		return code
	}
	//递归删除其下的所有标签和内容
	code, subs := GetCategories(uid, id, false)
	if code != ecode.SUCCESS {
		return code
	}
	for _, subCate := range subs {
		code = RemoveCategoryAndContent(uid, subCate.ID)
		if code != ecode.SUCCESS{
			return code
		}
	}
	return ecode.SUCCESS
}

func GetUserIDOfCategory(categoryId uint) (int, uint) {
	var category model.Category
	err := db.DB.Model(&model.Category{}).Select("uid").Where("id = ?", categoryId).First(&category).Error
	if err != nil {
		return ecode.FAIL, 0
	}
	return ecode.SUCCESS, category.Uid
}

func GetParentIDofCategory(categoryId uint) (int, uint) {
	var cate model.Category
	err := db.DB.Model(&model.Category{}).Select("pid").Where("id = ?", categoryId).Find(&cate).Error
	if err != nil {
		return ecode.FAIL, 0
	}
	return ecode.SUCCESS, cate.Uid
}