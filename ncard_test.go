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

func TestModule(t *testing.T) {

	//utils.CreateModel("ncard")
	//utils.CreateModel("ncard_give")
	//utils.CreateModel("ncard_single")
	//utils.CreateModel("shop_ncard")
	//utils.CreateModel("ncard_shop")
	utils.CreateModel("ncard_ext")
}

func TestLogicAddNCard(t *testing.T) {
	bsToken := getBusToken()
	ncl := new(logics.NCardLogic)
	t.Log(ncl.AddNCard(context.TODO(), 1, &cards.ArgsAddNCard{
		BsToken: bsToken,
		NCardBase: cards.NCardBase{
			Name:          "1122-洗剪吹套餐5次111",
			ShortDesc:     "1122-hello, kc111!",
			RealPrice:     1000,
			Price:         500,
			ServicePeriod: 10,
		},
		Notes: []cards.CardNote{
			{
				Notes: "1122温馨提示1",
			},
		},
		IncludeSingles: []cards.IncSingle{
			{
				SingleID: 11,
				Num:      2,
				SspId:    25,
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

func TestLogicEditNCard(t *testing.T) {
	bsToken := getBusToken()
	ncl := new(logics.NCardLogic)
	t.Log(ncl.EditNCard(context.TODO(), 51, &cards.ArgsEditNCard{
		BsToken: bsToken,
		CardID:  118,
		NCardBase: cards.NCardBase{
			Name:          "10次美容卡e",
			ShortDesc:     "10次美容卡e",
			RealPrice:    1000,
			Price:         2000,
			ServicePeriod: 1,
		},
		Notes: []cards.CardNote{
			{
				Notes: "十点十分十三-1",
			},
			{
				Notes: "十点十分十三-2",
			},
		},
		IncludeSingles: []cards.IncSingle{
			{
				SingleID: 167,
				Num:      10, // 5--> 10
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

func TestLogicNCardInfo(t *testing.T) {
	// 获取详情
	ncl := new(logics.NCardLogic)
	if info, err := ncl.NCardInfo(context.TODO(), 6,1); err == nil {
		bytes, _ := json.Marshal(info)
		t.Log(string(bytes))
	} else {
		t.Error(err)
	}
}

func TestLogicGetBusPage(t *testing.T) {
	ncl := new(logics.NCardLogic)
	isGround := "1"
	busId, shopId := 1, 2
	if page, err := ncl.GetBusPage(context.TODO(), busId, shopId, 0, 1, isGround,true); err == nil {
		bytes,_:=json.Marshal(page)
		t.Log(string(bytes))
	} else {
		t.Error(err)
	}
}

func TestLogicSetNCardShop(t *testing.T) {
	bsToken := getBusToken()
	ncl := new(logics.NCardLogic)
	t.Log(ncl.SetNCardShop(context.TODO(), 1, &cards.ArgsSetNCardShop{
		BsToken:   bsToken,
		NCardIDs:  []int{4},
		ShopIDs:   []int{1},
		IsAllShop: false,
	}))
}

func TestLogicDownUpNCard(t *testing.T) {
	ncl := new(logics.NCardLogic)
	t.Log(ncl.DownUpNCard(context.Background(), 1, &cards.ArgsDownUpNCard{
		BsToken:  getBusToken(),
		NCardIDs: []int{1, 2},
		OptType:  cards.OPT_UP,
	}))
}

func TestLogicShopGetBusNCardPage(t *testing.T) {
	ncl := new(logics.NCardLogic)
	if list, err := ncl.ShopGetBusNCardPage(context.TODO(), 1, 1, 0, 4); err == nil {
		fmt.Printf("%#v", list)
	} else {
		t.Error(err)
	}
}

func TestLogicShopAddNCard(t *testing.T) {
	ncl := new(logics.NCardLogic)
	t.Log(ncl.ShopAddNCard(context.TODO(), 10, 19, &cards.ArgsShopAddNCard{
		BsToken:  getBusToken(),
		NCardIDs: []int{96},
	}))
}

func TestLogicShopNCardPage(t *testing.T) {
	ncl := new(logics.NCardLogic)
	if list, err := ncl.ShopNCardPage(context.TODO(), 1, 0, 6, cards.STATUS_ON_SALE); err == nil {
		jbyte, _ := json.Marshal(list)
		fmt.Println(string(jbyte))
	} else {
		t.Error(err)
	}
}

func TestLogicShopDownUpNCard(t *testing.T) {
	ncl := new(logics.NCardLogic)
	t.Log(ncl.ShopDownUpNCard(context.TODO(), 1, &cards.ArgsShopDownUpNCard{
		BsToken: getBusToken(),
		CardIDs: []int{1, 2, 3},
		OptType: cards.OPT_UP,
	}))
}
