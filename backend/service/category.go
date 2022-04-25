package service

import (
	"bookmark/dao"
	"bookmark/ecode"
	"bookmark/model"
	"bookmark/utils/picture"
)

// CreateCategory 创建新标签
func CreateCategory(cate *model.Category) int {
	if cate.Pid > 0 {
		code, uid := dao.GetUserIDOfCategory(cate.Pid)
		if code != ecode.SUCCESS || uid != cate.Uid {
			return ecode.FAIL
		}
	}
	cate.Pic = picture.Get(cate.Name)
	return dao.CreateCategory(cate)
}

// GetCategories 获取标签(uid 用户ID, pid 父级标签ID, 为0表示顶层标签)
func GetCategories(uid uint, pid uint) (int, []*model.Category) {
	return dao.GetCategories(uid, pid, true)
}

// UpdateCategory 更新单个标签(update 更新信息结构体，包含名字和位置)
func UpdateCategory(uid, id uint, update *dao.CategoryUpdateStruct) int {
	if update.Pid > 0 {
		code, _uid := dao.GetUserIDOfCategory(uint(update.Pid))
		if code != ecode.SUCCESS || uid != _uid {
			return ecode.FAIL
		}
	}
	return dao.UpdateCategory(uid, id, update)
}

// RemoveCategory 删除单个标签(id 标签id, retainContent 是否保留内容)
func RemoveCategory(uid, id uint, retainContent bool) int {
	//保留内容，删除标签后，标签内所有内容并入上一级
	if retainContent {
		//获取要删除标签下的所有子标签
		code, subCategories := dao.GetCategories(uid, id, false)
		if code != ecode.SUCCESS {
			return code
		}
		//获取要删除的标签的父标签
		code, parentId := dao.GetParentIDofCategory(id)
		if code != ecode.SUCCESS {
			return code
		}
		//将所有子标签并入删除标签的父标签
		var ids []uint
		for _, sub := range subCategories {
			ids = append(ids, sub.ID)
		}
		code = dao.UpdateCategories(uid, ids, parentId)
		if code != ecode.SUCCESS {
			return code
		}

		//将所有书签并入上一级
		code, bookmarks := dao.GetBookmarks(uid, id)
		if code != ecode.SUCCESS {
			return code
		}
		ids = ids[0:0] //清空切片
		for _, bm := range bookmarks {
			ids = append(ids, bm.ID)
		}
		code = dao.UpdateBookmarks(uid, ids, parentId)
		if code != ecode.SUCCESS {
			return code
		}
		//删除该标签
		code = dao.RemoveOneCategory(id)
		if code != ecode.SUCCESS {
			return code
		}
		ClearCache(uid, id)
		ClearCache(uid, parentId)
	} else { //不保留内容，连同标签内容一并删除
		code := dao.RemoveCategoryAndContent(uid, id)
		if code != ecode.SUCCESS {
			return code
		}
	}
	return ecode.SUCCESS
}

/*********************************批量操作************************************/

// CateBatchUpdate 批量更新(移动位置)标签结构(Ids 要更新的标签id, ParentID 新的父级标签)
type CateBatchUpdate struct {
	Ids      []uint `json:"ids"`
	ParentID uint   `json:"pid"`
}

// CateBatchRemove 批量删除标签结构(Ids 要删除的标签id, RetainContent 是否保留内容)
type CateBatchRemove struct {
	Ids           []uint `json:"ids"`
	RetainContent bool   `json:"retain_content"`
}

// BatchUpdateCategory 批量更新(移动)标签
func BatchUpdateCategory(uid uint, update *CateBatchUpdate) int {
	if update.ParentID > 0 {
		code, _uid := dao.GetUserIDOfCategory(update.ParentID)
		if code != ecode.SUCCESS || uid != _uid {
			return code
		}
	}
	code := dao.UpdateCategories(uid, update.Ids, update.ParentID)
	if code != ecode.SUCCESS {
		return code
	}
	return ecode.SUCCESS
}

// BatchRemoveCategory 批量删除标签
func BatchRemoveCategory(uid uint, remove *CateBatchRemove) int {
	for _, id := range remove.Ids {
		code := RemoveCategory(uid, id, remove.RetainContent)
		if code != ecode.SUCCESS {
			return code
		}
	}
	return ecode.SUCCESS
}
