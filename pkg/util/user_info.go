package util

import (
	"context"
	"errors"
)

var userKey int

type UserInfo struct {
	Name string `json:"name"`
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	info, ok := ctx.Value(userKey).(*UserInfo)
	return info, ok
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	info, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	return info, nil
}

func SetUserInfo(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}
