package order

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
)

// 定义卡包关联表数据入参出参数据
// @author liyang<654516092@qq.com>
// @date  2020/7/23 9:37

//用户卡包二维码信息
type ArgsUserCardPageQrcode struct {
	common.Utoken     //用户信息
	RelationId    int //卡包ID
}

//用户卡包列表入参
type ArgsUserList struct {
	common.Utoken       //用户信息
	common.BsToken      //企业/商户/分店ID
	common.Paging       //分页信息
	Status          int //状态
	CardPackageType int //卡包类型
}

//用户所在分店卡包列表入参
type ArgsUserBusList struct {
	common.Utoken       //用户信息
	common.BsToken      //企业/商户/分店ID
	common.Paging       //分页信息
	BusId, ShopId   int //内部使用
	Uid             int //用户uid
	Status          int //状态
	CardPackageType int //卡包类型
}

//用户卡包返回信息
type ReplyUserList struct {
	Lists    []ReplyUserListData
	TotalNum int
}

//用户卡包列表具体信息
type ReplyUserListData struct {
	RelationId       int
	BusId            int
	ShopId           int
	Uid              int
	Status           int
	CardPackageId    int
	CardPackageSn    string
	CardPackageType  int
	CardPackageName  string
	InsuranceChannel int
	Price            float64
	RealPrice        float64
	ImgId            int
	PayTime          int64
	PayTimeStr       string
	ConsumePrice     float64
	TotalNum         int
	TransferNum      int
	ExpireDate       string
	SspId            int
	SspName          string
	CardId           int
	Disaccount       float64
	ExpireSurDay     int
	ServicePeriod    int
	DiscountSingle   string //项目折扣,如果是多个折扣则min~max
	DiscountGoods    string //商品折扣
}

//用户卡包所属商家列表
type ReplyUserBusList struct {
	Lists    []ReplyUserBusListData
	TotalNum int
}

//用户卡包所属商家列表
type ReplyUserBusListData struct {
	RelationId       int
	BusId            int
	ShopId           int
	Uid              int
	Status           int
	CardPackageId    int
	CardPackageSn    string
	CardPackageType  int
	CardPackageName  string
	InsuranceChannel int
	CanUse           int
	Price            float64
	RealPrice        float64
	ImgId            int
	PayTime          int64
	PayTimeStr       string
	ConsumePrice     float64
	TotalNum         int
	TransferNum      int
	ExpireDate       string
	SspId            int
	SsId             int //单项目所在门店的id
	SspName          string
	CardId           int
	Disaccount       float64
	DiscountSingle   string //项目折扣,如果是多个折扣则min~max
	DiscountGoods    string //商品折扣
	ExpireSurDay     int
	ServicePeriod    int
}

//用户卡包详情入参
type ArgsUserCardPackageInfo struct {
	common.Utoken     //用户信息
	ShopId        int //消费门店ID【可不传】
	RelationId    int //卡包关联ID
}

//用户卡包详情-公共部分
type ReplyCommonCardPackageInfo struct {
	RelationId      int    `mapstructure:"id"` //卡包关联ID
	BusId           int    //卡包所属发卡企业/商户ID
	ShopId          int    //卡包所属销售分店ID
	Uid             int    //卡包购卡者ID
	Status          int    //卡包状态
	CardPackageId   int    //卡包ID
	CardPackageSn   string `mapstructure:"card_package_card_sn"` //卡包编号
	CardPackageType int    `mapstructure:"order_type"`           //卡包类型
}

//单项目信息
type ReplyImgHash struct {
	Hash string
	Url  string
}

//卡包-单项目详情
type ReplySingleCardPackageInfo struct {
	CardBasic  ReplyCommonCardPackageInfo
	CardInfo   CardSingleUserInfo
	CardSingle []CardSingleUserSingle
}

//卡包-套餐详情
type ReplySmCardPackageInfo struct {
	CardBasic  ReplyCommonCardPackageInfo
	CardInfo   CardSmUserInfo
	CardSingle []CardUserSingle
	SingleImgs map[int]ReplyImgHash
}

//卡包-综合卡详情
type ReplyCardCardPackageInfo struct {
	CardBasic  ReplyCommonCardPackageInfo
	CardInfo   CardCardUserInfo
	CardSingle []CardUserSingle
	CardGoods  []CardUserGood
	SingleImgs map[int]ReplyImgHash
	GoodsImgs  map[int]ReplyImgHash
}

//卡包-限时卡详情
type ReplyHcardCardPackageInfo struct {
	CardBasic  ReplyCommonCardPackageInfo
	CardInfo   CardHcardUserInfo
	CardSingle []CardUserSingle
	SingleImgs map[int]ReplyImgHash
}

//卡包-限次卡详情
type ReplyNcardCardPackageInfo struct {
	CardBasic  ReplyCommonCardPackageInfo
	CardInfo   CardNcardUserInfo
	CardSingle []CardUserSingle
	SingleImgs map[int]ReplyImgHash
}

//卡包-限时限次卡详情
type ReplyHncardCardPackageInfo struct {
	CardBasic  ReplyCommonCardPackageInfo
	CardInfo   CardHncardUserInfo
	CardSingle []CardUserSingle
	SingleImgs map[int]ReplyImgHash
}

//卡包-充值卡详情
type ReplyRcardCardPackageInfo struct {
	CardBasic  ReplyCommonCardPackageInfo
	CardInfo   CardRcardUserInfo
	CardSingle []CardUserSingle
	CardGoods  []CardUserGood
	SingleImgs map[int]ReplyImgHash
	GoodsImgs  map[int]ReplyImgHash
}

//ReplyIcardCardPackageInfo 卡包-身份卡详情
type ReplyIcardCardPackageInfo struct {
	CardBasic  ReplyCommonCardPackageInfo
	CardInfo   CardRcardUserInfo
	CardSingle []CardUserSingle
	CardGoods  []CardUserGood
	SingleImgs map[int]ReplyImgHash
	GoodsImgs  map[int]ReplyImgHash
}

//用户卡包中包含的单项目
type CardUserSingle struct {
	SingleId       int     //单项目ID
	SsId           int     //单项目在门店的id
	SspId          int     //规格id
	SpecNames      string  //规格名称组合
	SingleName     string  `mapstructure:"name"`  //单项目名称
	TotalNum       int     `mapstructure:"count"` //总次数
	TransferNum    int     //消费总次数
	ReservationNum int     //预约次数
	Discount       float64 //身份卡折扣
	SpecPrice      string  //规格价格
	Price          string  //价格范围 CanUser=0时，价格为空
	ServiceTime    int     //项目服务时长
	CanUse         int     //0=未知 1=不可用 2=可用
	ShopStatus     int     //项目在门店状态
	ShopDelStatus  int     //项目在门店是否删除
	NotCanUseDescr string  //不可用描述信息
	ImgUrl         string  //单项目封面url
	ImgHash        string  //单项目封面hash
}

//用户卡包中包含的商品
type CardUserGood struct {
	SingleId       int     `mapstructure:"goods_id"` //商品ID
	SingleName     string  `mapstructure:"name"`     //商品名称
	Discount       float64 //身份卡折扣
	Price          string
	CanUse         int    //0=未知  1=不可用 2=可用
	ShopStatus     int    //商品在门店状态
	ShopDelStatus  int    //商品在门店是否删除
	NotCanUseDescr string //不可用描述信息
	ImgUrl         string //单项目封面url
	ImgHash        string //单项目封面hash
}

//获取卡包关联信息入参-用户
type ArgsGetUserRelation struct {
	common.Utoken     //用户信息
	Id            int //关联ID
}

//获取卡包关联信息入参-rpc
type ArgsGetSimpleRelation struct {
	Id int //关联ID
}

//关联信息详情
type CardRelationInfo struct {
	Id                int    //关联ID
	SubOrderId        int    //上级订单ID
	CardPackageCardSn string //卡编号
	CardPackageId     int    //卡包ID
	Uid               int    //下单用户ID
	BusId             int    //企业/商户ID
	ShopId            int    //分店ID
	Status            int    //状态
	OrderType         int    //卡包类型
	PayTime           int64  //支付时间
	OrderSource       int    //订单来源
	Deleted           int    //删除状态
	Ctime             int64  //生成时间
}

//获取用户关联表-rpc
type ArgsGetUserCardPackageByUser struct {
	Uid             int
	CardPackageType int //卡包类型
	RelationId      int //卡包关联ID
	ShopId          int
}

type ReplyGetUserCardPackageByUser struct {
	CardBasic  ReplyCommonCardPackageInfo
	CardInfo   CardSingleUserInfo
	CardSingle []CardUserSingle
}

//卡包返回适用门店shopIds
type ReplyCardPackageCableShop struct {
	CardPackageId int    //卡包ID
	CableShop     int    //适用范围 1=全部 2=部分
	num           int    //适用门店数量
	CableShopIds  string //使用分店ID,多个以","号隔开
}

//卡包使用分店入参
type ArgsCableShopInfo struct {
	common.Utoken         //用户信息
	RelationId    int     //卡包关联ID
	Lng           float64 //经度
	Lat           float64 //维度
}

//卡包分店返回信息
type ReplyCableShopInfo struct {
	ShopId                   int     //分店ID
	CompanyName              string  //分店工商营业执照名称
	ShopName                 string  //分店门店名称
	BranchName               string  //分店名称
	Address                  string  //分店详细地址
	IndustryId               int     //分店所属领域
	IndusName                string  //行业类型、多个字符串分割
	MainBindId               int     //分店所属主行业
	Status                   int     //分店状态 0=待审核 1=审核失败 2=审核通过 3=已下架
	Contact                  string  //负责人姓名
	ContactCall              string  //负责人联系电话
	BindId                   string  //分店所属兼营行业
	Longitude                float64 //经度
	Latitude                 float64 //维度
	BusinessHours            string  //分店营业时间 格式如：09:00-10:00
	Ctime                    string  //入驻时间
	ShopImg                  int     //分店门店照ID
	ShopImgUrl               string  //分店门店照URL
	ShopImgHash              string  //分店门店照HASH
	ReservationSettingEnable int     //分店是否开通预约 0=否 1=是
	Distance                 float64 //距离，默认单位：米
}

//根据卡包关联IDs获取卡包信息-出参
type ReplyGetCardPackageListByIds struct {
	Lists []ReplyUserListData
}

//获取用户持卡数量入参-rpc
type ArgsGetUserCardPackageCountRpc struct {
	BusId  int
	ShopId int
	Uid    int
}

//获取用户持卡数量出参-rpc
type ReplyGetUserCardPackageCountRpc struct {
	AllCard int // 累计卡数量
	UseCard int // 当前可用卡数量
}

// 获取卡包下的busId入参--rpc
type ArgsGetCardPackageIdRpc struct {
	CardPackageId int // 卡包id

}
type ReplyGetCardPackageIdRpc struct {
	BusId  int //商铺Id
	ShopId int //店铺Id
}

//获取卡包详情包含的单项目/商品-入参
type ArgsGetCardPackageInfoCardSingleGoods struct {
	ArgsUserCardPackageInfo
	common.Paging
}

//获取卡包详情包含的单项目-出参
type ReplyGetCardPackageInfoCardSingle struct {
	TotalNum   int
	Lists      []CardUserSingle     //卡包包含项目
	SingleImgs map[int]ReplyImgHash //项目图片
}

//获取卡包详情包含的商品出参
type ReplyGetCardPackageInfoCardGood struct {
	TotalNum  int
	Lists     []CardUserGood       //卡包包含商品
	GoodsImgs map[int]ReplyImgHash //商品图片
}

//根据relationId 获取卡包基础信息出参 --rpc
type ReplyGetCardPackageByRelation struct {
	Id               int     //卡包id
	CardSn           string  //卡包编号RealPrice
	RechargeSn       string  //充值订单编号,只有充值卡有
	BusId            int     //商户id
	ShopId           int     //店铺id
	Uid              int     //购卡用户Id
	RealPrice        float64 //实际金额
	PayTime          int     //付款时间
	FundMode         int     //资金管理方式
	DeposBankChannel int     //资金存管银行渠道 1=上海银行 2=交通银行
	InsuranceChannel int     //保险渠道
	DepositRatio     float64 //留存比例
	DepositAmount    float64 //留存金额
	CardType         int     //卡包种类
	RelationId       int     //卡包关联id
	SubOrderId       int     //合并支付订单ID
}

//用户预付卡/预约中单子数量
type ArgsGetUserCardPackageNum struct {
	common.Utoken
}
type ReplyGetUserCardPackageNum struct {
	CanUseCardPackageNum int //可使用的预付卡数量
	ReservatingOrderNum  int //预约中单子的数量
}

//根据单项目id获取包含此单项目的卡包
type ArgsGetCardPackgeListBySingleId struct {
	SingleId int
	Uid      int
	ShopId   int
}
type ReplyGetCardPackgeListBySingleId struct {
	CardPackageSms     []ReplyUserListData //套餐卡包
	CardPackageHcards  []ReplyUserListData //限次卡包
	CardPackageNcards  []ReplyUserListData //限时卡包
	CardPackageHNcards []ReplyUserListData //限时限次卡包
	CardPackageCards   []ReplyUserListData //综合卡包
	CardPackageRcard   []ReplyUserListData //充值卡包

}

//定义接口
type CardPackageRelation interface {
	//关联表详情-用户
	GetSingleByUser(ctx context.Context, args *ArgsGetUserRelation, reply *CardRelationInfo) error
	//关联表详情-rpc
	GetSingleByRpc(ctx context.Context, args *ArgsGetSimpleRelation, reply *CardRelationInfo) error
	//获取用户卡包列表
	GetCardPackageByUser(ctx context.Context, args *ArgsUserList, reply *ReplyUserList) error
	//获取用户关联表-rpc
	GetUserCardPackageByUser(ctx context.Context, args *ArgsGetUserCardPackageByUser, reply *ReplyGetUserCardPackageByUser) error
	//根据卡包关联IDs获取卡包信息-rpc
	GetCardPackageListByIdsRpc(ctx context.Context, ids *[]int, reply *ReplyGetCardPackageListByIds) error
	//获取用户持卡数量出参-rpc
	GetUserCardPackageCountRpc(ctx context.Context, args *ArgsGetUserCardPackageCountRpc, reply *ReplyGetUserCardPackageCountRpc) error
	//根据卡包Id获取的busId--rpc
	GetCardPackageBusIdRpc(ctx context.Context, args *ArgsGetCardPackageIdRpc, reply *ReplyGetCardPackageIdRpc) error
	//获取卡包详情包含的单项目
	GetCardPackageInfoCardSingle(ctx context.Context, args *ArgsGetCardPackageInfoCardSingleGoods, reply *ReplyGetCardPackageInfoCardSingle) error
	//获取卡包详情包含的商品
	GetCardPackageInfoCardGoods(ctx context.Context, args *ArgsGetCardPackageInfoCardSingleGoods, reply *ReplyGetCardPackageInfoCardGood) error
	//根据relationId 获取卡包基础信息 --rpc
	GetCardPackageByRelation(ctx context.Context, relationId *int, reply *ReplyGetCardPackageByRelation) error
	//用户预付卡/预约中单子数量
	GetUserCardPackageNum(ctx context.Context, args *ArgsGetUserCardPackageNum, reply *ReplyGetUserCardPackageNum) error
}
