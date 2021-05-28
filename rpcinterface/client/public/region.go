package public

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/public"
)

type Region struct {
	client.Baseclient
}

//实例化
func (r *Region) Init() *Region {
	r.ServiceName = "rpc_public"
	r.ServicePath = "Region"
	return r
}


//根据城市名获取城市ID
func (r *Region) GetRegionNameByRegionId(ctx context.Context,regionName *string,reply *public.RegionInfo) error{
	return r.Call(ctx,"GetRegionNameByRegionId",regionName,reply)
}

//获取城市区及区下面的商圈
func (r *Region) GetRegionAndDirect(ctx context.Context,regionId *int,reply *[]public.RegionSub) error{
	return r.Call(ctx,"GetRegionAndDirect",regionId,reply)
}

//获取省份/直辖市
func (r *Region) GetProvince(ctx context.Context,reply *[]public.RegionInfo) error{
	return r.Call(ctx,"GetProvince","",reply)
}

//获取省份/直辖市
func (r *Region) GetProvinceByPids(ctx context.Context,pids *[]int,reply *[]public.RegionInfo) error{
	return r.Call(ctx,"GetProvinceByPids",pids,reply)
}

//获取城市
func (r *Region) GetRegion(ctx context.Context,args *public.ArgsRegion,reply *[]public.RegionInfo) error{
	return r.Call(ctx,"GetRegion",args,reply)
}
//获取单区域信息
func (r *Region) GetByRegionId(ctx context.Context,args *public.ArgsRegion,reply *public.RegionInfo) error{
	return r.Call(ctx,"GetByRegionId",args,reply)
}

//获取多条区域信息
func (r *Region) GetByRegionIds(ctx context.Context,args *public.ArgsRegions,reply *[]public.RegionInfo)error{
	return r.Call(ctx,"GetByRegionIds",args,reply)
}

// 验证区域信息
func (r *Region)VerfiyRegion(ctx context.Context, args *public.ArgsVerfiyRegion, reply *bool) error{
	return r.Call(ctx,"VerfiyRegion",args,reply)
}


// 获取某区域下的商圈信息
func (r *Region)GetDistrictByRegionId(ctx context.Context, args *public.ArgsRegion, reply *[]public.ReplyDistrictInfo) error{
	return r.Call(ctx,"GetDistrictByRegionId",args,reply)
}

//根据多个区域获取商圈信息-RPC调用
func (r *Region) GetDistrictByRegionIds(ctx context.Context,args *public.ArgsRegions,reply *[]public.ReplyDistrictInfo) error{
	return r.Call(ctx,"GetDistrictByRegionIds",args,reply)
}

//根据多个区域获取商圈信息-RPC调用
func (r *Region) GetDistrictByDistrictIds(ctx context.Context,args *public.ArgsDistricts,reply *[]public.ReplyDistrictInfo) error{
	return r.Call(ctx,"GetDistrictByDistrictIds",args,reply)
}

//根据多个城市id获取多个城市名称
func (r *Region) GetCityNamesByCids(ctx context.Context,cids *[]int,reply *[]public.RegionInfo) error {
	return r.Call(ctx,"GetCityNamesByCids",cids,reply)
}

//根据cid获取关键字
func (r *Region)GetKeywordsByCid(ctx context.Context,cid *int,reply *[]public.ReplyGetKeywordsByCid)error{
	return r.Call(ctx,"GetKeywordsByCid",cid,reply)
}

//根据区域id获取区域名称
func (r *Region) GetDNameByDid(ctx context.Context, did int, dname *string) error {
	return r.Call(ctx,"GetDNameByDid",did,dname)
}
