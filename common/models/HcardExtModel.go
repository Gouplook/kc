//HcardExtModel 卡项管理-限时卡扩展信息表
//2020-04-26 15:10:11

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	)

//表结构体
type HcardExtModel struct {
	Model *base.Model
	Field HcardExtModelField
}

//表字段
type HcardExtModelField struct{
	T_table	string	`default:"hcard_ext"`
	F_id	string	`default:"id"`
	F_hcard_id	string	`default:"hcard_id"`
	F_notes	string	`default:"notes"`// 购卡须知 json数据
	F_service_subscribe	string	`default:"service_subscribe"`// 预约要求
}

// Init 初始化
func (h *HcardExtModel) Init(ormer ...orm.Ormer) *HcardExtModel{
	functions.ReflectModel(&h.Field)
	h.Model = base.NewModel(h.Field.T_table, ormer...)
	return h
}

// Insert 新增数据
func (h *HcardExtModel) Insert(data map[string]interface{}) (int){
	result,_ := h.Model.Data(data).Insert()
	return result
}
// UpdateByHcardID 根据现时卡id更新数据
func (h *HcardExtModel)UpdateByHcardID(hcardID int,data map[string]interface{})bool{
	if _, err :=h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id:hcardID,
	}).Data(data).Update();err!=nil{
		return false
	}
	return  true
}
// GetByHcardID 根据现时卡id获取数据
func (h *HcardExtModel)GetByHcardID(hcarID int)map[string]interface{}{
	if hcarID<=0{
		return  map[string]interface{}{}
	}
	return  h.Model.Where(map[string]interface{}{
		h.Field.F_hcard_id:hcarID,
	}).Find()
}