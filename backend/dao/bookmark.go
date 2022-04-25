package dao

import (
	"bookmark/db"
	"bookmark/ecode"
	"bookmark/model"
	"github.com/jinzhu/gorm"
)

func CreateBookmark(bookmark *model.Bookmark) int {
	err := db.DB.Model(&model.Bookmark{}).Create(&bookmark).Error
	if err != nil {
		return ecode.FAIL
	}
	return ecode.SUCCESS
}

//GetBookmarks 获取书签信息(cid为0表示未分类书签)
func GetBookmarks(uid, cid uint) (int, []*model.BookmarkResp) {
	var (
		bookmarks []*model.Bookmark
		resp      []*model.BookmarkResp
	)
	err := db.DB.Model(&model.Bookmark{}).Where("deleted = false and uid = ? and cid = ?", uid, cid).Find(&bookmarks).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return ecode.FAIL, nil
	}
	for _, b := range bookmarks {
		resp = append(resp, b.MakeRespStruct())
	}
	return ecode.SUCCESS, resp
}

//UpdateBookmark 更新单条书签
func UpdateBookmark(uid uint, bookmark *model.Bookmark) int {
	var bm model.Bookmark
	err := db.DB.Model(&model.Bookmark{}).Where("uid = ? and id = ?", uid, bookmark.ID).First(&bm).Error
	if err != nil || bm.Uid != uid {
		return ecode.FAIL
	}
	//cid小于0说明不修改该字段
	if bookmark.Cid >= 0 {
		bm.Cid = bookmark.Cid
	}
	//其他字段同理
	if bookmark.Title != "" {
		bm.Title = bookmark.Title
	}
	if bookmark.URL != "" {
		bm.URL = bookmark.URL
	}
	if bookmark.Icon != nil {
		bm.Icon = bookmark.Icon
	}
	err = db.DB.Model(&model.Bookmark{}).Where("id = ?", bookmark.ID).Updates(map[string]interface{}{
		"title": bm.Title,
		"url":   bm.URL,
		"cid":   bm.Cid,
		"icon":  bm.Icon,
	}).Error
	if err != nil {
		return ecode.FAIL
	}
	return ecode.SUCCESS
}

//UpdateBookmarks 批量更新书签(只能批量移动, 将id在ids中的书签移动到id为cid的标签下)
func UpdateBookmarks(uid uint, ids []uint, cid uint) int {
	err := db.DB.Model(&model.Bookmark{}).Where("uid = ?", uid).Where(ids).Update("cid", cid).Error
	if err != nil {
		return ecode.FAIL
	}
	return ecode.SUCCESS
}

//RemoveBookmarks 批量删除id在ids内的书签(放入回收站中)
func RemoveBookmarks(uid uint, ids []uint) int {
	err := db.DB.Model(&model.Bookmark{}).Where("uid = ?", uid).Where(ids).Updates(map[string]interface{}{
		"deleted":   true,
		"life_time": 10, //回收站保留时间10天
	}).Error
	if err != nil {
		return ecode.FAIL
	}
	return ecode.SUCCESS
}

/******************************回收站操作********************************/

// GetInRecycleBin 获取回收站中的书签
func GetInRecycleBin(uid uint) (int, []*model.BookmarkResp) {
	var (
		bookmarks []*model.Bookmark
		resp      []*model.BookmarkResp
	)
	err := db.DB.Model(&model.Bookmark{}).Where("deleted = true and uid = ?", uid).Find(&bookmarks).Error
	if err != nil {
		return ecode.FAIL, nil
	}
	for _, b := range bookmarks {
		resp = append(resp, b.MakeRespStruct())
	}
	return ecode.SUCCESS, resp
}

// BatchRestoreFromBin 批量从回收站中还原数据
func BatchRestoreFromBin(uid uint, ids []uint) int {
	err := db.DB.Model(&model.Bookmark{}).Where("uid = ?", uid).Where(ids).Updates(map[string]interface{}{
		"deleted":   false,
		"life_time": 0,
	}).Error
	if err != nil {
		return ecode.FAIL
	}
	return ecode.SUCCESS
}

// ClearRecycleBin 清空回收站
func ClearRecycleBin(uid uint) int {
	err := db.DB.Model(&model.Bookmark{}).Where("uid = ? and deleted = true", uid).Delete(&model.Bookmark{}).Error
	if err != nil {
		return ecode.FAIL
	}
	return ecode.SUCCESS
}
