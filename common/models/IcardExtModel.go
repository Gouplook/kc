//IcardModel
//2020-08-05 15:23:57

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//IcardExtModel 表结构体
type IcardExtModel struct {
	Model     *base.Model
	Field     IcardExtModelField
	TableName string
	PK        string
}

//IcardExtModelField 表字段
type IcardExtModelField struct {
	F_id                string `default:"id"`
	F_icard_id          string `default:"icard_id"`
	F_notes             string `default:"notes"`
	F_service_subscribe string `default:"service_subscribe"`
}

//Init 初始化
func (i *IcardExtModel) Init(ormer ...orm.Ormer) *IcardExtModel {
	functions.ReflectModel(&i.Field)
	i.TableName = "icard_ext"
	i.PK = "id"
	i.Model = base.NewModel(i.TableName, ormer...)
	return i
}

//Delete 批量删除
func (i *IcardExtModel) Delete(condition Condition) int {
	i.Model.Where(condition.Where)
	result, _ := i.Model.Delete()
	return result
}

//Insert 新增数据
func (i *IcardExtModel) Insert(data map[string]interface{}) (int, error) {
	result, err := i.Model.Data(data).Insert()
	return result, err
}

//CreateOrUpdate CreateOrUpdate
func (i *IcardExtModel) CreateOrUpdate(data map[string]interface{}) (result int, err error) {
	if pk := data[i.PK]; pk != nil {
		result, err = i.Update(data, pk)
	} else {
		result, err = i.Insert(data)
	}
	return
}

//Update update
func (i *IcardExtModel) Update(data map[string]interface{}, id interface{}) (int, error) {
	result, err := i.Model.Data(data).Where([]base.WhereItem{
		{i.PK, id},
	}).Update()
	return result, err
}
