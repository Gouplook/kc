/**
 * @Author: YangYun
 * @Date: 2020/4/13 19:13
 */
package bus

import (
    "context"
)
type Pay interface {
    // 订单支付成功
    PaySuc(ctx context.Context, orderSn *string, reply *bool) error
    //代付成功
    AgentSuc(ctx context.Context, clearId *string, reply *bool ) error
}