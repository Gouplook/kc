package client

import (
	"context"
	"fmt"
	"git.900sui.cn/kc/base/common/plugins/jaeger"
	"git.900sui.cn/kc/kcgin"
	"github.com/smallnest/rpcx/client"
	"net/http"
	"sync"
)

var (
	rpcPools map[string]map[string]*client.XClientPool
	lock     *sync.RWMutex
)

type Baseclient struct {
	ServiceName string
	ServicePath string
	discovery   client.ServiceDiscovery
	xClient     client.XClient
}

func init() {
	rpcPools = map[string]map[string]*client.XClientPool{}
	lock = new(sync.RWMutex)
}

func (cli *Baseclient) getPools(serviceName string, servicePath string) client.XClient {
	if service, ok := rpcPools[serviceName]; ok {
		if rpcpool, ok := service[servicePath]; ok {
			return rpcpool.Get()
		} else {
			lock.Lock()
			rpcpool, ok := service[servicePath]
			if !ok {
				rpcpool = client.NewXClientPool(kcgin.AppConfig.DefaultInt("rpc_pool_count", 10), cli.ServicePath, client.Failtry, client.RandomSelect, cli.GetDiscovery(), client.DefaultOption)
				rpcPools[serviceName][servicePath] = rpcpool
			}
			lock.Unlock()
			return rpcpool.Get()
		}
	} else {
		lock.Lock()
		service, ok := rpcPools[serviceName]
		if !ok {
			service = map[string]*client.XClientPool{
				servicePath: client.NewXClientPool(kcgin.AppConfig.DefaultInt("rpc_pool_count", 10), cli.ServicePath, client.Failtry, client.RandomSelect, cli.GetDiscovery(), client.DefaultOption),
			}
			rpcPools[serviceName] = service
		}
		lock.Unlock()
		return service[servicePath].Get()
	}
}

func (cli *Baseclient) GetDiscovery() client.ServiceDiscovery {
	if cli.discovery == nil {
		address := kcgin.AppConfig.String(cli.ServiceName)
		cli.discovery = client.NewPeer2PeerDiscovery(address, "")
	}
	return cli.discovery
}

func (cli *Baseclient) getXClient() client.XClient {
	return cli.getPools(cli.ServiceName, cli.ServicePath)
}

func (cli *Baseclient) Call(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error {
	span, ctx, spanErr := jaeger.RpcxSpanWithContext(ctx, fmt.Sprintf("调用%s服务的%s方法", cli.ServicePath, serviceMethod), &http.Request{})
	if spanErr == nil {
		span.SetTag("参数", args)
		defer span.Finish()
	}

	err := cli.getXClient().Call(ctx, serviceMethod, args, reply)
	if err != nil && spanErr == nil {
		span.SetTag("error", true)
		span.SetTag("错误信息", fmt.Sprint(err))
	}
	return err
}

func (cli *Baseclient) Close() {
	//cli.getXClient().Close()
}
