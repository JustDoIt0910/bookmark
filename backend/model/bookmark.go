package model

type Bookmark struct {
	ID        uint     `gorm:"primary_key;auto_increment" json:"ID"`
	Title     string   `gorm:"type:varchar(100)" json:"title"`
	URL       string   `gorm:"type:varchar(200)" json:"url"`
	Icon      []byte   `gorm:"type:blob" json:"icon"`
	Uid       uint     `json:"uid"`
	Cid       uint     `json:"cid"`
	Deleted   bool     `json:"deleted"`  //是否已经删除，相当于回收站
	LifeTime  uint     `gorm:"column:life_time" json:"life_time"` //剩余时间，只对已删除书签起作用，减到零该书签被删除
}

func (b *Bookmark) MakeRespStruct() *BookmarkResp {
	return &BookmarkResp{
		ID: b.ID,
		Title: b.Title,
		URL: b.URL,
		Icon: string(b.Icon),
		LifeTime: b.LifeTime,
	}
}

type BookmarkResp struct {
	ID     		uint     `json:"ID"`
	Title  		string   `json:"title"`
	URL    		string   `json:"url"`
	Icon   		string   `json:"icon"`
	LifeTime	uint     `json:"life_time"`
}