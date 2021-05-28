//IcardGiveDescModel
//2020-11-04 15:19:34

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	)

//表结构体
type IcardGiveDescModel struct {
	Model *base.Model
	Field IcardGiveDescModelField
}

//表字段
type IcardGiveDescModelField struct{
	T_table	string	`default:"icard_give_desc"`
	F_id	string	`default:"id"`
	F_icard_id	string	`default:"icard_id"`
	F_desc	string	`default:"desc"`
}

//初始化
func (i *IcardGiveDescModel) Init(ormer ...orm.Ormer) *IcardGiveDescModel{
	functions.ReflectModel(&i.Field)
	i.Model = base.NewModel(i.Field.T_table, ormer...)
	return i
}

//新增数据
func (i *IcardGiveDescModel) Insert(data map[string]interface{}) (int){
	result,_ := i.Model.Data(data).Insert()
	return result
}

//Delete 批量删除
func (i *IcardGiveDescModel) Delete(condition Condition) int {
	i.Model.Where(condition.Where)
	result, _ := i.Model.Delete()
	return result
}

//根据Icardid删除对应的赠品描述
func (c *IcardGiveDescModel) DelByicardId(icardId int) (err error) {
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_icard_id: icardId,
	}).Delete()

	return
}

//根据Icardid获取对应的赠品描述
func (c *IcardGiveDescModel) GetByicardId(icardId int) map[string]interface{} {
	res := c.Model.Where(map[string]interface{}{
		c.Field.F_icard_id :icardId,
	}).Find()
	return res
}

//根据Hncardids获取赠品描述
func (c *IcardGiveDescModel)GetByIcardids(Icardids []int) []map[string]interface{} {
	if len(Icardids) <= 0 {
		return []map[string]interface{}{}
	}
	return c.Model.Where(map[string]interface{}{
		c.Field.F_icard_id: []interface{}{"IN", Icardids},
	}).Select()
}
