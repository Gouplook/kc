//CardModel
//2020-04-24 09:26:28

package models

import (
	"time"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"git.900sui.cn/kc/rpcCards/common/tools"
)

//表结构体
type CardModel struct {
	Model  *base.Model
	Field  CardModelField
	Status tools.CardStatus
}

//表字段
type CardModelField struct {
	T_table                 string `default:"card"`
	F_card_id               string `default:"card_id"`
	F_bus_id                string `default:"bus_id"`
	F_bind_id               string `default:"bind_id"`
	F_name                  string `default:"name"`
	F_sort_desc             string `default:"sort_desc"`
	F_real_price            string `default:"real_price"`
	F_price                 string `default:"price"`
	F_service_period        string `default:"service_period"`
	F_is_permanent_validity string `default:"is_permanent_validity"`
	F_has_give_signle       string `default:"has_give_signle"`
	F_is_ground             string `default:"is_ground"`
	F_is_del                string `default:"is_del"`
	F_under_time            string `default:"under_time"`
	F_img_id                string `default:"img_id"`
	F_sales                 string `default:"sales"`
	F_sale_shop_num         string `default:"sale_shop_num"`
	F_ctime                 string `default:"ctime"`
	F_del_time              string `default:"del_time"`
}

//初始化
func (c *CardModel) Init(ormer ...orm.Ormer) *CardModel {
	functions.ReflectModel(&c.Field)
	c.Model = base.NewModel(c.Field.T_table, ormer...)
	return c
}

//新增数据
func (c *CardModel) Insert(data map[string]interface{}) (cardID int, err error) {
	cardID, err = c.Model.Data(data).Insert()
	return
}

func (c *CardModel) Select(where map[string]interface{}, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		c.Model.Field(fields)
	}
	return c.Model.Where(where).Where([]base.WhereItem{
		{c.Field.F_is_del, c.Status.NotDelStatus()},
	}).Select()
}

//根据CardID获取单条card数据
func (c *CardModel) GetByCardID(CardID int, fields ...string) (dataArray map[string]interface{}) {
	if CardID == 0 {
		return
	}
	if len(fields) > 0 {
		c.Model.Field(fields)
	}
	dataArray = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: CardID,
	}).Find()
	return
}

//批量修改
func (c *CardModel) UpdateByCardIDs(cardIDs []int, data map[string]interface{}) (err error) {
	if len(cardIDs) == 0 || len(data) == 0 {
		return
	}
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: []interface{}{"IN", cardIDs},
	}).Data(data).Update()

	return
}

//根据CardID修改数据
func (c *CardModel) UpdateByCardID(CardID int, data map[string]interface{}) (err error) {
	if CardID == 0 || len(data) == 0 {
		return
	}
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: CardID,
	}).Data(data).Update()

	return
}

//增加综合卡的销量
func (c *CardModel) IncrSalesByCardID(CardID int, step int) (err error) {
	if CardID == 0 || step <= 0 {
		return
	}
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: CardID,
	}).Data(map[string]interface{}{
		c.Field.F_sales: []interface{}{"inc", step},
	}).Update()

	return
}

//获取企业的综合卡列表
func (c *CardModel) GetPageByBusID(busId int, start, limit int, isGround ...int) (data []map[string]interface{}) {
	where := map[string]interface{}{
		c.Field.F_bus_id: busId,
	}
	if len(isGround) > 0 {
		where[c.Field.F_is_ground] = isGround[0]
	}
	//获取没有删除的数据
	where[c.Field.F_is_del] = c.Status.NotDelStatus()
	data = c.Model.Where(where).OrderBy(c.Field.F_card_id+" DESC ").Limit(start, limit).Select()
	return
}

func (c *CardModel) SelectCardsByWherePage(where []base.WhereItem, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		c.Model.Field(fields)
	}
	if limit > 0 {
		c.Model.Limit(start, limit)
	}
	return c.Model.Where(where).OrderBy(c.Field.F_card_id + " DESC ").Select()
}

func (c *CardModel) GetNumByWhere(where []base.WhereItem) int {
	return c.Model.Where(where).Count(c.Field.F_card_id)
}

//获取商家的综合卡数量
func (c *CardModel) GetNumByBusID(busId int, isGround ...int) (count int) {
	if busId <= 0 {
		return
	}
	where := map[string]interface{}{
		c.Field.F_bus_id: busId,
	}
	if len(isGround) > 0 {
		where[c.Field.F_is_ground] = isGround[0]
	}
	where[c.Field.F_is_del] = c.Status.NotDelStatus()
	count = c.Model.Where(where).Count(c.Field.F_card_id)

	return
}

//根据cardIDs批量获取数据
func (c *CardModel) GetByCardIDs(cardIDs []int, fields ...string) []map[string]interface{} {
	if len(cardIDs) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		c.Model.Field(fields)
	}
	where := []base.WhereItem{
		{c.Field.F_card_id, []interface{}{"IN", cardIDs}},
		{c.Field.F_is_del, c.Status.NotDelStatus()},
	}
	return c.Model.Where(where).Select()
}

//根据cardIDs批量获取数据
func (c *CardModel) GetByCardIDsAndGround(cardIDs []int, fields ...string) []map[string]interface{} {
	if len(cardIDs) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		c.Model.Field(fields)
	}
	where := []base.WhereItem{
		{c.Field.F_card_id, []interface{}{"IN", cardIDs}},
		{c.Field.F_is_ground, 1},
		{c.Field.F_is_del, c.Status.NotDelStatus()},
	}
	return c.Model.Where(where).Select()
}

//UpdateSaleShopNum 更新在售门店数量(增加/减少)
func (c *CardModel) UpdateSaleShopNum(cardIDs []int, decOrInc string /*可用数值为:dec和inc分别代表增减操作*/) bool {
	if len(cardIDs) == 0 || len(decOrInc) == 0 {
		return false
	}
	_, err := c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: []interface{}{"IN", cardIDs},
	}).Data(map[string]interface{}{
		c.Field.F_sale_shop_num: []interface{}{decOrInc, addRemoveSaleShopNum},
	}).Update()
	if err != nil {
		return false
	}
	return true
}

//软删除综合卡-记录删除时间unix
func (c *CardModel) DelByCardId(cardIds []int, busId int) (b bool, err error) {

	//调用模型
	if _, err := c.Model.Where([]base.WhereItem{
		{c.Field.F_bus_id, busId},
		{c.Field.F_card_id, []interface{}{"IN", cardIds}},
	}).Data(map[string]interface{}{
		c.Field.F_is_del:   c.Status.DelStatus(),
		c.Field.F_del_time: time.Now().Unix(),
	}).Update(); err != nil {
		return false, err
	}
	return true, nil
}
