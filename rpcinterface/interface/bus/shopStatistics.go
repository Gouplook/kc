/**
 * 门店数据统计
 * @Author: yangzhiwu
 * @Date: 2020/9/2 16:32
 */

package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//获取门店营业数据入参
type ArgsShopStatisData struct {
	common.BsToken
	Day string //日期 格式：YYYYmmdd 如：20200903
}

//获取门店营业数据返回数据
type ReplyShopStatisData struct {
	CashTotalAmount float64//兑付总金额
	TotalAmount  float64 //营业额
	MemberPayNum int     //会员购买笔数
	GuestPayNum  int     //散客购买笔数
	ConsumeNum   int     //会员消费笔数
}

//获取总店/分店新增会员数入参
type ArgsMemberStatisData struct {
	common.BsToken
	DayTime       string //今日日期，格式：20060102
	YesterdayTime string //昨日日期，格式：20060102
}

//获取总店/分店新增会员数出参
type ReplyMemberStatisData struct {
	TodayTotalNum     int //今日总增长数
	YesterdayTotalNum int //昨日总增长数
	UpDown               float64 //较昨日涨跌幅度
}

type ShopStatistics interface {
	//支付成功，统计门店的经营数据
	StatisBuy(ctx context.Context, orderSn *string, reply *bool) error
	//确认消费完成，统计门店的消费数据
	StatisConsume(ctx context.Context, consumeLogId *int, reply *bool) error
	//获取门店的营业数据
	ShopStatisData(ctx context.Context, args *ArgsShopStatisData, reply *ReplyShopStatisData) error
	//获取总店/分店新增会员数
	MemberStatisData(ctx context.Context, args *ArgsMemberStatisData, reply *ReplyMemberStatisData) error
}
