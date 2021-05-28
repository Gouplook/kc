package public

import "context"

//邮件相关接口
//@author yangzhiwu<wyz@900sui.com>
//@date 2020-03-09 14:55:23
//注意：此服务只能任务服务调用

type SendMailCaptchaParams struct {
	Email string //邮箱
	CheckSendNum bool //是否检查手机当日被发送次数
	FromMail string //发件人email 非必传
	FromName string //发件人名称 非必传
}

//验证邮箱验证码参数
type CheckMailCaptchaParams struct {
	Email string //邮箱
	Captcha string //验证码
}

//发送邮件参数
type SendMailParams struct {
	Email string //邮箱
	MsgStr string //邮件内容
	Subject string //邮件标题
	FromMail string //发件人email 非必传
	FromName string //发件人名称 非必传
}

type Mail interface {
	//添加邮件验证码任务到mq队列
	SendCaptcha2Mq( ctx context.Context, args *SendMailCaptchaParams, reply *bool ) error
	//发送邮件验证码,MQ消费者调用
	SendCaptcha( ctx context.Context, args *SendMailCaptchaParams, reply *bool ) error
	//验证邮件验证码
	CheckCaptcha( ctx context.Context, args *CheckMailCaptchaParams, reply *bool ) error
	//添加邮件消息到MQ队列
	SendMsg2Mq ( ctx context.Context, args *SendMailParams, reply *bool ) error
	//发送邮件消息,MQ消费者调用
	SendMsg ( ctx context.Context, args *SendMailParams, reply *bool ) error
}