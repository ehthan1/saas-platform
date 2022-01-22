package model

type RoleMenu struct {
	ID     int64 `gorm:"column:id" json:"id" form:"id"`
	Roleid int64 `gorm:"column:roleid" json:"roleid" form:"roleid"`
	Menuid int64 `gorm:"column:menuid" json:"menuid" form:"menuid"`
}
