//NCardExtModel
//2020-04-23 16:25:58

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	)

//表结构体
type NCardExtModel struct {
	Model *base.Model
	Field NCardExtModelField
}

//表字段
type NCardExtModelField struct{
	T_table	string	`default:"ncard_ext"`
	F_id	string	`default:"id"`
	F_ncard_id	string	`default:"ncard_id"`
	F_notes	string	`default:"notes"`
	F_service_subscribe	string	`default:"service_subscribe"`
}

//初始化
func (n *NCardExtModel) Init(ormer ...orm.Ormer) *NCardExtModel {
	functions.ReflectModel(&n.Field)
	n.Model = base.NewModel(n.Field.T_table, ormer...)
	return n
}

//新增数据
func (n *NCardExtModel) Insert(data map[string]interface{}) (result int, err error){
	result,err = n.Model.Data(data).Insert()
	return
}

//查询数据
func (n *NCardExtModel)Find( kv map[string]interface{})(data map[string]interface{}){
	data = n.Model.Where(kv).Find()
	return
}

//更新数据
func (n *NCardExtModel)Update(kv map[string]interface{}, data map[string]interface{})(result int, err error){
	result, err = n.Model.Where(kv).Data(data).Update()
	return
}