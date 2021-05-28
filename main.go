package main

import (
	"fmt"
	"git.900sui.cn/kc/base/common/components"
	"git.900sui.cn/kc/base/common/plugins/jaeger"
	"git.900sui.cn/kc/kcgin"
	"git.900sui.cn/kc/kcgin/logs"
	"git.900sui.cn/kc/rpcCards/common/plugins"
	"git.900sui.cn/kc/rpcCards/common/tools"
	"git.900sui.cn/kc/rpcCards/routers"
	"github.com/smallnest/rpcx/server"
)

//主执行函数kc/base
func main() {

	//打印环境变量
	// logs.EnableFuncCallDepth(true)
	// logs.SetLogFuncCallDepth(3)
	// logs.Info("Environment Variable:MSF_ENV:", kcgin.KcConfig.RunMode)

	//调用logger初始化
	components.InitLogger()

	//初始化项目分享链接
	tools.InitShareLink()
	//启动服务
	rpcServer := server.NewServer()
	routers.InitRpcRouters(rpcServer)
	address := fmt.Sprintf("%v:%v", kcgin.AppConfig.String("rpchost"), kcgin.AppConfig.String("rpcport"))
	logs.Info("rpcx Service address:", address)

	//启动链路追踪
	_, closer, err := jaeger.OpenJaeger()
	if err == nil && closer != nil {
		defer closer.Close()
		//添加jaeger链路追踪的一个中间件
		rpcServer.Plugins.Add(plugins.Jaeger{})
	}

	if err := rpcServer.Serve("tcp", address); err != nil {
		//rpc启动失败
		logs.Info("failed to rpcserve:%v", err)
	}
}
