package utils

import (
	"context"
	"encoding/json"
	"github.com/ut1221/micro-toolkit/define"
	"github.com/ut1221/micro-toolkit/pkg/constants"
)

func GetUserInfo(ctx context.Context) *define.UserInfo {
	value, _ := json.Marshal(ctx.Value(constants.UserCache))
	var user define.UserClaim
	_ = json.Unmarshal(value, &user)
	return &user.UserInfo
}

func GetUserId(ctx context.Context) string {
	return GetUserInfo(ctx).UserId
}

func GetUserName(ctx context.Context) string {
	return GetUserInfo(ctx).Username
}
