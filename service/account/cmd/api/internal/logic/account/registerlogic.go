package account

import (
	"context"

	"github.com/portgas-x/saas-platform/common/utils"
	"github.com/portgas-x/saas-platform/common/xerror"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/svc"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/types"
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/accountclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterRequest) (*types.RegisterResponse, error) {

	resp, err := l.svcCtx.AccountRpc.Register(l.ctx, &accountclient.RegisterRequest{
		Account:  req.Account,
		Nickname: req.NickName,
		Password: req.Password,
	})

	if err != nil {
		return nil, xerror.ErrorMsg(err)
	}

	token, err := utils.BuildToken(l.svcCtx.Config.Auth.AccessSecret, map[string]interface{}{
		"UUID": resp.UUID,
	}, int64(l.svcCtx.Config.Auth.AccessExpire))

	if err != nil {
		return nil, err
	}

	response := types.AccountReply{
		JwtToken: types.JwtToken{
			AccessToken:  token,
			AccessExpire: int64(l.svcCtx.Config.Auth.AccessExpire),
		},
		Account:  resp.Account,
		NickName: resp.NickName,
	}

	return &types.RegisterResponse{response}, nil
}
