package middleware

import (
	"context"
	"encoding/json"
	"github.com/ut1221/micro-toolkit/define"
	"github.com/ut1221/micro-toolkit/pkg/constants"
	"github.com/ut1221/micro-toolkit/pkg/response"
	"github.com/ut1221/micro-toolkit/utils"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle() rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get(constants.Authorization) == "" {
				w.WriteHeader(http.StatusOK)
				body := response.Body{Code: http.StatusUnauthorized, Message: constants.TokenNotExist}
				marshal, _ := json.Marshal(body)
				_, _ = w.Write(marshal)
				return
			}
			t := strings.Split(r.Header.Get(constants.Authorization), " ")
			if t[0] != constants.Bearer {
				w.WriteHeader(http.StatusOK)
				body := response.Body{Code: http.StatusUnauthorized, Message: constants.TokenInvalid}
				marshal, _ := json.Marshal(body)
				_, _ = w.Write(marshal)
				return
			}
			claim, err := utils.AnalyseToken(t[1], define.JwtSecret)
			if err != nil {
				w.WriteHeader(http.StatusOK)
				body := response.Body{Code: http.StatusUnauthorized, Message: constants.TokenInvalid}
				marshal, _ := json.Marshal(body)
				_, _ = w.Write(marshal)
				return
			}
			reqCtx := r.Context()
			ctx := context.WithValue(reqCtx, constants.UserCache, claim)
			newReq := r.WithContext(ctx)
			next(w, newReq)
		}
	}
}
