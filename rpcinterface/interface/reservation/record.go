package reservation

import (
	"context"
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	product2 "git.900sui.cn/kc/rpcinterface/interface/product"
	"git.900sui.cn/kc/rpcinterface/interface/staff"
)

//预约状态
func GetStatusList() []int {
	return []int{ReservationStatusWaitConfirm, ReservationStatusConfirmed,
		ReservationStatusOrdered, ReservationStatusCompleted, ReservationStatusCanceled}
}

func CheckStatusInArray(status int) bool {
	if !functions.InArray(status, GetStatusList()) {
		return false
	}
	return true
}

//预约类型
func GetTypeList() []int {
	return []int{ReservationTypeSingle, ReservationTypeSm, ReservationTypeCard,
		ReservationTypeHcard, ReservationTypeNcard, ReservationTypeHncard}
}

func CheckTypeInArray(typ int) bool {
	if !functions.InArray(typ, GetTypeList()) {
		return false
	}
	return true
}

//预约的项目基础结构
type ItemsBase struct {
	ReservationItemId int
	StaffIDs          string  //技师id
	SalesStaffIDs     string  //服务人员id
	SingleID          int     //单项目Id
	SingleImgUrl      string  //图片
	SingleImgId       int     // 单项目图片
	SingleName        string  //单项目name
	SingleNum         int     //单项目数量
	OriginalPrice     float64 //项目原价
	Price             float64 // 项目售价
	SkuID             int     // 门店单项目规格ID
	SsId              int     // 单项目在门店的ID

	SkuName string // 规格名
	SpecIds string //子规格组合

	EstimatedStartTime    int64 // 服务开始时间()
	EstimatedStartTimeStr string
	EstimatedDuration     int //　服务持续时间
}

//预约的商品基础结构
type ProductBase struct {
	ProductItemId   int     //
	ReservationID   int     //预约记录ID
	SalesStaffIds   string  //服务人员ids
	ProductId       int     //商品id
	ProductImgUrl   string  //商品图片
	ProductNum      int     //商品数量默认为1
	OriginalPrice   float64 //商品原价
	Price           float64 // 商品价格
	ProductName     string  // 商品名
	ProductImgId    int     // 商品图片
	ProductSpecName string  // 规格名

	ProductSpecId int    // 商品规格
	SpecIds       string //子规格组合id
}

//通用预约接口入参
type CommonReserveParams struct {
	common.BsToken
	common.Utoken
	ShopId               int    // 用户端预约传入
	CallType             int    // 类型：0-预约；1-开单；默认时预约
	RType                int    //预约类型
	RUserType            int    // 预约人类型
	RCardID              int    //卡包ID
	RChannel             int    // 预约渠道
	RRemark              string //预约备注len<1024
	Name                 string //用户昵称
	Phone                string //用户手机号码
	PeopleNum            int    // 预约人数
	IsCrossDay           int    //是否跨日 0=未跨日 1=跨日
	ReservationStartTime string //预约目标到店开始时间,格式：“2012-01-02”
	StartTimePoint       string //开始时间节点，格式：“10:00”
	EndTimePoint         string //结束时间节点，格式：“13:00”
	OccupyTime           int    // 服务占用时长m
	HealthID             int    // 健康ID
	Items                []ItemsBase
	ProductsItems        []ProductBase // 商品数据
}
type ReplyCommonReserveParams struct {
	SingleItemsId  int
	ProductItemsId int
}

//通用预约接口出参
type CommonReserveReplies struct {
	ReservationID int
}

//编辑预约记录入参
type EditReservationParams struct {
	common.BsToken
	common.Utoken
	ReservationID        int
	RRemark              string //预约备注len<1024
	RType                int    //预约类型
	RCardID              int    //如果预约类型不为 single则设置此参数
	PeopleNum            int    // 预约人数
	IsCrossDay           int    //是否跨日 0=未跨日 1=跨日
	ReservationStartTime string //预约目标到店开始时间,格式：“2012-01-02”
	StartTimePoint       string //开始时间节点，格式：“10:00”
	EndTimePoint         string //结束时间节点，格式：“13:00”
	OccupyTime           int    // 服务占用时长m
	Items                []ItemsBase
	ProductsItems        []ProductBase // 商品数据
}

//状态变更入参
type ModifyReservationParams struct {
	common.BsToken //需要店面登陆
	common.Utoken
	Code           string // 预约确认验证码
	ReservationID  int
	ReservationIDs []int
	Status         int
	CancelType     int // 取消类型
	Message        string
}

//获取预约记录入参
type GetReservationRecordListParams struct {
	common.BsToken
	common.Paging
	ShopId     int
	BusId      int
	Phone      string // 手机号
	RStatus    int    //预约状态
	OrderType  int    // 订单类型：0-商品；1-服务
	SettleType int    // 结算类型 1-待耗卡；2-待支付；3-结算完成
	GetSameDay bool   // 是否获取当天的数据
	StartTime  string // 到店开始时间
	EndTime    string // 到店结束时间
	Keyword    string // 手机号和名字
}

//根据预约ID获取预约记录入参
type GetReservationRecordParams struct {
	common.BsToken
	common.Utoken
	ReservationID int
}

type ReplyGetReservationRecordParams struct {
	GetReservationRecordReplies
	Specs        []cards.Specs                  // 项目的规格
	ProductSpecs []product2.GetSpecs            // 商品规格
	StaffIndex   []staff.ReplyGetListByStaffIds // 员工数据
}

type GetReservationRecordReplies struct {
	ReservationID           int     `mapstructure:"reservation_id"`     //预约记录ID
	ReservationStatus       int     `mapstructure:"reservation_status"` //预约记录状态
	ReservationType         int     `mapstructure:"reservation_type"`   //预约记录类型
	CardID                  int     `mapstructure:"card_id"`            //卡包关联ID
	CardPackageId           int     `mapstructure:"card_package_id"`    // 卡包ID
	CardPackageName         string  `mapstructure:"card_package_name"`  // 卡包名
	ReservationRemark       string  `mapstructure:"reservation_remark"` //预约记录注释
	Name                    string  `mapstructure:"name"`               //预约记录 用户昵称
	Phone                   string  `mapstructure:"phone"`              //预约记录 用户电话
	SettleType              int     `mapstructure:"settle_type"`        //结算类型：1-待耗卡；2-待支付
	UserImg                 string  // 会员头像
	Uid                     int     // 会员ID
	MemberName              string  // 会员名
	UserType                int     `mapstructure:"reservation_user_type"` // 用户类型
	ReservationChannel      int     `mapstructure:"reservation_channel"`
	BusID                   int     `mapstructure:"bus_id"`  //预约记录 目标商户ID
	ShopID                  int     `mapstructure:"shop_id"` //预约记录 目标店面ID
	ShopName                string  //分店门店名称
	BranchName              string  //分店名称
	PeopleNum               int     `mapstructure:"people_num"`             // 预约人数
	OrderNumber             string  `mapstructure:"order_number"`           // 预约单号
	ReservationStartTime    int64   `mapstructure:"reservation_start_time"` //预约记录 预约时间,格式：“2012-01-02”
	StartTimePoint          string  `mapstructure:"start_time_point"`       //开始时间节点，格式：“10:00”
	EndTimePoint            string  `mapstructure:"end_time_point"`         //结束时间节点，格式：“13:00”
	ReservationStartTimeStr string  //预约记录 预约时间
	OccupyTime              int     `mapstructure:"occupy_time"` // 服务占用时间
	CreateTime              int64   `mapstructure:"create_time"` //预约记录创建时间
	CreateTimeStr           string  //预约记录创建时间
	OrderTime               int64   `mapstructure:"order_time"` // 开单时间
	CallType                int     `mapstructure:"call_type"`  // 订单类型
	OrderTimeStr            string  // 开单时间
	CancelTime              int64   `mapstructure:"cancel_time"` //取消时间
	CancelTimeStr           string  //取消时间
	SettlementTime          int64   `mapstructure:"settlement_time"` //取消时间
	SettlementTimeStr       string  // 结算时间
	BusRemark               string  `mapstructure:"bus_remark"`     //商家备注
	CancelMessage           string  `mapstructure:"cancel_message"` //
	TotalPrice              float64 `mapstructure:"total_price"`    //总价
	Preferential            float64 `mapstructure:"preferential"`   // 优惠
	ReachHealthId           int     //到店前健康id
	LeaveHealthId           int     //离店前健康id
	Code                    string  // 预约验证码
	ReservationItems        []ReservationItemsBase
	ProductItems            []ProductBase // 预约的商品
}

type ReservationItemsBase struct {
	ReservationID     int    //预约记录ID
	ReservationItemID int    //预约条目ID
	StaffIDs          string //技师ID
	SalesStaffIDs     string //服务人员id
	SingleId          int    //
	SingleImgUrl      string //单项目图片
	SingleNum         int    //单项目数量
	SingleImgId       int    // 单项目图片ID
	//ShopId               int
	SsId                  int // 单项目在门店的ID
	SingleName            string
	SkuID                 int     // 规格ID
	SkuName               string  // 规格名
	SpecIds               string  //子规格组合id
	OriginalPrice         float64 //项目原价
	Price                 float64 // 规格价格
	EstimatedStartTime    int64   //预计项目开始时间
	EstimatedStartTimeStr string
	EstimatedDuration     int //预计项目持续时间
}

type ReservationItemsBase2 struct {
	ReservationID         int     `mapstructure:"reservation_id"`      //预约记录ID
	ReservationItemID     int     `mapstructure:"reservation_item_id"` //预约条目ID
	StaffIDs              string  `mapstructure:"staff_ids"`           //技师ID
	SalesStaffIDs         string  `mapstructure:"sales_staff_ids"`     //服务人员id
	SingleId              int     `mapstructure:"single_id"`           //
	SingleNum             int     `mapstructure:"single_num"`          //单项目数量
	SingleImgUrl          string  //单项目图片
	SingleImgId           int     `mapstructure:"single_img_id"`
	SsId                  int     `mapstructure:"ss_id"` // 单项目在门店的ID
	SingleName            string  `mapstructure:"single_name"`
	SkuID                 int     `mapstructure:"sku_id"`   // 规格ID
	SkuName               string  `mapstructure:"sku_name"` // 规格名
	SpecIds               string  //子规格组合id
	OriginalPrice         float64 `mapstructure:"original_price"`       //项目原价
	Price                 float64 `mapstructure:"price"`                // 规格价格
	EstimatedStartTime    int64   `mapstructure:"estimated_start_time"` //预计项目开始时间
	EstimatedStartTimeStr string
	EstimatedDuration     int `mapstructure:"estimated_duration"` //预计项目持续时间
}

//获取预约记录出参
type GetReservationRecordListReplies struct {
	TotalNum     int
	Data         []GetReservationRecordReplies
	Specs        []cards.Specs                  // 项目的规格
	ProductSpecs []product2.GetSpecs            // 商品规格
	StaffIndex   []staff.ReplyGetListByStaffIds // 员工数据
}

type InvalidateReservationItemParams struct {
	common.BsToken
	ItemID int
}

type GetReservationItemListParams struct {
	common.BsToken
	StaffID int    //员工ID
	Date    string //日期
}

type Item struct {
	ReservationItemID  int     `mapstructure:"reservation_item_id"`  //预约条目ID
	StaffID            int     `mapstructure:"staff_id"`             //技师ID
	GoodsId            int     `mapstructure:"goods_id"`             //商品ID
	OriginalPrice      float64 `mapstructure:"original_price"`       //项目原价
	Price              float64 `mapstructure:"price"`                // 规格价格
	SsId               int     `mapstructure:"ss_id"`                // 单项目在门店的ID
	SkuID              int     `mapstructure:"sku_id"`               // 商品规格ID(商品明细ID)
	EstimatedStartTime string  `mapstructure:"estimated_start_time"` //预计项目开始时间
	EstimatedDuration  int     `mapstructure:"estimated_duration"`   //预计项目持续时间
	Nick               string  `mapstructure:"nick"`                 //用户昵称
	Phone              string  `mapstructure:"phone"`                //用户手机号
}
type GetReservationItemListReplies struct {
	Data []Item
}

// 门店预约记录
type ArgsReservationRecordListByShopID struct {
	common.BsToken
	//ShopId          int
	ReservationTime string // 预约日期
}

type ReservationRecordListByShopIDBase struct {
	EstimatedDuration int     `mapstructure:"estimated_duration"` //预计项目持续时间
	ReservationId     int     `mapstructure:"reservation_id"`     //预计项目持续时间
	StaffIds          string  `mapstructure:"staff_ids"`          //手艺人id
	SalesStaffIDs     string  `mapstructure:"sales_staff_ids"`    //服务人员id
	SingleId          int     `mapstructure:"single_id"`          //单项目id
	SingleNum         int     `mapstructure:"single_num"`         //单项目数量
	SingleName        string  //单项目名字
	SkuID             int     `mapstructure:"sku_id"` //子规格id
	SsId              int     `mapstructure:"ss_id"`  // 单项目在门店的ID
	SingleImgUrl      string  //单项目图片
	SingleImgId       int     // 单项目图片ID
	SpecIds           string  //子规格组合ID
	OriginalPrice     float64 `mapstructure:"original_price"` //项目原价
	Price             float64 `mapstructure:"price"`          // 规格价格
}
type ReservationRecordListByShopIDBase2 struct {
	ReservationId        int     `mapstructure:"reservation_id"`
	ReservationRemark    string  `mapstructure:"reservation_remark"` //预约记录注释
	ReservationType      int     `mapstructure:"reservation_type"`
	PeopleNum            int     `mapstructure:"people_num"`   // 预约人数
	OrderNumber          string  `mapstructure:"order_number"` // 预约单号
	SettleType           int     `mapstructure:"settle_type"`  //结算类型：1-待耗卡；2-待支付
	Name                 string  `mapstructure:"name"`
	Phone                string  `mapstructure:"phone"`
	BusID                int     `mapstructure:"bus_id"`                 //预约记录 目标商户ID
	ShopID               int     `mapstructure:"shop_id"`                //预约记录 目标店面ID
	CardID               int     `mapstructure:"card_id"`                //卡包关联ID
	CardPackageId        int     `mapstructure:"card_package_id"`        // 卡包ID
	CardPackageName      string  `mapstructure:"card_package_name"`      // 卡包名
	CallType             int     `mapstructure:"call_type"`              // 订单类型
	ReservationStatus    int     `mapstructure:"reservation_status"`     // 预约状态
	OccupyTime           int     `mapstructure:"occupy_time"`            // 服务占用时间
	ReservationEndTime   string  `mapstructure:"reservation_end_time"`   //预约记录 预约结束时间
	ReservationStartTime string  `mapstructure:"reservation_start_time"` //预约记录 预约开始时间,格式：“2012-01-02”
	StartTimePoint       string  `mapstructure:"start_time_point"`       //开始时间节点，格式：“10:00”
	EndTimePoint         string  `mapstructure:"end_time_point"`         //结束时间节点，格式：“13:00”
	CreateTime           string  `mapstructure:"create_time"`            //预约记录创建时间
	OrderTime            int64   `mapstructure:"order_time"`             // 开单时间
	CancelTime           int64   `mapstructure:"cancel_time"`            //取消时间
	OrderTimeStr         string  // 开单时间
	CancelTimeStr        string  //取消时间
	SettlementTimeStr    string  // 结算时间
	CreateTimeStr        string  //创建时间
	TotalPrice           float64 `mapstructure:"total_price"`  //总价
	Preferential         float64 `mapstructure:"preferential"` // 优惠
	ReservationItem      []ReservationRecordListByShopIDBase
	ProductItems         []ProductBase // 预约的商品
}
type ReplyReservationRecordListByShopID struct {
	TotalNum     int
	Data         *[]ReservationRecordListByShopIDBase2
	Specs        []cards.Specs                  // 项目的规格
	ProductSpecs []product2.GetSpecs            // 商品规格
	StaffIndex   []staff.ReplyGetListByStaffIds // 员工数据
}

// 根据预约ids获取预约条目入参
type ArgsGetItemsByReservationIds struct {
	ShopID         int
	ReservationIDs []int
}

//根据预约ids获取预约条目出参
type ReplyGetItemsByReservationIds struct {
	Data         map[int][]ReservationItemsBase
	ProductData  map[int][]ProductBase
	Specs        []cards.Specs                  // 项目的规格
	ProductSpecs []product2.GetSpecs            // 商品规格
	StaffIndex   []staff.ReplyGetListByStaffIds // 员工数据
}

//根据用户信息查询预约条目列表入参
type ArgsReservationByUser struct {
	Status     int
	Uid        int
	SettleType int // 结算类型 1-待耗卡；2-待支付；3-结算完成
	common.BsToken
	common.Utoken
	common.Paging
}

//根据用户信息查询预约条目列表返回
type ReplyReservationByUser struct {
	TotalNum int
	Lists    []UserReservation //
	//StaffInfos []staff.ReplyGetListByStaffIds2 // 员工数据
}

//根据用户id查询预约
type UserReservation struct {
	ReservationId     int    `mapstructure:"reservation_id"` //预约id
	OrderNumber       string `mapstructure:"order_number"`   //订单号
	PeopleNum         int    `mapstructure:"people_num"`     //预约人数
	BusId             int    `mapstructure:"bus_id"`
	SettleType        int    //结算类型
	ShopImg           string //门店照片
	ShopId            int    `mapstructure:"shop_id"` //门店id
	ShopName          string //门店名称
	BranchName        string //分店名称
	Phone             string //门店电话
	Code              string //预约验证码
	ResTime           int    `mapstructure:"reservation_start_time"` //预约时间
	ResTimeStr        string // 格式：“2012-01-02”
	StartTimePoint    string `mapstructure:"start_time_point"`   //开始时间节点，格式：“10:00”
	EndTimePoint      string `mapstructure:"end_time_point"`     //结束时间节点，格式：“13:00”
	Status            int    `mapstructure:"reservation_status"` //预约状态
	OrderTime         int    `mapstructure:"order_time"`         //开单时间
	OrderTimeStr      string
	CreateTimeStr     string        //创建时间
	ReservationRemark string        `mapstructure:"reservation_remark"`
	Services          []ServiceItem //服务数组
	Products          []ProductItem //商品数组
}
type ServiceItem struct {
	SingleId         int    `mapstructure:"single_id"`
	SingName         string `mapstructure:"single_name"`
	SkuId            int    `mapstructure:"sku_id"`
	SsId             int    `mapstructure:"ss_id"` // 单项目在门店的ID
	SkuName          string `mapstructure:"sku_name"`
	SalesStaffIds    string `mapstructure:"sales_staff_ids"`
	SalesStaffNames  string
	StaffIds         string `mapstructure:"staff_ids"`
	StaffNames       string
	OriginalPrice    float64 `mapstructure:"original_price"` //项目原价
	Price            float64 `mapstructure:"price"`
	ServiceTime      int     `mapstructure:"estimated_duration"`
	SerdverStartTime int     `mapstructure:"estimated_start_time"`
	Num              int     `mapstructure:"single_num"`
}

type ReplyGetShopReservationBase struct {
	Uid         int
	UserImage   string
	UserName    string
	Phone       string
	ResTime     int
	MemberLevel int
	MemberName  string
	//Services
	ServiceItems []ServiceItem
	//Product
	ProductItems []ProductItem
}
type ProductItem struct {
	OriginalPrice   float64 `mapstructure:"original_price"` //项目原价
	Price           float64 `mapstructure:"price"`
	Num             int     `mapstructure:"product_num"`
	SalesStaffIds   string  `mapstructure:"sales_staff_ids"`
	SalesStaffNames string
	ProductId       int    `mapstructure:"product_id"`
	ProductName     string `mapstructure:"product_name"`
	DetailId        int    `mapstructure:"product_spec_id"`
	SpecNames       string `mapstructure:"product_spec_name"`
}
type ReplyGetReplyShopReservation struct {
	TotalNum int
	Lists    []ReplyGetShopReservationBase
	//ImgIndex map[int]file.ReplyFileInfo
}

//门店预约订单管理（待收银和已完成）入参
type ArgsGetKxbOrderedList struct {
	common.BsToken
	common.Paging
	Phone      string // 手机号
	Status     int
	GetSameDay bool // 是否获取当天的数据,内部使用
	Uid        int  // 会员uid
	SettleType int  // 结算类型 1-待耗卡；2-待支付；3-结算完成
	OrderType  int  // 订单类型：1-商品；2-服务

}
type ReplyGetKxbOrderedListBase struct {
	ReservationId        int    `mapstructure:"reservation_id"`
	ReservationRemark    string `mapstructure:"reservation_remark"` //预约记录注释
	ReservationType      int    `mapstructure:"reservation_type"`
	UserType             int    `mapstructure:"reservation_user_type"` // 用户类型
	CallType             int    `mapstructure:"call_type"`             // 订单类型
	SettleType           int    `mapstructure:"settle_type"`           // 结算类型：1-待耗卡；2-待支付
	PeopleNum            int    `mapstructure:"people_num"`            // 预约人数
	OrderNumber          string `mapstructure:"order_number"`          // 预约单号
	Name                 string `mapstructure:"name"`
	Phone                string `mapstructure:"phone"`
	CardID               int    `mapstructure:"card_id"`         //卡包关联ID
	CardPackageId        int    `mapstructure:"card_package_id"` // 卡包ID
	BusID                int    `mapstructure:"bus_id"`          //预约记录 目标商户ID
	ShopID               int    `mapstructure:"shop_id"`         //预约记录 目标店面ID
	ShopName             string // 分店门店名称
	BranchName           string //分店名称
	CardPackageName      string `mapstructure:"card_package_name"` // 卡包名
	UserImg              string // 会员头像
	Uid                  int    // 会员ID
	MemberName           string // 会员名
	ReservationStatus    int    `mapstructure:"reservation_status"` // 预约状态
	OccupyTime           int    `mapstructure:"occupy_time"`        // 服务占用时间
	ReservationStartTime string //预约记录 预约开始时间// 格式：“2012-01-02”
	StartTimePoint       string `mapstructure:"start_time_point"` //开始时间节点，格式：“10:00”
	EndTimePoint         string `mapstructure:"end_time_point"`   //结束时间节点，格式：“13:00”
	OrderTimeStr         string // 开单时间
	SettlementTimeStr    string // 结算时间
	CancelTimeStr        string //取消时间
	CreateTime           int64
	CreateTimeStr        string
	TotalPrice           float64 //总价
	Preferential         float64 // 优惠
	Code                 string  // 预约验证码
	ReservationItems     []ItemsBase
	ProductItems         []ProductBase // 预约的商品
}

//门店预约订单管理（待收银和已完成）出参
type ReplyGetKxbOrderedList struct {
	TotalNum     int
	Lists        []ReplyGetKxbOrderedListBase
	Specs        []cards.Specs                  // 项目的规格
	ProductSpecs []product2.GetSpecs            // 商品规格
	StaffIndex   []staff.ReplyGetListByStaffIds // 员工数据
}

//查询会员的待服务列表
type ArgsGetMemberReservation struct {
	common.BsToken
	Uid int
}

//查询预约中用户卡包单项目的预约次数（未结束的）入参
type ArgsGetCardPackageSingleNum struct {
	Uid             int
	CardId          int   // 卡包关联id
	CardPackageType int   // 卡包类型
	SingleIds       []int // 卡包的单项目
}

//查询预约中用户卡包单项目的预约次数（未结束的）出参
type ReplyGetCardPackageSingleNum struct {
	SingleNum map[int]int
}

//预约单的验证码
type CodeBase struct {
	Code           string // 验证码
	ResId          int    // 预约单id
	ExpirationTime int64  // 过期时间
}

type ArgsGetSimpleReserOrderRpc struct {
	Status  []int //过滤的状态
	Uid     int
	OnlyRes bool  //只获取预约单
	EndTime int64 //查询结束时间
}

type ReplyGetSimpleReserOrderRpc struct {
	ResIds   []int
	TotalNum int
}

type Record interface {
	//通用预约接口
	CommonReserve(ctx context.Context, params *CommonReserveParams, replies *ReplyCommonReserveParams) (err error)
	//编辑预约
	EditReservation(ctx context.Context, params *EditReservationParams, replies *bool) (err error)
	//更新预约记录状态
	ChangeStatus(ctx context.Context, params *ModifyReservationParams, replies *bool) (err error)
	//获取预约记录信息
	GetReservationRecord(ctx context.Context, params *GetReservationRecordParams, replies *ReplyGetReservationRecordParams) (err error)
	//根据预约ids获取预约条目-rpc
	GetItemsByReservationIdsRpc(ctx context.Context, params *ArgsGetItemsByReservationIds, replies *ReplyGetItemsByReservationIds) (err error)
	//获取预约管控记录列表-总店/分店
	GetReservationControlList(ctx context.Context, params *GetReservationRecordListParams, replies *GetReservationRecordListReplies) (err error)
	//管控RPC内部
	GetReservationControlListRpc(ctx context.Context, ids *[]int, replies *GetReservationRecordListReplies) (err error)
	// 门店预约看板
	GetReservationRecordListByShopID(ctx context.Context, params *ArgsReservationRecordListByShopID, replies *ReplyReservationRecordListByShopID) (err error)
	//设置预约Item失效
	InvalidateReservationItem(ctx context.Context, params *InvalidateReservationItemParams, replies *bool) (err error)
	//获取员工预约条目列表
	GetReservationItemList(ctx context.Context, params *GetReservationItemListParams, replies *GetReservationItemListReplies) (err error)
	//获取用户预约条目列表
	GetReservationItemListByUid(ctx context.Context, args *ArgsReservationByUser, reply *ReplyReservationByUser) error
	//门店预约订单管理
	GetKxbOrderedList(ctx context.Context, args *ArgsGetKxbOrderedList, reply *ReplyGetKxbOrderedList) error
	//查询预约中用户卡包单项目的预约次数（未结束的）
	GetCardPackageSingleNum(ctx context.Context, args *ArgsGetCardPackageSingleNum, reply *ReplyGetCardPackageSingleNum) error
	//根据条件获取预约基础信息
	GetSimpleReserOrderRpc(ctx context.Context, args *ArgsGetSimpleReserOrderRpc, reply *ReplyGetSimpleReserOrderRpc) error
}
