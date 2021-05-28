//ShopSingleModel
//2020-04-10 17:01:23

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"time"
)

//表结构体
type ShopSingleModel struct {
	Model *base.Model
	Field ShopSingleModelField
}

const (
	STATUS_OFF_SALE = 1 // 下架
	STATUS_ON_SALE  = 2 // 上架
	STATUS_DISABLE  = 3 // 禁用
)

//表字段
type ShopSingleModelField struct {
	T_table              string `default:"shop_single"`
	F_ss_id              string `default:"ss_id"`
	F_shop_id            string `default:"shop_id"`
	F_single_id          string `default:"single_id"`
	F_name               string `default:"name"`
	F_changed_real_price string `default:"changed_real_price"`
	F_changed_min_price  string `default:"changed_min_price"`
	F_changed_max_price  string `default:"changed_max_price"`
	F_status             string `default:"status"`
	F_is_del     string `default:"is_del"`
	F_under_time         string `default:"under_time"`
	F_del_time           string `default:"del_time"`
	F_sales              string `default:"sales"`
	F_ctime              string `default:"ctime"`
}

//初始化
func (s *ShopSingleModel) Init(ormer ...orm.Ormer) *ShopSingleModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *ShopSingleModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//批量添加
func (s *ShopSingleModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}

	result, _ := s.Model.InsertAll(data)
	return result
}

//根据单项目id获取项目信息
//@param singleId int 单项目id
//@param fields ...[]string 查询的字段
//@return map[string]interface{}
func (s *ShopSingleModel) GetShopSingleById(shopSingleId int, fields ...[]string) map[string]interface{} {
	if shopSingleId == 0 {
		return map[string]interface{}{}
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_ss_id: shopSingleId,
	}).Find()
}

//获取门店单项目信息
func (s *ShopSingleModel) GetByShopidAndSingleid(shopId, singleId int, fields ...[]string) map[string]interface{} {
	if shopId == 0 || singleId == 0 {
		return map[string]interface{}{}
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	return s.Model.Where([]base.WhereItem{{s.Field.F_single_id, singleId},
		{s.Field.F_shop_id, shopId}}).Find()

}

//获取门店多个单项目信息
func (s *ShopSingleModel) GetByShopidAndSingleids(shopId int, singleIds []int, fields ...[]string) []map[string]interface{} {
	if shopId == 0 || len(singleIds) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: []interface{}{"IN", singleIds},
		s.Field.F_shop_id:   shopId,
	}).Select()
}

//获取门店的单项目
func (s *ShopSingleModel) GetByShopid(shopId, start, limit int, status,isDel string,singleIds []int) []map[string]interface{} {
	if shopId <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}
	whereMap := map[string]interface{}{
		s.Field.F_shop_id: shopId,
	}
	if len(status) > 0 {
		whereMap[s.Field.F_status] = status
	}
	if len(isDel)>0{
		whereMap[s.Field.F_is_del]=isDel
	}
	if len(singleIds)>0{
		whereMap[s.Field.F_single_id]=[]interface{}{"IN",singleIds}
	}
	return s.Model.Where(whereMap).OrderBy(s.Field.F_ss_id+" DESC ").Limit(start, limit).Select()
}


func (s *ShopSingleModel) SelectRcardsByWherePage(where []base.WhereItem, start, limit int, fields ...[]string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	if limit > 0 {
		s.Model.Limit(start, limit)
	}
	return s.Model.Where(where).OrderBy(s.Field.F_single_id + " DESC ").Select()
}

//获取门店的单项目
func (s *ShopSingleModel) GetSingles(where map[string]interface{}, start, limit int) []map[string]interface{} {
	if len(where) <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}
	return s.Model.Where(where).OrderBy(s.Field.F_ss_id+" DESC ").Limit(start, limit).Select()
}

//获取门店的单项目
func (s *ShopSingleModel) GetSinglesOrderBy(where map[string]interface{}, start, limit int, orderby ...string) []map[string]interface{} {
	if len(where) <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}
	orderbyStr := orderby[0]
	if len(orderbyStr) == 0 {
		orderbyStr = s.Field.F_ss_id
	}
	return s.Model.Where(where).OrderBy(orderbyStr+" DESC ").Limit(start, limit).Select()
}

//获取门店的单项目数量
func (s *ShopSingleModel) GetNumByShopid(shopId int , status ,isDel string,singleIds []int) int {
	if shopId <= 0 {
		return 0
	}
	whereMap := map[string]interface{}{
		s.Field.F_shop_id: shopId,
	}
	if len(status) > 0 {
		whereMap[s.Field.F_status] = status
	}
	if len(isDel)>0{
		whereMap[s.Field.F_is_del]=isDel
	}
	if len(singleIds)>0{
		whereMap[s.Field.F_single_id]=[]interface{}{"IN",singleIds}
	}
	return s.Model.Where(whereMap).Count(s.Field.F_ss_id)
}

//获取门店的单项目数量
func (s *ShopSingleModel) GetTotal(where map[string]interface{}) int {
	if len(where) == 0 {
		return 0
	}
	return s.Model.Where(where).Count(s.Field.F_ss_id)
}

func (s *ShopSingleModel) GetTotalWhere(where []base.WhereItem) int {
	if len(where) == 0 {
		return 0
	}

	return s.Model.Where(where).Count(s.Field.F_ss_id)
}

//修改数据
func (s *ShopSingleModel) UpDateBySsid(ssid int, data map[string]interface{}) bool {
	if ssid <= 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_ss_id: ssid,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true

}

//批量修改数据
func (s *ShopSingleModel) UpDateBySsids(ssids []int, data map[string]interface{}) bool {
	if len(ssids) <= 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_ss_id: []interface{}{"IN", ssids},
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true

}

//根据多个单项目id 批量修改
func (s *ShopSingleModel) UpDateBySingleids(singleIds []int, data map[string]interface{}) bool {
	if len(singleIds) <= 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: []interface{}{"IN", singleIds},
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//根据多个单项目ids和shopId 批量修改
func (s *ShopSingleModel) UpDateBySingleidsAndShopId(singleIds []int, shopId int ,data map[string]interface{}) bool {
	if len(singleIds) <= 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: []interface{}{"IN", singleIds},
		s.Field.F_shop_id:shopId,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//根据单项目id和shopId 删除单项目
func (s *ShopSingleModel) DelBySingleIdAndShopId(singleId, shopId int) bool {
	if singleId <= 0 || shopId <= 0 {
		return false
	}
	_, err := s.Model.Where([]base.WhereItem{{s.Field.F_single_id, singleId},
		{s.Field.F_shop_id, shopId}}).Data(map[string]interface{}{
		s.Field.F_status: cards.STATUS_DISABLE, s.Field.F_del_time: time.Now().Unix(),
	}).Update()
	if err != nil {
		return false
	}
	return true
}

//根据单项目id 删除单项目
func (s *ShopSingleModel) DelBySingleId(singleId int) bool {
	if singleId <= 0 {
		return false
	}
	_, err := s.Model.Where([]base.WhereItem{{s.Field.F_single_id, singleId}}).Data(map[string]interface{}{
		s.Field.F_status: cards.STATUS_DISABLE, s.Field.F_del_time: time.Now().Unix(),
	}).Update()
	if err != nil {
		return false
	}
	return true
}

//根据主键获取信息
func (s *ShopSingleModel) GetBySsids(ssIds []int) []map[string]interface{} {
	if len(ssIds) == 0 {
		return []map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_ss_id: []interface{}{"IN", ssIds},
	}).Select()
}

//根据主键获取信息
func (s *ShopSingleModel) GetByShopIdAndSingleIds(shopId int, singleIds []int, status ...int) []map[string]interface{} {
	if len(singleIds) == 0 {
		return []map[string]interface{}{}
	}
	where := []base.WhereItem{
		{s.Field.F_shop_id, shopId},
		{s.Field.F_single_id, []interface{}{"IN", singleIds}},
	}
	if len(status) > 0 {
		where = append(where, base.WhereItem{s.Field.F_status, status[0]})
	}
	rs := s.Model.Where(where).Select()
	return rs
}

//获取指定单项目id在所有门店ssid
func (s *ShopSingleModel) GetBySingleid(singleId int) []map[string]interface{} {
	if singleId <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
	}).Select()
}

//根据指定单项目ids，获取门店项目信息
func (s *ShopSingleModel) GetBySingleids(singleIds []int) []map[string]interface{} {
	if len(singleIds) <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: []interface{}{"IN", singleIds},
	}).Select()
}

//门店单项目销量增加
func (s *ShopSingleModel) IncrSalesBySingleid(ssId int, step int) bool {
	if ssId <= 0 || step <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_ss_id: ssId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}

//门店单项目销量增加
func (s *ShopSingleModel) IncrSalesBySingleidAndShopid(singleId, shopId int, step int) bool {
	if singleId <= 0 || shopId <= 0 || step <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
		s.Field.F_shop_id:   shopId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}
