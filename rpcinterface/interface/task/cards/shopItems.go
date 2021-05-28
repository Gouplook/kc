package cards

import "context"

//门店项目

const (
	ITEMTYPE_single = 1 //单项目
	ITEMTYPE_sm = 2 //套餐
	ITEMTYPE_card = 3 //综合卡
	ITEMTYPE_hcard = 4 //限时卡
	ITEMTYPE_ncard = 5 //限次卡
	ITEMTYPE_hncard = 6 //限时限次卡
	ITEMTYPE_rcard = 7 //充值卡
	ITEMTYPE_icard = 8 //身份卡

	OPTTYPE_add = 1 //添加
	OPTTYPE_edit = 2 //修改
	OPTTYPE_del = 3 //删除

)

type ShopItems struct {
	ShopItemId int //项目在门店的id
	ItemId int 	//项目id
	ShopId int //门店的id
	ItemType int //项目类型
}

type ShopItemsIf interface {
	//设置项目到任务队列
	SetItems( ctx context.Context, args *ShopItems, reply *bool ) error
}