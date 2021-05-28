package user

import "context"

// ReplyUserAccountInfo ReplyUserAccountInfo
type ReplyUserAccountInfo struct {
	Uid   int
	Phone string
}

type UserAccount interface {
	// UserAccountInfo 根据uid获取用户帐号信息
	GetUserAccountInfo(ctx context.Context, uid *int, reply *ReplyUserAccountInfo) error
}
