//套餐赠送的单服务表模型
//2020-04-14 19:20:32

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type SmGiveModel struct {
	Model *base.Model
	Field SmGiveModelField
}

//表字段
type SmGiveModelField struct {
	T_table              string `default:"sm_give"`
	F_id                 string `default:"id"`
	F_sm_id              string `default:"sm_id"`
	F_single_id          string `default:"single_id"`
	F_num                string `default:"num"`
	F_period_of_validity string `default:"period_of_validity"`
}

//初始化
func (s *SmGiveModel) Init(ormer ...orm.Ormer) *SmGiveModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *SmGiveModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//批量添加
func (s *SmGiveModel) InsertAll(data []map[string]interface{}) int {
	result, _ := s.Model.InsertAll(data)
	return result
}

//获取套餐赠送的单项目
func (s *SmGiveModel) GetBySmid(smid int) []map[string]interface{} {
	if smid <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: smid,
	}).Select()
}

func (s *SmGiveModel) GetBySmids(smids []int) []map[string]interface{} {
	if len(smids) <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: []interface{}{"IN", smids},
	}).Select()
}

//删除赠送的项目
func (s *SmGiveModel) DelByIds(ids []int) bool {
	if len(ids) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_id: []interface{}{"IN", ids},
	}).Delete()
	if err != nil {
		return false
	}
	return true
}

//更新数量
func (s *SmGiveModel) UpdateNumById(id, num int) bool {
	if id <= 0 || num <= 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_id: id,
	}).Data(map[string]interface{}{
		s.Field.F_num: num,
	}).Update()
	if err != nil {
		return false
	}
	return true
}

//更新有效期天数
func (s *SmGiveModel) UpdateValidityById(id, periodOfValidity int) bool {
	if id <= 0 || periodOfValidity < 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_id: id,
	}).Data(map[string]interface{}{
		s.Field.F_period_of_validity: periodOfValidity,
	}).Update()
	if err != nil {
		return false
	}
	return true
}


func (s *SmGiveModel) SelectByPage(where []base.WhereItem, start, limit int) []map[string]interface{} {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(where).OrderBy(s.Field.F_id+" DESC ").Limit(start, limit).Select()
}

func (s *SmGiveModel) Count(where []base.WhereItem) int {
	if len(where) == 0 {
		return 0
	}
	return s.Model.Where(where).Count()
}
