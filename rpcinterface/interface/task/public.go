package task

import (
	"context"
)

const (
	SEND_SMS_TYPE_captcha = 1 //验证码类型
	SEND_SMS_TYPE_plain = 2 //普通类型

	SEND_MAIL_TYPE_captcha = 1 //验证码类型
	SEND_MAIL_TYPE_plain = 2 //普通类型

	DEL_ADD_IS_CRADS = 1 //卡项
	DEL_ADD_IS_PRODUCT =2 //商品
)

//发送短信的参数结构
type SendSmsParams struct {
	Phone string
	MsgStr string
	Type int //短信类型 1=验证码 2=普通消息
	CheckSendNum bool //是否检查手机当日被发送次数
}


type PublicSms interface {
	SendSms( ctx context.Context, args *SendSmsParams, reply *bool )
}

//发送邮件
type SendMailParams struct {
	Email string
	MsgStr string
	Type int //1=短信邮件 2=普通邮件
	CheckSendNum bool //是否检查邮箱当日被发送次数
	Subject string //邮件标题
	FromMail string //发件人email 非必传
	FromName string //发件人名称 非必传
}

type PublicMial interface {
	SendMail( ctx context.Context, args *SendMailParams, reply *bool )
}

//新增卡项或者商品推送到交换机入参
type ArgsAddDelGoods struct {
	Ids []int
	Type int
	CardType int
	/**
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	ITEM_TYPE_single = 1 //单项目
	ITEM_TYPE_sm     = 2 //套餐
	ITEM_TYPE_card   = 3 //综合卡
	ITEM_TYPE_hcard  = 4 //限时卡
	ITEM_TYPE_ncard  = 5 //限次卡
	ITEM_TYPE_hncard = 6 //限时限次卡
	ITEM_TYPE_rcard  = 7 //充值卡
	ITEM_TYPE_icard  = 8 //身份卡
	 */
}
type PublicAddDelGoods interface {
	AddDelGoods( ctx context.Context, args *ArgsAddDelGoods, reply *bool )error
}


