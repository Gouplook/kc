//规格
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/7 18:25
package service

import (
	"context"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/rpcCards/common/logics"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

type Spec struct {

}

//添加一个规格
func (s *Spec) AddSpec( ctx context.Context, args *cards.SpecInfo, reply *int ) error {
	//先验证商户信息
	busId, err := args.BsToken.GetBusId()
	if err != nil{
		return  toolLib.CreateKcErr( _const.POWER_ERR )
	}

	mSpec := new(logics.SpecLogic)
	*reply, err = mSpec.AddSpec( busId, args.ParentSpecId, args.Name )
	if err != nil{
		return err
	}

	return nil
}

//获取子规格
func (s *Spec)  GetByParentSpecId( ctx context.Context, args *cards.GetSubSpecParams, reply *[]cards.SubSpec ) error{
	//先验证商户信息
	busId, err := args.BsToken.GetBusId()
	if err != nil{
		return  toolLib.CreateKcErr( _const.POWER_ERR )
	}
	mSpec := new(logics.SpecLogic)
	*reply = mSpec.GetSubSpec( busId, args.PSpecId )

	return nil
}

//根据sspid 查询子规格名字和id
func (s *Spec)  GetBySspIds(ctx context.Context, args *[]int, reply *map[int][]cards.SubSpec ) (err error) {
	err = new(logics.SpecLogic).GetBySspIds(args, reply)
	return
}