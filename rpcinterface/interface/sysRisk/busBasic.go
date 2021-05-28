package sysRisk

import "context"

type ArgsChangeBusArea struct {
	RiskBusId  int //风控系统商家id
	Pid        int //省id
	Cid        int //市id
	Did        int //区id
	Tid        int //街道id
	DistrictId int //街道id
	SyntId     int //综合体id
}

//风控系统 商家
type BusBasic interface {
	//修改商家的区域信息
	ChangeBusArea(ctx context.Context, args *ArgsChangeBusArea, reply *bool) error
}
