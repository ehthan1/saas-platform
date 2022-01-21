package account

import (
	"net/http"

	"github.com/portgas-x/saas-platform/common/response"
	account "github.com/portgas-x/saas-platform/service/account/cmd/api/internal/logic/account"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/svc"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := account.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(req)
		response.Response(w, resp, err)

	}
}
