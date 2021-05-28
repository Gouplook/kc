package staff

import "context"

// @author liyang<654516092@qq.com>
// @date  2020/8/11 17:40

//rank基础详情
type RankModelInfo struct {
	StaffId int
	PriceTotal int
	ServiceTotal int
	EnvirTotal int
	GoodsNum int
	NormalNum int
	NegativeNum int
	TotalNum int
}

//rank详情
type RankInfo struct{
	StaffId int
	GoodsNum int
	NormalNum int
	NegativeNum int
	TotalNum int
	AvaStar float64
}

//商户服务-分店评价统计
type StaffRankBase struct {
	StaffId int
	ServiceScore int //服务评分
}

type StaffRank interface {
	//增加员工评价统计Rpc
	AddStaffRankRpc(ctx context.Context,args *StaffRankBase,reply *bool)error
}