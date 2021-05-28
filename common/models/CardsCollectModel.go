//CardsCollectModel
//2020-09-22 16:04:34

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type CardsCollectModel struct {
	Model *base.Model
	Field CardsCollectModelField
}

//表字段
type CardsCollectModelField struct {
	T_table     string `default:"cards_collect"`
	F_id        string `default:"id"`
	F_item_id   string `default:"item_id"`
	F_item_type string `default:"item_type"`
	F_bus_id    string `default:"bus_id"`
	F_ss_id    string `default:"ss_id"`
	F_shop_id   string `default:"shop_id"`
	F_uid       string `default:"uid"`
	F_ctime     string `default:"ctime"`
}

//初始化
func (c *CardsCollectModel) Init(ormer ...orm.Ormer) *CardsCollectModel {
	functions.ReflectModel(&c.Field)
	c.Model = base.NewModel(c.Field.T_table, ormer...)
	return c
}

//新增数据
func (c *CardsCollectModel) Insert(data map[string]interface{}) int {
	result, _ := c.Model.Data(data).Insert()
	return result
}

func (c *CardsCollectModel) Delete(where map[string]interface{}) bool {
	if len(where) == 0 {
		return false
	}
	result, _ := c.Model.Where(where).Delete()
	if result == 0 {
		return false
	}
	return true
}

func (c *CardsCollectModel) Find(where map[string]interface{}, fields ...string) map[string]interface{} {
	if len(where) == 0 {
		return make(map[string]interface{})
	}
	if len(fields) > 0 {
		c.Model.Field(fields)
	}
	return c.Model.Where(where).Find()
}

func (c *CardsCollectModel) SelectByPage(where map[string]interface{}, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0  {
		return make([]map[string]interface{}, 0)
	}
	if len(fields)>0{
		c.Model.Field(fields)
	}
	return c.Model.Where(where).OrderBy(c.Field.F_id+" DESC ").Limit(start, limit).Select()
}
func (c *CardsCollectModel)GetTotalNum(where map[string]interface{})int  {
	if len(where) == 0{
		return 0
	}
	return c.Model.Where(where).Count(c.Field.F_id)
}