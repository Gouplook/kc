//HcardSingleModel 卡项服务-限时卡包含的单项目
//2020-04-21 14:55:12

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type HcardSingleModel struct {
	Model *base.Model
	Field HcardSingleModelField
}

//表字段
type HcardSingleModelField struct {
	T_table     string `default:"hcard_single"`
	F_id        string `default:"id"`
	F_hcard_id  string `default:"hcard_id"`
	F_single_id string `default:"single_id"`
}

// Init 初始化
func (h *HcardSingleModel) Init(ormer ...orm.Ormer) *HcardSingleModel {
	functions.ReflectModel(&h.Field)
	h.Model = base.NewModel(h.Field.T_table, ormer...)
	return h
}

//Insert 新增数据
func (h *HcardSingleModel) Insert(data map[string]interface{}) int {
	result, _ := h.Model.Data(data).Insert()
	return result
}

// InsertAll 批量添加限时卡单项目
func (h *HcardSingleModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}
	result, _ := h.Model.InsertAll(data)
	return result
}

// GetByHcardID 根据hcard id获取限时卡包含的单项目
func (h *HcardSingleModel) GetByHcardID(hcardID int) []map[string]interface{} {
	if hcardID == 0 {
		return []map[string]interface{}{}
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id: hcardID,
	}).Select()
}

func (s *HcardSingleModel)SelectByPage(where map[string]interface{},start,limit int)[]map[string]interface{}  {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(where).OrderBy(s.Field.F_id+" DESC ").Limit(start,limit).Select()
}

func (s *HcardSingleModel)GetTotalNum(where map[string]interface{})int  {
	if len(where)==0{
		return 0
	}
	return s.Model.Where(where).Count(s.Field.F_id)
}

func (h *HcardSingleModel) GetByHcardIds(hcardIds []int) []map[string]interface{} {
	if len(hcardIds) == 0 {
		return []map[string]interface{}{}
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id: []interface{}{"IN", hcardIds},
	}).Select()
}

// DeleteByIDs 删除限时卡包含的单项目
func (h *HcardSingleModel) DeleteByIDs(ids []int) bool {
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
