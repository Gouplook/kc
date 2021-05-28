//HNCardShopModel
//2020-04-20 13:41:11

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	"github.com/wendal/errors"
)

//表结构体
type HNCardShopModel struct {
	Model *base.Model
	Field HNCardShopModelField
}

//表字段
type HNCardShopModelField struct{
	T_table	string	`default:"hncard_shop"`
	F_id	string	`default:"id"`
	F_hncard_id	string	`default:"hncard_id"`
	F_bus_id	string	`default:"bus_id"`
	F_shop_id	string	`default:"shop_id"`
}

//初始化
func (n *HNCardShopModel) Init(ormer ...orm.Ormer) *HNCardShopModel {
	functions.ReflectModel(&n.Field)
	n.Model = base.NewModel(n.Field.T_table, ormer...)
	return n
}

//新增数据
func (n *HNCardShopModel) Insert(data map[string]interface{}) (hNCardShopID int, err error){
	if result, insertErr := n.Model.Data(data).Insert() ; insertErr != nil || result == 0 {
		err = errors.New("insert failed")
	}else{
		hNCardShopID = result
	}

	return
}

//批量添加
func (n *HNCardShopModel) InsertAll( data []map[string]interface{}) (err error)  {
	if len(data) == 0{
		return
	}
	_, err = n.Model.InsertAll( data )
	return
}

//根据ids批量修改
func (n *HNCardShopModel) UpdateByIDs ( ids []int, data map[string]interface{} ) (err error) {
	if len(ids) == 0 || len(data) == 0{
		return
	}
	_, err = n.Model.Where(map[string]interface{}{
		n.Field.F_id: []interface{}{"IN", ids},
	}).Data(data).Update()
	return
}

//根据限时限次卡ids 获取所有可适用门店记录
func (n *HNCardShopModel) GetByHNCardIDs( hNCardID []int ) (data []map[string]interface{})  {
	if len(hNCardID) <= 0{
		return
	}
	data = n.Model.Where(map[string]interface{}{
		n.Field.F_hncard_id: []interface{}{"IN", hNCardID},
	}).Select()

	return
}

//根据门店id获取门店可添加的限时限次卡列表
func (n *HNCardShopModel) GetPageByShopID( busID, shopID, start, limit int ) []map[string]interface{} {
	if busID <= 0 || shopID <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}

	return  n.Model.Where(map[string]interface{}{
		n.Field.F_bus_id:  busID,
		n.Field.F_shop_id: []interface{}{"IN", []int{0, shopID}},
	}).OrderBy(n.Field.F_id+" DESC ").Limit( start, limit ).Select()
}
//根据限时限次卡id 和BusId 获取所有可适用门店记录
func (n * HNCardShopModel)GetByHNCardIdAndBusId(hncardId int ,busId int )[]map[string]interface{}{
	if hncardId <= 0 {
		return []map[string]interface{}{}
	}
	return n.Model.Where([]base.WhereItem{
		{n.Field.F_hncard_id,hncardId},
		{n.Field.F_bus_id,busId},
	}).Select()
}


//根据hNCardIDs和shopID获取适用门店的限时限次卡数据
func (n *HNCardShopModel) GetByShopIDAndHNCardIDs( busID, shopID int, hNCardIDs []int  ) []map[string]interface{} {
	if busID <= 0 || shopID <= 0 || len(hNCardIDs) == 0 {
		return []map[string]interface{}{}
	}
	return  n.Model.Where(map[string]interface{}{
		n.Field.F_bus_id:  busID,
		n.Field.F_shop_id: []interface{}{"IN", []int{0, shopID}},
		n.Field.F_hncard_id:   []interface{}{"IN", hNCardIDs},
	}).Select()

}

//根据门店id获取门店可添加的限时限次卡总数量
func (n *HNCardShopModel) GetNumByShopID( busID, shopID int ) int {
	if busID <= 0 || shopID <= 0 {
		return 0
	}
	return  n.Model.Where(map[string]interface{}{
		n.Field.F_bus_id:  busID,
		n.Field.F_shop_id: []interface{}{"IN", []int{0, shopID}},
	}).Count(n.Field.F_id)
}

//批量删除
func (n *HNCardShopModel) DelByHNCardIDs(hNCardIDs []int) (err error) {
	if len(hNCardIDs) == 0{
		return
	}
	_, err = n.Model.Where(map[string]interface{}{
		n.Field.F_hncard_id: []interface{}{"IN", hNCardIDs},
	}).Delete()

	return 
}