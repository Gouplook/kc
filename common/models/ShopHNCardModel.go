//ShopHNCardModel
//2020-04-20 13:37:51

package models

import (
	"fmt"
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"github.com/wendal/errors"
)

//表结构体
type ShopHNCardModel struct {
	Model *base.Model
	Field ShopHNCardModelField
}

//表字段
type ShopHNCardModelField struct {
	T_table      string `default:"shop_hncard"`
	F_id         string `default:"id"`
	F_shop_id    string `default:"shop_id"`
	F_hncard_id  string `default:"hncard_id"`
	F_status     string `default:"status"`
	F_is_del     string `default:"is_del"`
	F_del_time   string `default:"del_time"`
	F_under_time string `default:"under_time"`
	F_sales      string `default:"sales"`
	F_ctime      string `default:"ctime"`
}

//初始化
func (s *ShopHNCardModel) Init(ormer ...orm.Ormer) *ShopHNCardModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *ShopHNCardModel) Insert(data map[string]interface{}) (shopHNCardID int, err error) {
	if result, insertErr := s.Model.Data(data).Insert(); insertErr != nil || result == 0 {
		err = errors.New("insert failed")
	} else {
		shopHNCardID = result
	}
	return
}

//批量添加
func (s *ShopHNCardModel) InsertAll(data []map[string]interface{}) (result int, err error) {
	result, err = s.Model.InsertAll(data)
	return
}

//获取多条信息
func (s *ShopHNCardModel) GetByShopIDAndHNCardIDs(shopID int, hNCardID []int) (data []map[string]interface{}) {
	if shopID <= 0 || len(hNCardID) == 0 {
		return
	}
	data = s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:   shopID,
		s.Field.F_hncard_id: []interface{}{"in", hNCardID},
	}).Select()
	return
}

//获取门店的限时限次卡列表
func (s *ShopHNCardModel) GetPageByShopID(shopID, start, limit, status int) (data []map[string]interface{}) {
	if shopID <= 0 || start < 0 || limit <= 0 {
		return
	}
	whereMap := map[string]interface{}{
		s.Field.F_shop_id: shopID,
		s.Field.F_is_del: cards.IS_BUS_DEL_no,
	}
	if status > 0 {
		whereMap[s.Field.F_status] = status
	}
	data = s.Model.Where(whereMap).OrderBy(fmt.Sprintf("%s %s", s.Field.F_id, "desc")).Limit(start, limit).Select()

	return
}

//获取门店的限时限次卡列表
func (s *ShopHNCardModel) GetList(where map[string]interface{}, start, limit int) (data []map[string]interface{}) {
	if len(where) <= 0 || start < 0 || limit <= 0 {
		return
	}
	data = s.Model.Where(where).OrderBy(s.Field.F_id+" desc ").Limit(start, limit).Select()
	return
}

//获取门店限时限次卡数量
func (s *ShopHNCardModel) GetNumByShopID(shopID int, status int) (result int) {
	if shopID <= 0 {
		return
	}
	whereMap := map[string]interface{}{
		s.Field.F_shop_id: shopID,
		s.Field.F_is_del:cards.IS_BUS_DEL_no,
	}

	if status > 0 {
		whereMap[s.Field.F_status] = status
	}
	result = s.Model.Where(whereMap).Count(s.Field.F_id)

	return
}

//获取门店限时限次卡数量
func (s *ShopHNCardModel) GetTotalNum(where map[string]interface{}) (result int) {
	if len(where) <= 0 {
		return
	}
	result = s.Model.Where(where).Count(s.Field.F_id)
	return
}

//批量修改
func (s *ShopHNCardModel) UpdateByIDs(ids []int, data map[string]interface{}) (err error) {
	if len(ids) == 0 || len(data) == 0 {
		return
	}
	_, err = s.Model.Where(map[string]interface{}{
		s.Field.F_id: []interface{}{"IN", ids},
	}).Data(data).Update()

	return
}

//根据主键ids 获取信息
func (s *ShopHNCardModel) GetByIDs(ids []int) (data []map[string]interface{}) {
	if len(ids) == 0 {
		return
	}

	data = s.Model.Where(map[string]interface{}{
		s.Field.F_id: []interface{}{"IN", ids},
	}).Select()

	return
}

//根据限时限次卡ids获取数据
func (s *ShopHNCardModel) GetByHNCardIDs(hNCardIDs []int) (data []map[string]interface{}) {
	if len(hNCardIDs) == 0 {
		return
	}

	data = s.Model.Where(map[string]interface{}{
		s.Field.F_hncard_id: []interface{}{"IN", hNCardIDs},
	}).Select()

	return
}

//根据限时限次卡ids修改数据
func (s *ShopHNCardModel) UpdateByHNCardIDs(hNCardIDs []int, data map[string]interface{}) (err error) {
	if len(hNCardIDs) == 0 || len(data) == 0 {
		return
	}
	_, err = s.Model.Where(map[string]interface{}{
		s.Field.F_hncard_id: []interface{}{"IN", hNCardIDs},
	}).Data(data).Update()

	return
}
//根据限时限次卡ids 和shopId 修改数据
func (s *ShopHNCardModel)UpdateByHNCardIDsAndShopId(hNCardIDs []int,shopId int , data map[string]interface{}) (err error) {
	if len(hNCardIDs) == 0 || len(data) == 0 {
		return
	}
	_, err = s.Model.Where(map[string]interface{}{
		s.Field.F_hncard_id: []interface{}{"IN", hNCardIDs},
		s.Field.F_shop_id:shopId,
	}).Data(data).Update()

	return
}

//增加限时限次卡的销量
func (s *ShopHNCardModel) IncrSalesByShopidAndCardid(shopId, cardId int, step int) bool {
	if cardId <= 0 || step <= 0 || shopId <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:   shopId,
		s.Field.F_hncard_id: cardId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}
// 根据hncardIds获取多条数据
func (h *ShopHNCardModel)FindHNcardIdsAndBusId(hncardIds []int, shopId int , fields ...string) []map[string]interface{}{
	if len(hncardIds) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	return h.Model.Where(map[string]interface{}{
		h.Field.F_hncard_id: []interface{}{"IN", hncardIds},
		h.Field.F_is_del: cards.IS_BUS_DEL_no,
		h.Field.F_shop_id:   shopId,
	}).Select()
}

func (h *ShopHNCardModel) SelectRcardsByWherePage(where []base.WhereItem, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		h.Model.Field(fields)
	}
	if limit > 0 {
		h.Model.Limit(start, limit)
	}
	return h.Model.Where(where).OrderBy(h.Field.F_hncard_id + " DESC ").Select()
}