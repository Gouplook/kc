package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"git.900sui.cn/kc/rpcinterface/interface/file"
	"git.900sui.cn/kc/rpcinterface/interface/public"
)

const (
	// 未收藏
	COLLECT_NO = 0
	// 已收藏
	COLLECT_YES = 1
)

//日期
type WeekDateParams struct {
	WeekMonStatus  int //周一状态，0：关闭，1-开启
	WeekTuesStatus int //周二状态，0：关闭，1-开启
	WeekWedStatus  int //周三状态，0：关闭，1-开启
	WeekThurStatus int //周四状态，0：关闭，1-开启
	WeekFriStatus  int //周五状态，0：关闭，1-开启
	WeekSatStatus  int //周六状态，0：关闭，1-开启
	WeekSunStatus  int //周日状态，0：关闭，1-开启
}

//分店申请入参
type ArgsBusShopReg struct {
	common.Utoken                     //用户信息
	BusId              int            //内部使用
	CompanyName        string         //分店营业执照名称
	ShopName           string         //分店门店名称
	BranchName         string         //分店名称
	BindId             string         //分店所属兼营行业
	DistrictId         string         //分店所属商圈，多个使用逗号隔开
	Pid                int            //分店经营所属省份/直辖市
	Cid                int            //分店经营所属城市
	Did                int            //分店经营所属区/街道
	Address            string         //分店经营详细地址
	Contact            string         //分店联系人(负责人)
	ContactCall        string         //分店联系电话(手机号或固话)
	BusinessImgHash    string         //分店营业执照图片Hash值
	ShopImgHash        string         //分店门店照图片Hash值
	ScanImgHash        string         //分店食品卫生许可证Hash值 当领域为餐饮领域必传
	EduImgHash         string         //分店教育许可证Hash值 当领域为教育领域必传
	BusinessHoursType  int            //分店营业时间类型 1=非全天营业 2=全天营业
	WeekDate           WeekDateParams //营业日期
	BusinessHoursStart string         //分店营业开始时间 格式如：09:00
	BusinessHoursEnd   string         //分店营业结束时间 格式如：22:00
	GovBusId           int            //监管平台已信息对接的商家同步过来的商家这张表的id
}

//审核失败入参
type ArgsBusShopRepeat struct {
	common.Utoken                     //用户信息
	ShopId        int                 //分店ID
	CompanyName   string              //分店工商名称
	GovBusId      int                 //监管平台已信息对接的商家同步过来的商家这张表的id
	Common        CommonBusShopUpdate //分店公共参数
}

//分店更新公共参数更新
type CommonBusShopUpdate struct {
	ShopName           string         //分店门店名称
	BranchName         string         //分店名称
	BindId             string         //分店所属兼营行业
	DistrictId         string         //分店所属商圈，多个使用逗号隔开
	BusinessHoursType  int            //分店营业时间类型 1=非全天营业 2=全天营业
	WeekDate           WeekDateParams //营业日期
	BusinessHoursStart string         //分店营业开始时间 格式如：09:00
	BusinessHoursEnd   string         //分店营业结束时间 格式如：22:00
	Pid                int            //分店经营所属省份/直辖市
	Cid                int            //分店经营所属城市
	Did                int            //分店经营所属区/街道
	Status             int            //分店审核状态 0=待审核 1=审核失败 2=已通过审核 3=下架
	Address            string         //分店经营详细地址
	Contact            string         //分店联系人(负责人)
	ContactCall        string         //分店联系电话(手机号或固话)
	BusinessImgHash    string         //分店营业执照图片Hash值
	ShopImgHash        string         //分店门店照图片Hash值
	ScanImgHash        string         //分店食品卫生许可证Hash值 当领域为餐饮领域必传
	EduImgHash         string         //分店教育许可证Hash值 当领域为教育领域必传
	AdvantageIds       string         //服务标签ID，多个使用逗号隔开
}

//分店信息详情
type BusShopDetail struct {
	BusinessYear  float64 //分店经营年限
	BusinessIntro string  `mapstructure:"business_introduction"` //分店简介/公告
	BusinessArea  float64 //分店经营面积 //平方米
}

//分店申请/更新返回信息
type ReplyBusShopRegUp struct {
	ShopId int //分店ID
	BusId  int //总店ID
}

//获取门店入参
type ArgsGetShop struct {
	ShopId int //分店ID
}

//获取门店入参-需要验证
type ArgsGetFrontShop struct {
	BusId  int //总店ID
	ShopId int //分店ID
}

//批量获取门店信息入参
type ArgsGetShops struct {
	ShopIds []int   //分店ID数组
	Lng     float64 //经度
	Lat     float64 //维度
}

//批量获取企业/商户下的可用分店
type ArgsAvaBusId struct {
	BusId int     //企业/商户ID
	Lng   float64 //经度
	Lat   float64 //维度
}

//批量获取审核后的门店信息入参
type ArgsAvaGetShops struct {
	ShopIds []int   //分店ID数组
	Lng     float64 //经度
	Lat     float64 //维度
}

//获取门店简易信息-切换使用
type ReplySimpleShopInfo struct {
	BusId       int    //企业/商户ID
	ShopId      int    //分店ID
	CompanyName string //分店工商营业执照名称
	ShopName    string //分店门店名称
	BranchName  string //分店名称
	Contact     string //负责人姓名
	ContactCall string //负责人联系电话
	Ctime       int64  //入驻时间时间戳
	CtimeStr    string //入驻时间字符串
}

//
type ReplyShopName struct {
	ShopId     int
	ShopName   string
	BranchName string
}

//获取门店详情返回信息
type ReplyShopInfo struct {
	ShopId             int               //分店ID
	BusId              int               //总店ID
	CompanyName        string            //分店工商营业执照名称
	ShopName           string            //分店门店名称
	BranchName         string            //分店名称
	IndustryId         int               //分店所属领域
	MainBindId         int               //分店所属主行业
	BindId             string            //分店所属兼营行业
	DistrictId         string            //分店所属商圈，多个使用逗号隔开
	SyntId             int               //分店所属综合体id
	BusinessHoursStart string            //分店营业开始时间 格式：09:00
	BusinessHoursEnd   string            //分店营业结束时间 格式：10:00
	BusinessHours      string            //分店营业时间 格式如：09:00-10:00
	BusinessHoursType  int               //营业时间范围 1=非全天 1=全天
	WeekDate           WeekDateParams    //营业日期
	BusinessImg        int               //营业执照ID
	ShopImg            int               //分店门店照ID
	SanImg             int               //食品卫生许可证ID
	EduImg             int               //教育许可证ID
	Pid                int               //分店经营所属省份/直辖市
	Cid                int               //分店经营所属城市
	Did                int               //分店经营所属区/街道
	Tid                int               //分店经营所属区下的镇/乡ID
	Status             int               //分店审核状态 0=待审核 1=审核失败 2=已通过审核 3=下架
	Address            string            //分店经营详细地址
	Contact            string            //分店联系人(负责人)
	ContactCall        string            //分店联系电话(手机号或固话)
	Longitude          float64           //经度
	Latitude           float64           //维度
	ReviewTime         int64             //审核通过时间戳
	Ctime              int64             //分店入驻时间
	FundMode           int               //资金管理方式 0=未知 1=存管 2=保险
	IsCredit           int               //是否申请过信用报告 0=否 1=是
	SafeCode           int               //商家安全码颜色 1=黑色 2=红色 3=黄色  4=绿色
	DenialReason       string            //分店拒绝审核理由
	Assoc              ReplyBusAssocInfo //卡协会员
	ShopDetail         BusShopDetail     //分店详细信息
}

//批量获取分店入参(根据企业/商户ID)
type ArgsBusId struct {
	BusId int     //企业/商户ID
	Lng   float64 //经度
	Lat   float64 //维度
}

//批量获取分店返回信息
type ReplyShopInfos struct {
	ShopId        int    //分店ID
	BusId         int    //总店ID
	CompanyName   string //分店工商营业执照名称
	ShopName      string //分店门店名称
	BranchName    string //分店名称
	Address       string //分店详细地址
	IndustryId    int    //分店所属领域
	MainBindId    int    //分店所属主行业
	IndusName     string //分店所属行业字符串
	Status        int    //分店状态 0=待审核 1=审核失败 2=审核通过 3=已下架
	Contact       string //负责人姓名
	ContactCall   string //负责人联系电话
	BindId        string //分店所属兼营行业
	Did           int
	DistrictId    string  //商圈id,多个使用逗号隔开
	SyntId        int     //分店所属综合体id
	CommentScore  float64 // 店铺综合评分
	Longitude     float64 //经度
	Latitude      float64 //维度
	BusinessHours string  //分店营业时间 格式如：09:00-10:00
	WeekDate      WeekDateParams
	Ctime         string  //入驻时间
	ShopImg       int     //分店门店照ID
	ShopImgUrl    string  //图片地址
	Distance      float64 //距离，默认单位：米
	SafeCode      int     //商家安全码颜色 1=黑色 2=红色 3=黄色  4=绿色
	DenialReason  string  //分店拒绝审核理由
}

//rpc验证门店的正确性入参
type ArgsCheckShop struct {
	BusId   int   //企业商户ID
	ShopIds []int //待验证的门店ID数组
}

//rpc返回验证门店正确性信息
type ReplyCheckShop struct {
}

type ArgsGetNearbyShopInfo struct {
	Lng    float64
	Lat    float64
	ShopId int
}

type ReplyGetNearbyShopInfo struct {
	ReplyShopInfo
	ShareLink          string
	ReguInfoDesc       string
	DistrictNames      string  //商圈名称
	DName              string  //区名称
	IndustryName       string  // 所属领域
	Distance           float64 //距离，默认单位：米
	CommentNum         int     // 评论量
	AvgPrice           float64 // 平均消费价格
	PriceScore         float64 // 技师评分
	EnvirScore         float64 // 环境评分
	ServiceScore       float64 // 服务评分
	ComprehensiveScore float64 // 店铺综合评分
	OthersShopNum      int     //其它分店数量
	IndexImgs          map[int]file.ReplyFileInfo
	BindIndus          map[int]string
	Advantages         []public.AdvantageBase
	//	todo 商家权益
}

// 其他门店
type ArgsOthersShopList struct {
	ArgsGetShops
	NowShopId int
}
type ReplyOthersShopList struct {
	TotalNum int
	Lists    []ReplyShopInfos
	//IndexImgs map[int]file.ReplyFileInfo
	//BindIndus map[int]string
}

//商户服务-分店评价统计
type BusShopReportRankBase struct {
	ShopId       int //分店ID
	BusId        int //企业/商户主体ID
	PriceScore   int //价格评分
	ServiceScore int //服务评分
	EnvirScore   int //环境评分
}
type ArgsAddBusShopReportRpc struct {
	BusShopReportRankBase
}

//根据分店ID获取分店评价数据出参
type ReplyGetShopReportRankByShopID struct {
	ShopId       int //分店ID
	BusId        int //企业/商户主体ID
	PriceTotal   int //价格评分
	ServiceTotal int //服务评分
	EnvirTotal   int //环境评分
	TotalNum     int // 总记录
}

// 获取分店人均消费数据入参
type ArgsGetShopReportCapita struct {
	ShopId int
}
type ReplyGetShopReportCapita struct {
	ShopId            int
	BusId             int
	TotalSalesPrice   float64 //总销售金额（实际支付金额）
	TotalSalesNum     int     //销量笔数
	TotalConsumePrice float64 //总消费金额（实际消费金额）
	TotalConsumeNum   int     //总消费人次
}

// 商户服务-分店人均消费统计
type ArgsAddShopReportCapita struct {
	ShopId       int
	BusId        int
	SalesPrice   float64 //销售金额（实际支付金额）
	SalesNum     int     //销量笔数
	ConsumePrice float64 //消费金额（实际消费金额）
	ConsumeNum   int     //消费人次
}

//添加门店包含的优势标签
type ArgsAddShopAdvantage struct {
	common.BsToken
	Id           int
	AdvantageIds string // 标签IDs
}

//获取门店包含的标签
type ArgsGetAdvantageIdByShopId struct {
	ShopId int
}
type GetAdvantageIdByShopId struct {
	Id           int
	AdvantageIds string
}
type ReplyGetAdvantageIdByShopId struct {
	GetAdvantageIdByShopId
	Advantages []public.AdvantageBase
}

//用户门店收藏
type ArgsUserShopCollect struct {
	common.Utoken
	ShopId int //用户收藏的ID
}

//收藏基础结构体
type CollectBase struct {
	Id        int `mapstructure:id`
	ShopId    int `mapstructure:"shopid"`
	Uid       int `mapstructure:"uid"`
	Attention int `mapstructure:"attention"`
}

//收藏门店列表 入参
type ArgsShopCollectList struct {
	common.Utoken
	common.Paging
	Lng float64 //经度
	Lat float64 //维度
}

//门店基本信息
type ShopCollectBase struct {
	ShopId             int     //分店ID
	CompanyName        string  //分店工商营业执照名称
	ShopName           string  //分店门店名称
	BranchName         string  //分店名称
	Address            string  //分店经营详细地址
	IndustryId         int     //分店所属领域
	IndusName          string  //行业类型、多个字符串分割
	DistrictNames      string  //商圈名称
	DName              string  //区名称
	MainBindId         int     //分店所属主行业
	Status             int     //分店审核状态 0=待审核 1=审核失败 2=已通过审核 3=下架
	BindId             string  //分店所属兼营行业
	Contact            string  //分店联系人(负责人)
	Longitude          float64 //经度
	Latitude           float64 //纬度
	ContactCall        string  //分店联系电话(手机号或固话)
	BusinessHours      string  //分店营业时间 格式如：09:00-10:00
	Ctime              string  //分店入驻平台时间
	ShopImg            int     //分店门店照ID
	ShopImgUrl         string  //分店门店照url
	ShopImgHash        string  //分店门店照hash
	Distance           float64 //距离，默认单位：米
	ComprehensiveScore float64 // 店铺综合评分
	SafeCode           int     //商家安全码颜色 1=黑色 2=红色 3=黄色  4=绿色
}

//店铺收藏返回信息
type ReplyShopCollectList struct {
	TotalNum int
	Lists    []ShopCollectBase
}

//店铺状态入参
type ArgsShopCollectStatus struct {
	common.Utoken
	ShopId int //用户收藏的ID
}

//店铺返回状态
type ReplyShopCollectStatus struct {
	CollectStatus int //用户收藏状态  0：表示未收藏，1：已收藏
}

//更新分店设置入参
type ArsgUpdateBusShopSetting struct {
	common.BsToken
	ShopId int
	BusShopDetail
	Common CommonBusShopUpdate //分店公共参数
}

// 门店总数 入参数
type ArgsRiskShopNum struct {
	BusId int // 商铺ID
}

// 获取风控系统商铺ID 门店总数返回值
type ReplyRiskShopNum struct {
	RiskBusId int // 风控系统商铺的ID
	ShopNum   int // 门店总数量
}

//监管平台直连接口-商户下的分店
type ArgsGetGovShopLists struct {
	Cid       int
	RiskBusId int
}
type GetGovShopListsBase struct {
	ShopId      int    //分店ID
	CompanyName string //分店工商营业执照名称
	ShopName    string //分店门店名称
	BranchName  string //分店名称
	Address     string //分店经营详细地址
	MainBindId  int    //分店所属主行业
	BindId      string //分店所属兼营行业
	IndusName   string //行业类型、多个字符串分割
	ShopImg     int
	ShopImgUrl  string //分店门店照url
}
type ReplyGetGovShopLists struct {
	TotalNum int
	Lists    []GetGovShopListsBase
}

type Shop interface {
	//分店信息申请
	ShopSettled(ctx context.Context, args *ArgsBusShopReg, reply *ReplyBusShopRegUp) error
	//分店信息更新
	ShopUpdate(ctx context.Context, args *ArgsBusShopRepeat, reply *ReplyBusShopRegUp) error
	//获取分店-根据分店ID
	GetShopByShopid(ctx context.Context, args *ArgsGetShop, reply *ReplyShopInfo) error
	//批量获取分店-根据分店IDS
	GetShopByShopids(ctx context.Context, args *ArgsGetShops, reply *[]ReplyShopInfos) error
	//批量获取分店-企业/商户ID
	GetShopByBusId(ctx context.Context, args *ArgsBusId, reply *[]ReplyShopInfos) error
	//验证分店的正确性
	CheckBusShop(ctx context.Context, args *ArgsCheckShop, reply *ReplyCheckShop) error
	//更新分店设置入参
	UpdateBusShopSetting(ctx context.Context, args *ArsgUpdateBusShopSetting, reply *bool) error
	//实现批量获取分店详细地址和封面图片-根据批量分店id
	GetShopAddressAndImgByIdS(ctx context.Context, args *[]int, reply *map[int]map[string]interface{}) error

	// 附近门店详情
	GetNearbyShopInfo(ctx context.Context, args *ArgsGetNearbyShopInfo, reply *ReplyGetNearbyShopInfo) error
	// 门店详情-其他分店
	GetOthersShopList(ctx context.Context, args *ArgsOthersShopList, reply *ReplyOthersShopList) error

	//增加分店评价统计Rpc
	AddBusShopReportRpc(ctx context.Context, args *ArgsAddBusShopReportRpc, reply *bool)
	//根据分店ID获取分店评价数据
	GetShopReportRankByShopID(ctx context.Context, shopId *int, reply *ReplyGetShopReportRankByShopID) error
	//获取分店人均消费数据
	GetShopReportCapitaByShopId(ctx context.Context, args *ArgsGetShopReportCapita, reply *ReplyGetShopReportCapita) error
	//商户服务-分店人均消费统计
	AddShopReportCapitaRpc(ctx context.Context, args *ArgsAddShopReportCapita, reply *bool) error
	//添加门店包含的优势标签
	AddShopAdvantage(ctx context.Context, args *ArgsAddShopAdvantage, reply *bool) error
	//获取门店包含的标签
	GetAdvantageIdByShopId(ctx context.Context, args *ArgsGetAdvantageIdByShopId, reply *ReplyGetAdvantageIdByShopId) error
	//添加门店收藏
	UserShopCollect(ctx context.Context, args *ArgsUserShopCollect, reply *bool) error
	//用户收藏列表
	ShopCollectList(ctx context.Context, args *ArgsShopCollectList, reply *ReplyShopCollectList) error
	//店铺收藏状态
	ShopCollectStatus(ctx context.Context, args *ArgsShopCollectStatus, reply *ReplyShopCollectStatus) error
	// 获取风控系统商铺ID 门店总数
	GetRiskShopNum(ctx context.Context, args *ArgsRiskShopNum, reply *ReplyRiskShopNum) error
	//监管平台直连接口-商户下的分店
	GetGovShopLists(ctx context.Context, args *ArgsGetGovShopLists, reply *ReplyGetGovShopLists) error
}
