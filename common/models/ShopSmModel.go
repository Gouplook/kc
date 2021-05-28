//已添加到门店的套餐
//2020-04-14 19:20:32

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"git.900sui.cn/kc/rpcCards/common/tools"
)

//表结构体
type ShopSmModel struct {
	Model *base.Model
	Field ShopSmModelField
	tools.CardStatus
}

//表字段
type ShopSmModelField struct {
	T_table      string `default:"shop_sm"`
	F_id         string `default:"id"`
	F_shop_id    string `default:"shop_id"`
	F_sm_id      string `default:"sm_id"`
	F_status     string `default:"status"`
	F_under_time string `default:"under_time"`
	F_sales      string `default:"sales"`
	F_ctime      string `default:"ctime"`
	F_is_del     string `default:"is_del"`
	F_del_time   string `default:"del_time"`
}

//初始化
func (s *ShopSmModel) Init(ormer ...orm.Ormer) *ShopSmModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *ShopSmModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//批量添加
func (s *ShopSmModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}
	result, _ := s.Model.InsertAll(data)
	return result
}

//获取单条数据
func (s *ShopSmModel) GetByShopidAdSmid(shopId, smId int) map[string]interface{} {
	if shopId <= 0 || smId <= 0 {
		return map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id: shopId,
		s.Field.F_sm_id:   smId,
		s.Field.F_is_del:  s.NotDelStatus(),
	}).Find()
}

//获取多条数据
func (s *ShopSmModel) GetByShopidAdSmids(shopId int, smIds []int) []map[string]interface{} {
	if shopId <= 0 || len(smIds) == 0 {
		return []map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id: shopId,
		s.Field.F_sm_id:   []interface{}{"IN", smIds},
		s.Field.F_is_del:  s.NotDelStatus(),
	}).Select()
}

func (s *ShopSmModel) GetByShopIdBySmids(shopId int, smIds []int) []map[string]interface{} {
	if shopId <= 0 || len(smIds) == 0 {
		return []map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id: shopId,
		s.Field.F_sm_id:   []interface{}{"IN", smIds},
	}).Select()
}
//获取门店的套餐列表
func (s *ShopSmModel) GetPageByShopId(shopid, start, limit, status int) []map[string]interface{} {
	if shopid <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}
	wheremap := map[string]interface{}{
		s.Field.F_shop_id: shopid,
		s.Field.F_is_del:  s.NotDelStatus(),
	}
	if status > 0 {
		wheremap[s.Field.F_status] = status
	}
	return s.Model.Where(wheremap).OrderBy(s.Field.F_id+" DESC ").Limit(start, limit).Select()
}

//获取门店的套餐列表
func (s *ShopSmModel) GetSms(where map[string]interface{}, start, limit int) []map[string]interface{} {
	if len(where) <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}
	where[s.Field.F_is_del] = s.NotDelStatus()
	return s.Model.Where(where).OrderBy(s.Field.F_id+" DESC ").Limit(start, limit).Select()
}

//获取门店套餐数量
func (s *ShopSmModel) GetNumByShopId(shopId int, status int) int {
	if shopId <= 0 {
		return 0
	}

	wheremap := map[string]interface{}{
		s.Field.F_shop_id: shopId,
		s.Field.F_is_del:  s.NotDelStatus(),
	}
	if status > 0 {
		wheremap[s.Field.F_status] = status
	}

	return s.Model.Where(wheremap).Count(s.Field.F_id)
}
func (s *ShopSmModel) GetTotalNum(where map[string]interface{}) int {
	if len(where) <= 0 {
		return 0
	}
	return s.Model.Where(where).Count(s.Field.F_id)
}

//批量修改
func (s *ShopSmModel) UpdateByIds(ids []int, data map[string]interface{}) bool {
	if len(ids) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_id:     []interface{}{"IN", ids},
		s.Field.F_is_del: s.NotDelStatus(),
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//根据主键ids 获取信息
func (s *ShopSmModel) GetByIds(ids []int) []map[string]interface{} {
	if len(ids) == 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_id:     []interface{}{"IN", ids},
		s.Field.F_is_del: s.NotDelStatus(),
	}).Select()
}

//根据套餐ids获取数据
func (s *ShopSmModel) GetBySmids(smIds []int) []map[string]interface{} {
	if len(smIds) == 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id:  []interface{}{"IN", smIds},
		s.Field.F_is_del: s.NotDelStatus(),
	}).Select()

}

//根据套餐ids修改数据
func (s *ShopSmModel) UpdateBySmids(smIds []int, data map[string]interface{}) bool {
	if len(smIds) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id:  []interface{}{"IN", smIds},
		s.Field.F_is_del: s.NotDelStatus(),
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//增加套餐的销量
func (s *ShopSmModel) IncrSalesByShopidAndSmid(shopId, smId int, step int) bool {
	if smId <= 0 || step <= 0 || shopId <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id: shopId,
		s.Field.F_sm_id:   smId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}

//分店-增加批量软删除
func (s *ShopSmModel) UpdateByShopSmids(smIds []int, shop_id int, data map[string]interface{}) bool {
	if len(smIds) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id:   []interface{}{"IN", smIds},
		s.Field.F_shop_id: shop_id,
		s.Field.F_is_del:  s.NotDelStatus(),
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

// 批量更新数据
func (s *ShopSmModel) UpdateShopIdBySmids(smIds []int, shop_id int, data map[string]interface{}) bool {
	if len(smIds) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id:   []interface{}{"IN", smIds},
		s.Field.F_shop_id: shop_id,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}


func (s *ShopSmModel) SelectRcardsByWherePage(where []base.WhereItem, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		s.Model.Field(fields)
	}
	if limit > 0 {
		s.Model.Limit(start, limit)
	}
	return s.Model.Where(where).OrderBy(s.Field.F_sm_id + " DESC ").Select()
}