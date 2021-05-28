package admin

import "context"

type ArgsAddSugger struct {
	Name string //姓名
	Phone string //手机号
	Email string //邮箱
	Content string //反馈内容
}


type Sugger interface {
	AddSugger( ctx context.Context, args *ArgsAddSugger, reply *bool ) error //添加反馈意见
}
