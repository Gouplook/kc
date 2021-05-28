//单项目规格
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/7 11:09
package logics

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/mapstructure"
	"git.900sui.cn/kc/rpcCards/common/models"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"strconv"
)

type SpecLogic struct {
}

//添加规格
//@Param int busId 商家id
//@Param int parentId 规格的父id 顶级规格为0
//@Param name string 规格名称
//@return int 新增规格id
func (s *SpecLogic) AddSpec(busId int, parentId int, name string) (int, error) {
	mSpec := new(models.SingleBusSpecModel).Init()
	defer func() {
		mSpec = nil
	}()
	//如果parentId > 0 , 检查parentId是否属于busId
	if parentId > 0 {
		parentSpec := mSpec.GetBySpecid(parentId)
		if len(parentSpec) == 0 {
			return 0, toolLib.CreateKcErr(_const.SPEC_PARENT_NO)
		}
		var specInfo = struct {
			BusId int
		}{0}
		mapstructure.WeakDecode(parentSpec, &specInfo)
		if specInfo.BusId != busId {
			return 0, toolLib.CreateKcErr(_const.POWER_ERR)
		}
	}

	//先检查同名的规格是否已经存在
	r := mSpec.GetByBusidAndName(busId, parentId, name)
	if len(r) > 0 {
		return 0, toolLib.CreateKcErr(_const.SPEC_HASED)
	}
	//插入数据
	specId := mSpec.Insert(map[string]interface{}{
		mSpec.Field.F_p_spec_id: parentId,
		mSpec.Field.F_name:      name,
		mSpec.Field.F_bus_id:    busId,
	})

	if specId == 0 {
		return 0, toolLib.CreateKcErr(_const.DB_ERR)
	}
	return specId, nil
}

//获取子规格
//@param int busId 企业id
//@param int parentId 父规格id
//@return []cards.SubSpec
func (s *SpecLogic) GetSubSpec(busId int, parentId int) []cards.SubSpec {
	mSpec := new(models.SingleBusSpecModel).Init()
	defer func() {
		mSpec = nil
	}()
	r := mSpec.GetByParentSpecId(parentId, busId)
	var res []cards.SubSpec
	mapstructure.WeakDecode(r, &res)
	return res
}

//获取多个规格数据
func (s *SpecLogic) GetBySpecIds(specIds []int) ([]cards.SubSpec, error) {
	if len(specIds) == 0 {
		return []cards.SubSpec{}, toolLib.CreateKcErr(_const.PARAM_ERR)
	}
	mSpec := new(models.SingleBusSpecModel).Init()
	defer func() {
		mSpec = nil
	}()
	r := mSpec.GetBySpecids(specIds)
	var res []cards.SubSpec
	mapstructure.WeakDecode(r, &res)
	return res, nil
}

//检查传递的规格id是否是商家的
func (s *SpecLogic) CheckBusSpecIds(busId int, specIds []int) error {
	r, err := s.GetBySpecIds(specIds)
	if err != nil {
		return err
	}
	for _, spec := range r {
		if spec.BusId != busId {
			return toolLib.CreateKcErr(_const.SPEC_NOT_IN_BUS)
		}
	}
	return nil
}

//根据多个父规格id获取子规格
func (s *SpecLogic) GetByParentSpecIds(busId int, parentIds []int) []cards.SubSpec {
	mSpec := new(models.SingleBusSpecModel).Init()
	defer func() {
		mSpec = nil
	}()
	r := mSpec.GetByParentSpecIds(parentIds, busId)
	var res []cards.SubSpec
	mapstructure.WeakDecode(r, &res)
	return res
}

//根据sspid 查询子规格名字和id
func (s *SpecLogic) GetBySspIds(args *[]int, reply *map[int][]cards.SubSpec) error {
	sspModel := new(models.SingleSpecPriceModel).Init()
	sspMaps := sspModel.GetBySspids(*args)

	var specIds []int
	var specMap = make(map[int][]int)
	var ids []int
	for _, sspMap := range sspMaps {
		ids = functions.StrExplode2IntArr(sspMap[sspModel.Field.F_spec_ids].(string), ",")
		specIds = append(specIds, ids...)
		sspId, _ := strconv.Atoi(sspMap[sspModel.Field.F_ssp_id].(string))
		specMap[sspId] = ids
	}
	specIds = functions.ArrayUniqueInt(specIds)
	res, err := s.GetBySpecIds(specIds)
	if err != nil {
		return err
	}

	var reply2 = make(map[int][]cards.SubSpec)
	for sspId, subSpecs := range specMap {
		var specs []cards.SubSpec
		for _, spec := range subSpecs {
			for _, re := range res {
				if re.SpecId == spec {
					specs = append(specs, re)
				}
			}
			reply2[sspId] = specs
		}
	}
	*reply = reply2
	return nil
}
