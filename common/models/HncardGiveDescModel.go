//HncardGiveDescModel
//2020-11-04 15:19:34

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	)

//表结构体
type HncardGiveDescModel struct {
	Model *base.Model
	Field HncardGiveDescModelField
}

//表字段
type HncardGiveDescModelField struct{
	T_table	string	`default:"hncard_give_desc"`
	F_id	string	`default:"id"`
	F_hncard_id	string	`default:"hncard_id"`
	F_desc	string	`default:"desc"`
}

//初始化
func (h *HncardGiveDescModel) Init(ormer ...orm.Ormer) *HncardGiveDescModel{
	functions.ReflectModel(&h.Field)
	h.Model = base.NewModel(h.Field.T_table, ormer...)
	return h
}

//新增数据
func (h *HncardGiveDescModel) Insert(data map[string]interface{}) (int){
	result,_ := h.Model.Data(data).Insert()
	return result
}


//根据Hcardid删除对应的赠品描述
func (c *HncardGiveDescModel) DelByHncardId(hncardId int) (err error) {
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_hncard_id: hncardId,
	}).Delete()

	return
}

//根据Hcardid获取对应的赠品描述
func (c *HncardGiveDescModel) GetByHncardId(hncardId int) map[string]interface{} {
	res := c.Model.Where(map[string]interface{}{
		c.Field.F_hncard_id :hncardId,
	}).Find()
	return res
}


//根据Hncardids获取赠品描述
func (c *HncardGiveDescModel)GetByHncardids(Hncardids []int) []map[string]interface{} {
	if len(Hncardids) <= 0 {
		return []map[string]interface{}{}
	}
	return c.Model.Where(map[string]interface{}{
		c.Field.F_hncard_id: []interface{}{"IN", Hncardids},
	}).Select()
}
