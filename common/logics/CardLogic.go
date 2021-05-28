//综合卡业务处理
//@author loop
//@date 2020/4/15 14:55
package logics

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"git.900sui.cn/kc/rpcCards/common/tools"
	"git.900sui.cn/kc/rpcCards/constkey"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/mapstructure"
	"git.900sui.cn/kc/redis"
	"git.900sui.cn/kc/rpcCards/common/models"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/client/bus"
	"git.900sui.cn/kc/rpcinterface/client/product"
	bus2 "git.900sui.cn/kc/rpcinterface/interface/bus"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	product2 "git.900sui.cn/kc/rpcinterface/interface/product"
	redis2 "github.com/gomodule/redigo/redis"
)

type CardLogic struct {
}

//添加Card描述
func (c *CardLogic) AddCardExt(cm *models.CardModel, cardID int, notes []cards.CardNote) (err error) {
	nem := new(models.CardExtModel).Init(cm.Model.GetOrmer())
	if len(notes) > constkey.CardNotesMaxNum {
		return toolLib.CreateKcErr(_const.NOTES_NUM_MAX_10)
	}
	for _, note := range notes {
		if functions.Mb4Strlen(note.Notes) > constkey.CardNotesSimpleMaxLength {
			return toolLib.CreateKcErr(_const.NOTES_LEN_MAX_30)
		}
	}
	notesStr, _ := json.Marshal(notes)
	if _, err = nem.Insert(map[string]interface{}{nem.Field.F_card_id: cardID, nem.Field.F_notes: string(notesStr)}); err != nil {
		err = toolLib.CreateKcErr(_const.DB_ERR)
	}
	return
}

//编辑Card描述
func (c *CardLogic) EditCardExt(cm *models.CardModel, cardID int, notes []cards.CardNote) (err error) {
	nem := new(models.CardExtModel).Init(cm.Model.GetOrmer())
	if len(notes) > constkey.CardNotesMaxNum {
		return toolLib.CreateKcErr(_const.NOTES_NUM_MAX_10)
	}
	for _, note := range notes {
		if functions.Mb4Strlen(note.Notes) > constkey.CardNotesSimpleMaxLength {
			return toolLib.CreateKcErr(_const.NOTES_LEN_MAX_30)
		}
	}
	notesStr, _ := json.Marshal(notes)
	smExMap := nem.Find(map[string]interface{}{nem.Field.F_card_id: cardID})
	if len(smExMap) > 0 {
		if _, updateErr := nem.Update(map[string]interface{}{nem.Field.F_card_id: cardID},
			map[string]interface{}{nem.Field.F_notes: string(notesStr)}); updateErr != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	} else {
		if _, err = nem.Insert(map[string]interface{}{nem.Field.F_card_id: cardID, nem.Field.F_notes: string(notesStr)}); err != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	return
}

//获取Card描述
func (c *CardLogic) GetCardExt(cardID int, notes *[]cards.CardNote) {
	nem := new(models.CardExtModel).Init()
	dataMap := nem.Find(map[string]interface{}{nem.Field.F_card_id: cardID})
	if len(dataMap) > 0 {
		data := dataMap[nem.Field.F_notes].(string)
		json.Unmarshal([]byte(data), notes)
	}
	return
}

//添加card
func (c *CardLogic) AddCard(ctx context.Context, busID int, args *cards.ArgsAddCard) (cardID int, err error) {
	//验证参数
	err = c.checkCardData(busID, args.CardBase, args.IsAllSingle, args.IsAllProduct, args.IncludeSingles, args.GiveSingles, args.IncludeProducts)
	if err != nil {
		return
	}

	if !args.IsAllProduct && len(args.IncludeProducts) > 0 { //包含部分商品
		//商品 id集合
		var productIDs []int
		for _, good := range args.IncludeProducts {
			productIDs = append(productIDs, good.ProductID)
		}
		if len(productIDs) > 0 {
			if err = checkProducts(ctx, busID, productIDs); err != nil {
				return
			}
		}
	}

	//验证图片
	var imgID int
	imgID, err = checkImg(ctx, args.ImgHash)
	if err != nil {
		return
	}
	var hasGive bool
	if len(args.GiveSingles) > 0 {
		hasGive = true
	}

	// 购买或者充值须100起
	if tools.RunMode == "prod" {
		if args.RealPrice < cards.BUY_CRARD_MIN_AMOUNT {
			err = toolLib.CreateKcErr(_const.BUY_CRARD_MIN_AMOUNT_ERR)
			return
		}
	}
	cm := new(models.CardModel).Init()
	csm := new(models.CardSingleModel).Init(cm.Model.GetOrmer())
	cpm := new(models.CardGoodsModel).Init(cm.Model.GetOrmer())

	// 卡项有效期 1-是；2-否；如果永久有效则service_period字段可忽略，否者必填
	ispermanentvalidity := cards.IS_PERMANENT_YES
	if args.ServicePeriod != 0 {
		ispermanentvalidity = cards.IS_PERMANENT_NO
	}

	//添加基本信息
	cm.Model.Begin()
	if cardID, err = cm.Insert(map[string]interface{}{
		cm.Field.F_bus_id:                busID,
		cm.Field.F_price:                 args.Price,
		cm.Field.F_is_ground:             cards.SINGLE_IS_GROUND_yes,
		cm.Field.F_real_price:            args.RealPrice,
		cm.Field.F_ctime:                 time.Now().Unix(),
		cm.Field.F_img_id:                imgID,
		cm.Field.F_name:                  args.Name,
		cm.Field.F_sort_desc:             args.ShortDesc,
		cm.Field.F_bind_id:               getBusMainBindId(ctx, busID),
		cm.Field.F_has_give_signle:       hasGive,
		cm.Field.F_service_period:        args.ServicePeriod,
		cm.Field.F_is_permanent_validity: ispermanentvalidity,
	}); err != nil {
		cm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//添加单项目
	if args.IsAllSingle { //走全部适用逻辑
		if err = csm.Insert(map[string]interface{}{
			csm.Field.F_card_id:   cardID,
			csm.Field.F_single_id: 0,
		}); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	} else if !args.IsAllSingle && len(args.IncludeSingles) > 0 { //走部分适用逻辑
		var cardSingleData []map[string]interface{}
		for _, single := range args.IncludeSingles {
			cardSingleData = append(cardSingleData, map[string]interface{}{
				csm.Field.F_single_id: single.SingleID,
				csm.Field.F_card_id:   cardID,
			})
		}
		if err = csm.InsertAll(cardSingleData); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//添加商品
	if args.IsAllProduct {
		if _, err = cpm.Insert(map[string]interface{}{
			cpm.Field.F_card_id:    cardID,
			cpm.Field.F_product_id: 0,
		}); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	} else if !args.IsAllProduct && len(args.IncludeProducts) > 0 {
		var cardProductData []map[string]interface{}
		for _, incProduct := range args.IncludeProducts {
			cardProductData = append(cardProductData, map[string]interface{}{
				cpm.Field.F_product_id: incProduct.ProductID,
				cpm.Field.F_card_id:    cardID,
			})
		}
		if _, err = cpm.InsertAll(cardProductData); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//添加赠送单项目数据
	if len(args.GiveSingles) > 0 {
		ngm := new(models.CardGiveModel).Init(cm.Model.GetOrmer())
		var giveSingleData []map[string]interface{}
		for _, single := range args.GiveSingles {
			giveSingleData = append(giveSingleData, map[string]interface{}{
				ngm.Field.F_single_id:          single.SingleID,
				ngm.Field.F_card_id:            cardID,
				ngm.Field.F_num:                single.Num,
				ngm.Field.F_period_of_validity: single.PeriodOfValidity,
			})
		}
		if err = ngm.InsertAll(giveSingleData); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//添加赠品描述
	if hasGive && len(args.GiveSingleDesc) > 0 {
		descm := new(models.CardGiveDescModel).Init(cm.Model.GetOrmer())
		giveSingleDescStr, _ := json.Marshal(args.GiveSingleDesc)
		descData := map[string]interface{}{
			descm.Field.F_card_id: cardID,
			descm.Field.F_desc:    string(giveSingleDescStr),
		}

		if _, err = descm.Model.Data(descData).Insert(); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//添加CardExt
	if err = c.AddCardExt(cm, cardID, args.Notes); err != nil {
		cm.Model.RollBack()
		return
	}
	cm.Model.Commit()
	//添加风控统计任务
	new(ItemLogic).AddXCardTask(ctx, cardID, cards.ITEM_TYPE_card)
	return
}

//编辑Card数据
func (c *CardLogic) EditCard(ctx context.Context, busID int, args *cards.ArgsEditCard) (err error) {

	cm := new(models.CardModel).Init()
	csm := new(models.CardSingleModel).Init(cm.Model.GetOrmer())
	cpm := new(models.CardGoodsModel).Init(cm.Model.GetOrmer())
	gdescm := new(models.CardGiveDescModel).Init(cm.Model.GetOrmer())

	//验证Card数据
	cardInfo := cm.GetByCardID(args.CardID, cm.Field.F_bus_id)
	if len(cardInfo) == 0 {
		err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
		return
	}
	if cardInfo[cm.Field.F_bus_id] != strconv.Itoa(busID) {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}

	//验证参数
	err = c.checkCardData(busID, args.CardBase, args.IsAllSingle, args.IsAllProduct, args.IncludeSingles, args.GiveSingles, args.IncludeProducts)
	if err != nil {
		return
	}
	if !args.IsAllProduct && len(args.IncludeProducts) > 0 { //包含部分商品
		//商品 id集合
		var productIDs []int
		for _, good := range args.IncludeProducts {
			productIDs = append(productIDs, good.ProductID)
		}
		if len(productIDs) > 0 {
			if err = checkProducts(ctx, busID, productIDs); err != nil {
				return
			}
		}
	}
	// 购买或者充值须100起
	if tools.RunMode == "prod" {
		if args.RealPrice < cards.BUY_CRARD_MIN_AMOUNT {
			err = toolLib.CreateKcErr(_const.BUY_CRARD_MIN_AMOUNT_ERR)
			return
		}
	}
	var imgID int
	imgID, err = checkImg(ctx, args.ImgHash)
	if err != nil {
		return
	}
	var hasGive bool
	if len(args.GiveSingles) > 0 {
		hasGive = true
	}

	//计算单项目改动
	includeSingles := csm.GetByCardID(args.CardID) //包含的单项目
	includeSingleIDs := functions.ArrayValue2Array(csm.Field.F_single_id, includeSingles)
	//需要新增的项目
	var addIncSingles []map[string]interface{}
	//需要删除的ids
	var delCardSingleIDs []int

	//获取需要添加的项目
	if args.IsAllSingle {
		addIncSingles = append(addIncSingles, map[string]interface{}{
			csm.Field.F_card_id:   args.CardID,
			csm.Field.F_single_id: 0,
		})
		args.IncludeSingles = []cards.IncInfSingle{}
	} else if !args.IsAllSingle && len(args.IncludeSingles) > 0 {
		for _, single := range args.IncludeSingles {
			if functions.InArray(single.SingleID, includeSingleIDs) {
				continue
			}
			addIncSingles = append(addIncSingles, map[string]interface{}{
				csm.Field.F_card_id:   args.CardID,
				csm.Field.F_single_id: single.SingleID,
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
		id, _ := strconv.Atoi(includeSingles[index][csm.Field.F_id].(string))
		delCardSingleIDs = append(delCardSingleIDs, id)
	}

	//计算商品改动
	includeProducts := cpm.GetByCardID(args.CardID)
	includeProductIDs := functions.ArrayValue2Array(cpm.Field.F_product_id, includeProducts)
	//需要新增的productIDs
	var addIncProducts []map[string]interface{}
	//需要删除的productIDs
	var delCardProductIDs []int

	//获取需要添加的product
	if args.IsAllProduct { //包含全部
		addIncProducts = append(addIncProducts, map[string]interface{}{
			cpm.Field.F_card_id:    args.CardID,
			cpm.Field.F_product_id: 0,
		})
		args.IncludeProducts = []cards.IncProduct{}
	} else if !args.IsAllProduct && len(args.IncludeProducts) > 0 {
		for _, includeProduct := range args.IncludeProducts {
			if functions.InArray(includeProduct.ProductID, includeProductIDs) {
				continue
			}
			addIncProducts = append(addIncProducts, map[string]interface{}{
				cpm.Field.F_card_id:    args.CardID,
				cpm.Field.F_product_id: includeProduct.ProductID,
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
		id, _ := strconv.Atoi(includeProducts[index][cpm.Field.F_id].(string))
		delCardProductIDs = append(delCardProductIDs, id)
	}

	//计算赠送项目改动
	ngm := new(models.CardGiveModel).Init(cm.Model.GetOrmer())
	giveSingles := ngm.GetByCardID(args.CardID)

	//需要新增的项目
	var addGiveSingles []map[string]interface{}
	//需要修改的单项目
	var updateGiveSingles = map[int]int{} //id:num
	//需要删除的ids
	var delCardGiveIDs []int
	//获取需要修改的项目以及需要添加的项目
out3:
	for _, single := range args.GiveSingles {
		for _, dbSingle := range giveSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[ngm.Field.F_single_id].(string) {
				//需要修改次数
				if strconv.Itoa(single.Num) != dbSingle[ngm.Field.F_num].(string) {
					dbID, _ := strconv.Atoi(dbSingle[ngm.Field.F_id].(string))
					updateGiveSingles[dbID] = single.Num
				}
				continue out3
			}
		}

		addGiveSingles = append(addGiveSingles, map[string]interface{}{
			ngm.Field.F_card_id:            args.CardID,
			ngm.Field.F_num:                single.Num,
			ngm.Field.F_single_id:          single.SingleID,
			ngm.Field.F_period_of_validity: single.PeriodOfValidity,
		})
	}
	//获取需要删除的id
out4:
	for _, dbSingle := range giveSingles {
		for _, single := range args.GiveSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[ngm.Field.F_single_id].(string) {
				continue out4
			}
		}
		sid, _ := strconv.Atoi(dbSingle[ngm.Field.F_id].(string))
		delCardGiveIDs = append(delCardGiveIDs, sid)
	}

	//修改主表信息
	cm.Model.Begin()
	if err = cm.UpdateByCardID(args.CardID, map[string]interface{}{
		cm.Field.F_price:                 args.Price,
		cm.Field.F_real_price:            args.RealPrice,
		cm.Field.F_img_id:                imgID,
		cm.Field.F_name:                  args.Name,
		cm.Field.F_sort_desc:             args.ShortDesc,
		cm.Field.F_has_give_signle:       hasGive,
		cm.Field.F_service_period:        args.ServicePeriod,
		cm.Field.F_is_permanent_validity: args.IsPermanentValidity,
	}); err != nil {
		cm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	//修改包含的单项
	//1. 添加需要添加的项目
	if len(addIncSingles) > 0 {
		if err = csm.InsertAll(addIncSingles); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//2. 删除需要删除的项目
	if len(delCardSingleIDs) > 0 {
		if err = csm.DelByIds(delCardSingleIDs); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//修改包含的商品
	//1. 添加需要添加的商品
	if len(addIncProducts) > 0 {
		if _, err = cpm.InsertAll(addIncProducts); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//2. 删除需要删除的商品
	if len(delCardProductIDs) > 0 {
		if err = cpm.DelByIds(delCardProductIDs); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//修改赠送的单项目
	//1. 修改需要修改的项目
	if len(updateGiveSingles) > 0 {
		for id, num := range updateGiveSingles {
			if err = ngm.UpdateNumById(id, num); err != nil {
				cm.Model.RollBack()
				err = toolLib.CreateKcErr(_const.DB_ERR)
				return
			}
		}
	}
	//2. 添加需要添加的项目
	if len(addGiveSingles) > 0 {
		if err = ngm.InsertAll(addGiveSingles); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//3. 删除需要删除的项目
	if len(delCardGiveIDs) > 0 {
		if err = ngm.DelByIds(delCardGiveIDs); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//修改赠品描述
	//1.删除原有赠品描述
	if err = gdescm.DelByCardId(args.CardID); err != nil {
		cm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//2.新增赠品描述
	if len(args.GiveSingles) > 0 && len(args.GiveSingleDesc) > 0 {
		giveSingleDescStr, _ := json.Marshal(args.GiveSingleDesc)
		descData := map[string]interface{}{
			gdescm.Field.F_card_id: args.CardID,
			gdescm.Field.F_desc:    string(giveSingleDescStr),
		}
		if _, err = gdescm.Model.Data(descData).Insert(); err != nil {
			cm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//编辑cardExt
	if err = c.EditCardExt(cm, args.CardID, args.Notes); err != nil {
		cm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	cm.Model.Commit()

	return
}

//验证商品是否属于商家
func checkProducts(ctx context.Context, busId int, productIDs []int) (err error) {
	client := new(product.Product).Init()
	defer client.Close()
	reply := new([]product2.ReplyProductGetByIds)
	if err = client.GetProductByIds(ctx, &product2.ArgsProductGetByIds{
		Ids: productIDs,
	}, reply); err != nil {
		return
	}
	if len(*reply) != len(productIDs) {
		err = toolLib.CreateKcErr(_const.PRODUCT_REPEAT_ERR)
		return
	}
	for _, good := range *reply {
		if good.BusId != busId {
			err = toolLib.CreateKcErr(_const.OPT_OTHER_BUS_ITEM)
			return
		}
	}

	return
}

//检查综合卡的入参数据
func (c *CardLogic) checkCardData(busID int, cardBase cards.CardBase, isAllSingle, isAllProduct bool, incSingles []cards.IncInfSingle,
	giveSingles []cards.IncSingle, incProducts []cards.IncProduct) (err error) {
	if err = cards.VerfiyName(cardBase.Name); err != nil {
		return
	}
	if err = cards.VerfiyPrice(cardBase.RealPrice, cardBase.Price); err != nil {
		return
	}
	if err = cards.VerfiyServicePeriod(cardBase.ServicePeriod); err != nil {
		return
	}
	if (isAllSingle == false && len(incSingles) == 0) && (isAllProduct == false && len(incProducts) == 0) {
		err = toolLib.CreateKcErr(_const.CARD_SINGLE_PRODUCT_NIL)
		return
	}
	if isAllSingle == false && len(incSingles) > 0 {
		if err = cards.VerifySinglesNum(len(incSingles)); err != nil {
			return
		}
	}
	if err = cards.VerifyGiveSinglesNum(len(giveSingles)); err != nil {
		return
	}
	if isAllProduct == false && len(incProducts) > 0 {
		if err = cards.VerifyProductsNum(len(incProducts)); err != nil {
			return
		}
	}

	if isAllSingle == false && len(incSingles) > 0 { //包含部分单项目
		//单项目id集合
		var singleIDs []int
		for _, single := range incSingles {
			singleIDs = append(singleIDs, single.SingleID)
		}
		if err = checkSingles(busID, singleIDs); err != nil {
			return
		}
	}

	if len(giveSingles) > 0 {
		var giveIDs []int
		for _, giveSingle := range giveSingles {
			if giveSingle.Num <= 0 {
				err = toolLib.CreateKcErr(_const.PARAM_ERR)
				return
			}
			if giveSingle.PeriodOfValidity <= 0 {
				err = toolLib.CreateKcErr(_const.PERIOD_OF_VALIDITY_IS_NIL)
				return
			}
			giveIDs = append(giveIDs, giveSingle.SingleID)
		}
		if err = checkSingles(busID, giveIDs); err != nil {
			return
		}
	}

	return
}

//获取包含的单项目
func getIncInfSingles(ctx context.Context, singleIds []int) (reply map[int]cards.IncInfSingleDetail) {
	reply = make(map[int]cards.IncInfSingleDetail, len(singleIds))
	mSingle := new(models.SingleModel).Init()
	singles := mSingle.GetBySingleids(singleIds, []string{
		mSingle.Field.F_single_id,
		mSingle.Field.F_name,
		mSingle.Field.F_price,
		mSingle.Field.F_real_price,
		mSingle.Field.F_min_price,
		mSingle.Field.F_max_price,
		mSingle.Field.F_img_id,
		mSingle.Field.F_service_time,
	})
	//获取图片
	imgIds := functions.ArrayValue2Array(mSingle.Field.F_img_id, singles)
	imgs := getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_single)
	var singlesStruct []struct {
		SingleId    int
		Name        string
		Price       string
		RealPrice   string
		MinPrice    string
		MaxPrice    string
		ImgId       int
		ServiceTime int
	}
	_ = mapstructure.WeakDecode(singles, &singlesStruct)
	for _, single := range singlesStruct {
		realPrice := single.RealPrice
		if single.MinPrice != "0" || single.MaxPrice != "0" {
			realPrice = fmt.Sprintf("%s-%s", single.MinPrice, single.MaxPrice)
		}
		reply[single.SingleId] = cards.IncInfSingleDetail{
			IncInfSingle: cards.IncInfSingle{
				SingleID: single.SingleId,
			},
			Name:        single.Name,
			Price:       single.Price,
			RealPrice:   realPrice,
			ImgId:       single.ImgId,
			ImgUrl:      imgs[single.ImgId],
			ServiceTime: single.ServiceTime,
		}
	}
	return
}

//获取包含的单项目
func getIncProducts(ctx context.Context, productIDs []int) (result map[int]cards.IncProductDetail, err error) {

	result = make(map[int]cards.IncProductDetail, len(productIDs))
	client := new(product.Product).Init()
	defer client.Close()

	if len(productIDs) > 0 {
		reply := new([]product2.ReplyProductGetByIds)
		if err = client.GetProductByIds(ctx, &product2.ArgsProductGetByIds{
			Ids: productIDs,
		}, reply); err != nil {
			return
		}
		imgIds := []int{}
		for _, productDesc := range *reply {
			result[productDesc.Id] = cards.IncProductDetail{
				IncProduct: cards.IncProduct{ProductID: productDesc.Id},
				Name:       productDesc.Name,
				SpecPrice:  fmt.Sprintf("%.2f-%.2f", productDesc.MinPrice, productDesc.MaxPrice),
				ImgId:      productDesc.ImgId,
				ImgUrl:     "",
			}
			imgIds = append(imgIds, productDesc.ImgId)
		}
		//获取图片
		imgs := getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_single)
		for k, v := range result {
			v.ImgUrl = imgs[v.ImgId]
			result[k] = v
		}
	}

	return
}

//获取包含的商品
func getIncProducts2(ctx context.Context, productIDs []int) (result map[int]cards.IncProductDetail2, err error) {

	result = make(map[int]cards.IncProductDetail2, len(productIDs))
	client := new(product.Product).Init()
	defer client.Close()

	if len(productIDs) > 0 {
		reply := new([]product2.ReplyProductGetByIds)
		if err = client.GetProductByIds(ctx, &product2.ArgsProductGetByIds{
			Ids: productIDs,
		}, reply); err != nil {
			return
		}
		imgIds := []int{}
		for _, productDesc := range *reply {
			result[productDesc.Id] = cards.IncProductDetail2{
				IncProduct: cards.IncProduct{ProductID: productDesc.Id},
				Name:       productDesc.Name,
				SpecPrice:  fmt.Sprintf("%.2f-%.2f", productDesc.MinPrice, productDesc.MaxPrice),
				ImgId:      productDesc.ImgId,
				ImgUrl:     "",
			}
			imgIds = append(imgIds, productDesc.ImgId)
		}
		//获取图片
		imgs := getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_single)
		for k, v := range result {
			v.ImgUrl = imgs[v.ImgId]
			result[k] = v
		}
	}

	return
}

//获取综合卡的详情
//@param int cardID 综合卡id
//@param int shopID 门店id
//@return  cards.ReplyCardInfo reply
func (c *CardLogic) CardInfo(ctx context.Context, cardID int, shopID ...int) (reply cards.ReplyCardInfo, err error) {
	reply.IncludeSingles = make([]cards.IncInfSingleDetail, 0)
	reply.IncProducts = make([]cards.IncProductDetail, 0)
	reply.GiveSingles = make([]cards.IncSingleDetail, 0)
	reply.GiveSingleDesc = make([]cards.GiveSingleDesc, 0)
	ncm := new(models.CardModel).Init()
	cardInfo := ncm.GetByCardID(cardID)
	if len(cardInfo) == 0 {
		err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
		return
	}
	var shopStatus = 0
	var shopSales = 0
	var shopid = 0
	if len(shopID) > 0 && shopID[0] > 0 {
		snm := new(models.ShopCardModel).Init()
		shopNcInfo := snm.GetByShopIDAndCardIDs(shopID[0], []int{cardID})
		if len(shopNcInfo) == 0 {
			err = toolLib.CreateKcErr(_const.NO_IN_SHOP)
			return
		}
		shopStatus, _ = strconv.Atoi(shopNcInfo[0][snm.Field.F_status].(string))
		reply.SsId, _ = strconv.Atoi(shopNcInfo[0][snm.Field.F_id].(string))
		shopSales, _ = strconv.Atoi(shopNcInfo[0][snm.Field.F_sales].(string))
		shopid = shopID[0]
	}
	imgId, _ := strconv.Atoi(cardInfo[ncm.Field.F_img_id].(string))
	imgHash, imgUrl := getImg(ctx, imgId, cards.ITEM_TYPE_card)
	reply.ImgHash = imgHash
	reply.ImgUrl = imgUrl
	reply.ShopStatus = shopStatus
	_ = mapstructure.WeakDecode(cardInfo, &reply.CardBase)
	reply.CtimeStr = time.Unix(int64(reply.Ctime), 0).Format("2006/01/02 15:04:05")
	_ = mapstructure.WeakDecode(cardInfo, &reply)
	if len(shopID) > 0 && shopID[0] > 0 {
		reply.Sales = shopSales
	}

	//商户信息
	if err = getBusInfo(ctx, reply.BusID, &reply.BusInfo); err != nil {
		err = toolLib.CreateKcErr(_const.SHOP_INFO_ERR)
		return
	}
	//.....
	//获取包含的项目
	csm := new(models.CardSingleModel).Init()
	cardSingles := csm.Find(map[string]interface{}{csm.Field.F_card_id: cardID})
	if len(cardSingles) > 0 && cardSingles[csm.Field.F_single_id].(string) == "0" { //包含全部
		reply.IsAllSingle = true
	}

	//获取包含的商品
	cpm := new(models.CardGoodsModel).Init()
	cardProducts := cpm.Find(map[string]interface{}{cpm.Field.F_card_id: cardID})
	if len(cardProducts) > 0 {
		if cardProducts[cpm.Field.F_product_id].(string) == "0" { //包含全部
			reply.IsAllProduct = true
		}
	}

	//获取包含的赠送项目信息
	var allGiveSingles map[int]cards.IncSingleDetail
	cgm := new(models.CardGiveModel).Init()
	var cardGives []map[string]interface{}
	if cardInfo[ncm.Field.F_has_give_signle].(string) == strconv.Itoa(cards.HAS_GIVE_SINGLE_yes) {
		cardGives = cgm.GetByCardID(cardID)
		allGiveSingles = getIncSingles(ctx, functions.ArrayValue2Array(cgm.Field.F_single_id, cardGives))
	}
	if len(cardGives) > 0 {
		for _, single := range cardGives {
			singleId, _ := strconv.Atoi(single[cgm.Field.F_single_id].(string))
			sInfo := allGiveSingles[singleId]
			sInfo.Num, _ = strconv.Atoi(single[cgm.Field.F_num].(string))
			sInfo.PeriodOfValidity, _ = strconv.Atoi(single[cgm.Field.F_period_of_validity].(string))
			reply.GiveSingles = append(reply.GiveSingles, sInfo)
		}

		//获取赠品描述信息
		gdescm := new(models.CardGiveDescModel).Init()
		desc, ok := gdescm.GetByCardId(cardID)[gdescm.Field.F_desc].(string)
		if ok {
			json.Unmarshal([]byte(desc), &reply.GiveSingleDesc)
		} else {
			reply.GiveSingleDesc = []cards.GiveSingleDesc{}
		}
	}
	//获取CardExt
	c.GetCardExt(cardID, &reply.Notes)
	reply.ShareLink = tools.GetShareLink(cardID, shopid, cards.ITEM_TYPE_card)

	// 获取限时卡门店添加详情  15 -- []int {3,4,6}
	busId, _ := strconv.Atoi(cardInfo[ncm.Field.F_bus_id].(string))
	cardModel := new(models.CardShopModel).Init()
	cardLists := cardModel.GetByCardIdAndBusId(cardID, busId)

	cardShopIds := make([]int, 0)
	for _, cardInfoValue := range cardLists {
		sshopId, _ := strconv.Atoi(cardInfoValue[cardModel.Field.F_shop_id].(string))
		cardShopIds = append(cardShopIds, sshopId)
	}
	var replyShop []bus2.ReplyShopName
	rLists := make([]cards.ReplyShopName, 0)
	rpcBus := new(bus.Shop).Init()
	defer rpcBus.Close()
	err = rpcBus.GetShopNameByShopIds(ctx, &cardShopIds, &replyShop)
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
	_ = redis.RedisGlobMgr.Hincrby(constkey.CARD_CLIKS, strconv.Itoa(cardID), 1)
	return
}

//获取商家的综合卡列表
func (c *CardLogic) GetBusPage(ctx context.Context, busId, shopId, start, limit int, isGround string, filterShopHasAdd bool) (list cards.ReplyCardPage, err error) {

	if busId <= 0 || start < 0 || limit < 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	ncm := new(models.CardModel).Init()
	cardInfos := make([]map[string]interface{}, 0)
	list.List = make([]cards.CardDesc,0)
	list.IndexImg = make(map[int]string)

	// 获取门店已添加的综合卡
	var shopAddCards []map[string]interface{}
	scModel := new(models.ShopCardModel).Init()
	where := make([]base.WhereItem, 0)
	where = append(where, base.WhereItem{ncm.Field.F_bus_id, busId})
	where = append(where, base.WhereItem{ncm.Field.F_is_del, cards.IS_BUS_DEL_no})

	if shopId > 0 {
		shopAddCards = scModel.SelectRcardsByWherePage([]base.WhereItem{{scModel.Field.F_shop_id, shopId},{scModel.Field.F_is_del, cards.IS_BUS_DEL_no}}, 0, 0)
		if filterShopHasAdd && len(shopAddCards) > 0 {
			shopHasAddRcardIds := functions.ArrayValue2Array(scModel.Field.F_card_id, shopAddCards)
			where = append(where, base.WhereItem{ncm.Field.F_card_id, []interface{}{"NOT IN", shopHasAddRcardIds}})
		}
	}

	//获取总数量
	//if isGround == "" {
	cardInfos = ncm.SelectCardsByWherePage(where, start, limit)
	list.TotalNum = ncm.GetNumByWhere(where)
	//} else {
	//	isground, _ := strconv.Atoi(isGround)
	//	isground = isground - 1
	//	cardInfos = ncm.SelectCardsByWherePage(where,start,limit,isGround)
	//	list.TotalNum = ncm.GetNumByWhere(where)
	//}

	if len(cardInfos) == 0 {
		return
	}
	list.List = make([]cards.CardDesc, len(cardInfos))
	for index, nc := range cardInfos {
		_ = mapstructure.WeakDecode(nc, &list.List[index].CardBase)
		list.List[index].CtimeStr = time.Unix(int64(list.List[index].Ctime), 0).Format("2006/01/02 15:04:05")
		for _, shopCard := range shopAddCards {
			if nc[ncm.Field.F_card_id].(string) == shopCard[scModel.Field.F_card_id].(string) { //表明当前子店已添加该卡项
				nc["ShopHasAdd"] = 1
				nc["Status"] = shopCard[scModel.Field.F_status].(string)
				nc["ShopItemId"] = shopCard[scModel.Field.F_id].(string)
				nc["ShopDelStatus"] = shopCard[scModel.Field.F_is_del].(string)
				break
			}
		}
		_ = mapstructure.WeakDecode(nc, &list.List[index])
		list.List[index].Clicks, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.CARD_CLIKS, nc[ncm.Field.F_card_id].(string)))
	}

	//获取图片信息
	imgIds := functions.ArrayValue2Array(ncm.Field.F_img_id, cardInfos)
	list.IndexImg = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_card)

	return
}

//设置综合卡的适用门店（废用）
func (c *CardLogic) SetCardShop(ctx context.Context, busId int, args *cards.ArgsSetCardShop) (err error) {
	if len(args.CardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.CARD_ID_IS_NIL)
		return
	}
	if args.IsAllShop == false && len(args.ShopIDs) == 0 {
		err = toolLib.CreateKcErr(_const.SHOPID_NTL)
		return
	}

	//限次卡id重复提交判断
	realCardIDs := functions.ArrayUniqueInt(args.CardIDs)
	if len(realCardIDs) != len(args.CardIDs) {
		err = toolLib.CreateKcErr(_const.DATA_REPEAT_ERR)
		return
	}

	//店铺id重复提交判断
	realShopIds := functions.ArrayUniqueInt(args.ShopIDs)
	if args.IsAllShop == false && len(realShopIds) != len(args.ShopIDs) {
		err = toolLib.CreateKcErr(_const.DATA_REPEAT_ERR)
		return
	}

	//检查限次卡
	//1.是否属于企业
	//2.是否上架
	hcm := new(models.CardModel).Init()
	dataArr := hcm.GetByCardIDs(realCardIDs)
	if len(dataArr) != len(realCardIDs) {
		err = toolLib.CreateKcErr(_const.PARAM_ERR, "has invalid cardID")
		return
	}
	busIdStr := strconv.Itoa(busId)
	for _, data := range dataArr {
		if busIdStr != data[hcm.Field.F_bus_id].(string) {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
	}

	hsm := new(models.CardShopModel).Init()
	var insertData []map[string]interface{}
	//走全部适用逻辑
	if args.IsAllShop == true {
		for _, hCardID := range realCardIDs {
			insertData = append(insertData, map[string]interface{}{
				hsm.Field.F_card_id: hCardID,
				hsm.Field.F_bus_id:  busId,
				hsm.Field.F_shop_id: 0,
			})
		}
	} else {
		//部分门店适用
		//检查门店id是否合法
		rpcShop := new(bus.Shop).Init()
		defer rpcShop.Close()
		checkArgs := &bus2.ArgsCheckShop{
			BusId:   busId,
			ShopIds: realShopIds,
		}
		checkReply := &bus2.ReplyCheckShop{}
		if err = rpcShop.CheckBusShop(ctx, checkArgs, checkReply); err != nil {
			return
		}
		for _, hCardID := range realCardIDs {
			for _, shopID := range realShopIds {
				insertData = append(insertData, map[string]interface{}{
					hsm.Field.F_card_id: hCardID,
					hsm.Field.F_bus_id:  busId,
					hsm.Field.F_shop_id: shopID,
				})
			}
		}
	}

	//处理规格
	if len(insertData) > 0 {
		hsm.Model.Begin()
		if err = hsm.DelByCardIDs(realCardIDs); err != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			hsm.Model.RollBack()
			return
		}
		if err = hsm.InsertAll(insertData); err != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			hsm.Model.RollBack()
			return
		}
		hsm.Model.Commit()
	}

	var disableIds []int
	var downIds []int

	shm := new(models.ShopCardModel).Init()
	shopCards := shm.GetByCardIDs(realCardIDs)
	if args.IsAllShop {
		//2.当总店设置适用门店包含下架的卡项时需要使用下面的逻辑当总店设置适用门店包含下架的卡项时需要使用下面的逻辑
		for _, hcard := range dataArr {
			hcardStatus, _ := strconv.Atoi(hcard[hcm.Field.F_is_ground].(string)) // 总店中卡的状态
			hcardID, _ := strconv.Atoi(hcard[hcm.Field.F_card_id].(string))
			for _, shopHcard := range shopCards {
				shopStatus, _ := strconv.Atoi(shopHcard[shm.Field.F_status].(string)) // 子店中卡的状态
				id, _ := strconv.Atoi(shopHcard[shm.Field.F_id].(string))
				shopHcardID, _ := strconv.Atoi(shopHcard[shm.Field.F_card_id].(string)) // 子店中卡的id
				if hcardID != shopHcardID {
					continue
				}
				if shopStatus == cards.STATUS_DISABLE && hcardStatus == cards.SINGLE_IS_GROUND_yes { // 子店中卡的状态为禁用并且总店卡的状态为上架时,子店才可以上架
					downIds = append(downIds, id)
				}
			}
		}
	} else {
		// 1.当总店设置适用门店包含下架的卡项时需要使用下面的逻辑当总店设置适用门店包含下架的卡项时需要使用下面的逻辑
		for _, hcard := range dataArr {
			hcardStatus, _ := strconv.Atoi(hcard[hcm.Field.F_is_ground].(string)) // 总店中卡的状态
			hcardID, _ := strconv.Atoi(hcard[hcm.Field.F_card_id].(string))
			for _, shopHcard := range shopCards {
				shopStatus, _ := strconv.Atoi(shopHcard[shm.Field.F_status].(string)) // 子店中卡的状态
				id, _ := strconv.Atoi(shopHcard[shm.Field.F_id].(string))
				shopID, _ := strconv.Atoi(shopHcard[shm.Field.F_shop_id].(string))      // 已经存在的子店id
				shopHcardID, _ := strconv.Atoi(shopHcard[shm.Field.F_card_id].(string)) // 子店中卡的id
				if hcardID != shopHcardID {
					continue
				}
				if functions.InArray(shopID, realShopIds) { // 以前已经添加过的子店并且status="总店禁用";现在适用时需要将status恢复为下架;前提为总店卡的状态为上架时
					if hcardStatus == cards.SINGLE_IS_GROUND_yes && shopStatus == cards.STATUS_DISABLE {
						downIds = append(downIds, id)
					}
				} else { // 以前已经添加过的子店并且status为"上架"或"下架";如果现在不再适用现时卡时，status需更改为"总店禁用"
					if shopStatus != cards.STATUS_DISABLE {
						disableIds = append(disableIds, id)
					}
				}
			}
		}
	}

	if len(disableIds) > 0 {
		_ = shm.UpdateByIDs(disableIds, map[string]interface{}{
			shm.Field.F_status:     cards.STATUS_DISABLE,
			shm.Field.F_under_time: time.Now().Unix(),
		})
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, []int{}, 0, cards.ITEM_TYPE_card, disableIds)
	}
	if len(downIds) > 0 {
		_ = shm.UpdateByIDs(downIds, map[string]interface{}{
			shm.Field.F_status: cards.STATUS_OFF_SALE,
		})
	}
	return
}

//总店上下架综合卡
func (c *CardLogic) DownUpCard(ctx context.Context, busId int, args *cards.ArgsDownUpCard) (err error) {
	// 综合卡ID列表交验
	if len(args.CardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	realCardIDs := functions.ArrayUniqueInt(args.CardIDs)
	//交验是否有重复的id
	if len(realCardIDs) != len(args.CardIDs) {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//检查综合卡是否属于企业
	ncm := new(models.CardModel).Init()
	dataArr := ncm.GetByCardIDs(realCardIDs, ncm.Field.F_bus_id, ncm.Field.F_card_id, ncm.Field.F_is_ground)
	busIdStr := strconv.Itoa(busId)
	for _, data := range dataArr {
		if busIdStr != data[ncm.Field.F_bus_id].(string) {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
	}
	//将id中已经上架的综合卡id和未上架的综合卡id分开处理
	var cardDescList []struct {
		CardID   int `mapstructure:"card_id"`
		IsGround int `mapstructure:"is_ground"`
	}
	_ = mapstructure.WeakDecode(dataArr, &cardDescList)
	var downIds, upIds []int
	for _, cardDesc := range cardDescList {
		if cardDesc.IsGround == cards.IS_GROUND_no {
			downIds = append(downIds, cardDesc.CardID)
		} else {
			upIds = append(upIds, cardDesc.CardID)
		}
	}

	snm := new(models.ShopCardModel).Init(ncm.Model.GetOrmer())
	//下架操作, 只处理已经上架的综合卡id
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		ncm.Model.Begin()
		if err = ncm.UpdateByCardIDs(upIds, map[string]interface{}{
			ncm.Field.F_is_ground:     cards.IS_GROUND_no,
			ncm.Field.F_under_time:    time.Now().Unix(),
			ncm.Field.F_sale_shop_num: 0,
		}); err != nil {
			snm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//将分店的综合卡设置为总店禁用
		if err = snm.UpdateByCardIDs(upIds, map[string]interface{}{
			snm.Field.F_status:     cards.STATUS_DISABLE,
			snm.Field.F_under_time: time.Now().Unix(),
		}); err != nil {
			snm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		ncm.Model.Commit()
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, upIds, 0, cards.ITEM_TYPE_card)
	}

	//上架操作, 只处理未上架的综合卡id
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		ncm.Model.Begin()
		if err = ncm.UpdateByCardIDs(downIds, map[string]interface{}{
			ncm.Field.F_is_ground:  cards.IS_GROUND_yes,
			ncm.Field.F_under_time: 0,
		}); err != nil {
			snm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//2.解除总店禁用状态
		if err = snm.UpdateByCardIDs(downIds, map[string]interface{}{
			snm.Field.F_status: cards.STATUS_OFF_SALE,
		}); err != nil {
			snm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		ncm.Model.Commit()
	}
	return
}

//子店获取适用本店的综合卡列表
func (c *CardLogic) ShopGetBusCardPage(ctx context.Context, busId, shopId, start, limit int) (list cards.ReplyCardPage, err error) {
	if busId <= 0 || shopId < 0 || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	nsm := new(models.CardShopModel).Init()
	cardShops := nsm.GetPageByShopID(busId, shopId, start, limit)
	if len(cardShops) == 0 {
		return
	}
	//获取综合卡基本信息
	ncm := new(models.CardModel).Init()
	snm := new(models.ShopCardModel).Init()

	cardIds := functions.ArrayValue2Array(nsm.Field.F_card_id, cardShops)
	cardInfos := ncm.GetByCardIDs(cardIds)
	if len(cardInfos) == 0 {
		return
	}

	list.TotalNum = nsm.GetNumByShopID(busId, shopId)
	list.List = make([]cards.CardDesc, len(cardInfos))
	// 店面已添加综合卡列表
	shopCards := snm.GetByShopIDAndCardIDs(shopId, cardIds)
	shopCardIds := functions.ArrayValue2Array(snm.Field.F_card_id, shopCards)
	for index, card := range cardInfos {
		cardId, _ := strconv.Atoi(card[ncm.Field.F_card_id].(string))
		_ = mapstructure.WeakDecode(card, &list.List[index].CardBase)
		list.List[index].CtimeStr = time.Unix(int64(list.List[index].Ctime), 0).Format("2006/01/02 15:04:05")
		_ = mapstructure.WeakDecode(card, &list.List[index])
		list.List[index].Clicks, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.CARD_CLIKS, card[ncm.Field.F_card_id].(string)))
		shopHasAdd := 0
		if functions.InArray(cardId, shopCardIds) {
			shopHasAdd = 1
		}
		list.List[index].ShopHasAdd = shopHasAdd
		for _, shopCard := range shopCards {
			cardID, _ := strconv.ParseInt(shopCard[snm.Field.F_card_id].(string), 10, 64)
			if list.List[index].CardID == int(cardID) {
				list.List[index].ShopHasAdd = 1
				status, _ := strconv.ParseInt(shopCard[snm.Field.F_status].(string), 10, 64)
				list.List[index].ShopStatus = int(status)
			}
		}
	}
	//获取图片信息
	imgIds := functions.ArrayValue2Array(ncm.Field.F_img_id, cardInfos)
	list.IndexImg = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_card)

	return
}

//子店添加综合卡到自己的门店
func (c *CardLogic) ShopAddCard(ctx context.Context, busId, shopId int, args *cards.ArgsShopAddCard) (err error) {
	args.CardIDs = functions.ArrayUniqueInt(args.CardIDs)
	if len(args.CardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	//检查综合卡id是否适用当前门店
	//nsm := new(models.CardShopModel).Init()
	//cardShop := nsm.GetByShopIDAndCardIDs(busId, shopId, args.CardIDs)
	//if len(cardShop) != len(args.CardIDs) {
	//	err = toolLib.CreateKcErr(_const.NO_IN_SHOP)
	//	return
	//}

	// 提取门店已经添加过的卡项
	snm := new(models.ShopCardModel).Init()
	shopCardLists := snm.GetByShopByCardIDs(shopId,args.CardIDs)
	shopCardIds := functions.ArrayValue2Array(snm.Field.F_card_id, shopCardLists)

	// 刷选出已经添加过并且删除的数据
	delCardIdSlice := make([]int, 0)
	for _, hcardMap := range shopCardLists {
		isDel, _ := strconv.Atoi(hcardMap[snm.Field.F_is_del].(string))
		if isDel == cards.IS_BUS_DEL_yes {
			delCardId, _ := strconv.Atoi(hcardMap[snm.Field.F_card_id].(string))
			delCardIdSlice = append(delCardIdSlice, delCardId)
		}
	}

	// 更新门店之前添加过并删除的数据
	if len(delCardIdSlice) > 0 {
		if updateBool := snm.UpdateDelByHcardIdsAndshopId(delCardIdSlice, shopId, map[string]interface{}{
			snm.Field.F_is_del: cards.IS_BUS_DEL_no,
			snm.Field.F_status: cards.STATUS_OFF_SALE,
		}); !updateBool {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
		//更新卡项关联表
		sirModel := new(models.ShopItemRelationModel).Init()
		if b := sirModel.UpdateByItemIdsAndShopId(delCardIdSlice, cards.ITEM_TYPE_card, shopId, map[string]interface{}{
			sirModel.Field.F_is_del: cards.IS_BUS_DEL_no,
			sirModel.Field.F_status: cards.STATUS_OFF_SALE,
		}); !b {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}

	//需要添加的CardID列表
	addCardIds := make([]int, 0)
	for _, cardID := range args.CardIDs {
		if functions.InArray(cardID, shopCardIds) == false {
			addCardIds = append(addCardIds, cardID)
		}
	}

	//判断本店是否存在包含商品
	err, checkSuccess := new(ItemLogic).CheckProductInShop(ctx, busId, shopId, cards.ITEM_TYPE_card, addCardIds)
	if err != nil {
		return err
	}
	if !checkSuccess {
		return toolLib.CreateKcErr(_const.SHOP_PRODUCT_NOT_CONTAIN_BUS_PRODUCT)
	}
	//校验当前门店是否已经将卡项内涉及到的单项目添加到自己的门店内
	allSingle, singleIds, err := new(ItemLogic).getItemCardIncSingleIds(addCardIds, cards.ITEM_TYPE_card)
	if err != nil {
		return
	}
	if err = new(ItemLogic).validShopSingleContainItemCardSingles(shopId, busId, allSingle, singleIds); err != nil {
		return
	}

	var addData []map[string]interface{}
	shopItemRelationData := make([]map[string]interface{}, 0)
	shopItemRelationModel := new(models.ShopItemRelationModel).Init()
	for _, cardId := range addCardIds {
		status := cards.STATUS_OFF_SALE
		ctime := time.Now().Local().Unix()
		addData = append(addData, map[string]interface{}{
			snm.Field.F_card_id: cardId,
			snm.Field.F_status:  status,
			snm.Field.F_shop_id: shopId,
			snm.Field.F_ctime:   ctime,
		})
		shopItemRelationData = append(shopItemRelationData, map[string]interface{}{
			shopItemRelationModel.Field.F_item_id:   cardId,
			shopItemRelationModel.Field.F_item_type: cards.ITEM_TYPE_card,
			shopItemRelationModel.Field.F_status:    cards.STATUS_OFF_SALE,
			shopItemRelationModel.Field.F_shop_id:   shopId,
			shopItemRelationModel.Field.F_is_del:    cards.ITEM_IS_DEL_NO,
		})
	}
	//过滤的数据添加到门店综合卡表
	if len(addData) > 0 {
		id, _ := snm.InsertAll(addData)
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
	cardShopModel := new(models.CardShopModel).Init()
	smshopSms := cardShopModel.GetByCardIDs(args.CardIDs)
	cardshopIds := functions.ArrayValue2Array(cardShopModel.Field.F_card_id, smshopSms)
	var addCardShopIds = []int{}
	for _, cardId := range args.CardIDs {
		if functions.InArray(cardId, cardshopIds) == false {
			addCardShopIds = append(addCardShopIds, cardId)
		}
	}
	if len(addCardShopIds) == 0 {
		return
	}

	var addCardShopData []map[string]interface{} // 添加适用门店表的数据
	for _, cardId := range addCardShopIds {
		addCardShopData = append(addCardShopData, map[string]interface{}{
			cardShopModel.Field.F_card_id: cardId,
			cardShopModel.Field.F_shop_id: shopId,
			cardShopModel.Field.F_bus_id:  busId,
		})
	}
	// 过滤的数据添加到适用综合卡表
	if len(addCardShopData) > 0 {
		err = cardShopModel.InsertAll(addCardShopData)
		if err != nil {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}

	//获取综合卡在总店的上下架状态
	//ncm := new(models.CardModel).Init()
	//cardInfos := ncm.GetByCardIDs(addCardIds, ncm.Field.F_card_id, ncm.Field.F_is_ground)
	//cardMap := functions.ArrayRebuild(ncm.Field.F_card_id, cardInfos)
	//var addData []map[string]interface{}
	//for _, cardID := range addCardIds {
	//	status := cards.STATUS_OFF_SALE
	//	if _, ok := cardMap[strconv.Itoa(cardID)]; ok {
	//		if card, ok := cardMap[strconv.Itoa(cardID)].(map[string]interface{}); ok {
	//			isGround, _ := strconv.Atoi(card[ncm.Field.F_is_ground].(string))
	//			if isGround == cards.IS_GROUND_no {
	//				status = cards.STATUS_DISABLE
	//			}
	//		}
	//	}
	//	ctime := time.Now().Local().Unix()
	//	addData = append(addData, map[string]interface{}{
	//		snm.Field.F_card_id: cardID,
	//		snm.Field.F_status:  status,
	//		snm.Field.F_shop_id: shopId,
	//		snm.Field.F_ctime:   ctime,
	//	})
	//}

	return
}

//获取子店的综合卡列表
func (c *CardLogic) ShopCardPage(ctx context.Context, shopId, start, limit int, status int) (list cards.ReplyCardPage, err error) {
	if shopId <= 0 || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	list.List = []cards.CardDesc{}
	list.IndexImg = make(map[int]string)
	//获取门店的综合卡数据
	snm := new(models.ShopCardModel).Init()
	shopCards := snm.GetPageByShopID(shopId, start, limit, status)
	if len(shopCards) <= 0 {
		return
	}
	cardIDs := functions.ArrayValue2Array(snm.Field.F_card_id, shopCards)
	//获取综合卡基本信息
	ncm := new(models.CardModel).Init()
	cardInfos := ncm.GetByCardIDs(cardIDs)
	if len(cardInfos) == 0 {
		return
	}
	var cardArr []cards.CardBase
	mapstructure.WeakDecode(cardInfos, &cardArr)
	cardsMap := map[string]cards.CardDesc{}

	//获取不同卡项-适用单项目的个数和赠送单项目的个数
	gaagsNumMap := GetApplyAndGiveSingleNum(cardIDs, cards.ITEM_TYPE_card)

	for k, card := range cardInfos {
		cardId, _ := strconv.Atoi(card[ncm.Field.F_card_id].(string))
		//isGround, _ := strconv.Atoi(card[ncm.Field.F_is_ground].(string))
		clicks, _ := redis2.Int(redis.RedisGlobMgr.Hget(constkey.CARD_CLIKS, card[ncm.Field.F_card_id].(string)))
		bindId, _ := strconv.Atoi(card[ncm.Field.F_bind_id].(string))
		sales, _ := strconv.Atoi(card[ncm.Field.F_sales].(string))

		cardsMap[card[ncm.Field.F_card_id].(string)] = cards.CardDesc{
			CardBase: cardArr[k],
			CardID:   cardId,
			BindID:   bindId,
			Clicks:   clicks,
			Sales:    sales,
			//IsGround:   isGround,
			ShopStatus:     0,
			ShopHasAdd:     1,
			ShopItemId:     0,
			IsAllSingle:    gaagsNumMap[cardId].IsAllSingle,
			ApplySingleNum: gaagsNumMap[cardId].ApplySingleNum,
			GiveSingleNum:  gaagsNumMap[cardId].GiveSingleNum,
		}
	}

	for _, shopcard := range shopCards {
		listShopCard := cardsMap[shopcard[snm.Field.F_card_id].(string)]
		listShopCard.ShopItemId, _ = strconv.Atoi(shopcard[snm.Field.F_id].(string))
		listShopCard.ShopStatus, _ = strconv.Atoi(shopcard[snm.Field.F_status].(string))
		listShopCard.Sales, _ = strconv.Atoi(shopcard[snm.Field.F_sales].(string))

		list.List = append(list.List, listShopCard)
	}

	//获取图片信息
	imgIds := functions.ArrayValue2Array(ncm.Field.F_img_id, cardInfos)
	if len(imgIds) > 0 {
		list.IndexImg = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_card)
	}
	//获取数量
	list.TotalNum = snm.GetNumByShopID(shopId, status)
	return
}

//门店上下架综合卡
func (c *CardLogic) ShopDownUpCard(ctx context.Context, shopId int, args *cards.ArgsShopDownUpCard) (err error) {

	args.CardIDs = functions.ArrayUniqueInt(args.CardIDs)
	if len(args.CardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.CARD_ID_IS_NIL)
		return
	}
	//获取门店综合卡信息
	snm := new(models.ShopCardModel).Init()
	snm.Model.Begin()
	shopCards := snm.GetByShopIDAndCardIDs(shopId, args.CardIDs)
	var shopCardStruct []struct {
		Id     int
		ShopId int
		Status int
		CardId int
	}
	var upIds, downIds, cardIds []int
	_ = mapstructure.WeakDecode(shopCards, &shopCardStruct)
	for _, shopCardDesc := range shopCardStruct {
		if shopCardDesc.Status == cards.STATUS_OFF_SALE {
			downIds = append(downIds, shopCardDesc.Id)
			cardIds = append(cardIds, shopCardDesc.CardId)
		} else if shopCardDesc.Status == cards.STATUS_ON_SALE {
			upIds = append(upIds, shopCardDesc.Id)
			cardIds = append(cardIds, shopCardDesc.CardId)
		}
	}
	cardModel := new(models.CardModel).Init(snm.Model.GetOrmer())
	var decOrInc string
	//综合卡下架
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		if err = snm.UpdateByIDs(upIds, map[string]interface{}{
			snm.Field.F_status:     cards.STATUS_OFF_SALE,
			snm.Field.F_under_time: time.Now().Unix(),
		}); err != nil {
			snm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//同步下架门店卡项关联表数据
		shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
		if !shopItemRealtionModel.UpdateStatusByItemIds(cardIds, cards.ITEM_TYPE_card, cards.STATUS_OFF_SALE, shopId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		decOrInc = "dec"
	}
	//综合卡上架
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		if err = snm.UpdateByIDs(downIds, map[string]interface{}{
			snm.Field.F_status:     cards.STATUS_ON_SALE,
			snm.Field.F_under_time: 0,
		}); err != nil {
			snm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//同步上架门店卡项关联表数据
		shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
		if !shopItemRealtionModel.UpdateStatusByItemIds(cardIds, cards.ITEM_TYPE_card, cards.STATUS_ON_SALE, shopId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		decOrInc = "inc"
	}
	if len(decOrInc) > 0 {
		//	更新总店中对应综合卡的在售门店数量
		if !cardModel.UpdateSaleShopNum(cardIds, decOrInc) {
			snm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	snm.Model.Commit()
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, []int{}, shopId, cards.ITEM_TYPE_card, upIds)
	}
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, []int{}, shopId, cards.ITEM_TYPE_card, downIds)
	}
	return
}

func (c *CardLogic) GetCardsInfo(ctx context.Context, shopId int, cardIds []int) (map[int]*cards.ReplyCardsInfo, error) {
	shopCardModel := new(models.ShopCardModel).Init()
	shopCards := shopCardModel.GetByShopIDAndCardIDs(shopId, cardIds)
	if len(shopCards) != len(cardIds) {
		return nil, toolLib.CreateKcErr(_const.NO_IN_SHOP)
	}
	for _, shopCard := range shopCards {
		status, _ := strconv.Atoi(shopCard[shopCardModel.Field.F_status].(string))

		// 验证门店是否允许销售该卡
		if err := cards.VerifyStatus(status); err != nil {
			return nil, err
		}
	}

	cardModel := new(models.CardModel).Init()
	cardInfos := cardModel.GetByCardIDs(cardIds)
	if len(cardInfos) != len(cardIds) {
		return nil, toolLib.CreateKcErr(_const.NO_IN_SHOP)
	}

	var GiveCardIds []int
	for _, cardInfo := range cardInfos {
		ground, _ := strconv.Atoi(cardInfo[cardModel.Field.F_is_ground].(string))

		// 验证是否上架
		if err := cards.VerifyGround(ground); err != nil {
			return nil, err
		}

		// 赠送项目
		if cardInfo[cardModel.Field.F_has_give_signle].(string) == strconv.Itoa(cards.HAS_GIVE_SINGLE_yes) {
			cardId, _ := strconv.Atoi(cardInfo[cardModel.Field.F_card_id].(string))
			GiveCardIds = append(GiveCardIds, cardId)
		}
	}

	replyCardsInfo := map[int]*cards.ReplyCardsInfo{}
	for _, cardInfo := range cardInfos {
		cardId, _ := strconv.Atoi(cardInfo[cardModel.Field.F_card_id].(string))
		realPrice, _ := strconv.ParseFloat(cardInfo[cardModel.Field.F_real_price].(string), 64)
		price, _ := strconv.ParseFloat(cardInfo[cardModel.Field.F_price].(string), 64)
		replyCardsInfo[cardId] = &cards.ReplyCardsInfo{
			CardID:         cardId,
			Name:           cardInfo[cardModel.Field.F_name].(string),
			RealPrice:      realPrice,
			Price:          price,
			IncludeSingles: nil,
			IncProducts:    nil,
			GiveSingles:    nil,
		}
	}

	//获取包含的商品
	cardGoodsModel := new(models.CardGoodsModel).Init()
	cardProducts := cardGoodsModel.GetByCardIds(cardIds)

	if len(cardProducts) > 0 {
		productIds := functions.ArrayValue2Array(cardGoodsModel.Field.F_product_id, cardProducts)
		var allProducts map[int]cards.IncProductDetail
		var err error
		if allProducts, err = getIncProducts(ctx, productIds); err != nil {
			return nil, err
		}

		for _, cardProduct := range cardProducts {
			cardId, _ := strconv.Atoi(cardProduct[cardGoodsModel.Field.F_card_id].(string))
			productId, _ := strconv.Atoi(cardProduct[cardGoodsModel.Field.F_product_id].(string))
			replyCardsInfo[cardId].IncProducts = append(replyCardsInfo[cardId].IncProducts, cards.IncProductDetail{
				IncProduct: cards.IncProduct{
					ProductID: productId,
				},
				Name:      allProducts[productId].Name,
				SpecPrice: allProducts[productId].SpecPrice,
			})
		}
	}

	////获取包含的单项目
	cardSingleModel := new(models.CardSingleModel).Init()
	cardSingles := cardSingleModel.GetByCardIds(cardIds)
	if len(cardSingles) > 0 {
		singleIds := functions.ArrayValue2Array(cardSingleModel.Field.F_single_id, cardSingles)
		allSingles := getIncInfSingles(ctx, singleIds)

		for _, cardSingle := range cardSingles {
			cardId, _ := strconv.Atoi(cardSingle[cardSingleModel.Field.F_card_id].(string))
			singleId, _ := strconv.Atoi(cardSingle[cardSingleModel.Field.F_single_id].(string))
			replyCardsInfo[cardId].IncludeSingles = append(replyCardsInfo[cardId].IncludeSingles, cards.IncSingleDetail{
				IncSingle: cards.IncSingle{
					SingleID: singleId,
					Num:      0,
				},
				Name:      allSingles[singleId].Name,
				Price:     allSingles[singleId].Price,
				RealPrice: allSingles[singleId].RealPrice,
			})
		}
	}

	//获取包含的赠送项目信息
	cardGiveModel := new(models.CardGiveModel).Init()
	if len(GiveCardIds) > 0 {
		cardGives := cardGiveModel.GetByCardIds(GiveCardIds)
		allGiveSingles := getIncSingles(ctx, functions.ArrayValue2Array(cardGiveModel.Field.F_single_id, cardGives))

		for _, cardGive := range cardGives {
			cardId, _ := strconv.Atoi(cardGive[cardGiveModel.Field.F_card_id].(string))
			singleId, _ := strconv.Atoi(cardGive[cardGiveModel.Field.F_single_id].(string))
			num, _ := strconv.Atoi(cardGive[cardGiveModel.Field.F_num].(string))
			replyCardsInfo[cardId].GiveSingles = append(replyCardsInfo[cardId].GiveSingles, cards.IncSingleDetail{
				IncSingle: cards.IncSingle{
					SingleID: singleId,
					Num:      num,
				},
				Name:      allGiveSingles[singleId].Name,
				Price:     allGiveSingles[singleId].Price,
				RealPrice: allGiveSingles[singleId].RealPrice,
			})
		}
	}
	return replyCardsInfo, nil
}

// 门店综合卡数据rpc内部调用
func (n *CardLogic) ShopCardListRpc(ctx context.Context, params *cards.ArgsShopCardListRpc, list *cards.ReplyShopCardListRpc) (err error) {
	if params.ShopId <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//获取门店的综合卡数据
	snm := new(models.ShopCardModel).Init()
	shopCards := snm.GetByShopIDAndCardIDs(params.ShopId, params.CardIds)
	shopCards = SortMapByIntField(snm.Field.F_card_id, shopCards)
	cardIDs := functions.ArrayValue2Array(snm.Field.F_card_id, shopCards)
	//获取综合卡基本信息
	ncm := new(models.CardModel).Init()
	cardInfos := ncm.GetByCardIDs(cardIDs)
	if len(cardInfos) == 0 {
		return
	}
	list.List = make([]cards.CardDesc, len(cardInfos))
	for index, card := range cardInfos {
		_ = mapstructure.WeakDecode(card, &list.List[index].CardBase)
		_ = mapstructure.WeakDecode(card, &list.List[index])
	}
	for index := 0; index < len(list.List); index++ {
		_ = mapstructure.WeakDecode(shopCards[index], &list.List[index])
		list.List[index].ShopHasAdd = 1
	}
	return
}

// 获取所有卡的发布数量 rpc内部调用
func (n *CardLogic) GetAllCardsNum(ctx context.Context, args *cards.ArgsAllCardsNum, reply *cards.ReplyAllCardsNum) (err error) {
	// 综合卡
	cardModel := new(models.CardModel).Init()
	cardNum := cardModel.GetNumByBusID(args.BusId)

	// 限时卡
	hcardModle := new(models.HcardModel).Init()
	hcardNum := hcardModle.GetNumByBusID(args.BusId)

	// 身份卡表
	icardModle := new(models.IcardModel).Init()
	icardNum := icardModle.GetNumByBusID(args.BusId)

	//限次卡表
	ncardModel := new(models.NCardModel).Init()
	ncardNum := ncardModel.GetNumByBusID(args.BusId)

	// 充值卡
	rcardModel := new(models.RcardModel).Init()
	rcardNum := rcardModel.GetNumByBusId(args.BusId)

	// 单项目
	singleModel := new(models.SingleModel).Init()
	singleNum := singleModel.GetNumByBusId(args.BusId, "")

	//套餐
	smModel := new(models.SmModel).Init()
	smNum := smModel.GetNumByBusId(args.BusId)

	allNum := cardNum + hcardNum + icardNum + ncardNum + rcardNum + singleNum + smNum
	reply.AllCardsNum = allNum

	return
}

//总店-删除-综合卡
func (n *CardLogic) DeleteCard(ctx context.Context, args *cards.ArgsDeleteCard, reply *bool) (err error) {
	//总店-软删除
	card := new(models.CardModel).Init()
	bId, _ := args.GetBusId()
	card.Model.Begin()
	if _, err = card.DelByCardId(args.CardIds, bId); err != nil {
		card.Model.RollBack()
		return
	}
	//分店-软删除
	shopCard := new(models.ShopCardModel).Init()
	where := []base.WhereItem{
		{shopCard.Field.F_card_id, []interface{}{"IN", args.CardIds}},
	}
	if _, err = shopCard.DelShopCardById(where); err != nil {
		card.Model.RollBack()
		return
	}
	//同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.CardIds, cards.ITEM_TYPE_card) {
		card.Model.RollBack()
		return
	}
	card.Model.Commit()
	return
}

//分店-删除-综合卡 DeleteShopCard()
func (n *CardLogic) DeleteShopCard(ctx context.Context, args *cards.ArgsDeleteShopCard, reply *bool) (err error) {
	//分店-软删除
	shopCard := new(models.ShopCardModel).Init()
	shopId, _ := args.GetShopId()
	where := []base.WhereItem{
		{shopCard.Field.F_card_id, []interface{}{"IN", args.CardIds}},
		{shopCard.Field.F_shop_id, shopId},
	}
	shopCard.Model.Begin()

	if _, err = shopCard.DelShopCardById(where); err != nil {
		shopCard.Model.RollBack()
		return
	}
	//同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.CardIds, cards.ITEM_TYPE_card, shopId) {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	shopCard.Model.Commit()
	return
}
