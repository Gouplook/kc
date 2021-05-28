//ShopCardModel
//2020-04-24 09:26:28

package models

import (
	"fmt"
	"time"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"git.900sui.cn/kc/rpcCards/common/tools"
)

//表结构体
type ShopCardModel struct {
	Model *base.Model
	Field ShopCardModelField
	tools.CardStatus
}

//表字段
type ShopCardModelField struct {
	T_table      string `default:"shop_card"`
	F_id         string `default:"id"`
	F_shop_id    string `default:"shop_id"`
	F_card_id    string `default:"card_id"`
	F_status     string `default:"status"`
	F_is_del     string `default:"is_del"`
	F_under_time string `default:"under_time"`
	F_sales      string `default:"sales"`
	F_ctime      string `default:"ctime"`
	F_del_time   string `default:"del_time"`
}

//初始化
func (s *ShopCardModel) Init(ormer ...orm.Ormer) *ShopCardModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *ShopCardModel) Insert(data map[string]interface{}) (shopCardID int, err error) {
	shopCardID, err = s.Model.Data(data).Insert()
	return
}

//批量添加
func (s *ShopCardModel) InsertAll(data []map[string]interface{}) (result int, err error) {
	result, err = s.Model.InsertAll(data)
	return
}

//获取多条信息
func (s *ShopCardModel) GetByShopIDAndCardIDs(shopID int, cardID []int) (data []map[string]interface{}) {
	if shopID <= 0 || len(cardID) == 0 {
		return
	}
	where := []base.WhereItem{
		{s.Field.F_shop_id, shopID},
		{s.Field.F_card_id, []interface{}{"in", cardID}},
		{s.Field.F_is_del, s.NotDelStatus()},
	}
	data = s.Model.Where(where).Select()
	return
}


func (s *ShopCardModel) GetByShopByCardIDs(shopID int, cardID []int) (data []map[string]interface{}) {
	if shopID <= 0 || len(cardID) == 0 {
		return
	}
	where := []base.WhereItem{
		{s.Field.F_shop_id, shopID},
		{s.Field.F_card_id, []interface{}{"in", cardID}},
	}
	data = s.Model.Where(where).Select()
	return
}


//获取门店的综合卡列表
func (s *ShopCardModel) GetPageByShopID(shopID, start, limit, status int) (data []map[string]interface{}) {
	if shopID <= 0 || start < 0 || limit <= 0 {
		return
	}
	/*whereMap := map[string]interface{}{
		s.Field.F_shop_id: shopID,
	}
	if status > 0 {
		whereMap[s.Field.F_status] = status
	}*/
	whereMap := []base.WhereItem{
		{Field: s.Field.F_shop_id, Value: shopID},
	}
	if status > 0 {
		whereMap = append(whereMap, base.WhereItem{Field: s.Field.F_status, Value: status})
	}
	whereMap = append(whereMap, base.WhereItem{Field: s.Field.F_is_del, Value: s.NotDelStatus()})
	data = s.Model.Where(whereMap).OrderBy(fmt.Sprintf("%s %s", s.Field.F_id, "desc")).Limit(start, limit).Select()

	return
}

//获取门店的综合卡列表
func (s *ShopCardModel) GetCards(where map[string]interface{}, start, limit int) (data []map[string]interface{}) {
	if len(where) <= 0 || start < 0 || limit <= 0 {
		return
	}
	data = s.Model.Where(where).OrderBy(fmt.Sprintf("%s %s", s.Field.F_id, "desc")).Limit(start, limit).Select()
	return
}

//获取门店综合卡数量
func (s *ShopCardModel) GetNumByShopID(shopID int, status int) (result int) {
	if shopID <= 0 {
		return
	}

	/*whereMap := map[string]interface{}{
		s.Field.F_shop_id: shopID,
	}
	if status > 0 {
		whereMap[s.Field.F_status] = status
	}*/

	whereMap := []base.WhereItem{
		{Field: s.Field.F_shop_id, Value: shopID},
	}
	if status > 0 {
		whereMap = append(whereMap, base.WhereItem{Field: s.Field.F_status, Value: status})
	}
	whereMap = append(whereMap, base.WhereItem{Field: s.Field.F_is_del, Value: s.NotDelStatus()})
	result = s.Model.Where(whereMap).Count(s.Field.F_id)

	return
}

//获取门店综合卡数量
func (s *ShopCardModel) GetTotalNum(where map[string]interface{}) (result int) {
	if len(where) <= 0 {
		return
	}
	where[s.Field.F_is_del] = s.NotDelStatus()
	result = s.Model.Where(where).Count(s.Field.F_id)
	return
}

//批量修改
func (s *ShopCardModel) UpdateByIDs(ids []int, data map[string]interface{}) (err error) {
	if len(ids) == 0 || len(data) == 0 {
		return
	}
	_, err = s.Model.Where(map[string]interface{}{
		s.Field.F_id: []interface{}{"IN", ids},
	}).Data(data).Update()

	return
}

//根据主键ids 获取信息
func (s *ShopCardModel) GetByIDs(ids []int) (data []map[string]interface{}) {
	if len(ids) == 0 {
		return
	}

	data = s.Model.Where(map[string]interface{}{
		s.Field.F_id: []interface{}{"IN", ids},
	}).Select()

	return
}

//根据综合卡ids获取数据
func (s *ShopCardModel) GetByCardIDs(cardIDs []int) (data []map[string]interface{}) {
	if len(cardIDs) == 0 {
		return
	}
	where := []base.WhereItem{
		{s.Field.F_card_id, []interface{}{"IN", cardIDs}},
		{s.Field.F_is_del, s.NotDelStatus()},
	}
	data = s.Model.Where(where).Select()

	return
}

//根据综合卡ids修改数据
func (s *ShopCardModel) UpdateByCardIDs(cardIDs []int, data map[string]interface{}) (err error) {
	if len(cardIDs) == 0 || len(data) == 0 {
		return
	}
	_, err = s.Model.Where(map[string]interface{}{
		s.Field.F_card_id: []interface{}{"IN", cardIDs},
	}).Data(data).Update()

	return
}

// 批量软删除数据
func (s *ShopCardModel) UpdateDelByHcardIdsAndshopId(cardIds []int, shopId int ,data map[string]interface{}) bool {
	if len(cardIds) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_card_id: []interface{}{"IN", cardIds},
		s.Field.F_shop_id:shopId,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//增加套餐的销量
func (s *ShopCardModel) IncrSalesByShopidAndCardid(shopId, cardId int, step int) bool {
	if cardId <= 0 || step <= 0 || shopId <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id: shopId,
		s.Field.F_card_id: cardId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}

//分店-软删除
func (s *ShopCardModel) DelShopCardById(where []base.WhereItem) (b bool, err error) {

	if _, err = s.Model.Where(where).Data(map[string]interface{}{
		s.Field.F_is_del:   s.DelStatus(),
		s.Field.F_del_time: time.Now().Unix(),
	}).Update(); err != nil {
		return false, err
	}
	return true, nil
}

func (r *ShopCardModel) SelectRcardsByWherePage(where []base.WhereItem, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		r.Model.Field(fields)
	}
	if limit > 0 {
		r.Model.Limit(start, limit)
	}
	return r.Model.Where(where).OrderBy(r.Field.F_card_id + " DESC ").Select()
}