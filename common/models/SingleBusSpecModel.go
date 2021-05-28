//SingleBusSpecModel
//2020-03-27 15:37:30

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
)

//表结构体
type SingleBusSpecModel struct {
	Model *base.Model
	Field SingleBusSpecModelField
}

//表字段
type SingleBusSpecModelField struct {
	T_table     string `default:"single_bus_spec"`
	F_spec_id   string `default:"spec_id"`
	F_name      string `default:"name"`
	F_bus_id    string `default:"bus_id"`
	F_p_spec_id string `default:"p_spec_id"`
}

//初始化
func (s *SingleBusSpecModel) Init() *SingleBusSpecModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table)
	return s
}

//新增数据
func (s *SingleBusSpecModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//根据specId获取单条信息
func (s *SingleBusSpecModel) GetBySpecid(specId int) map[string]interface{} {
	if specId <= 0 {
		return map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_spec_id: specId,
	}).Limit(0, 1).Find()
}

//根据busid和name获取数据
func (s *SingleBusSpecModel) GetByBusidAndName(busId int, parentSpecId int, name string) map[string]interface{} {
	if busId <= 0 || parentSpecId <= 0 || name == "" {
		return map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_bus_id:    busId,
		s.Field.F_p_spec_id: parentSpecId,
		s.Field.F_name:      name,
	}).Limit(0, 1).Find()
}

//根据父id获取子规格
func (s *SingleBusSpecModel) GetByParentSpecId(parentSpecId int, busId int) []map[string]interface{} {
	if parentSpecId < 0 || busId <= 0 {
		return []map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_bus_id:    busId,
		s.Field.F_p_spec_id: parentSpecId,
	}).OrderBy(s.Field.F_spec_id + " ASC ").Select()
}

//根据多个父id获取子规格信息
func (s *SingleBusSpecModel) GetByParentSpecIds(parentSpecIds []int, busId int) []map[string]interface{} {
	return s.Model.Where(map[string]interface{}{
		s.Field.F_bus_id: busId,
		s.Field.F_p_spec_id: []interface{}{
			"IN",
			parentSpecIds,
		},
	}).OrderBy(s.Field.F_spec_id + " ASC ").Select()
}

//根据specIds获取多条信息
func (s *SingleBusSpecModel) GetBySpecids(specIds []int) []map[string]interface{} {
	if len(specIds) == 0 {
		return []map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_spec_id: []interface{}{
			"IN",
			specIds,
		},
	}).Select()
}
