//NCardGiveModel
//2020-04-16 11:27:02

package models

import (
	"errors"
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type NCardGiveModel struct {
	Model *base.Model
	Field NCardGiveModelField
}

//表字段
type NCardGiveModelField struct {
	T_table              string `default:"ncard_give"`
	F_id                 string `default:"id"`
	F_ncard_id           string `default:"ncard_id"`
	F_single_id          string `default:"single_id"`
	F_num                string `default:"num"`
	F_period_of_validity string `default:"period_of_validity"`
}

//初始化
func (n *NCardGiveModel) Init(ormer ...orm.Ormer) *NCardGiveModel {
	functions.ReflectModel(&n.Field)
	n.Model = base.NewModel(n.Field.T_table, ormer...)
	return n
}

//新增数据
func (n *NCardGiveModel) Insert(data map[string]interface{}) (err error) {
	if result, insertErr := n.Model.Data(data).Insert(); insertErr != nil || result == 0 {
		err = errors.New("insert failed")
	}
	return
}

//批量添加
func (n *NCardGiveModel) InsertAll(data []map[string]interface{}) (err error) {
	_, err = n.Model.InsertAll(data)
	return
}

//获取ncard赠送的单项目
func (n *NCardGiveModel) GetByNCardID(nCardID int) (dataArray []map[string]interface{}) {
	if nCardID == 0 {
		return []map[string]interface{}{}
	}

	dataArray = n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id: nCardID,
	}).Select()
	return
}

func (n *NCardGiveModel) GetByNCardIds(nCardIds []int) (dataArray []map[string]interface{}) {
	if len(nCardIds) == 0 {
		return []map[string]interface{}{}
	}

	dataArray = n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id: []interface{}{"IN", nCardIds},
	}).Select()
	return
}

//删除赠送的项目
func (n *NCardGiveModel) DelByIds(ids []int) (err error) {
	if len(ids) == 0 {
		return
	}
	_, err = n.Model.Where(map[string]interface{}{
		n.Field.F_id: []interface{}{"IN", ids},
	}).Delete()

	return
}

//更新数量
func (n *NCardGiveModel) UpdateNumById(id int, num int) (err error) {
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

func (n *NCardGiveModel) SelectByPage(where []base.WhereItem, start, limit int) []map[string]interface{} {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return n.Model.Where(where).OrderBy(n.Field.F_id+" DESC ").Limit(start, limit).Select()
}

func (n *NCardGiveModel) Count(where []base.WhereItem) int {
	if len(where) == 0 {
		return 0
	}
	return n.Model.Where(where).Count()
}
