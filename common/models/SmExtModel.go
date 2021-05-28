//SmExtModel
//2020-11-18 14:25:42

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type SmExtModel struct {
	Model *base.Model
	Field SmExtModelField
}

//表字段
type SmExtModelField struct {
	T_table             string `default:"sm_ext"`
	F_id                string `default:"id"`
	F_sm_id             string `default:"sm_id"`
	F_notes             string `default:"notes"`
	F_service_subscribe string `default:"service_subscribe"`
}

//初始化
func (s *SmExtModel) Init(ormer ...orm.Ormer) *SmExtModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *SmExtModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//查询数据
func (s *SmExtModel) Find(kv map[string]interface{}) (data map[string]interface{}) {
	data = s.Model.Where(kv).Find()
	return
}

//更新数据
func (s *SmExtModel) Update(kv map[string]interface{}, data map[string]interface{}) (result int, err error) {
	result, err = s.Model.Where(kv).Data(data).Update()
	return
}

// UpdateBySmID 根据套餐id更新数据
func (s *SmExtModel) UpdateBySmID(smId int, data map[string]interface{}) bool {
	if _, err := s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: smId,
	}).Data(data).Update(); err != nil {
		return false
	}
	return true
}

// GetBySmID 根据套餐id获取数据
func (s *SmExtModel) GetBySmID(smId int) map[string]interface{} {
	if smId <= 0 {
		return map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: smId,
	}).Find()
}
