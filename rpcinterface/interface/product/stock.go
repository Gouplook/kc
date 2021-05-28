package product

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//库存查询
type ArgsStockGet struct {
	BusId         int //商户信息
	common.Paging     //分页信息
	ShopId        int //门店id
	CateId        int //分类id
}

//新增入库 入参
type ArgsStockInAdd struct {
	common.BsToken               //商户信息
	common.Utoken
	Time           string        //入库时间
	Remark         string        //备注
	Details        []StockDetail //商品list
}

//出入库 商品明细 详情
type StockDetail struct {
	ProductId int     //商品id       //有新建入库 查询商品明细得到 并传过来
	DetailId  int     //商品明细id
	FrontNum int	//盘点前数量		//盘点 专用
	Num       int     //出入库 盘盈亏 数量   //正数 表示入库 或者盘盈  负数表示 出库 或者盘亏
	LastNum int			//盘点后数量	//盘点 专用
	Price     float64 //采购单价    //采购 专用
	Remark    string  //备注       //盘点 专用
}

//查询入库 入参
type ArgsInStockGet struct {
	Uid int
	BusId         int      //商户信息
	common.Paging          //分页信息
	Time          []string //入库时间 开始 到 结束
	Flag          int      //1表示 今天  2表示 近七天  3表示 本月
}

//商品入库返回信息
type ReplyInStockPage struct {
	TotalNum int       //总条数
	List     []InStock //入库信息
}

//商品入库信息
type InStock struct {
	Id         int
	BillNum    string //单据号
	ShopId     int    //入库仓库   0表示总部 / 其他是分店id
	ShopName   string //门店名称
	BranchName string //分店名称
	Time       string //入库日期
	Type       int    `mapstructure:"type_id"` //入库类型id
	TypeName   string //入库类型名
	CreateTime string //创建时间
	Remark     string //备注
	Uid int
	Uname string
}

//查询一条出入库详情
type ArgsGetOneStock struct {
	BusId         int    //商户信息
	common.Paging        //分页信息
	BillNum       string //单据号
}

//一条 出入库详情返回
type ReplyOneStockPage struct {
	TotalNum int                //总条数
	List     []InOutStockDetail //出入库记录
}

//出入库记录
type InOutStockDetail struct {
	//基础明细信息
	DetailBase2
	Num int //出入库数量
}

//新增出库 入参
type ArgsOutStockAdd struct {
	common.BsToken               //商户信息
	common.Utoken
	Time           string        //入库时间
	Remark         string        //备注
	TypeId         int           //入库类型
	Details        []StockDetail //
}

//查询出库 入参
type ArgsOutStockGet struct {
	Uid int
	BusId         int      //商户信息
	common.Paging          //分页信息
	TypeId        int      //出库类型
	Time          []string //出库时间 开始 到 结束
	Flag          int      //1表示 今天  2表示 近七天  3表示 本月
}

//商品出库返回信息
type ReplyOutStockPage struct {
	TotalNum int       //总条数
	List     []InStock //入库信息
}

//商品出库信息
type OutStock struct {
	BillNum    string //单据号
	ShopId     int    //入库仓库   0表示总部 / 其他是分店id
	ShopName   string //门店名称
	BranchName string //分店名称
	Time       string //出库日期
	Type       int    //入库类型
	CreateTime string //创建时间
	Remark     string //备注
}

//新建盘点入参
type ArgsCheckAdd struct {
	common.BsToken               //商户信息
	common.Utoken
	Remark         string        //备注
	Status         int           //提交的状态。 1暂停  2完成  3作废
	Details        []StockDetail //商品明细详情
	BillNum string //单据号 //继续盘点入参
}

//继续盘点
type ContinueAdd struct {

}

//查询盘点单
type ArgsCheckGet struct {
	Uid int
	BusId         int      //商户信息
	common.Paging          //分页信息
	Status        int      //盘点状态
	IsBalance     int      //是否盘平 1是  2否
	Time          []string //出库时间 开始 到 结束
	Flag          int      //1表示 今天  2表示 近七天  3表示 本月
}

//盘点单查询返回
type ReplyCheckPage struct {
	TotalNum int         //总条数
	List     []CheckInfo //盘点单信息
}

type CheckInfo struct {
	BillNum    string //单据号
	Time       string //盘点日期
	Num        int    //盘点商品个数
	Status     int    //盘点状态
	IsBalance  int    //是否盘平
	CreateTime string //完成时间
	Remark     string //备注
	Uid  int
	Uname string
}

//盘点单作废
type ArgsCheckSet struct {
	common.BsToken
	BillNum string //单据号
}

//一条盘点单 详情
type ArgsCheckDetail struct {
	BusId int
	common.Paging
	BillNum string //单据号
	Flag    int    // 1查看详情  2继续盘点单
}

//一条 盘点详情返回
type ReplyCheckDetailPage struct {
	TotalNum int           //总条数
	List     []CheckDetail //出入库记录
}

//出入库记录
type CheckDetail struct {
	InOutStockDetail        //基础 进出库数据
	Remark           string //备注
	FrontNum	int			//盘点前的数量
	LastNum          int    //盘点后的数量
}

//按条件查询 全部 出入库
type ArgsInOutStockGet struct {
	BusId  int
	ShopId int //商户信息
	common.Paging
	DetailId int
	TypeId int      //出入库信息
	Time   []string //出库时间 开始 到 结束
	Flag   int      //1表示 今天  2表示 近七天  3表示 本月

}

//出入库信息返回
type ReplyInOutStockPage struct {
	TotalNum int
	List     []InOutStock
}

//出入库明细
type InOutStock struct {
	DetailBase2        //基础明细信息
	Num        int    //出入库数量 				// 出入库表
	BillNum    string //单据号					// 出入库表
	ShopId     int    //0表示总部   有值表示 门店	// 出入库表
	ShopName   string //门店名称
	BranchName string //分店名称
	TypeId     int    //出入库类型id					// 出入库表
	TypeName   string //出入库类型名称
	CreateTime string //创建时间				// 出入库表
	Surplus    int    //剩余库存				//明细库存表
	Uid int
	Uname string
}

//统一设置库存预警数量
type ArgsWarnNumSet struct {
	common.BsToken
	IsBus bool
	Num   int //数量
}

//查询所有预警 商品明细
type ArgsWarnDetail struct {
	//IsBus  bool
	BusId  int
	ShopId int
	common.Paging
	CateId int    //分类id
	Status int    //0开启预警 默认项   1关闭  2低库存  3正常
	//Name   string //产品名
}

//预警 商品明细 返回
type ReplyWarnDetailPage struct {
	TotalNum int
	//WarnNum  int
	List     []WarnDetail
}
type WarnDetail struct {
	// 基础明细数据
	Id         int    //商品明细id
	ProductId  int    //商品 id   //出入库添加商品需要传参
	Name       string //商品名称
	SpecIds    string //规格组合id
	SpecNames   string
	BarCode    string //条形码
	CateId     int    //商品分类id
	CateName   string //商品分类名称
	UnitId int
	UnitName string
	ImgId int
	ImgPath string

	// 商品明细库存查询
	Stock       int //库存数量
	UsableStock int //可用库存

	// 预警表查询
	Lowest   int //最低预警数量
	IsOpend  int //'预警开关 0开启 1关闭
	IsCustom int //'是否开启自定义预警数量 0不开启 1开启 默认不开启 开启后 统一设置数量绕过这个',
	Status int  //预警状态 0正常 1低库存
	WarnId int    //预警明细id
}

//关闭或开启 明细商品 预警
type ArgsDetailSet struct {
	common.BsToken
	WarnId int
	IsOpend  int //0开启 1关闭
}

//自定义预警值
type ArgsCustom struct {
	common.BsToken
	DetailId int //商品明细id
	Custom   int //预警值
}

//恢复默认预警值
type ArgsRegainDefault struct {
	common.BsToken
	DetailId int //商品明细id
}

//添加供应商信息
type ArgsSubAdd struct {
	common.BsToken
	SubDetail //供应商信息
}

//供应商信息
type SubDetail struct {
	Name     string //公司名称
	Code     string //公司编码
	Linkman  string //联系人
	Phone    string //手机
	Tel      string //公司电话
	Address  string //公司地址
	Remark   string //备注
	IsOpened int    //是否开启  0不开启 1开启
}

//编辑供应商信息
type ArgsSubUpdate struct {
	common.BsToken
	Id        int
	SubDetail //供应商信息
}

//查询一条详情
type ArgsSubGetOne struct {
	BusId int
	Id    int
}

//查询供应商
type ArgsSubGet struct {
	BusId int
	common.Paging
}

//供应商返回
type ReplySup struct {
	TotalNum int
	List     []SupInfo
}
type SupInfo struct {
	Id       int
	Name     string
	Code     string
	Linkman  string
	Phone    int
	Tel      string
	Address  string
	Remark   string
	IsOpened int `mapstructure:"is_opened"`
}

//采购入库
type ArgsPurAdd struct {
	common.BsToken
	common.Utoken
	Time    string        //入库日期
	SupId   int           //供应商id
	Remark  string        //备注
	Details []StockDetail //采购明细
}

//采购查询
type ArgsPurGet struct {
	Uid int
	BusId int
	common.Paging
	SupId int
	Time  []string //出库时间 开始 到 结束
	Flag  int      //1表示 今天  2表示 近七天  3表示 本月
}

//查询采购返回
type ReplyPur struct {
	TotalNum int
	List     []PurInfo
}
type PurInfo struct {
	BillNum    string //单据号
	Time       string //采购日期
	ShopId     int    //0总部 有值代表分店
	ShopName   string //门店名称
	BranchName string //分店名称
	TypeId     int    //入库类型
	TypeName   string
	SupId      int    //供应商
	SupName    string //供应商名
	CreateTime string //完成时间
	Remark     string //备注
	Uid int
	Uname string
}

//查询一条采购详情
type ArgsPurDetail struct {
	BusId int
	common.Paging
	BillNum string //单据号
}

//一条采购详情返回
type ReplyPurDetail struct {
	TotalNum int
	List     []PurDetailInfo
}
type PurDetailInfo struct {
	//基础商品明细
	DetailBase

	//出入库查询
	Num   int     //采购数量
	Price float64 //采购单价
}

//审核门店要货申请
type ArgsCheckRequire struct {
	common.BsToken
	common.Utoken
	BillNum string
	Details []StockDetail
}

//门店要货申请 驳回入参
type ArgsSetRequire struct {
	common.BsToken
	BillNum string
}

//查询所有供应商的名字和id
type ReplySubInfo struct {
	Id   int    //供应商id
	Name string //供应商名称
}

//查询出入库类型返回
type ReplyInoutType struct {
	Id   int    //出入库类型id
	Name string //出入库类型名称
}

//查询门店要货入库记录
type ArgsRequireInGet struct {
	common.Paging
	BusId  int
	ShopId int
	Time   []string //出库时间 开始 到 结束
	Flag   int      //1表示 今天  2表示 近七天  3表示 本月
}

//门店要货入库记录返回
type ReplyRequirePage struct {
	TotalNum int
	List     []ShopRequireIn
}
type ShopRequireIn struct {
	Id         int    `mapstructure:"id"`          //id
	BillNum    string `mapstructure:"bill_num_yh"` //要货单据号
	ShopId     int    `mapstructure:"shop_id"`     //门店id
	ShopName   string //门店名称
	BranchName string //分店名称
	TypeId     int    `mapstructure:"type_id"`     //出入库类型id
	TypeName   string `mapstructure:"TypeName"`    //出入库类型名称
	MakeTime   string `mapstructure:"create_time"` //制单日期
	Remark     string `mapstructure:"remark"`      //备注
	BillNumSQ  string `mapstructure:"bill_num"`    //申请单据号
	InTime     string `mapstructure:"time"`        // 入库时间
}

//根据单据号查询一条 要货详情
type ArgsRequireOneGet struct {
	BillNum string
	common.Paging
}

//查询一条要货入库 详情
type ReplyRequireDetail struct {
	TotalNum int
	List     []RequireDetail2
}
type RequireDetail2 struct {
	DetailBase

	//出入库查询
	ReqNum int
	ReaNum int
	ReaPrice float64
	ReaTotalMoney float64
	RecNum int
	RecPrice float64
	RecTotalMoney float64

}

//开启预警或者关闭预警
type ArgsWarnSet struct {
	common.BsToken
	IsBus    bool
	IsOpened int //是否开启预警 默认 0关闭  1开启
}

//查询预警状态
type ArgsWarnGet struct {
	IsBus  bool
	BusId  int
	ShopId int
}

//预警总状态和总数量返回
type ReplyWarnGet struct {
	IsOpened bool
	WarnNum int
}

//获取 待处理要货和预警 入参
type ArgsGetPending struct {
	common.BsToken
}

//获取 待处理要货和预警 返回值
type ReplyGetPending struct {
	WarnPendingNum int
	RequirePendingNum int
}

type Stock interface {

	//库存查询
	GetStock(ctx context.Context, args *ArgsStockGet, reply *ReplyDetailPage) error

	//新增入库
	AddInStock(ctx context.Context, args *ArgsStockInAdd, reply *int) error
	//查询所有入库
	GetInStock(ctx context.Context, args *ArgsInStockGet, reply *ReplyInStockPage) error
	//查询一条出入库详情
	GetOneInOutStockDetail(ctx context.Context, args *ArgsGetOneStock, reply *ReplyOneStockPage) error

	//新增出库
	AddOutStock(ctx context.Context, args *ArgsOutStockAdd, reply *int) error
	//查询所有出库
	GetOutStock(ctx context.Context, args *ArgsOutStockGet, reply *ReplyOutStockPage) error

	//新建 盘点 单
	AddStockCheck(ctx context.Context, args *ArgsCheckAdd, reply *int) error
	//查询盘点单
	GetStockCheck(ctx context.Context, args *ArgsCheckGet, reply *ReplyCheckPage) error
	//作废盘点单
	SetStockCheck(ctx context.Context, args *ArgsCheckSet, reply *bool) error
	//盘点单详情 - 恢复盘点单  共用接口
	GetStockCheckDetail(ctx context.Context, args *ArgsCheckDetail, reply *ReplyCheckDetailPage) error

	//按条件查询出入库明细
	GetInOutStock(ctx context.Context, args *ArgsInOutStockGet, reply *ReplyInOutStockPage) error

	//统一设置库存预警数量
	SetWarnNum(ctx context.Context, args *ArgsWarnNumSet, reply *bool) error
	//查询所有预警 商品明细
	GetWarn(ctx context.Context, args *ArgsWarnDetail, reply *ReplyWarnDetailPage) error
	//关闭或开启 明细商品 预警
	SetWarnDetail(ctx context.Context, args *ArgsDetailSet, reply *bool) error
	//自定义预警值
	SetCustom(ctx context.Context, args *ArgsCustom, reply *bool) error
	//恢复默认值
	RegainDefault(ctx context.Context, args *ArgsRegainDefault, reply *bool) error
	//开启预警或者关闭预警
	SetWarn(ctx context.Context, args *ArgsWarnSet, reply *bool) error
	//返回预警状态
	WarnGet(ctx context.Context, args *ArgsWarnGet, reply *ReplyWarnGet) error

	//添加供应商
	AddSub(ctx context.Context, args *ArgsSubAdd, reply *int) error
	//修改供应商
	UpdateSub(ctx context.Context, args *ArgsSubUpdate, reply *bool) error
	//查询一条供应商信息
	GetSubOne(ctx context.Context, args *ArgsSubGetOne, reply *SupInfo) error
	//查询供应商
	GetSub(ctx context.Context, args *ArgsSubGet, reply *ReplySup) error
	//根据供应商Id 查询供应商名
	GetSupName(ctx context.Context, args *[]int, reply *map[int]string) error
	//采购入库
	AddPur(ctx context.Context, args *ArgsPurAdd, reply *int) error
	//查询采购信息
	GetPur(ctx context.Context, args *ArgsPurGet, reply *ReplyPur) error
	//查询一条采购详情信息
	GetPurDetail(ctx context.Context, args *ArgsPurDetail, reply *ReplyPurDetail) error

	//总部要货申请修改状态 //审核操作  待审核 修改为 待入库
	CheckRequire(ctx context.Context, args *ArgsCheckRequire, reply *bool) error
	//总部要货申请 审核驳回   待审核 修改为  已驳回
	SetRequire(ctx context.Context, args *ArgsSetRequire, reply *bool) error
	//查询供应商名称和id接口
	GetSubNameAndId(ctx context.Context, busId *int, reply *[]ReplySubInfo) error
	/*//查询门店要货入库
	GetRequireInStock(ctx context.Context, args *ArgsRequireInGet, reply *ReplyRequirePage) error*/
	//查询一条要货入库详情
	GetOneRequireInStock(ctx context.Context, args *ArgsRequireOneGet, reply *ReplyRequireDetail) error

	//获取要货提货待处理数量和预警待处理数量
	GetPending(ctx context.Context,args *ArgsGetPending,reply *ReplyGetPending) error
}
