package main_test

import (
	"context"
	"encoding/json"
	"fmt"
	"git.900sui.cn/kc/base/utils"
	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"testing"
)

func TestHHNCardModule(t *testing.T) {

	//utils.CreateModel("hncard")
	//utils.CreateModel("hncard_give")
	//utils.CreateModel("hncard_single")
	//utils.CreateModel("shop_hncard")
	//utils.CreateModel("hncard_shop")
	utils.CreateModel("hncard_ext")
}

func TestLogicAddHNCard(t *testing.T) {
	bsToken := getBusToken()
	ncl := new(logics.HNCardLogic)
	t.Log(ncl.AddHNCard(context.TODO(), 1, &cards.ArgsAddHNCard{
		BsToken: bsToken,
		HNCardBase: cards.HNCardBase{
			Name:          "洗剪吹套餐5次ggggg",
			ShortDesc:     "hello, kc!",
			RealPrice:     100,
			Price:         150,
			ServicePeriod: 3,
		},
		Notes: []cards.CardNote{
			{
				Notes: "xxxx",
			},
		},
		IncludeSingles: []cards.IncSingle{
			{
				SingleID: 11,
				Num:      2,
			},
		},
		GiveSingles: []cards.IncSingle{
			{
				SingleID: 12,
				Num:      3,
			},
		},
		ImgHash: "b82275a6c75a43fd836d25be534932ad",
	}))
}

func TestLogicEditHNCard(t *testing.T) {
	bsToken := getBusToken()
	ncl := new(logics.HNCardLogic)
	t.Log(ncl.EditHNCard(context.TODO(), 1, &cards.ArgsEditHNCard{
		BsToken: bsToken,
		CardID:  5,
		HNCardBase: cards.HNCardBase{
			Name:          "洗剪吹套餐6次",
			ShortDesc:     "hello, kc!",
			RealPrice:     200,
			Price:         300,
			ServicePeriod: 10,
		},
		Notes: []cards.CardNote{
			{
				Notes: "提示1",
			},
		},
		IncludeSingles: []cards.IncSingle{
			{
				SingleID: 1,
				Num:      6,
			},
		},
		// GiveSingles: []cards.IncSingle{
		// 	{
		// 		SingleID: 11,
		// 		Num:      5,
		// 	},
		// },
		ImgHash: "b82275a6c75a43fd836d25be534932ad",
	}))
}

func TestLogicHNCardInfo(t *testing.T) {
	ncl := new(logics.HNCardLogic)
	if info, err := ncl.HNCardInfo(context.TODO(), 107,399); err == nil {
		bytes, _ := json.Marshal(info)
		t.Log(string(bytes))
	} else {
		t.Error(err)
	}
}

func TestLogicHNCardGetBusPage(t *testing.T) {
	ncl := new(logics.HNCardLogic)
	isGround := ""

	busId, shopId := 1, 2
	if page, err := ncl.GetBusPage(context.TODO(), busId, shopId, 0, 10, isGround,true); err == nil {
		bytes, _ := json.Marshal(page)
		t.Log(string(bytes))
	} else {
		t.Error(err)
	}
}

func TestLogicSetHNCardShop(t *testing.T) {
	//bsToken := getBusToken()
	ncl := new(logics.HNCardLogic)
	t.Log(ncl.SetHNCardShop(context.TODO(), 1, &cards.ArgsSetHNCardShop{
		//BsToken:   bsToken,
		HNCardIDs: []int{1, 2},
		ShopIDs:   []int{5},
		IsAllShop: false,
	}))
}

func TestLogicDownUpHNCard(t *testing.T) {
	ncl := new(logics.HNCardLogic)
	t.Log(ncl.DownUpHNCard(context.Background(), 1, &cards.ArgsDownUpHNCard{
		BsToken:   getBusToken(),
		HNCardIDs: []int{1, 2},
		OptType:   cards.OPT_UP,
	}))
}

func TestLogicShopGetBusHNCardPage(t *testing.T) {
	ncl := new(logics.HNCardLogic)
	if list, err := ncl.ShopGetBusHNCardPage(context.TODO(), 1, 1, 0, 100); err == nil {
		fmt.Printf("%#v", list)
	} else {
		t.Error(err)
	}
}

func TestLogicShopAddHNCard(t *testing.T) {
	ncl := new(logics.HNCardLogic)
	t.Log(ncl.ShopAddHNCard(context.TODO(), 10, 19, &cards.ArgsShopAddHNCard{
		BsToken:   getBusToken(),
		HNCardIDs: []int{126},
	}))
}

func TestLogicShopHNCardPage(t *testing.T) {
	ncl := new(logics.HNCardLogic)
	if list, err := ncl.ShopHNCardPage(context.TODO(), 399, 0, 20, 2); err == nil {
		jbyte, _ := json.Marshal(list)
		fmt.Println(string(jbyte))
	} else {
		t.Error(err)
	}
}

func TestLogicShopDownUpHNCard(t *testing.T) {
	ncl := new(logics.HNCardLogic)
	t.Log(ncl.ShopDownUpHNCard(context.TODO(), 399, &cards.ArgsShopDownUpHNCard{
		BsToken: getBusToken(),
		CardIDs: []int{114},
		OptType: 2, //1-下架；2-上架
	}))
}

//ShopHncardRpc(ctx context.Context,args *cards.ArgsShopHncardRpc,list *cards.ReplyShopHncardRpc)
func TestShopHncardRpc(t *testing.T) {
	args := &cards.ArgsShopHncardRpc{
		ShopId:    1,
		HNCardIds: []int{1, 2, 4},
	}
	var reply cards.ReplyShopHncardRpc
	if err := new(logics.HNCardLogic).ShopHncardRpc(context.Background(), args, &reply); err != nil {
		return
	}
	t.Logf("%#v", reply)
}

func TestDelteHncard(t *testing.T) {
	//bsToken := getBusToken()
	ncl := new(logics.HNCardLogic)
	err := ncl.DeleteHNCard(context.TODO(), &cards.ArgsDelHNCard{
		//BsToken:bsToken,
		HNCardIds: []int{1, 2},
	})
	t.Log(err)

}

func TestShopDelHncard(t *testing.T) {
	bsToken := getBusToken()
	ncl := new(logics.HNCardLogic)
	err := ncl.ShopDeleteHNCard(context.TODO(), &cards.ArgsDelHNCard{
		BsToken:   bsToken,
		HNCardIds: []int{1, 2},
	})
	t.Log(err)
}
