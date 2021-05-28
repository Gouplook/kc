package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

/*
招聘/简历信箱
*/

const (
	InviteStatusNo  = 1 //待邀约
	InviteStatusYes = 2 //已邀约
)

//获取简历信箱列表入参
type ArgsGetResumeMailList struct {
	common.BsToken
	common.Paging
	PositionId int //应聘职位id
	Status     int //状态
}
type GetResumeMailListBase struct {
	Id           int
	Uid          int
	Name         string //面试者姓名
	Phone        string //面试手机号
	PositionId   int    //应聘职位id
	PositionName string //应聘职位
	Status       int    //状态
	Ctime        int64  //投递时间
	CtimeStr     string
}

//获取简历信箱列表出参
type ReplyGetResumeMailList struct {
	TotalNum int
	Lists    []GetResumeMailListBase
}

//邀约面试入参
type ArgsInviteInterview struct {
	common.BsToken
	Id               int
	Phone            string //手机号
	SmsContent       string //短信内容
	Contact          string //联系人
	ContactCall      string //联系方式
	InterviewAddress string //面试地址
}

type ResumeMail interface {
	//简历信箱列表
	GetResumeMailList(ctx context.Context, args *ArgsGetResumeMailList, reply *ReplyGetResumeMailList) error
	//邀约面试入参
	InviteInterview(ctx context.Context, args *ArgsInviteInterview, reply *bool) error
}
