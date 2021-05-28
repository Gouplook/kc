//SmGiveDescModel
//2020-11-04 15:19:34

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	)

//表结构体
type SmGiveDescModel struct {
	Model *base.Model
	Field SmGiveDescModelField
}

//表字段
type SmGiveDescModelField struct{
	T_table	string	`default:"sm_give_desc"`
	F_id	string	`default:"id"`
	F_sm_id	string	`default:"sm_id"`
	F_desc	string	`default:"desc"`
}

//初始化
func (s *SmGiveDescModel) Init(ormer ...orm.Ormer) *SmGiveDescModel{
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *SmGiveDescModel) Insert(data map[string]interface{}) (int){
	result,_ := s.Model.Data(data).Insert()
	return result
}

//根据smcardid删除对应的赠品描述
func (s *SmGiveDescModel) DelBySmId(smId int) (err error) {
	_, err = s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: smId,
	}).Delete()

	return
}

//根据smcardid获取对应的赠品描述
func (s *SmGiveDescModel) GetBySmId(smId int) map[string]interface{} {
	res := s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id :smId,
	}).Find()
	return res
}

//根据smids获取赠品描述
func (s *SmGiveDescModel)GetBySmids(smids []int) []map[string]interface{} {
	if len(smids) <= 0 {
		return []map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: []interface{}{"IN", smids},
	}).Select()
}
