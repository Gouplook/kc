//HcardGiveDescModel
//2020-11-04 15:19:34

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	)

//表结构体
type HcardGiveDescModel struct {
	Model *base.Model
	Field HcardGiveDescModelField
}

//表字段
type HcardGiveDescModelField struct{
	T_table	string	`default:"hcard_give_desc"`
	F_id	string	`default:"id"`
	F_hcard_id	string	`default:"hcard_id"`
	F_desc	string	`default:"desc"`
}

//初始化
func (h *HcardGiveDescModel) Init(ormer ...orm.Ormer) *HcardGiveDescModel{
	functions.ReflectModel(&h.Field)
	h.Model = base.NewModel(h.Field.T_table, ormer...)
	return h
}

//新增数据
func (h *HcardGiveDescModel) Insert(data map[string]interface{}) (int){
	result,_ := h.Model.Data(data).Insert()
	return result
}

//根据Hcardid删除对应的赠品描述
func (c *HcardGiveDescModel) DelByHcardId(HcardId int) (err error) {
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_hcard_id: HcardId,
	}).Delete()

	return
}

//根据Hcardid获取对应的赠品描述
func (c *HcardGiveDescModel) GetByHcardId(hcardId int) map[string]interface{} {
	res := c.Model.Where(map[string]interface{}{
		c.Field.F_hcard_id :hcardId,
	}).Find()
	return res
}

//根据Hcardids获取赠品描述
func (c *HcardGiveDescModel)GetByHcardids(Hcardids []int) []map[string]interface{} {
	if len(Hcardids) <= 0 {
		return []map[string]interface{}{}
	}
	return c.Model.Where(map[string]interface{}{
		c.Field.F_hcard_id: []interface{}{"IN", Hcardids},
	}).Select()
}
