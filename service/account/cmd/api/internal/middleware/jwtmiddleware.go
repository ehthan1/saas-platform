package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/portgas-x/saas-platform/common/response"
	"github.com/portgas-x/saas-platform/common/utils"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/config"
)

const (
	jwtAudience    = "aud"
	jwtExpire      = "exp"
	jwtId          = "jti"
	jwtIssueAt     = "iat"
	jwtIssuer      = "iss"
	jwtNotBefore   = "nbf"
	jwtSubject     = "sub"
	noDetailReason = "no detail reason"
)

var (
	errInvalidToken = errors.New("invalid auth token")
	errNoClaims     = errors.New("no auth params")
)

type JwtMiddleware struct {
	Config config.Config
	ctx    context.Context
}

func NewJwtMiddleware(c config.Config) *JwtMiddleware {
	return &JwtMiddleware{
		Config: c,
	}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parser := utils.NewTokenParser()
		tok, err := parser.ParseToken(r, m.Config.Auth.AccessSecret, "")

		if err != nil {
			response.Response(w, nil, err)
			return
		}

		if !tok.Valid {
			response.Response(w, nil, errInvalidToken)
			return
		}

		claims, ok := tok.Claims.(jwt.MapClaims)

		if !ok {
			response.Response(w, nil, errNoClaims)
			return
		}

		ctx := r.Context()
		for k, v := range claims {
			switch k {
			case jwtAudience, jwtExpire, jwtId, jwtIssueAt, jwtIssuer, jwtNotBefore, jwtSubject:
				// ignore the standard claims
			default:
				ctx = context.WithValue(ctx, k, v)
			}
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
