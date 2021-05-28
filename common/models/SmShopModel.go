//套餐的适用门店模型
//2020-04-14 19:20:32

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

//表结构体
type SmShopModel struct {
	Model *base.Model
	Field SmShopModelField
}

//表字段
type SmShopModelField struct{
	T_table	string	`default:"sm_shop"`
	F_id	string	`default:"id"`
	F_sm_id	string	`default:"sm_id"`
	F_bus_id	string	`default:"bus_id"`
	F_shop_id	string	`default:"shop_id"`
}

//初始化
func (s *SmShopModel) Init(ormer ...orm.Ormer) *SmShopModel{
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table, ormer...)
	return s
}

//新增数据
func (s *SmShopModel) Insert(data map[string]interface{}) (int){
	result,_ := s.Model.Data(data).Insert()
	return result
}

//批量添加
func (s *SmShopModel) InsertAll( data []map[string]interface{}) int  {
	if len(data) == 0{
		return 0
	}

	r, _ := s.Model.InsertAll( data )
	return r
}

//根据ids批量修改
func (s *SmShopModel) UpdateByIds ( ids []int, data map[string]interface{} ) bool {
	if len(ids) == 0 || len(data) == 0{
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_id:[]interface{}{"IN", ids},
	}).Data(data).Update()

	if err != nil{
		return false
	}
	return true
}

//根据套餐ids 获取所有可适用门店记录
func (s *SmShopModel) GetBySmids( smIds []int ) []map[string]interface{}  {
	if len(smIds) <= 0{
		return []map[string]interface{}{}
	}

	return  s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id:[]interface{}{"IN", smIds},
	}).Select()
}

//根据套餐id 获取所有可适用门店记录
func (s *SmShopModel) GetBySmId( smId int ) []map[string]interface{}  {
	if smId <= 0{
		return []map[string]interface{}{}
	}

	return  s.Model.Field([]string{s.Field.F_shop_id,s.Field.F_bus_id}).Where(map[string]interface{}{
		s.Field.F_sm_id: smId,
	}).Select()
}

//根据套餐id和busId 获取所有可适用门店记录
func (s *SmShopModel) GetBySmIdByBusId( smId int ,busId int ) []map[string]interface{}  {
	if smId <= 0{
		return []map[string]interface{}{}
	}
	return  s.Model.Field([]string{s.Field.F_shop_id,s.Field.F_bus_id}).Where(map[string]interface{}{
		s.Field.F_sm_id: smId,
		s.Field.F_bus_id:busId,
	}).Select()
}


//根据门店id获取门店可添加的套餐列表
func (s *SmShopModel) GetPageByShopId( busId, shopId, start, limit int ) []map[string]interface{} {
	if busId <= 0 || shopId <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}

	return  s.Model.Where(map[string]interface{}{
		s.Field.F_bus_id:busId,
		s.Field.F_shop_id:[]interface{}{"IN", []int{0, shopId}},
	}).OrderBy(s.Field.F_id+" DESC ").Limit( start, limit ).Select()
}

//根据smids和shopid获取适用门店的套餐数据
func (s *SmShopModel) GetByShopIdAndSmids( busId, shopId int , smIds []int  ) []map[string]interface{} {
	if busId <= 0 || shopId <= 0 || len(smIds) == 0 {
		return []map[string]interface{}{}
	}
	return  s.Model.Where(map[string]interface{}{
		s.Field.F_bus_id:busId,
		s.Field.F_shop_id:[]interface{}{"IN", []int{0, shopId}},
		s.Field.F_sm_id:[]interface{}{"IN", smIds},
	}).Select()

}

//根据门店id获取门店可添加的套餐总数量
func (s *SmShopModel) GetNumByShopId( busId, shopId int ) int {
	if busId <= 0 || shopId <= 0 {
		return 0
	}
	return  s.Model.Where(map[string]interface{}{
		s.Field.F_bus_id:busId,
		s.Field.F_shop_id:[]interface{}{"IN", []int{0, shopId}},
	}).Count(s.Field.F_id)
}

//批量删除
func (s *SmShopModel) DelBySmIds(smIds []int) bool {
	if len(smIds) == 0{
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_sm_id:[]interface{}{"IN", smIds},
	}).Delete()
	if err != nil{
		return false
	}
	return true
}