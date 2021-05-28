package bus

import (
	"context"
)

//分店面积更新风控统计
type ArgsShopAreaUpdate struct {
	ShopId       int
	BusinessArea float64
}
type Shop interface {
	//商户更改面积风控统计
	ShopAreaUpdate(ctx context.Context, args *ArgsShopAreaUpdate, reply *bool) error
}
