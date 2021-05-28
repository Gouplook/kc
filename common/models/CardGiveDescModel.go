//CardGiveDescModel
//2020-11-04 15:19:34

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	)

//表结构体
type CardGiveDescModel struct {
	Model *base.Model
	Field CardGiveDescModelField
}

//表字段
type CardGiveDescModelField struct{
	T_table	string	`default:"card_give_desc"`
	F_id	string	`default:"id"`
	F_card_id	string	`default:"card_id"`
	F_desc	string	`default:"desc"`
}

//初始化
func (c *CardGiveDescModel) Init(ormer ...orm.Ormer) *CardGiveDescModel{
	functions.ReflectModel(&c.Field)
	c.Model = base.NewModel(c.Field.T_table, ormer...)
	return c
}

//新增数据
func (c *CardGiveDescModel) Insert(data map[string]interface{}) (int){
	result,_ := c.Model.Data(data).Insert()
	return result
}

//根据cardid删除对应的赠品描述
func (c *CardGiveDescModel) DelByCardId(CardId int) (err error) {
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: CardId,
	}).Delete()

	return
}

//根据cardid获取对应的赠品描述
func (c *CardGiveDescModel) GetByCardId(CardId int) map[string]interface{} {
	res := c.Model.Where(map[string]interface{}{
		c.Field.F_card_id :CardId,
	}).Find()
	return res
}


//根据cardids获取赠品描述
func (c *CardGiveDescModel)GetByCardids(CardId []int) []map[string]interface{} {
	if len(CardId) <= 0 {
		return []map[string]interface{}{}
	}
	return c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: []interface{}{"IN", CardId},
	}).Select()
}
