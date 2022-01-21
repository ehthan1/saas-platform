package logic

import (
	"context"
	"errors"
	"time"

	"github.com/portgas-x/saas-platform/common/utils"
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/account"
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/internal/svc"
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/model"
	"gorm.io/gorm"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *account.LoginRequest) (*account.Response, error) {
	accountResult, err := GetAccount(&model.Account{Account: in.Account}, l.svcCtx.Db)
	if err != nil {
		return nil, errors.New("account check error")
	}
	if accountResult.Account == "" {
		return nil, errors.New("account not exists")
	}

	match := utils.CheckPasswordHash(in.Password, accountResult.Password)

	if !match {
		return nil, errors.New("password error")
	}

	if err := UpdateAccount(&model.Account{UUID: accountResult.UUID, LastLoggedIn: time.Now()}, l.svcCtx.Db); err != nil {
		return nil, errors.New("update account login error")
	}

	return &account.Response{Account: accountResult.Account, NickName: accountResult.NickName, UUID: accountResult.UUID}, nil
}

func UpdateAccount(accountModel *model.Account, db *gorm.DB) error {
	account := model.Account{}
	result := db.Debug().Model(&account).Where("uuid = ?", accountModel.UUID).Select("last_logged_in").Updates(&accountModel)
	return result.Error
}
