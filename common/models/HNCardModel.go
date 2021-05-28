// HNCardModel
// 2020-04-16 11:27:02

package models

import (
	"errors"
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

// 表结构体
type HNCardModel struct {
	Model *base.Model
	Field HNCardModelField
}

// 表字段
type HNCardModelField struct {
	T_table                 string `default:"hncard"`
	F_hncard_id             string `default:"hncard_id"`
	F_bus_id                string `default:"bus_id"`
	F_bind_id               string `default:"bind_id"`
	F_name                  string `default:"name"`
	F_sort_desc             string `default:"sort_desc"`
	F_real_price            string `default:"real_price"`
	F_price                 string `default:"price"`
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
	F_validcount            string `default:"validcount"`
	// F_is_permanent_validity string `default:"is_permanent_validity"`
}

// 初始化
func (h *HNCardModel) Init(ormer ...orm.Ormer) *HNCardModel {
	functions.ReflectModel(&h.Field)
	h.Model = base.NewModel(h.Field.T_table, ormer...)
	return h
}

// 新增数据
func (h *HNCardModel) Insert(data map[string]interface{}) (hNCardID int, err error) {
	if result, insertErr := h.Model.Data(data).Insert(); insertErr != nil || result == 0 {
		err = errors.New("insert failed")
	} else {
		hNCardID = result
	}
	return
}

func (h *HNCardModel) Select(where map[string]interface{}, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	return h.Model.Where(where).Select()
}

// 根据HNCardID获取单条hncard数据
func (h *HNCardModel) GetByHNCardID(HNCardID int, fields ...string) (dataArray map[string]interface{}) {
	if HNCardID == 0 {
		return
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	dataArray = h.Model.Where(map[string]interface{}{
		h.Field.F_hncard_id: HNCardID,
	}).Find()
	return
}

// 批量修改
func (h *HNCardModel) UpdateByHNCardIDs(hNCardIDs []int, data map[string]interface{}) (err error) {
	if len(hNCardIDs) == 0 || len(data) == 0 {
		return
	}
	_, err = h.Model.Where(map[string]interface{}{
		h.Field.F_hncard_id: []interface{}{"IN", hNCardIDs},
	}).Data(data).Update()

	return
}

// 根据HNCardID和busId修改数据
func (h *HNCardModel) UpdateByHNCardIDsAndBusId(hNCardIDs []int, busId int, data map[string]interface{}) (err error) {
	if len(hNCardIDs) == 0 || len(data) == 0 {
		return
	}
	_, err = h.Model.Where(map[string]interface{}{
		h.Field.F_hncard_id: []interface{}{"IN", hNCardIDs},
		h.Field.F_bus_id:    busId,
	}).Data(data).Update()

	return
}

// 根据HNCardID修改数据
func (h *HNCardModel) UpdateByHNCardID(HNCardID int, data map[string]interface{}) (err error) {
	if HNCardID == 0 || len(data) == 0 {
		return
	}
	_, err = h.Model.Where(map[string]interface{}{
		h.Field.F_hncard_id: HNCardID,
	}).Data(data).Update()

	return
}

// 增加限时限次卡的销量
func (h *HNCardModel) IncrSalesByHNCardID(HNCardID int, step int) (err error) {
	if HNCardID == 0 || step <= 0 {
		return
	}
	_, err = h.Model.Where(map[string]interface{}{
		h.Field.F_hncard_id: HNCardID,
	}).Data(map[string]interface{}{
		h.Field.F_sales: []interface{}{"inc", step},
	}).Update()

	return
}

// 获取企业的限时限次卡列表
func (h *HNCardModel) GetPageByBusID(busId int, start, limit int, isGround ...int) (data []map[string]interface{}) {
	where := map[string]interface{}{
		h.Field.F_bus_id: busId,
	}
	if len(isGround) > 0 {
		where[h.Field.F_is_ground] = isGround[0]
	}
	data = h.Model.Where(where).OrderBy(h.Field.F_hncard_id+" DESC ").Limit(start, limit).Select()
	return
}

// 获取限时限次卡列表数据 过滤掉是否删除的数据(分页)
func (h *HNCardModel) GetBusHncardBusId(busId int, start, limit, isDe int) (data []map[string]interface{}) {
	where := map[string]interface{}{
		h.Field.F_bus_id: busId,
		h.Field.F_is_del: isDe,
	}
	data = h.Model.Where(where).OrderBy(h.Field.F_hncard_id+" DESC ").Limit(start, limit).Select()
	return
}

// 获取商家限时限次卡卡数量 过滤掉是否删除的数据
func (h *HNCardModel) GetNumHncardBusId(busId int, isDel int) (count int) {
	if busId <= 0 {
		return
	}
	where := map[string]interface{}{
		h.Field.F_bus_id: busId,
		h.Field.F_is_del: isDel,
	}
	count = h.Model.Where(where).Count(h.Field.F_hncard_id)
	return
}

func (h *HNCardModel) SelectHNCardsByWherePage(where []base.WhereItem, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	if limit > 0 {
		h.Model.Limit(start, limit)
	}
	return h.Model.Where(where).OrderBy(h.Field.F_hncard_id + " DESC ").Select()
}

func (h *HNCardModel) GetNumByWhere(where []base.WhereItem) int {
	return h.Model.Where(where).Count(h.Field.F_hncard_id)
}

// 获取商家的限时限次卡数量
func (h *HNCardModel) GetNumByBusID(busId int, isGround ...int) (count int) {
	if busId <= 0 {
		return
	}
	where := map[string]interface{}{
		h.Field.F_bus_id: busId,
	}
	if len(isGround) > 0 {
		where[h.Field.F_is_ground] = isGround[0]
	}
	count = h.Model.Where(where).Count(h.Field.F_hncard_id)

	return
}

// 根据hNCardIDs批量获取数据
func (h *HNCardModel) GetByHNCardIDs(hNCardIDs []int, fields ...string) []map[string]interface{} {
	if len(hNCardIDs) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hncard_id: []interface{}{"IN", hNCardIDs},
	}).Select()
}

// 根据hNCardIDs批量获取数据
func (h *HNCardModel) GetByHNCardIDsAndGround(hNCardIDs []int, fields ...string) []map[string]interface{} {
	if len(hNCardIDs) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hncard_id: []interface{}{"IN", hNCardIDs},
		h.Field.F_is_ground: 1,
	}).Select()
}

// UpdateSaleShopNum 更新在售门店数量(增加/减少)
func (h *HNCardModel) UpdateSaleShopNum(hncardIDs []int, decOrInc string /*可用数值为:dec和inc分别代表增减操作*/) bool {
	if len(hncardIDs) == 0 || len(decOrInc) == 0 {
		return false
	}
	_, err := h.Model.Where(map[string]interface{}{
		h.Field.F_hncard_id: []interface{}{"IN", hncardIDs},
	}).Data(map[string]interface{}{
		h.Field.F_sale_shop_num: []interface{}{decOrInc, addRemoveSaleShopNum},
	}).Update()
	if err != nil {
		return false
	}
	return true
}

// 根据hncardIds获取多条数据
func (h *HNCardModel) FindHNcardIdsAndBusId(hncardIds []int, busId int, fields ...string) []map[string]interface{} {
	if len(hncardIds) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hncard_id: []interface{}{"IN", hncardIds},
		h.Field.F_bus_id:    busId,
	}).Select()
}
