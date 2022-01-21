package account

import (
	"context"
	"errors"

	"github.com/portgas-x/saas-platform/common/utils"
	"github.com/portgas-x/saas-platform/common/xerror"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/svc"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/types"
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/accountclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginRequest) (*types.LoginResponse, error) {
	resp, err := l.svcCtx.AccountRpc.Login(l.ctx, &accountclient.LoginRequest{
		Account:  req.Account,
		Password: req.Password,
	})

	if err != nil {
		return nil, xerror.ErrorMsg(err)
	}

	token, err := utils.BuildToken(l.svcCtx.Config.Auth.AccessSecret, map[string]interface{}{
		"UUID": resp.UUID,
	}, int64(l.svcCtx.Config.Auth.AccessExpire))

	if err != nil {
		return nil, errors.New("GenerateToken Error:" + err.Error())
	}

	response := types.AccountReply{
		JwtToken: types.JwtToken{
			AccessToken:  token,
			AccessExpire: int64(l.svcCtx.Config.Auth.AccessExpire),
		},
		Account:  resp.Account,
		NickName: resp.NickName,
	}

	return &types.LoginResponse{response}, nil
}
