//RcardModel
//2020-08-05 15:23:57

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type RcardModel struct {
	Model *base.Model
	Field RcardModelField
}

//表字段
type RcardModelField struct {
	T_table                 string `default:"rcard"`
	F_rcard_id              string `default:"rcard_id"`
	F_bus_id                string `default:"bus_id"`
	F_bind_id               string `default:"bind_id"`
	F_name                  string `default:"name"`
	F_sort_desc             string `default:"sort_desc"`
	F_real_price            string `default:"real_price"`
	F_price                 string `default:"price"`
	F_discount_type         string `default:"discount_type"`
	F_discount              string `default:"discount"`
	F_is_have_discount      string `default:"is_have_discount"`
	F_is_permanent_validity string `default:"is_permanent_validity"`
	F_service_period        string `default:"service_period"`
	F_has_give_signle       string `default:"has_give_signle"`
	F_is_ground             string `default:"is_ground"`
	F_is_del                string `default:"is_del"`
	F_del_time              string `default:"del_time"`
	F_under_time            string `default:"under_time"`
	F_img_id                string `default:"img_id"`
	F_sales                 string `default:"sales"`
	F_sale_shop_num         string `default:"sale_shop_num"`
	F_ctime                 string `default:"ctime"`
}

//初始化
func (r *RcardModel) Init(ormer ...orm.Ormer) *RcardModel {
	functions.ReflectModel(&r.Field)
	r.Model = base.NewModel(r.Field.T_table, ormer...)
	return r
}

//新增数据
func (r *RcardModel) Insert(data map[string]interface{}) int {
	result, _ := r.Model.Data(data).Insert()
	return result
}

func (r *RcardModel) GetRcardsByRcardIds(rcardIds []int, fields ...string) []map[string]interface{} {
	if len(rcardIds) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		r.Model.Field(fields)
	}
	return r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", rcardIds},
	}).OrderBy(r.Field.F_rcard_id + " DESC ").Select()
}

//根据充值卡id获取单条充值卡数据
func (r *RcardModel) GetByRcardId(rcardId int, fields ...[]string) map[string]interface{} {
	if rcardId <= 0 {
		return map[string]interface{}{}
	}
	if len(fields) > 0 {
		r.Model.Field(fields[0])
	}
	return r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Find()
}

//批量修改
func (r *RcardModel) UpdateByRcardids(rcardIds []int, data map[string]interface{}) bool {
	if len(rcardIds) == 0 || len(data) == 0 {
		return false
	}
	_, err := r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", rcardIds},
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//根据rcardid修改数据
func (r *RcardModel) UpdateByRcardid(rcardId int, data map[string]interface{}) bool {
	if rcardId <= 0 || len(data) == 0 {
		return false
	}
	_, err := r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//增加充值卡的销量
func (r *RcardModel) IncrSalesByRcardid(rcardId int, step int) bool {
	if rcardId <= 0 || step <= 0 {
		return false
	}

	_, err := r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Data(map[string]interface{}{
		r.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}

//增加在售店铺数量
func (r *RcardModel) IncrSaleShopNumByRcardid(rcardId, step int) bool {
	if rcardId <= 0 || step <= 0 {
		return false
	}

	_, err := r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Data(map[string]interface{}{
		r.Field.F_sale_shop_num: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}
	return true
}

//减少在售店铺数量
func (r *RcardModel) DecrSaleShopNumByRcardid(rcardId, step int) bool {
	if rcardId <= 0 || step <= 0 {
		return false
	}

	_, err := r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Data(map[string]interface{}{
		r.Field.F_sale_shop_num: []interface{}{"dec", step},
	}).Update()
	if err != nil {
		return false
	}
	return true
}

//获取企业的充值卡列表
func (r *RcardModel) GetPageByBusId(busId int, start, limit int, isGround ...int) []map[string]interface{} {
	if busId <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}
	where := map[string]interface{}{
		r.Field.F_bus_id: busId,
	}
	if len(isGround) > 0 {
		where[r.Field.F_is_ground] = isGround[0]
	}
	return r.Model.Where(where).OrderBy(r.Field.F_rcard_id+" DESC ").Limit(start, limit).Select()
}

// 获取充值卡列表 根据是否删除状态
func (r *RcardModel) GetPageIsDelByBusId(busId int, start, limit int, isDel ...int) []map[string]interface{} {
	if busId <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}
	where := map[string]interface{}{
		r.Field.F_bus_id: busId,
	}
	if len(isDel) > 0 {
		where[r.Field.F_is_del] = isDel
	}
	return r.Model.Where(where).OrderBy(r.Field.F_rcard_id+" DESC ").Limit(start, limit).Select()
}

func (r *RcardModel) SelectRcardsByWherePage(where []base.WhereItem, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		r.Model.Field(fields)
	}
	if limit > 0 {
		r.Model.Limit(start, limit)
	}
	return r.Model.Where(where).OrderBy(r.Field.F_rcard_id + " DESC ").Select()
}

func (r *RcardModel) GetNumByWhere(where []base.WhereItem) int {
	return r.Model.Where(where).Count(r.Field.F_rcard_id)
}

//获取商家的充值卡数量
func (r *RcardModel) GetNumByBusId(busId int, isGround ...int) int {
	if busId <= 0 {
		return 0
	}
	where := map[string]interface{}{
		r.Field.F_bus_id: busId,
	}
	if len(isGround) > 0 {
		where[r.Field.F_is_ground] = isGround[0]
	}
	return r.Model.Where(where).Count(r.Field.F_rcard_id)
}

// 获取商家的充值卡数量 根据是否删除状态
func (r *RcardModel) GetNumIsDelByBusId(busId int, isDel ...int) int {
	if busId <= 0 {
		return 0
	}
	where := map[string]interface{}{
		r.Field.F_bus_id: busId,
		r.Field.F_is_del: isDel,
	}
	return r.Model.Where(where).Count(r.Field.F_rcard_id)
}

//根据rcardIds批量获取数据
func (r *RcardModel) GetByRcardids(rcardIds []int, fields ...[]string) []map[string]interface{} {
	if len(rcardIds) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		r.Model.Field(fields[0])
	}
	return r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", rcardIds},
	}).Select()
}

//根据cardIDs批量获取数据
func (r *RcardModel) GetByRcardIDs(cardIDs []int, fields ...string) []map[string]interface{} {
	if len(cardIDs) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		r.Model.Field(fields)
	}
	return r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", cardIDs},
	}).Select()
}

//增加限时限次卡的销量
func (r *RcardModel) IncrSalesByRcardID(rcardID int, step int) (err error) {
	if rcardID == 0 || step <= 0 {
		return
	}
	_, err = r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardID,
	}).Data(map[string]interface{}{
		r.Field.F_sales: []interface{}{"inc", step},
	}).Update()

	return
}

func (r *RcardModel) Select(where map[string]interface{}, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		r.Model.Field(fields)
	}
	return r.Model.Where(where).Select()
}
