package comtreeData

import "context"

type ArgsMerchantSettled struct {
	BusId int
}

type MerchantSettled interface {
	AddArgsMerchantSettledRpc(ctx context.Context, args *ArgsMerchantSettled, reply *bool) error
}
