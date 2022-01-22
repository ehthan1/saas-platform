package model

type Menu struct {
	ID         int64  `gorm:"column:id" json:"id" form:"id"`
	Name       string `gorm:"column:name" json:"name" form:"name"`
	Parentid   int64  `gorm:"column:parentid" json:"parentid" form:"parentid"`
	Parentpath string `gorm:"column:parentpath" json:"parentpath" form:"parentpath"`
	Status     int64  `gorm:"column:status" json:"status" form:"status"`
	Icon       string `gorm:"column:icon" json:"icon" form:"icon"`
}
