//ShopIcardModel
//2020-08-05 15:23:57

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

const (
	//SHOP_ICARD_STATUS_OFF 身份卡在门店的状态
	//SHOP_ICARD_STATUS_OFF     = 1 //下架
	//SHOP_ICARD_STATUS_ON      = 2 //上架
	//SHOP_ICARD_STATUS_DISABLE = 3 //禁用
)

//ShopIcardModel 表结构体
type ShopIcardModel struct {
	Model     *base.Model
	Field     ShopIcardModelField
	TableName string
	PK        string
}

//ShopIcardModelField 表字段
type ShopIcardModelField struct {
	F_id         string `default:"id"`
	F_shop_id    string `default:"shop_id"`
	F_icard_id   string `default:"icard_id"`
	F_status     string `default:"status"`
	F_is_del     string `default:"is_del"`
	F_del_time   string `default:"del_time"`
	F_under_time string `default:"under_time"`
	F_sales      string `default:"sales"`
	F_ctime      string `default:"ctime"`
}

//Init 初始化
func (s *ShopIcardModel) Init(ormer ...orm.Ormer) *ShopIcardModel {
	functions.ReflectModel(&s.Field)
	s.TableName = "shop_icard"
	s.PK = "id"
	s.Model = base.NewModel(s.TableName, ormer...)
	return s
}

//GetNumByShopidAndStatus GetNumByShopidAndStatus
func (s *ShopIcardModel) GetNumByShopidAndStatus(shopId, status int) int {
	if shopId <= 0 {
		return 0
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id: shopId,
		s.Field.F_status:  status,
	}).Count(s.Field.F_id)
}

//GetAll GetAll
func (s *ShopIcardModel) GetAll(condition Condition, fields ...string) (data []map[string]interface{}) {
	if len(fields) > 0 {
		s.Model.Field(fields)
	}
	if len(condition.Order) > 0 {
		s.Model.OrderBy(condition.Order)
	}
	s.Model.Where(condition.Where)
	return s.Model.Select()
}

//Insert 新增数据
func (s *ShopIcardModel) Insert(data map[string]interface{}) int {
	result, _ := s.Model.Data(data).Insert()
	return result
}

//InsertAll 批量新增数据
func (s *ShopIcardModel) InsertAll(data []map[string]interface{}) int {
	result, _ := s.Model.InsertAll(data)
	return result
}

//DeleteAll DeleteAll
func (s *ShopIcardModel) DeleteAll(where map[string]interface{}) int {
	res, _ := s.Model.Where(where).Delete()
	return res
}

func (s *ShopIcardModel) GetRcards(where map[string]interface{}, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		s.Model.Field(fields)
	}
	return s.Model.Where(where).Limit(start, limit).OrderBy(s.Field.F_id + " DESC ").Select()
}

func (s *ShopIcardModel) GetTotalNum(where map[string]interface{}) int {
	if len(where) == 0 {
		return 0
	}
	return s.Model.Where(where).Count(s.Field.F_id)
}

//Update update
func (s *ShopIcardModel) Update(data map[string]interface{}, pk interface{}) int {
	result, _ := s.Model.Data(data).Where([]base.WhereItem{
		{s.PK, pk},
	}).Update()
	return result
}

//UpdateAll UpdateAll
func (s *ShopIcardModel) UpdateAll(data map[string]interface{}, where map[string]interface{}) int {
	result, _ := s.Model.Data(data).Where(where).Update()
	return result
}

// 批量更新数据
func (s *ShopIcardModel) UpdateShopIdBySmids(icrads []int, shop_id int, data map[string]interface{}) bool {
	if len(icrads) == 0 || len(data) == 0 {
		return false
	}
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_icard_id:   []interface{}{"IN", icrads},
		s.Field.F_shop_id: shop_id,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}


//增加充值卡的销量
func (s *ShopIcardModel) IncrSalesByShopidAndIcardid(shopId, icardId int, step int) bool {
	if icardId <= 0 || step <= 0 || shopId <= 0 {
		return false
	}

	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:  shopId,
		s.Field.F_icard_id: icardId,
	}).Data(map[string]interface{}{
		s.Field.F_sales: []interface{}{"inc", step},
	}).Update()
	if err != nil {
		return false
	}

	return true
}

//获取单条数据
func (s *ShopIcardModel) GetByShopidAndRcardid(shopId, icardId int) map[string]interface{} {
	if shopId <= 0 || icardId <= 0 {
		return map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:  shopId,
		s.Field.F_icard_id: icardId,
	}).Find()
}

//根据主键ids 获取信息
func (s *ShopIcardModel) GetByIds(ids []int) (data []map[string]interface{}) {
	if len(ids) == 0 {
		return
	}

	data = s.Model.Where(map[string]interface{}{
		s.Field.F_id: []interface{}{"IN", ids},
	}).Select()

	return
}

//获取多条数据
func (s *ShopIcardModel) GetByShopidAndIcardids(shopId int, icardIds []int) []map[string]interface{} {
	if shopId <= 0 || len(icardIds) == 0 {
		return []map[string]interface{}{}
	}
	return s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:  shopId,
		s.Field.F_icard_id: []interface{}{"IN", icardIds},
	}).Select()
}

//根据充值卡ids获取数据
func (s *ShopIcardModel) GetByIcardids(icardIds []int) []map[string]interface{} {
	if len(icardIds) == 0 {
		return []map[string]interface{}{}
	}

	return s.Model.Where(map[string]interface{}{
		s.Field.F_icard_id: []interface{}{"IN", icardIds},
	}).Select()

}
//根据充值卡ids and shopId 获取数据
func (s *ShopIcardModel) GetByShopIDAndHNCardIDs(shopID int, icardIds []int) (data []map[string]interface{}) {
	if shopID <= 0 || len(icardIds) == 0 {
		return
	}
	data = s.Model.Where(map[string]interface{}{
		s.Field.F_shop_id:   shopID,
		s.Field.F_icard_id: []interface{}{"in", icardIds},
	}).Select()
	return
}

//修改
func (s *ShopIcardModel) UpdateByIds(ids []int, data map[string]interface{}) bool {
	if len(ids) == 0 {
		return false
	}

	_, err := s.Model.Where([]base.WhereItem{
		{
			Field: s.Field.F_id,
			Value: []interface{}{"IN", ids},
		},
	}).Data(data).Update()

	if err != nil {
		return false
	}

	return true
}
