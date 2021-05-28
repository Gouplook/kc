//ShopHcardModel 卡项服务-已添加到门店的限时卡
//2020-04-21 14:55:12

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"strconv"
)

//表结构体
type ShopHcardModel struct {
	Model *base.Model
	Field ShopHcardModelField
}

//表字段
type ShopHcardModelField struct {
	T_table      string `default:"shop_hcard"`
	F_id         string `default:"id"`
	F_shop_id    string `default:"shop_id"`  // 门店id
	F_hcard_id   string `default:"hcard_id"` // 限时卡id
	F_status     string `default:"status"`   // 限时卡在门店状态 1=下架 2=上架 3=被总店禁用
	F_under_time string `default:"under_time"`
	F_is_del     string `default:"is_del"`   // 删除状态：0-否，1-是
	F_del_time   string `default:"del_time"` // 删除时间
	F_sales      string `default:"sales"`    // 销量
	F_ctime      string `default:"ctime"`    // 限时卡上店时间

}

// Init 初始化
func (s *ShopHcardModel) Init(ormer ...orm.Ormer) *ShopHcardModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

// Insert 新增数据
func (s *ShopHcardModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

// InsertAll 批量添加
func (s *ShopHcardModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}
	result, _ := s.Model.InsertAll(data)
	return result
}

// GetByShopIDAndHcardIDs 根据子店id和限时卡ids获取多条信息
func (s *ShopHcardModel) GetByShopIDAndHcardIDs(shopID int, hcardIDs []int) []map[string]interface{} {
	//if shopID == 0 || len(hcardIDs) == 0 {
	//	return []map[string]interface{}{}
	//}
	var whereMap = map[string]interface{}{
		s.Field.F_hcard_id: []interface{}{"IN", hcardIDs},
	}
	if len(hcardIDs) == 0 || shopID < 0 {
		return []map[string]interface{}{}
	} else if shopID > 0 {
		whereMap[s.Field.F_shop_id] = shopID
	}
	return s.Model.Where(whereMap).Select()
}

// GetPageByShopID 根据门店id获取门店的限时卡列表
func (s *ShopHcardModel) GetPageByShopID(status string, shopID, start, limit int) []map[string]interface{} {
	if shopID <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}
	var where = map[string]interface{}{}
	where[s.Field.F_shop_id] = shopID
	where[s.Field.F_is_del] = cards.IS_BUS_DEL_no
	if len(status) > 0 {
		intStatus, err := strconv.Atoi(status)
		if err == nil {
			where[s.Field.F_status] = intStatus
		}
	}

	return s.Model.Where(where).OrderBy(s.Field.F_shop_id+" DESC ").
		Limit(start, limit).Select()
}

// GetNumByShopID 根据门店id获取门店限次卡总数量
func (s *ShopHcardModel) GetNumByShopID(shopID int, status string) int {
	var where = map[string]interface{}{}
	if shopID > 0 {
		where[s.Field.F_shop_id] = shopID
		where[s.Field.F_is_del] = cards.IS_BUS_DEL_no
	}

	if len(status) > 0 {
		intStatus, err := strconv.Atoi(status)
		if err == nil {
			where[s.Field.F_status] = intStatus
		}
	}

	return s.Model.Where(where).Count(s.Field.F_id)
}

// GetNumByShopID 根据门店id获取门店限次卡总数量
func (s *ShopHcardModel) GetTotalNum(where map[string]interface{}) int {

	if len(where) == 0 {
		return 0
	}
	return s.Model.Where(where).Count(s.Field.F_id)
}

// UpdateByIDs 根据ids批量修改
func (s *ShopHcardModel) UpdateByIDs(ids []int, data map[string]interface{}) bool {
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

// GetByIDs 根据id获取数据
func (s *ShopHcardModel) GetByIDs(ids []int) []map[string]interface{} {
	if len(ids) == 0 {
		return []map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_id: []interface{}{"IN", ids},
	}).Select()
}

// GetByHcardIDs 根据限时卡ids获取数据
func (s *ShopHcardModel) GetByHcardIDs(hcardIDs []int) []map[string]interface{} {
	if len(hcardIDs) == 0 {
		return []map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_hcard_id: []interface{}{"IN", hcardIDs},
	}).Select()
}

// UpdateByHcardIDs 根据限时卡ids批量修改数据
func (s *ShopHcardModel) UpdateByHcardIDs(hcardIDs []int, data map[string]interface{}) bool {
	if len(hcardIDs) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_hcard_id: []interface{}{"IN", hcardIDs},
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//增加限时卡的销量
func (s *ShopHcardModel) IncrSalesByShopidAndCardid(shopId, cardId int, step int) bool {
	if cardId <= 0 || step <= 0 || shopId <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:  shopId,
		s.Field.F_hcard_id: cardId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}
// 批量软删除门店限时卡
func (s *ShopHcardModel) UpdateDelByHcardIds(hcardIds []int ,data map[string]interface{}) bool {
	if len(hcardIds) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_hcard_id: []interface{}{"IN", hcardIds},
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

// 批量软删除门店限时卡
func (s *ShopHcardModel) UpdateDelByHcardIdsAndshopId(hcardIds []int, shopId int ,data map[string]interface{}) bool {
	if len(hcardIds) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_hcard_id: []interface{}{"IN", hcardIds},
		s.Field.F_shop_id:shopId,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}
// 根据hcardIds获取多条数据
func (s *ShopHcardModel)FindHcardIdsAndBusId(hcardIds []int, shopId int , fields ...string) []map[string]interface{}{
	if len(hcardIds) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		s.Model.Field(fields)
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_hcard_id: []interface{}{"IN", hcardIds},
		s.Field.F_is_del: cards.IS_BUS_DEL_no,
		s.Field.F_shop_id:   shopId,
	}).Select()
}


func (s *ShopHcardModel) SelectRcardsByWherePage(where []base.WhereItem, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		s.Model.Field(fields)
	}
	if limit > 0 {
		s.Model.Limit(start, limit)
	}
	return s.Model.Where(where).OrderBy(s.Field.F_hcard_id + " DESC ").Select()
}

func (s *ShopHcardModel)FindIsDelByHcarId(hcardId int ,isDel int )map[string]interface{}{
	res := s.Model.Where(map[string]interface{}{
		s.Field.F_hcard_id :hcardId,
		s.Field.F_is_del:isDel,
	}).Find()
	return res
}