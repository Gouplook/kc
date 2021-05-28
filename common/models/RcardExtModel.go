//RcardExtModel
//2020-10-20 17:18:26

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type RcardExtModel struct {
	Model *base.Model
	Field RcardExtModelField
}

//表字段
type RcardExtModelField struct {
	T_table             string `default:"rcard_ext"`
	F_id                string `default:"id"`
	F_rcard_id          string `default:"rcard_id"`
	F_notes             string `default:"notes"`
	F_service_subscribe string `default:"service_subscribe"`
}

//初始化
func (r *RcardExtModel) Init(ormer ...orm.Ormer) *RcardExtModel {
	functions.ReflectModel(&r.Field)
	r.Model = base.NewModel(r.Field.T_table, ormer...)
	return r
}

//新增数据
func (r *RcardExtModel) Insert(data map[string]interface{}) int {
	result, _ := r.Model.Data(data).Insert()
	return result
}
//更新数据
func (r *RcardExtModel)Update(kv map[string]interface{}, data map[string]interface{})(result int, err error){
	result, err = r.Model.Where(kv).Data(data).Update()
	return
}
//新增Ext数据
func (r *RcardExtModel) InsertExt(data map[string]interface{}) (result int, err error){
	result,err = r.Model.Data(data).Insert()
	return
}

//查询数据
func (r *RcardExtModel)Find(kv map[string]interface{})(data map[string]interface{}){
	data = r.Model.Where(kv).Find()
	return
}

//查询数据
func (r *RcardExtModel) GetByRcardid(rcardId int) (data map[string]interface{}) {
	data = r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Find()
	return
}

//更新数据
func (r *RcardExtModel) UpdateByRcardid(rcardId int, data map[string]interface{}) bool {
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
