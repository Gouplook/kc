package public

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/public"
)

//注意：此服务只能任务服务调用

type Email struct {
	client.Baseclient
}

func (m *Email) Init() *Email {
	m.ServiceName = "rpc_public"
	m.ServicePath = "Mail"
	return m
}

//添加邮件验证码任务到mq队列
func (m *Email) SendCaptcha2Mq( ctx context.Context, args *public.SendMailCaptchaParams, reply *bool ) error {
	return m.Call( ctx, "SendCaptcha2Mq", args, reply )
}

//发送短信,MQ消费者调用
func (m *Email) SendCaptcha( ctx context.Context, args *public.SendMailCaptchaParams, reply *bool ) error {
	return m.Call( ctx, "SendCaptcha", args, reply )
}

//验证验证码
func (m *Email) CheckCaptcha( ctx context.Context, args *public.CheckMailCaptchaParams, reply *bool ) error{
	return m.Call( ctx, "CheckCaptcha", args, reply )
}

//添加邮件消息到MQ队列
func (m *Email) SendMsg2Mq ( ctx context.Context, args *public.SendMailParams, reply *bool ) error{
	return m.Call( ctx, "SendMsg2Mq", args, reply )
}

//发送普通短信,MQ消费者调用
func (m *Email) SendMsg ( ctx context.Context, args *public.SendMailParams, reply *bool ) error{
	return m.Call( ctx, "SendMsg", args, reply )
}






