package public

import "context"

//区域入参
type ArgsRegion struct {
	RegionId int
}
//区域返回
type RegionInfo struct {
   RegionId int
   RegionName string
}

// 验证区域
type ArgsVerfiyRegion struct {
	Pid int // 省份
	Cid int // 市区
	Did int // 县区
}

//商圈返回
type ReplyDistrictInfo struct {
	DistrictId int //商圈ID
	DistrictName string //商圈名称
}

//根据多个区域Id获取商圈入参
type ArgsRegions struct {
	RegionIds []int
}

//根据多个商圈Id获取商圈入参
type ArgsDistricts struct {
	DistrictIds []int
}


//region
type RegionSub struct {
	RegionId int
	RegionName string
	Sub [] RegionDistrict
}
//district
type RegionDistrict struct {
	DistrictId int
	RegionId int
	DistrictName string
}


//根据cid获取关键字出参
type ReplyGetKeywordsByCid struct {
	Id int
	Keywords string // 关键字
	//Num int // 搜索频率
	//IsRec int // 是否推荐
}


type Region interface {
	//获取省份/直辖市
	GetProvince(ctx context.Context,reply *[]RegionInfo) error
	//获取城市/区街道
	GetRegion(ctx context.Context,args *ArgsRegion,reply *[]RegionInfo) error
	//获取单条区域信息
	GetByRegionId(ctx context.Context,args *ArgsRegion,reply *RegionInfo) error
	//获取多条区域信息
	GetByRegionIds(ctx context.Context,args *ArgsRegions,reply *[]RegionInfo)error
	//验证区域信息
	VerfiyRegion(ctx context.Context,args *ArgsVerfiyRegion,reply *bool) error
	//获取某区域下的商圈信息
	GetDistrictByRegionId(ctx context.Context,args *ArgsRegion,reply *[]ReplyDistrictInfo) error
	//根据多个区域获取商圈信息-RPC调用
	GetDistrictByRegionIds(ctx context.Context,args *ArgsRegions,reply *[]ReplyDistrictInfo) error
	//根据多个商圈id获取多个商圈信息
	GetDistrictByDistrictIds(ctx context.Context,args *ArgsDistricts,reply *[]ReplyDistrictInfo) error
	//根据多个城市id获取多个城市名称
	GetCityNamesByCids(ctx context.Context,cids *[]int,reply *[]RegionInfo) error
	//根据cid获取关键字
	GetKeywordsByCid(ctx context.Context,cid *int,reply *[]ReplyGetKeywordsByCid)error
}
