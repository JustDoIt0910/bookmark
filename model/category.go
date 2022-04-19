package model

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"ID"`
	Uid  uint   `json:"uid"`
	Pid  uint   `json:"pid"`
	Name string `gorm:"type:varchar(20)" json:"name"`
	Pic  string `gorm:"type:varchar(50)" json:"pic"`
}
