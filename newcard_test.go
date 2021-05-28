package main_test

import (
	"context"
	"fmt"
	"testing"

	"git.900sui.cn/kc/rpcinterface/client/cards"
	cards2 "git.900sui.cn/kc/rpcinterface/interface/cards"
)

//综合卡-总店删除卡
func TestCardDel(t *testing.T) {
	card := new(cards.Card).Init()
	args := &cards2.ArgsDeleteCard{
		CardIds: []int{4},
	}
	reply := true
	if err := card.DeleteCard(context.Background(), args, &reply); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("success")
}

//综合卡-分店删除卡
func TestShopCardDel(t *testing.T) {
	card := new(cards.Card).Init()
	args := &cards2.ArgsDeleteShopCard{
		CardIds: []int{29, 30},
	}
	reply := true
	if err := card.DeleteShopCard(context.Background(), args, &reply); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("success")
}

//套餐卡-总店删除
func TestSmDel(t *testing.T) {
	card := new(cards.Sm).Init()
	args := &cards2.ArgsDeleteSm{
		SmIds: []int{1, 2},
	}
	reply := true
	if err := card.DeleteSm(context.Background(), args, &reply); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("success")
}

//套餐卡-分店删除
func TestShopSmDel(t *testing.T) {
	card := new(cards.Sm).Init()
	args := &cards2.ArgsDeleteShopSm{
		SmIds: []int{5, 6},
	}
	reply := true
	if err := card.DeleteShopSm(context.Background(), args, &reply); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("success")
}

//总店-限次卡
func TestNCardDel(t *testing.T) {
	card := new(cards.NCard).Init()
	args := &cards2.ArgsDeleteNCard{
		NcardIds: []int{1, 2, 3},
	}
	reply := true
	if err := card.DeleteNCard(context.Background(), args, &reply); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("success")
}

//分店-限次卡
func TestShopNCardDel(t *testing.T) {
	card := new(cards.NCard).Init()
	args := &cards2.ArgsDeleteShopNCard{
		NcardIds: []int{7, 8},
	}
	reply := true
	if err := card.DeleteShopNCard(context.Background(), args, &reply); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("success")
}
