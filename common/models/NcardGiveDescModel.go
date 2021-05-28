//NcardGiveDescModel
//2020-11-04 15:19:34

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	)

//表结构体
type NcardGiveDescModel struct {
	Model *base.Model
	Field NcardGiveDescModelField
}

//表字段
type NcardGiveDescModelField struct{
	T_table	string	`default:"ncard_give_desc"`
	F_id	string	`default:"id"`
	F_ncard_id	string	`default:"ncard_id"`
	F_desc	string	`default:"desc"`
}

//初始化
func (n *NcardGiveDescModel) Init(ormer ...orm.Ormer) *NcardGiveDescModel{
	functions.ReflectModel(&n.Field)
	n.Model = base.NewModel(n.Field.T_table, ormer...)
	return n
}

//新增数据
func (n *NcardGiveDescModel) Insert(data map[string]interface{}) (int){
	result,_ := n.Model.Data(data).Insert()
	return result
}

//根据Hcardid删除对应的赠品描述
func (n *NcardGiveDescModel) DelByNcardId(ncardId int) (err error) {
	_, err = n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id: ncardId,
	}).Delete()

	return
}

//根据Hcardid获取对应的赠品描述
func (n *NcardGiveDescModel) GetByNcardId(ncardId int) map[string]interface{} {
	res := n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id :ncardId,
	}).Find()
	return res
}

//根据Ncardids获取赠品描述
func (n *NcardGiveDescModel)GetByNcardids(Ncardids []int) []map[string]interface{} {
	if len(Ncardids) <= 0 {
		return []map[string]interface{}{}
	}
	return n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id: []interface{}{"IN", Ncardids},
	}).Select()
}
