//ShopSingleSpecPriceModel
//2020-04-02 17:36:45

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

//表结构体
type ShopSingleSpecPriceModel struct {
	Model *base.Model
	Field ShopSingleSpecPriceModelField
}

//表字段
type ShopSingleSpecPriceModelField struct {
	T_table   string `default:"shop_single_spec_price"`
	F_shop_id string `default:"shop_id"`
	F_id      string `default:"id"`
	F_ss_id   string `default:"ss_id"`
	F_ssp_id  string `default:"ssp_id"`
	F_price   string `default:"price"`
	F_sales   string `default:"sales"`
	F_is_del   string `default:"is_del"`
}

//初始化
func (s *ShopSingleSpecPriceModel) Init() *ShopSingleSpecPriceModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table)
	return s
}

//新增数据
func (s *ShopSingleSpecPriceModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//批量添加
func (s *ShopSingleSpecPriceModel) InsertAll(data []map[string]interface{}) int {
	if len(data) <= 0 {
		return 0
	}
	result, _ := s.Model.InsertAll(data)
	return result
}

//根据商户规格价格表id获取信息
//@param singleSpecPriceId int 规格价格表id
//@param fields ...[]string 查询的字段
//@return map[string]interface{}
func (s *ShopSingleSpecPriceModel) GetShopSingleSpecPriceById(shopSingleSpecPriceId int, fields ...[]string) map[string]interface{} {
	if shopSingleSpecPriceId == 0 {
		return map[string]interface{}{}
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_id: shopSingleSpecPriceId,
	}).Find()
}

//删除门店多规格价格
func (s *ShopSingleSpecPriceModel) DelBySsids(ssids []int) bool {
	if len(ssids) == 0 {
		return false
	}
	if _, err := s.Model.Where([]base.WhereItem{{s.Field.F_ss_id, []interface{}{"IN", ssids}}}).
		Data(map[string]interface{}{
			s.Field.F_is_del: cards.IS_DEL_YES,
		}).Update(); err != nil {
			return false
	}
	return true
}

//获取门店多规格价格
func (s *ShopSingleSpecPriceModel) GetBySsid(ssid int, field ...string) []map[string]interface{} {
	if ssid <= 0 {
		return []map[string]interface{}{}
	}
	if len(field) > 0 {
		s.Model.Field(field)
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_ss_id: ssid,
	}).Select()
}

//获取门店多规格价格
func (s *ShopSingleSpecPriceModel) GetBySsids(ssids []int) []map[string]interface{} {
	if len(ssids) <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_ss_id: []interface{}{"IN", ssids},
	}).Select()
}

//获取门店多规格价格
func (s *ShopSingleSpecPriceModel) GetBySspids(sspids []int) []map[string]interface{} {
	if len(sspids) <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_ssp_id: []interface{}{"IN", sspids},
	}).Select()
}

//修改数据
func (s *ShopSingleSpecPriceModel) UpdateById(id int, data map[string]interface{}) bool {
	if id < 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_id: id,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//修改多条规格的数据
func (s *ShopSingleSpecPriceModel) UpdateBySspids(sspids []int, data map[string]interface{}) bool {
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_ssp_id: []interface{}{"IN", sspids},
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//获取单条数据
func (s *ShopSingleSpecPriceModel) GetBySsidAndSspid(ssId, sspId int) map[string]interface{} {
	if sspId <= 0 || ssId <= 0 {
		return map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_ss_id:  ssId,
		s.Field.F_ssp_id: sspId,
	}).Find()
}

//根据门店id和规格组合ids获取数据
func (s *ShopSingleSpecPriceModel) GetByShopidAndSspids(shopId int, sspIds []int) []map[string]interface{} {
	if shopId <= 0 || len(sspIds) <= 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id: shopId,
		s.Field.F_ssp_id:  []interface{}{"IN", sspIds},
	}).Select()
}

//获取门店多规格价格
func (s *ShopSingleSpecPriceModel) GetByShopSspIds(shopId int, sspIds []int) []map[string]interface{} {
	if shopId <= 0 {
		return []map[string]interface{}{}
	}
	rs := s.Model.Where([]base.WhereItem{
		{s.Field.F_shop_id, shopId},
		{s.Field.F_ssp_id, []interface{}{"IN", sspIds}},
	}).Select()
	return rs
}

//增加单项目的销量
func (s *ShopSingleSpecPriceModel) IncrSalesByShopidAndSspid(shopId, sspId int, step int) bool {
	if sspId <= 0 || step <= 0 || shopId <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id: shopId,
		s.Field.F_ssp_id:  sspId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}
