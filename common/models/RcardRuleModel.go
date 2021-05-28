//RcardRuleModel
//2020-10-20 17:21:50

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type RcardRuleModel struct {
	Model *base.Model
	Field RcardRuleModelField
}

//表字段
type RcardRuleModelField struct {
	T_table      string `default:"rcard_rule"`
	F_id         string `default:"id"`
	F_rcard_id   string `default:"rcard_id"`
	F_price      string `default:"price"`
	F_give_price string `default:"give_price"`
	F_is_del     string `default:"is_del"`
}

//初始化
func (r *RcardRuleModel) Init(ormer ...orm.Ormer) *RcardRuleModel {
	functions.ReflectModel(&r.Field)
	r.Model = base.NewModel(r.Field.T_table, ormer...)
	return r
}

//新增数据
func (r *RcardRuleModel) Insert(data map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}

	result, _ := r.Model.Data(data).Insert()
	return result
}

//批量添加
func (r *RcardRuleModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}

	rs, _ := r.Model.InsertAll(data)
	return rs
}
//根据充值卡id删除
func (r *RcardRuleModel) DelByRcardid(rcardId int) bool {
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

//批量删除充值卡规则
func (r *RcardRuleModel) DelByIds(ids []int) bool {
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

//修改数据
func (r *RcardRuleModel) UpdateById(id int, data map[string]interface{}) bool {
	if id <= 0 || len(data) == 0 {
		return false
	}

	_, err := r.Model.Where([]base.WhereItem{
		{
			Field: r.Field.F_id,
			Value: id,
		},
	}).Data(data).Update()
	if err != nil {
		return false
	}

	return true
}

//获取卡项规则
func (r *RcardRuleModel) GetByRcardid(rcardId int) []map[string]interface{} {
	if rcardId <= 0 {
		return []map[string]interface{}{}
	}

	return r.Model.Where([]base.WhereItem{
		{
			Field: r.Field.F_rcard_id,
			Value: rcardId,
		},
	}).OrderBy(r.Field.F_id + " DESC ").Select()
}


func (r *RcardRuleModel) Select(where map[string]interface{}) []map[string]interface{} {
	if len(where) > 0 {
		r.Model.Where(where)
	}
	return r.Model.OrderBy(r.Field.F_id + " DESC ").Select()
}

// 获取充值卡 单条数据
func (r *RcardRuleModel) GetBycrardId(id int, files ...[]string) map[string]interface{} {
	if id <= 0 {
		return map[string]interface{}{}
	}
	if len(files) > 0 {
		r.Model.Field(files[0])
	}
	return r.Model.Where(map[string]interface{}{
		r.Field.F_id: id,
	}).Find()
}

// 根据充值卡规则Id 更新数据
func (r *RcardRuleModel) UpdateRelusId(id int, data map[string]interface{}) bool {
	if id <= 0 || len(data) == 0 {
		return false
	}
	_, err := r.Model.Where(map[string]interface{}{
		r.Field.F_id: id,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}
// 根据充值卡(rcardId) 更新数据
func (r *RcardRuleModel) UpdateRcardId(rcardId int, data map[string]interface{}) bool {
	if rcardId <= 0 || len(data) == 0 {
		return false
	}
	_, err := r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

// 获取充值卡规则列表
func (r *RcardRuleModel) GetPageByRcardId(rcardId int ,start,limit int, status int)[]map[string]interface{}{
	if rcardId <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}
	return  r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id:rcardId,
		r.Field.F_is_del:status,
	}).OrderBy(r.Field.F_id+ " DESC ").Limit(start,limit).Select()
}

// 根据充值卡规则Id 获取充值卡规则rcard_id下总数量
func ( r *RcardRuleModel) GetNumByRcardId(rcardId int,status int) int {
	if rcardId <= 0 {
		return  0
	}

	return r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id:rcardId,
		r.Field.F_is_del:status,
	}).Count(r.Field.F_id)
}

