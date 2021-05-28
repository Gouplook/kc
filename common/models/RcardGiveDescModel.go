//RcardGiveDescModel
//2020-11-04 15:19:34

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	)

//表结构体
type RcardGiveDescModel struct {
	Model *base.Model
	Field RcardGiveDescModelField
}

//表字段
type RcardGiveDescModelField struct{
	T_table	string	`default:"rcard_give_desc"`
	F_id	string	`default:"id"`
	F_rcard_id	string	`default:"rcard_id"`
	F_desc	string	`default:"desc"`
}

//初始化
func (r *RcardGiveDescModel) Init(ormer ...orm.Ormer) *RcardGiveDescModel{
	functions.ReflectModel(&r.Field)
	r.Model = base.NewModel(r.Field.T_table, ormer...)
	return r
}

//新增数据
func (r *RcardGiveDescModel) Insert(data map[string]interface{}) (int){
	result,_ := r.Model.Data(data).Insert()
	return result
}

//根据rcardId删除对应的赠品描述
func (r *RcardGiveDescModel) DelByRcardId(rcardId int) (err error) {
	_, err = r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Delete()

	return
}

//根据rcardId获取对应的赠品描述
func (r *RcardGiveDescModel) GetByRcardId(rcardId int) map[string]interface{} {
	res := r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id :rcardId,
	}).Find()
	return res
}

//根据rcardids获取赠品描述
func (r *RcardGiveDescModel)GetByRcardids(rcardids []int) []map[string]interface{} {
	if len(rcardids) <= 0 {
		return []map[string]interface{}{}
	}
	return r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", rcardids},
	}).Select()
}
