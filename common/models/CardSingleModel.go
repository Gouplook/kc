//CardSingleModel
//2020-04-24 09:26:28

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"github.com/wendal/errors"
)

//表结构体
type CardSingleModel struct {
	Model *base.Model
	Field CardSingleModelField
}

//表字段
type CardSingleModelField struct {
	T_table     string `default:"card_single"`
	F_id        string `default:"id"`
	F_card_id   string `default:"card_id"`
	F_single_id string `default:"single_id"`
}

//初始化
func (c *CardSingleModel) Init(ormer ...orm.Ormer) *CardSingleModel {
	functions.ReflectModel(&c.Field)
	c.Model = base.NewModel(c.Field.T_table, ormer...)
	return c
}

//新增数据
func (c *CardSingleModel) Insert(data map[string]interface{}) (err error) {
	if result, insertErr := c.Model.Data(data).Insert(); insertErr != nil || result == 0 {
		err = errors.New("insert failed")
	}
	return
}

//批量添加
func (c *CardSingleModel) InsertAll(data []map[string]interface{}) (err error) {
	if len(data) == 0 {
		return
	}
	_, err = c.Model.InsertAll(data)
	return
}

//获取card包含的单项目
func (c *CardSingleModel) GetByCardID(cardID int) (data []map[string]interface{}) {
	if cardID == 0 {
		return []map[string]interface{}{}
	}
	data = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: cardID,
	}).Select()

	return
}

func (s *CardSingleModel)SelectByPage(where map[string]interface{},start,limit int)[]map[string]interface{}  {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(where).OrderBy(s.Field.F_id+" DESC ").Limit(start,limit).Select()
}

func (s *CardSingleModel)Find(where map[string]interface{})map[string]interface{}  {
	if len(where) <= 0 {
		return map[string]interface{}{}
	}

	return s.Model.Where(where).OrderBy(s.Field.F_id+" DESC ").Find()
}

func (s *CardSingleModel)GetTotalNum(where map[string]interface{})int  {
	if len(where)==0{
		return 0
	}
	return s.Model.Where(where).Count(s.Field.F_id)
}

//删除包含的单项目
func (c *CardSingleModel) DelByIds(ids []int) (err error) {
	if len(ids) == 0 {
		return
	}
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_id: []interface{}{"IN", ids},
	}).Delete()

	return
}

//根据cardIds批量获取单项目
func (c *CardSingleModel) GetByCardIds(cardIds []int) (data []map[string]interface{}) {
	if len(cardIds) == 0 {
		return []map[string]interface{}{}
	}
	data = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: []interface{}{"IN", cardIds},
	}).Select()

	return
}
