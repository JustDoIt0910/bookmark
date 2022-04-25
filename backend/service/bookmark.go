package service

import (
	"bookmark/cache"
	"bookmark/dao"
	"bookmark/ecode"
	"bookmark/model"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/antchfx/htmlquery"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func CreateBookmark(bookmark *model.Bookmark) int {
	if bookmark.Cid != 0 {
		code, uid := dao.GetUserIDOfCategory(bookmark.Cid)
		if code != ecode.SUCCESS || uid != bookmark.Uid {
			return ecode.FAIL
		}
	}
	//没有收到网页标题，就进行一个爬
	if bookmark.Title == "" {
		err, title := GetPageTitle(bookmark.URL)
		if err != nil {
			return ecode.ErrGetPageTitleFailed
		}
		bookmark.Title = title
	}
	//进行一个icon的爬
	err, icon := GetPageIcon(bookmark.URL)
	if err == nil {
		bookmark.Icon = []byte(icon)
	}

	code := dao.CreateBookmark(bookmark)
	//删除缓存
	ClearCache(bookmark.Uid, bookmark.Cid)
	return code
}

func GetBookmarks(uid uint, cid uint) (int, []*model.BookmarkResp) {
	var (
		code = ecode.SUCCESS
		resp []*model.BookmarkResp
	)
	countKey, err := cache.Redis.BookmarkKeyGen(uid, cid)
	if err != nil {
		return ecode.ErrRedisKeyGenFailed, nil
	}
	countStr, _ := cache.Redis.GetString(countKey)
	//如果countStr查找不到，说明需要重新查询数据库，更新缓存
	if countStr == "" {
		//fmt.Println("come from mysql")
		code, resp = dao.GetBookmarks(uid, cid)
		if code != ecode.SUCCESS {
			return code, nil
		}
		var (
			keys []string
			objs []map[string]interface{}
		)
		for index, bookmark := range resp {
			key, err := cache.Redis.BookmarkKeyGen(uid, cid, index)
			if err != nil {
				return ecode.ErrRedisKeyGenFailed, nil
			}
			keys = append(keys, key)
			objs = append(objs, map[string]interface{}{
				"ID":    bookmark.ID,
				"title": bookmark.Title,
				"url":   bookmark.URL,
				"icon":  bookmark.Icon,
			})
		}
		cache.Redis.SetString(countKey, strconv.Itoa(len(resp)), 0)
		err = cache.Redis.BatchSetHash(keys, objs)
		if err != nil {
			cache.Redis.Delete(countKey)
			return ecode.ErrCacheInsertFailed, nil
		}
	} else { //若缓存中的数据有效
		//fmt.Println("come from cache")
		var (
			keys    []string
			results []map[string]string
		)
		count, _ := strconv.Atoi(countStr)
		for i := 0; i < count; i++ {
			key, err := cache.Redis.BookmarkKeyGen(uid, cid, i)
			if err != nil {
				return ecode.ErrRedisKeyGenFailed, nil
			}
			keys = append(keys, key)
		}
		results, err = cache.Redis.BatchGetHash(keys)
		if err != nil {
			return ecode.ErrCacheGetFailed, nil
		}
		for _, res := range results {
			id, _ := strconv.Atoi(res["ID"])
			resp = append(resp, &model.BookmarkResp{
				ID:    uint(id),
				Title: res["title"],
				URL:   res["url"],
				Icon:  res["icon"],
			})
		}
	}
	return code, resp
}

func UpdateBookmark(uid uint, bookmark *model.Bookmark) (int, *model.Bookmark) {
	if bookmark.Cid > 0 {
		code, _uid := dao.GetUserIDOfCategory(bookmark.Cid)
		if code != ecode.SUCCESS || _uid != uid {
			return ecode.FAIL, nil
		}
	}
	//如果修改了书签的URL,需要重新获取网页title和icon
	if bookmark.URL != "" {
		if bookmark.Title == "" {
			err, title := GetPageTitle(bookmark.URL)
			if err != nil {
				return ecode.ErrGetPageTitleFailed, nil
			}
			bookmark.Title = title
		}
		err, icon := GetPageIcon(bookmark.URL)
		if err == nil {
			bookmark.Icon = []byte(icon)
		}
	}
	code := dao.UpdateBookmark(uid, bookmark)
	return code, bookmark
}

/******************************批量操作**********************************/

type BookmarkBatchRemove struct {
	Ids []uint `json:"ids"`
}

func BatchRemoveBookmarks(uid uint, remove *BookmarkBatchRemove) int {
	var code = ecode.SUCCESS
	code = dao.RemoveBookmarks(uid, remove.Ids)
	if code != ecode.SUCCESS {
		return code
	}
	//删除缓存
	code = CacheClearAll(uid)
	return code
}

type BookmarkBatchUpdate struct {
	Ids        []uint `json:"ids"`
	CategoryID uint   `json:"cid"`
}

func BatchUpdateBookmarks(uid uint, update *BookmarkBatchUpdate) int {
	var code = ecode.SUCCESS
	code = dao.UpdateBookmarks(uid, update.Ids, update.CategoryID)
	if code != ecode.SUCCESS {
		return code
	}
	//删除缓存
	code = CacheClearAll(uid)
	return code
}

// GetRecycleBin 获取回收站中数据
func GetRecycleBin(uid uint) (int, []*model.BookmarkResp) {
	return dao.GetInRecycleBin(uid)
}

type RestoreStruct struct {
	Ids []uint `json:"ids"`
}

// RestoreRecycleBin 还原
func RestoreRecycleBin(uid uint, restore *RestoreStruct) int {
	return dao.BatchRestoreFromBin(uid, restore.Ids)
}

// ClearRecycleBin 清空
func ClearRecycleBin(uid uint) int {
	return dao.ClearRecycleBin(uid)
}

/******************************缓存操作*********************************/

// ClearCache 清除指定uid和cid对应的缓存数据，便于下次获取时进行更新
func ClearCache(uid, cid uint) int {
	countKey, err := cache.Redis.BookmarkKeyGen(uid, cid)
	if err != nil {
		return ecode.ErrRedisKeyGenFailed
	}
	countStr, _ := cache.Redis.GetString(countKey)
	count, _ := strconv.Atoi(countStr)
	var keys []string
	for i := 0; i < count; i++ {
		key, _ := cache.Redis.BookmarkKeyGen(uid, cid, i)
		keys = append(keys, key)
		err = cache.Redis.BatchDelete(keys)
		if err != nil {
			return ecode.ErrCacheDeleteFailed
		}
	}
	cache.Redis.Delete(countKey)
	return ecode.SUCCESS
}

// CacheClearAll 清除某个用户的所有缓存
func CacheClearAll(uid uint) int {
	pattern := fmt.Sprintf("bookmark:%d:*", uid)
	keys, err := cache.Redis.GetKeys(pattern)
	if err != nil {
		return ecode.ErrGetRedisKeyFailed
	}
	err = cache.Redis.BatchDelete(keys)
	if err != nil {
		return ecode.ErrCacheDeleteFailed
	}
	return ecode.SUCCESS
}

/**********************************************************************/

func GetPageTitle(url string) (error, string) {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return err, ""
	}
	title := htmlquery.Find(doc, "//title")
	if len(title) != 1 {
		return errors.New("Cannot find page title. "), ""
	}
	return nil, htmlquery.InnerText(title[0])
}

func GetPageIcon(url string) (error, string) {
	protocol := strings.Split(url, "//")[0]
	domain := strings.Split(strings.Split(url, "//")[1], "/")[0]
	iconUrl := protocol + "//" + domain + "/favicon.ico"
	resp, err := http.Get(iconUrl)
	if err != nil {
		return err, ""
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}
	base64Str := base64.StdEncoding.EncodeToString(body)
	return nil, base64Str
}
