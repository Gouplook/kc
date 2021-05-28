//ShopItemRelationModel
//2021-04-25 10:19:25

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

//表结构体
type ShopItemRelationModel struct {
	Model *base.Model
	Field ShopItemRelationModelField
}

//表字段
type ShopItemRelationModelField struct {
	T_table     string `default:"shop_item_relation"`
	F_id        string `default:"id"`
	F_item_id   string `default:"item_id"`
	F_item_type string `default:"item_type"`
	F_status    string `default:"status"`
	F_is_del    string `default:"is_del"`
	F_shop_id   string `default:"shop_id"`
}

//初始化
func (s *ShopItemRelationModel) Init(ormer ...orm.Ormer) *ShopItemRelationModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *ShopItemRelationModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

func (s *ShopItemRelationModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}

	result, _ := s.Model.InsertAll(data)
	return result
}

func (s *ShopItemRelationModel) DelByItemIds(itemIds []int, itemType int, shopId ...int) bool {
	if len(itemIds) == 0 || itemType == 0 {
		return false
	}
	where := []base.WhereItem{
		{s.Field.F_item_id, []interface{}{"IN", itemIds}},
		{s.Field.F_item_type, itemType},
	}
	if len(shopId) > 0 {
		where = append(where, base.WhereItem{s.Field.F_shop_id, shopId[0]})
	}
	_, err := s.Model.Where(where).Data(map[string]interface{}{s.Field.F_is_del: cards.ITEM_IS_DEL_YES}).Update()
	if err != nil {
		return false
	}
	return true
}

func (s *ShopItemRelationModel) UpdateStatusByItemIds(itemIds []int, itemType int, status int, shopId ...int) bool {
	if len(itemIds) == 0 || itemType == 0 || status == 0 {
		return false
	}
	where := []base.WhereItem{
		{s.Field.F_item_id, []interface{}{"IN", itemIds}},
		{s.Field.F_item_type, itemType},
	}
	if len(shopId) > 0 {
		where = append(where, base.WhereItem{s.Field.F_shop_id, shopId[0]})
	}
	_, err := s.Model.Where(where).Data(map[string]interface{}{s.Field.F_status: status}).Update()
	if err != nil {
		return false
	}
	return true
}

func (s *ShopItemRelationModel) UpdateByItemIdsAndShopId(itemIds []int, itemType, shopId int, data map[string]interface{}) bool {
	if len(itemIds) == 0 || itemType == 0 || shopId == 0 {
		return false
	}
	where := []base.WhereItem{
		{s.Field.F_item_id, []interface{}{"IN", itemIds}},
		{s.Field.F_item_type, itemType},
		{s.Field.F_shop_id, shopId},
	}
	_, err := s.Model.Where(where).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

func (s *ShopItemRelationModel) FindByWhere(where []base.WhereItem, fields ...[]string) map[string]interface{} {
	if len(where) == 0 {
		return map[string]interface{}{}
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	return s.Model.Where(where).Find()
}

func (s *ShopItemRelationModel) SelectByWhere(where []base.WhereItem, fields ...[]string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	return s.Model.Where(where).Select()
}

func (s *ShopItemRelationModel) SelectByWhereByPage(where []base.WhereItem, start, limit int, fields ...[]string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	return s.Model.Where(where).Limit(start, limit).OrderBy(s.Field.F_id + " DESC ").Select()
}

func (s *ShopItemRelationModel) GetTotalNumByWhere(where []base.WhereItem) int {
	if len(where) == 0 {
		return 0
	}
	return s.Model.Where(where).Count(s.Field.F_id)
}
