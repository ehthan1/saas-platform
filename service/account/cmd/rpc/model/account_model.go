package model

import (
	"time"

	"github.com/tal-tech/go-zero/core/utils"
	"gorm.io/gorm"
)

type Account struct {
	UUID         string    `gorm:"primaryKey" json:"uuid" form:"uuid"`
	ParentID     string    `gorm:"column:parent_id" json:"parent_id" form:"parent_id"`
	Account      string    `gorm:"index:uniq_account,unique" json:"account" form:"account"`
	Mobile       string    `gorm:"column:mobile" json:"mobile" form:"mobile"`
	Password     string    `gorm:"column:password" json:"password" form:"password"`
	Gender       string    `gorm:"column:gender" json:"gender" form:"gender"`
	Birthdate    time.Time `gorm:"column:birthdate" json:"birthdate" form:"birthdate"`
	Status       int64     `gorm:"column:status" json:"status" form:"status"`
	LastLoginIp  string    `gorm:"column:last_login_ip" json:"last_login_ip" form:"last_login_ip"`
	LastLoggedIn time.Time `gorm:"column:last_logged_in" json:"last_logged_in" form:"last_logged_in"`
	RealName     string    `gorm:"column:real_name" json:"real_name" form:"real_name"`
	NickName     string    `gorm:"column:nick_name" json:"nick_name" form:"nick_name"`
	PausedAt     time.Time `gorm:"column:paused_at" json:"paused_at" form:"paused_at"`
	ClosedAt     time.Time `gorm:"column:closed_at" json:"closed_at" form:"closed_at"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt    time.Time `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at"`
}

func (u *Account) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = utils.NewUuid()
	return
}
func (u *Account) BeforeUpdate(tx *gorm.DB) (err error) {

	return
}
