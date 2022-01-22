package model

type AccountRole struct {
	ID        int64 `gorm:"column:id" json:"id" form:"id"`
	Accountid int64 `gorm:"column:accountid" json:"accountid" form:"accountid"`
	Roleid    int64 `gorm:"column:roleid" json:"roleid" form:"roleid"`
}
