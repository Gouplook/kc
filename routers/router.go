package routers

import (
	"log"

	"git.900sui.cn/kc/rpcCards/service"
	"github.com/smallnest/rpcx/server"
)

//InitRpcRouters 初始化rpc路由
func InitRpcRouters(rpcServer *server.Server) {
	//注册标签服务
	if err := rpcServer.Register(new(service.Tag), ""); err != nil {
		log.Fatalf("failed to register rpcRouter: %v", err)
	}
	//注册规格服务
	if err := rpcServer.Register(new(service.Spec), ""); err != nil {
		log.Fatalf("failed to register rpcRouter: %v", err)
	}
	//注册单项目服务
	if err := rpcServer.Register(new(service.Single), ""); err != nil {
		log.Fatalf("failed to register rpcRouter: %v", err)
	}
	//注册套餐服务
	if err := rpcServer.Register(new(service.Sm), ""); err != nil {
		log.Fatalf("failed to register rpcRouter: %v", err)
	}
	//注册综合卡服务
	if err := rpcServer.Register(new(service.Card).Init(), ""); err != nil {
		log.Fatalf("failed to register rpcRouter: %v", err)
	}
	//注册限次卡服务
	if err := rpcServer.Register(new(service.NCard).Init(), ""); err != nil {
		log.Fatalf("failed to register rpcRouter: %v", err)

	}
	//注册限时限次卡服务
	if err := rpcServer.Register(new(service.HNCard).Init(), ""); err != nil {
		log.Fatalf("failed to register rpcRouter: %v", err)
	}
	//注册CardExt描述选项获取服务
	if err := rpcServer.Register(new(service.Note).Init(), ""); err != nil {
		log.Fatalf("failed to register rpcRouter: %v", err)
	}
	// 注册限时卡服务
	if err := rpcServer.Register(new(service.Hcard), ""); err != nil {
		log.Fatalf("failed to register rpcRouter: %v", err)
	}
	//注册项目通用服务
	if err := rpcServer.Register(new(service.Item), ""); err != nil {
		log.Fatalf("failed to register rpcRouter: %v", err)
	}
	//注册充值卡服务
	if err := rpcServer.Register(new(service.Rcard), ""); err != nil {
		log.Fatalf("failed to register rpcRouter: %v", err)
	}
	//折扣卡管理（身份卡）
	if err := rpcServer.Register(new(service.ICard), ""); err != nil {
		log.Fatalf("failed to register rpcRouter: %v", err)
	}
}
