//HcardShopModel 卡项服务-限时卡适用门店
//2020-04-21 14:55:12

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

// HcardShopModel 表结构体
type HcardShopModel struct {
	Model *base.Model
	Field HcardShopModelField
}

// HcardShopModelField 表字段
type HcardShopModelField struct {
	T_table    string `default:"hcard_shop"`
	F_id       string `default:"id"`
	F_hcard_id string `default:"hcard_id"` // 限时卡id代表全部适用
	F_bus_id   string `default:"bus_id"`   // 所属商家企业id
	F_shop_id  string `default:"shop_id"`  // 限时卡适用门店的id,0
}

// Init 初始化
func (h *HcardShopModel) Init(ormer ...orm.Ormer) *HcardShopModel {
	functions.ReflectModel(&h.Field)
	h.Model = base.NewModel(h.Field.T_table, ormer...)
	return h
}

// Insert 新增数据
func (h *HcardShopModel) Insert(data map[string]interface{}) int {
	result, _ := h.Model.Data(data).Insert()
	return result
}

// InsertAll 批量添加数据
func (h *HcardShopModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}
	resutlt, _ := h.Model.InsertAll(data)
	return resutlt
}

// UpdateByIDs 根据ids批量修改数据
func (h *HcardShopModel) UpdateByIDs(ids []int, data map[string]interface{}) bool {
	if len(ids) == 0 || len(data) == 0 {
		return false
	}
	_, err := h.Model.Where(map[string]interface{}{
		h.Field.F_id: []interface{}{"IN", ids},
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

// GetByHcardIDs 根据限时卡ids获取所有可适应的门店记录
func (h *HcardShopModel) GetByHcardIDs(hcardIDs []int) []map[string]interface{} {
	if len(hcardIDs) == 0 {
		return []map[string]interface{}{}
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id: []interface{}{"IN", hcardIDs},
	}).Select()
}
//GetByHcardIDs 根据限时卡ids和busId获取门店添加记录
func (h *HcardShopModel) GetByHcardIDsByBusId(hcardId int, busId int ) []map[string]interface{} {
	if hcardId == 0 {
		return []map[string]interface{}{}
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id:hcardId,
		h.Field.F_bus_id:busId,
	}).Select()
}

// GetPageByShopID 根据门店id获取门店可添加的限时卡列表(即分店限时卡模块下的总部限时卡部分)
func (h *HcardShopModel) GetPageByShopID(busID, shopID, start, limit int) []map[string]interface{} {
	//if busID <= 0 || shopID <= 0 || start < 0 || limit <= 0 {
	//	return []map[string]interface{}{}
	//}
	 var whereMap =map[string]interface{}{
		 h.Field.F_bus_id: busID,
	}
	if busID <= 0 || shopID < 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}else if shopID>0{
		whereMap[h.Field.F_shop_id]=[]interface{}{"IN", []int{0, shopID}}
	}

	return h.Model.Where(whereMap).OrderBy(h.Field.F_id+" DESC ").Limit(start, limit).Select()
}

// GetByShopIDAndHcardIDs 根据hcardIDs和shopID获取适用门店的限时卡数据
func (h *HcardShopModel) GetByShopIDAndHcardIDs(busID, shopID int, hcardIDs []int) []map[string]interface{} {
	if busID == 0 || shopID == 0 || len(hcardIDs) == 0 {
		return []map[string]interface{}{}
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_bus_id:   busID,
		h.Field.F_shop_id:       []interface{}{"IN", []int{0, shopID}},
		h.Field.F_hcard_id: []interface{}{"IN", hcardIDs},
	}).Select()
}

// GetNumByBusIDAndShopID 根据门店id和总店id获取门店可添加的限时卡总数量
func (h *HcardShopModel) GetNumByBusIDAndShopID(busID, shopID int) int {
	//if busID == 0 || shopID == 0 {
	//	return 0
	//}
	var whereMap=map[string]interface{}{
		h.Field.F_bus_id:  busID,
	}
	if busID == 0 || shopID == 0 {
		return 0
	}else if shopID>0{
		whereMap[h.Field.F_shop_id]=[]interface{}{"IN", []int{0, shopID}}
	}

	return h.Model.Where(whereMap).Count(h.Field.F_id)
}

// DelByHcardIDs 批量删除
func (h *HcardShopModel) DelByHcardIDs(hcardIDs []int) bool {
	if len(hcardIDs) == 0 {
		return false
	}
	_, err := h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id: []interface{}{"IN", hcardIDs},
	}).Delete()
	if err != nil {
		return false
	}
	return true
}
