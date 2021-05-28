package v1

import (
	"context"
	"mime/multipart"
)

//开放接口API

const (
	//签约状态
	CONTACT_STATUS_WAIT_SIGNED    = 0 //待签约
	CONTACT_STATUS_ALREADY_SIGNED = 1 //已签约
	CONTACT_STATUS_SIGNED_FAILURE = 2 //签约失效

	//企业类型
	BUS_TYPE_GT = "1" //个体工商户
	BUS_TYPE_QY = "2" //企业

)

//开放接口公共字段，必填
type OpenRequestCommonFields struct {
	CnsmrSeqNo string `length:"22" must:"Y"` //请求流水号
	MerchId    string `length:"20" must:"Y"` //商户号，康存这边依据商户号获取对应的busId
	Version    string `length:"5" must:"Y"`  //版本号，固定 v1
	Sign       string `must:"Y"`             //签名字符串
}

//获取指定的商户签约信息-入参
type ArgsGetOneOpenMerchantContract struct {
	OpenRequestCommonFields
}

//商户签约信息
type OpenMerchantContractBase struct {
	BusId         string `mapstructure:"bus_id"`
	MerchId       string `mapstructure:"merch_id"`
	ContractId    string `mapstructure:"contract_id"`
	ContactStatus string `mapstructure:"contact_status"`
	//PrivatePem    string `mapstructure:"private_pem"`
	//PublicPem     string `mapstructure:"public_pem"`
	CreateTime string `mapstructure:"create_time"`
}

//获取指定的商户签约信息-出参
type ReplyGetOneOpenMerchantContract struct {
	OpenMerchantContractBase
}

//上传影像-入参
type ArgsUploadFile struct {
	OpenRequestCommonFields
	Type       int //上传文件分类
	Context    []byte
	FileHeader *multipart.FileHeader
}

//上传影像-出参
type ReplyUploadFile struct {
	Hash string //影像 Hash 串
	Path string //影像地址
}

//商家发布预付充值卡-入参
type ArgsMerchantCreateRcard struct {
	OpenRequestCommonFields
	RcardName     string `length:"30" must:"Y"`  //预付充值卡名称
	RcardSummary  string `length:"30"`           //预付充值卡描述
	BuyPrice      string `length:"15" must:"Y"`  //充值金额，消费者购卡付款金 额，单位：元
	GivePrice     string `length:"15" must:"Y"`  //赠送金额，单位：元
	ServicePeriod string `length:"10"`           //保险有效期，单位：月，商户信 息对接资金管理方式为保证保 险时，必传
	Tips          string `length:"500" must:"Y"` //温馨提示 json 数组字符串，格式：["提示一","提示 二"]
	ImgHash       string `length:"255" must:"Y"` //预付充值卡封面图 HASH 串
	RcardRule     string `length:"500" must:"Y"` //充值金额选项，json 数组，格 式：[{“RechargeAmount”: 1000 // 充值金额}, {“DonationAmount”: 200 // 赠送金额） }]
}

//商家发布预付充值卡-出参
type ReplyMerchantCreateRcard struct {
	RcardId string //预付充值卡 ID
}

//商家获取预付充值卡列表-入参
type ArgsMerchantRcardLists struct {
	OpenRequestCommonFields
	Page string `length:"10"` //页码,默认值为 1，每页最多返 回 10 条记录
}

//商家获取预付充值卡列表-出参
type MerchantRcardListBase struct {
	RcardId   string //预付充值卡 ID
	Name      string //预付充值卡名称
	SortDesc  string //预付充值卡短描述
	RealPrice string //售价
	Price     string //赠送金额
	Sales     string //销量
	Clicks    string //点击率
	ImgUrl    string //封面影像地址
	Ctime     string //发卡时间戳，格式：1603880239
	CtimeStr  string //发卡日期，格式：2020/10/28 18:17:19
}
type ReplyMerchantRcardLists struct {
	TotalNum string                  //总记录数
	Lists    []MerchantRcardListBase //数据
}

//商家删除预付充值卡-入参
type ArgsMerchantDelRcard struct {
	OpenRequestCommonFields
	RcardId string `length:"10" must:"Y"` //预付充值卡 ID，格式：1
}

//商家删除预付充值卡-出参
type ReplyMerchantDelRcard struct {
	RcardId string //预付充值卡 ID，格式：1
}

//商家查询账户信息及资金-入参
type ArgsMerchantDeposInfo struct {
	OpenRequestCommonFields
	BusId string
}

//商家存管账户信息
type AcctInfoBase struct {
	AcctNo       string //存管账号
	AcctName     string //存管户名
	DepositRatio string //存管比例，例如：0.4
	BankName     string //结算户开户行，例如：平安银行
	BankCardNo   string //结算账号
	BankCardName string //结算户名
	BankNo       string //结算户开户行联行号
}

//商家存管资金信息
type DepInfoBase struct {
	TotalAmount        string //总金额
	DepositoryAmount   string //存管金额
	UsableAmount       string //可提现金额
	UndischargedAmount string //待清算金额
	CashingAmount      string //提现中的金额
}

//商家查询账户信息及资金-出参
type ReplyMerchantDeposInfo struct {
	AcctInfo AcctInfoBase
	DepInfo  DepInfoBase
}

//商家查询留存资金明细-入参
type ArgsMerchantDeposLogs struct {
	OpenRequestCommonFields
	Page      string //页码,默认值为 1，每页最多返 回 10 条记录
	StartTime string //查询开始日期，格式：2021-01-20 16:40:01
	EndTime   string //查询结束日期，格式：2021-01-20 16:40:01
	BusId     string
}
type MerchantDeposLogsBase struct {
	Amount        string //金额
	Type          string //记录类型 1=入账 2=出账
	OrderType     string //账单类型 1=清算入账 2=耗卡 出账
	CreateTime    string //记录时间，格式：1603880239
	CreateTimeStr string //记录时间格式化，格式： 2021/01/06 17:09:46
	RecordTime    string //出入账时间，格式: 1603880239
	RecordTimeStr string //出入账时间格式化，格式： 2021/01/06 17:09:46
}

//商家查询留存资金明细-入参
type ReplyMerchantDeposLogs struct {
	TotalNum string //总记录数
	Lists    []MerchantDeposLogsBase
}

//商家绑定结算账户-个体-入参
type ArgsBindEntitySettleAcct struct {
	OpenRequestCommonFields
	MemberName        string `length:"30" must:"Y"` //持卡人姓名
	MemberGlobalId    string `length:"30" must:"Y"` //持卡人证件号码，身份证
	MemberAcctNo      string `length:"22" must:"Y"` //银行卡号
	BankType          string `length:"1" must:"Y"`  //银行类型 1：本行 2：它行
	EiconBankBranchId string `length:"12" must:"Y"` //结算账号开户行 305544585012
	Mobile            string `length:"11" must:"Y"` //预留手机号
	BusId             string //商户id
	BusType           string //商户类型
}

//商家绑定结算账户-企业-入参
type ArgsBindCompanySettleAcct struct {
	OpenRequestCommonFields
	MemberAcctNo           string `length:"22" must:"Y"` //银行卡号
	BankType               string `length:"1" must:"Y"`  //银行类型 1：本行 2：它行
	EiconBankBranchId      string `length:"12" must:"Y"` //结算账号开户行 305544585012
	Mobile                 string `length:"11" must:"Y"` //预留手机号
	AgencyClientFlag       string `length:"1" must:"Y"`  //是否存在经办人
	AgencyClientName       string `length:"11" must:"N"` //经办人姓名
	AgencyClientGlobalType string `length:"1" must:"N"`  //经办人证件类型
	AgencyClientGlobalId   string `length:"22" must:"N"` //经办人证件号
	AgencyClientMobile     string `length:"11" must:"N"` //经办人手机号
	ReprName               string `length:"11" must:"Y"` //法人名称
	ReprGlobalType         string `length:"2" must:"Y"`  //法人证件类型
	ReprGlobalId           string `length:"22" must:"Y"` //法人证件号码
	BusId                  string //商户id
	BusType                string //商户类型
}

//商家绑定结算账户-小额鉴权回填-入参
type ArgsSmalAmountAuthBackfill struct {
	OpenRequestCommonFields
	OrderNo string `length:"10" must:"Y"` //验证码，测试环境传 1234
	AuthAmt string `length:"10" must:"Y"` //小额鉴权金额
	BusId   string
}

//商家解绑结算账户-入参
type ArgsMerchantUnbindRelateAcct struct {
	OpenRequestCommonFields
	BusId string
}

//商家申请提现-入参
type ArgsMerchantApplyWithdraw struct {
	OpenRequestCommonFields
	Amount string `length:"10" must:"Y"` //提现金额
	BusId  string
}

//商家查询提现记录-入参
type ArgsGetMerchantWithdrawLogs struct {
	OpenRequestCommonFields
	Page   string `length:"10" must:"N"` //页码,默认值为 1，每页最多返 回 10 条记录
	Status string `length:"1" must:"N"`  //1=处理中 2=提现成功 3=提现 失败，默认获取全状态
	BusId  string
}

//提现记录信息
type GetMerchantWithdrawLogsBase struct {
	Id         string //提现 ID
	OrderSn    string //提现订单号
	CashAmount string //提现金额
	Status     string //提现状态 1=处理中 2=成功 3= 失败
	BankCardNo string //收款银行卡
	Ctime      string //申请提现时间戳 1603880239
	CtimeStr   string //申请提现时间格式化 2020/10/28 18:17:19
	Ntime      string //审核时间戳 1603880239
	NtimeStr   string //审核时间格式化 2020/10/28 18:17:19
	FailReason string //提现失败原因
}

//商家查询提现记录-出参
type ReplyGetMerchantWithdrawLogs struct {
	TotalNum string //总记录数
	Lists    []GetMerchantWithdrawLogsBase
}

//门店进件-入参
type ArgsApplyShop struct {
	OpenRequestCommonFields
	CompanyName        string `length:"155" must:"Y"` //分店营业执照名称
	ShopName           string `length:"155" must:"Y"` //分店门店名称
	BranchName         string `length:"155"`          //分店名称
	Pid                string `length:"10" must:"Y"`  //分店经营所属省份/直辖市
	Cid                string `length:"10" must:"Y"`  //分店经营所属城市
	Did                string `length:"10" must:"Y"`  //分店经营所属区/街道
	DistrictId         string `length:"100" must:"Y"` //经营区域所属商圈，最多可传 3 个商圈 ID,多商圈 ID 以逗号 隔开，例如：177,22,39
	Address            string `length:"255" must:"Y"` //分店经营详细地址
	Contact            string `length:"20" must:"Y"`  //分店联系人(负责人)
	ContactCall        string `length:"11" must:"Y"`  //分店联系电话(手机号或固话)
	BusinessImgHash    string `length:"255" must:"Y"` //分店营业执照图片Hash值
	ShopImgHash        string `length:"255" must:"Y"` //分店门店照图片Hash值
	ScanImgHash        string `length:"255"`          //分店食品卫生许可证Hash值 当领域为餐饮领域必传
	EduImgHash         string `length:"255"`          //分店教育许可证Hash值 当领域为教育领域必传
	BusinessHoursType  string `length:"1" must:"Y"`   //分店营业时间类型 1=非全天营业 2=全天营业
	BusinessHoursStart string `length:"5"`            //每天营业开始时间 BusinessHoursType=1 时必传 格式如：09:00
	BusinessHoursEnd   string `length:"5"`            //每天营业结束时间 BusinessHoursType=1 时必传 格式如：22:00
	WeekDate           string `length:"255" must:"Y"` //营业日期
	GovBusId           string `length:"10" must:"Y"`  //监管平台已信息对接的商家同步过来的商家这张表的id
}

//门店进件-出参
type ReplyApplyShop struct {
	ShopId int //分店ID
}

//门店获取信息对接唯一ID-入参
type ArgsGetGovBusId struct {
	OpenRequestCommonFields
	CompanyName string `length:"100" must:"Y"` //门店工商营业执照名称
}

//门店获取信息对接唯一ID-出参
type ReplyGetGovBusId struct {
	GovBusId string //门店信息对接唯一 ID
}

//门店添加预付充值卡-入参
type ArgsShopAddRcard struct {
	OpenRequestCommonFields
	ShopId   string `length:"10" must:"Y"`  //门店 ID
	RcardIds string `length:"100" must:"Y"` //预付充值卡 ID,多个以逗号隔开，例如：“12,13”
}

//门店获取预付充值卡列表-入参
type ArgsGetShopRcardLists struct {
	OpenRequestCommonFields
	Page   string `length:"10" must:"Y"` //页码,默认值为 1，每页最多返 回 10 条记录
	ShopId string `length:"10" must:"Y"` //门店id
	Status string `length:"1" must:"Y"`  //充值卡在门店的状态
}
type GetShopRcardListsBase struct {
	RcardId   string //预付充值卡 ID
	Name      string //预付充值卡名称
	SortDesc  string //预付充值卡短描述
	RealPrice string //售价
	Price     string //赠送金额
	Sales     string //销量
	Clicks    string //点击率
	ImgUrl    string //封面影像地址
	Ctime     string //发卡时间戳 1603880239
	CtimeStr  string //发卡日期 2020/10/28 18:17:19
}
type ReplyGetShopRcardLists struct {
	TotalNum string //总记录数
	Lists    []GetShopRcardListsBase
}

//门店上、下架预付充值卡-入参
type ArgsShopDownUpRcard struct {
	OpenRequestCommonFields
	ShopId   string `length:"10" must:"Y"`  //门店 ID
	RcardIds string `length:"100" must:"Y"` //预付充值卡 ID,多个以逗号隔开，例如：“12,13”
	FuncFlag string `length:"1" must:"Y"`   //1:下架 2：上架
}

//门店删除预付卡充值卡-入参
type ArgsShopDelRcard struct {
	OpenRequestCommonFields
	ShopId   string `length:"10" must:"Y"`  //门店 ID
	RcardIds string `length:"100" must:"Y"` //预付充值卡 ID,多个以逗号隔开，例如：“12,13”
}

//查询会员卡包信息-入参
type ArgsGetUserCardPackageLists struct {
	OpenRequestCommonFields
	Page              string `length:"10" must:"Y"` //页码,默认值为 1，每页最多返 回 10 条记录
	ShopId            string `length:"10" must:"Y"` //门店 ID
	Mobile            string `length:"11" must:"Y"` //会员手机号， 该会员手机号必须为九百岁平 台用户
	CardPackageType   string `length:"2" must:"Y"`  //卡包类型，固定传 7 7:充值卡
	CardPackageStatus string `length:"2"`           //卡包状态，不传则默认全部状 态1：待消费 2：消费中 3：已完成 4：关闭卡
}
type GetUserCardPackageListsBase struct {
	RelationId      string //卡包关联 ID
	Uid             string //用户 ID
	ShopId          string //门店 ID
	Status          string //卡包状态 1：待消费 2：消费中 3：已完成 4：已关闭
	CardPackageSn   string //卡包卡号
	CardPackageType string //卡包类型 7:充值卡
	CardPackageName string //卡包名称
	CanUse          string //是否在该门店可用 0=不可用 1=可用
	Price           string //卡包面值
	RealPrice       string //实际购卡金额
	ImgId           string //卡包封面影像 ID
	ImgUrl          string //卡包封面影像路径
	PayTime         string //支付时间戳 1586663332
	PayTimeStr      string //支付时间字符串 2020-04-12
	ConsumePrice    string //消费面值金额
	Disaccount      string //折扣率
}

//查询会员卡包信息-出参
type ReplyGetUserCardPackageLists struct {
	TotalNum string //总记录数
	Lists    []GetUserCardPackageListsBase
}

//门店兑付预付充值卡-入参
type ArgsShopConsumeRcardPackage struct {
	OpenRequestCommonFields
	ShopId       string `length:"10" must:"Y"` //门店 ID
	RelationId   string `length:"10" must:"Y"` //预付充值卡卡包 ID
	ConsumePrice string `length:"10" must:"Y"` //消费金额
}

//查询预付充值卡兑付记录-入参
type ArgsGetRardPackageConsumeLog struct {
	OpenRequestCommonFields
	RelationId string `length:"10" must:"Y"` //用户卡包 ID
	Page       string `length:"10" must:"Y"` //分页
}

//ConsumeData节点数据
type ReplyConsumeData struct {
	ConsumePrice       string //消费金额
	ActualConsumePrice string //实际消费金额
}

//ConsumeDataConf 节点数据
type ReplyConsumeDataConf struct {
	SingleId   string  //项目/产品 ID
	SingleName string  //项目/产品名称
	SspId      string  //项目/产品规格 ID ,无规格 返回 0
	SspName    string  //项目/产品规格名称、无规格 返回""
	Type       string  //类型 0=消费项目 1=消费产 品
	Num        string  //消费次数
	Price      float64 //项目/产品单价
}
type GetRardPackageConsumeLogBase struct {
	LogId           string //消费记录关联 ID
	RelationId      string //预付充值卡卡包关联 ID
	CardPackageId   string //预付充值卡 ID
	CardPackageSn   string //预付充值卡卡号
	CardPackageType string //类型：7=充值卡
	CardLogId       string //消费记录 ID
	ShopId          string //消费门店 ID
	Ctime           string //消费时间戳
	CtimeStr        string //消费时间字符串 2020/10/28 18:17:19
	ConsumeData     ReplyConsumeData
	ConsumeDataConf []ReplyConsumeDataConf
}

//查询预付充值卡兑付记录-出参
type ReplyGetRardPackageConsumeLog struct {
	TotalNum string //总记录数
	Lists    []GetRardPackageConsumeLogBase
}

type Open interface {
	//获取指定的商户签约信息
	GetOpenMerchantContract(ctx context.Context, args *string, reply *ReplyGetOneOpenMerchantContract) error
	//上传影像
	UploadFile(ctx context.Context, args *ArgsUploadFile, reply *ReplyUploadFile) error
	//商家发布预付充值卡
	MerchantCreateRcard(ctx context.Context, args *string, reply *ReplyMerchantCreateRcard) error
	//商家获取预付充值卡列表
	GetMerchantRcardLists(ctx context.Context, args *string, reply *ReplyMerchantRcardLists) error
	//商家删除预付充值卡
	MerchantDelRcard(ctx context.Context, args *string, reply *ReplyMerchantDelRcard) error
	//商家查询账户信息及资金
	GetMerchantDeposInfo(ctx context.Context, args *string, reply *ReplyMerchantDeposInfo) error
	//商家查询留存资金明细
	GetMerchantDeposLogs(ctx context.Context, args *string, reply *ReplyMerchantDeposLogs) error
	//商家绑定结算账户-个体
	BindEntitySettleAcct(ctx context.Context, args *string, reply *bool) error
	//商家绑定结算账户-企业
	BindCompanySettleAcct(ctx context.Context, args *string, reply *bool) error
	//商家绑定结算账户-小额鉴权回填
	SmalAmountAuthBackfill(ctx context.Context, args *string, reply *bool) error
	//商家解绑结算账户
	MerchantUnbindRelateAcct(ctx context.Context, args *string, reply *bool) error
	//商家申请提现
	MerchantApplyWithdraw(ctx context.Context, args *string, reply *bool) error
	//商家查询提现记录
	GetMerchantWithdrawLogs(ctx context.Context, args *string, reply *ReplyGetMerchantWithdrawLogs) error
	//门店进件
	ApplyShop(ctx context.Context, args *string, reply *ReplyApplyShop) error
	//门店获取信息对接唯一 ID
	GetGovBusId(ctx context.Context, args *string, reply *ReplyGetGovBusId) error
	//门店添加预付充值卡
	ShopAddRcard(ctx context.Context, args *string, reply *bool) error
	//门店获取预付充值卡列表
	GetShopRcardLists(ctx context.Context, args *string, reply *ReplyGetShopRcardLists) error
	//门店上、下架预付充值卡
	ShopDownUpRcard(ctx context.Context, args *string, reply *bool) error
	//门店删除预付卡充值卡
	ShopDelRcard(ctx context.Context, args *string, reply *bool) error
	//查询会员卡包信息
	GetUserCardPackageLists(ctx context.Context, args *string, reply *ReplyGetUserCardPackageLists) error
	//门店兑付预付充值卡
	ShopConsumeRcardPackage(ctx context.Context, args *string, reply *bool) error
	//查询预付充值卡兑付记录
	GetRardPackageConsumeLog(ctx context.Context, args *string, reply *ReplyGetRardPackageConsumeLog) error
}
