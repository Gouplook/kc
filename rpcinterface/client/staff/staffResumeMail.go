package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/staff"
)

type ResumeMail struct {
	client.Baseclient
}

func (r *ResumeMail) Init() *ResumeMail {
	r.ServiceName = "rpc_staff"
	r.ServicePath = "ResumeMail"
	return r
}

//简历信箱列表
func (r *ResumeMail) GetResumeMailList(ctx context.Context, args *staff.ArgsGetResumeMailList, reply *staff.ReplyGetResumeMailList) (err error) {
	return r.Call(ctx, "GetResumeMailList", args, reply)
}

//邀约面试入参
func (r *ResumeMail) InviteInterview(ctx context.Context, args *staff.ArgsInviteInterview, reply *bool) (err error) {
	return r.Call(ctx, "InviteInterview", args, reply)
}
