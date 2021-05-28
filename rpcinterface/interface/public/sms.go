package public

import "context"

//短信相关接口
//@author yangzhiwu<wyz@900sui.com>
//@date 2020-03-09 14:55:23
//注意：此服务只能任务服务调用

type SendCaptchaParams struct {
	Phone string //手机号
	CheckSendNum bool //是否检查手机当日被发送次数
}

type SendSmsParams struct {
	Phone string //手机号
	MsgStr string  //发送的短信内容
}

//验证短信验证码参数
type CheckCaptchaParams struct {
	Phone string
	Captcha string
}

type Sms interface {
	//添加短信验证码任务到mq队列
	SendCaptcha2Mq(ctx context.Context, args *SendCaptchaParams, reply *bool ) error
	//发送短信验证码,MQ任务消费者调用
	SendCaptcha( ctx context.Context, args *SendCaptchaParams, reply *bool ) error
	//验证短信验证码
	CheckCaptcha( ctx context.Context, args *CheckCaptchaParams, reply *bool ) error
	//添加短信消息到MQ队列
	SendSms2Mq( ctx context.Context, args *SendSmsParams, reply *bool ) error
	//发送短信消息,MQ任务消费者调用
	SendSms ( ctx context.Context, args *SendSmsParams, reply *bool ) error
}