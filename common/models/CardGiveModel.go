//CardGiveModel
//2020-04-24 09:26:28

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"github.com/wendal/errors"
)

//表结构体
type CardGiveModel struct {
	Model *base.Model
	Field CardGiveModelField
}

//表字段
type CardGiveModelField struct {
	T_table     string `default:"card_give"`
	F_id        string `default:"id"`
	F_card_id   string `default:"card_id"`
	F_single_id string `default:"single_id"`
	F_num       string `default:"num"`
	F_period_of_validity string `default:"period_of_validity"`
}

//初始化
func (c *CardGiveModel) Init(ormer ...orm.Ormer) *CardGiveModel {
	functions.ReflectModel(&c.Field)
	c.Model = base.NewModel(c.Field.T_table, ormer...)
	return c
}

//新增数据
func (c *CardGiveModel) Insert(data map[string]interface{}) (err error) {
	if result, insertErr := c.Model.Data(data).Insert(); insertErr != nil || result == 0 {
		err = errors.New("insert failed")
	}
	return
}

//批量添加
func (c *CardGiveModel) InsertAll(data []map[string]interface{}) (err error) {
	_, err = c.Model.InsertAll(data)
	return
}

//获取card赠送的单项目
func (c *CardGiveModel) GetByCardID(cardID int) (dataArray []map[string]interface{}) {
	if cardID == 0 {
		return []map[string]interface{}{}
	}

	dataArray = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: cardID,
	}).Select()
	return
}

//删除赠送的项目
func (c *CardGiveModel) DelByIds(ids []int) (err error) {
	if len(ids) == 0 {
		return
	}
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_id: []interface{}{"IN", ids},
	}).Delete()

	return
}

//更新数量
func (c *CardGiveModel) UpdateNumById(id int, num int) (err error) {
	if id == 0 || num <= 0 {
		return
	}
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_id: id,
	}).Data(map[string]interface{}{
		c.Field.F_num: num,
	}).Update()

	return
}

//获取cardIds批量获取赠送的单项目
func (c *CardGiveModel) GetByCardIds(cardIds []int) (dataArray []map[string]interface{}) {
	if len(cardIds) == 0 {
		return []map[string]interface{}{}
	}

	dataArray = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: []interface{}{"IN", cardIds},
	}).Select()
	return
}

func (c *CardGiveModel) SelectByPage(where []base.WhereItem, start, limit int) []map[string]interface{} {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return c.Model.Where(where).OrderBy(c.Field.F_id+" DESC ").Limit(start, limit).Select()
}

func (c *CardGiveModel) Count(where []base.WhereItem) int {
	if len(where) == 0 {
		return 0
	}
	return c.Model.Where(where).Count()
}