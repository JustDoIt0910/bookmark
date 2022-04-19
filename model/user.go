package model

type User struct {
	ID          uint     `gorm:"primary_key;auto_increment" json:"ID"`
	Username    string   `gorm:"type:varchar(50)" json:"username"`
	Password    string   `gorm:"type:varchar(200)" json:"password"`
	Avatar      string   `gorm:"type:varchar(200)" json:"avatar"`
	Email       string   `gorm:"type:varchar(50)" json:"email"`
}
