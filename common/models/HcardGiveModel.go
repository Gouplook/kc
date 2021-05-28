//HcardGiveModel 卡项-限时卡赠送的单服务表
//2020-04-21 14:55:12

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

// HcardGiveModel 表结构体
type HcardGiveModel struct {
	Model *base.Model
	Field HcardGiveModelField
}

//HcardGiveModelField 表字段
type HcardGiveModelField struct {
	T_table              string `default:"hcard_give"`
	F_id                 string `default:"id"`
	F_hcard_id           string `default:"hcard_id"`
	F_single_id          string `default:"single_id"`
	F_num                string `default:"num"`
	F_period_of_validity string `default:"period_of_validity"`
}

// Init 初始化
func (h *HcardGiveModel) Init(ormer ...orm.Ormer) *HcardGiveModel {
	functions.ReflectModel(&h.Field)
	h.Model = base.NewModel(h.Field.T_table, ormer...)
	return h
}

//Insert 新增数据
func (h *HcardGiveModel) Insert(data map[string]interface{}) int {
	result, _ := h.Model.Data(data).Insert()
	return result
}

// InsertAll 批量新增数据
func (h *HcardGiveModel) InsertAll(data []map[string]interface{}) int {
	result, _ := h.Model.InsertAll(data)
	return result
}

// UpdateNumByID 根据id更新hcard赠送的单项目数量
func (h *HcardGiveModel) UpdateNumByID(id int, num int) bool {
	if id == 0 || num <= 0 {
		return false
	}
	_, err := h.Model.Where(map[string]interface{}{
		h.Field.F_id: id,
	}).Data(map[string]interface{}{
		h.Field.F_num: num,
	}).Update()
	if err != nil {
		return false
	}
	return true
}

// DeleteByIDs 根据ids批量删除赠送的项目
func (h *HcardGiveModel) DeleteByIDs(ids []int) bool {
	if len(ids) == 0 {
		return false
	}
	_, err := h.Model.Where(map[string]interface{}{
		h.Field.F_id: []interface{}{"IN", ids},
	}).Delete()
	if err != nil {
		return false
	}
	return true
}

// GetByHcardID 根据hcard id获取hcard赠送的单项目
func (h *HcardGiveModel) GetByHcardID(hcardid int) []map[string]interface{} {
	if hcardid <= 0 {
		return []map[string]interface{}{}
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id: hcardid,
	}).Select()
}

func (h *HcardGiveModel) GetByHcardIds(hcardids []int) []map[string]interface{} {
	if len(hcardids) <= 0 {
		return []map[string]interface{}{}
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id: []interface{}{"IN", hcardids},
	}).Select()
}

func (h *HcardGiveModel) SelectByPage(where []base.WhereItem, start, limit int) []map[string]interface{} {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return h.Model.Where(where).OrderBy(h.Field.F_id+" DESC ").Limit(start, limit).Select()
}

func (h *HcardGiveModel) Count(where []base.WhereItem) int {
	if len(where) == 0 {
		return 0
	}
	return h.Model.Where(where).Count()
}
