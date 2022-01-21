package logic

import (
	"context"
	"errors"

	"github.com/portgas-x/saas-platform/common/utils"
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/account"
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/internal/svc"
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/model"
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/hints"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *account.RegisterRequest) (*account.Response, error) {
	accountResult, err := GetAccount(&model.Account{Account: in.Account}, l.svcCtx.Db)
	if err != nil {
		return nil, errors.New("account check error")
	}
	if accountResult.Account != "" {
		return nil, errors.New("account exists")
	}
	password, err := utils.HashPassword(in.Password)
	if err != nil {
		return nil, errors.New("password generate error")
	}
	newAccount := &model.Account{
		Account:  in.Account,
		Password: password,
		NickName: in.Nickname,
	}

	l.svcCtx.Db.Where(&model.Account{Account: in.Account}).FirstOrCreate(&newAccount)

	return &account.Response{Account: newAccount.Account, NickName: newAccount.NickName, UUID: newAccount.UUID}, nil
}

func GetAccount(accountModel *model.Account, db *gorm.DB) (queryResult model.Account, err error) {
	query := db.Clauses(hints.UseIndex("uniq_account")).Table("account")
	if accountModel.Account != "" {
		query.Where(&model.Account{Account: accountModel.Account})
	}
	if accountModel.UUID != "" {
		query.Where(&model.Account{UUID: accountModel.UUID})
	}
	result := query.Limit(1).Find(&queryResult)
	return queryResult, result.Error
}
