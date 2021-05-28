//SingleSpecPriceModel
//2020-03-27 15:39:48

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

//表结构体
type SingleSpecPriceModel struct {
	Model *base.Model
	Field SingleSpecPriceModelField
}

//表字段
type SingleSpecPriceModelField struct {
	T_table     string `default:"single_spec_price"`
	F_ssp_id    string `default:"ssp_id"`
	F_single_id string `default:"single_id"`
	F_spec_ids  string `default:"spec_ids"`
	F_hash      string `default:"hash"`
	F_price     string `default:"price"`
	F_sales     string `default:"sales"`
	F_is_del    string `default:"is_del"`
}

//初始化
func (s *SingleSpecPriceModel) Init() *SingleSpecPriceModel {
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table)
	return s
}

//新增数据
func (s *SingleSpecPriceModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//批量新增数据
func (s *SingleSpecPriceModel) InsertAll(data []map[string]interface{}) int {
	result, _ := s.Model.InsertAll(data)
	return result
}

//根据规格价格表id获取信息
//@param singleSpecPriceId int 规格价格表id
//@param fields ...[]string 查询的字段
//@return map[string]interface{}
func (s *SingleSpecPriceModel) GetSingleSpecPriceById(singleSpecPriceId int, fields ...[]string) map[string]interface{} {
	if singleSpecPriceId == 0 {
		return map[string]interface{}{}
	}
	if len(fields) > 0 {
		s.Model.Field(fields[0])
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_ssp_id: singleSpecPriceId,
	}).Find()
}

//获取不同规格的价格
func (s *SingleSpecPriceModel) GetBySingleid(singleId int) []map[string]interface{} {
	return s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
		s.Field.F_is_del:    cards.SPEC_PRICE_IS_DEL_no,
	}).Select()
}

//获取不同规格的价格
func (s *SingleSpecPriceModel) GetBySingleids(singleIds []int) []map[string]interface{} {
	return s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: []interface{}{"IN", singleIds},
		s.Field.F_is_del:    cards.SPEC_PRICE_IS_DEL_no,
	}).Select()
}

//获取不同规格的价格
func (s *SingleSpecPriceModel) GetBySspids(sspIds []int, field ...string) []map[string]interface{} {
	if len(sspIds) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(field) > 0 {
		s.Model.Field(field)
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_ssp_id: []interface{}{"IN", sspIds},
		s.Field.F_is_del: cards.SPEC_PRICE_IS_DEL_no,
	}).Select()
}

//根据多个hash值查询
func (s *SingleSpecPriceModel) GetByHashs(hashs []string) []map[string]interface{} {
	return s.Model.Where(map[string]interface{}{
		s.Field.F_hash:   []interface{}{"IN", hashs},
		s.Field.F_is_del: cards.SPEC_PRICE_IS_DEL_no,
	}).Select()
}

//修改多条规格的数据
func (s *SingleSpecPriceModel) UpdateBySspids(sspids []int, data map[string]interface{}) bool {
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_ssp_id: []interface{}{"IN", sspids},
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//修改单条规格数据
func (s *SingleSpecPriceModel) UpdateBySspid(sspid int, data map[string]interface{}) bool {
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_ssp_id: sspid,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

//增加单项目的销量
func (s *SingleSpecPriceModel) IncrSalesBySspid(sspId int, step int) bool {
	if sspId <= 0 || step <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_ssp_id: sspId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}

//修改单条规格数据
func (s *SingleSpecPriceModel) UpdateBySingleid(singleId int, data map[string]interface{}) bool {
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_single_id: singleId,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

// 根据规格sspId 查询
func (s *SingleSpecPriceModel)FindBySspId(sspId int ) map[string]interface{} {
	return s.Model.Where(map[string]interface{}{
		s.Field.F_ssp_id:sspId,
		s.Field.F_is_del:cards.SPEC_PRICE_IS_DEL_no,
	}).Find()
}
