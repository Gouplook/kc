//CardExtModel
//2020-04-24 09:26:28

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	)

//表结构体
type CardExtModel struct {
	Model *base.Model
	Field CardExtModelField
}

//表字段
type CardExtModelField struct{
	T_table	string	`default:"card_ext"`
	F_id	string	`default:"id"`
	F_card_id	string	`default:"card_id"`
	F_notes	string	`default:"notes"`
	F_service_subscribe	string	`default:"service_subscribe"`
}

//初始化
func (c *CardExtModel) Init(ormer ...orm.Ormer) *CardExtModel{
	functions.ReflectModel(&c.Field)
	c.Model = base.NewModel(c.Field.T_table, ormer...)
	return c
}

//新增数据
func (c *CardExtModel) Insert(data map[string]interface{}) (result int, err error){
	result,err = c.Model.Data(data).Insert()
	return
}

//查询数据
func (c *CardExtModel)Find(kv map[string]interface{})(data map[string]interface{}){
	data = c.Model.Where(kv).Find()
	return
}

//更新数据
func (c *CardExtModel)Update(kv map[string]interface{}, data map[string]interface{})(result int, err error){
	result, err = c.Model.Where(kv).Data(data).Update()
	return
}