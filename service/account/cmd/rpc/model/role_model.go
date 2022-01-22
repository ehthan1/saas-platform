package model

type Role struct {
	ID     int64  `gorm:"column:id" json:"id" form:"id"`
	Name   string `gorm:"column:name" json:"name" form:"name"`
	Remark string `gorm:"column:remark" json:"remark" form:"remark"`
	Status int64  `gorm:"column:status" json:"status" form:"status"`
}
