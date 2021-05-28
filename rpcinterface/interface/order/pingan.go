package order

import "context"

const (
	// 4=见证+收单的冻结资金解冻 7=在途充值解冻
	PINGAN_THAW_FLAG_jz = 4
	PINGAN_THAW_FLAG_zt = 7
)

type PingAn interface {
	//确认消费 - 平安银行解冻存管资金
	ThawDeposAmount(ctx context.Context, args *int, reply *bool) error
}
