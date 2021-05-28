//RcardGoodsModel
//2020-10-20 17:21:03

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type RcardGoodsModel struct {
	Model *base.Model
	Field RcardGoodsModelField
}

//表字段
type RcardGoodsModelField struct {
	T_table      string `default:"rcard_goods"`
	F_id         string `default:"id"`
	F_rcard_id   string `default:"rcard_id"`
	F_product_id string `default:"product_id"`
}

//初始化
func (r *RcardGoodsModel) Init(ormer ...orm.Ormer) *RcardGoodsModel {
	functions.ReflectModel(&r.Field)
	r.Model = base.NewModel(r.Field.T_table, ormer...)
	return r
}

//新增数据
func (r *RcardGoodsModel) Insert(data map[string]interface{}) int {
	result, _ := r.Model.Data(data).Insert()
	return result
}

//批量增加数据
func (r *RcardGoodsModel) InsertAll(data []map[string]interface{}) (result int, err error) {
	result, err = r.Model.InsertAll(data)
	return
}

//获取充值卡包含的商品
func (r *RcardGoodsModel) GetByRcardid(rcardId int) (data []map[string]interface{}) {
	if rcardId == 0 {
		return []map[string]interface{}{}
	}

	data = r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: rcardId,
	}).Select()

	return
}

func (r *RcardGoodsModel) DelByIds(ids []int) (err error) {
	_, err = r.Model.Where(map[string]interface{}{
		r.Field.F_id: []interface{}{"IN", ids},
	}).Delete()
	return
}

//获取多个充值卡包含的商品
func (r *RcardGoodsModel) GetByRcardids(rcardIds []int) (data []map[string]interface{}) {
	if len(rcardIds) == 0 {
		return []map[string]interface{}{}
	}
	data = r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", rcardIds},
	}).Select()
	return
}

//根据充值卡id删除
func (r *RcardGoodsModel) DelByRcardid(rcardId int) bool {
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

//获取cards包含的商品
func (r *RcardGoodsModel) GetByRcardIds(cardIds []int) (data []map[string]interface{}) {
	if len(cardIds) == 0 {
		return []map[string]interface{}{}
	}
	data = r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: []interface{}{"IN", cardIds},
	}).Select()

	return
}

//获取rcard包含的商品,带分页
func (r *RcardGoodsModel) GetByRcardIdPage(cardID,start,pageSize int,) (data []map[string]interface{}) {
	if cardID == 0 {
		return []map[string]interface{}{}
	}
	data = r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id: cardID,
	}).Limit(start, pageSize).Select()
	return
}

//根据rcardId获取总数
func (r *RcardGoodsModel) CountByCardId(rcardId int) int {
	count := r.Model.Where(map[string]interface{}{
		r.Field.F_rcard_id:rcardId,
	}).Count("*")
	return count
}

func (s *RcardGoodsModel)Find(where map[string]interface{})map[string]interface{}  {
	if len(where) <= 0 {
		return map[string]interface{}{}
	}

	return s.Model.Where(where).OrderBy(s.Field.F_id+" DESC ").Find()
}