/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/4/16 11:03
@Description:

*********************************************/
package main_test

import (
	"context"
	"encoding/json"
	"fmt"
	logics2 "git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"testing"
)

// 获取身份卡详情
func TestInfo(t *testing.T) {
	var args cards.InputParamsICardInfo
	var reply cards.OutputParamsICardInfo
	args.IcardID = 240
	//args.IcardID = 305

	logic := new(logics2.ICardLogic)
	err := logic.Info(context.Background(), args, &reply)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(reply)

}

func TestAddToShop(t *testing.T) {
	var args cards.InputParamsICardAddToShop
	var reply cards.OutputParamsICardAddToShop
	args.IcardIds = "[59]"
	// 1618559586
	logic := new(logics2.ICardLogic)
	busId,shopId:=23,1
	err := logic.AddToShop(context.Background(),busId,shopId, args, &reply)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(reply)
}

//List 列表
func TestList(t *testing.T) {
	busId,shopId:=23,1
	args:=cards.InputParamsICardList{
		PaginationInput:  common.PaginationInput{Page: 1,PageSize: 10},
		BsToken:          common.BsToken{},
		Status:           0,
		Ground:           0,
		ShopID:           0,
		IsDel:            "",
		FilterShopHasAdd: true,
	}
	var reply cards.OutputICardList
	if err:=new(logics2.ICardLogic).List(context.Background(),busId,shopId,args,&reply);err!=nil{
		t.Log(err.Error())
		return
	}
	bytes,_:=json.Marshal(reply)
	t.Log(string(bytes))
}