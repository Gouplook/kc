//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/9 10:02
package main_test

import (
	"context"
	"encoding/json"
	"fmt"
	"git.900sui.cn/kc/kcgin/logs"
	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcCards/service"
	"git.900sui.cn/kc/rpcinterface/client/bus"
	cards2 "git.900sui.cn/kc/rpcinterface/client/cards"
	"git.900sui.cn/kc/rpcinterface/client/user"
	bus2 "git.900sui.cn/kc/rpcinterface/interface/bus"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	user2 "git.900sui.cn/kc/rpcinterface/interface/user"
	"testing"
)

func TestGetSingle(t *testing.T) {
	mSingle := new(service.Single)
	args := &[]int{11}
	reply := []cards.ReplyCommonSingle{}
	if err := mSingle.GetBySingle(context.Background(), args, &reply); err != nil {
		logs.Info("err:", err)
		return
	}
	logs.Info("reply:", reply)
	replyJson, _ := json.Marshal(reply)
	logs.Info("replyJson:", string(replyJson))
}

func TestAddSingle(t *testing.T) {
	ctx := context.Background()
	rpcUser := new(user.UserLogin).Init()
	loginParams := user2.CheckLoginParams{
		Channel: 1,
		Token:   "YzJpYjczYWhnZGMyMmlwZWNhYjAxNjIxNDA2NjA1OTUyMjAwMDIy",
	}
	loginRep := user2.CheckLoginReply{}

	rpcUser.CheckLogin(ctx, &loginParams, &loginRep)
	rpcBus := new(bus.BusAuth).Init()
	defer rpcBus.Close()
	utoken := common.Utoken{UidEncodeStr: loginRep.UidEncodeStr}
	args := bus2.ArgsBusAuth{
		BusId:  1,
		Utoken: utoken,
	}
	reply := bus2.ReplyBusAuth{}
	rpcBus.BusAuth(ctx, &args, &reply)
	bstoken := common.BsToken{EncodeStr: reply.EncodeStr}

	mSingle := new(logics.SingleLogic)
	singleInfo := cards.ArgsAddSingle{
		BsToken: bstoken,
		SingleBase: cards.SingleBase{
			Name:        "context2021",
			SortDesc:    "2011精油开背，放松自己2011",
			BindId:      7,
			TagIds:      []int{1, 2},
			RealPrice:   80,
			Price:       100,
			ServiceTime: 60,
			ImgHash:     "02457676043548aa8f8f7b274b2e4d41",
			//Sex:            1,
			//AgeBracket:     []int{1, 2, 3},
			//TailorIndus:    1,
			//TailorSubIndus: []int{1, 2, 3, 4},
			Pictures: []string{"d356d82039fc4463bb0c7e55c1acdd01"},
			//Effect:         []string{"淡斑", "改善问题性肌肤"},
			//UseProducts:    []string{"洁面产品", "洁面品"},
			//UseTools:       []string{"刮痧板", "暗疮针"},
			//UseInstrument:  "电吹风,理发剪",
			//StaffLevel:     []string{"初级美容师", "中级美容师", "高级美容师"},
			//Taboo:          []string{"静脉曲张", "疤痕体质"},
			Subscribe: "提前一天预约",
			// Reminder: map[string]interface{}{
			// 	"ReminderName":    "服务效果",
			// 	"ReminderContent": "不适合敏感体质",
			// },
			// SingleContext: map[string]interface{}{
			// 	"Content": "222",
			// },

			SpecIds: []cards.SingleSpecIds{
				{
					SpecId: 1,
					Sub:    []int{2},
				},
				{
					SpecId: 5,
					Sub:    []int{6, 7, 8},
				},
			},
			SpecPrices: []cards.SingleSpecPrice{
				{
					SpecIds: []int{7, 2},
					Price:   90,
				},
				{
					SpecIds: []int{6, 2},
					Price:   80,
				},
				{
					SpecIds: []int{8, 2},
					Price:   100,
				},
			},
			EffectImgs: []string{"b82275a6c75a43fd836d25be534932ad"},
			ToolsImgs:  []string{},
		},
	}

	logs.Info(mSingle.AddSingle(context.Background(), &singleInfo))
}

func TestSingleInfo(t *testing.T) {
	mSingle := new(logics.SingleLogic)
	sinfo, err := mSingle.SingleInfo(context.Background(), 160, 0, 0)
	logs.Info(err)
	da, _ := json.Marshal(sinfo)
	logs.Info(string(da))
}

func TestUpdate(t *testing.T) {
	bsToken := getBusToken()
	mSingle := new(logics.SingleLogic)
	args := cards.ArgsEditSingle{
		BsToken: bsToken,
		SingleBase: cards.SingleBase{
			Name:           "头部舒缓",
			SortDesc:       "精油开背，放松自己",
			BindId:         7,
			TagIds:         []int{1, 2},
			RealPrice:      88,
			Price:          108,
			ServiceTime:    60,
			ImgHash:        "02457676043548aa8f8f7b274b2e4d41",
			Sex:            1,
			AgeBracket:     []int{1, 2, 3},
			TailorIndus:    1,
			TailorSubIndus: []int{1, 2, 3, 4},
			Pictures:       []string{"02457676043548aa8f8f7b274b2e4d41"},
			// Reminder: map[string]interface{}{
			// 	"effect":         []interface{}{"淡斑", "改善问题性肌肤"},
			// 	"staff_level":    []interface{}{"初级美容师", "中级美容师", "高级美容师"},
			// 	"taboo":          []interface{}{"静脉曲张", "疤痕体质"},
			// 	"use_instrument": "电吹风,理发剪",
			// 	"subscribe":      []interface{}{"提前一天预约"},
			// },
			//Effect:         []string{"淡斑", "改善问题性肌肤"},
			//UseProducts:    []string{"洁面产品", "洁面品"},
			//UseTools:       []string{"刮痧板", "暗疮针"},
			//UseInstrument:  "电吹风,理发剪",
			//StaffLevel:     []string{"初级美容师", "中级美容师", "高级美容师"},
			//Taboo:          []string{"静脉曲张", "疤痕体质"},
			//Subscribe:      "提前一天预约",
			SpecIds: []cards.SingleSpecIds{
				{
					SpecId: 1,
					Sub:    []int{2, 3},
				},
				{
					SpecId: 5,
					Sub:    []int{7, 8},
				},
			},
			SpecPrices: []cards.SingleSpecPrice{
				{
					SpecIds: []int{7, 2},
					Price:   90,
				},
				{
					SpecIds: []int{8, 2},
					Price:   100,
				},
				{
					SpecIds: []int{7, 3},
					Price:   108,
				},
				{
					SpecIds: []int{8, 3},
					Price:   118,
				},
			},
			EffectImgs: []string{"b82275a6c75a43fd836d25be534932ad", "02457676043548aa8f8f7b274b2e4d41"},
			ToolsImgs:  []string{"d356d82039fc4463bb0c7e55c1acdd01"},
		},
		SingleId: 12,
	}
	logs.Info(mSingle.EditSingle(context.Background(), &args))
}

func TestShopAdd(t *testing.T) {
	bsToken := getBusToken()
	mSingle := new(logics.SingleLogic)
	args := cards.ArgsShopAddSingle{
		BsToken:  bsToken,
		SingleId: []int{11},
	}
	logs.Info(mSingle.ShopAddSingle(&args))
}

//获取商家的单项目列表
func TestBusSingles(t *testing.T) {
	mSingle := new(logics.SingleLogic)
	isGround := ""
	busId, shopId := 51, 429
	r, err := mSingle.GetBusSingles(context.Background(), busId, shopId, 0, 10, isGround, "0", true)
	logs.Info(err)
	str, _ := json.Marshal(r)
	logs.Info(string(str))
}

//测试子店修改单项目价格
func TestChangeSpecPrice(t *testing.T) {
	bsToken := getBusToken()
	mSingle := new(logics.SingleLogic)
	arg := cards.ArgsShopChangePrice{
		BsToken:   bsToken,
		SingleId:  11,
		RealPrice: 90,
		SpecPrice: []cards.SingleSpecPrice{
			{
				Price: 91,
				SspId: 25,
			},
			{
				Price: 101,
				SspId: 27,
			},
			{
				Price: 109,
				SspId: 32,
			},
			{
				Price: 119,
				SspId: 33,
			},
			{
				Price: 129,
				SspId: 37,
			},
			{
				Price: 139,
				SspId: 38,
			},
		},
	}
	logs.Info(mSingle.ShopChangePrice(&arg))
}

/*//测试总店设置项目上下架
func TestBusUpDownSingle(t *testing.T) {
	bsToken := getBusToken()
	mSingle := new(logics.SingleLogic)
	args := cards.ArgsDownUpSingle{
		BsToken:   bsToken,
		SingleIds: []int{11, 12},
		OptType:   cards.OPT_UP,
	}

	logs.Info(mSingle.BusUpDownSingle(context.Background(), &args))
}*/

//测试总店 和分店删除 单项目

//func TestDelSingle(t *testing.T) {
//	var reply bool
//	err := new(logics.SingleLogic).DelSingle(context.Background(), 0, 403, &cards.ArgsDelSingle{
//		SingleId: 117,
//	}, &reply)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log(reply)
//}

//获取子店的单项目列表
func TestShopSingle(t *testing.T) {
	mSingle := new(logics.SingleLogic)
	r, _ := mSingle.GetShopSingles(context.Background(), 414, 0, 10, []int{}, "", "")
	str, _ := json.Marshal(r)
	logs.Info(string(str))
}

//测试子店上下架
func TestShopUpDown(t *testing.T) {
	mSingle := new(logics.SingleLogic)
	bsToken := getBusToken()
	args := cards.ArgsShopDownUpSingle{
		BsToken: bsToken,
		SsIds:   []int{180},
		OptType: cards.STATUS_ON_SALE,
	}
	logs.Info(mSingle.ShopDownUpSingle(context.Background(), &args))

}

//获取单项目子店价格
func TestGetShopPrice(t *testing.T) {
	mSingle := new(logics.SingleLogic)
	args := cards.ArgsGetShopSinglePrice{
		SsId:  3,
		SspId: 0,
	}

	logs.Info(mSingle.GetShopSinglePrice(args.SsId, args.SspId))
}

//测试获取属性
func TestGetAttrs(t *testing.T) {
	mSingle := new(logics.SingleLogic)
	r := mSingle.GetAttrs()
	str, _ := json.Marshal(r)
	logs.Info(string(str))
}

/*
func TestGetSignlesByStaffId(t *testing.T) {
	mSingle := new(logics.SingleLogic)
	//	args *cards.ArgsGetSignlesByStaffID,reply *cards.ReplyGetSignlesByStaffID
	args := &cards.ArgsGetSignlesByStaffID{
		//BsToken:   common.BsToken{},
		Paging:    common.Paging{},
		GetAll:    false,
		BindId:    0,
		StaffId:   50,
		Name:      "",
		SspIds:    []int{50, 51, 52},
		SingleIds: []int{16},
	}
	var reply cards.ReplyGetSignlesByStaffID
	if err := mSingle.GetSignlesByStaffId(context.Background(), args, &reply); err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("%#v", reply)
}*/

//GetShopSingleBySingleIdsRpc(ctx context.Context,args *cards.ArgsGetShopSingleBySingleIdsRpc,reply *cards.ReplyGetShopSingleBySingleIdsRpc)
func TestGetShopSingleBySingleIdsRpc(t *testing.T) {
	args := &cards.ArgsGetShopSingleBySingleIdsRpc{
		ShopId:    1,
		SingleIds: []int{11, 12, 14},
	}
	var reply cards.ReplyGetShopSingleBySingleIdsRpc
	if err := new(logics.SingleLogic).GetShopSingleBySingleIdsRpc(context.Background(), args, &reply); err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("%#v", reply)
}

func TestGetSingleByShopIdAndTagId(t *testing.T) {
	var reply cards.ReplyShopSingle
	if err := new(logics.SingleLogic).GetSingleByShopIdAndTagId(context.Background(), &cards.ArgsShopSingleByPage{
		Paging: common.Paging{
			Page:     1,
			PageSize: 10,
		}, ShopId: 1,
	}, &reply); err != nil {
		t.Fatalf("出错了， err = %v", err)
	}
	bytes, _ := json.Marshal(reply)
	t.Logf("%v", string(bytes))
}

func TestSsids(t *testing.T) {
	//var ssIds = []int{3, 4, 5, 6}
	//var reply = []cards.ReplyGetBySsidsRpc{}
	//new(logics.SingleLogic).GetBySsidsRpc(&ssIds, &reply)
	//logs.Info(reply)

	var args = cards.ArgsGetShopSpecs{
		ShopId: 1,
		SspIds: []int{25},
	}
	var reply = []cards.ReplyGetShopSpecs{}
	if err := new(logics.SingleLogic).GetShopSpecs(&args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Logf("%v", string(bytes))
	//args := cards.ArgsGetItemsBySsids{
	//	SsIds:    []int{1, 2, 3, 4},
	//	ItemType: cards.ITEM_TYPE_hncard,
	//}
	//logs.Info(new(logics.ItemLogic).GetItemsBySsids(context.Background(), &args))

}

func TestA(t *testing.T) {
	var reply map[int][]cards.SubSpec
	err := new(logics.SpecLogic).GetBySspIds(&[]int{25, 27}, &reply)
	if err != nil {
		t.Fatalf("出错了 err = %#v", err)
	}
	t.Logf("%#v", reply)
}

//GetSingleSpecBySspId(ctx context.Context,args *cards.ArgsSubSpecID,reply *cards.ReplySubServer)
func TestGetSingleSpecBySspId(t *testing.T) {
	args := &cards.ArgsSubSpecID{
		SspIds: []int{25},
		ShopId: 1,
	}
	var reply cards.ReplySubServer2
	if err := new(logics.SingleLogic).GetSingleSpecBySspId(args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//GetSingleByShopIdAndSingleIds(ctx context.Context,args *cards.ArgsGetSingleByShopIdAndSingleIds,reply *cards.ReplyShopSingle)
func TestGetSingleByShopIdAndSingleIds(t *testing.T) {
	args := &cards.ArgsGetSingleByShopIdAndSingleIds{
		ShopId:    399,
		SingleIds: []int{111},
	}
	var reply cards.ReplyShopSingle
	//rpcCards := new(cards2.Single).Init()
	if err := new(logics.SingleLogic).GetSingleByShopIdAndSingleIds(context.Background(), args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

//GetByShopSingle2(ctx context.Context, args *cards.ArgsCommonShopSingle, reply *[]cards.ReplyCommonShopSingle)
func TestGetByShopSingle2(t *testing.T) {
	args := &cards.ArgsCommonShopSingle{
		ShopId: 1,
		//SingleIds: []int{11,12,14,15,16},
		SspIds: []int{45, 44},
		Status: 0,
	}
	//14
	var reply []cards.ReplyGetPriceByShopIdAndSsspId
	rpcS := new(cards2.Single).Init()
	if err := rpcS.GetPriceByShopIdAndSsspId(context.Background(), args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
	return
}

//GetByShopSspIds(ctx context.Context, args *cards.ArgsGetShopSpecs, reply *[]cards.ReplyCommonSingleSpec)
func TestGetByShopSspIds(t *testing.T) {
	args := &cards.ArgsGetShopSpecs{
		ShopId: 1,
		SspIds: []int{45, 44},
	}
	var reply []cards.ReplyCommonSingleSpec
	if err := new(service.Single).GetByShopSspIds(context.Background(), args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
	return
}

//ShopSinglePage(ctx context.Context, args *cards.ArgsShopSinglePage, reply *cards.ReplyShopSinglePage)
func TestShopSinglePage(t *testing.T) {
	args := &cards.ArgsShopSinglePage{
		Paging: common.Paging{Page: 2, PageSize: 10},
		ShopId: 403,
		Status: fmt.Sprint(cards.IS_SYNC_YES),
	}
	var reply cards.ReplyShopSinglePage
	if err := new(service.Single).ShopSinglePage(context.Background(), args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(reply)
	t.Log(string(bytes))
}

func TestShopChangePrice(t *testing.T) {
	args := &cards.ArgsShopChangePrice{
		BsToken:   common.BsToken{},
		SingleId:  101,
		RealPrice: 0.01,
		SpecPrice: []cards.SingleSpecPrice{
			{
				Price: 50,
				SspId: 172,
			},
			{
				Price: 60,
				SspId: 173,
			},
			{
				Price: 70,
				SspId: 174,
			},
			{
				Price: 80,
				SspId: 175,
			},
		},
	}
	if err := new(logics.SingleLogic).ShopChangePrice(args); err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(args)
}

//BusSinglePage(ctx context.Context, args *cards.ArgsBusSinglePage, reply *cards.ReplyBusSinglePage)
func TestBusSinglePage(t *testing.T) {
	args := &cards.ArgsBusSinglePage{
		Paging:   common.Paging{Page: 1, PageSize: 8},
		BusId:    6,
		IsGround: "1",
	}
	var reply cards.ReplyBusSinglePage
	if err := new(service.Single).BusSinglePage(context.Background(), args, &reply); err != nil {
		t.Log(err.Error())
		return
	}
	res, _ := json.Marshal(reply)
	t.Log(string(res))
}

func TestDelSingle(t *testing.T) {
	args := cards.ArgsDelSingle{
		SingleIds: []int{45, 44},
	}
	if err := new(logics.SingleLogic).DelSingle(context.Background(), &args); err != nil {
		t.Log("err", err)
	}
	t.Log(args)
}

func TestShopDelSingle(t *testing.T) {
	args := cards.ArgsDelSingle{
		SingleIds: []int{48, 47},
	}
	if err := new(logics.SingleLogic).DelShopSingle(context.Background(), 7, &args); err != nil {
		t.Log("err", err)
	}
	t.Log(args)
}
