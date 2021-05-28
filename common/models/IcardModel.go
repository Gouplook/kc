//IcardModel
//2020-08-05 15:23:57

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
)

const (
	//是否上架 0=否 1=是
	//IS_GROUND_OFF = 0
	//IS_GROUND_ON  = 1
)

//IcardModel 表结构体
type IcardModel struct {
	Model     *base.Model
	Field     IcardModelField
	TableName string
	PK        string
}

//Condition 条件
type Condition struct {
	Where  map[string]interface{}
	Offset int
	Limit  int
	Order  string
}

//IcardModelField 表字段
type IcardModelField struct {
	F_icard_id        string `default:"icard_id"`
	F_bus_id          string `default:"bus_id"`
	F_bind_id         string `default:"bind_id"`
	F_name            string `default:"name"`
	F_sort_desc       string `default:"sort_desc"`
	F_real_price      string `default:"real_price"`
	F_price           string `default:"price"`
	F_service_period  string `default:"service_period"`
	F_has_give_signle string `default:"has_give_signle"`
	F_is_ground       string `default:"is_ground"`
	F_is_del          string `default:"is_del"`
	F_del_time        string `default:"del_time"`
	F_under_time      string `default:"under_time"`
	F_img_id          string `default:"img_id"`
	F_sales           string `default:"sales"`
	F_sale_shop_num   string `default:"sale_shop_num"`
	F_ctime           string `default:"ctime"`
}

//Init 初始化
func (i *IcardModel) Init(ormer ...orm.Ormer) *IcardModel {
	functions.ReflectModel(&i.Field)
	i.TableName = "icard"
	i.PK = "icard_id"
	i.Model = base.NewModel(i.TableName, ormer...)
	return i
}

//GetOneByPK 获取单条数据
func (i *IcardModel) GetOneByPK(id int) (reply map[string]interface{}) {
	reply = make(map[string]interface{})
	if id <= 0 {
		return
	}
	reply = i.Model.Where([]base.WhereItem{
		{i.Field.F_icard_id, id},
	}).Find()
	return
}

//Insert 新增数据
func (i *IcardModel) Insert(data map[string]interface{}) int {
	result, _ := i.Model.Data(data).Insert()
	return result
}

//Delete 批量删除
func (i *IcardModel) Delete(condition Condition) int {
	i.Model.Where(condition.Where)
	result, _ := i.Model.Delete()
	return result
}

//GetAll GetAll
func (i *IcardModel) GetAll(condition Condition, fields ...string) (data []map[string]interface{}) {
	if len(fields) > 0 {
		i.Model.Field(fields)
	}
	if len(condition.Order) > 0 {
		i.Model.OrderBy(condition.Order)
	}
	if condition.Limit > 0 {
		i.Model.Limit(condition.Offset, condition.Limit)
	}
	i.Model.Where(condition.Where)
	return i.Model.Select()
}

//GetCount GetCount
func (i *IcardModel) GetCount(condition Condition, fields ...string) int {
	if len(fields) > 0 {
		i.Model.Field(fields)
	}
	if len(condition.Order) > 0 {
		i.Model.OrderBy(condition.Order)
	}
	i.Model.Where(condition.Where)
	i.Model.Limit(condition.Offset, condition.Limit)
	return i.Model.Count()
}

//GetPaginationData GetPaginationData
func (i *IcardModel) GetPaginationData(condition Condition, fields ...string) (data []map[string]interface{}) {
	if len(fields) > 0 {
		i.Model.Field(fields)
	}
	if len(condition.Order) > 0 {
		i.Model.OrderBy(condition.Order)
	}
	i.Model.Where(condition.Where)
	i.Model.Limit(condition.Offset, condition.Limit)
	return i.Model.Select()
}

//GetRcardsByRcardIds GetRcardsByRcardIds
func (i *IcardModel) GetRcardsByIcardIds(rcardIds []int, fields ...string) []map[string]interface{} {
	if len(rcardIds) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		i.Model.Field(fields)
	}
	return i.Model.Where(map[string]interface{}{
		i.Field.F_icard_id: []interface{}{"IN", rcardIds},
	}).OrderBy(i.Field.F_icard_id + " DESC ").Select()
}

//CreateOrUpdate CreateOrUpdate
func (i *IcardModel) CreateOrUpdate(data map[string]interface{}) (result int) {
	if pk := data[i.PK]; pk != nil {
		result = i.Update(data, pk)
		result = pk.(int)
	} else {
		result = i.Insert(data)
	}
	return
}

//UpdateAll UpdateAll
func (i *IcardModel) UpdateAll(data map[string]interface{}, condition Condition) int {
	i.Model.Where(condition.Where)
	i.Model.Data(data)
	res, _ := i.Model.Update()
	return res
}

//Update update
func (i *IcardModel) Update(data map[string]interface{}, pk interface{}) int {
	result, _ := i.Model.Data(data).Where([]base.WhereItem{
		{i.PK, pk},
	}).Update()
	return result
}

//增加限时限次卡的销量
func (i *IcardModel) IncrSalesByIcardID(icardId int, step int) (err error) {
	if icardId == 0 || step <= 0 {
		return
	}
	_, err = i.Model.Where(map[string]interface{}{
		i.Field.F_icard_id: icardId,
	}).Data(map[string]interface{}{
		i.Field.F_sales: []interface{}{"inc", step},
	}).Update()

	return
}

//根据充值卡id获取单条充值卡数据
func (i *IcardModel) GetByIcardId(icardId int, fields ...[]string) map[string]interface{} {
	if icardId <= 0 {
		return map[string]interface{}{}
	}
	if len(fields) > 0 {
		i.Model.Field(fields[0])
	}
	return i.Model.Where(map[string]interface{}{
		i.Field.F_icard_id: icardId,
	}).Find()
}

func (i *IcardModel) Select(where map[string]interface{}, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		i.Model.Field(fields)
	}
	return i.Model.Where(where).Select()
}

func (i *IcardModel) GetNumByBusID(busId int) (count int) {
	if busId <= 0 {
		return
	}
	count = i.Model.Where(map[string]interface{}{
		i.Field.F_bus_id: busId,
	}).Count(i.Field.F_icard_id)
	return count
}

//UpdateSaleShopNum 更新在售门店数量(增加/减少)
func (i *IcardModel) UpdateSaleShopNum(cardIDs []int, decOrInc string /*可用数值为:dec和inc分别代表增减操作*/) bool {
	if len(cardIDs) == 0 || len(decOrInc) == 0 {
		return false
	}
	_, err := i.Model.Where(map[string]interface{}{
		i.Field.F_icard_id: []interface{}{"IN", cardIDs},
	}).Data(map[string]interface{}{
		i.Field.F_sale_shop_num: []interface{}{decOrInc, addRemoveSaleShopNum},
	}).Update()
	if err != nil {
		return false
	}
	return true
}
