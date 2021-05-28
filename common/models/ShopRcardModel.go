//ShopRcardModel
//2020-08-05 15:23:57

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

//表结构体
type ShopRcardModel struct {
	Model *base.Model
	Field ShopRcardModelField
}

//表字段
type ShopRcardModelField struct {
	T_table      string `default:"shop_rcard"`
	F_id         string `default:"id"`
	F_shop_id    string `default:"shop_id"`
	F_rcard_id   string `default:"rcard_id"`
	F_status     string `default:"status"`
	F_is_del     string `default:"is_del"`
	F_del_time   string `default:"del_time"`
	F_under_time string `default:"under_time"`
	F_sales      string `default:"sales"`
	F_ctime      string `default:"ctime"`
}

//初始化
func (s *ShopRcardModel) Init(ormer ...orm.Ormer) *ShopRcardModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *ShopRcardModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

func (s *ShopRcardModel) GetByShopidAndStatus(shopId int, status int, start, limit int, fields ...string) []map[string]interface{} {

	if len(fields) > 0 {
		s.Model.Field(fields)
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id: shopId,
		s.Field.F_status:  status,
	}).Limit(start, limit).OrderBy(s.Field.F_id + " DESC ").Select()
}

func (s *ShopRcardModel) GetNumByShopidAndStatus(shopId, status int) int {
	if shopId <= 0 {
		return 0
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id: shopId,
		s.Field.F_status:  status,
	}).Count(s.Field.F_id)
}

//批量添加
func (s *ShopRcardModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}
	result, _ := s.Model.InsertAll(data)
	return result
}

//获取单条数据
func (s *ShopRcardModel) GetByShopidAndRcardid(shopId, rcardId int) map[string]interface{} {
	if shopId <= 0 || rcardId <= 0 {
		return map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:  shopId,
		s.Field.F_rcard_id: rcardId,
	}).Find()
}

//获取多条数据
func (s *ShopRcardModel) GetByShopidAndRcardids(shopId int, rcardIds []int) []map[string]interface{} {
	if shopId <= 0 || len(rcardIds) == 0 {
		return []map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:  shopId,
		s.Field.F_rcard_id: []interface{}{"IN", rcardIds},
	}).Select()
}

//获取门店的充值卡列表
func (s *ShopRcardModel) GetPageByShopId(shopid, start, limit, status int) []map[string]interface{} {

	if shopid <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}
	wheremap := map[string]interface{}{
		s.Field.F_shop_id: shopid,
		s.Field.F_is_del:  cards.IS_BUS_DEL_no,
	}
	if status > 0 {
		wheremap[s.Field.F_status] = status
	}
	return s.Model.Where(wheremap).OrderBy(s.Field.F_id+" DESC ").Limit(start, limit).Select()
}

//获取门店冲值卡数量
func (s *ShopRcardModel) GetNumByShopId(shopId int, status int) int {
	if shopId <= 0 {
		return 0
	}

	wheremap := map[string]interface{}{
		s.Field.F_shop_id: shopId,
		s.Field.F_is_del:  cards.TAG_IS_DEL_no,
	}
	if status > 0 {
		wheremap[s.Field.F_status] = status
	}

	return s.Model.Where(wheremap).Count(s.Field.F_id)
}

//批量修改
func (s *ShopRcardModel) UpdateByIds(ids []int, data map[string]interface{}) bool {
	if len(ids) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_id: []interface{}{"IN", ids},
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//根据主键ids 获取信息
func (s *ShopRcardModel) GetByIds(ids []int) []map[string]interface{} {
	if len(ids) == 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_id: []interface{}{"IN", ids},
	}).Select()
}

//获取多条信息
func (s *ShopRcardModel) GetByShopIDAndCardIDs(shopID int, rcardID []int) (data []map[string]interface{}) {
	if shopID <= 0 || len(rcardID) == 0 {
		return
	}
	data = s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:  shopID,
		s.Field.F_rcard_id: []interface{}{"in", rcardID},
	}).Select()
	return
}

//根据充值卡rcardIds获取数据
func (s *ShopRcardModel) GetByRcardids(rcardIds []int) []map[string]interface{} {
	if len(rcardIds) == 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_rcard_id: []interface{}{"IN", rcardIds},
	}).Select()

}

//根据充值卡ids修改数据
func (s *ShopRcardModel) UpdateByRcardids(rcardIds []int, data map[string]interface{}) bool {
	if len(rcardIds) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_rcard_id: []interface{}{"IN", rcardIds},
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//根据充值卡ids和shopId修改数据
func (s *ShopRcardModel) UpdateByRcardidsAndShopId(rcardIds []int, shopId int, data map[string]interface{}) bool {
	if len(rcardIds) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_rcard_id: []interface{}{"IN", rcardIds},
		s.Field.F_shop_id:  shopId,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//增加充值卡的销量
func (s *ShopRcardModel) IncrSalesByShopidAndRcardid(shopId, rcardId int, step int) bool {
	if rcardId <= 0 || step <= 0 || shopId <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:  shopId,
		s.Field.F_rcard_id: rcardId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}

//根据主键ids 获取信息
func (s *ShopRcardModel) GetByIDs(ids []int) (data []map[string]interface{}) {
	if len(ids) == 0 {
		return
	}

	data = s.Model.Where(map[string]interface{}{
		s.Field.F_id: []interface{}{"IN", ids},
	}).Select()

	return
}

func (s *ShopRcardModel) GetTotalNum(where map[string]interface{}) (result int) {
	if len(where) <= 0 {
		return
	}
	result = s.Model.Where(where).Count(s.Field.F_id)
	return
}

func (r *ShopRcardModel) SelectRcardsByWherePage(where []base.WhereItem, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		r.Model.Field(fields)
	}
	if limit > 0 {
		r.Model.Limit(start, limit)
	}
	return r.Model.Where(where).OrderBy(r.Field.F_rcard_id + " DESC ").Select()
}
