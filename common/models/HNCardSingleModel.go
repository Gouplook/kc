//HNCardSingleModel
//2020-04-16 11:27:02

package models

import (
	"errors"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type HNCardSingleModel struct {
	Model *base.Model
	Field HNCardSingleModelField
}

//表字段
type HNCardSingleModelField struct {
	T_table     string `default:"hncard_single"`
	F_id        string `default:"id"`
	F_hncard_id string `default:"hncard_id"`
	F_single_id string `default:"single_id"`
	F_ssp_id	string `default:"ssp_id"`
	F_num       string `default:"num"`
}

//初始化
func (n *HNCardSingleModel) Init(ormer ...orm.Ormer) *HNCardSingleModel {
	functions.ReflectModel(&n.Field)
	n.Model = base.NewModel(n.Field.T_table, ormer...)
	return n
}

//新增数据
func (n *HNCardSingleModel) Insert(data map[string]interface{}) (err error) {
	if result, insertErr := n.Model.Data(data).Insert(); insertErr != nil || result == 0 {
		err = errors.New("insert failed")
	}
	return
}

//批量添加
func (n *HNCardSingleModel) InsertAll(data []map[string]interface{}) (err error) {
	if len(data) == 0 {
		return
	}
	_, err = n.Model.InsertAll(data)
	return
}

//获取hncard包含的单项目
func (n *HNCardSingleModel) GetByHNCardID(hNCardID int) (data []map[string]interface{}) {
	if hNCardID == 0 {
		return []map[string]interface{}{}
	}
	data = n.Model.Where(map[string]interface{}{
		n.Field.F_hncard_id: hNCardID,
	}).Select()

	return
}

func (n *HNCardSingleModel) Count(where []base.WhereItem) int {
	if len(where) == 0 {
		return 0
	}
	return n.Model.Where(where).Count()
}

func (n *HNCardSingleModel) GetByHNCardIds(hNCardIds []int) (data []map[string]interface{}) {
	if len(hNCardIds) == 0 {
		return []map[string]interface{}{}
	}
	data = n.Model.Where(map[string]interface{}{
		n.Field.F_hncard_id: []interface{}{"IN", hNCardIds},
	}).Select()

	return
}

//删除包含的单项目
func (n *HNCardSingleModel) DelByIds(ids []int) (err error) {
	if len(ids) == 0 {
		return
	}
	_, err = n.Model.Where(map[string]interface{}{
		n.Field.F_id: []interface{}{"IN", ids},
	}).Delete()

	return
}

//修改项目数量
func (n *HNCardSingleModel) UpdateNumById(id int, num int) (err error) {
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


func (s *HNCardSingleModel)SelectByPage(where map[string]interface{},start,limit int)[]map[string]interface{}  {
	if len(where) <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(where).OrderBy(s.Field.F_id+" DESC ").Limit(start,limit).Select()
}