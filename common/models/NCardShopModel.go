//NCardShopModel
//2020-04-20 13:41:11

package models

import(
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/base/common/models/base"
    "git.900sui.cn/kc/kcgin/orm"
	"github.com/wendal/errors"
)

//表结构体
type NCardShopModel struct {
	Model *base.Model
	Field NCardShopModelField
}

//表字段
type NCardShopModelField struct{
	T_table	string	`default:"ncard_shop"`
	F_id	string	`default:"id"`
	F_ncard_id	string	`default:"ncard_id"`
	F_bus_id	string	`default:"bus_id"`
	F_shop_id	string	`default:"shop_id"`
}

//初始化
func (n *NCardShopModel) Init(ormer ...orm.Ormer) *NCardShopModel {
	functions.ReflectModel(&n.Field)
	n.Model = base.NewModel(n.Field.T_table, ormer...)
	return n
}

//新增数据
func (n *NCardShopModel) Insert(data map[string]interface{}) (nCardShopID int, err error){
	if result, insertErr := n.Model.Data(data).Insert() ; insertErr != nil || result == 0 {
		err = errors.New("insert failed")
	}else{
		nCardShopID = result
	}

	return
}

//批量添加
func (n *NCardShopModel) InsertAll( data []map[string]interface{}) (err error)  {
	if len(data) == 0{
		return
	}
	_, err = n.Model.InsertAll( data )
	return
}

//根据ids批量修改
func (n *NCardShopModel) UpdateByIDs ( ids []int, data map[string]interface{} ) (err error) {
	if len(ids) == 0 || len(data) == 0{
		return
	}
	_, err = n.Model.Where(map[string]interface{}{
		n.Field.F_id: []interface{}{"IN", ids},
	}).Data(data).Update()
	return
}

//根据限次卡ids 获取所有可适用门店记录
func (n *NCardShopModel) GetByNCardIDs( nCardID []int ) (data []map[string]interface{})  {
	if len(nCardID) <= 0{
		return
	}
	data = n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id: []interface{}{"IN", nCardID},
	}).Select()

	return
}
//根据限次卡id and busId 获取所有可适用门店记录
func (n *NCardShopModel)GetByNcardIdAndBusId(ncardId int ,busId int )[]map[string]interface{}{
	if ncardId <= 0 {
		return []map[string]interface{}{}
	}
	return n.Model.Where([]base.WhereItem{
		{n.Field.F_ncard_id,ncardId},
		{n.Field.F_bus_id,busId},
	}).Select()
}

//根据门店id获取门店可添加的限次卡列表
func (n *NCardShopModel) GetPageByShopID( busID, shopID, start, limit int ) []map[string]interface{} {
	if busID <= 0 || shopID <= 0 || start < 0 || limit <= 0 {
		return []map[string]interface{}{}
	}

	return  n.Model.Where(map[string]interface{}{
		n.Field.F_bus_id:  busID,
		n.Field.F_shop_id: []interface{}{"IN", []int{0, shopID}},
	}).OrderBy(n.Field.F_id+" DESC ").Limit( start, limit ).Select()
}

//根据nCardIDs和shopID获取适用门店的限次卡数据
func (n *NCardShopModel) GetByShopIDAndNCardIDs( busID, shopID int, nCardIDs []int  ) []map[string]interface{} {
	if busID <= 0 || shopID <= 0 || len(nCardIDs) == 0 {
		return []map[string]interface{}{}
	}
	return  n.Model.Where(map[string]interface{}{
		n.Field.F_bus_id:  busID,
		n.Field.F_shop_id: []interface{}{"IN", []int{0, shopID}},
		n.Field.F_ncard_id:   []interface{}{"IN", nCardIDs},
	}).Select()

}

//根据门店id获取门店可添加的限次卡总数量
func (n *NCardShopModel) GetNumByShopID( busID, shopID int ) int {
	if busID <= 0 || shopID <= 0 {
		return 0
	}
	return  n.Model.Where(map[string]interface{}{
		n.Field.F_bus_id:  busID,
		n.Field.F_shop_id: []interface{}{"IN", []int{0, shopID}},
	}).Count(n.Field.F_id)
}

//批量删除
func (n *NCardShopModel) DelByNCardIDs(nCardIDs []int) (err error) {
	if len(nCardIDs) == 0{
		return
	}
	_, err = n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id: []interface{}{"IN", nCardIDs},
	}).Delete()

	return 
}