//NCardModel
//2020-04-16 11:27:02

package models

import (
	"errors"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"git.900sui.cn/kc/rpcCards/common/tools"
)

//表结构体
type NCardModel struct {
	Model *base.Model
	Field NCardModelField
	tools.CardStatus
}

//表字段
type NCardModelField struct {
	T_table                 string `default:"ncard"`
	F_ncard_id              string `default:"ncard_id"`
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
	F_under_time            string `default:"under_time"`
	F_img_id                string `default:"img_id"`
	F_sales                 string `default:"sales"`
	F_sale_shop_num         string `default:"sale_shop_num"`
	F_ctime                 string `default:"ctime"`
	F_validcount            string `default:"validcount"`
	F_is_del                string `default:"is_del"`
	F_del_time              string `default:"del_time"`
}

//初始化
func (n *NCardModel) Init(ormer ...orm.Ormer) *NCardModel {
	functions.ReflectModel(&n.Field)
	n.Model = base.NewModel(n.Field.T_table, ormer...)
	return n
}

//新增数据
func (n *NCardModel) Insert(data map[string]interface{}) (nCardID int, err error) {
	if result, insertErr := n.Model.Data(data).Insert(); insertErr != nil || result == 0 {
		err = errors.New("insert failed")
	} else {
		nCardID = result
	}
	return
}

func (n *NCardModel) Select(where map[string]interface{}, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		n.Model.Field(fields)
	}
	where[n.Field.F_is_del] = n.NotDelStatus()
	return n.Model.Where(where).Select()
}

//根据NCardID获取单条ncard数据
func (n *NCardModel) GetByNCardID(NCardID int, fields ...string) (dataArray map[string]interface{}) {
	if NCardID == 0 {
		return
	}
	if len(fields) > 0 {
		n.Model.Field(fields)
	}
	dataArray = n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id: NCardID,
		n.Field.F_is_del:   n.NotDelStatus(),
	}).Find()
	return
}

//批量修改
func (n *NCardModel) UpdateByNCardIDs(nCardIDs []int, data map[string]interface{}) (err error) {
	if len(nCardIDs) == 0 || len(data) == 0 {
		return
	}
	_, err = n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id: []interface{}{"IN", nCardIDs},
		n.Field.F_is_del:   n.NotDelStatus(),
	}).Data(data).Update()

	return
}

//根据NCardID修改数据
func (n *NCardModel) UpdateByNCardID(NCardID int, data map[string]interface{}) (err error) {
	if NCardID == 0 || len(data) == 0 {
		return
	}
	_, err = n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id: NCardID,
		n.Field.F_is_del:   n.NotDelStatus(),
	}).Data(data).Update()

	return
}

//增加限次卡的销量
func (n *NCardModel) IncrSalesByNCardID(NCardID int, step int) (err error) {
	if NCardID == 0 || step <= 0 {
		return
	}
	_, err = n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id: NCardID,
		n.Field.F_is_del:   n.NotDelStatus(),
	}).Data(map[string]interface{}{
		n.Field.F_sales: []interface{}{"inc", step},
	}).Update()

	return
}

//获取企业的限次卡列表
func (n *NCardModel) GetPageByBusID(busId int, start, limit int, isGround ...int) (data []map[string]interface{}) {
	where := map[string]interface{}{
		n.Field.F_bus_id: busId,
		n.Field.F_is_del: n.NotDelStatus(),
	}
	if len(isGround) > 0 {
		where[n.Field.F_is_ground] = isGround[0]
	}
	data = n.Model.Where(where).OrderBy(n.Field.F_ncard_id+" DESC ").Limit(start, limit).Select()
	return
}

func (n *NCardModel) SelectNCardsByWherePage(where []base.WhereItem, start, limit int, fields ...string) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	if len(fields) > 0 {
		n.Model.Field(fields)
	}
	if limit > 0 {
		n.Model.Limit(start, limit)
	}
	return n.Model.Where(where).OrderBy(n.Field.F_ncard_id + " DESC ").Select()
}

func (h *NCardModel) GetNumByWhere(where []base.WhereItem) int {
	return h.Model.Where(where).Count(h.Field.F_ncard_id)
}

//获取商家的限次卡数量
func (n *NCardModel) GetNumByBusID(busId int, isGround ...int) (count int) {
	if busId <= 0 {
		return
	}
	where := map[string]interface{}{
		n.Field.F_bus_id: busId,
		n.Field.F_is_del: n.NotDelStatus(),
	}
	if len(isGround) > 0 {
		where[n.Field.F_is_ground] = isGround[0]
	}
	count = n.Model.Where(where).Count(n.Field.F_ncard_id)

	return
}

//根据nCardIDs批量获取数据
func (n *NCardModel) GetByNCardIDs(nCardIDs []int, fields ...string) []map[string]interface{} {
	if len(nCardIDs) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		n.Model.Field(fields)
	}
	return n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id: []interface{}{"IN", nCardIDs},
		n.Field.F_is_del:   n.NotDelStatus(),
	}).Select()
}

//根据nCardIDs批量获取数据
func (n *NCardModel) GetByNCardIDsAndGround(nCardIDs []int, fields ...string) []map[string]interface{} {
	if len(nCardIDs) == 0 {
		return []map[string]interface{}{}
	}
	if len(fields) > 0 {
		n.Model.Field(fields)
	}
	return n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id:  []interface{}{"IN", nCardIDs},
		n.Field.F_is_ground: 1,
		n.Field.F_is_del:    n.NotDelStatus(),
	}).Select()
}

//UpdateSaleShopNum 更新在售门店数量(增加/减少)
func (n *NCardModel) UpdateSaleShopNum(ncardIDs []int, decOrInc string /*可用数值为:dec和inc分别代表增减操作*/) bool {
	if len(ncardIDs) == 0 || len(decOrInc) == 0 {
		return false
	}
	_, err := n.Model.Where(map[string]interface{}{
		n.Field.F_ncard_id: []interface{}{"IN", ncardIDs},
	}).Data(map[string]interface{}{
		n.Field.F_sale_shop_num: []interface{}{decOrInc, addRemoveSaleShopNum},
	}).Update()
	if err != nil {
		return false
	}
	return true
}
