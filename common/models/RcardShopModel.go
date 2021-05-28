//RcardShopModel
//2020-10-20 18:07:12

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type RcardShopModel struct {
	Model *base.Model
	Field RcardShopModelField
}

//表字段
type RcardShopModelField struct {
	T_table    string `default:"rcard_shop"`
	F_id       string `default:"id"`
	F_rcard_id string `default:"rcard_id"`
	F_bus_id   string `default:"bus_id"`
	F_shop_id  string `default:"shop_id"`
}

//初始化
func (r *RcardShopModel) Init(ormer ...orm.Ormer) *RcardShopModel {
	functions.ReflectModel(&r.Field)
	r.Model = base.NewModel(r.Field.T_table, ormer...)
	return r
}

//新增数据
func (r *RcardShopModel) Insert(data map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}
	result, _ := r.Model.Data(data).Insert()
	return result
}

//批量添加
func (r *RcardShopModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}

	rs, _ := r.Model.InsertAll(data)
	return rs
}

//根据ids批量修改
func (r *RcardShopModel) UpdateByIds(ids []int, data map[string]interface{}) bool {
	if len(ids) == 0 || len(data) == 0 {
		return false
	}
	_, err := r.Model.Where(map[string]interface{}{
		r.Field.F_id: []interface{}{"IN", ids},
	}).Data(data).Update()

	if err != nil {
		return false
	}
	return true
}

//根据充值卡ids 获取所有可适用门店记录
func (r *RcardShopModel) GetByRcardids(rcardIds []int) []map[string]interface{} {
	if len(rcardIds) <= 0 {
		return []map[string]interface{}{}
	}

	return r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", rcardIds},
	}).Select()
}

//根据充值卡id 获取所有可适用门店记录
func (r *RcardShopModel) GetByRcardId(rcardId int) []map[string]interface{} {
	if rcardId <= 0 {
		return []map[string]interface{}{}
	}

	return r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Select()
}
//根据充值卡id and busId 获取所有可适用门店记录
func (r *RcardShopModel)GetByRcardIdAndBusId(rcardId int ,busId int )[]map[string]interface{}{
	if rcardId <= 0 {
		return []map[string]interface{}{}
	}
	return r.Model.Where([]base.WhereItem{
		{r.Field.F_rcard_id,rcardId},
		{r.Field.F_bus_id,busId},
	}).Select()
}


//根据门店id获取门店可添加的套餐列表
func (r *RcardShopModel) GetPageByShopId(busId, shopId, start, limit int) []map[string]interface{} {
	if busId <= 0 || shopId <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}

	return r.Model.Where(map[string]interface{}{
		r.Field.F_bus_id:  busId,
		r.Field.F_shop_id: []interface{}{"IN", []int{0, shopId}},
	}).OrderBy(r.Field.F_id+" DESC ").Limit(start, limit).Select()
}

//根据rcardIds和shopid获取适用门店的充值卡数据
func (r *RcardShopModel) GetByShopIdAndRcardids(busId, shopId int, rcardIds []int) []map[string]interface{} {
	if busId <= 0 || shopId <= 0 || len(rcardIds) == 0 {
		return []map[string]interface{}{}
	}
	return r.Model.Where(map[string]interface{}{
		r.Field.F_bus_id:   busId,
		r.Field.F_shop_id:  []interface{}{"IN", []int{0, shopId}},
		r.Field.F_rcard_id: []interface{}{"IN", rcardIds},
	}).Select()

}

//根据门店id获取门店可添加的冲值卡总数量
func (r *RcardShopModel) GetNumByShopId(busId, shopId int) int {
	if busId <= 0 || shopId <= 0 {
		return 0
	}
	return r.Model.Where(map[string]interface{}{
		r.Field.F_bus_id:  busId,
		r.Field.F_shop_id: []interface{}{"IN", []int{0, shopId}},
	}).Count(r.Field.F_id)
}

//批量删除
func (s *RcardShopModel) DelByRcardids(rcardIds []int) bool {
	if len(rcardIds) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_rcard_id: []interface{}{"IN", rcardIds},
	}).Delete()
	if err != nil {
		return false
	}
	return true
}

//根据限时综合卡ids 获取所有可适用门店记录
func (r *RcardShopModel) GetByRcardIDs( cardID []int ) (data []map[string]interface{})  {
	if len(cardID) <= 0{
		return
	}
	data = r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", cardID},
	}).Select()

	return
}