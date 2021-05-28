//RcardSingleModel
//2020-10-20 18:08:16

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type RcardSingleModel struct {
	Model *base.Model
	Field RcardSingleModelField
}

//表字段
type RcardSingleModelField struct {
	T_table     string `default:"rcard_single"`
	F_id        string `default:"id"`
	F_rcard_id  string `default:"rcard_id"`
	F_single_id string `default:"single_id"`
	F_discount  string `default:"discount"`
}

//初始化
func (r *RcardSingleModel) Init(ormer ...orm.Ormer) *RcardSingleModel {
	functions.ReflectModel(&r.Field)
	r.Model = base.NewModel(r.Field.T_table, ormer...)
	return r
}

//新增数据
func (r *RcardSingleModel) Insert(data map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}
	result, _ := r.Model.Data(data).Insert()
	return result
}

//批量添加
func (r *RcardSingleModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}

	rs, _ := r.Model.InsertAll(data)
	return rs
}

//获取充值卡包含的单项目
func (r *RcardSingleModel) GetByRcardid(rcardId int) []map[string]interface{} {
	if rcardId <= 0 {
		return []map[string]interface{}{}
	}

	return r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Select()
}

//获取多个充值卡包含的单项目
func (r *RcardSingleModel) GetByRcardids(rcardIds []int) []map[string]interface{} {
	if len(rcardIds) <= 0 {
		return []map[string]interface{}{}
	}

	return r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", rcardIds},
	}).Select()
}

//删除包含的单项目
func (r *RcardSingleModel) DelByIds(ids []int) bool {
	if len(ids) == 0 {
		return false
	}
	_, err := r.Model.Where(map[string]interface{}{
		r.Field.F_id: []interface{}{"IN", ids},
	}).Delete()
	if err != nil {
		return false
	}
	return true
}

//根据充值卡id删除
func (r *RcardSingleModel) DelByRcardid(rcardId int) bool {
	if rcardId <= 0 {
		return false
	}
	_, err := r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Delete()
	if err != nil {
		return false
	}
	return true
}

//根据cardIds批量获取单项目
func (r *RcardSingleModel) GetByRcardIds(cardIds []int) (data []map[string]interface{}) {
	if len(cardIds) == 0 {
		return []map[string]interface{}{}
	}
	data = r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", cardIds},
	}).Select()

	return
}

func (s *RcardSingleModel) SelectByPage(where map[string]interface{}, start, limit int) []map[string]interface{} {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(where).OrderBy(s.Field.F_id+" DESC ").Limit(start, limit).Select()
}

func (s *RcardSingleModel) GetTotalNum(where map[string]interface{}) int {
	if len(where) == 0 {
		return 0
	}
	return s.Model.Where(where).Count(s.Field.F_id)
}

func (s *RcardSingleModel) Find(where map[string]interface{}) map[string]interface{} {
	if len(where) <= 0 {
		return map[string]interface{}{}
	}

	return s.Model.Where(where).OrderBy(s.Field.F_id + " DESC ").Find()
}
