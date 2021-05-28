package comtreeData

import "context"

type ArgsMerchantShop struct {
	ShopId int
}

type MerchantShop interface {
	AddMerchantShopRpc(ctx context.Context, args *ArgsMerchantShop, reply *bool) error
}
