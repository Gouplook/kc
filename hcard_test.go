package main_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"git.900sui.cn/kc/base/utils"
	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

// TestHcardModels 生成限时卡model
func TestHcardModels(t *testing.T) {
	utils.CreateModel("hcard")        // 卡项服务-限时卡表
	utils.CreateModel("hcard_give")   // 卡项-限时卡赠送的单服务表
	utils.CreateModel("hcard_shop")   // 卡项服务-限时卡适用门店
	utils.CreateModel("hcard_single") // 卡项服务-限时卡包含的单项目
	utils.CreateModel("shop_hcard")   // 卡项服务-已添加到门店的限时卡
	utils.CreateModel("hcard_ext")    // 卡项管理-限时卡扩展信息表
}

// TestAddHcard TestAddHcard
func TestAddHcard(t *testing.T) {
	busToken := getBusToken()
	hcardLogic := new(logics.HcardLogic)
	args := &cards.ArgsAddHcard{
		BsToken: busToken,
		HcardBase: cards.HcardBase{
			Name:          "优化-------333限时卡名称bbbb",
			SortDesc:      "this is hcard two",
			RealPrice:     10,
			Price:         50,
			ServicePeriod: 12,
		},
		IncludeSingles: []cards.IncInfSingle{
			{
				SingleID: 12,
			},
		},
		GiveSingles: []cards.HcardSingle{
			{
				SingleID: 11,
				Num:      3,
			},
		},
		Notes: []cards.CardNote{
			{
				Notes: "温馨提示1",
			},
		},
		ImgHash: "b82275a6c75a43fd836d25be534932ad",
	}
	t.Log(hcardLogic.AddHcard(context.Background(), args))
}

// TestEditHcard 修改限时卡数据
func TestEditHcard(t *testing.T) {
	busToken := getBusToken()
	hcardLogic := new(logics.HcardLogic)
	args := &cards.ArgsEditHcard{
		BsToken: busToken,
		HcardID: 19,
		HcardBase: cards.HcardBase{
			Name:          "修改限时卡hcard one",
			SortDesc:      "修改限时卡描述",
			RealPrice:     66,
			Price:         150,
			ServicePeriod: 13,
		},
		IncludeSingles: []cards.IncInfSingle{
			{
				SingleID: 12,
			},
		},
		GiveSingles: []cards.HcardSingle{
			{
				SingleID: 13,
				Num:      12,
			},
		},
		Notes: []cards.CardNote{
			{
				Notes: "温馨提示1",
			},
		},
		ImgHash: "b82275a6c75a43fd836d25be534932ad",
	}
	t.Log(hcardLogic.EditHcard(context.Background(), args))
}

func TestDeleteHcard(t *testing.T) {
	hcardLogics := new(logics.HcardLogic)
	args := &cards.ArgsDeleteHcard{
		BsToken:  getBusToken(),
		HcardIds: []int{15, 16},
	}
	if err := hcardLogics.DeleteHcard(context.Background(), args); err != nil {
		t.Error(err.Error())
	}

}

// TestHcardInfo 获取限时卡详情数据
func TestHcardInfo(t *testing.T) {
	hcardLogic := new(logics.HcardLogic)
	hcardID := 39
	reply, err := hcardLogic.HcardInfo(context.Background(), hcardID)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

// TestBusHcardPage 获取总店的限时卡列表
func TestBusHcardPage(t *testing.T) {
	busID, start, limit, shopId := 51, 0, 10, 429
	isGround := ""
	hcardLogic := new(logics.HcardLogic)
	reply, _ := hcardLogic.BusHcardPage(context.Background(), busID, shopId, start, limit, isGround,true)
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//　TestSetHcardShop 总店设置限时卡的适用门店
func TestSetHcardShop(t *testing.T) {
	args := &cards.ArgsSetHcardShop{
		BsToken:   getBusToken(),
		HcardIDs:  []int{29},
		ShopIDs:   []int{1},
		IsAllShop: false,
	}
	hcardLogic := new(logics.HcardLogic)
	t.Log(hcardLogic.SetHcardShop(context.Background(), args))
}

// TestDownUpHcard 总店上下架限时卡
func TestDownUpHcard(t *testing.T) {
	hcardLogic := new(logics.HcardLogic)
	args := &cards.ArgsDownUpHcard{
		BsToken:  getBusToken(),
		HcardIDs: []int{39},
		OptType:  cards.STATUS_ON_SALE, // cards.OPT_DOWN/OPT_UP
	}
	t.Log(hcardLogic.DownUpHcard(context.Background(), args))
}

// TestShopGetBusHcardPage 子店获取适用本店的限时卡列表
func TestShopGetBusHcardPage(t *testing.T) {
	busID, shopID, start, limit := 1, 1, 0, 10
	hcardLogic := new(logics.HcardLogic)
	t.Log(hcardLogic.ShopGetBusHcardPage(context.Background(), busID, shopID, start, limit))
}

// TestShopHcardPage 子店限时卡列表
func TestShopHcardPage(t *testing.T) {
	hcardLogic := new(logics.HcardLogic)
	args := &cards.ArgsShopHcardPage{
		Status:   "",
		ShopID:   14,
		ShopCall: false, // false:总店调用此时shopID可以为0
		Paging:   common.Paging{PageSize: 10, Page: 1},
	}
	lists, _ := hcardLogic.ShopHcardPage(context.Background(), args)
	lby, _ := json.Marshal(lists)
	fmt.Println(string(lby))
}

// TestShopAddHcard 子店添加总部限时卡到自己的店铺
func TestShopAddHcard(t *testing.T) {
	args := &cards.ArgsShopAddHcard{
		//BsToken:  getBusToken(),
		HcardIDs: []int{108},
	}
	hcardLogic := new(logics.HcardLogic)
	t.Log(hcardLogic.ShopAddHcard(context.Background(), args))
}

// TestShopDownUpHcard 子店上下架自己店铺中的限时卡
func TestShopDownUpHcard(t *testing.T) {
	args := &cards.ArgsShopDownUpHcard{
		BsToken:      getBusToken(),
		ShopHcardIDs: []int{16},
		OptType:      cards.OPT_DOWN, // cards.OPT_DOWN/OPT_UP
	}
	hcardLogic := new(logics.HcardLogic)
	hcardLogic.ShopDownUpHcard(context.Background(), args)
}

//ShopHcardListRpc(ctx context.Context,args *cards.ArgsShopHcardListRpc,reply *cards.ReplyShopHcardListRpc)
func TestShopHcardListRpc(t *testing.T) {
	args := &cards.ArgsShopHcardListRpc{
		ShopId:   1,
		HcardIds: []int{16},
	}
	var reply cards.ReplyShopHcardListRpc
	if err := new(logics.HcardLogic).ShopHcardListRpc(context.Background(), args, &reply); err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("%#v", reply)
}

func TestShopDelHcard(t *testing.T) {
	bsToken := getBusToken()
	ncl := new(logics.HcardLogic)
	err := ncl.ShopDeleteHcard(context.TODO(), &cards.ArgsDeleteHcard{
		BsToken:  bsToken,
		HcardIds: []int{15, 16},
	})
	t.Log(err)
}
