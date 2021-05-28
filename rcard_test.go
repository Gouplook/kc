package main_test

import (
	"context"
	"encoding/json"
	"fmt"
	"git.900sui.cn/kc/kcgin/logs"
	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"testing"
)

func TestAddRcard(t *testing.T) {

	bsToken := getBusToken()

	addLogic := new(logics.RcardLogic)
	args := cards.ArgsAddRcard{
		BsToken: bsToken,
		RcardBase: cards.RcardBase{
			BusID:         51,
			Name:          "51充值卡",
			SortDesc:      "51充值卡",
			RealPrice:     2000,
			Price:         0,
			DiscountType:  1, //
			Discount:      1.5,
			ServicePeriod: 3, // 3月
			SaleShopNum:   2,
		},
		Notes: []cards.CardNote{
			{
				Notes: "提示1",
			},
		},
		IsAllSingle: true,
		IncludeSingles: []cards.IncSingle{
			{
				SingleID: 140,
				Num:      3,
			},
		},
		GiveSingles: []cards.IncSingle{
			{
				SingleID: 140,
				Num:      2,
				PeriodOfValidity:50, // 30--50
			},
		},
		// IncludeProducts: []cards.IncProduct{
		// 	{
		// 		ProductID: 0,
		// 	},
		// },
		ImgHash: "b82275a6c75a43fd836d25be534932ad",
	}
	busId := 1
	rcardId, err := addLogic.AddRcard(context.Background(), busId, &args)
	if err != nil {
		fmt.Println("err---------", err)
	}
	fmt.Println(rcardId)

}

func TestEditRcard(t *testing.T) {
	bsToken := getBusToken()

	editLogic := new(logics.RcardLogic)
	args := cards.ArgsEditRcard{
		BsToken: bsToken,
		RcardId: 199,
		RcardBase: cards.RcardBase{
			BusID:         51,
			Name:          "万能充值卡e",
			SortDesc:      "万能充值卡e",
			RealPrice:     1000,
			Price:         1000,
			DiscountType:  1, // 需要讨论
			Discount:      10,  //
			ServicePeriod: 3, // 36 --> 3
			SaleShopNum:   1,
		},
		Notes: []cards.CardNote{
			{
				Notes: "万能充值卡,无限使用",
			},
		},
		IsAllSingle: true,
		// IncludeSingles: []cards.IncSingle{
		// 	{
		// 		SingleID: 15,
		// 		Num:      3,
		// 	},
		// },
		// GiveSingles: []cards.IncSingle{
		// 	{
		// 		SingleID: 12,
		// 		Num:      3,
		// 	},
		// },
		IsAllProduct: false,
		// IncludeProducts: []cards.IncProduct{
		// 	{
		// 		ProductID: 15,
		// 	},
		// },
		ImgHash: "b82275a6c75a43fd836d25be534932ad",
	}
	busId := 51
	err := editLogic.EditRcard(context.Background(), busId, &args)
	if err != nil {
		fmt.Println("err---------", err)
	}
	t.Log(111)

}

func TestAddRechargeRules(t *testing.T) {

	rulesLogic := new(logics.RcardLogic)
	args := cards.ArgsAddRechargeRules{
		RcardId: 2,
		ListsRechargeRules: cards.ListsRechargeRules{
			RechargeAmount: 500,
			DonationAmount: 100,
		},
	}
	rcardId, err := rulesLogic.AddRechargeRules(context.Background(), &args)
	if err != nil {
		fmt.Println("----------:", err)
		return
	}
	fmt.Println(rcardId)

}
func TestEditRules(t *testing.T) {
	bsToken := getBusToken()
	editLogic := new(logics.RcardLogic)
	args := &cards.ArgsEditRechargeRules{
		BsToken: bsToken,
		RcardId: 20,
		ListsRechargeRules: cards.ListsRechargeRules{
			RechargeAmount: 888,
			DonationAmount: 77,
		},
	}
	err := editLogic.EditRechargeRules(context.Background(), args)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, _ := json.Marshal(args)

	fmt.Println(string(data))
}

func TestRechargeRulesInfo(t *testing.T) {
	infoLogic := new(logics.RcardLogic)
	args := cards.ArgsRechargeRulesInfo{
		Id: 1,
	}
	reply := cards.ReplyRechargeRulesInfo{}
	err := infoLogic.RechargeRulesInfo(context.Background(), &args, &reply)
	if err != nil {
		fmt.Println("----------:", err)
		return
	}
	fmt.Println(reply.DonationAmount)
	fmt.Println(reply.RechargeAmount)

}

//获取总店的充值卡列表
func TestBusRcardPage(t *testing.T) {
	args := &cards.ArgsBusRcardPage{
		Paging:           common.Paging{Page: 1, PageSize: 10},
		BusId:            1,
		ShopId:           2,
		IsGround:         "",
		IsDel:            0,
		FilterShopHasAdd: true,
	}
	reply, err := new(logics.RcardLogic).BusRcardPage(context.Background(), args)
	if err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//func (r *Rcard) BusRcardPage(ctx context.Context, args *cards.ArgsBusRcardPage, reply *cards.ReplyRcardPage) (err error) {

func TestRcardInfo(t *testing.T) {
	args := &cards.ArgsRcardInfo{
		RcardId: 98,
		ShopId:  0,
	}
	reply, err := new(logics.RcardLogic).RcardInfo(context.Background(), args)
	if err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

func TestShopAddRcard(t *testing.T) {
	shopId, busId, rcardIds := 1, 1, []int{116}
	if err := new(logics.RcardLogic).ShopAddRcard(context.Background(), shopId, busId, rcardIds); err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(1)
}

func TestItems(t *testing.T) {
	args := cards.ArgsGetItemsByShopId{
		Paging: common.Paging{
			Page:     1,
			PageSize: 3,
		},
		ItemType: 1,
		ShopId:   429,
		OrderBy:  "",
	}

	reply := cards.ReplyGetItemsByShopId{}
	err := new(logics.ItemLogic).GetItemsByShopId(context.Background(), &args, &reply)

	logs.Info(err)
	logs.Info(reply)
}
