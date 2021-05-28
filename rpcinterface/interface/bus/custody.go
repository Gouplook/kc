package bus

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//AccountParams 入参
type AccountParams struct {
	common.Input
}

//AccountReply 出参
type AccountReply struct {
	common.Output
}

//Custody 银行管存
type Custody interface {
	OpenAccount(ctx context.Context, args *AccountParams, reply *AccountReply) error
	SendSms(ctx context.Context, args *AccountParams, reply *AccountReply) error
	ActiveAccount(ctx context.Context, args *AccountParams, reply *AccountReply) error
	ApplyContractSign(ctx context.Context, args *AccountParams, reply *AccountReply) error
	ContractSign(ctx context.Context, args *AccountParams, reply *AccountReply) error
	OpenAcctNotify(ctx context.Context, args *AccountParams, reply *AccountReply) error
}
