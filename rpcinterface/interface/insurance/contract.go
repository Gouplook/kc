package insurance

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//ContractParams 合同签约入参
type ContractParams struct {
	common.Input
}

//ContractReply 合同签约出参
type ContractReply struct {
	common.Output
}

//Contract Contract
type Contract interface {
	//保险签约
	InsuranceSignUp(ctx context.Context, args *ContractParams, reply *ContractReply) error
}
