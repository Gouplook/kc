//HcardModel 卡项服务-限时卡表
//2020-04-21 14:55:12

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

// 在售门店数量修改的量
var addRemoveSaleShopNum = 1

// HcardModel 表结构体
type HcardModel struct {
	Model *base.Model
	Field HcardModelField
}

// HcardModelField 表字段
type HcardModelField struct {
	T_table           string `default:"hcard"`
	F_hcard_id        string `default:"hcard_id"`
	F_bus_id          string `default:"bus_id"`
	F_bind_id         string `default:"bind_id"`
	F_name            string `default:"name"`
	F_sort_desc       string `default:"sort_desc"`
	F_real_price      string `default:"real_price"`
	F_price           string `default:"price"`
	F_service_period  string `default:"service_period"`
	F_has_give_signle string `default:"has_give_signle"`
	F_is_ground       string `default:"is_ground"` // 是否上架 0=否 1=是
	F_is_del          string `default:"is_del"`
	F_del_time        string `default:"del_time"`
	F_under_time      string `default:"under_time"`
	F_img_id          string `default:"img_id"`
	F_sales           string `default:"sales"`
	F_sale_shop_num   string `default:"sale_shop_num"`
	F_ctime           string `default:"ctime"`
}

// Init 初始化
func (h *HcardModel) Init(ormer ...orm.Ormer) *HcardModel {
	functions.ReflectModel(&h.Field)
	h.Model = base.NewModel(h.Field.T_table, ormer...)
	return h
}

// Insert 新增数据
func (h *HcardModel) Insert(data map[string]interface{}) int {
	result, _ := h.Model.Data(data).Insert()
	return result
}

func (h *HcardModel) Select(where map[string]interface{}, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	return h.Model.Where(where).Select()
}

// GetHcardByID 根据hcard id获取获取单条限时卡数据
func (h *HcardModel) GetHcardByID(hcardID int, fields ...string) map[string]interface{} {
	if hcardID == 0 {
		return make(map[string]interface{})
	}
	return h.Model.Field(fields).Where(map[string]interface{}{
		h.Field.F_hcard_id: hcardID,
	}).Find()
}

// IncrSalesByHcardID 增加限时卡销量
func (h *HcardModel) IncrSalesByHcardID(sales, hcardID int) bool {
	if sales <= 0 || hcardID <= 0 {
		return false
	}
	_, err := h.Model.Where(map[string]interface{}{h.Field.F_hcard_id: hcardID}).Data(map[string]interface{}{
		h.Field.F_sales: []interface{}{"inc", sales}}).Update()
	if err != nil {
		return false
	}
	return true
}

// UpdateByHcardIDs 批量更新限时卡数据(上下架等)
func (h *HcardModel) UpdateByHcardIDs(hcardIDs []int, data map[string]interface{}) bool {
	if len(hcardIDs) == 0 || len(data) == 0 {
		return false
	}
	_, err := h.Model.Where(map[string]interface{}{h.Field.F_hcard_id: []interface{}{"IN", hcardIDs}}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

// 批量软删除限时卡
func (h *HcardModel) UpdateDelByHcardIds(hcardIds []int, busId int, data map[string]interface{}) bool {
	if len(hcardIds) == 0 || len(data) == 0 {
		return false
	}
	_, err := h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id: []interface{}{"IN", hcardIds},
		h.Field.F_bus_id:   busId,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

// UpdateByHcardID 根据限时卡id更新数据
func (h *HcardModel) UpdateByHcardID(hcardID int, data map[string]interface{}) bool {
	if hcardID == 0 || len(data) == 0 {
		return false
	}
	_, err := h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id: hcardID,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

// GetPageByBusID 获取限时卡列表数据(分页)
func (h *HcardModel) GetPageByBusID(busID int, start, limit int, isGround ...int) []map[string]interface{} {
	where := map[string]interface{}{
		h.Field.F_bus_id: busID,
	}
	if len(isGround) > 0 {
		where[h.Field.F_is_ground] = isGround[0]
	}
	return h.Model.Where(where).OrderBy(h.Field.F_hcard_id+" DESC ").Limit(start, limit).Select()
}

// 获取限时卡列表数据 过滤掉是否删除的数据(分页)
func (h *HcardModel) GetBusHcardBusId(busId int, start, limit, isDe int) (data []map[string]interface{}) {
	if busId <= 0 {
		return
	}
	where := map[string]interface{}{
		h.Field.F_bus_id: busId,
		h.Field.F_is_del: isDe,
	}
	data = h.Model.Where(where).OrderBy(h.Field.F_hcard_id+" DESC ").Limit(start, limit).Select()
	return
}

func (h *HcardModel) SelectHcardsByWherePage(where []base.WhereItem, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	if limit > 0 {
		h.Model.Limit(start, limit)
	}
	return h.Model.Where(where).OrderBy(h.Field.F_hcard_id + " DESC ").Select()
}

func (h *HcardModel) GetNumByWhere(where []base.WhereItem) int {
	return h.Model.Where(where).Count(h.Field.F_hcard_id)
}

//获取商家限时卡数量 过滤掉是否删除的数据
func (h *HcardModel) GetNumHcardBusId(busId int, isDel int) (count int) {
	if busId <= 0 {
		return
	}
	where := map[string]interface{}{
		h.Field.F_bus_id: busId,
		h.Field.F_is_del: isDel,
	}
	count = h.Model.Where(where).Count(h.Field.F_hcard_id)
	return
}

// GetNumByBusID 获取商家限时卡数量
func (h *HcardModel) GetNumByBusID(busID int, isGround ...int) int {
	if busID == 0 {
		return 0
	}
	where := map[string]interface{}{
		h.Field.F_bus_id: busID,
	}
	if len(isGround) > 0 {
		where[h.Field.F_is_ground] = isGround[0]
	}
	return h.Model.Where(where).Count(h.Field.F_hcard_id)
}

// GetHcardByIDs 根据hcard id批量获取数据
func (h *HcardModel) GetHcardByIDs(hardIDs []int, fields ...string) []map[string]interface{} {
	if len(hardIDs) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id: []interface{}{"IN", hardIDs},
	}).Select()
}

//根据hcardIDs批量获取上架的限时卡数据
func (h *HcardModel) GetByHcardIDsAndGround(hcardIDs []int, fields ...string) []map[string]interface{} {
	if len(hcardIDs) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id:  []interface{}{"IN", hcardIDs},
	}).Select()
}

//UpdateSaleShopNum 更新在售门店数量
func (h *HcardModel) UpdateSaleShopNum(hcardIDs []int, decOrInc string /*可用数值为:dec和inc分别代表增减操作*/) bool {
	if len(hcardIDs) == 0 || len(decOrInc) == 0 {
		return false
	}
	_, err := h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id: []interface{}{"IN", hcardIDs},
	}).Data(map[string]interface{}{
		h.Field.F_sale_shop_num: []interface{}{decOrInc, addRemoveSaleShopNum},
	}).Update()
	if err != nil {
		return false
	}
	return true
}

//FindHcardByHcardIDAndBusID 根据hcardID busID找寻数据
func (h *HcardModel) FindHcardByHcardIDAndBusID(hcardIDs []int, busID int, fields ...string) []map[string]interface{} {
	if len(hcardIDs) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id: []interface{}{"IN", hcardIDs},
		h.Field.F_bus_id:   busID,
	}).Select()
}

// DeleteHcardByHcardIDs 删除限时卡
func (h *HcardModel) DeleteHcardByHcardIDs(hcardIDs []int, busID int) bool {
	if len(hcardIDs) == 0 {
		return false
	}
	_, err := h.Model.Where(map[string]interface{}{
		h.Field.F_bus_id:   busID,
		h.Field.F_hcard_id: []interface{}{"IN", hcardIDs},
	}).Delete()
	if err != nil {
		return false
	}
	return true
}
