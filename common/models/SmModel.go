//套餐基本信息模型
//2020-04-14 19:20:32

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"git.900sui.cn/kc/rpcCards/common/tools"
)

//表结构体
type SmModel struct {
	Model *base.Model
	Field SmModelField
	tools.CardStatus
}

//表字段
type SmModelField struct {
	T_table                 string `default:"sm"`
	F_sm_id                 string `default:"sm_id"`
	F_bus_id                string `default:"bus_id"`
	F_bind_id               string `default:"bind_id"`
	F_name                  string `default:"name"`
	F_sort_desc             string `default:"sort_desc"`
	F_real_price            string `default:"real_price"`
	F_price                 string `default:"price"`
	F_service_period        string `default:"service_period"`
	F_has_give_signle       string `default:"has_give_signle"`
	F_is_permanent_validity string `default:"is_permanent_validity"`
	F_is_ground             string `default:"is_ground"`
	F_under_time            string `default:"under_time"`
	F_img_id                string `default:"img_id"`
	F_sales                 string `default:"sales"`
	F_sale_shop_num         string `default:"sale_shop_num"`
	F_validcount            string `default:"validcount"`
	F_ctime                 string `default:"ctime"`
	F_is_del                string `default:"is_del"`
	F_del_time              string `default:"del_time"`
}

//初始化
func (s *SmModel) Init(ormer ...orm.Ormer) *SmModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *SmModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//根据套餐id获取单条套餐数据
func (s *SmModel) GetBySmid(smId int, fields ...[]string) map[string]interface{} {
	if smId <= 0 {
		return map[string]interface{}{}
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id:  smId,
		s.Field.F_is_del: s.NotDelStatus(),
	}).Find()
}

//批量修改
func (s *SmModel) UpdateBySmids(smIds []int, data map[string]interface{}) bool {
	if len(smIds) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: []interface{}{"IN", smIds},
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//根据smid修改数据
func (s *SmModel) UpdateBySmid(smId int, data map[string]interface{}) bool {
	if smId <= 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: smId,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//增加套餐的销量
func (s *SmModel) IncrSalesBySmid(smId int, step int) bool {
	if smId <= 0 || step <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: smId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}

//增加在售店铺数量
func (s *SmModel) IncrSaleShopNumBySmid(smid, step int) bool {
	if smid <= 0 || step <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: smid,
	}).Data(map[string]interface{}{
		s.Field.F_sale_shop_num: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}
	return true
}

//减少在售店铺数量
func (s *SmModel) DecrSaleShopNumBySmid(smid, step int) bool {
	if smid <= 0 || step <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id: smid,
	}).Data(map[string]interface{}{
		s.Field.F_sale_shop_num: []interface{}{"dec", step},
	}).Update()
	if err != nil {
		return false
	}
	return true
}

//获取企业的套餐列表
func (s *SmModel) GetPageByBusId(busId int, start, limit int, isGround ...int) []map[string]interface{} {
	if busId <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}
	where := map[string]interface{}{
		s.Field.F_bus_id: busId,
		s.Field.F_is_del: s.NotDelStatus(),
	}
	if len(isGround) > 0 {
		where[s.Field.F_is_ground] = isGround[0]
	}
	return s.Model.Where(where).OrderBy(s.Field.F_sm_id+" DESC ").Limit(start, limit).Select()
}

func (s *SmModel) SelectSmsByWherePage(where []base.WhereItem, start, limit int, fields ...string) []map[string]interface{} {
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

func (s *SmModel) GetNumByWhere(where []base.WhereItem) int {
	return s.Model.Where(where).Count(s.Field.F_sm_id)
}

//获取商家的套餐数量
func (s *SmModel) GetNumByBusId(busId int, isGround ...int) int {
	if busId <= 0 {
		return 0
	}
	where := map[string]interface{}{
		s.Field.F_bus_id: busId,
		s.Field.F_is_del: s.NotDelStatus(),
	}
	if len(isGround) > 0 {
		where[s.Field.F_is_ground] = isGround[0]
	}
	return s.Model.Where(where).Count(s.Field.F_sm_id)
}

//根据smids批量获取数据
func (s *SmModel) GetBySmids(smIds []int, fields ...[]string) []map[string]interface{} {
	if len(smIds) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id:  []interface{}{"IN", smIds},
		s.Field.F_is_del: s.NotDelStatus(),
	}).Select()
}

func (s *SmModel) Select(where map[string]interface{}, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		s.Model.Field(fields)
	}
	where[s.Field.F_is_del] = s.NotDelStatus()
	return s.Model.Where(where).Select()
}
