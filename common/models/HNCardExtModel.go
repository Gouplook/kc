//HNCardExtModel
//2020-04-23 16:26:23

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	)

//表结构体
type HNCardExtModel struct {
	Model *base.Model
	Field HNCardExtModelField
}

//表字段
type HNCardExtModelField struct{
	T_table	string	`default:"hncard_ext"`
	F_id	string	`default:"id"`
	F_hncard_id	string	`default:"hncard_id"`
	F_notes	string	`default:"notes"`
	F_service_subscribe	string	`default:"service_subscribe"`
}

//初始化
func (h *HNCardExtModel) Init(ormer ...orm.Ormer) *HNCardExtModel {
	functions.ReflectModel(&h.Field)
	h.Model = base.NewModel(h.Field.T_table, ormer...)
	return h
}

//新增数据
func (h *HNCardExtModel) Insert(data map[string]interface{}) (result int, err error){
	result,err = h.Model.Data(data).Insert()
	return
}

//查询数据
func (h *HNCardExtModel)Find(kv map[string]interface{})(data map[string]interface{}){
	data = h.Model.Where(kv).Find()
	return
}

//更新数据
func (h *HNCardExtModel)Update(kv map[string]interface{}, data map[string]interface{})(result int, err error){
	result, err = h.Model.Where(kv).Data(data).Update()
	return
}