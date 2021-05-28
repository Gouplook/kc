//套餐测试用例
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/15 18:43
package main_test

import (
	"context"
	"encoding/json"
	"git.900sui.cn/kc/base/utils"
	"git.900sui.cn/kc/kcgin/logs"
	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcCards/common/tools"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"testing"
)

func TestMain(m *testing.M) {
	//初始化项目分享链接
	tools.InitShareLink()
	m.Run()
}

func TestMode(t *testing.T) {
	utils.CreateModel("sm_ext")
}

func TestAddSm(t *testing.T) {
	bsToken := getBusToken()
	mSm := new(logics.SmLogic)
	args := cards.ArgsAddSm{
		BsToken: bsToken,
		SmBase: cards.SmBase{
			Name:          "洗剪吹套餐5次zxxx",
			SortDesc:      "洗剪吹套餐5次,good",
			RealPrice:     100,
			Price:         150,
			ServicePeriod: 3,
		},
		Notes: []cards.CardNote{
			{
				Notes: "提示1",
			},
		},
		//IncludeSingles: []cards.IncSingle2{
		//	{
		//		SingleID: 11,
		//		Num:      2,
		//		SspId: 25,
		//	},
		//	{
		//		SingleID: 11,
		//		Num:      2,
		//		SspId: 26,
		//	},
		//	{
		//		SingleID: 18,
		//		Num:      3,
		//		SspId: 0,
		//	},
		//},
		GiveSingles: []cards.IncSingle{
			{
				SingleID: 12,
				Num:      2,
			},
			{
				SingleID: 18,
				Num:      3,
			},
		},
		ImgHash: "",
	}

	logs.Info(mSm.AddSm(context.Background(), &args))
}

//测试修改
func TestEditSm(t *testing.T) {
	bsToken := getBusToken()
	mSm := new(logics.SmLogic)
	args := cards.ArgsEditSm{
		BsToken: bsToken,
		SmId: 200,
		SmBase: cards.SmBase{
			Name:          "水泉套餐V3",
			SortDesc:      "水泉套餐V3",
			RealPrice:     0.01,
			Price:         128,
			ServicePeriod: 0,
		},
		Notes: []cards.CardNote{
			{
				Notes: "提示水泉套餐V311",
			},
			{
				Notes: "提示水泉套餐V311",
			},

		},
		IncludeSingles: []cards.IncSingle{
			{
				SingleID: 144,
				Num:      120, // 120--->122
			},
		},
		GiveSingles: []cards.IncSingle{
			{
				SingleID: 144,
				Num:      10, // 10 -12
				PeriodOfValidity:30, // 30--50
			},
		},
		ImgHash: "b82275a6c75a43fd836d25be534932ad",
	}

	logs.Info(mSm.EditSm(context.Background(), &args))
}

func TestSmInfo(t *testing.T) {
	mSm := new(logics.SmLogic)
	r, err := mSm.SmInfo(context.Background(), 200, 429)
	logs.Info(err)
	rbyte, _ := json.Marshal(r)
	logs.Info(string(rbyte))
}

//测试商家的套餐列表
func TestBusSmPage(t *testing.T) {
	mSm := new(logics.SmLogic)
	isGround := "2"
	busId, shopId := 1, 2
	rs, _ := mSm.GetBusPage(context.Background(), busId, shopId, 0, 10, isGround,true)
	byters, _ := json.Marshal(rs)
	t.Log(string(byters))
}

//测试设置适用门店
func TestSetSmShop(t *testing.T) {
	bsToken := getBusToken()
	mSm := new(logics.SmLogic)
	args := &cards.ArgsSetSmShop{
		BsToken:   bsToken,
		SmIds:     []int{1, 2, 3, 4, 5},
		ShopIds:   []int{1},
		IsAllShop: false,
	}

	logs.Info(mSm.SetSmShop(context.Background(), args))
}

//测试总店上下架套餐
func TestBusUpDown(t *testing.T) {
	bsToken := getBusToken()
	mSm := new(logics.SmLogic)
	args := &cards.ArgsDownUpSm{
		BsToken: bsToken,
		SmIds:   []int{1, 2, 5},
		OptType: cards.OPT_DOWN,
	}

	logs.Info(mSm.DownUpSm(context.Background(), args))
}

//测试门店适用的套餐数据
func TestShopGetBusSm(t *testing.T) {
	mSm := new(logics.SmLogic)
	r, err := mSm.ShopGetBusSmPage(context.Background(), 1, 1, 0, 10)
	logs.Info(err)
	rbyte, _ := json.Marshal(r)
	logs.Info(string(rbyte))
}

//测试门店添加套餐
func TestShopAddSm(t *testing.T) {
	bsToken := getBusToken()
	mSm := new(logics.SmLogic)
	args := cards.ArgsShopAddSm{
		BsToken: bsToken,
		SmIds:   []int{207},
	}

	logs.Info(mSm.ShopAddSm(&args))
}

func TestShoptoAddSm(t *testing.T) {
	// bsToken := getBusToken()
	mSm := new(logics.SmLogic)
	args := cards.ArgsShopAddSm{
		// BsToken: bsToken,
		SmIds:   []int{207},
	}

	logs.Info(mSm.ShopAddToSm(context.Background(),&args))
}

//获取门店的套餐列表
func TestShopSmPage(t *testing.T) {
	mSm := new(logics.SmLogic)
	r, err := mSm.ShopSmPage(context.Background(), 1, 0, 10, cards.STATUS_ON_SALE)
	logs.Info(err)
	rbyte, _ := json.Marshal(r)
	logs.Info(string(rbyte))
}

//门店上下架套餐
func TestShopDownUp(t *testing.T) {
	bsToken := getBusToken()
	mSm := new(logics.SmLogic)
	args := cards.ArgsShopDownUpSm{
		BsToken:   bsToken,
		ShopSmIds: []int{1, 2, 3, 4, 5},
		OptType:   cards.OPT_DOWN,
	}

	logs.Info(mSm.ShopDownUpSm(context.Background(), &args))

}

func TestOrder(t *testing.T) {
	//orderSn := "JS10001"
	//logs.Info(fmt.Sprintf("%s%s", string(orderSn[0]), string(orderSn[1])))

	orderSn := "JP4634331631339700230"
	reply := false
	err := new(logics.ItemLogic).IncrItemSales(context.Background(), orderSn, &reply)
	logs.Info(err)
}
