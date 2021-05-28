//HNCardGiveModel
//2020-04-16 11:27:02

package models

import (
	"errors"
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type HNCardGiveModel struct {
	Model *base.Model
	Field HNCardGiveModelField
}

//表字段
type HNCardGiveModelField struct {
	T_table              string `default:"hncard_give"`
	F_id                 string `default:"id"`
	F_hncard_id          string `default:"hncard_id"`
	F_single_id          string `default:"single_id"`
	F_num                string `default:"num"`
	F_period_of_validity string `default:"period_of_validity"`
}

//初始化
func (n *HNCardGiveModel) Init(ormer ...orm.Ormer) *HNCardGiveModel {
	functions.ReflectModel(&n.Field)
	n.Model = base.NewModel(n.Field.T_table, ormer...)
	return n
}

//新增数据
func (n *HNCardGiveModel) Insert(data map[string]interface{}) (err error) {
	if result, insertErr := n.Model.Data(data).Insert(); insertErr != nil || result == 0 {
		err = errors.New("insert failed")
	}
	return
}

//批量添加
func (n *HNCardGiveModel) InsertAll(data []map[string]interface{}) (err error) {
	_, err = n.Model.InsertAll(data)
	return
}

//获取hncard赠送的单项目
func (n *HNCardGiveModel) GetByHNCardID(hNCardID int) (dataArray []map[string]interface{}) {
	if hNCardID == 0 {
		return []map[string]interface{}{}
	}

	dataArray = n.Model.Where(map[string]interface{}{
		n.Field.F_hncard_id: hNCardID,
	}).Select()
	return
}

func (n *HNCardGiveModel) GetByHNCardIds(hncardIds []int) (dataArray []map[string]interface{}) {
	if len(hncardIds) == 0 {
		return []map[string]interface{}{}
	}

	dataArray = n.Model.Where(map[string]interface{}{
		n.Field.F_hncard_id: []interface{}{"IN", hncardIds},
	}).Select()
	return
}

//删除赠送的项目
func (n *HNCardGiveModel) DelByIds(ids []int) (err error) {
	if len(ids) == 0 {
		return
	}
	_, err = n.Model.Where(map[string]interface{}{
		n.Field.F_id: []interface{}{"IN", ids},
	}).Delete()

	return
}

//更新数量
func (n *HNCardGiveModel) UpdateNumById(id int, num int) (err error) {
	if id == 0 || num <= 0 {
		return
	}
	_, err = n.Model.Where(map[string]interface{}{
		n.Field.F_id: id,
	}).Data(map[string]interface{}{
		n.Field.F_num: num,
	}).Update()

	return
}

func (n *HNCardGiveModel) SelectByPage(where []base.WhereItem, start, limit int) []map[string]interface{} {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return n.Model.Where(where).OrderBy(n.Field.F_id+" DESC ").Limit(start, limit).Select()
}

func (n *HNCardGiveModel) Count(where []base.WhereItem) int {
	if len(where) == 0 {
		return 0
	}
	return n.Model.Where(where).Count()
}
