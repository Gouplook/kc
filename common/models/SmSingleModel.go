//SmSingleModel
//2020-04-14 19:20:32

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type SmSingleModel struct {
	Model *base.Model
	Field SmSingleModelField
}

//表字段
type SmSingleModelField struct {
	T_table     string `default:"sm_single"`
	F_id        string `default:"id"`
	F_sm_id     string `default:"sm_id"`
	F_single_id string `default:"single_id"`
	F_ssp_id	 string `default:"ssp_id"`
	F_num       string `default:"num"`
}

//初始化
func (s *SmSingleModel) Init(ormer ...orm.Ormer) *SmSingleModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *SmSingleModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//批量添加
func (s *SmSingleModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}

	r, _ := s.Model.InsertAll(data)
	return r
}

//获取套餐包含的单项目
func (s *SmSingleModel) Count(where []base.WhereItem) int {
	if len(where) == 0 {
		return 0
	}
	return s.Model.Where(where).Count()
}

//获取套餐包含的单项目
func (s *SmSingleModel) GetBySmid(smid int) []map[string]interface{} {
	if smid <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: smid,
	}).Select()
}

func (s *SmSingleModel)SelectByPage(where map[string]interface{},start,limit int)[]map[string]interface{}  {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(where).OrderBy(s.Field.F_id+" DESC ").Limit(start,limit).Select()
}

//获取多个套餐包含的单项目
func (s *SmSingleModel) GetBySmids(smids []int) []map[string]interface{} {
	if len(smids) <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: []interface{}{"IN", smids},
	}).Select()
}

//删除包含的单项目
func (s *SmSingleModel) DelByIds(ids []int) bool {
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

//修改项目数量
func (s *SmSingleModel) UpdateNumById(id int, num int) bool {
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
