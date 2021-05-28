//CardGoodsModel
//2020-05-08 08:58:34

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type CardGoodsModel struct {
	Model *base.Model
	Field CardGoodsModelField
}

//表字段
type CardGoodsModelField struct {
	T_table      string `default:"card_goods"`
	F_id         string `default:"id"`
	F_card_id    string `default:"card_id"`
	F_product_id string `default:"product_id"`
}

//初始化
func (c *CardGoodsModel) Init(ormer ...orm.Ormer) *CardGoodsModel {
	functions.ReflectModel(&c.Field)
	c.Model = base.NewModel(c.Field.T_table, ormer...)
	return c
}

//新增数据
func (c *CardGoodsModel) Insert(data map[string]interface{}) (result int, err error) {
	result, err = c.Model.Data(data).Insert()
	return
}

//批量增加数据
func (c *CardGoodsModel) InsertAll(data []map[string]interface{}) (result int, err error) {
	result, err = c.Model.InsertAll(data)
	return
}

//获取card包含的商品
func (c *CardGoodsModel) GetByCardID(cardID int) (data []map[string]interface{}) {
	if cardID == 0 {
		return []map[string]interface{}{}
	}
	data = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: cardID,
	}).Select()

	return
}

func (c *CardGoodsModel) DelByIds(ids []int) (err error) {
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_id: []interface{}{"IN", ids},
	}).Delete()
	return
}

//获取cards包含的商品
func (c *CardGoodsModel) GetByCardIds(cardIds []int) (data []map[string]interface{}) {
	if len(cardIds) == 0 {
		return []map[string]interface{}{}
	}
	data = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: []interface{}{"IN", cardIds},
	}).Select()

	return
}

//获取card包含的商品,带分页
func (c *CardGoodsModel) GetByCardIdPage(cardID,start,pageSize int,) (data []map[string]interface{}) {
	if cardID == 0 {
		return []map[string]interface{}{}
	}
	data = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: cardID,
	}).Limit(start, pageSize).Select()

	return
}

//根据cardId获取总数
func (c *CardGoodsModel) CountByCardId(cardId int) int {
	count := c.Model.Where(map[string]interface{}{
		c.Field.F_card_id:cardId,
	}).Count("*")
	return count
}

func (s *CardGoodsModel)Find(where map[string]interface{})map[string]interface{}  {
	if len(where) <= 0 {
		return map[string]interface{}{}
	}

	return s.Model.Where(where).OrderBy(s.Field.F_id+" DESC ").Find()
}