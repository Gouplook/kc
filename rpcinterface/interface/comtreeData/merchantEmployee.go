package comtreeData

import (
	"context"
)

type ArgsMerchantEmployee struct {
	EId int `mapstructure:"employee_id"`
}

type MerchantEmployee interface {

	// 添加 预付卡消费 信息
	AddMerchantEmployeeRpc(ctx context.Context, args *ArgsMerchantEmployee, reply *bool) error
}
