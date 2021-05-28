//SingleSpecModel
//2020-03-27 15:38:59

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
)

//表结构体
type SingleSpecModel struct {
	Model *base.Model
	Field SingleSpecModelField
}

//表字段
type SingleSpecModelField struct {
	T_table     string `default:"single_spec"`
	F_id        string `default:"id"`
	F_single_id string `default:"single_id"`
	F_spec_info string `default:"spec_info"`
}

//初始化
func (s *SingleSpecModel) Init() *SingleSpecModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table)
	return s
}

//新增数据
func (s *SingleSpecModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//批量添加
func (s *SingleSpecModel) InsertAll(data []map[string]interface{}) int {
	result, _ := s.Model.InsertAll(data)
	return result
}

//获取规格数据
func (s *SingleSpecModel) GetBySingleid(singleId int, field ...string) map[string]interface{} {
	if singleId <= 0 {
		return make(map[string]interface{})
	}
	if len(field) > 0 {
		s.Model.Field(field)
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
	}).Find()
}

//获取规格数据
func (s *SingleSpecModel) GetBySingleids(singleIds []int) []map[string]interface{} {
	return s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: []interface{}{"IN", singleIds},
	}).Select()
}

//修改数据
func (s *SingleSpecModel) UpdateById(id int, data map[string]interface{}) int {
	r, _ := s.Model.Where(map[string]interface{}{
		s.Field.F_id: id,
	}).Data(data).Update()
	return r
}

//删除规格
func (s *SingleSpecModel) DeleteBySingleid(singleId int) bool {
	r, _ := s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
	}).Delete()
	if r <= 0 {
		return false
	}
	return true

}
