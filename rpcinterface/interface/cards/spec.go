//商家单项目规格管理接口
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/7 11:00
package cards

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//单项目规格
type SpecInfo struct {
	common.BsToken
	Name string //规格名称
	ParentSpecId int //父规格id
}

//子规格
type SubSpec struct {
	Name string
	SpecId int
	PSpecId int
	BusId int
}

//获取子规格的参数结构体
type GetSubSpecParams struct {
	common.BsToken
	PSpecId int
}

type Spec interface {
	//增加规格
	AddSpec( ctx context.Context, args *SpecInfo, reply *int ) error
	//根据父规格id获取子规格
	GetByParentSpecId( ctx context.Context, args *GetSubSpecParams, reply *[]SubSpec ) error
	//根据sspid 查询子规格名字和id
	GetBySspIds(ctx context.Context, args *[]int, reply *map[int][]SubSpec ) error
}