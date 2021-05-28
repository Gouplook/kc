//充值卡业务处理
//@author yangzhiwu<578154898@qq.com>
//@date 2020/10/21 14:55
package logics

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/mapstructure"
	"git.900sui.cn/kc/redis"
	"git.900sui.cn/kc/rpcCards/common/models"
	"git.900sui.cn/kc/rpcCards/common/tools"
	"git.900sui.cn/kc/rpcCards/constkey"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/client/bus"
	bus2 "git.900sui.cn/kc/rpcinterface/interface/bus"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	redis2 "github.com/gomodule/redigo/redis"
	"github.com/shopspring/decimal"
)

type RcardLogic struct {
}

//添加RcardExt描述
func (r *RcardLogic) AddRcardExt(mRcard *models.RcardModel, rcardId int, notes []cards.CardNote) (err error) {
	if len(notes) > constkey.CardNotesMaxNum {
		return toolLib.CreateKcErr(_const.NOTES_NUM_MAX_10)
	}
	for _, note := range notes {
		if functions.Mb4Strlen(note.Notes) > constkey.CardNotesSimpleMaxLength {
			return toolLib.CreateKcErr(_const.NOTES_LEN_MAX_30)
		}
	}
	notesStr, _ := json.Marshal(notes)
	nem := new(models.RcardExtModel).Init(mRcard.Model.GetOrmer())
	if _, err = nem.InsertExt(map[string]interface{}{nem.Field.F_rcard_id: rcardId, nem.Field.F_notes: string(notesStr)}); err != nil {
		err = toolLib.CreateKcErr(_const.DB_ERR)
	}
	return
}

//编辑RcardExt描述
func (r *RcardLogic) EditRcardExt(rcardId int, notes []cards.CardNote) (err error) {
	if len(notes) > constkey.CardNotesMaxNum {
		return toolLib.CreateKcErr(_const.NOTES_NUM_MAX_10)
	}
	for _, note := range notes {
		if functions.Mb4Strlen(note.Notes) > constkey.CardNotesSimpleMaxLength {
			return toolLib.CreateKcErr(_const.NOTES_LEN_MAX_30)
		}
	}
	notesStr, _ := json.Marshal(notes)
	nem := new(models.RcardExtModel).Init()
	rcardExMap := nem.Find(map[string]interface{}{nem.Field.F_rcard_id: rcardId})
	if len(rcardExMap) > 0 {
		if _, updateErr := nem.Update(map[string]interface{}{nem.Field.F_rcard_id: rcardId},
			map[string]interface{}{nem.Field.F_notes: string(notesStr)}); updateErr != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	} else {
		if _, err = nem.InsertExt(map[string]interface{}{nem.Field.F_rcard_id: rcardId, nem.Field.F_notes: string(notesStr)}); err != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	return
}

//获取RcardExt描述
func (r *RcardLogic) GetRcardExt(rcardId int, notes *[]cards.CardNote) {
	nem := new(models.RcardExtModel).Init()
	dataMap := nem.Find(map[string]interface{}{nem.Field.F_rcard_id: rcardId})
	if len(dataMap) > 0 {
		note := dataMap[nem.Field.F_notes].(string)
		json.Unmarshal([]byte(note), notes)
	}
	return
}

//添加充值卡
func (r *RcardLogic) AddRcard(ctx context.Context, busId int, args *cards.ArgsAddRcard) (rcardId int, err error) {

	//验证卡项数据
	err = r.checkRcardData(ctx, busId, args.RcardBase, args.IsAllSingle, args.IsAllProduct, args.IncludeSingles, args.GiveSingles, args.IncludeProducts)
	if err != nil {
		return
	}

	//验证图片
	imgId, err := checkImg(ctx, args.ImgHash)
	if err != nil {
		return
	}
	var hasGive uint8 = 0
	if len(args.GiveSingles) > 0 {
		hasGive = 1
	}

	//添加基本信息
	mRcard := new(models.RcardModel).Init()
	discount := args.Discount
	if args.DiscountType == cards.DISCOUNT_TYPE_price {
		discount, _ = decimal.NewFromFloat(args.RealPrice).Div(decimal.NewFromFloat(args.Price + args.RealPrice)).Truncate(2).Float64()
		discount = discount * 10 // 折扣率转换成折扣 100
		args.Price += args.RealPrice
	} else {
		args.Price = args.RealPrice
	}
	if discount < cards.DICOUNTMIN {
		err = toolLib.CreateKcErr(_const.DISCOUNT_ERR)
		return
	}
	// 购买或者充值须100起
	if tools.RunMode == "prod" {
		if args.RealPrice < cards.BUY_CRARD_MIN_AMOUNT {
			err = toolLib.CreateKcErr(_const.BUY_CRARD_MIN_AMOUNT_ERR)
			return
		}
	}

	// 卡项有效期 1-是；2-否；如果永久有效则service_period字段可忽略，否者必填
	ispermanentvalidity := cards.IS_PERMANENT_YES
	if args.ServicePeriod != 0 {
		ispermanentvalidity = cards.IS_PERMANENT_NO
	}
	mRcard.Model.Begin()
	defer func() {
		if err != nil {
			mRcard.Model.RollBack()
			return
		}
	}()
	rcardId = mRcard.Insert(map[string]interface{}{
		mRcard.Field.F_bus_id:                busId,
		mRcard.Field.F_bind_id:               getBusMainBindId(ctx, busId),
		mRcard.Field.F_name:                  args.Name,
		mRcard.Field.F_sort_desc:             args.SortDesc,
		mRcard.Field.F_real_price:            args.RealPrice,
		mRcard.Field.F_price:                 args.Price,
		mRcard.Field.F_discount_type:         args.DiscountType,
		mRcard.Field.F_discount:              discount,
		mRcard.Field.F_is_have_discount:      args.IsHaveDiscount,
		mRcard.Field.F_service_period:        args.ServicePeriod,
		mRcard.Field.F_has_give_signle:       hasGive,
		mRcard.Field.F_is_ground:             cards.SINGLE_IS_GROUND_yes,
		mRcard.Field.F_img_id:                imgId,
		mRcard.Field.F_ctime:                 time.Now().Local().Unix(),
		mRcard.Field.F_is_permanent_validity: ispermanentvalidity,
	})
	if rcardId == 0 {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	// 添加充值卡规则
	if len(args.RcardRule) > 0 {
		mRcardmode := new(models.RcardRuleModel).Init(mRcard.Model.GetOrmer())
		rcardRuleData := []map[string]interface{}{}
		for _, rule := range args.RcardRule {
			if args.DiscountType == cards.DISCOUNT_TYPE_item {
				rule.DonationAmount = 0.0
			}
			rcardRuleData = append(rcardRuleData, map[string]interface{}{
				mRcardmode.Field.F_rcard_id:   rcardId,
				mRcardmode.Field.F_price:      rule.RechargeAmount,
				mRcardmode.Field.F_give_price: rule.DonationAmount,
			})
		}
		if len(rcardRuleData) > 0 {
			if mRcardmode.InsertAll(rcardRuleData) < 0 {
				err = toolLib.CreateKcErr(_const.DB_ERR)
				return
			}
		}
	}
	if args.IsHaveDiscount == cards.IS_HAVE_NO {
		args.SingleDiscount = 10.0 // 表示没有折扣
	}

	//添加单项目
	mRS := new(models.RcardSingleModel).Init(mRcard.Model.GetOrmer())
	if args.IsAllSingle { //全部单项目
		if mRS.Insert(map[string]interface{}{
			mRS.Field.F_rcard_id:  rcardId,
			mRS.Field.F_single_id: 0,
			mRS.Field.F_discount:  args.SingleDiscount,
		}) <= 0 {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	} else {
		if len(args.IncludeSingles) > 0 {
			rcardSingleData := []map[string]interface{}{}
			for _, single := range args.IncludeSingles {
				rcardSingleData = append(rcardSingleData, map[string]interface{}{
					mRS.Field.F_rcard_id:  rcardId,
					mRS.Field.F_single_id: single.SingleID,
					mRS.Field.F_discount:  args.SingleDiscount,
				})
			}
			if len(rcardSingleData) > 0 {
				if mRS.InsertAll(rcardSingleData) < 0 {
					err = toolLib.CreateKcErr(_const.DB_ERR)
					return
				}
			}
		}
	}

	//添加商品
	mRP := new(models.RcardGoodsModel).Init(mRcard.Model.GetOrmer())
	if args.IsAllProduct { //全部商品
		if mRP.Insert(map[string]interface{}{
			mRP.Field.F_rcard_id:   rcardId,
			mRP.Field.F_product_id: 0,
		}) <= 0 {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	} else {
		if len(args.IncludeProducts) > 0 {
			rcardProductData := []map[string]interface{}{}
			for _, product := range args.IncludeProducts {
				rcardProductData = append(rcardProductData, map[string]interface{}{
					mRP.Field.F_rcard_id:   rcardId,
					mRP.Field.F_product_id: product.ProductID,
				})
			}
			if len(rcardProductData) > 0 {
				if _, err = mRP.InsertAll(rcardProductData); err != nil {
					err = toolLib.CreateKcErr(_const.DB_ERR)
					return
				}
			}
		}
	}

	//添加赠送项目
	if len(args.GiveSingles) > 0 {
		mRG := new(models.RcardGiveModel).Init(mRcard.Model.GetOrmer())
		rcardGiveData := []map[string]interface{}{}
		for _, give := range args.GiveSingles {
			rcardGiveData = append(rcardGiveData, map[string]interface{}{
				mRG.Field.F_rcard_id:           rcardId,
				mRG.Field.F_single_id:          give.SingleID,
				mRG.Field.F_num:                give.Num,
				mRG.Field.F_period_of_validity: give.PeriodOfValidity,
			})
		}
		if len(rcardGiveData) > 0 {
			if mRG.InsertAll(rcardGiveData) < 0 {
				err = toolLib.CreateKcErr(_const.DB_ERR)
				return
			}
		}
	}
	//添加赠品描述
	if len(args.GiveSingles) > 0 && len(args.GiveSingleDesc) > 0 {
		rdescm := new(models.RcardGiveDescModel).Init(mRcard.Model.GetOrmer())
		giveSingleDescStr, _ := json.Marshal(args.GiveSingleDesc)
		descData := map[string]interface{}{
			rdescm.Field.F_rcard_id: rcardId,
			rdescm.Field.F_desc:     string(giveSingleDescStr),
		}
		if _, err = rdescm.Model.Data(descData).Insert(); err != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	if err = r.AddRcardExt(mRcard, rcardId, args.Notes); err != nil {
		return
	}
	mRcard.Model.Commit()
	//添加风控统计任务
	new(ItemLogic).AddXCardTask(ctx, rcardId, cards.ITEM_TYPE_rcard)
	return

}

//编辑充值卡
func (r *RcardLogic) EditRcard(ctx context.Context, busId int, args *cards.ArgsEditRcard) (err error) {

	mRcard := new(models.RcardModel).Init()
	mRcardRule := new(models.RcardRuleModel).Init(mRcard.Model.GetOrmer())
	mRcardSingle := new(models.RcardSingleModel).Init(mRcard.Model.GetOrmer())
	mRcardGoods := new(models.RcardGoodsModel).Init(mRcard.Model.GetOrmer())
	mRcardGive := new(models.RcardGiveModel).Init(mRcard.Model.GetOrmer())
	rdesm := new(models.RcardGiveDescModel).Init(mRcard.Model.GetOrmer())
	rcardInfo := mRcard.GetByRcardId(args.RcardId)
	if len(rcardInfo) == 0 {
		err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
		return
	}
	if rcardInfo[mRcard.Field.F_bus_id] != strconv.Itoa(busId) {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}

	args.DiscountType, _ = strconv.Atoi(rcardInfo[mRcard.Field.F_discount_type].(string))
	//验证参数
	err = r.checkRcardData(ctx, busId, args.RcardBase, args.IsAllSingle, args.IsAllProduct, args.IncludeSingles, args.GiveSingles, args.IncludeProducts)
	if err != nil {
		return
	}
	imgId, err := checkImg(ctx, args.ImgHash)
	if err != nil {
		return
	}
	var hasGive uint8 = 0
	if len(args.GiveSingles) > 0 {
		hasGive = 1
	}
	discount := args.Discount
	if args.DiscountType == cards.DISCOUNT_TYPE_price {
		discount, _ = decimal.NewFromFloat(args.RealPrice).Div(decimal.NewFromFloat(args.Price + args.RealPrice)).Truncate(2).Float64()
		discount = discount * 10 // 折扣率转换成折扣
		args.Price += args.RealPrice
	} else {
		args.Price = args.RealPrice
	}
	if discount < cards.DICOUNTMIN {
		err = toolLib.CreateKcErr(_const.DISCOUNT_ERR)
		return
	}
	// 购买或者充值须100起
	if tools.RunMode == "prod" {
		if args.RealPrice < cards.BUY_CRARD_MIN_AMOUNT {
			err = toolLib.CreateKcErr(_const.BUY_CRARD_MIN_AMOUNT_ERR)
			return
		}
	}
	//修改包含的单项目
	includeSingles := mRcardSingle.GetByRcardid(args.RcardId)
	includeSingleIDs := functions.ArrayValue2Array(mRcardSingle.Field.F_single_id, includeSingles)
	//需要新增的项目
	var addIncSingles []map[string]interface{}
	//需要删除的ids
	var delCardSingleIDs []int
	if args.IsHaveDiscount == cards.IS_HAVE_NO {
		args.SingleDiscount = 10.0 // 表示没有折扣
	}
	//获取需要添加的项目
	if args.IsAllSingle { //全部单项目
		addIncSingles = append(addIncSingles, map[string]interface{}{
			mRcardSingle.Field.F_rcard_id:  args.RcardId,
			mRcardSingle.Field.F_single_id: 0,
			mRcardSingle.Field.F_discount:  args.SingleDiscount,
		})
		args.IncludeSingles = []cards.IncSingle{}
	} else { //部分单项目
		for _, single := range args.IncludeSingles {
			if functions.InArray(single.SingleID, includeSingleIDs) {
				continue
			}
			addIncSingles = append(addIncSingles, map[string]interface{}{
				mRcardSingle.Field.F_rcard_id:  args.RcardId,
				mRcardSingle.Field.F_single_id: single.SingleID,
				mRcardSingle.Field.F_discount:  args.SingleDiscount,
			})
		}
	}
	//获取需要删除的id
out1:
	for index, dbSingleID := range includeSingleIDs {
		for _, single := range args.IncludeSingles {
			if single.SingleID == dbSingleID {
				continue out1
			}
		}
		id, _ := strconv.Atoi(includeSingles[index][mRcardSingle.Field.F_id].(string))
		delCardSingleIDs = append(delCardSingleIDs, id)
	}

	//计算商品改动
	includeProducts := mRcardGoods.GetByRcardid(args.RcardId)
	includeProductIDs := functions.ArrayValue2Array(mRcardGoods.Field.F_product_id, includeProducts)
	//需要新增的productIDs
	var addIncProducts []map[string]interface{}
	//需要删除的productIDs
	var delCardProductIDs []int

	//获取需要添加的product
	if args.IsAllProduct { //包含全部
		addIncProducts = append(addIncProducts, map[string]interface{}{
			mRcardGoods.Field.F_rcard_id:   args.RcardId,
			mRcardGoods.Field.F_product_id: 0,
		})
		args.IncludeProducts = []cards.IncProduct{}
	} else { //部分你商品
		for _, includeProduct := range args.IncludeProducts {
			if functions.InArray(includeProduct.ProductID, includeProductIDs) {
				continue
			}
			addIncProducts = append(addIncProducts, map[string]interface{}{
				mRcardGoods.Field.F_rcard_id:   args.RcardId,
				mRcardGoods.Field.F_product_id: includeProduct.ProductID,
			})
		}
	}

	//获取需要删除的product
out2:
	for index, productID := range includeProductIDs {
		for _, includeProduct := range args.IncludeProducts {
			if includeProduct.ProductID == productID {
				continue out2
			}
		}
		id, _ := strconv.Atoi(includeProducts[index][mRcardGoods.Field.F_id].(string))
		delCardProductIDs = append(delCardProductIDs, id)
	}

	// 卡项有效期 1-是；2-否；如果永久有效则service_period字段可忽略，否者必填
	ispermanentvalidity := cards.IS_PERMANENT_YES
	if args.ServicePeriod != 0 {
		ispermanentvalidity = cards.IS_PERMANENT_NO
	}
	mRcard.Model.Begin()
	defer func() {
		if err != nil {
			mRcard.Model.RollBack()
			return
		}
	}()
	//修改主表信息
	rs := mRcard.UpdateByRcardid(args.RcardId, map[string]interface{}{
		mRcard.Field.F_name:                  args.Name,
		mRcard.Field.F_sort_desc:             args.SortDesc,
		mRcard.Field.F_real_price:            args.RealPrice,
		mRcard.Field.F_price:                 args.Price,
		mRcard.Field.F_discount:              discount,
		mRcard.Field.F_is_have_discount:      args.IsHaveDiscount,
		mRcard.Field.F_service_period:        args.ServicePeriod,
		mRcard.Field.F_has_give_signle:       hasGive,
		mRcard.Field.F_img_id:                imgId,
		mRcard.Field.F_is_permanent_validity: ispermanentvalidity,
	})
	if rs == false {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	// 修改充值卡规则
	if len(args.RcardRule) == 0 {
		if !mRcardRule.DelByRcardid(args.RcardId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	} else {
		var incRcardRuleNum int // 原来充值卡规则条数
		var delIdsNum int       //  删除充值卡规则条数
		var addRcardRuleNum int // 新增充值卡规则条数
		// 定义需要新增充值卡规则
		var addRcardRule []map[string]interface{}
		// 定义需要删除的Ids
		var delIds []int
		incRcardRule := mRcardRule.GetByRcardid(args.RcardId)
		incRcardRuleNum = mRcardRule.GetNumByRcardId(args.RcardId, cards.RECHARGE_TYPE_NO)
		for _, rule := range args.RcardRule {
			hasd := 0
			if args.DiscountType == cards.DISCOUNT_TYPE_item {
				rule.DonationAmount = 0.0
			}
			for _, dbRcardRule := range incRcardRule {
				if strconv.FormatFloat(rule.RechargeAmount, 'E', -1, 64) == dbRcardRule[mRcardRule.Field.F_price].(string) &&
					strconv.FormatFloat(rule.DonationAmount, 'E', -1, 64) == dbRcardRule[mRcardRule.Field.F_give_price].(string) {
					hasd = 1
					break
				}
			}
			if hasd == 0 {
				addRcardRule = append(addRcardRule, map[string]interface{}{
					mRcardRule.Field.F_rcard_id:   args.RcardId,
					mRcardRule.Field.F_price:      rule.RechargeAmount,
					mRcardRule.Field.F_give_price: rule.DonationAmount,
				})
				addRcardRuleNum++
			}
		}
		// 计算需要删除的规则
		for _, dbRcardRule := range incRcardRule {
			hasd := 0
			for _, rule := range args.RcardRule {
				if strconv.FormatFloat(rule.RechargeAmount, 'E', -1, 64) == dbRcardRule[mRcardRule.Field.F_price].(string) &&
					strconv.FormatFloat(rule.DonationAmount, 'E', -1, 64) == dbRcardRule[mRcardRule.Field.F_give_price].(string) {
					hasd = 1
					break
				}
			}
			if hasd == 0 {
				id, _ := strconv.Atoi(dbRcardRule[mRcardRule.Field.F_id].(string))
				delIds = append(delIds, id)
				delIdsNum++
			}
		}

		if len(addRcardRule) > 0 {
			if (incRcardRuleNum - delIdsNum + addRcardRuleNum) > 10 {
				err = toolLib.CreateKcErr(_const.RCARD_MAX_NUM_ERR)
				return
			} else {
				if mRcardRule.InsertAll(addRcardRule) <= 0 {
					err = toolLib.CreateKcErr(_const.DB_ERR)
					return
				}
			}

		}
		if len(delIds) > 0 {
			if !mRcardRule.DelByIds(delIds) {
				err = toolLib.CreateKcErr(_const.DB_ERR)
				return
			}
		}
	}

	//修改包含的单项
	//1. 添加需要添加的项目
	if len(addIncSingles) > 0 {
		if mRcardSingle.InsertAll(addIncSingles) <= 0 {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//2. 删除需要删除的项目
	if len(delCardSingleIDs) > 0 {
		if !mRcardSingle.DelByIds(delCardSingleIDs) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//修改包含的商品
	//1. 添加需要添加的商品
	if len(addIncProducts) > 0 {
		if _, err = mRcardGoods.InsertAll(addIncProducts); err != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//2. 删除需要删除的商品
	if len(delCardProductIDs) > 0 {
		if err = mRcardGoods.DelByIds(delCardProductIDs); err != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//赠送项目修改
	gives := mRcardGive.GetByRcardid(args.RcardId)
	if len(args.GiveSingles) == 0 {
		if !mRcardGive.DelByRcardid(args.RcardId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	} else {
		var addGives []map[string]interface{}
		var delGiveIds []int
		//定义需要修改的单项目
		var updateGiveSingles = map[int]int{} //id=>num
		//计算要新增的赠项目
		for _, give := range args.GiveSingles {
			hasd := 0
			for _, dbGive := range gives {
				if strconv.Itoa(give.SingleID) == dbGive[mRcardGive.Field.F_single_id].(string) {
					hasd = 1
				}
				//需要修改次数
				if strconv.Itoa(give.Num) != dbGive[mRcardGive.Field.F_num].(string) {
					dbId, _ := strconv.Atoi(dbGive[mRcardGive.Field.F_id].(string))
					updateGiveSingles[dbId] = give.Num
				}
			}

			if hasd == 0 {
				addGives = append(addGives, map[string]interface{}{
					mRcardGive.Field.F_rcard_id:           args.RcardId,
					mRcardGive.Field.F_single_id:          give.SingleID,
					mRcardGive.Field.F_num:                give.Num,
					mRcardGive.Field.F_period_of_validity: give.PeriodOfValidity,
				})
			}
		}
		//计算要删除的
		for _, dbGive := range gives {
			hasd := 0
			for _, give := range args.GiveSingles {
				if dbGive[mRcardGive.Field.F_single_id].(string) == strconv.Itoa(give.SingleID) {
					hasd = 1
				}
			}
			if hasd == 0 {
				dbId, _ := strconv.Atoi(dbGive[mRcardGive.Field.F_id].(string))
				delGiveIds = append(delGiveIds, dbId)
			}
		}

		if len(addGives) > 0 {
			if mRcardGive.InsertAll(addGives) <= 0 {
				err = toolLib.CreateKcErr(_const.DB_ERR)
				return
			}
		}
		if len(updateGiveSingles) > 0 {
			for id, num := range updateGiveSingles {
				if !mRcardGive.UpdateNumById(id, num) {
					err = toolLib.CreateKcErr(_const.DB_ERR)
					return
				}
			}
		}
		if len(delGiveIds) > 0 {
			if !mRcardGive.DelByIds(delGiveIds) {
				err = toolLib.CreateKcErr(_const.DB_ERR)
				return
			}
		}
	}
	//修改赠品描述
	//1.删除原有赠品描述
	if err = rdesm.DelByRcardId(args.RcardId); err != nil {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//2.新增赠品描述
	if hasGive == 1 && len(args.GiveSingleDesc) > 0 {
		giveSingleDescStr, _ := json.Marshal(args.GiveSingleDesc)
		descData := map[string]interface{}{
			rdesm.Field.F_rcard_id: args.RcardId,
			rdesm.Field.F_desc:     string(giveSingleDescStr),
		}
		if rdesm.Insert(descData) <= 0 {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	mRcard.Model.Commit()
	err = r.EditRcardExt(args.RcardId, args.Notes)

	return
}

//获取充值卡详情
func (r *RcardLogic) RcardInfo(ctx context.Context, args *cards.ArgsRcardInfo) (reply cards.ReplyRcardInfo, err error) {
	reply = cards.ReplyRcardInfo{
		RcardId:        0,
		RcardBase:      cards.RcardBase{},
		ShareLink:      "",
		Notes:          []cards.CardNote{},
		SsId:           0,
		ImgHash:        "",
		ImgUrl:         "",
		IncludeSingles: []cards.IncSingleDetail{},
		GiveSingles:    []cards.IncSingleDetail{},
		IncProducts:    []cards.IncProductDetail{},
		RcardRules:     []cards.ListsRechargeRules{},
		//IsGround:       0,
		ShopStatus:     0,
		BusInfo:        cards.BusInfo{},
		IsHaveDiscount: 1,
		SingleDiscount: 10.0,  // 10.0 表示不打折
		IsAllSingle:    false, //适用于全部单项目
		IsAllProduct:   false, //适用于全部商品
	}

	if args.RcardId <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	//获取卡项信息
	mRcard := new(models.RcardModel).Init()
	rcardInfo := mRcard.GetByRcardId(args.RcardId)
	if len(rcardInfo) == 0 {
		err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
		return
	}

	var shopStatus int = 0
	var shopSales int = 0
	if args.ShopId > 0 {
		mShopRcard := new(models.ShopRcardModel).Init()
		shopRcard := mShopRcard.GetByShopidAndRcardid(args.ShopId, args.RcardId)
		if len(shopRcard) == 0 {
			err = toolLib.CreateKcErr(_const.NO_IN_SHOP)
			return
		}
		shopStatus, _ = strconv.Atoi(shopRcard[mShopRcard.Field.F_status].(string))
		reply.SsId, _ = strconv.Atoi(shopRcard[mShopRcard.Field.F_id].(string))
		shopSales, _ = strconv.Atoi(shopRcard[mShopRcard.Field.F_sales].(string))
	}
	reply.ShareLink = tools.GetShareLink(args.RcardId, args.ShopId, cards.ITEM_TYPE_rcard)
	rcardBase := cards.RcardBase{}
	mapstructure.WeakDecode(rcardInfo, &rcardBase)
	imgId, _ := strconv.Atoi(rcardInfo[mRcard.Field.F_img_id].(string))
	//isGround, _ := strconv.Atoi(rcardInfo[mRcard.Field.F_is_ground].(string))
	isHaveDiscount, _ := strconv.Atoi(rcardInfo[mRcard.Field.F_is_have_discount].(string))

	imgHash, imgUrl := getImg(ctx, imgId, cards.ITEM_TYPE_rcard)
	if args.ShopId > 0 {
		rcardBase.Sales = shopSales
	}
	reply.RcardBase = rcardBase
	reply.RcardId = args.RcardId
	reply.ImgHash = imgHash
	reply.ImgUrl = imgUrl
	//reply.IsGround = isGround
	reply.ShopStatus = shopStatus
	reply.IsHaveDiscount = isHaveDiscount

	//商户信息
	if err = getBusInfo(ctx, reply.BusID, &reply.BusInfo); err != nil {
		err = toolLib.CreateKcErr(_const.SHOP_INFO_ERR)
		return
	}

	//获取包含的套餐和赠送的项目
	var singleIds []int
	giveSingles := []map[string]interface{}{}
	mRcardSingle := new(models.RcardSingleModel).Init()
	rcardSingles := mRcardSingle.Find(map[string]interface{}{mRcardSingle.Field.F_rcard_id: args.RcardId})
	// 获取单项目折扣率   0 或 10.0表示没有折扣
	if len(rcardSingles) > 0 {
		singleDiscount, _ := strconv.ParseFloat(rcardSingles[mRcardSingle.Field.F_discount].(string), 64)
		reply.SingleDiscount = singleDiscount
		//singleIds = functions.ArrayValue2Array(mRcardSingle.Field.F_single_id, rcardSingles)
		if rcardSingles[mRcardSingle.Field.F_single_id].(string) == "0" {
			reply.IsAllSingle = true
		}
	}

	mRcardGive := new(models.RcardGiveModel).Init()
	if rcardInfo[mRcard.Field.F_has_give_signle].(string) == strconv.Itoa(cards.HAS_GIVE_SINGLE_yes) {
		giveSingles = mRcardGive.GetByRcardid(args.RcardId)
		giveSingleIds := functions.ArrayValue2Array(mRcardGive.Field.F_single_id, giveSingles)
		singleIds = append(singleIds, giveSingleIds...)
	}
	if len(singleIds) > 0 {
		allSingles := getIncSingles(ctx, singleIds)

		if len(giveSingles) > 0 {
			for _, single := range giveSingles {
				singleId, _ := strconv.Atoi(single[mRcardGive.Field.F_single_id].(string))
				sinfo := allSingles[singleId]
				sinfo.Num, _ = strconv.Atoi(single[mRcardGive.Field.F_num].(string))
				sinfo.PeriodOfValidity, _ = strconv.Atoi(single[mRcardGive.Field.F_period_of_validity].(string))
				reply.GiveSingles = append(reply.GiveSingles, sinfo)
			}
			//获取赠品描述信息
			rdescm := new(models.RcardGiveDescModel).Init()
			desc, ok := rdescm.GetByRcardId(args.RcardId)[rdescm.Field.F_desc].(string)
			if ok {
				json.Unmarshal([]byte(desc), &reply.GiveSingleDesc)
			} else {
				reply.GiveSingleDesc = []cards.GiveSingleDesc{}
			}
		}
	}

	// 获取充值卡规则
	mRcardRule := new(models.RcardRuleModel).Init()
	rcardRuleInfo := mRcardRule.GetByRcardid(args.RcardId)

	for _, listRcard := range rcardRuleInfo {
		rechargeAmountStr := listRcard[mRcardRule.Field.F_price].(string)
		donationAmountStr := listRcard[mRcardRule.Field.F_give_price].(string)
		rechargeAmount, _ := strconv.ParseFloat(rechargeAmountStr, 64)
		donationAmount, _ := strconv.ParseFloat(donationAmountStr, 64)
		id, _ := strconv.Atoi(listRcard[mRcardRule.Field.F_id].(string))
		name := fmt.Sprintf("充%s,赠送%s元", rechargeAmountStr, donationAmountStr)
		reply.RcardRules = append(reply.RcardRules, cards.ListsRechargeRules{
			Id:             id,
			RechargeAmount: rechargeAmount,
			DonationAmount: donationAmount,
			Name:           name,
		})
	}

	//获取包含的产品
	mRcardGoods := new(models.RcardGoodsModel).Init()
	rcardProducts := mRcardGoods.Find(map[string]interface{}{mRcardGoods.Field.F_rcard_id: args.RcardId})
	if len(rcardProducts) > 0 {
		if rcardProducts[mRcardGoods.Field.F_product_id].(string) == "0" {
			reply.IsAllProduct = true
		}
	}
	//	if !reply.IsAllProduct{
	//		productIds := functions.ArrayValue2Array(mRcardGoods.Field.F_product_id, rcardProducts)
	//		var allProducts map[int]cards.IncProductDetail
	//		if allProducts, err = getIncProducts(ctx, productIds); err != nil {
	//			return
	//		}
	//		for _, rcardProduct := range rcardProducts {
	//			productId, _ := strconv.Atoi(rcardProduct[mRcardGoods.Field.F_product_id].(string))
	//			pInfo := allProducts[productId]
	//			reply.IncProducts = append(reply.IncProducts, pInfo)
	//		}
	//	}
	//}

	//获取CardExt
	r.GetRcardExt(args.RcardId, &reply.Notes)

	// 获取充值卡门店添加详情  15 -- []int {3,4,6}
	busId, _ := strconv.Atoi(rcardInfo[mRcard.Field.F_bus_id].(string))
	mRardShopModel := new(models.RcardShopModel).Init()
	rCardShopLists := mRardShopModel.GetByRcardIdAndBusId(args.RcardId, busId)

	rCardShopIds := make([]int, 0)
	for _, rInfoValue := range rCardShopLists {
		sshopId, _ := strconv.Atoi(rInfoValue[mRardShopModel.Field.F_shop_id].(string))
		rCardShopIds = append(rCardShopIds, sshopId)
	}
	var replyShop []bus2.ReplyShopName
	rLists := make([]cards.ReplyShopName, 0) // 临时存放返回数据
	rpcBus := new(bus.Shop).Init()
	defer rpcBus.Close()
	err = rpcBus.GetShopNameByShopIds(ctx, &rCardShopIds, &replyShop)
	if err != nil {
		return
	}
	for _, v := range replyShop {
		rLists = append(rLists, cards.ReplyShopName{
			ShopId:     v.ShopId,
			ShopName:   v.ShopName,
			BranchName: v.BranchName,
		})
	}
	reply.ShopLists = rLists

	//浏览次数加1
	redis.RedisGlobMgr.Hincrby(constkey.RCARD_CLIKS, strconv.Itoa(args.RcardId), 1)
	return
}

//获取总店充值卡列表
func (r *RcardLogic) BusRcardPage(ctx context.Context, args *cards.ArgsBusRcardPage) (reply cards.ReplyRcardPage, err error) {
	start := args.GetStart()
	limit := args.GetPageSize()

	reply = cards.ReplyRcardPage{
		TotalNum:  0,
		Lists:     []cards.ListRcard{},
		IndexImgs: map[int]string{},
	}
	if args.BusId <= 0 || start < 0 || limit < 0 {
		return
	}

	mRcard := new(models.RcardModel).Init()
	rcards := make([]map[string]interface{}, 0)
	reply.IndexImgs = make(map[int]string)

	//子店已添加的卡项
	var shopAddCards []map[string]interface{}
	scModel := new(models.ShopRcardModel).Init()
	where := []base.WhereItem{
		{mRcard.Field.F_bus_id, args.BusId},
		{mRcard.Field.F_is_del, cards.IS_BUS_DEL_no},
	}
	if args.ShopId > 0 {
		shopAddCards = scModel.SelectRcardsByWherePage([]base.WhereItem{{scModel.Field.F_shop_id, args.ShopId},{scModel.Field.F_is_del, cards.IS_BUS_DEL_no}}, 0, 0)
		if args.FilterShopHasAdd && len(shopAddCards) > 0 {
			shopHasAddRcardIds := functions.ArrayValue2Array(scModel.Field.F_rcard_id, shopAddCards)
			where = append(where, base.WhereItem{mRcard.Field.F_rcard_id, []interface{}{"NOT IN", shopHasAddRcardIds}})
		}
	}
	// 获取未删除充值卡总数量
	rcards = mRcard.SelectRcardsByWherePage(where, start, limit)
	reply.TotalNum = mRcard.GetNumByWhere(where)

	if len(rcards) == 0 {
		return
	}

	rcardsArr := []cards.RcardBase{}
	mapstructure.WeakDecode(rcards, &rcardsArr)
	for k, rcard := range rcards {
		rcardId, _ := strconv.Atoi(rcard[mRcard.Field.F_rcard_id].(string))
		sales, _ := strconv.Atoi(rcard[mRcard.Field.F_sales].(string))
		isGround, _ := strconv.Atoi(rcard[mRcard.Field.F_is_ground].(string))
		clicks, _ := redis2.Int(redis.RedisGlobMgr.Hget(constkey.RCARD_CLIKS, rcard[mRcard.Field.F_rcard_id].(string)))
		imgId, _ := strconv.Atoi(rcard[mRcard.Field.F_img_id].(string))
		shopStatus, shopHasAdd, shopItemId, shopDelStatus := 0, 0, 0, 0
		for _, shopCard := range shopAddCards {
			if rcard[mRcard.Field.F_rcard_id].(string) == shopCard[scModel.Field.F_rcard_id].(string) { //表明当前子店已添加该卡项
				shopItemId, _ = strconv.Atoi(shopCard[scModel.Field.F_id].(string))
				shopStatus, _ = strconv.Atoi(shopCard[scModel.Field.F_status].(string))
				shopHasAdd = 1
				shopDelStatus, _ = strconv.Atoi(shopCard[scModel.Field.F_is_del].(string))
				break
			}
		}
		reply.Lists = append(reply.Lists, cards.ListRcard{
			RcardId:       rcardId,
			RcardBase:     rcardsArr[k],
			CtimeStr:      functions.TimeToStr(int64(rcardsArr[k].Ctime)),
			Clicks:        clicks,
			Sales:         sales,
			IsGround:      isGround,
			ShopStatus:    shopStatus,
			ShopHasAdd:    shopHasAdd,
			ImgId:         imgId,
			ShopItemId:    shopItemId,
			ShopDelStatus: shopDelStatus,
		})
	}

	//获取图片信息
	imgIds := functions.ArrayValue2Array(mRcard.Field.F_img_id, rcards)
	reply.IndexImgs = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_rcard)
	return
}

//设置适用门店(废用）
func (r *RcardLogic) SetRcardShop(ctx context.Context, args *cards.ArgsSetRcardShop) (err error) {
	busId, err := checkBus(args.BsToken, true)
	if err != nil {
		return
	}
	if len(args.RcardIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	if args.IsAllShop == false && len(args.ShopIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	//充值卡id重复提交判断
	realRcardIds := functions.ArrayUniqueInt(args.RcardIds)
	if len(realRcardIds) != len(args.RcardIds) {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//店铺id重复提交判断
	realShopIds := functions.ArrayUniqueInt(args.ShopIds)
	if len(realShopIds) != len(args.ShopIds) && args.IsAllShop == false {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//检查充值卡是否属于商家
	busRcards, err := r.checkRcardids(busId, realRcardIds)
	if err != nil {
		return
	}
	mRcardShop := new(models.RcardShopModel).Init()
	var insertData = []map[string]interface{}{}
	//适用全部门店
	if args.IsAllShop == true {
		for _, rcardId := range realRcardIds {
			insertData = append(insertData, map[string]interface{}{
				mRcardShop.Field.F_rcard_id: rcardId,
				mRcardShop.Field.F_bus_id:   busId,
				mRcardShop.Field.F_shop_id:  0,
			})
		}
	} else {
		//检查门店id是否合法
		rpcShop := new(bus.Shop).Init()
		checkArgs := &bus2.ArgsCheckShop{
			BusId:   busId,
			ShopIds: realShopIds,
		}
		checkReply := &bus2.ReplyCheckShop{}
		err = rpcShop.CheckBusShop(ctx, checkArgs, checkReply)
		if err != nil {
			return
		}
		for _, rcardId := range realRcardIds {
			for _, shopid := range realShopIds {
				insertData = append(insertData, map[string]interface{}{
					mRcardShop.Field.F_rcard_id: rcardId,
					mRcardShop.Field.F_bus_id:   busId,
					mRcardShop.Field.F_shop_id:  shopid,
				})
			}
		}
	}

	//处理适用门店
	if len(insertData) > 0 {
		mRcardShop.DelByRcardids(realRcardIds)
		mRcardShop.InsertAll(insertData)
	}

	//提取目前是上架状态的充值卡id
	//var onGroundRcardids []int
	//for _, rcard := range busRcards {
	//	if rcard["is_ground"].(string) == strconv.Itoa(cards.IS_GROUND_yes) {
	//		tmpRcardid, _ := strconv.Atoi(rcard["rcard_id"].(string))
	//		onGroundRcardids = append(onGroundRcardids, tmpRcardid)
	//	}
	//}

	// 提前未删除的
	var noDelIds []int
	for _, rcard := range busRcards {
		if rcard["is_del"].(string) == strconv.Itoa(cards.IS_BUS_DEL_no) {
			tmpRcardId, _ := strconv.Atoi(rcard["rcard_id"].(string))
			noDelIds = append(noDelIds, tmpRcardId)
		}
	}

	mShopRcard := new(models.ShopRcardModel).Init()
	shopRcards := mShopRcard.GetByRcardids(realRcardIds)
	var downIds = []int{}
	for _, shoprcard := range shopRcards {
		shopStatus, _ := strconv.Atoi(shoprcard[mShopRcard.Field.F_status].(string))
		shopRcardId, _ := strconv.Atoi(shoprcard[mShopRcard.Field.F_id].(string))
		shopid, _ := strconv.Atoi(shoprcard[mShopRcard.Field.F_shop_id].(string))
		rcardId, _ := strconv.Atoi(shoprcard[mShopRcard.Field.F_rcard_id].(string))
		// 一期优化----yinjinlin-2021-04-08
		// 全部适用：之前门店已经添加过并且分： 1>.上架/或下架 状态 ：都不需要更新
		//                                2>.被总店禁用 状态 ：需要更新为下架状态
		if shopStatus == cards.STATUS_DISABLE && functions.InArray(rcardId, noDelIds) {
			if args.IsAllShop == true || functions.InArray(shopid, realShopIds) {
				downIds = append(downIds, shopRcardId)
			}
		}
	}

	if len(downIds) > 0 {
		mShopRcard.UpdateByIds(downIds, map[string]interface{}{
			mShopRcard.Field.F_status:     cards.STATUS_OFF_SALE,
			mShopRcard.Field.F_under_time: time.Now().Unix(),
		})
	}

	//如果是设置成部分门店适用，需要把在门店已上架的但是现在改为不适用的门店套餐改为总店警用
	if args.IsAllShop == false {
		var disableIds = []int{}
		for _, shoprcard := range shopRcards {
			shopStatus, _ := strconv.Atoi(shoprcard[mShopRcard.Field.F_status].(string))
			// 被总店已删除
			if shopStatus == cards.STATUS_DISABLE {
				continue
			}
			shopid, _ := strconv.Atoi(shoprcard[mShopRcard.Field.F_shop_id].(string))
			if functions.InArray(shopid, realShopIds) == false {
				disableId, _ := strconv.Atoi(shoprcard[mShopRcard.Field.F_id].(string))
				disableIds = append(disableIds, disableId)
			}
		}
		if len(disableIds) > 0 {
			mShopRcard.UpdateByIds(disableIds, map[string]interface{}{
				mShopRcard.Field.F_status:     cards.STATUS_DISABLE,
				mShopRcard.Field.F_under_time: time.Now().Unix(),
			})
			//添加维护es的shop-item文档的任务
			setShopItem(ctx, []int{}, 0, cards.ITEM_TYPE_rcard, disableIds)
		}
	}

	return
}

//总店上下架充值卡
func (r *RcardLogic) DownUpRcard(ctx context.Context, args *cards.ArgsDownUpRcard) (err error) {
	busId, err := checkBus(args.BsToken, true)
	if err != nil {
		return
	}

	//检查套餐是否属于企业
	rcards, err := r.checkRcardids(busId, args.RcardIds)
	if err != nil {
		return
	}
	var rcardsStruct []struct {
		RcardId  int
		IsGround int
	}
	mapstructure.WeakDecode(rcards, &rcardsStruct)
	var downIds, upIds []int
	for _, rcard := range rcardsStruct {
		if rcard.IsGround == cards.IS_GROUND_no {
			downIds = append(downIds, rcard.RcardId)
		} else {
			upIds = append(upIds, rcard.RcardId)
		}
	}

	mRcard := new(models.RcardModel).Init()
	mShopRcard := new(models.ShopRcardModel).Init()
	//下架操作
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		r := mRcard.UpdateByRcardids(upIds, map[string]interface{}{
			mRcard.Field.F_is_ground:     cards.IS_GROUND_no,
			mRcard.Field.F_under_time:    time.Now().Unix(),
			mRcard.Field.F_sale_shop_num: 0,
		})
		if r == false {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//将分店的套餐设置为总店禁用
		mShopRcard.UpdateByRcardids(upIds, map[string]interface{}{
			mShopRcard.Field.F_status:     cards.STATUS_DISABLE,
			mShopRcard.Field.F_under_time: time.Now().Unix(),
		})
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, upIds, 0, cards.ITEM_TYPE_rcard)
	}

	//上架操作
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		r := mRcard.UpdateByRcardids(downIds, map[string]interface{}{
			mRcard.Field.F_is_ground:  cards.IS_GROUND_yes,
			mRcard.Field.F_under_time: 0,
		})
		if r == false {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//将已添加到子店的套餐解除总店禁用状态 需要将不适用的门店套餐id 过滤掉
		shopRcards := mShopRcard.GetByRcardids(downIds)
		if len(shopRcards) == 0 {
			return
		}
		mRcardShop := new(models.RcardShopModel).Init()
		rcardShops := mRcardShop.GetByRcardids(downIds)

		var rcardShopArr = []string{} //rcardId_shopid
		for _, rcardsshop := range rcardShops {
			rcardShopArr = append(rcardShopArr, fmt.Sprintf("%s_%s", rcardsshop[mRcardShop.Field.F_rcard_id].(string), rcardsshop[mRcardShop.Field.F_shop_id].(string)))
		}
		var unDisableIds = []int{}
		for _, shoprcard := range shopRcards {
			rcardidShopidStr := fmt.Sprintf("%s_%s", shoprcard[mShopRcard.Field.F_rcard_id].(string), shoprcard[mShopRcard.Field.F_shop_id].(string))
			rcardidAllStr := fmt.Sprintf("%s_0", shoprcard[mShopRcard.Field.F_rcard_id].(string))

			if functions.InArray(rcardidAllStr, rcardShopArr) || functions.InArray(rcardidShopidStr, rcardShopArr) {
				shoprcardId, _ := strconv.Atoi(shoprcard[mShopRcard.Field.F_id].(string))
				unDisableIds = append(unDisableIds, shoprcardId)
			}
		}
		if len(unDisableIds) > 0 {
			mShopRcard.UpdateByIds(unDisableIds, map[string]interface{}{
				mShopRcard.Field.F_status: cards.STATUS_OFF_SALE,
			})
		}
	}

	return
}

// 总店删除充值卡
func (r *RcardLogic) DeleteRcard(ctx context.Context, args *cards.ArgsDelRcard) (err error) {

	//更新总店
	rcalModel := new(models.RcardModel).Init()
	rBool := rcalModel.UpdateByRcardids(args.RcardIds, map[string]interface{}{
		rcalModel.Field.F_is_del:   cards.IS_BUS_DEL_yes,
		rcalModel.Field.F_del_time: time.Now().Unix(),
	})
	if rBool == false {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	// 更新分店
	shopRcalModel := new(models.ShopRcardModel).Init()
	shopBool := shopRcalModel.UpdateByRcardids(args.RcardIds, map[string]interface{}{
		shopRcalModel.Field.F_is_del:   cards.IS_BUS_DEL_yes,
		shopRcalModel.Field.F_del_time: time.Now().Unix(),
	})
	if shopBool == false {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.RcardIds, cards.ITEM_TYPE_rcard) {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	return
}

//子店获取适用本店的充值卡列表
func (r *RcardLogic) ShopGetBusRcardPage(ctx context.Context, args *cards.ArgsShopGetBusRcardPage) (reply cards.ReplyRcardPage, err error) {
	start := args.GetStart()
	limit := args.GetPageSize()
	if args.BusId <= 0 || args.ShopId <= 0 || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	reply = cards.ReplyRcardPage{
		TotalNum:  0,
		Lists:     []cards.ListRcard{},
		IndexImgs: map[int]string{},
	}

	mRcardShop := new(models.RcardShopModel).Init()
	rcardShops := mRcardShop.GetPageByShopId(args.BusId, args.ShopId, start, limit)
	if len(rcardShops) == 0 {
		return
	}
	//获取充值卡基本信息
	rcardIds := functions.ArrayValue2Array(mRcardShop.Field.F_rcard_id, rcardShops)
	mRcard := new(models.RcardModel).Init()
	rcards := mRcard.GetByRcardids(rcardIds)
	if len(rcards) == 0 {
		return
	}
	var rcardArr []cards.RcardBase
	_ = mapstructure.WeakDecode(rcards, &rcardArr)
	rcardsMap := map[string]cards.ListRcard{}
	//获取门店添加状态
	mShopRcardModel := new(models.ShopRcardModel).Init()
	shopRcards := mShopRcardModel.GetByShopidAndRcardids(args.ShopId, rcardIds)
	shopRcardids := functions.ArrayValue2Array(mShopRcardModel.Field.F_rcard_id, shopRcards)

	for k, rcard := range rcards {
		rcardId, _ := strconv.Atoi(rcard[mRcard.Field.F_rcard_id].(string))
		sales, _ := strconv.Atoi(rcard[mRcard.Field.F_sales].(string))
		//isGround, _ := strconv.Atoi(rcard[mRcard.Field.F_is_ground].(string))
		clicks, _ := redis2.Int(redis.RedisGlobMgr.Hget(constkey.RCARD_CLIKS, rcard[mRcard.Field.F_rcard_id].(string)))
		imgId, _ := strconv.Atoi(rcard[mRcard.Field.F_img_id].(string))
		shopHasAdd := 0
		if functions.InArray(rcardId, shopRcardids) {
			shopHasAdd = 1
		}

		rcardsMap[rcard[mRcard.Field.F_rcard_id].(string)] = cards.ListRcard{
			RcardId:   rcardId,
			RcardBase: rcardArr[k],
			CtimeStr:  functions.TimeToStr(int64(rcardArr[k].Ctime)),
			Clicks:    clicks,
			Sales:     sales,
			//IsGround:   isGround,
			ShopStatus: 0,
			ShopHasAdd: shopHasAdd,
			ImgId:      imgId,
			ShopItemId: 0,
		}
	}

	for _, rcardsshop := range rcardShops {
		reply.Lists = append(reply.Lists, rcardsMap[rcardsshop[mRcardShop.Field.F_rcard_id].(string)])
	}
	//获取图片信息
	imgIds := functions.ArrayValue2Array(mRcard.Field.F_img_id, rcards)
	reply.IndexImgs = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_rcard)
	//获取数量
	reply.TotalNum = mRcardShop.GetNumByShopId(args.BusId, args.ShopId)

	return
}

//子店添加充值卡到门店
func (r *RcardLogic) ShopAddRcard(ctx context.Context, shopId, busId int, rcardIds []int) (err error) {
	rcardIds = functions.ArrayUniqueInt(rcardIds)
	if len(rcardIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//检查充值卡id是否适用当前门店
	//mRcardShop := new(models.RcardShopModel).Init()
	//rcardShops := mRcardShop.GetByShopIdAndRcardids(busId, shopId, rcardIds)
	//if len(rcardShops) != len(rcardIds) {
	//	err = toolLib.CreateKcErr(_const.NO_IN_SHOP)
	//	return
	//}
	//过滤掉已经添加了的充值卡id
	mShopRcard := new(models.ShopRcardModel).Init()
	shopRcardLists := mShopRcard.GetByShopidAndRcardids(shopId, rcardIds)
	shopRcardIds := functions.ArrayValue2Array(mShopRcard.Field.F_rcard_id, shopRcardLists)

	// 刷选出已经添加过并且删除的数据
	delHcardIdSlice := make([]int, 0)
	for _, hcardMap := range shopRcardLists {
		isDel, _ := strconv.Atoi(hcardMap[mShopRcard.Field.F_is_del].(string))
		if isDel == cards.IS_BUS_DEL_yes {
			delHcardId, _ := strconv.Atoi(hcardMap[mShopRcard.Field.F_rcard_id].(string))
			delHcardIdSlice = append(delHcardIdSlice, delHcardId)
		}
	}

	// 更新门店之前添加过并删除的数据
	if len(delHcardIdSlice) > 0 {
		// 更新数据删除和上下架状态
		if updateBool := mShopRcard.UpdateByRcardidsAndShopId(delHcardIdSlice, shopId, map[string]interface{}{
			mShopRcard.Field.F_is_del: cards.IS_BUS_DEL_no,
			mShopRcard.Field.F_status: cards.STATUS_OFF_SALE,
		}); !updateBool {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
		//更新卡项关联表
		sirModel := new(models.ShopItemRelationModel).Init()
		if b := sirModel.UpdateByItemIdsAndShopId(delHcardIdSlice, cards.ITEM_TYPE_rcard, shopId, map[string]interface{}{
			sirModel.Field.F_is_del: cards.IS_BUS_DEL_no,
			sirModel.Field.F_status: cards.STATUS_OFF_SALE,
		}); !b {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}


	// 实际需要添加的
	addRcardids := make([]int, 0)
	for _, rcardid := range rcardIds {
		if functions.InArray(rcardid, shopRcardIds) == false {
			addRcardids = append(addRcardids, rcardid)
		}
	}

	//校验当前门店是否已经将卡项内涉及到的单项目添加到自己的门店内
	allSingle, singleIds, err := new(ItemLogic).getItemCardIncSingleIds(addRcardids, cards.ITEM_TYPE_rcard)
	if err != nil {
		return
	}
	if err = new(ItemLogic).validShopSingleContainItemCardSingles(shopId, busId, allSingle, singleIds); err != nil {
		return
	}

	//判断本店是否存在包含商品
	err, checkSuccess := new(ItemLogic).CheckProductInShop(ctx, busId, shopId, cards.ITEM_TYPE_rcard, addRcardids)
	if err != nil {
		return err
	}
	if !checkSuccess {
		return toolLib.CreateKcErr(_const.SHOP_PRODUCT_NOT_CONTAIN_BUS_PRODUCT)
	}

	// 需要添加数据
	var addData []map[string]interface{}
	shopItemRelationData := make([]map[string]interface{}, 0)
	shopItemRelationModel := new(models.ShopItemRelationModel).Init()
	for _, rcardId := range addRcardids {
		status := cards.STATUS_OFF_SALE
		ctime := time.Now().Local().Unix()
		addData = append(addData, map[string]interface{}{
			mShopRcard.Field.F_rcard_id: rcardId,
			mShopRcard.Field.F_status:   status,
			mShopRcard.Field.F_shop_id:  shopId,
			mShopRcard.Field.F_ctime:    ctime,
		})
		shopItemRelationData = append(shopItemRelationData, map[string]interface{}{
			shopItemRelationModel.Field.F_item_id:   rcardId,
			shopItemRelationModel.Field.F_item_type: cards.ITEM_TYPE_rcard,
			shopItemRelationModel.Field.F_status:    cards.STATUS_OFF_SALE,
			shopItemRelationModel.Field.F_shop_id:   shopId,
			shopItemRelationModel.Field.F_is_del:    cards.ITEM_IS_DEL_NO,
		})
	}
	//过滤的数据添加到门店充值卡表
	if len(addData) > 0 {
		id := mShopRcard.InsertAll(addData)
		if id < 0 {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}
	//门店卡项关联表数据插入
	if len(shopItemRelationData) > 0 {
		if shopItemRelationModel.InsertAll(shopItemRelationData) <= 0 {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	// 初始化适用门店模型
	rcardShopModel := new(models.RcardShopModel).Init()
	rcardShopSms := rcardShopModel.GetByRcardids(rcardIds)
	rcShopIds := functions.ArrayValue2Array(rcardShopModel.Field.F_rcard_id, rcardShopSms)
	addRcShopIds := make([]int, 0)
	for _, rcardId := range rcardIds {
		if functions.InArray(rcardId, rcShopIds) == false {
			addRcShopIds = append(addRcShopIds, rcardId)
		}
	}
	if len(addRcShopIds) == 0 {
		return
	}

	var addRcShopData []map[string]interface{} // 添加适用门店表的数据
	for _, addSmShopId := range addRcShopIds {
		addRcShopData = append(addRcShopData, map[string]interface{}{
			rcardShopModel.Field.F_rcard_id: addSmShopId,
			rcardShopModel.Field.F_shop_id:  shopId,
			rcardShopModel.Field.F_bus_id:   busId,
		})
	}
	// 过滤的数据添加到适用套餐表
	if len(addRcShopData) > 0 {
		id := rcardShopModel.InsertAll(addRcShopData)
		if id < 0 {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}

	//获取充值卡在总店的上下架状态
	//mRcard := new(models.RcardModel).Init()
	//rcards := mRcard.GetByRcardids(addRcardids, []string{
	//	mRcard.Field.F_rcard_id,
	//	mRcard.Field.F_is_ground,
	//})

	// 获取充值卡在总店的状态
	//rcards := mRcard.GetByRcardids(addRcardids, []string{
	//	mRcard.Field.F_rcard_id,
	//	mRcard.Field.F_is_del,
	//})
	//
	//// 数组转map
	//rcardsMap := functions.ArrayRebuild(mRcard.Field.F_rcard_id, rcards)
	//var addData []map[string]interface{}
	//for _, addRcardid := range addRcardids {
	//	status := cards.STATUS_OFF_SALE
	//	if _, ok := rcardsMap[strconv.Itoa(addRcardid)]; ok {
	//		if newrcardmap, ok2 := rcardsMap[strconv.Itoa(addRcardid)].(map[string]interface{}); ok2 {
	//			//isground, _ := strconv.Atoi(newrcardmap[mRcard.Field.F_is_ground].(string))
	//			isDel, _ := strconv.Atoi(newrcardmap[mRcard.Field.F_is_del].(string))
	//			//if isground == cards.IS_GROUND_no {
	//			//	status = cards.STATUS_DISABLE
	//			//}
	//			// 如果总店删除 更改门店状态为被总店删除
	//			if isDel == cards.IS_BUS_DEL_no {
	//				status = cards.STATUS_OFF_SALE
	//			}
	//		}
	//	}
	//
	//	ctime := time.Now().Local().Unix()
	//	addData = append(addData, map[string]interface{}{
	//		mShopRcard.Field.F_rcard_id: addRcardid,
	//		mShopRcard.Field.F_status:   status,
	//		mShopRcard.Field.F_shop_id:  shopId,
	//		mShopRcard.Field.F_ctime:    ctime,
	//	})
	//}
	//
	//if len(addData) > 0 {
	//	mShopRcard.InsertAll(addData)
	//}

	return
}

//获取子店充值卡列表
func (r *RcardLogic) ShopRcardPage(ctx context.Context, args *cards.ArgsShopRcardPage) (reply cards.ReplyShopRcardPage, err error) {
	start := args.GetStart()
	limit := args.GetPageSize()
	if args.ShopId <= 0 || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	reply = cards.ReplyShopRcardPage{
		TotalNum:  0,
		Lists:     []cards.ListRcard{},
		IndexImgs: map[int]string{},
	}
	reply.Lists = make([]cards.ListRcard, 0)
	//获取门店的充值卡数据
	mShopRcard := new(models.ShopRcardModel).Init()
	shopRcards := mShopRcard.GetPageByShopId(args.ShopId, start, limit, args.Status)
	rcardIds := functions.ArrayValue2Array(mShopRcard.Field.F_rcard_id, shopRcards)
	//获取充值卡基本信息
	mRcard := new(models.RcardModel).Init()
	rcards := mRcard.GetByRcardids(rcardIds)
	if len(rcards) == 0 {
		return
	}
	var rcardsArr []cards.RcardBase
	_ = mapstructure.WeakDecode(rcards, &rcardsArr)
	rcardsMap := map[string]cards.ListRcard{}

	//获取不同卡项-适用单项目的个数和赠送单项目的个数
	gaagsNumMap := GetApplyAndGiveSingleNum(rcardIds, cards.ITEM_TYPE_rcard)
	for k, rcard := range rcards {
		rcardId, _ := strconv.Atoi(rcard[mRcard.Field.F_rcard_id].(string))
		//isGround, _ := strconv.Atoi(rcard[mRcard.Field.F_is_ground].(string))
		//isDel, _ := strconv.Atoi(rcard[mRcard.Field.F_is_del].(string))
		clicks, _ := redis2.Int(redis.RedisGlobMgr.Hget(constkey.RCARD_CLIKS, rcard[mRcard.Field.F_rcard_id].(string)))
		imgId, _ := strconv.Atoi(rcard[mRcard.Field.F_img_id].(string))

		rcardsMap[rcard[mRcard.Field.F_rcard_id].(string)] = cards.ListRcard{
			RcardId:   rcardId,
			RcardBase: rcardsArr[k],
			CtimeStr:  functions.TimeToStr(int64(rcardsArr[k].Ctime)),
			Clicks:    clicks,
			Sales:     0,
			//IsGround:   isGround,
			//IsDel:      isDel,
			ShopStatus:     0,
			ShopHasAdd:     1,
			ImgId:          imgId,
			ShopItemId:     0,
			IsAllSingle:    gaagsNumMap[rcardId].IsAllSingle,
			ApplySingleNum: gaagsNumMap[rcardId].ApplySingleNum,
			GiveSingleNum:  gaagsNumMap[rcardId].GiveSingleNum,
		}
	}

	for _, shoprcard := range shopRcards {
		listShopRcard := rcardsMap[shoprcard[mShopRcard.Field.F_rcard_id].(string)]
		listShopRcard.ShopItemId, _ = strconv.Atoi(shoprcard[mShopRcard.Field.F_id].(string))
		listShopRcard.ShopStatus, _ = strconv.Atoi(shoprcard[mShopRcard.Field.F_status].(string))
		listShopRcard.Sales, _ = strconv.Atoi(shoprcard[mShopRcard.Field.F_sales].(string))
		reply.Lists = append(reply.Lists, listShopRcard)
	}

	//获取图片信息
	imgIds := functions.ArrayValue2Array(mRcard.Field.F_img_id, rcards)
	reply.IndexImgs = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_rcard)
	//获取数量
	reply.TotalNum = mShopRcard.GetNumByShopId(args.ShopId, args.Status)

	return
}

//子店上下架充值卡
func (r *RcardLogic) ShopDownUpRcard(ctx context.Context, args *cards.ArgsShopDownUpRcard) (err error) {
	shopId := args.ShopId
	args.ShopRcardIds = functions.ArrayUniqueInt(args.ShopRcardIds)
	if len(args.ShopRcardIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//获取门店充值卡信息
	mShopRcard := new(models.ShopRcardModel).Init()
	shopRcards := mShopRcard.GetByShopIDAndCardIDs(shopId, args.ShopRcardIds)
	var shopRcardStruct []struct {
		Id      int
		ShopId  int
		Status  int
		RcardId int
	}
	var upIds, downIds, upRcardids, downRcardids []int
	_ = mapstructure.WeakDecode(shopRcards, &shopRcardStruct)
	for _, shopRcard := range shopRcardStruct {
		if shopId != shopRcard.ShopId {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
		if shopRcard.Status == cards.STATUS_OFF_SALE {
			downIds = append(downIds, shopRcard.Id)
			downRcardids = append(downRcardids, shopRcard.RcardId)
		} else if shopRcard.Status == cards.STATUS_ON_SALE {
			upIds = append(upIds, shopRcard.Id)
			upRcardids = append(upRcardids, shopRcard.RcardId)
		}
	}
	mRcard := new(models.RcardModel).Init()
	//充值卡下架
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		mShopRcard.UpdateByIds(upIds, map[string]interface{}{
			mShopRcard.Field.F_status:     cards.STATUS_OFF_SALE,
			mShopRcard.Field.F_under_time: time.Now().Unix(),
		})
		for _, rcardId := range upRcardids {
			mRcard.DecrSaleShopNumByRcardid(rcardId, 1)
		}
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, upRcardids, shopId, cards.ITEM_TYPE_rcard)
		//同步下架门店卡项关联表数据
		shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
		if !shopItemRealtionModel.UpdateStatusByItemIds(upRcardids, cards.ITEM_TYPE_rcard, cards.STATUS_OFF_SALE, shopId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//充值卡上架
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		mShopRcard.UpdateByIds(downIds, map[string]interface{}{
			mShopRcard.Field.F_status:     cards.STATUS_ON_SALE,
			mShopRcard.Field.F_under_time: 0,
		})
		for _, rcardId := range downRcardids {
			mRcard.IncrSaleShopNumByRcardid(rcardId, 1)
		}
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, downRcardids, shopId, cards.ITEM_TYPE_rcard)
		//同步上架门店卡项关联表数据
		shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
		if !shopItemRealtionModel.UpdateStatusByItemIds(downRcardids, cards.ITEM_TYPE_rcard, cards.STATUS_ON_SALE, shopId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	return
}

// 子店删除充值卡
func (r *RcardLogic) ShopDeleteRcard(ctx context.Context, args *cards.ArgsDelRcard) (err error) {
	shopId := args.ShopId
	mRcardShop := new(models.ShopRcardModel).Init()
	mRcardShop.UpdateByRcardidsAndShopId(args.RcardIds, shopId, map[string]interface{}{
		mRcardShop.Field.F_is_del:   cards.IS_BUS_DEL_yes,
		mRcardShop.Field.F_del_time: time.Now().Unix(),
	})
	//同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.RcardIds, cards.ITEM_TYPE_rcard, shopId) {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	return
}

// 店铺新增充值规则
func (r *RcardLogic) AddRechargeRules(ctx context.Context, args *cards.ArgsAddRechargeRules) (id int, err error) {
	//验证商铺，必须是总店，才可以添加充值卡规则
	_, err = checkBus(args.BsToken, true)
	if err != nil {
		return
	}

	if args.RcardId <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	rcardModel := new(models.RcardModel).Init()
	rcardInfo := rcardModel.GetByRcardId(args.RcardId)
	if len(rcardInfo) == 0 {
		err = toolLib.CreateKcErr(_const.RCARD_NOT_EXIST)
		return
	}

	discountType, _ := strconv.Atoi(rcardInfo[rcardModel.Field.F_discount_type].(string))
	if discountType == cards.DISCOUNT_TYPE_price {
		if args.DonationAmount < 0 || args.RechargeAmount < 0 || args.RechargeAmount < args.DonationAmount {
			err = toolLib.CreateKcErr(_const.PARAM_ERR)
			return
		}
	}

	if discountType == cards.DISCOUNT_TYPE_item {
		if args.RechargeAmount < 0 {
			err = toolLib.CreateKcErr(_const.PARAM_ERR)
			return
		}
	}

	//添加充值值
	addRechar := new(models.RcardRuleModel).Init()

	// 	充值卡规则，最多添加10条
	num := addRechar.GetNumByRcardId(args.RcardId, cards.RECHARGE_TYPE_NO)
	if num < cards.RCARD_MAX_NUM {
		id = addRechar.Insert(map[string]interface{}{
			addRechar.Field.F_rcard_id:   args.RcardId,
			addRechar.Field.F_price:      args.RechargeAmount,
			addRechar.Field.F_give_price: args.DonationAmount,
		})
	} else {
		err = toolLib.CreateKcErr(_const.RCARD_MAX_NUM_ERR)
		return
	}

	if id == 0 {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	return
}

// 充值卡规则编辑
func (r *RcardLogic) EditRechargeRules(ctx context.Context, args *cards.ArgsEditRechargeRules) (err error) {
	//验证商铺，必须是总店，才可以编辑充值卡规则
	_, err = checkBus(args.BsToken, true)
	if err != nil {
		return
	}
	//更新数据（rcardrule表）
	if args.RcardId <= 0 || args.Id <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	rcardModel := new(models.RcardModel).Init()
	rcardInfo := rcardModel.GetByRcardId(args.RcardId)
	discountType, _ := strconv.Atoi(rcardInfo[rcardModel.Field.F_discount_type].(string))
	if discountType == cards.DISCOUNT_TYPE_price {
		if args.DonationAmount < 0 || args.RechargeAmount < 0 || args.RechargeAmount < args.DonationAmount {
			err = toolLib.CreateKcErr(_const.PARAM_ERR)
			return
		}
	}

	if discountType == cards.DISCOUNT_TYPE_item {
		if args.RechargeAmount < 0 {
			err = toolLib.CreateKcErr(_const.PARAM_ERR)
			return
		}
	}
	editRechar := new(models.RcardRuleModel).Init()
	rs := editRechar.UpdateRelusId(args.Id, map[string]interface{}{
		editRechar.Field.F_rcard_id:   args.RcardId,
		editRechar.Field.F_price:      args.RechargeAmount,
		editRechar.Field.F_give_price: args.DonationAmount,
	})

	if rs == false {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return err
	}
	return
}

// 获取充值卡规则详情
func (r *RcardLogic) RechargeRulesInfo(ctx context.Context, args *cards.ArgsRechargeRulesInfo, reply *cards.ReplyRechargeRulesInfo) (err error) {

	if args.Id <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	dataRechar := new(models.RcardRuleModel).Init()
	dataInfo := dataRechar.GetBycrardId(args.Id)
	rcardId, _ := strconv.Atoi(dataInfo[dataRechar.Field.F_rcard_id].(string))

	//rcardModle := new(models.RcardModel).Init()
	//rcardInfo := rcardModle.GetByRcardId(rcardId)
	//discountType, _ := strconv.Atoi(rcardInfo[rcardModle.Field.F_discount_type].(string))
	//discount, _ := strconv.ParseFloat(rcardInfo[rcardModle.Field.F_discount].(string), 64)
	//reply.DiscountType = discountType
	//reply.Discount = discount
	reply.RcardId = rcardId
	reply.RechargeAmount, _ = strconv.ParseFloat(dataInfo[dataRechar.Field.F_price].(string), 64)
	reply.DonationAmount, _ = strconv.ParseFloat(dataInfo[dataRechar.Field.F_give_price].(string), 64)
	reply.IsDel, _ = strconv.Atoi(dataInfo[dataRechar.Field.F_is_del].(string))
	return
}

// 删除充值卡规则
func (r *RcardLogic) DeleRechargeRules(ctx context.Context, args *cards.ArgsDeleRechargeRules) (err error) {
	//验证商铺，必须是总店，才可以删除充值卡规则
	_, err = checkBus(args.BsToken, true)
	if err != nil {
		return
	}
	if args.Id <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	deleRechar := new(models.RcardRuleModel).Init()
	rs := deleRechar.UpdateRelusId(args.Id, map[string]interface{}{
		deleRechar.Field.F_is_del: cards.RECHARGE_TYPE_YES,
	})
	if rs == false {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	return
}

// 获取充值规则列表
func (r *RcardLogic) BusRechargeRulesList(ctx context.Context, args *cards.ArgsRechargeRulesList, reply *cards.ReplyRechargerRulesList) (err error) {
	if args.RcardId <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	rcardModle := new(models.RcardModel).Init()
	rcardInfo := rcardModle.GetByRcardId(args.RcardId)
	discountType, _ := strconv.Atoi(rcardInfo[rcardModle.Field.F_discount_type].(string))
	discount, _ := strconv.ParseFloat(rcardInfo[rcardModle.Field.F_discount].(string), 64)

	// 如果类型是卡项折扣，返回折扣率，面值不需要折扣率
	if discountType == cards.DISCOUNT_TYPE_item {
		reply.DiscountType = cards.DISCOUNT_TYPE_item
		reply.Discount = discount

	} else {
		reply.DiscountType = cards.DISCOUNT_TYPE_price
		reply.Discount = 0

	}

	// 查找rcard_id为所有的充值卡规则。
	start := args.GetStart()
	limit := args.GetPageSize()
	listRecharModel := new(models.RcardRuleModel).Init()
	// 获取充值卡规则数据 显示未删除数据
	listRcards := listRecharModel.GetPageByRcardId(args.RcardId, start, limit, cards.RECHARGE_TYPE_NO)
	if len(listRcards) == 0 {
		return
	}
	for _, listRcard := range listRcards {
		rechargeAmountStr := listRcard[listRecharModel.Field.F_price].(string)
		donationAmountStr := listRcard[listRecharModel.Field.F_give_price].(string)
		rechargeAmount, _ := strconv.ParseFloat(rechargeAmountStr, 64)
		donationAmount, _ := strconv.ParseFloat(donationAmountStr, 64)
		id, _ := strconv.Atoi(listRcard[listRecharModel.Field.F_id].(string))
		name := fmt.Sprintf("充%s,赠送%s元", rechargeAmountStr, donationAmountStr)
		reply.Lists = append(reply.Lists, cards.ListsRechargeRules{
			Id:             id,
			RechargeAmount: rechargeAmount,
			DonationAmount: donationAmount,
			Name:           name,
		})
	}
	reply.TotalNum = listRecharModel.GetNumByRcardId(args.RcardId, cards.RECHARGE_TYPE_NO)
	return
}

//检查充值卡是否属于商家
func (r *RcardLogic) checkRcardids(busId int, rcardIds []int) (rcards []map[string]interface{}, err error) {
	mRcard := new(models.RcardModel).Init()
	rcards = mRcard.GetByRcardids(rcardIds)
	busIdStr := strconv.Itoa(busId)
	for _, rcard := range rcards {
		if rcard[mRcard.Field.F_bus_id].(string) != busIdStr {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
	}
	return
}

//检查充值卡的入参数据
func (r *RcardLogic) checkRcardData(ctx context.Context, busId int, rcardBase cards.RcardBase, isAllSingle, isAllProduct bool, incSingles []cards.IncSingle, giveSingles []cards.IncSingle,
	incProducts []cards.IncProduct) (err error) {
	//卡项折扣类型判断：现在只有面值折扣
	if rcardBase.DiscountType != cards.DISCOUNT_TYPE_price {
		return toolLib.CreateKcErr(_const.RCARD_DISCOUNT_TYPE_ERR)
	}
	if err = cards.VerfiyName(rcardBase.Name); err != nil {
		return
	}
	if err = cards.VerfiyPrice(rcardBase.RealPrice, rcardBase.Price); err != nil {
		return
	}
	if err = cards.VerfiyServicePeriod(rcardBase.ServicePeriod); err != nil {
		return
	}
	if (isAllSingle == false && len(incSingles) == 0) && (isAllProduct == false && len(incProducts) == 0) {
		err = toolLib.CreateKcErr(_const.RCARD_SINGLE_PRODUCT_NIL)
		return
	}
	if isAllSingle == false && len(incSingles) > 0 {
		if err = cards.VerifySinglesNum(len(incSingles)); err != nil {
			return
		}
	}
	if isAllProduct == false && len(incProducts) > 0 {
		if err = cards.VerifyProductsNum(len(incProducts)); err != nil {
			return
		}
	}
	if err = cards.VerifyGiveSinglesNum(len(giveSingles)); err != nil {
		return
	}
	if err = cards.VerifyRcardDiscount(rcardBase.DiscountType, rcardBase.Price, rcardBase.RealPrice, rcardBase.Discount); err != nil {
		return
	}

	//单项目id集合
	singleIds := []int{}
	if isAllSingle == false && len(incSingles) > 0 { //适用于部分单项目
		for _, single := range incSingles {
			singleIds = append(singleIds, single.SingleID)
		}
		//检查包含的项目里面是否有重复的
		incSingleNum := len(functions.ArrayUniqueInt(singleIds))
		if len(incSingles) != incSingleNum {
			err = toolLib.CreateKcErr(_const.SINGLE_REPEAT_ERR)
			return
		}
	}

	giveIds := []int{}
	if len(giveSingles) > 0 {
		for _, gsingle := range giveSingles {
			if gsingle.Num <= 0 {
				err = toolLib.CreateKcErr(_const.PARAM_ERR)
				return
			}
			if gsingle.PeriodOfValidity <= 0 {
				err = toolLib.CreateKcErr(_const.PERIOD_OF_VALIDITY_IS_NIL)
				return
			}
			singleIds = append(singleIds, gsingle.SingleID)
			giveIds = append(giveIds, gsingle.SingleID)
		}
	}
	//检查赠送的项目是否有重复
	if len(giveSingles) != len(functions.ArrayUniqueInt(giveIds)) {
		err = toolLib.CreateKcErr(_const.SINGLE_REPEAT_ERR)
		return
	}
	//检查单项目
	if len(singleIds) > 0 {
		if err = checkSingles(busId, singleIds); err != nil {
			return
		}
	}

	//检查产品
	if isAllProduct == false && len(incProducts) > 0 {
		productIds := []int{}
		for _, product := range incProducts {
			productIds = append(productIds, product.ProductID)
		}
		//检查包含的产品里面是否有重复的
		incProductNum := len(functions.ArrayUniqueInt(productIds))
		if len(incProducts) != incProductNum {
			err = toolLib.CreateKcErr(_const.PRODUCT_REPEAT_ERR)
			return
		}
		if err = checkProducts(ctx, busId, productIds); err != nil {
			return
		}
	}

	return
}

//获取充值卡基础数据
func (r *RcardLogic) GetRcardBaseInfo(ctx context.Context, args *cards.ArgsGetRcardBaseInfo, reply *cards.ReplyGetRcardBaseInfo) (err error) {
	mRcard := new(models.RcardModel).Init()
	rcardMaps := mRcard.GetByRcardids(args.RcardIds)
	if len(rcardMaps) == 0 {
		return
	}

	//获取充值卡规则
	rcardIds := functions.ArrayValue2Array(mRcard.Field.F_rcard_id, rcardMaps)
	rrm := new(models.RcardRuleModel).Init()
	rrmWhere := map[string]interface{}{rrm.Field.F_rcard_id: []interface{}{"IN", rcardIds}, rrm.Field.F_is_del: cards.RECHARGE_TYPE_NO}
	if args.RuleId > 0 {
		rrmWhere[rrm.Field.F_id] = args.RuleId
	}
	rrmMaps := rrm.Select(rrmWhere)
	if len(rrmMaps) > 0 {
		ruleMaps := make(map[string][]map[string]interface{})
		for _, rrmMap := range rrmMaps {
			rcardId := rrmMap[rrm.Field.F_rcard_id].(string)
			if _, ok := ruleMaps[rcardId]; !ok {
				ruleMaps[rcardId] = make([]map[string]interface{}, 0)
			}
			ruleMaps[rcardId] = append(ruleMaps[rcardId], rrmMap)
		}
		for index, rcardMap := range rcardMaps {
			rcardId := rcardMap[rrm.Field.F_rcard_id].(string)
			rcardMaps[index]["Rules"] = ruleMaps[rcardId]
		}
	}

	var rcardBase []cards.GetRcardBaseInfoBase
	_ = mapstructure.WeakDecode(rcardMaps, &rcardBase)
	reply.Lists = rcardBase
	return
}
