package logics

import (
	"git.900sui.cn/kc/rpcCards/common/models"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"strconv"
)

type ApplyAndGiveSingleNum struct {
	IsAllSingle    bool //适用于全部单项目
	ApplySingleNum int  //适用单项目的个数
	GiveSingleNum  int  //赠送单项目的个数
}

//获取不同卡项-适用单项目的个数和赠送单项目的个数
func GetApplyAndGiveSingleNum(itemCardIds []int, itemType int) (aagsmMap map[int]ApplyAndGiveSingleNum) {
	aagsmMap = make(map[int]ApplyAndGiveSingleNum) //key:卡项id
	switch itemType {
	case cards.ITEM_TYPE_sm: //套餐
		ssm := new(models.SmSingleModel).Init()
		ssmMaps := ssm.GetBySmids(itemCardIds)
		for _, v := range ssmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_sm_id].(string))
			value, ok := aagsmMap[smId]
			if v[ssm.Field.F_single_id].(string) == "0" { //适用于全部项目
				value.IsAllSingle = true
				aagsmMap[smId] = value
				continue
			} else {
				if !ok {
					value.ApplySingleNum = 1
				} else {
					value.ApplySingleNum += 1
				}
				aagsmMap[smId] = value
			}
		}
		sgm := new(models.SmGiveModel).Init()
		sgmMaps := sgm.GetBySmids(itemCardIds)
		for _, v := range sgmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_sm_id].(string))
			value, ok := aagsmMap[smId]
			if !ok {
				value.GiveSingleNum = 1
			} else {
				value.GiveSingleNum += 1
			}
			aagsmMap[smId] = value
		}
	case cards.ITEM_TYPE_card: //综合卡
		ssm := new(models.CardSingleModel).Init()
		ssmMaps := ssm.GetByCardIds(itemCardIds)
		for _, v := range ssmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_card_id].(string))
			value, ok := aagsmMap[smId]
			if v[ssm.Field.F_single_id].(string) == "0" { //适用于全部项目
				value.IsAllSingle = true
				aagsmMap[smId] = value
				continue
			} else {
				if !ok {
					value.ApplySingleNum = 1
				} else {
					value.ApplySingleNum += 1
				}
				aagsmMap[smId] = value
			}
		}
		sgm := new(models.CardGiveModel).Init()
		sgmMaps := sgm.GetByCardIds(itemCardIds)
		for _, v := range sgmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_card_id].(string))
			value, ok := aagsmMap[smId]
			if !ok {
				value.GiveSingleNum = 1
			} else {
				value.GiveSingleNum += 1
			}
			aagsmMap[smId] = value
		}
	case cards.ITEM_TYPE_hcard: //限时卡
		ssm := new(models.HcardSingleModel).Init()
		ssmMaps := ssm.GetByHcardIds(itemCardIds)
		for _, v := range ssmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_hcard_id].(string))
			value, ok := aagsmMap[smId]
			if v[ssm.Field.F_single_id].(string) == "0" { //适用于全部项目
				value.IsAllSingle = true
				aagsmMap[smId] = value
				continue
			} else {
				if !ok {
					value.ApplySingleNum = 1
				} else {
					value.ApplySingleNum += 1
				}
				aagsmMap[smId] = value
			}
		}
		sgm := new(models.HcardGiveModel).Init()
		sgmMaps := sgm.GetByHcardIds(itemCardIds)
		for _, v := range sgmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_hcard_id].(string))
			value, ok := aagsmMap[smId]
			if !ok {
				value.GiveSingleNum = 1
			} else {
				value.GiveSingleNum += 1
			}
			aagsmMap[smId] = value
		}
	case cards.ITEM_TYPE_ncard: //限次卡
		ssm := new(models.NCardSingleModel).Init()
		ssmMaps := ssm.GetByNCardIds(itemCardIds)
		for _, v := range ssmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_ncard_id].(string))
			value, ok := aagsmMap[smId]
			if v[ssm.Field.F_single_id].(string) == "0" { //适用于全部项目
				value.IsAllSingle = true
				aagsmMap[smId] = value
				continue
			} else {
				if !ok {
					value.ApplySingleNum = 1
				} else {
					value.ApplySingleNum += 1
				}
				aagsmMap[smId] = value
			}
		}
		sgm := new(models.NCardGiveModel).Init()
		sgmMaps := sgm.GetByNCardIds(itemCardIds)
		for _, v := range sgmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_ncard_id].(string))
			value, ok := aagsmMap[smId]
			if !ok {
				value.GiveSingleNum = 1
			} else {
				value.GiveSingleNum += 1
			}
			aagsmMap[smId] = value
		}
	case cards.ITEM_TYPE_hncard: //限时限次卡
		ssm := new(models.HNCardSingleModel).Init()
		ssmMaps := ssm.GetByHNCardIds(itemCardIds)
		for _, v := range ssmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_hncard_id].(string))
			value, ok := aagsmMap[smId]
			if v[ssm.Field.F_single_id].(string) == "0" { //适用于全部项目
				value.IsAllSingle = true
				aagsmMap[smId] = value
				continue
			} else {
				if !ok {
					value.ApplySingleNum = 1
				} else {
					value.ApplySingleNum += 1
				}
				aagsmMap[smId] = value
			}
		}
		sgm := new(models.HNCardGiveModel).Init()
		sgmMaps := sgm.GetByHNCardIds(itemCardIds)
		for _, v := range sgmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_hncard_id].(string))
			value, ok := aagsmMap[smId]
			if !ok {
				value.GiveSingleNum = 1
			} else {
				value.GiveSingleNum += 1
			}
			aagsmMap[smId] = value
		}
	case cards.ITEM_TYPE_rcard: //充值卡
		ssm := new(models.RcardSingleModel).Init()
		ssmMaps := ssm.GetByRcardIds(itemCardIds)
		for _, v := range ssmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_rcard_id].(string))
			value, ok := aagsmMap[smId]
			if v[ssm.Field.F_single_id].(string) == "0" { //适用于全部项目
				value.IsAllSingle = true
				aagsmMap[smId] = value
				continue
			} else {
				if !ok {
					value.ApplySingleNum = 1
				} else {
					value.ApplySingleNum += 1
				}
				aagsmMap[smId] = value
			}
		}
		sgm := new(models.RcardGiveModel).Init()
		sgmMaps := sgm.GetByRcardIds(itemCardIds)
		for _, v := range sgmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_rcard_id].(string))
			value, ok := aagsmMap[smId]
			if !ok {
				value.GiveSingleNum = 1
			} else {
				value.GiveSingleNum += 1
			}
			aagsmMap[smId] = value
		}
	case cards.ITEM_TYPE_icard: //身份卡
		ssm := new(models.IcardSingleModel).Init()
		ssmMaps := ssm.GetByIcardIds(itemCardIds)
		for _, v := range ssmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_icard_id].(string))
			value, ok := aagsmMap[smId]
			if v[ssm.Field.F_single_id].(string) == "0" { //适用于全部项目
				value.IsAllSingle = true
				aagsmMap[smId] = value
				continue
			} else {
				if !ok {
					value.ApplySingleNum = 1
				} else {
					value.ApplySingleNum += 1
				}
				aagsmMap[smId] = value
			}
		}
		sgm := new(models.IcardGiveModel).Init()
		sgmMaps := sgm.GetAll(models.Condition{
			Where:  map[string]interface{}{sgm.Field.F_icard_id: []interface{}{"IN", itemCardIds}},
			Offset: 0,
			Limit:  0,
			Order:  "",
		})
		for _, v := range sgmMaps {
			smId, _ := strconv.Atoi(v[ssm.Field.F_icard_id].(string))
			value, ok := aagsmMap[smId]
			if !ok {
				value.GiveSingleNum = 1
			} else {
				value.GiveSingleNum += 1
			}
			aagsmMap[smId] = value
		}
	}
	return
}
