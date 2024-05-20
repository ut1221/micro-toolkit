package define

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	UserInfo
	jwt.RegisteredClaims
}

type UserInfo struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}

const (
	JwtSecret = "micro-platform"
)
