//单项目基本信息表
//@author yangzhiwu<578154898@qq.com>
//@date 2020-03-27 14:22

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"time"
)

//表结构体
type SingleModel struct {
	Model *base.Model
	Field SingleModelField
}

//表字段
type SingleModelField struct {
	T_table         string `default:"single"`
	F_single_id     string `default:"single_id"`
	F_bus_id        string `default:"bus_id"`
	F_name          string `default:"name"`
	F_sort_desc     string `default:"sort_desc"`
	F_bind_id       string `default:"bind_id"`
	F_tag_ids       string `default:"tag_ids"`
	F_real_price    string `default:"real_price"`
	F_price         string `default:"price"`
	F_service_time  string `default:"service_time"`
	F_img_id        string `default:"img_id"`
	F_sales         string `default:"sales"`
	F_is_ground     string `default:"is_ground"`
	F_is_del        string `default:"is_del"`
	F_del_time      string `default:"del_time"`
	F_has_spec      string `default:"has_spec"`
	F_min_price     string `default:"min_price"`
	F_max_price     string `default:"max_price"`
	F_sale_shop_num string `default:"sale_shop_num"`
	F_subscribe     string `default:"subscribe"`
	F_ctime         string `default:"ctime"`
}

//初始化
func (s *SingleModel) Init(ormer ...orm.Ormer) *SingleModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *SingleModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//根据单项目id获取项目信息
//@param singleId int 单项目id
//@param fields ...[]string 查询的字段
//@return map[string]interface{}
func (s *SingleModel) GetBySingleId(singleId int, fields ...[]string) map[string]interface{} {
	if singleId == 0 {
		return map[string]interface{}{}
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
	}).Find()
}

//根据单项目id修改项目信息
func (s *SingleModel) UpdateBySingleid(singleId int, data map[string]interface{}) bool {
	if singleId <= 0 || len(data) == 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
	}).Data(data).Update()

	if err != nil {
		return false
	}

	return true
}

//根据单项目id和busId 删除单项目信息
func (s *SingleModel) DelBySingleId(singleId, busId int) bool {
	if singleId <= 0 || busId <= 0 {
		return false
	}
	_, err := s.Model.Where([]base.WhereItem{
		{s.Field.F_single_id, singleId},
		{s.Field.F_bus_id, busId},
	}).Data(map[string]interface{}{
		s.Field.F_is_del:   cards.IS_DEL_YES,
		s.Field.F_del_time: time.Now().Unix(),
	}).Update()

	if err != nil {
		return false
	}

	return true
}

//根据单项目ids修改项目信息
func (s *SingleModel) UpdateBySingleids(singleIds []int, data map[string]interface{}) bool {
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

//根据多个单项id，获取单项目信息
func (s *SingleModel) GetBySingleids(singleIds []int, fields ...[]string) []map[string]interface{} {
	if len(singleIds) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: []interface{}{"IN", singleIds},
	}).Select()

}

//根据多个单项id，获取单项目信息Li
func (s *SingleModel) GetList(where map[string]interface{}, fields ...[]string) []map[string]interface{} {
	if len(where) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	return s.Model.Where(where).Select()
}

//获取商家的一页数据
func (s *SingleModel) GetPageByBusId(busId, start, limit int, isDel string, isGround ...int) []map[string]interface{} {
	where := map[string]interface{}{
		s.Field.F_bus_id: busId,
	}
	if len(isGround) > 0 {
		where[s.Field.F_is_ground] = isGround[0]
	}
	if len(isDel) > 0 {
		where[s.Field.F_is_del] = isDel
	}

	return s.Model.Where(where).Limit(start, limit).OrderBy(s.Field.F_single_id + " DESC ").Select()
}

func (s *SingleModel) SelectSinglesByWherePage(where []base.WhereItem, start, limit int, fields ...[]string) []map[string]interface{} {
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

func (s *SingleModel)GetNumByWhere(where []base.WhereItem) int {
	return s.Model.Where(where).Count(s.Field.F_single_id)
}

//获取商家单项目总条数
func (s *SingleModel) GetNumByBusId(busId int, isDel string, isGround ...int) int {
	where := map[string]interface{}{
		s.Field.F_bus_id: busId,
	}
	if len(isGround) > 0 {
		where[s.Field.F_is_ground] = isGround[0]
	}
	if len(isDel) > 0 {
		where[s.Field.F_is_del] = isDel
	}

	return s.Model.Where(where).Count(s.Field.F_single_id)
}

//增加单项目的销量
func (s *SingleModel) IncrSalesBySingleid(singleId int, step int) bool {
	if singleId <= 0 || step <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}

//增加在售门店数量
func (s *SingleModel) IncrSaleshopnumBySingleid(singleId int, step int) bool {
	if singleId <= 0 || step <= 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
	}).Data(map[string]interface{}{
		s.Field.F_sale_shop_num: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}

//减少在售门店数量
func (s *SingleModel) DecrSaleshopnumBySingleid(singleId int, step int) bool {
	if singleId <= 0 || step <= 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
	}).Data(map[string]interface{}{
		s.Field.F_sale_shop_num: []interface{}{"dec", step},
	}).Update()

	if err != nil {
		return false
	}

	return true
}


