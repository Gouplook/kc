//SingleExtModel
//2020-03-27 15:37:58

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
)

//表结构体
type SingleExtModel struct {
	Model *base.Model
	Field SingleExtModelField
}

//表字段
type SingleExtModelField struct {
	T_table            string `default:"single_ext"`
	F_id               string `default:"id"`
	F_single_id        string `default:"single_id"`
	F_sex              string `default:"sex"`
	F_age_bracket      string `default:"age_bracket"`
	F_tailor_indus     string `default:"tailor_indus"`
	F_tailor_sub_indus string `default:"tailor_sub_indus"`
	F_service_reminder string `default:"service_reminder"`
	F_single_context   string `default:"single_context"`
}

//初始化
func (s *SingleExtModel) Init() *SingleExtModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table)
	return s
}

//新增数据
func (s *SingleExtModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//获取单条单项目详情
func (s *SingleExtModel) GetBySingleid(singleId int) map[string]interface{} {
	return s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
	}).Limit(0, 1).Find()

}

//修改扩展数据
func (s *SingleExtModel) UpdateBySingleid(singleId int, data map[string]interface{}) bool {
	if singleId <= 0 || len(data) == 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
	}).Data(data).Update()
	if err != nil {
		return false
	}

	return true
}
