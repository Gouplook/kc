package gov

const (
	//# 第三方在途充值支付类型
	//微信支付
	PAY_channel_wxpay = "0001"
	//支付宝支付
	PAY_channel_alipay = "0002"

	//#证件类型
	//组织机构代码证
	TYPE_Global_zzjg = "52"
	//统一社会信用代码
	TYPE_Global_tysh = "73"
	//身份证
	TYPE_Global_sfz = "1"
	//港澳台居民通行证（即回乡证）
	TYPE_Global_gat = "3"
	//台湾居民来往大陆通行证（即台胞证）
	TYPE_Global_tw = "5"
	//外国护照
	TYPE_Global_hz = "19"

	//#银行类型
	//平安银行本行
	TYPE_BANK_pingan = "1"
	//平安银行他行
	TYPE_BANK_ta = "2"

	//是个体工商户
	INDIV_BUSINESS_yes = "1"
	//不是个体工商户
	INDIV_BUSINESS_no = "2"

	//会员名称是法人
	REP_FLAG_yes = "1"
	//会员名称不是法人
	REP_FLAG_no = "2"

	//#是否存在经办人
	//存在经办人
	HAVE_AGENCY_CLIENT_yes = "1"
	//不存在经办人
	HAVE_AGENCY_CLIENT_no = "2"
)

//公共返回参数
type ResultCommonData struct {
	CnsmrSeqNo      string //请求流水号
	ReservedMsg     string //固化字符串
	RsaSign         string //签名字符串
	TxnReturnCode   string //返回码    TxnReturnCode = 000000 代表交易成功
	TxnReturnMsg    string //返回信息
	tokenExpiryFlag string //token是否过期
}

//公共入参参数
type CommonData struct {
	ApiVersionNo   string //API 版本
	AppAccessToken string //API Token字符串
	ApplicationID  string //APPID
	RequestMode    string //请求报文格式，默认传"json"
	RsaSign        string //签名字符串名称
	SDKType        string //sdk形式 默认=api
	SdkSeid        string //sdk 版本序列号
	SdkVersionNo   string //sdk 版本号
	TranStatus     string //交易状态 固定写法 "0"
	TxnTime        string //交易时间
	ValidTerm      string //交易Term
	TxnCode        string //请求code
	TxnClientNo    string //请求客户端编号
	CnsmrSeqNo     string //请求唯一流水号
	MrchCode       string //商户号
}

//会员绑定提现账户小额鉴权参数
type BindAmountData struct {
	CommonData
	FundSummaryAcctNo      string //资金汇总账号
	SubAcctNo              string //子账户账号
	TranNetMemberCode      string //交易网会员代码
	MemberName             string //会员名称
	MemberGlobalType       string //会员证件类型
	MemberGlobalId         string //会员证件号码
	MemberAcctNo           string //会员账号
	BankType               string //银行类型
	AcctOpenBranchName     string //开户行名称
	CnapsBranchId          string //大小额行号
	EiconBankBranchId      string //超级网银行号
	Mobile                 string //手机号码
	IndivBusinessFlag      string //个体工商户标志
	CompanyName            string //公司名称
	CompanyGlobalType      string //公司证件类型
	CompanyGlobalId        string //公司证件号码
	ShopId                 string //店铺id
	ShopName               string //店铺名称
	AgencyClientFlag       string //是否存在经办人
	AgencyClientName       string //经办人姓名
	AgencyClientGlobalType string //经办人证件类型
	AgencyClientGlobalId   string //经办人证件号
	AgencyClientMobile     string //经办人手机号
	RepFlag                string //会员名称是否是法人
	ReprName               string //法人名称
	ReprGlobalType         string //法人证件类型
	ReprGlobalId           string //法人证件号码
}

//会员绑定提现账户小额鉴权返回参数
type ReplyBindAmountData struct {
	ResultCommonData
	ReservedMsg string //保留域
}

//小额鉴权回填金额入参
type CheckAmountData struct {
	CommonData
	FundSummaryAcctNo string //资金汇总账号
	SubAcctNo         string //子账户账号
	TranNetMemberCode string //交易网会员代码
	TakeCashAcctNo    string //提现账号
	AuthAmt           string //鉴权金额
	OrderNo           string //指令号
	Ccy               string //币种
	ReservedMsg       string //保留域

}

//小额鉴权回填金额返回参数
type ReplyCheckAmount struct {
	ResultCommonData
	FrontSeqNo  string //前置流水号
	ReservedMsg string //保留域
}

//解绑提现账户
type UnbindRelateAcctData struct {
	CommonData
	FunctionFlag      string //功能标志 1=解绑
	FundSummaryAcctNo string //资金汇总账号
	TranNetMemberCode string //交易网会员代码
	SubAcctNo         string //见证子账户的账号
	MemberAcctNo      string //待解绑的提现账户的账号
	ReservedMsg       string //保留域
}

//解绑提现账户返回参数
type ReplyUnbindRelateAcct struct {
	ResultCommonData
	FrontSeqNo  string //前置流水号
	ReservedMsg string //保留域
}

//查询单笔交易状态（提现）
type QueryTransactionStatus struct {
	CommonData
	FundSummaryAcctNo string //资金汇总账号
	FunctionFlag      string //功能标志 目前只使用 3=提现
	TranNetSeqNo      string //交易网流水号 对应提现请求公共参数中的CnsmrSeqNo
	SubAcctNo         string //见证子帐户的帐号 （暂未启用）
	TranDate          string //交易日期 （暂未启用）
	ReservedMsg       string //保留域

}

//查询单笔交易状态返回参数
type ReplyQueryTransactionStatus struct {
	ResultCommonData
	BookingFlag       string //记账标志 1：登记挂账 2：支付 3：提现
	TranStatus        string //交易状态 0：成功，1：失败，若系统返回状态不为0或者1， 返回其他任何状态均为交易状态不明，5分钟后重新查询。
	TranAmt           string //交易金额
	TranDate          string //交易日期
	TranTime          string //交易时间
	InSubAcctNo       string //转入子账户账号
	OutSubAcctNo      string //转出子账户账号
	FailMsg           string //失败信息 当提现失败时，返回交易失败原因
	OldTranFrontSeqNo string //原交易前置流水号
}

//会员提现
type MembershipWithdrawCash struct {
	CommonData
	TranWebName      string //交易网名称-市场名称
	SubAcctNo        string //见证子账户的账号
	MemberGlobalType string //会员证件类型 1-身份证 73-统一社会信用代码 3-港澳台居民通行证（即回乡证）
	// 5-台湾居民来往大陆通行证（即台胞证）19-外国护照 68-营业执照 52-组织机构代码证
	MemberGlobalId    string //会员证件号码
	TranNetMemberCode string //交易网会员代码
	MemberName        string //会员名称
	FundSummaryAcctNo string //资金汇总账号
	TakeCashAcctNo    string //提现账号 银行卡
	OutAmtAcctName    string //出金账户名称 银行卡户名
	Ccy               string //币种 默认为RMB
	CashAmt           string //可提现金额
	Remark            string //备注 （非必填）
	ReservedMsg       string //手续费 格式0.00（非必填）
	WebSign           string //网银签名 （非必填）
}

//会员提现返回参数
type ReplyMembershipWithdrawCash struct {
	ResultCommonData
	FrontSeqNo  string //见证系统流水号
	TransferFee string //转账手续费
	ReservedMsg string //保留域
}

//查询充值明细入参
type QueryChargeDetail struct {
	CommonData
	FundSummaryAcctNo    string //资金汇总账号
	AcquiringChannelType string //收单渠道类型01-橙E收款 02-跨行快收（非T0） 03-跨行快收（T0） 04-聚合支付 YST1-云收款
	OrderNo              string //订单号 下单时的子订单号，不是总订单号 详见《电商见证宝开发说明V1.0.docx》的2.4章节
	ReservedMsg          string //保留域
}

//查询充值明细返回参数
type ReplyQueryChargeDetail struct {
	ResultCommonData
	TranStatus            string //交易状态 0：成功，1：失败，2：异常,3:冲正，5：待处理
	TranAmt               string //交易金额
	CommissionAmt         string //佣金费
	PayMode               string //支付方式 0-冻结支付 1-普通支付
	TranDate              string //交易日期
	TranTime              string //交易时间
	OrderInSubAcctNo      string //订单转入见证子账户的帐号
	OrderInSubAcctName    string //订单转入见证子账户的名称
	OrderActInSubAcctNo   string //订单实际转入见证子账户的帐号
	OrderActInSubAcctName string //订单实际转入见证子账户的名称
	FrontSeqNo            string //见证系统流水号
	TranDesc              string //交易描述 当充值失败时，返回交易失败原因
}

//会员资金冻结入参
type MembershipTrancheFreeze struct {
	CommonData
	FunctionFlag string //功能标志 1：冻结（会员→担保） 2：解冻（担保→会员）4：见证+收单的冻结资金解冻 5: 可提现冻结（会员→担保）
	// 6: 可提现解冻（担保→会员） 7: 在途充值解冻（担保→会员）
	FundSummaryAcctNo string //资金汇总账号
	SubAcctNo         string //见证子账户的账号
	TranNetMemberCode string //交易网会员代码
	ConsumeAmt        string //消费金额
	TranAmt           string //交易金额-释放金额
	TranCommission    string //交易手续费
	Ccy               string //币种
	OrderNo           string //订单号
	OrderContent      string //订单内容
	Remark            string //备注
	ReservedMsg       string //保留域
	Cid               int    `json:"-"` //商家所在城市id
}

//会员资金冻结出参
type ReplyMembershipTrancheFreeze struct {
	ResultCommonData
	FrontSeqNo  string //见证系统流水号
	ReservedMsg string //保留域
}

//查询小额鉴权转账结果
type SmallAmountTransferQuery struct {
	CommonData
	FundSummaryAcctNo string //资金汇总账号
	OldTranSeqNo      string //原交易流水号
	TranDate          string //交易日期
}

//查询小额鉴权转账结果返回参数
type ReplySmallAmountTransferQuery struct {
	ResultCommonData
	ReturnStatu string //返回状态
	ReturnMsg   string //返回信息
	ReservedMsg string //保留域
}

//登记挂账（增加可提现金额）
type RegisterBillSupportWithdraw struct {
	CommonData
	FundSummaryAcctNo string //资金汇总账号
	SubAcctNo         string //见证子账户的账号
	TranNetMemberCode string //交易网会员代码
	OrderNo           string //订单号
	SuspendAmt        string //挂账金额
	TranFee           string //交易费用
	Remark            string //备注
}

//登记挂账返回参数
type ReplyRegisterBillSupportWithdraw struct {
	ResultCommonData
	FrontSeqNo  string //见证系统流水号
	ReservedMsg string //保留域
}

//查询企业/商户子账户资金余额入参
type QueryAccountBalance struct {
	CommonData
	FundSummaryAcctNo string //资金归集账号
	TranNetMemberCode string //企业/商户唯一ID
}

//查询企业/商户子账户资金余额返回参数
type ReplyQueryAccountBalance struct {
	ResultCommonData
	SubAcctNo        string //子账户号
	SubAcctCashBal   string //子账户可提现资金
	SubAcctAvailBal  string //子账户可用资金
	SubAcctFreezeAmt string //子账户冻结资金
}

//第三方支付渠道在途充值（分账）入参
type RechargeWay struct {
	BusId int //企业/商户ID
	Data  RechargeWayThirdPay
}

//第三方支付渠道在途充值数据
type RechargeWayThirdPay struct {
	CommonData
	FundSummaryAcctNo     string                     //资金归集账号
	PayChannelType        string                     //支付渠道类型 0001=微信 0002=支付宝
	PayChannelAssignMerNo string                     //支付渠道所分配的商户号
	TotalOrderNo          string                     //总订单号
	TranTotalAmt          string                     //交易总金额
	OrdersCount           string                     //订单数量
	TranItemArray         []RechargeWayTranItemArray //子订单信息
}

//第三方支付渠道在途充值入参TranItemArray信息
type RechargeWayTranItemArray struct {
	RechargeSubAcctNo    string //充值子账户
	SubOrderFillMemberCd string //子订单充值会员代码
	SubOrderTranAmt      string //子订单交易金额
	SubOrderTranFee      string //子订单交易费用
	SubOrderNo           string //子订单号
	SubOrderPayMode      string //支付模式 0-冻结支付 1=普通支付
	SubOrderContent      string //订单描述
}

//第三方支付渠道在途充值返回结果
type ReplyRechargeWay struct {
	ResultCommonData
	OrdersCount   string           //子订单数量
	TranItemArray []TransItemArray //前置流水号数组
}

//第三方支付渠道在途充值前置流水号数组
type TransItemArray struct {
	FrontSeqNo string //前置流水号
}

//撤销第三方支付渠道在途充值（分账）入参
type UndoRechargeWay struct {
	CommonData
	FundSummaryAcctNo string                         //资金归集账号
	OldPayChannelType string                         //支付渠道类型
	OldTotalOrderNo   string                         //原总订单号
	TotalRefundAmt    string                         //退款总金额
	RefundOrderNum    string                         //退款总订单数量
	TranItemArray     []UndoRechargeWayTranItemArray //子订单信息
}

//撤销第三方支付渠道在途充值（分账）入参TranItemArray信息
type UndoRechargeWayTranItemArray struct {
	SubOrderRefundSubAcctNo string //子订单退款子账户
	SubOrderRefundMemberCd  string //子订单退款会员代码
	SubOrderMemberRefundAmt string //子订单退款金额
	SubOrderFeeRefundAmt    string //子订单手续费退款金额
	SubOrderRefundOrderNo   string //子订单退款订单号
}

//撤销第三方支付渠道在途充值返回结果
type ReplyUndoRechargeWay struct {
	ResultCommonData
	OrdersCount   string               //子订单数量
	TranItemArray []UndoTransItemArray //前置流水号数组
}

//撤销第三方支付渠道在途充值前置流水号数组
type UndoTransItemArray struct {
	FrontSeqNo string //前置流水号
}

//会员绑定信息查询
type MemberBindQuery struct {
	CommonData
	QueryFlag         string //查询标志 1：全部会员 2：单个会员
	FundSummaryAcctNo string //资金汇总账号
	SubAcctNo         string //见证子账户的账号 若QueryFlag为2时， 子账户账号必输
	PageNum           string //起始值为1，每次最多返回20条记录，第二页返回的记录数为第21至40条记录，第三页为41至60条记录，顺序均按照建立时间的先后
}


//预付卡充值查询请求参数
type QueryPrepaidCardRecharge struct {
	CommonData
	QueryFlag         		string //查询标志 1:见证+收单订单（云收款接口发起的） 2:三方充值订单（6251接口发起的）
	AcquiringChannelType 	string //收单渠道类型 当FuncFlag为1时，必输； YST1-云收款T1
	OrderNo         		string //订单号 根据所填渠道所返回的订单号
}

//预付卡充值查询返回参数
type ReplyPrepaidCardRecharge struct {
	ResultCommonData
	TranAmt 			string	//交易金额
	Commission 			string	//手续费
	PayMode 			string	//支付方式
	TranDate 			string	//交易日期
	TranTime 			string	//交易时间
	OrderRemainAmt 		string	//订单剩余金额
	OrderInSubAcctNo 	string	//订单转入子账户
	DepositRatio 		string	//缴存比例

}