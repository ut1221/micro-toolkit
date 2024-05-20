package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/ut1221/micro-toolkit/define"
	"strings"
	"time"
)

// GetMd5
//
//	@Description: 生成 md5
//	@Author PTJ 2024-05-14 17:58:28
//	@param s
//	@return string
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GenerateToken
//
//	@Description: 生成 token
//	@Author PTJ 2024-05-14 17:58:36
//	@param identity
//	@param name
//	@return string
//	@return error
func GenerateToken(userId, username string, expireTime time.Duration, secret string) (string, error) {
	var userInfo = define.UserInfo{
		UserId:   userId,
		Username: username,
	}
	UserClaim := &define.UserClaim{
		UserInfo: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
//
//	@Description: 解析 token
//	@Author PTJ 2024-05-14 17:58:44
//	@param tokenString
//	@return *define.UserClaim
//	@return error
func AnalyseToken(tokenString string, secret string) (*define.UserClaim, error) {
	userClaim := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// GetUUID
//
//	@Description: 生成唯一码
//	@Author PTJ 2024-05-14 17:58:52
//	@return string
func GetUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
