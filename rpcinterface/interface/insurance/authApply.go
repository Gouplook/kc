package insurance

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

/**
 * @desc 后台审核申请投保商家
 * @className authApply
 * @author yangzhiwu<578154898@qq.com>
 * @date 2020/9/25 16:39
 */

const (
	//保险申请处理状态： 0=待处理  1=处理中 2=成功 3=失败
	STATUS_wait = 0
	STATUS_doing = 1
	STATUS_suc = 2
	STATUS_fail = 3
	//回调监管平台的状态: 0=未回调 1=回调成功 2=回调失败
	CALLBACK_STATUS_none  = 0
	CALLBACK_STATUS_suc  = 1
	CALLBACK_STATUS_fail  = 2
	//商家同意签约通知状态 0=未处理 1=同意签约成功 2=同意签约失败
	AGREE_STATUS_none = 0
	AGREE_STATUS_suc = 1
	AGREE_STATUS_fail = 2

	//保险类型 1=安信保险 2=人保
	TYPE_anx = 1
	TYPE_picc = 2

	//星级来源 0=大众点评  1=美团
	STAR_SOURCE_dianp = 0
	STAR_SOURCE_mt = 1

	INSURANCE_FLAG_fail = 0
	INSURANCE_FLAG_suc = 1
)

type ApplyInfo struct {
	Id int //申请id
	SerialNo string //流水号
	CntractId int //签约号
	CompanyName string //公司名称
	Type int //保险类型
	Status int //处理状态
	CallbackStatus int //回调状态
	AgreeStatus int //商家同意状态
	Contract string //签约内容
	Insure string //签约成功返回的信息
	Callback string //回调地址
	CreateTime int64 //创建时间
}

type ArgsGetApplyLists struct {
	Status int //审核状态
	common.Paging
}

type ReplyGetApplyLists struct {
	Lists []ApplyInfo //列表数据
	TotalNum int //总条数
}

type ArgsApplyDo struct {
	Id int //申请id
	StarSource int //星级来源 0=大众点评  1=美团
	StarNum float64 //星级数量
	RiskType int //风险等级 1=低风险 2=中风险 3=高 风险
	RiskScore float64 //风险状况分 0-100分
}


type ArgsGovAuthSucBindToBus struct {
	BusId int //商家id
	MerchantId string //企业商户编号
}

type ReplyGovAuthSucBindToBus struct {
	InsuranceType int //保险公司类型 1=安信 2=人保
	InsuranceFlag int //承保成功标价 0=未承保 1=已承保
}

type AuthApply interface {
	//获取申请列表
	GetApplyLists(ctx context.Context, args *ArgsGetApplyLists, reply *ReplyGetApplyLists) error
	//申请详情
	ApplyDetail(ctx context.Context, id *int, reply *ApplyInfo) error
	//提交投保申请
	ApplyDo(ctx context.Context, args *ArgsApplyDo , reply *bool) error
	//手动通知监管平台
	RetryNotifyToGov(ctx context.Context, id *int, reply *bool) error
	//手动同意签约意向
	RetryAgreeToAnx(ctx context.Context, id *int, reply *bool) error
	//处理承保回调业务
	AaicMerchantResultNotify(ctx context.Context, notifyData *string, reply *bool) error
	//监管平台商家同步到平台，绑定商家承保信息
	GovAuthSucBindToBus(ctx context.Context, args *ArgsGovAuthSucBindToBus, reply *ReplyGovAuthSucBindToBus ) error
}


