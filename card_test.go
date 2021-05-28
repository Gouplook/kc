package main_test

import (
	"context"
	"encoding/json"
	"fmt"
	"git.900sui.cn/kc/base/utils"
	"git.900sui.cn/kc/kcgin/logs"
	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcCards/service"
	cards2 "git.900sui.cn/kc/rpcinterface/client/cards"
	"git.900sui.cn/kc/rpcinterface/client/product"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"git.900sui.cn/kc/rpcinterface/interface/order"
	product2 "git.900sui.cn/kc/rpcinterface/interface/product"
	"strings"
	"testing"
)

func TestCardModule(t *testing.T) {
	utils.CreateModel("shop_item_relation")
}

func TestLogicAddCard(t *testing.T) {
	bsToken := getBusToken()
	ncl := new(logics.CardLogic)
	cardId, err := ncl.AddCard(context.TODO(), 1, &cards.ArgsAddCard{
		BsToken: bsToken,
		CardBase: cards.CardBase{
			Name:          "新改版综合卡2",
			ShortDesc:     "hello, kc!",
			RealPrice:     100,
			Price:         150,
			ServicePeriod: 3,
		},
		Notes: []cards.CardNote{
			{
				Notes: "综合卡温馨提示1111111111111111111111111111111111111111111111111111111111111111",
			},
		},
		IsAllSingle: false,
		IncludeSingles: []cards.IncInfSingle{
			{
				SingleID: 11,
			},
		},
		GiveSingles: []cards.IncSingle{
			{
				SingleID: 12,
				Num:      3,
			},
		},
		IsAllProduct: false,
		IncludeProducts: []cards.IncProduct{
			{
				ProductID: 15,
			},
			{
				ProductID: 16,
			},
		},
		ImgHash: "b82275a6c75a43fd836d25be534932ad",
	})
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log("cardId=", cardId)
}

func TestLogicEditCard(t *testing.T) {
	bsToken := getBusToken()
	ncl := new(logics.CardLogic)
	err := ncl.EditCard(context.TODO(), 51, &cards.ArgsEditCard{
		BsToken: bsToken,
		CardID:  169,
		CardBase: cards.CardBase{
			Name:          "万能卡e",
			ShortDesc:     "万能卡",
			RealPrice:     0.1,
			Price:         500,
			ServicePeriod: 0,
		},
		Notes: []cards.CardNote{
			{
				Notes: "万能卡e温馨提示",
			},
		},
		IsAllSingle: true,
		// IncludeSingles: []cards.IncInfSingle{
		// 	{
		// 		SingleID: 12,
		// 	},
		// },
		// GiveSingles: []cards.IncSingle{
		// 	{
		// 		SingleID: 11,
		// 		Num:      5,
		// 	},
		// },
		IsAllProduct: false,
		// IncludeProducts: []cards.IncProduct{
		// 	{
		// 		ProductID: 15,
		// 	},
		// 	{
		// 		ProductID: 16,
		// 	},
		// },
		ImgHash: "b82275a6c75a43fd836d25be534932ad",
	})
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(1)
}

func TestLogicCardInfo(t *testing.T) {
	ncl := new(logics.CardLogic)
	if info, err := ncl.CardInfo(context.TODO(), 161); err == nil {
		bytes, _ := json.Marshal(info)
		t.Log(string(bytes))
	} else {
		t.Error(err)
	}
}

func TestGetBusPage(t *testing.T) {
	ncl := new(logics.CardLogic)
	isGround := ""
	busId, shopId := 1, 2
	if page, err := ncl.GetBusPage(context.TODO(), busId, shopId, 0, 10, isGround, true); err == nil {
		bytes, _ := json.Marshal(page)
		t.Log(string(bytes))
	} else {
		t.Error(err)
	}
}

func TestLogicSetCardShop(t *testing.T) {
	//bsToken := getBusToken()
	ncl := new(logics.CardLogic)
	t.Log(ncl.SetCardShop(context.Background(), 1, &cards.ArgsSetCardShop{
		//BsToken:   bsToken,
		CardIDs:   []int{29}, // 下架的数据
		ShopIDs:   []int{1},
		IsAllShop: false,
	}))
}

func TestLogicDownUpCard(t *testing.T) {
	ncl := new(logics.CardLogic)
	t.Log(ncl.DownUpCard(context.Background(), 1, &cards.ArgsDownUpCard{
		BsToken: getBusToken(),
		CardIDs: []int{1},
		OptType: cards.OPT_UP,
	}))
}

func TestLogicShopGetBusCardPage(t *testing.T) {
	ncl := new(logics.CardLogic)
	if list, err := ncl.ShopGetBusCardPage(context.TODO(), 1, 2, 0, 20); err == nil {
		bytes, _ := json.Marshal(list)

		fmt.Printf(string(bytes))
	} else {
		t.Error(err)
	}
}
func TestServiceShopGetBusCardPage(t *testing.T) {
	args := &cards.ArgsShopGetBusCardPage{
		Paging:  common.Paging{1, 10},
		BsToken: getBusToken(),
	}
	reply := &cards.ReplyCardPage{}
	sc := new(service.Card)
	if err := sc.ShopGetBusCardPage(context.TODO(), args, reply); err == nil {
		fmt.Printf("%#v", args)
	} else {
		t.Error(err)
	}
}

func TestLogicShopAddCard(t *testing.T) {
	ncl := new(logics.CardLogic)
	t.Log(ncl.ShopAddCard(context.TODO(), 1, 1, &cards.ArgsShopAddCard{
		BsToken: getBusToken(),
		CardIDs: []int{56},
	}))
}

func TestLogicShopCardPage(t *testing.T) {
	ncl := new(logics.CardLogic)
	list, err := ncl.ShopCardPage(context.Background(), 414, 0, 20, 2)
	if err != nil {
		logs.Info("err:", err)
	}
	logs.Info("list:", list)
}

func TestLogicShopDownUpCard(t *testing.T) {
	ncl := new(logics.CardLogic)
	t.Log(ncl.ShopDownUpCard(context.TODO(), 1, &cards.ArgsShopDownUpCard{
		BsToken: getBusToken(),
		CardIDs: []int{1},
		OptType: cards.OPT_UP,
	}))
}

func TestProduct(t *testing.T) {
	client := new(product.Product).Init()
	defer client.Close()
	reply := new([]product2.ReplyProductGetByIds)
	param := new(product2.ArgsProductGetByIds)
	param.Ids = []int{28, 29, 30}
	if err := client.GetProductByIds(context.TODO(), param, reply); err != nil {
		t.Error(err)
	} else {
		t.Log(reply)
	}
}

func TestShopCardListRpc(t *testing.T) {
	args := &cards.ArgsShopCardListRpc{
		ShopId:  1,
		CardIds: []int{1, 2, 3},
	}
	var reply cards.ReplyShopCardListRpc
	if err := new(logics.CardLogic).ShopCardListRpc(context.Background(), args, &reply); err != nil {
		return
	}
	t.Logf("%#v", reply)
}

func TestGetInfos(t *testing.T) {
	reply, err := new(logics.ItemLogic).GetInfos(context.Background(), &cards.ArgsAppInfos{
		//Id:  2,
		Cid:    321,
		Paging: common.Paging{Page: 0, PageSize: 1},
	})
	if err != nil {
		t.Fatalf("出错了 err = %#v", err)
	}
	bytes, _ := json.Marshal(reply)
	shopIds := []int{}
	for _, list := range reply.Lists {
		shopIds = append(shopIds, list.ShopId)
	}
	ss := "(" + strings.Replace(strings.Trim(fmt.Sprint(shopIds), "[]"), " ", ",", -1) + ")"
	t.Log(ss)
	t.Log(string(bytes))
}

func TestGetDetailById(t *testing.T) {
	var reply []order.ReplyCableShopInfo
	err := new(logics.ItemLogic).GetDetailById(context.Background(), &cards.ArgsShopList{
		Lng: 121.514, Lat: 31.235, ItemType: 1, ItemId: 11,
	}, &reply)
	if err != nil {
		t.Fatalf("出错了 err = %v", err)
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//GetItemsByShopId(ctx context.Context,args *cards.ArgsGetItemsByShopId,reply *cards.ReplyGetItemsByShopId)
func TestGetItemsByShopId(t *testing.T) {
	typs := []int{4}
	for _, v := range typs {
		args := &cards.ArgsGetItemsByShopId{
			Paging: common.Paging{
				Page:     1,
				PageSize: 10,
			},
			ItemType: v,
			ShopId:   429,
		}
		var reply cards.ReplyGetItemsByShopId
		if err := new(logics.ItemLogic).GetItemsByShopId(context.Background(), args, &reply); err != nil {
			return
		}
		bytes, _ := json.Marshal(reply)
		t.Logf(string(bytes))
	}
}

//GetRecommendSingles(ctx context.Context,args *cards.ArgsGetRecommendSingles,reply *cards.ReplyGetRecommendSingles)
func TestGetRecommendSingles(t *testing.T) {
	args := &cards.ArgsGetRecommendSingles{
		TotalNum: 10, ShopId: 1,
	}
	var reply cards.ReplyGetRecommendSingles
	if err := new(service.Item).GetRecommendSingles(context.Background(), args, &reply); err != nil {
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Logf(string(bytes))
}

func TestAS(t *testing.T) {
	var reply cards.ReplySubServer
	if err := new(logics.SingleLogic).GetSubServerBySingleId(&cards.ArgsSubServer{
		SingleId: 12,
	}, &reply); err != nil {
		t.Fatalf("出错了 err = %v", err)
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

func TestAAA(t *testing.T) {
	var reply map[int][]cards.SubSpec
	err := new(logics.SpecLogic).GetBySspIds(&[]int{25, 26, 27}, &reply)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reply)
}

//CollectItems(ctx context.Context, args *cards.ArgsCollectItems, reply *bool)
func TestCollectItems(t *testing.T) {
	args := &cards.ArgsCollectItems{
		Uid:      7,
		ItemType: cards.ITEM_TYPE_hncard,
		ItemId:   23,
		SsId:     12,
	}
	var reply bool
	if err := new(logics.ItemLogic).CollectItems(context.Background(), args, &reply); err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(reply)
}

//GetCollectItems(ctx context.Context, args *cards.ArgsGetCollectItems, reply *cards.ReplyGetCollectItems)
func TestGetCollectItems(t *testing.T) {
	args := &cards.ArgsGetCollectItems{
		Paging: common.Paging{Page: 1, PageSize: 7},
		Uid:    54, ItemType: cards.ITEM_TYPE_single,
	}
	var reply cards.ReplyGetCollectItems
	if err := new(logics.ItemLogic).GetCollectItems(context.Background(), args, &reply); err != nil {
		t.Error(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//GetCollectStatus(ctx context.Context, args *cards.ArgsCollectStatus, reply *cards.ReplyCollectStatus) (err error)
func TestGetCollectStatus(t *testing.T) {
	args := &cards.ArgsCollectStatus{
		Uid: 7, ItemId: 3, ItemType: 2, SsId: 3,
	}
	var reply cards.ReplyCollectStatus
	if err := new(logics.ItemLogic).GetCollectStatus(context.Background(), args, &reply); err != nil {
		t.Error(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//GetRcardBaseInfo(ctx context.Context,args *cards.ArgsGetRcardBaseInfo,reply *cards.ReplyGetRcardBaseInfo)
func TestGetRcardBaseInfo(t *testing.T) {
	args := &cards.ArgsGetRcardBaseInfo{RcardIds: []int{98}, RuleId: 0}
	var reply cards.ReplyGetRcardBaseInfo
	if err := new(logics.RcardLogic).GetRcardBaseInfo(context.Background(), args, &reply); err != nil {
		t.Error(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

func TestGetItemsBySsidsTypeSm(t *testing.T) {
	//args *cards.ArgsGetItemsBySsids
	args := &cards.ArgsGetItemsBySsids{
		SsIds:    []int{44},
		ItemType: cards.ITEM_TYPE_rcard,
	}
	itemLogicnew := new(logics.ItemLogic)
	reply, err := itemLogicnew.GetItemsBySsids(context.Background(), args)
	if err != nil {
		fmt.Printf("err :%#v\n", err)
		return
	}
	js, _ := json.Marshal(reply)
	jsStr := string(js)
	t.Log(jsStr)

}

func TestGetItemsBySsidsTypeCard(t *testing.T) {
	//args *cards.ArgsGetItemsBySsids
	args := &cards.ArgsGetItemsBySsids{
		SsIds:    []int{22, 23},
		ItemType: cards.ITEM_TYPE_card,
	}
	itemLogicnew := new(logics.ItemLogic)
	reply, err := itemLogicnew.GetItemsBySsids(context.Background(), args)
	if err != nil {
		fmt.Printf("err :%#v\n", err)
		return
	}
	js, _ := json.Marshal(reply)
	jsStr := string(js)
	fmt.Printf("jsStr:%#v\n", jsStr)
}

func TestGetItemsBySsidsTypeHcard(t *testing.T) {
	//args *cards.ArgsGetItemsBySsids
	args := &cards.ArgsGetItemsBySsids{
		SsIds:    []int{20, 21},
		ItemType: cards.ITEM_TYPE_hcard,
	}
	itemLogicnew := new(logics.ItemLogic)
	reply, err := itemLogicnew.GetItemsBySsids(context.Background(), args)
	if err != nil {
		fmt.Printf("err :%#v\n", err)
		return
	}
	js, _ := json.Marshal(reply)
	jsStr := string(js)
	fmt.Printf("jsStr:%#v\n", jsStr)
}
func TestGetItemsBySsidsTypeNcard(t *testing.T) {
	//args *cards.ArgsGetItemsBySsids
	args := &cards.ArgsGetItemsBySsids{
		SsIds:    []int{25, 26, 27},
		ItemType: cards.ITEM_TYPE_ncard,
	}
	itemLogicnew := new(logics.ItemLogic)
	reply, err := itemLogicnew.GetItemsBySsids(context.Background(), args)
	if err != nil {
		fmt.Printf("err :%#v\n", err)
		return
	}
	js, _ := json.Marshal(reply)
	jsStr := string(js)
	fmt.Printf("jsStr:%#v\n", jsStr)
}

func TestGetItemsBySsidsTypeHncard(t *testing.T) {
	//args *cards.ArgsGetItemsBySsids
	args := &cards.ArgsGetItemsBySsids{
		SsIds:    []int{25, 26, 27, 28, 29},
		ItemType: cards.ITEM_TYPE_hncard,
	}
	itemLogicnew := new(logics.ItemLogic)
	reply, err := itemLogicnew.GetItemsBySsids(context.Background(), args)
	if err != nil {
		fmt.Printf("err :%#v\n", err)
		return
	}
	js, _ := json.Marshal(reply)
	jsStr := string(js)
	fmt.Printf("jsStr:%#v\n", jsStr)
}
func TestGetItemsBySsidsTypeIcard(t *testing.T) {
	//args *cards.ArgsGetItemsBySsids
	args := &cards.ArgsGetItemsBySsids{
		SsIds:    []int{3, 4},
		ItemType: cards.ITEM_TYPE_icard,
	}
	itemLogicnew := new(logics.ItemLogic)
	reply, err := itemLogicnew.GetItemsBySsids(context.Background(), args)
	if err != nil {
		fmt.Printf("err :%#v\n", err)
		return
	}
	js, _ := json.Marshal(reply)
	jsStr := string(js)
	fmt.Printf("jsStr:%#v\n", jsStr)
}
func TestGetItemsBySsidsTypeRcard(t *testing.T) {
	//args *cards.ArgsGetItemsBySsids
	args := &cards.ArgsGetItemsBySsids{
		SsIds:    []int{9, 12, 13, 14},
		ItemType: cards.ITEM_TYPE_rcard,
	}
	itemLogicnew := new(logics.ItemLogic)
	reply, err := itemLogicnew.GetItemsBySsids(context.Background(), args)
	if err != nil {
		fmt.Printf("err :%#v\n", err)
		return
	}
	js, _ := json.Marshal(reply)
	jsStr := string(js)
	fmt.Printf("jsStr:%#v\n", jsStr)
}

func TestGetItemAllXCradsNumByShopId(t *testing.T) {
	shopId := 2
	var reply cards.ReplyGetItemAllXCradsNumByShopId
	rpcCards := new(cards2.Item).Init()
	err := rpcCards.GetItemAllXCradsNumByShopId(context.Background(), &shopId, &reply)
	if err != nil {
		t.Log(err.Error())
		return
	}
	by, _ := json.Marshal(reply)
	t.Log(string(by))
}

func TestGetBusBaseInfoRpc(t *testing.T) {
	logic := new(logics.ICardLogic)

	reply, err := logic.GetBusBaseInfoRpc(45)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(reply)

}

func TestGetSelectServices(t *testing.T) {
	var reply []cards.ReplyGetSelectServices
	err := new(logics.SingleLogic).GetSelectServices(context.Background(), &cards.ArgsGetSelectServices{
		Cid: 321, Num: 3,
	}, &reply)
	if err != nil {
		t.Fatal(err)
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//包含的单项目
type IncSingle struct {
	SingleID int `mapstructure:"single_id" json:"SingleId"` //单项目id
	Num      int `mapstructure:"num"`                       //单项目次数
	SspId    int `mapstructure:"ssp_id"`                    //规格id
}

/*
[{“SingleID”:12,“Num”:3 },{“SingleID”:14,“Num”:3}]
*/
func TestAddXCardTask(t *testing.T) {
	insingleStr := "[{\"SingleID\":12,\"Num\":3 },{\"SingleID\":14,\"Num\":3}]"
	var is []IncSingle
	var err error
	if err = json.Unmarshal([]byte(insingleStr), &is); err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(is)
	isAllSingle := 1 == 0
	t.Log(isAllSingle)

	//new(logics.ItemLogic).AddXCardTask(context.Background(),11,1)
}

/*func TestAAAA(t *testing.T) {
	reply, err := new(logics.ItemLogic).GetSms([]int{29, 30})
	if err != nil {
		t.Fatal(err)
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}*/

//func (i *ItemLogic)GetCardProductsInfo(ctx context.Context,args *cards.ArgsGetCardProductsInfo,reply *cards.ReplyGetCardProductsInfo)(err error){
func TestGetCardProductsInfo(t *testing.T) {
	//获取部分
	args := &cards.ArgsGetCardProductsSinglesInfo{
		//ItemId: 70,  //card
		ItemId: 268, //rcard
		//ItemType: cards.ITEM_TYPE_card,
		ItemType: cards.ITEM_TYPE_icard,
		ShopId:   0,
		Paging: common.Paging{
			Page:     1,
			PageSize: 10,
		},
	}
	////获取全部
	//args := &cards.ArgsGetCardProductsInfo{
	//	ItemId: 113,  //rcard
	//	ItemType: cards.ITEM_TYPE_rcard,
	//	ShopId: 1,
	//	Paging :common.Paging{
	//		Page: 0,
	//		PageSize: 10,
	//	},
	//}

	var reply cards.ReplyGetCardProductsInfo
	err := new(logics.ItemLogic).GetCardProductsInfo(context.Background(), args, &reply)
	if err != nil {
		t.Fatal(err)
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//GetItemIncludeSingles(ctx context.Context, args *cards.ArgsGetItemIncludeSingles, reply *cards.ReplyGetItemIncludeSingles)
func TestGetItemIncludeSingles(t *testing.T) {
	args := &cards.ArgsGetCardProductsSinglesInfo{
		Paging:   common.Paging{Page: 1, PageSize: 10},
		ItemId:   202,
		ItemType: 2,
		ShopId:   19,
	}
	var reply cards.ReplyGetItemIncludeSingles
	if err := new(service.Item).GetItemIncludeSingles(context.Background(), args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//Info(ctx context.Context, args cards.InputParamsICardInfo, reply *cards.OutputParamsICardInfo)
func TestIcardInfo(t *testing.T) {
	args := cards.InputParamsICardInfo{
		BsToken: common.BsToken{},
		IcardID: 191,
		ShopID:  0,
	}
	var reply cards.OutputParamsICardInfo
	if err := new(service.ICard).Info(context.Background(), args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//func (i *ItemLogic) GetItemIncProductsId(ctx context.Context,ItemType int,ItemIds []int)(err error, isAll bool,productIds []int){
func TestGetItemIncProductsId(t *testing.T) {
	//err,isAll,productIds := new(logics.ItemLogic).GetItemIncProductsId(context.Background(),cards.ITEM_TYPE_card,[]int{1,17,19})
	//err,isAll,productIds := new(logics.ItemLogic).GetItemIncProductsId(context.Background(),cards.ITEM_TYPE_rcard,[]int{100,102})
	err, isAll, productIds := new(logics.ItemLogic).GetItemIncProductsId(context.Background(), cards.ITEM_TYPE_icard, []int{2, 3})
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(isAll)
	t.Log(productIds)
}

func TestCheckProductInShop(t *testing.T) {
	err, reply := new(logics.ItemLogic).CheckProductInShop(context.Background(), 1, 1, cards.ITEM_TYPE_rcard, []int{100, 102})
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(reply)
}

func TestGetItemDefaultImgs(t *testing.T) {
	args := &cards.ArgsGetItemDefaultImgs{ImgType: 2, ItemType: 1}
	var reply cards.ReplyGetItemDefaultImgs
	rpc := new(cards2.Item).Init()
	if err := rpc.GetItemDefaultImgs(context.Background(), args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//获取身份卡的折扣信息
func TestGetIcardDiscountById(t *testing.T) {
	ICardId := 59
	var reply cards.ReplyGetIcardDiscountById
	if err := new(logics.ICardLogic).GetICardDiscountById(context.Background(), ICardId, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//Delete 删除身份卡
func TestDelete(t *testing.T) {
	busId, shopId, icardIdsStr := 1, 0, "[285]"
	res, err := new(logics.ICardLogic).Delete(context.Background(), busId, shopId, icardIdsStr)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(res)
}

//Save 总店添加/编辑身份卡
func TestSaveICard(t *testing.T) {
	busId := 1
	args := cards.InputParamsICardSave{
		BsToken: common.BsToken{},
		ICardBase: cards.ICardBase{
			Name:           "身份卡一期优化测试",
			BusID:          1,
			ShortDesc:      "身份卡一期优化测试折扣变更",
			RealPrice:      0.01,
			Price:          100,
			ServicePeriod:  12,
			SaleShopNum:    0,
			ImgID:          0,
			Sales:          0,
			Ctime:          "",
			Click:          0,
			SaleShopCount:  0,
			HasGiveSignle:  0,
			IsGround:       0,
			IcardID:        289,
			Status:         0,
			IsSelfShop:     0,
			SsID:           0,
			DiscountSingle: 0,
			DiscountGoods:  0,
			IsAllSingle:    false,
			IsAllProduct:   false,
			ShopItemId:     0,
		},
		Notes:    "[{\"Notes\":\"身份卡一期优化测试提示\"}]",
		IsGround: 0,
		//IncludeSingles:  "[{\"singleID\":0,\"discount\":5.5}]",
		//IncludeProducts: "[{\"productID\":0,\"discount\":5.5}]",
		IncludeSingles:  "[{\"singleID\":3,\"discount\":2.5},{\"singleID\":4,\"discount\":2.5}]",
		IncludeProducts: "[{\"productID\":3,\"discount\":2.5},{\"productID\":4,\"discount\":2.6}]",
		GiveSingles:     "[{\"singleID\":3,\"num\":1},{\"singleID\":4,\"num\":1}]",
		ImgHash:         "02457676043548aa8f8f7b274b2e4d41",
		GiveSingleDesc:  "身份卡一期优化测试赠送描述",
		IsSync:          cards.IS_SYNC_NO,
	}
	var reply cards.OutputParamsICardSave
	if err := new(logics.ICardLogic).Save(context.Background(), busId, args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//获取身份卡备份表中的项目折扣
func TestGetICardSingleDiscount(t *testing.T) {
	args := &cards.ArgsGetICardSingleDiscount{
		ICardId: 265,
		IsSync:  fmt.Sprint(cards.IS_SYNC_NO),
	}
	var reply cards.ReplyGetICardSingleDiscount
	if err := new(logics.ICardLogic).GetICardSingleDiscount(context.Background(), args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

func TestGetApplyAndGiveSingleNum(t *testing.T) {
	itemCardIds, itemType := []int{122, 128}, cards.ITEM_TYPE_rcard
	res := logics.GetApplyAndGiveSingleNum(itemCardIds, itemType)
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

//预付卡赠送的单项目
func TestGetItemGiveSingles(t *testing.T) {
	args := &cards.ArgsGetCardProductsSinglesInfo{
		Paging:   common.Paging{Page: 1, PageSize: 10},
		ItemId:   179,
		ItemType: cards.ITEM_TYPE_sm,
		ShopId:   399,
		IsDel: "0",
		ShopStatus: "2",

	}
	var reply cards.ReplyGetItemGiveSingles
	if err := new(logics.ItemLogic).GetItemGiveSingles(context.Background(), args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}
