//IcardModel
//2020-08-05 15:23:57

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//IcardSingleModel 表结构体
type IcardSingleModel struct {
	Model     *base.Model
	Field     IcardSingleModelField
	TableName string
	PK        string
}

//IcardSingleModelField 表字段
type IcardSingleModelField struct {
	F_id        string `default:"id"`
	F_icard_id  string `default:"icard_id"`
	F_single_id string `default:"single_id"`
	F_discount  string `default:"discount"`
}

//Init 初始化
func (i *IcardSingleModel) Init(ormer ...orm.Ormer) *IcardSingleModel {
	functions.ReflectModel(&i.Field)
	i.TableName = "icard_single"
	i.PK = "id"
	i.Model = base.NewModel(i.TableName, ormer...)
	return i
}

//GetAll GetAll
func (i *IcardSingleModel) GetAll(condition Condition, fields ...string) (data []map[string]interface{}) {
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

// 根据IcardIds获取身份卡信息
func (i *IcardSingleModel) GetByIcardIds(icardIds []int, field ...string) []map[string]interface{} {
	if len(icardIds) == 0 {
		return nil
	}
	if len(field) > 0 {
		i.Model.Field(field)
	}
	return i.Model.Where([]base.WhereItem{{i.Field.F_icard_id, []interface{}{"IN", icardIds}}}).OrderBy(i.Field.F_discount+" asc ").Select()
}


//Delete 批量删除
func (i *IcardSingleModel) Delete(condition Condition) int {
	i.Model.Where(condition.Where)
	result, _ := i.Model.Delete()
	return result
}

//InsertAll 批量新增数据
func (i *IcardSingleModel) InsertAll(data []map[string]interface{}) int {
	result, _ := i.Model.InsertAll(data)
	return result
}

//Insert 新增数据
func (i *IcardSingleModel) Insert(data map[string]interface{}) int {
	result, _ := i.Model.Data(data).Insert()
	return result
}

//CreateOrUpdate CreateOrUpdate
func (i *IcardSingleModel) CreateOrUpdate(data map[string]interface{}) (result int) {
	if pk := data[i.PK]; pk != nil {
		result = i.Update(data, pk)
	} else {
		result = i.Insert(data)
	}
	return
}

//Update update
func (i *IcardSingleModel) Update(data map[string]interface{}, id interface{}) int {
	result, _ := i.Model.Data(data).Where([]base.WhereItem{
		{i.PK, id},
	}).Update()
	return result
}

func (i *IcardSingleModel) SelectByPage(where map[string]interface{}, start, limit int) []map[string]interface{} {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return i.Model.Where(where).Limit(start, limit).OrderBy("id DESC ").Select()
}

func (s *IcardSingleModel) Find(where map[string]interface{}) map[string]interface{} {
	if len(where) <= 0 {
		return map[string]interface{}{}
	}

	return s.Model.Where(where).OrderBy("id DESC ").Find()
}

func (s *IcardSingleModel) GetTotalNum(where map[string]interface{}) int {
	if len(where) == 0 {
		return 0
	}
	return s.Model.Where(where).Count()
}
