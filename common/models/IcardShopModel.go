//IcardModel
//2020-08-05 15:23:57

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

const (
	//AllShopID AllShopID
	AllShopID = 0
)

//IcardShopModel 表结构体
type IcardShopModel struct {
	Model     *base.Model
	Field     IcardShopModelField
	TableName string
	PK        string
}

//IcardShopModelField 表字段
type IcardShopModelField struct {
	F_id       string `default:"id"`
	F_icard_id string `default:"icard_id"`
	F_bus_id   string `default:"bus_id"`
	F_shop_id  string `default:"shop_id"`
}

//Init 初始化
func (i *IcardShopModel) Init(ormer ...orm.Ormer) *IcardShopModel {
	functions.ReflectModel(&i.Field)
	i.TableName = "icard_shop"
	i.PK = "id"
	i.Model = base.NewModel(i.TableName, ormer...)
	return i
}

//Delete 批量删除
func (i *IcardShopModel) Delete(condition Condition) int {
	i.Model.Where(condition.Where)
	result, _ := i.Model.Delete()
	return result
}

//GetAll GetAll
func (i *IcardShopModel) GetAll(condition Condition, fields ...string) (data []map[string]interface{}) {
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
func (i *IcardShopModel) InsertAll(data []map[string]interface{}) int {
	result, _ := i.Model.InsertAll(data)
	return result
}

//Insert 新增数据
func (i *IcardShopModel) Insert(data map[string]interface{}) int {
	result, _ := i.Model.Data(data).Insert()
	return result
}

//CreateOrUpdate CreateOrUpdate
func (i *IcardShopModel) CreateOrUpdate(data map[string]interface{}) (result int) {
	if pk := data[i.PK]; pk != nil {
		result = i.Update(data, pk)
	} else {
		result = i.Insert(data)
	}
	return
}

//Update update
func (i *IcardShopModel) Update(data map[string]interface{}, id interface{}) int {
	result, _ := i.Model.Data(data).Where([]base.WhereItem{
		{i.PK, id},
	}).Update()
	return result
}

//根据身份卡id 获取所有可适用门店记录
func (i *IcardShopModel) GetByIcardId(icardId int) []map[string]interface{} {
	if icardId <= 0 {
		return []map[string]interface{}{}
	}

	return i.Model.Where(map[string]interface{}{
		i.Field.F_icard_id: icardId,
	}).Select()
}

//根据身份卡id and busId 获取所有可适用门店记录
func (i *IcardShopModel)GetByIcardIdAndBusId(icardId int ,busId int )[]map[string]interface{}{
	if icardId <= 0 {
		return []map[string]interface{}{}
	}
	return i.Model.Where([]base.WhereItem{
		{i.Field.F_icard_id,icardId},
		{i.Field.F_bus_id,busId},
	}).Select()
}

func (i *IcardShopModel) GetByIcardIds(icardIds []int) []map[string]interface{} {
	if len(icardIds) <= 0 {
		return []map[string]interface{}{}
	}

	return i.Model.Where(map[string]interface{}{
		i.Field.F_icard_id: []interface{}{"IN", icardIds},
	}).Select()
}
