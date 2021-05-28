//CardShopModel
//2020-04-24 09:26:28

package models

import(
	"errors"
	"git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	)

//表结构体
type CardShopModel struct {
	Model *base.Model
	Field CardShopModelField
}

//表字段
type CardShopModelField struct{
	T_table	string	`default:"card_shop"`
	F_id	string	`default:"id"`
	F_card_id	string	`default:"card_id"`
	F_bus_id	string	`default:"bus_id"`
	F_shop_id	string	`default:"shop_id"`
}

//初始化
func (c *CardShopModel) Init(ormer ...orm.Ormer) *CardShopModel{
	functions.ReflectModel(&c.Field)
	c.Model = base.NewModel(c.Field.T_table, ormer...)
	return c
}

//新增数据
func (c *CardShopModel) Insert(data map[string]interface{}) (cardShopID int, err error){
	if result, insertErr := c.Model.Data(data).Insert() ; insertErr != nil || result == 0 {
		err = errors.New("insert failed")
	}else{
		cardShopID = result
	}

	return
}

//批量添加
func (c *CardShopModel) InsertAll( data []map[string]interface{}) (err error)  {
	if len(data) == 0{
		return
	}
	_, err = c.Model.InsertAll( data )
	return
}

//根据ids批量修改
func (c *CardShopModel) UpdateByIDs ( ids []int, data map[string]interface{} ) (err error) {
	if len(ids) == 0 || len(data) == 0{
		return
	}
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_id: []interface{}{"IN", ids},
	}).Data(data).Update()
	return
}

//根据限时综合卡ids 获取所有可适用门店记录
func (c *CardShopModel) GetByCardIDs( cardID []int ) (data []map[string]interface{})  {
	if len(cardID) <= 0{
		return
	}
	data = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: []interface{}{"IN", cardID},
	}).Select()

	return
}

//根据综合卡id 获取所有可适用门店记录
func (c *CardShopModel) GetByCardId(cardId int) []map[string]interface{} {
	if cardId <= 0 {
		return []map[string]interface{}{}
	}
	return c.Model.Where([]base.WhereItem{{c.Field.F_card_id,cardId}}).Select()
}
//根据综合卡id 和BusId 获取所有可适用门店记录
func (c *CardShopModel) GetByCardIdAndBusId(cardId int, busId int) []map[string]interface{} {
	if cardId <= 0 {
		return []map[string]interface{}{}
	}
	return c.Model.Where([]base.WhereItem{
		{c.Field.F_card_id,cardId},
		{c.Field.F_bus_id,busId},
	}).Select()
}

//根据门店id获取门店可添加的限时综合卡列表
func (c *CardShopModel) GetPageByShopID( busID, shopID, start, limit int ) []map[string]interface{} {
	if busID <= 0 || shopID <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}

	return  c.Model.Where(map[string]interface{}{
		c.Field.F_bus_id:  busID,
		c.Field.F_shop_id: []interface{}{"IN", []int{0, shopID}},
	}).OrderBy(c.Field.F_id+" DESC ").Limit( start, limit ).Select()
}

//根据cardIDs和shopID获取适用门店的限时综合卡数据
func (c *CardShopModel) GetByShopIDAndCardIDs( busID, shopID int, cardIDs []int  ) []map[string]interface{} {
	if busID <= 0 || shopID <= 0 || len(cardIDs) == 0 {
		return []map[string]interface{}{}
	}
	return  c.Model.Where(map[string]interface{}{
		c.Field.F_bus_id:  busID,
		c.Field.F_shop_id: []interface{}{"IN", []int{0, shopID}},
		c.Field.F_card_id: []interface{}{"IN", cardIDs},
	}).Select()

}

//根据门店id获取门店可添加的限时综合卡总数量
func (c *CardShopModel) GetNumByShopID( busID, shopID int ) int {
	if busID <= 0 || shopID <= 0 {
		return 0
	}
	return  c.Model.Where(map[string]interface{}{
		c.Field.F_bus_id:  busID,
		c.Field.F_shop_id: []interface{}{"IN", []int{0, shopID}},
	}).Count(c.Field.F_id)
}

//批量删除
func (c *CardShopModel) DelByCardIDs(cardIDs []int) (err error) {
	if len(cardIDs) == 0{
		return
	}
	_, err = c.Model.Where(map[string]interface{}{
		c.Field.F_card_id: []interface{}{"IN", cardIDs},
	}).Delete()

	return
}