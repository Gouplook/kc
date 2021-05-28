/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/30 下午4:14

*******************************************/
package insurance

import "context"

// 保单出单成功入参数
type ArgsCardPackPolicySuc struct {
	TransNo string // 保单流水号
	IsRenew int    // 是否续保 0=正常出单，1= 续保保单
}

type CardPackPolicy interface {
	// 保单出单成功
	CardPackPolicySuc(ctx context.Context, args *ArgsCardPackPolicySuc, reply *bool) error
}
