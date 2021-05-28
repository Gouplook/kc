package reservation

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/reservation"
)

type Record struct {
	client.Baseclient
}

func (r *Record) Init() *Record {
	r.ServiceName = "rpc_reservation"
	r.ServicePath = "Record"
	return r
}

//通用预约接口
func (r *Record) CommonReserve(ctx context.Context, params *reservation.CommonReserveParams, replies *reservation.ReplyCommonReserveParams) (err error) {
	err = r.Call(ctx, "CommonReserve", params, replies)
	return
}

//编辑预约
func (r *Record) EditReservation(ctx context.Context, params *reservation.EditReservationParams, replies *bool) (err error) {
	err = r.Call(ctx, "EditReservation", params, replies)
	return
}

//更新预约记录状态(取消,确认,完成)
func (r *Record) ChangeStatus(ctx context.Context, params *reservation.ModifyReservationParams, replies *bool) (err error) {
	return r.Call(ctx, "ChangeStatus", params, replies)
}

//获取预约记录信息
func (r *Record) GetReservationRecord(ctx context.Context, params *reservation.GetReservationRecordParams, replies *reservation.ReplyGetReservationRecordParams) (err error) {
	err = r.Call(ctx, "GetReservationRecord", params, replies)
	return
}

//根据预约ids获取预约条目-rpc
func (r *Record) GetItemsByReservationIdsRpc(ctx context.Context, params *reservation.ArgsGetItemsByReservationIds, replies *reservation.ReplyGetItemsByReservationIds) (err error) {
	err = r.Call(ctx, "GetItemsByReservationIdsRpc", params, replies)
	return
}

//预约管控列表
func (r *Record) GetReservationControlList(ctx context.Context, params *reservation.GetReservationRecordListParams, replies *reservation.GetReservationRecordListReplies) (err error) {
	err = r.Call(ctx, "GetReservationControlList", params, replies)
	return
}

//管控RPC内部
func (r *Record) GetReservationControlListRpc(ctx context.Context, ids *[]int, replies *reservation.GetReservationRecordListReplies) (err error) {
	err = r.Call(ctx, "GetReservationControlListRpc", ids, replies)
	return
}

//门店预约看板
func (r *Record) GetReservationRecordListByShopID(ctx context.Context, params *reservation.ArgsReservationRecordListByShopID, replies *reservation.ReplyReservationRecordListByShopID) (err error) {
	err = r.Call(ctx, "GetReservationRecordListByShopID", params, replies)
	return
}

//设置预约Item失效
func (r *Record) InvalidateReservationItem(ctx context.Context, params *reservation.InvalidateReservationItemParams, replies *bool) (err error) {
	err = r.Call(ctx, "InvalidateReservationItem", params, replies)
	return
}

//获取员工预约条目列表
func (r *Record) GetReservationItemList(ctx context.Context, params *reservation.GetReservationItemListParams, replies *reservation.GetReservationItemListReplies) (err error) {
	err = r.Call(ctx, "GetReservationItemList", params, replies)
	return
}

//获取用户预约条目列表
func (r *Record) GetReservationItemListByUid(ctx context.Context, args *reservation.ArgsReservationByUser, reply *reservation.ReplyReservationByUser) error {
	return r.Call(ctx, "GetReservationItemListByUid", args, reply)
}

//门店预约订单管理
func (r *Record)GetKxbOrderedList(ctx context.Context,args *reservation.ArgsGetKxbOrderedList,reply *reservation.ReplyGetKxbOrderedList)error{
	return r.Call(ctx, "GetKxbOrderedList", args, reply)
}

//查询预约中用户卡包单项目的预约次数（未结束的）
func (r *Record)GetCardPackageSingleNum(ctx context.Context,args *reservation.ArgsGetCardPackageSingleNum,reply *reservation.ReplyGetCardPackageSingleNum)error{
	return r.Call(ctx, "GetCardPackageSingleNum", args, reply)
}

//根据条件获取预约基础信息
func (r *Record)GetSimpleReserOrderRpc(ctx context.Context,args *reservation.ArgsGetSimpleReserOrderRpc,reply *reservation.ReplyGetSimpleReserOrderRpc)error{
	return r.Call(ctx, "GetSimpleReserOrderRpc", args, reply)
}