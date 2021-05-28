package bus

//企业/商户主体相关接口定义
//@author liyang<654516092@qq.com>
//@date   2020-03-19 10:03:15
import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	// AuditActionAgree 审核通过
	AuditActionAgree = 1
	// AuditActionRefuse 审核拒绝
	AuditActionRefuse = 2
	// AuditActionUnder 审核下架
	AuditActionUnder = 3

	//资金管理方式 0=无 1=存管 2=保证保险
	FUND_MODE_NO     = 0
	FUND_MODE_CUST   = 1
	FUND_MODE_INSURE = 2

	//安全码颜色 1=灰色 2-红色 3=黄色 4=绿色 5=黑色
	SAFE_CODE_grey   = 1
	SAFE_CODE_red    = 2
	SAFE_CODE_yellow = 3
	SAFE_CODE_blue   = 4
	SAFE_CODE_black  = 5

	//保险渠道 0=无保险 1=长安保险 2=宁波人保 3=上海安信保险
	INSURANCE_CHANNEL_none = 0
	INSURANCE_CHANNEL_can  = 1
	INSURANCE_CHANNEL_ccb  = 2
	INSURANCE_CHANNEL_aaic = 3
)

//企业/商户申请/更新主体入参
type ArgsBusReg struct {
	common.Utoken          //用户信息
	CompanyName     string //企业/商户营业执照名称
	BrandName       string //企业/商户品牌名称
	IndustryId      int    //企业/商户所属领域
	MainBindId      int    //企业/商户所属主行业
	BindId          string //企业/商户所属兼营行业
	BusinessType    int    //企业/商户类型 1=中小微企业 2=个体工商户
	Pid             int    //企业/商户所属省份/直辖市
	Cid             int    //企业/商户经营所属城市
	Did             int    //企业/商户经营所属区/街道
	Address         string //企业/商户经营详细地址
	Contact         string //企业/商户联系人(负责人)
	ContactCall     string //企业/商户联系电话(手机号或固话)
	BusinessImgHash string //企业/商户营业执照图片Hash值
	CardYImgHash    string //企业/商户法人代表身份证正面照片
	CardNImgHash    string //企业/商户法人代表身份证背面照片
	AccountType     int    //企业/商户账户类型 1=对私(个人) 2=对公(企业)
	Account         string //企业/商户账户名
	AccountNo       string //企业/商户账户号
	AccountBank     int    //企业/商户账户号所属银行
	BrandImgHash    string //企业/商户logo
	OpType          int    //企业/商户主体申请操作类型 1=申请 2=更新
}

//企业/商户申请/更新主体返回信息
type ReplyBusReg struct {
	BusId int
}

//企业/商户主体详情入参-单个
type ArgsSingleBus struct {
	BusId int
}

//企业/商户主体详情入参-需要登录
type ArgsSingleBusUser struct {
	common.Utoken
	BusId int
}

//企业/商户主体返回信息
type ReplySingleBus struct {
	BusId            int                   //企业/商户ID
	MerchantId       string                //企业/商户编号
	RiskBusId        int                   //风控系统busId
	CompanyName      string                //企业/商户营业执照名称
	BrandName        string                //企业/商户品牌名称
	IndustryId       int                   //企业/商户所属领域
	MainBindId       int                   //企业/商户所属主行业
	BindId           string                //企业/商户所属兼营行业
	DistrictId       int                   //商圈id,多个使用逗号隔开
	SyntId           int                   //所属综合体id
	BusinessType     int                   //企业/商户类型 1=中小微企业 2=个体工商户
	HairpinType      int                   //企业/商户发卡类型 1=个体工商户 2=其他发卡企业 3=集团发卡企业 4=品牌发卡企业 5=规模发卡企业
	Pid              int                   //企业/商户所属省份/直辖市
	Cid              int                   //企业/商户经营所属城市
	Did              int                   //企业/商户经营所属区/街道
	Tid              int                   //区域下属镇/街道ID
	DeposBankChannel int                   //资金存管银行渠道 1=上海银行 2=交通银行
	Status           int                   //企业/商户审核状态 0=待审核 1=审核失败 2=已通过审核 3=下架
	Address          string                //企业/商户经营详细地址
	Contact          string                //企业/商户联系人(负责人)
	ContactCall      string                //企业/商户联系电话(手机号或固话)
	BusinessImg      int                   //企业/商户营业执照图片ID
	CardYImg         int                   //企业/商户法人代表身份证正面照片ID
	CardNImg         int                   //企业/商户法人代表身份证背面照片ID
	BrandImg         int                   //企业/商户品牌LOGO ID
	Ctime            int                   //企业/商户入驻时间戳
	Account          ReplySingleBusAccount //企业/商户主体银行账户
	InsuranceChannel int                   //保险渠道
	PayChannel       int                   //支付渠道
	FundMode         int                   //资金管理方式
	SafeCode         int                   //商家安全码颜色 1=黑色 2=红色 3=黄色  4=绿色
}

//企业/商户主体返回信息-主体银行账户信息
type ReplySingleBusAccount struct {
	BankAccount     string //账户
	BankAccountNo   string //账号
	BankType        int    //银行类型
	BankAccountType int    //账户类型
}

//检测用户类型入参
type ArgsBusUserType struct {
	common.Utoken //用户信息
}

//检测用户类型返回信息
type ReplyBusUserType struct {
	UserType    int    //1=普通用户 2=企业/商户总账号
	BusId       int    //企业/商户主体信息 当UserType=2时有值
	BusStatus   int    //企业/主体状态 当UserType=2时有值 0=待审核 1=审核失败 2=审核通过 3=下架
	IsindividualBusiness int //是否为个体工商户 1=是 2=否
	FundMode    int    //资金管理方式 当UserType=2时有值 1=资金存管 2=保证保险
	DeposBankChannel int  //资金存管银行
	Shops       []int  //分店ID信息
	ShopLists   []ReplySimpleShopInfo
}

//批量获取企业/商户入参
type ArgsBatchBus struct {
	BusIds []int
}

//批量获取企业/商户返回信息
type ReplyBatchBus struct {
	BusId        int    //企业/商户ID
	CompanyName  string //企业/商户营业执照名称
	BrandName    string //企业/商户品牌名称
	IndustryId   int    //企业/商户所属领域
	MainBindId   int    //企业/商户所属主行业
	BindId       string //企业/商户所属兼营行业
	BusinessType int    //企业/商户类型 1=中小微企业 2=个体工商户
	Pid          int    //企业/商户所属省份/直辖市
	Cid          int    //企业/商户经营所属城市
	Did          int    //企业/商户经营所属区/街道
	Address      string //企业/商户经营详细地址
}

//企业/商户性质返回信息
type ReplyBusinessType struct {
	BusinessTypeId   int
	BusinessTypeName string
}

//审核/修改企业/商户信息入参-平台管理员
type ArgsAuditBus struct {
	common.Utoken          //平台后台管理员信息
	common.BsToken         //企业/商户ID信息
	CompanyName     string //企业/商户营业执照名称
	BrandName       string //企业/商户品牌名称
	IndustryId      int    //企业/商户所属领域
	MainBindId      int    //企业/商户所属主行业
	BindId          string //企业/商户所属兼营行业
	BusinessType    int    //企业/商户类型 1=中小微企业 2=个体工商户
	Pid             int    //企业/商户所属省份/直辖市
	Cid             int    //企业/商户经营所属城市
	Did             int    //企业/商户经营所属区/街道
	Status          int    //企业/商户审核状态 0=待审核 1=审核失败 2=已通过审核 3=下架
	Address         string //企业/商户经营详细地址
	Contact         string //企业/商户联系人(负责人)
	ContactCall     string //企业/商户联系电话(手机号或固话)
	BusinessImgHash string //企业/商户营业执照图片HASH
	CardYImgHash    string //企业/商户法人代表身份证正面照片HASH
	CardNImgHash    string //企业/商户法人代表身份证背面照片HASH
	BranImgHash     string //企业/商户logo

}

//审核/修改企业/商户信息返回
type ReplyAuditBus struct {
	BusId int //企业/商户ID
}

//总账号替换另一个用户入参
type ArgsReplaceAccount struct {
	common.Utoken
	common.BsToken
	Phone      string
	CaptchCode string
	Channel    int // 登录渠道  0=未知， 1=pc网站 2=900岁app 3=康享宝app 4=900岁wap版，5=卡D兜小程序
	Device     int // 登录设备  0=未知  1=PC 2=Android 3=Ios 4=WinPhone  5=Mac
}

//总账号替换另一个用户返回信息
type ReplyReplaceAccount struct {
	Ok bool
}

//卡项服务获取Bus显示信息入参
type ArgsBusInfo struct {
	BusID int
}

//卡项服务获取Bus显示信息出参
type ReplyBusInfo struct {
	CompanyName string `mapstructure:"company_name"`
	BrandName   string `mapstructure:"brand_name"`
	BrandImg    int    `mapstructure:"brand_img"`
}

//检查行业id是否属于商家入参
type ArgsCheckBindid struct {
	BusId  int
	BindId int
}

// ArgsAdminBusAudit 后台商户审核入餐
type ArgsAdminBusAudit struct {
	common.Autoken     // 后台管理员信息
	BusId          int // 商户id
	Action         int // 1:通过;2:拒绝
}

// ReplyAdminBusAudit 商户审核返回信息
type ReplyAdminBusAudit struct {
	BusId int // 商户id
}

// ArgsAdminBusAuditPage 后台审核商户列表入参
type ArgsAdminBusAuditPage struct {
	common.Autoken // 后台管理员信息
	common.Paging
	CompanyName string //　企业名称
	BrandName   string //　店铺名称
	Pid         int    // 直属省份/城市
	Cid         int    // 所属城市ID
	Status      string // 商户审核状态 0=待审核 1=审核失败 2=已通过审核 3=下架
	CtimeStart  int64  //　提交开始时间戳
	CtimeEnd    int64  //　提交结束时间戳
}

// ReplyAdminBusAuditPage 后台审核商户列表返回信息
type ReplyAdminBusAuditPage struct {
	BusList  []AdminBusAuditBase
	TotalNum int // 总数
}

//AdminBusAuditBase AdminBusAuditBase
type AdminBusAuditBase struct {
	BusId               int     // 企业/商户ID
	CompanyName         string  // 企业名称
	BrandName           string  // 店铺名称
	BusinessType        int     //企业/商户类型 1=中小微企业 2=个体工商户
	HairpinType         int     //企业/商户发卡类型 1=个体工商户 2=其他发卡企业 3=集团发卡企业 4=品牌发卡企业 5=规模发卡企业
	Status              int     // 企业/商户审核状态 0=待审核 1=审核失败 2=已通过审核 3=下架
	Pid                 int     //　企业/商户所属省份/直辖市
	Cid                 int     //　企业/商户经营所属城市
	Did                 int     //　企业/商户经营所属区/街道
	Tid                 int     //区域下属镇/街道ID
	FundMode            int     //资金管理方式 0=无管理方式 1=资金存管 2=保证保险
	DeposBankChannel    int     //资金存管银行渠道 1=上海银行 2=交通银行
	Ctime               int64   // 提交时间
	InstallmentStatus   int     // 分期合作状态
	Pattern             int     // 模式
	Rate                float64 // 费率
	Limit               int     // 是否限收
	LimitAmount         float64 // 限收额度
	MarketingProportion float64 // 二次营销比例
	PayChannel          int     // 支付渠道
	InsuranceChannel    int     // 保险渠道
	YearMdateStr        string  // 转化好的日期 年-月-日
	CityName            string  // Cid转换好的城市名
}

// ArgsEsAdminBusInfo ArgsEsAdminBusInfo
type ArgsEsAdminBusInfo struct {
	BusId int
}

// ReplyEsAdminBusInfo ReplyEsAdminBusInfo
type ReplyEsAdminBusInfo struct {
	AdminBusAuditBase
}

//获取商家-服务设置选项入参
type ArgsGetBusServiceSwitch struct {
	BusId  int
	Status string //0-开启；1-关闭
}
type ReplyGetBusServiceSwitch struct {
	Id          int
	Name        string
	ServiceType int
	Status      int //0-开启；1-关闭
}

//更新商家-服务设置选项入参
type ArgsUpdateBusServiceSwitch struct {
	common.Autoken
	Id     int
	Status int
}

//分店管理页面入参数
type ArgsBranchPageMgt struct {
	common.Autoken     // 后台管理员信息
	common.Paging      //
	Pid            int // 企业/商户所属省份/直辖市
	Cid            int // 所属城市ID
	BusId          int // 商户ID
	ShopId         int
}
type ReplyBranchPageMgt struct {
	TotalNum int
	Lists    []ReplyBranchInfo
}

//分店管理页面入返回
type ReplyBranchInfo struct {
	BusId        int    //  企业/商户ID
	ShopId       int    //	分店/店铺ID
	ShopName     string //　店铺名称
	Status       int    //  企业/商户审核状态 0=待审核 1=审核失败 2=已通过审核 3=下架
	BusinessType int    //  企业/商户类型 1=中小微企业 2=个体工商户
	Pid          int    //　企业/商户所属省份/直辖市
	Cid          int    //　企业/商户经营所属城市
}

//分店/店铺详情入参
type ArgsBranchExamine struct {
	BusId          int
	ShopId         int // 店铺id
	Status         int //
	common.Autoken     // 后台管理员信息
}

//分店/店铺详情返回信息
type ReplyBranchExamine struct {
	BusId        int     //企业/商户ID
	ShopId       int     //分店/店铺ID
	ShopName     string  //店铺名称
	Address      string  //企业/商户经营详细地址
	BusinessType int     //企业/商户类型 1=中小微企业 2=个体工商户
	Contact      string  //企业/商户联系人(负责人)
	ContactCall  string  //企业/商户联系电话(手机号或固话)
	IndustryId   int     //企业/商户所属领域
	Ctime        int     //企业/商户入驻时间戳
	BusinessImg  int     //企业/商户营业执照图片ID
	BrandImg     int     //企业/商户品牌LOGO ID
	MainBindId   int     //企业/商户所属主行业
	BindId       string  //企业/商户所属兼营行业
	CardYImg     int     //企业/商户法人代表身份证正面照片ID
	CardNImg     int     //企业/商户法人代表身份证背面照片ID
	Status       int     //企业/商户审核状态 0=待审核 1=审核失败 2=已通过审核 3=下架 4=上架
	Pid          int     //企业/商户所属省份/直辖市
	Cid          int     //企业/商户经营所属城市
	Did          int     //企业/商户经营所属区/街道
	Longitude    float64 //经度
	Latitude     float64 //维度
}

// 分店/店铺 审核入参
type ArgsUpdatPage struct {
	common.Autoken        // 后台管理员信息
	ShopId         int    // 店铺id
	Status         int    // 1=同意审核 2=否决审核 3=下架 4=上架
	BusId          int    // 企业/商户ID
	DenialReason   string // 拒绝审核原因

}

//分店/店铺 审核返回信息
type ReplyUpdatPage struct {
	ShopId int // 店铺id
}

//InputParams 入参
type InputParams struct {
	common.Input
	common.BsToken
}

//InputParamsReg InputParamsReg
type InputParamsReg struct {
	RecordsID        int     `form:"recordsId" json:"recordsId"`               // recordsId		记录id
	CompanyName      string  `form:"companyName" json:"companyName"`           // companyName	公司名称
	ShopName         string  `form:"shopName" json:"shopName"`                 // shopName		企业/商户主体信息
	HaripinType      int     `form:"hairpinType" json:"hairpinType"`           // hairpinType		企业/商户发卡类型
	Pid              int     `form:"pid" json:"pid"`                           // pid			省/直辖市ID
	Cid              int     `form:"cid" json:"cid"`                           // cid			城市ID
	Did              int     `form:"did" json:"did"`                           // did			区/县ID
	Tid              int     `form:"tid" json:"tid"`                           // tid			街道ID
	DistrictId       int     `form:"districtId" json:"districtId"`             //district_id 商圈id
	SyntId           int     `form:"syntId" json:"syntId"`                     //synt_id 综合体id
	Address          string  `form:"address" json:"address"`                   // address		企业/商户实际经营地址
	BindID           int     `form:"bindId" json:"bindId"`                     // bindId		企业/商户所属行业ID
	IndustryID       int     `form:"industryId" json:"industryId"`             // industryId	企业/商户所属领域ID
	Phone            string  `form:"phone" json:"phone"`                       // phone			法人代表/负责人联系手机号码
	ApplyPhone       string  `form:"applyPhone" json:"applyPhone"`             // applyPhone	企业/商户备案申请人电话
	Buslic           string  `form:"buslic" json:"buslic"`                     // buslic		经营场所营业执照
	CardY            string  `form:"cardY" json:"cardY"`                       // cardY 		法人代表/负责人身份证正面照片
	CardN            string  `form:"cardN" json:"cardN"`                       // cardN 		法人代表/负责人身份证反面照片
	RealName         string  `form:"realName" json:"realName"`                 // realName
	DepositRatio     float32 `form:"depositRatio" json:"depositRatio"`         // depositRatio	资金存管留存比例
	BoscID           int     `form:"boscId" json:"boscId"`                     // boscId		提交存管银行ID@cs_bosc表中的bosc_id字段
	GovType          int     `form:"govType" json:"govType"`                   // govType		监管平台类型 1=宝山监管平台
	MerchantID       string  `form:"merchantId" json:"merchantId"`             // merchantId	企业/商户编号
	RiskBusID        int     `form:"riskBusId" json:"riskBusId"`               // riskBusId		企业/商户风控id
	FundManageMode   int     `form:"fundManageMode" json:"fundManageMode"`     // fundManageMode 资金管理方式 1=保证保险 2=资金存管
	DeposType        int     `form:"deposType" json:"deposType"`               // deposType		存管账户类型  0=无存管 1=上海银行 2=交通银行 3=工商银行
	DeposAccountInfo string  `form:"deposAccountInfo" json:"deposAccountInfo"` // deposAccountInfo		存管账户信息
	common.BsToken
}

//OutputReplyReg OutputReplyReg
type OutputReplyReg struct {
	//common.Output
	Sign  string //商家入住的sign
	BusId int    //商家id
}

//OutputReply 出参
type OutputReply struct {
	common.Output
}

//更改商家安全码颜色
type ArgsUpdateRiskBusSafeCode struct {
	RiskBusId     int
	Rank          int //风险状况
	SafeCodeColor int //安全码颜色值
	Pid           int //省id
	Cid           int //市id
	Did           int //区id
	Tid           int //街道id
}
type ReplyUpdateRiskBusSafeCode struct {
	BusId int
}

//监管平台直连接口-商户主体用户评论
type ArgsGetGovBusComment struct {
	RiskBusId int
}
type ReplyGetGovBusComment struct {
	CommentTotal       int     //评论量
	ComprehensiveScore float64 // 综合评分
	GoodRate           float64 //好评率
}

//Bus Bus
type Bus interface {
	//企业/商户入驻
	Reg(ctx context.Context, args *InputParamsReg, reply *OutputReplyReg) error
	//企业/商户申请/更新主体
	BusSettled(ctx context.Context, args *ArgsBusReg, reply *ReplyBusReg) error
	//获取企业/商户主体详情
	GetByBusid(ctx context.Context, args *ArgsSingleBus, reply *ReplySingleBus) error
	//批量获取企业/商户主体信息
	GetByBusids(ctx context.Context, args *ArgsBatchBus, reply *[]ReplyBatchBus) error
	//用户绑定企业/商户、门店信息
	GetBusUserType(ctx context.Context, args *ArgsBusUserType, reply *ReplyBusUserType) error
	//获取企业/商户性质
	GetBusinessType(ctx context.Context, reply *[]ReplyBusinessType) error
	//变更企业/商户总账号
	ReplaceAccount(ctx context.Context, args *ArgsReplaceAccount, reply *ReplyReplaceAccount) error
	//检查行业id是否属于商家
	CheckBindid(ctx context.Context, args *ArgsCheckBindid, reply *bool) error
	//卡项服务获取Bus显示信息
	BusInfo(ctx context.Context, args *ArgsBusInfo, reply *ReplyBusInfo) error
	// 后台商户审核
	AdminBusAudit(ctx context.Context, args *ArgsAdminBusAudit, reply *ReplyAdminBusAudit) error
	// 后台审核商户列表
	AdminBusAuditPage(ctx context.Context, args *ArgsAdminBusAuditPage, reply *ReplyAdminBusAuditPage) error
	//	后台获取商户详情(更新Es的时候调用)
	EsAdminBusInfo(ctx context.Context, args *ArgsEsAdminBusInfo, reply *ReplyEsAdminBusInfo) error
	//获取商家-服务设置选项
	GetBusServiceSwitch(ctx context.Context, args *ArgsGetBusServiceSwitch, reply *[]ReplyGetBusServiceSwitch) error
	//更新商家-服务设置选项入参
	UpdateBusServiceSwitch(ctx context.Context, args *ArgsUpdateBusServiceSwitch, reply *bool) error
	//获取分店管理页面
	GetBranchPageManagement(ctx context.Context, args *ArgsBranchPageMgt, reply *ReplyBranchPageMgt) error
	//分店/店铺详情审核
	GetBranchExamine(ctx context.Context, args *ArgsBranchExamine, reply *ReplyBranchExamine) error
	//更新分店/店铺详情
	UpadatebranchPage(ctx context.Context, args *ArgsUpdatPage, reply *ReplyUpdatPage) error
	//根据企业编号获取企业/商户主体详情
	GetByMerchantId(ctx context.Context, merchantId *string, reply *ReplySingleBus) error
	//更改商家安全码颜色
	UpdateRiskBusSafeCode(ctx context.Context, args *ArgsUpdateRiskBusSafeCode, reply *ReplyUpdateRiskBusSafeCode) error
	//监管平台直连接口-商户主体用户评论
	GetGovBusComment(ctx context.Context, args *ArgsGetGovBusComment, reply *ReplyGetGovBusComment) error
}
