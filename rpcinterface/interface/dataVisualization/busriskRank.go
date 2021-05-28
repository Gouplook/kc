/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/1 16:23
@Description:

*********************************************/
package dataVisualization

import "context"

//更改商家安全码颜色
type ArgsBusRisRandSafeCode struct {
	RiskBusId     int //企业/商户ID
	Rank          int //风险状况
	SafeCodeColor int //安全码颜色值
	Pid           int //省id
	Cid           int //市id
	Did           int //区id
	Tid           int //街道id
	DistrictId	  int //商圈id
	SyntId 		  int //综合体id
}

//修改商家的区域信息
type ArgsChangeBusArea struct {
	RiskBusId     int //企业/商户ID
	Pid           int //省id
	Cid           int //市id
	Did           int //区id
	Tid           int //街道id
	DistrictId	  int //商圈id
	SyntId 		  int //综合体id
}

type BusRiskRank interface {
	//更改商户风险等级和安全码信息
	UpdateBusRisRandSafeCode(ctx context.Context, args *ArgsBusRisRandSafeCode, reply *bool) error
	//修改商家的区域信息
	ChangeBusArea(ctx context.Context, aegs *ArgsChangeBusArea, reply *bool) error
}
