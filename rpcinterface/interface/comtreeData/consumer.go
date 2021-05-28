package comtreeData

import "context"

type ArgsAddConsumer struct {
	Uid int `mapstructure:"card_id"`
}

type Consumer interface {

	// 添加 预付卡消费 信息
	AddConsumerRpc(ctx context.Context, args *ArgsAddConsumer, reply *bool) error
}
