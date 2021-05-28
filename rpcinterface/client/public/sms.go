package public

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/public"
)

//注意：此服务只能任务服务调用

type Sms struct {
	client.Baseclient
}

func (s *Sms) Init() *Sms {
	s.ServiceName = "rpc_public"
	s.ServicePath = "Sms"
	return s
}

//添加短信验证码任务到mq队列
func ( s *Sms ) SendCaptcha2Mq( ctx context.Context, args *public.SendCaptchaParams, reply *bool  ) error {
	return s.Call( ctx, "SendCaptcha2Mq", args, reply )
}

//发送短信验证码,供MQ消费者调用
func ( s *Sms ) SendCaptcha( ctx context.Context, args *public.SendCaptchaParams, reply *bool  ) error {
	return s.Call( ctx, "SendCaptcha", args, reply )
}

//验证短信验证码
func ( s *Sms ) CheckCaptcha( ctx context.Context, args *public.CheckCaptchaParams, reply *bool  ) error {
	return s.Call( ctx, "CheckCaptcha", args, reply )
}

//添加短信任务到MQ队列
func ( s *Sms ) SendSms2Mq( ctx context.Context, args *public.SendSmsParams, reply *bool  ) error {
	return s.Call( ctx, "SendSms2Mq", args, reply )
}

//发送短信,供MQ消费者调用
func ( s *Sms ) SendSms( ctx context.Context, args *public.SendSmsParams, reply *bool  ) error {
	return s.Call( ctx, "SendSms", args, reply )
}

