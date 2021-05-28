//IcardModel
//2020-08-05 15:23:57

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//IcardGiveModel 表结构体
type IcardGiveModel struct {
	Model     *base.Model
	Field     IcardGiveModelField
	TableName string
	PK        string
}

//IcardGiveModelField 表字段
type IcardGiveModelField struct {
	F_id                string `default:"id"`
	F_icard_id          string `default:"icard_id"`
	F_single_id string `default:"single_id"`
	F_num       string `default:"num"`
}

//Init 初始化
func (i *IcardGiveModel) Init(ormer ...orm.Ormer) *IcardGiveModel {
	functions.ReflectModel(&i.Field)
	i.TableName = "icard_give"
	i.PK = "id"
	i.Model = base.NewModel(i.TableName, ormer...)
	return i
}

//Delete 批量删除
func (i *IcardGiveModel) Delete(condition Condition) int {
	i.Model.Where(condition.Where)
	result, _ := i.Model.Delete()
	return result
}

//GetAll GetAll
func (i *IcardGiveModel) GetAll(condition Condition, fields ...string) (data []map[string]interface{}) {
	if len(fields) > 0 {
		i.Model.Field(fields)
	}
	if len(condition.Order) > 0 {
		i.Model.OrderBy(condition.Order)
	}
	if condition.Limit > 0 {
		i.Model.Limit(condition.Offset, condition.Limit)
	}
	i.Model.Where(condition.Where)
	return i.Model.Select()
}

//InsertAll 批量新增数据
func (i *IcardGiveModel) InsertAll(data []map[string]interface{}) int {
	result, _ := i.Model.InsertAll(data)
	return result
}

//Insert 新增数据
func (i *IcardGiveModel) Insert(data map[string]interface{}) int {
	result, _ := i.Model.Data(data).Insert()
	return result
}

//CreateOrUpdate CreateOrUpdate
func (i *IcardGiveModel) CreateOrUpdate(data map[string]interface{}) (result int) {
	if pk := data[i.PK]; pk != nil {
		result = i.Update(data, pk)
	} else {
		result = i.Insert(data)
	}
	return
}

//Update update
func (i *IcardGiveModel) Update(data map[string]interface{}, id interface{}) int {
	result, _ := i.Model.Data(data).Where([]base.WhereItem{
		{i.PK, id},
	}).Update()
	return result
}

func (i *IcardGiveModel) SelectByPage(where []base.WhereItem, start, limit int) []map[string]interface{} {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return i.Model.Where(where).OrderBy(i.Field.F_id+" DESC ").Limit(start, limit).Select()
}

func (i *IcardGiveModel) Count(where []base.WhereItem) int {
	if len(where) == 0 {
		return 0
	}
	return i.Model.Where(where).Count()
}
