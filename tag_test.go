//标签测试
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/10 16:17
package main_test

import (
	"context"
	"fmt"
	"git.900sui.cn/kc/kcgin/logs"
	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcinterface/client/bus"
	"git.900sui.cn/kc/rpcinterface/client/user"
	bus2 "git.900sui.cn/kc/rpcinterface/interface/bus"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	user2 "git.900sui.cn/kc/rpcinterface/interface/user"
	"testing"
)

func getBusToken() common.BsToken {
	ctx := context.Background()
	rpcUser := new(user.UserLogin).Init()
	loginParams := user2.CheckLoginParams{
		Channel: 1,
		Token:   "YzJqbmFzaWhnZGNmcm4xbGRpZTAxNjIxNTg3MzE0MTQwMDI4NjE2",
	}
	loginRep := user2.CheckLoginReply{}

	if err := rpcUser.CheckLogin(ctx, &loginParams, &loginRep); err != nil {
		fmt.Println(err)
	}

	rpcBus := new(bus.BusAuth).Init()
	defer rpcBus.Close()
	utoken := common.Utoken{UidEncodeStr: loginRep.UidEncodeStr}
	args := bus2.ArgsBusAuth{
		//BusId:  1,
		Utoken: utoken,
		ShopId: 429,
	}
	reply := bus2.ReplyBusAuth{}
	if err := rpcBus.BusAuth(ctx, &args, &reply); err != nil {
		fmt.Println(err)
	}
	return common.BsToken{EncodeStr: reply.EncodeStr}
}

func TestAdd(t *testing.T) {
	mTag := new(logics.TagsLogic)
	bsToken := getBusToken()
	args := cards.ArgAddTag{
		BsToken: bsToken,
		Name:    "漂亮",
	}
	logs.Info(mTag.AddTag(&args))
}

func TestEdit(t *testing.T) {
	mTag := new(logics.TagsLogic)
	bsToken := getBusToken()
	args := cards.ArgEditTag{
		BsToken: bsToken,
		Name:    "美丽",
		TagId:   3,
	}
	logs.Info(mTag.EditTag(&args))
}

func TestDel(t *testing.T) {
	mTag := new(logics.TagsLogic)
	bsToken := getBusToken()
	args := cards.ArgDelTag{
		BsToken: bsToken,
		TagId:   4,
	}
	logs.Info(mTag.DelTag(&args))
}

func TestGetBus(t *testing.T) {
	mTag := new(logics.TagsLogic)
	bsToken := getBusToken()
	logs.Info(mTag.GetBusTags(bsToken))
}
