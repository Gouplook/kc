/**
 * @Author: Gosin
 * @Date: 2020/3/5 16:30
 */
package user

import (
	"context"
)

type Userinfo interface {
	// 用户注册
	UserReg(ctx context.Context, id *int, reply *bool) error
	// 用户信息
	SetUserInfo(ctx context.Context, uid *int, reply *bool) error
}
