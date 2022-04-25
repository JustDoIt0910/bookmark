package cron

import (
	"bookmark/db"
	"bookmark/logger"
	"bookmark/model"
)

// ClearRecycleBin 每天更新回收站
func ClearRecycleBin() {
	rows, err := db.DB.Model(&model.Bookmark{}).Where("deleted = true").Rows()
	if err != nil {
		return
	}
	var bookmark model.Bookmark
	for rows.Next() {
		err = db.DB.ScanRows(rows, &bookmark)
		if err != nil {
			logger.SugarLogger.Error(err)
			return
		}
		if bookmark.LifeTime == 0 {
			err = db.DB.Delete(&bookmark).Error
			if err != nil {
				logger.SugarLogger.Error(err)
			}
		} else {
			err = db.DB.Model(&bookmark).Update("life_time", bookmark.LifeTime - 1).Error
			if err != nil {
				logger.SugarLogger.Error(err)
			}
		}
	}
}