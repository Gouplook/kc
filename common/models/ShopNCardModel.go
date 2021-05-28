//ShopNCardModel
//2020-04-20 13:37:51

package models

import (
	"fmt"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"git.900sui.cn/kc/rpcCards/common/tools"
	"github.com/wendal/errors"
)

//表结构体
type ShopNCardModel struct {
	Model *base.Model
	Field ShopNCardModelField
	tools.CardStatus
}

//表字段
type ShopNCardModelField struct {
	T_table      string `default:"shop_ncard"`
	F_id         string `default:"id"`
	F_shop_id    string `default:"shop_id"`
	F_ncard_id   string `default:"ncard_id"`
	F_status     string `default:"status"`
	F_under_time string `default:"under_time"`
	F_sales      string `default:"sales"`
	F_ctime      string `default:"ctime"`
	F_is_del     string `default:"is_del"`
	F_del_time   string `default:"del_time"`
}

//初始化
func (s *ShopNCardModel) Init(ormer ...orm.Ormer) *ShopNCardModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *ShopNCardModel) Insert(data map[string]interface{}) (shopNCardID int, err error) {
	if result, insertErr := s.Model.Data(data).Insert(); insertErr != nil || result == 0 {
		err = errors.New("insert failed")
	} else {
		shopNCardID = result
	}
	return
}

//批量添加
func (s *ShopNCardModel) InsertAll(data []map[string]interface{}) (result int, err error) {
	result, err = s.Model.InsertAll(data)
	return
}

//获取多条信息
func (s *ShopNCardModel) GetByShopIDAndNCardIDs(shopID int, nCardID []int) (data []map[string]interface{}) {
	if shopID <= 0 || len(nCardID) == 0 {
		return
	}
	data = s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:  shopID,
		s.Field.F_ncard_id: []interface{}{"in", nCardID},
		s.Field.F_is_del:   s.NotDelStatus(),
	}).Select()
	return
}

//获取多条信息(含删除的的）
func (s *ShopNCardModel) GetByShopIdByNCardIDs(shopID int, nCardID []int) (data []map[string]interface{}) {
	if shopID <= 0 || len(nCardID) == 0 {
		return
	}
	data = s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:  shopID,
		s.Field.F_ncard_id: []interface{}{"in", nCardID},
	}).Select()
	return
}

//获取门店的限次卡列表
func (s *ShopNCardModel) GetPageByShopID(shopID, start, limit, status int) (data []map[string]interface{}) {
	if shopID <= 0 || start < 0 || limit <= 0 {
		return
	}
	whereMap := map[string]interface{}{
		s.Field.F_shop_id: shopID,
		s.Field.F_is_del:  s.NotDelStatus(),
	}
	if status > 0 {
		whereMap[s.Field.F_status] = status
	}

	data = s.Model.Where(whereMap).OrderBy(fmt.Sprintf("%s %s", s.Field.F_id, "desc")).Limit(start, limit).Select()

	return
}

//获取门店的限次卡列表
func (s *ShopNCardModel) GetList(where map[string]interface{}, start, limit int) (data []map[string]interface{}) {
	if len(where) <= 0 || start < 0 || limit <= 0 {
		return
	}
	where[s.Field.F_is_del] = s.NotDelStatus()
	data = s.Model.Where(where).OrderBy(s.Field.F_id+" desc ").Limit(start, limit).Select()
	return
}

//获取门店限次卡数量
func (s *ShopNCardModel) GetTotalNum(where map[string]interface{}) (result int) {
	if len(where) <= 0 {
		return
	}
	where[s.Field.F_is_del] = s.NotDelStatus()

	result = s.Model.Where(where).Count(s.Field.F_id)
	return
}

//获取门店限次卡数量
func (s *ShopNCardModel) GetNumByShopID(shopID int, status int) (result int) {
	if shopID <= 0 {
		return
	}

	whereMap := map[string]interface{}{
		s.Field.F_shop_id: shopID,
		s.Field.F_is_del:  s.NotDelStatus(),
	}
	if status > 0 {
		whereMap[s.Field.F_status] = status
	}
	result = s.Model.Where(whereMap).Count(s.Field.F_id)

	return
}

//批量修改
func (s *ShopNCardModel) UpdateByIDs(ids []int, data map[string]interface{}) (err error) {
	if len(ids) == 0 || len(data) == 0 {
		return
	}
	_, err = s.Model.Where(map[string]interface{}{
		s.Field.F_id: []interface{}{"IN", ids},
	}).Data(data).Update()

	return
}

//根据主键ids 获取信息
func (s *ShopNCardModel) GetByIDs(ids []int) (data []map[string]interface{}) {
	if len(ids) == 0 {
		return
	}

	data = s.Model.Where(map[string]interface{}{
		s.Field.F_id:     []interface{}{"IN", ids},
		s.Field.F_is_del: s.NotDelStatus(),
	}).Select()

	return
}

//根据限次卡ids获取数据
func (s *ShopNCardModel) GetByNCardIDs(nCardIDs []int) (data []map[string]interface{}) {
	if len(nCardIDs) == 0 {
		return
	}

	data = s.Model.Where(map[string]interface{}{
		s.Field.F_ncard_id: []interface{}{"IN", nCardIDs},
		s.Field.F_is_del:   s.NotDelStatus(),
	}).Select()

	return
}

//根据限次卡ids修改数据
func (s *ShopNCardModel) UpdateByNCardIDs(nCardIDs []int, data map[string]interface{}) (err error) {
	if len(nCardIDs) == 0 || len(data) == 0 {
		return
	}
	_, err = s.Model.Where(map[string]interface{}{
		s.Field.F_ncard_id: []interface{}{"IN", nCardIDs},
		s.Field.F_is_del:   s.NotDelStatus(),
	}).Data(data).Update()

	return
}

//根据商店id及限次卡ids修改数据
func (s *ShopNCardModel) UpdateShopIdByNCardIDs(nCardIDs []int, shopId int, data map[string]interface{}) (err error) {
	if len(nCardIDs) == 0 || len(data) == 0 {
		return
	}
	_, err = s.Model.Where(map[string]interface{}{
		s.Field.F_ncard_id: []interface{}{"IN", nCardIDs},
		s.Field.F_shop_id:  shopId,
	}).Data(data).Update()
	return
}

//增加限次卡的销量
func (s *ShopNCardModel) IncrSalesByShopidAndCardid(shopId, cardId int, step int) bool {
	if cardId <= 0 || step <= 0 || shopId <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:  shopId,
		s.Field.F_ncard_id: cardId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}


func (s *ShopNCardModel) SelectRcardsByWherePage(where []base.WhereItem, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		s.Model.Field(fields)
	}
	if limit > 0 {
		s.Model.Limit(start, limit)
	}
	return s.Model.Where(where).OrderBy(s.Field.F_ncard_id + " DESC ").Select()
}