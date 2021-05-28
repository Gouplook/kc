package bus

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	ZHIFUBAO = 1
	WEIXIN   = 2
	XIANJIN  = 3
	OTHER    = 4
)
const (
	//消毒地点
	DISINFECT_ADDRESS_SITTING_ROOT = 1 //客厅
	DISINFECT_ADDRESS_CORRIDOR     = 2 //走廊
	DISINFECT_ADDRESS_RECEPTION    = 3 //接待处
	DISINFECT_ADDRESS_RESTAREA     = 4 //休息区

	//消毒类型
	DISINFECT_TYPE_UV           = 1 //紫外线照射
	DISINFECT_TYPE_DISINFECTANT = 2 //消毒剂喷洒
	DISINFECT_TYPE_ALCOHOL      = 3 //酒精消毒

	//消毒物品
	DISINFECT_GOODS_TOWEL   = 1 //毛巾
	DISINFECT_GOODS_SCISSOR = 2 //剪刀
	DISINFECT_GOODS_FOOTTUB = 3 //足浴盆

	//垃圾处理方式
	RUBBISH_DISPOSE_INCINERATION = 1 //焚烧
	RUBBISH_DISPOSE_ISOLATION    = 2 //隔离

	//垃圾类型
	RUBBISH_TYPE_DO  = 1 //干垃圾
	RUBBISH_TYPE_WET = 2 //湿垃圾

	COMPANY_TYPE_GROUP = 1  //集团
	COMPANY_TYPE_BRAND = 2	 //品牌发卡
	COMPANY_TYPE_SCALE = 3	//规格发卡企业/商户
	COMPANY_TYPE_OTHER = 4	//其他发卡企业/商户
	COMPANY_TYPE_INDIVIDUAL = 5	//个体工商户

	DisinfectNoDel=1//未删除
	DisinfectIsDel=2//删除
)

var CompanyTypeMap = map[int]string {
	COMPANY_TYPE_GROUP : "集团",
	COMPANY_TYPE_BRAND : "品牌发卡",
	COMPANY_TYPE_SCALE : "规格发卡企业/商户",
	COMPANY_TYPE_OTHER : "其他发卡企业/商户",
	COMPANY_TYPE_INDIVIDUAL : "个体工商户",
}

//消毒地点
var DisinfectAddressMap = map[int]string{
	DISINFECT_ADDRESS_SITTING_ROOT: "客厅",
	DISINFECT_ADDRESS_CORRIDOR:     "走廊",
	DISINFECT_ADDRESS_RECEPTION:    "接待处",
	DISINFECT_ADDRESS_RESTAREA:     "休息区",
}

func DisinfectAddressList() []int {
	return []int{DISINFECT_ADDRESS_SITTING_ROOT, DISINFECT_ADDRESS_CORRIDOR, DISINFECT_ADDRESS_RECEPTION,
		DISINFECT_ADDRESS_RESTAREA}
}

//消毒类型
var DisinfectTypeMap = map[int]string{
	DISINFECT_TYPE_UV:           "紫外线照射",
	DISINFECT_TYPE_DISINFECTANT: "消毒剂喷洒",
	DISINFECT_TYPE_ALCOHOL:      "酒精消毒",
}

func DisinfectTypeList() []int {
	return []int{DISINFECT_TYPE_UV, DISINFECT_TYPE_DISINFECTANT, DISINFECT_TYPE_ALCOHOL}
}

//消毒物品
var DisinfectGoodsMap = map[int]string{
	DISINFECT_GOODS_TOWEL:   "毛巾",
	DISINFECT_GOODS_SCISSOR: "剪刀",
	DISINFECT_GOODS_FOOTTUB: "足浴盆",
}

func DisinfectGoodsList() []int {
	return []int{DISINFECT_GOODS_TOWEL, DISINFECT_GOODS_SCISSOR, DISINFECT_GOODS_FOOTTUB}
}

//垃圾处理方式
var RubbishDisposeMap = map[int]string{
	RUBBISH_DISPOSE_INCINERATION: "焚烧",
	RUBBISH_DISPOSE_ISOLATION:    "隔离",
}

func RubbishDisposeList() []int {
	return []int{RUBBISH_DISPOSE_INCINERATION, RUBBISH_DISPOSE_ISOLATION}
}

//垃圾类型
var RubbishTypeMap = map[int]string{
	RUBBISH_TYPE_DO:  "干垃圾",
	RUBBISH_TYPE_WET: "湿垃圾",
}

func RubbishTypeList() []int {
	return []int{RUBBISH_TYPE_DO, RUBBISH_TYPE_WET}
}

var DisberseTypeMap = map[int]string{
	ZHIFUBAO: "支付宝",
	WEIXIN:   "微信",
	XIANJIN:  "现金",
	OTHER:    "其他",
}

//前台支出类目入参
type ArgsDisburseCategoryInfo struct {
	common.BsToken        //企业信息
	CategoryId     int    //类目id
	CategoryName   string //类目名
}

//前台支出类目返回
type ReplyDisburseCategoryRes struct {
	Status int //返回状态
}

//前台支出类目信息
type ReplyCategoryType struct {
	CategoryId   int    //类目id
	CategoryName string //类目名
}

//新增支出明细入参
type ArgsDisburseDetailInfo struct {
	common.BsToken         //企业信息
	DisburseId     int     //支出明细id
	CategoryId     int     //支出类目id
	DisburseMoney  float64 //支出金额
	//DisburseDept   string  //支出部门
	DisburseType int    //支出方式
	StaffId      int    //技师id
	DisburseTime string //支出时间
	Remark       string //备注
}

//新增支出明细返回
type ReplyDisburseDetailRes struct {
	Status int //支出明细返回
}

//查询支出明细入参
type ArgsDisburseDetailReq struct {
	common.Paging
	BusId      int
	CategoryId int    //支出类目id
	ShopId     int    //分店id
	StartTime  string //开始时间
	EndTime    string //结束时间
}

//查询支出明细返回
type ReplyDisburseDetailInfo struct {
	TotalNum int
	Lists    []DisburseDetailInfo
}

//支出明细
type DisburseDetailInfo struct {
	DisburseId       int //支出明细id
	ShopId           int //门店id
	ShopName         string
	BranchName       string
	CategoryId       int //支出类目id
	CategoryName     string
	DisburseMoney    float64 //支出金额
	DisburseDeptId   int     //支出岗位id
	DisburseDeptName string  //支出岗位名称
	DisburseType     int     //支出方式
	StaffId          int     //经办人
	StaffName        string
	StaffNickName    string
	DisburseTime     int    //支出时间
	DisburseTimeStr  string //支出时间
	Remark           string //备注
}

//前台 疫情防控 环境消毒入参
type ArgsEpidemicSettingInfo struct {
	common.BsToken        //企业信息
	common.Paging
	Id             int    //id
	ShopId         int    //分店id
	Address        int    //消毒地点
	Time           string //消毒时间,时间格式：2006-01-02 15:00
	Type           int    //消毒类型
	Duration       int    //消毒持续时间,单位分钟
	Executor       int    //执行人
}

//前台 疫情防控 环境消毒 返回
type ReplyEpidemicSetting struct {
	Status int //返回状态码
}
type EpidemicSettingGetBase struct {
	Id           int   //id
	ShopId       int   //分店id
	BranchName string
	Address      string   //消毒地点
	Time         int64 //消毒时间
	TimeStr      string
	Type         string    //消毒类型
	Duration     int    //消毒持续时间
	Executor     int //执行人
	ExecutorName string
}
//前台 疫情防控 环境消毒 查询返回
type ReplyEpidemicSettingGet struct {
	TotalNum int
	Lists []EpidemicSettingGetBase
}

//前台 疫情防控 用品消毒 信息入参
type ArgsEpidemicTackleInfo struct {
	common.BsToken        //企业信息
	common.Paging
	Id             int    //用品消毒id
	ShopId         int    //分店id
	Time           string //消毒时间,时间格式：2006-01-02 15:00
	Good           int    //消毒物品
	Type           int    //消毒类型
	Duration       int    //持续时间,单位分钟
	Executor       int    //执行人
}

//前台 疫情防控 用品消毒 状态返回
type ReplyEpidemicTackle struct {
	Status int
}
type EpidemicTackleGetBase struct {
	Id           int   //用品消毒id
	ShopId       int   //分店id
	BranchName string
	Time         int64 //消毒时间
	TimeStr      string
	Good         string    //消毒物品
	Type         string    //消毒类型
	Duration     int    //持续时间
	Executor     int //执行人
	ExecutorName string
}
//前台 疫情防控 用品消毒 信息返回
type ReplyEpidemicTackleGet struct {
	TotalNum int
	Lists []EpidemicTackleGetBase
}

//前台 疫情防控 垃圾处理 信息入参
type ArgsEpidemicGarbageInfo struct {
	common.BsToken        //企业信息
	common.Paging
	Id             int    //垃圾处理id
	ShopId         int    //分店id
	Time           string //处理时间
	Type           int    //垃圾类型
	Method         int    //处理方式
	Executor       int    //执行人
}

//前台 疫情防控 垃圾处理 状态返回
type ReplyEpidemicGarbage struct {
	Status int
}
type EpidemicGarbageGetBase struct {
	Id           int   //垃圾处理id
	BranchName string
	ShopId       int   //分店id
	Time         int64 //检查时间
	TimeStr      string
	Type         string //垃圾类型
	Method       string //处理方式
	Executor     int //执行人
	ExecutorName string
}
//前台 疫情防控 垃圾处理 信息返回
type ReplyEpidemicGarbageGet struct {
	TotalNum int
	Lists []EpidemicGarbageGetBase
}

//前台 疫情防控 技师健康 信息入参
type ArgsEpidemicTechnicianInfo struct {
	common.BsToken        //企业信息
	common.Paging
	Id             int    //
	ShopId         int    //门店id
	TechId         int    //技师id
	Time           string //检查时间
	IsHeat         int    //是否发热 1是 0否
	IsCough        int    //是否咳嗽 1是 0否
	IsSymptom      int    //是否有呼吸道症状 1是 0否
	Temperature    string //体温
}

//前台 疫情防控 技师健康 状态返回
type ReplyEpidemicTechnician struct {
	Status int
}

type EpidemicTechnicianGetBase struct {
	Id          int //
	ShopId      int //门店id
	BranchName string
	TechId      int //技师id
	TechName    string
	Phone       string
	Time        int64 //检查时间
	TimeStr     string
	IsHeat      int    //是否发热 1是 0否
	IsCough     int    //是否咳嗽 1是 0否
	IsSymptom   int    //是否有呼吸道症状 1是 0否
	Temperature string //体温
}
//前台 疫情防控 技师健康 信息返回
type ReplyEpidemicTechnicianGet struct {
	TotalNum int
	Lists []EpidemicTechnicianGetBase
}
type EpidemicTechnician struct {
	OnlineStaff       int // 在职员工人数
	HealthNormalNum   int //健康状况正常人数
	HealthAbnormalNum int // 将抗状况异常人数
}

//查看门店当天的防疫情况出参
type ReplyGetShopEpidemic struct {
	EpidemicSetting []EpidemicSettingGetBase // 环境消毒
	EpidemicTackle  []EpidemicTackleGetBase  // 物品消毒
	StaffHealth     EpidemicTechnician        // 员工健康状况
}

//前台 复工管理 信息入参
type ArgsReturnWork struct {
	common.BsToken
	Id	int
	CompanyName	string
	CompanyType	int
	Did	int
	CompanyAddress	string
	UscCode	string
	ReturnWorkTime	string
	LegalRepresentative	string
	LrContactDetails	string
	EnterpriseContact	string
	EcContactDetails	string
	TotalPeopleNum	int
	ShanghaiHousehold	int
	NotShanghaiHousehold	int
	Thermometer	int
	ThermometerImg	string
	Masks	int
	MasksImg	string
	DisinfectantFluid	int
	DisinfectantFluidImg	string
	LatexGloves	int
	LatexGlovesImg	string
	HandSoap	int
	HandSoapImg	string
	Fungicide	int
	FungicideImg	string
	SchemeFile	string
	IsRead	int
}

//前台 复工管理 返回
type ReplyReturnWork struct {
	Status int
}

//前台 复工管理 查询返回
type ReplyReturnWorkGet struct {
	Verify int
}

//前台 修改支出汇总入参
type ArgsUpdateBusDisburse struct {
	BusId int
	ShopId int
	Cid            int
	Money          float64
	Month int
}

//前台 查询支出汇总入参
type ArgsGetBusDisburse struct {
	common.BsToken
	StartTime string
	EndTime string
	CategoryId int
}

//前台 查询支出汇总返回
type ReplyGetBusDisburse struct {
	Details []CateDetail
	TotalDisburse string
}

type CateDetail struct {
	Cid int
	CName string
	Money string
}

//添加 门店企业复工员工登记 入参
type ArgsAddStaffRegister struct {
	common.BsToken
	Name	string
	Age	int
	Sex	int
	IdNum	string
	Address	string
	Phone	int
	IsLeave	int
	Remark	string
}
//查询 门店企业复工员工登记 入参
type ArgsGetStaffRegister struct {
	common.Paging
	common.BsToken
}
//添加 门店企业复工防控日报 入参
type ArgsAddDefendDaily struct {
	common.BsToken
	OfficeNum	int
	Performance	int
	Gt372	int
	Disisolation	int
	HomeQuarantine	int
}
//查询 门店企业复工防控日报 入参
type ArgsGetDefendDaily struct {
	common.BsToken
	common.Paging
}
//查询 门店企业复工员工登记 返回值
type ReplyGetStaffRegister struct {
	TotalNum int
	Lists []GetStaffRegister
}
type GetStaffRegister struct {
	Id int
	Name	string
	Age	int
	Sex	int
	IdNum	string
	Address	string
	Phone	int
	IsLeave	int
	Remark	string
	Status	int
	TrafficPermit	string
	CreateTime string
}
//添加 门店企业复工防控日报 返回值
type ReplyGetDefendDaily struct {
	TotalNum int
	Lists []GetDefendDaily
}
type GetDefendDaily struct {
	Id             int
	OfficeNum      int
	Performance    int
	Gt372          int
	Disisolation   int
	HomeQuarantine int
	CreateTime     string
}

//企业类型
type CompayType struct {
	Id int
	Name string
}

//===========
type DisinfectFrontBase struct {
	Id int
	Name string
	Ctime int64
	CtimeStr string
}
type ArgsAddOrUpdateDelDisinfectFrontBase struct {
	common.BsToken
	Id int
	Name string
}
type GetDisinfectFrontBase struct {
	common.BsToken
	common.Paging
	IsLimit bool //是否分页
}

type ReplyGetDisinfectFrontCommon struct {
	TotalNum int
	Lists []DisinfectFrontBase
}



//前台服务接口
type Front interface {
	//前台总部支出汇总查询
	GetBusDisburseSum(ctx context.Context, args *ArgsGetBusDisburse, reply *ReplyGetBusDisburse) error

	//获取 支出类目
	GetDisburseCategory(ctx context.Context, args *ArgsDisburseCategoryInfo, reply *[]ReplyCategoryType) error
	//添加 支出类目
	AddDisburseCategory(ctx context.Context, args *ArgsDisburseCategoryInfo, reply *ReplyDisburseCategoryRes) error
	//修改 支出类目
	UpdateDisburseCategory(ctx context.Context, args *ArgsDisburseCategoryInfo, reply *ReplyDisburseCategoryRes) error
	//删除 支出类目
	DeleteDisburseCategory(ctx context.Context, args *ArgsDisburseCategoryInfo, reply *ReplyDisburseCategoryRes) error

	//实现 支出明细的 添加
	AddDisburseDetail(ctx context.Context, args *ArgsDisburseDetailInfo, reply *ReplyDisburseDetailRes) error
	//实现 支出明细的 获取
	GetDisburseDetail(ctx context.Context, args *ArgsDisburseDetailReq, reply *ReplyDisburseDetailInfo) error
	//实现 获取u一条支出明细
	GetOneDisburseDetail(ctx context.Context, args int, reply *DisburseDetailInfo) error
	//实现 支出明细的 修改
	UpdateDisburseDetail(ctx context.Context, args *ArgsDisburseDetailInfo, reply *ReplyDisburseDetailRes) error
	//实现 支出明细的 删除
	DeleteDisburseDetail(ctx context.Context, args *ArgsDisburseDetailInfo, reply *ReplyDisburseDetailRes) error

	//添加/更新消毒地点
	AddOrUpdateDisinfectAddress(ctx context.Context,args *ArgsAddOrUpdateDelDisinfectFrontBase,replyId *int)error
	//删除消毒地点-软删除
	DelDisinfectAddress(ctx context.Context,args *ArgsAddOrUpdateDelDisinfectFrontBase,reply *bool)error
	//消毒地点数据获取
	GetDisinfectAddress(ctx context.Context,args *GetDisinfectFrontBase,reply *ReplyGetDisinfectFrontCommon)error
	//添加/更新消毒类型
	AddOrUpdateDisinfectType(ctx context.Context,args *ArgsAddOrUpdateDelDisinfectFrontBase,replyId *int)error
	//删除消毒类型-软删除
	DelDisinfectType(ctx context.Context,args *ArgsAddOrUpdateDelDisinfectFrontBase,reply *bool)error
	//消毒类型数据获取
	GetDisinfectType(ctx context.Context,args *GetDisinfectFrontBase,reply *ReplyGetDisinfectFrontCommon)error
	//添加/更新消毒物品
	AddOrUpdateDisinfectGood(ctx context.Context,args *ArgsAddOrUpdateDelDisinfectFrontBase,replyId *int)error
	//删除消毒物品-软删除
	DelDisinfectGood(ctx context.Context,args *ArgsAddOrUpdateDelDisinfectFrontBase,reply *bool)error
	//消毒物品数据获取
	GetDisinfectGoods(ctx context.Context,args *GetDisinfectFrontBase,reply *ReplyGetDisinfectFrontCommon)error

	//添加 环境消毒明细
	AddSetting(ctx context.Context, args *ArgsEpidemicSettingInfo, reply *ReplyEpidemicSetting) error
	//删除 一条环境消毒明细
	DeleteSetting(ctx context.Context, args *ArgsEpidemicSettingInfo, reply *ReplyEpidemicSetting) error
	//查询 环境消毒明细
	GetSetting(ctx context.Context, args *ArgsEpidemicSettingInfo, reply *ReplyEpidemicSettingGet) error

	//添加 用品消毒明细
	AddTackle(ctx context.Context, args *ArgsEpidemicTackleInfo, reply *ReplyEpidemicTackle) error
	//删除 一条用品消毒记录
	DeleteTackle(ctx context.Context, args *ArgsEpidemicTackleInfo, reply *ReplyEpidemicTackle) error
	//查询 用品消毒 记录
	GetTackle(ctx context.Context, args *ArgsEpidemicTackleInfo, reply *ReplyEpidemicTackleGet) error


	//添加/更新垃圾处理方式
	AddOrUpdateGarbageDisposetType(ctx context.Context,args *ArgsAddOrUpdateDelDisinfectFrontBase,replyId *int)error
	//删除垃圾处理方式-软删除
	DelGarbageDisposetType(ctx context.Context,args *ArgsAddOrUpdateDelDisinfectFrontBase,reply *bool)error
	//垃圾处理方式数据获取
	GetGarbageDisposetType(ctx context.Context,args *GetDisinfectFrontBase,reply *ReplyGetDisinfectFrontCommon)error
	//添加/更新垃圾类型
	AddOrUpdateGarbageType(ctx context.Context,args *ArgsAddOrUpdateDelDisinfectFrontBase,replyId *int)error
	//删除垃圾类型-软删除
	DelGarbageType(ctx context.Context,args *ArgsAddOrUpdateDelDisinfectFrontBase,reply *bool)error
	//垃圾类型数据获取
	GetGarbageType(ctx context.Context,args *GetDisinfectFrontBase,reply *ReplyGetDisinfectFrontCommon)error

	//添加 垃圾处理 记录
	AddGarbage(ctx context.Context, args *ArgsEpidemicGarbageInfo, reply *ReplyEpidemicGarbage) error
	//删除 一条垃圾处理 记录
	DeleteGarbage(ctx context.Context, args *ArgsEpidemicGarbageInfo, reply *ReplyEpidemicGarbage) error
	//查询 垃圾处理 记录
	GetGarbage(ctx context.Context, args *ArgsEpidemicGarbageInfo, reply *ReplyEpidemicGarbageGet) error

	//添加 技师健康 记录
	AddTechnician(ctx context.Context, args *ArgsEpidemicTechnicianInfo, reply *ReplyEpidemicTechnician) error
	//删除 一条技师健康 记录
	DeleteTechnician(ctx context.Context, args *ArgsEpidemicTechnicianInfo, reply *ReplyEpidemicTechnician) error
	//查询 技师健康 记录
	GetTechnician(ctx context.Context, args *ArgsEpidemicTechnicianInfo, reply *ReplyEpidemicTechnicianGet) error
	//查看门店当天的防疫情况
	GetShopEpidemic(ctx context.Context, shopId *int, reply *ReplyGetShopEpidemic) error

	//添加 复工申请
	AddWork(ctx context.Context, args *ArgsReturnWork, reply *ReplyReturnWork) error
	//查询 审核状态
	GetVerify(ctx context.Context, args *ArgsReturnWork, reply *ReplyReturnWorkGet) error
	//门店添加员工登记
	AddStaffRegister(ctx context.Context, args *ArgsAddStaffRegister, reply *int) error
	//门店查询员工登记
	GetStaffRegister(ctx context.Context, args *ArgsGetStaffRegister, reply *ReplyGetStaffRegister) error
	//门店添加防控日报
	AddDefendDaily(ctx context.Context, args *ArgsAddDefendDaily, reply *int) error
	//门店查询防控日报
	GetDefendDaily(ctx context.Context, args *ArgsGetDefendDaily, reply *ReplyGetDefendDaily) error
}
