/*
 * @Author: your name
 * @Date: 2021-05-19 16:29:33
 * @LastEditTime: 2021-05-19 17:19:28
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \rpcOrder\rpcinterface\interface\order\consume.go
 */
//卡项消费接口定义
//@author yangzhiwu<578154898@qq.com>
//@date 2020/8/7 10：30

package order

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
	v1 "git.900sui.cn/kc/rpcinterface/interface/order/openPlatform/v1"
)

const (
	//消费确认类型 1=短信验证 2=动态码验证
	CONFIRM_TYPE_phone   = 1
	CONFIRM_TYPE_code    = 2
	CONFIRM_TYPE_DEFAULT = 3
)

//确认消费的单项目
type ConsumeSingle struct {
	PackageRelationId int    //卡包主表id-当前项目的主卡包id
	SingleId          int    //单项目id -当前项目id，
	SingleName        string //单项目名称
	SspId             int    //单项目规格组合id
	SspName           string //单项目规格名称
	Num               int    //数量
	StaffIds          []int  //服务技师
	MarketingType     int    //营销类型： 1=身份卡，2=套餐 3=综合卡 4=限时卡 5=限次卡 6=限时限次卡 7=充值卡 8=身份卡 9=权益
	MarketingTypeId   int    //权益id，只有MarketingType=1的时候该字段才有值
	//一期优化
	ChangePrice   float64 //更改后的价格
	Discount      float64 //身份卡/充值卡折扣
	CardPackageId int     //内部使用

}

//确认消费的实物产品
type ConsumeGood struct {
	GoodsId         int    //产品id
	GoodsName       string //产品名称
	Num             int    //消费数量
	GoodsDetailId   int    //商品明细id 不同规格组合的表主键id
	GoodsDetailName string //规格名称
	StaffIds        []int  //销售技师
	//TODO 下面的字段二期可能会用到
	MarketingType     int     //营销类型： 1=身份卡，2=套餐 3=综合卡 4=限时卡 5=限次卡 6=限时限次卡 7=充值卡 8=身份卡 9=权益
	MarketingTypeId   int     //权益id，只有MarketingType=1的时候该字段才有值
	ChangePrice       float64 //更改后的价格
	Discount          float64 //身份卡/充值卡折扣
	PackageRelationId int     //卡包主表id-当前项目的主卡包id
	CardPackageId     int     //内部使用
}

//确认消费需要的参数格式
type ArgsConsumeService struct {
	common.BsToken                    //商家登录信息
	UId               int             //用户id
	ReservationId     int             //预约id
	ConfirmType       int             //消费确认类型 1=短信验证 2=动态码验证
	Captcha           string          //短语验证码或者动态码
	ConsumePrice      float64         //确认消费金额
	Singles           []ConsumeSingle //消费的单项目数据
	Goods             []ConsumeGood   //消费的实物产品
	PackageRelationId int64           //卡包主表id 可删除字段
}

//充值卡释放金额信息
type RcardRelaseAmount struct {
	RechargeLogId int     //充值记录id
	RelaseAmount  float64 //释放金额
	Amount        string  //实际消费金额
	PayOrderSn    string  //支付订单号
}

//确认消费修改卡包释放金额 入参
type ArgsChangeRelaseAmount struct {
	CardPackageRelationId  int                 //卡包关系id
	RelaseAmount           float64             //释放金额
	Amount                 string              //实际消费金额
	PayOrderSn             string              //支付订单号
	RcardRelaseAmountInfos []RcardRelaseAmount //充值卡释放存管金额信息
}

type Consume interface {
	//确认消费接口
	ConsumeService(ctx context.Context, args *ArgsConsumeService, reply *bool) error
	//确认消费修改卡包释放金额【rpc】
	ChangeRelaseAmount(ctx context.Context, args *ArgsChangeRelaseAmount, reply *bool) error
	//开放平台-v1-充值卡确认消费
	OPV1RcardConsumeSrv(ctx context.Context, args *v1.OPV1RcardConsumeSrvRequest, reply *bool) error
}
