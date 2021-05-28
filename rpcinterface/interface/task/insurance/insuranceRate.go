/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/18 上午11:32

*******************************************/
package insurance

import "context"

// 投保率入场
type ArgsInsuranceRate struct {
	RelationId      int // 卡包关联ID
	CardPackageId   int // 卡包ID
	CardPackageType int // 卡包类型

}

type Insurance interface {
	// 投保率
	SetInsuranceRate(ctx context.Context, args *ArgsInsuranceRate, reply *bool) error
}
