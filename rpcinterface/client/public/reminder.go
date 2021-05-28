//温馨提示
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/3 19:16
package public

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/public"
)

type Reminder struct {
	client.Baseclient
}

//实例化
func (r *Reminder) Init() *Reminder {
	r.ServiceName = "rpc_public"
	r.ServicePath = "Reminder"
	return r
}

//获取行业的温馨提示信息
func (r *Reminder) GetReminderInfo(ctx context.Context, indusId *int, reply *[]public.ReminderInfo) error {
	return r.Call(ctx, "GetReminderInfo", indusId, reply)
}
