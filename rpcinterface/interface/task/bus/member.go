/**
 * @Author: YangYun
 * @Date: 2020/4/13 19:13
 */
package bus

import (
    "context"
)
type Member interface {
    // 店铺添加会员
    MemberAdd(ctx context.Context, memberId *int, reply *bool) error
}