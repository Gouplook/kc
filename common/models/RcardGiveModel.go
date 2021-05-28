//RcardGiveModel
//2020-10-20 17:20:38

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type RcardGiveModel struct {
	Model *base.Model
	Field RcardGiveModelField
}

//表字段
type RcardGiveModelField struct {
	T_table     string `default:"rcard_give"`
	F_id        string `default:"id"`
	F_rcard_id  string `default:"rcard_id"`
	F_single_id string `default:"single_id"`
	F_num       string `default:"num"`
	F_period_of_validity string `default:"period_of_validity"`
}

//初始化
func (r *RcardGiveModel) Init(ormer ...orm.Ormer) *RcardGiveModel {
	functions.ReflectModel(&r.Field)
	r.Model = base.NewModel(r.Field.T_table, ormer...)
	return r
}

//新增数据
func (r *RcardGiveModel) Insert(data map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}
	result, _ := r.Model.Data(data).Insert()
	return result
}

//批量添加
func (r *RcardGiveModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}
	result, _ := r.Model.InsertAll(data)
	return result
}

//获取充值卡赠送的单项目
func (r *RcardGiveModel) GetByRcardid(rcardId int) []map[string]interface{} {
	if rcardId <= 0 {
		return []map[string]interface{}{}
	}

	return r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Select()
}

func (r *RcardGiveModel) GetByRcardids(rcardIds []int) []map[string]interface{} {
	if len(rcardIds) <= 0 {
		return []map[string]interface{}{}
	}

	return r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", rcardIds},
	}).Select()
}

//删除赠送的项目
func (r *RcardGiveModel) DelByIds(ids []int) bool {
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
func (r *RcardGiveModel) DelByRcardid(rcardId int) bool {
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

//更新数量
func (r *RcardGiveModel) UpdateNumById(id, num int) bool {
	if id <= 0 || num <= 0 {
		return false
	}
	_, err := r.Model.Where(map[string]interface{}{
		r.Field.F_id: id,
	}).Data(map[string]interface{}{
		r.Field.F_num: num,
	}).Update()
	if err != nil {
		return false
	}
	return true
}


//获取cardIds批量获取赠送的单项目
func (r *RcardGiveModel) GetByRcardIds(cardIds []int) (dataArray []map[string]interface{}) {
	if len(cardIds) == 0 {
		return []map[string]interface{}{}
	}

	dataArray = r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", cardIds},
	}).Select()
	return
}

func (r *RcardGiveModel) SelectByPage(where []base.WhereItem, start, limit int) []map[string]interface{} {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return r.Model.Where(where).OrderBy(r.Field.F_id+" DESC ").Limit(start, limit).Select()
}

func (r *RcardGiveModel) Count(where []base.WhereItem) int {
	if len(where) == 0 {
		return 0
	}
	return r.Model.Where(where).Count()
}
