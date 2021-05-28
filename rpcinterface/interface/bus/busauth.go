package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)
//统一鉴权接口定义
//@author liyang<654516092@qq.com>
//@date  2020/3/25 16:21

//九百岁saas鉴权输入参数
//BusId-ShopId 至少传一个
type ArgsBusAuth struct {
	common.Utoken  //用户信息
	BusId  int     //企业/商户ID
	ShopId int     //分店ID
	Path   string  //路径
}

//九百岁saas鉴权返回参数
type ReplyBusAuth struct{
	EncodeStr  string  //加密字符串
}

type BusAuth interface {
	 //九百岁saas统一鉴权
	 BusAuth(ctx context.Context,args *ArgsBusAuth,reply *ReplyBusAuth) error
}


