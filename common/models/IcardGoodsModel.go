//IcardModel
//2020-08-05 15:23:57

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//IcardGoodsModel 表结构体
type IcardGoodsModel struct {
	Model     *base.Model
	Field     IcardGoodsModelField
	TableName string
	PK        string
}

//IcardGoodsModelField 表字段
type IcardGoodsModelField struct {
	F_id       string `default:"id"`
	F_icard_id string `default:"icard_id"`
	F_goods_id string `default:"goods_id"`
	F_discount string `default:"discount"`
}

//Init 初始化
func (i *IcardGoodsModel) Init(ormer ...orm.Ormer) *IcardGoodsModel {
	functions.ReflectModel(&i.Field)
	i.TableName = "icard_goods"
	i.PK = "id"
	i.Model = base.NewModel(i.TableName, ormer...)
	return i
}

//Delete 批量删除
func (i *IcardGoodsModel) Delete(condition Condition) int {
	i.Model.Where(condition.Where)
	result, _ := i.Model.Delete()
	return result
}

//GetAll GetAll
func (i *IcardGoodsModel) GetAll(condition Condition, fields ...string) (data []map[string]interface{}) {
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

func (i *IcardGoodsModel) GetByIcardIds(icardIds []int, field ...string) []map[string]interface{} {
	if len(icardIds) == 0 {
		return nil
	}
	if len(field) > 0 {
		i.Model.Field(field)
	}
	return i.Model.Where([]base.WhereItem{{i.Field.F_icard_id, []interface{}{"IN", icardIds}}}).Select()
}

//InsertAll 批量新增数据
func (i *IcardGoodsModel) InsertAll(data []map[string]interface{}) int {
	result, _ := i.Model.InsertAll(data)
	return result
}

//Insert 新增数据
func (i *IcardGoodsModel) Insert(data map[string]interface{}) int {
	result, _ := i.Model.Data(data).Insert()
	return result
}

//CreateOrUpdate CreateOrUpdate
func (i *IcardGoodsModel) CreateOrUpdate(data map[string]interface{}) (result int) {
	if pk := data[i.PK]; pk != nil {
		result = i.Update(data, pk)
	} else {
		result = i.Insert(data)
	}
	return
}

//Update update
func (i *IcardGoodsModel) Update(data map[string]interface{}, id interface{}) int {
	result, _ := i.Model.Data(data).Where([]base.WhereItem{
		{i.PK, id},
	}).Update()
	return result
}

func (i *IcardGoodsModel) GetTotalNum(where map[string]interface{}) int {
	if len(where) == 0 {
		return 0
	}
	return i.Model.Where(where).Count()
}
