package comtreeData

import (
	"context"
)

type ArgsMerchantSecurityCode struct {
	RiskBusId          int
	SecurityCodeStatus int
	CTime              string
	Pid                int
	Cid                int
	Did                int
	Tid                int
	IndustryId         int
	BindId             int
	Rank               int
	DistrictId int//商圈
	SyntId int //综合体
	FundMode int //资金管理方式  0=无管理方式 1=资金存管 2=保证保险
}

type MerchantSecurityCode interface {
	AddMerchantSecurityCodeRpc(ctx context.Context, args *ArgsMerchantSecurityCode, reply *bool) error
}
