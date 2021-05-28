//限次卡业务处理
//@author loop
//@date 2020/4/15 14:55
package logics

import (
	"context"
	"encoding/json"
	"git.900sui.cn/kc/base/common/models/base"
	"strconv"
	"time"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/kcgin"
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
)

type NCardLogic struct {
}

//添加NCard描述
func (n *NCardLogic) AddCardExt(ncm *models.NCardModel, nCardID int, notes []cards.CardNote) (err error) {
	if len(notes) > constkey.CardNotesMaxNum {
		return toolLib.CreateKcErr(_const.NOTES_NUM_MAX_10)
	}
	for _, note := range notes {
		if functions.Mb4Strlen(note.Notes) > constkey.CardNotesSimpleMaxLength {
			return toolLib.CreateKcErr(_const.NOTES_LEN_MAX_30)
		}
	}
	notesStr, _ := json.Marshal(notes)
	nem := new(models.NCardExtModel).Init(ncm.Model.GetOrmer())
	if _, err = nem.Insert(map[string]interface{}{nem.Field.F_ncard_id: nCardID, nem.Field.F_notes: string(notesStr)}); err != nil {
		err = toolLib.CreateKcErr(_const.DB_ERR)
	}
	return
}

//编辑NCard描述
func (n *NCardLogic) EditCardExt(ncm *models.NCardModel, nCardID int, notes []cards.CardNote) (err error) {
	if len(notes) > constkey.CardNotesMaxNum {
		return toolLib.CreateKcErr(_const.NOTES_NUM_MAX_10)
	}
	for _, note := range notes {
		if functions.Mb4Strlen(note.Notes) > constkey.CardNotesSimpleMaxLength {
			return toolLib.CreateKcErr(_const.NOTES_LEN_MAX_30)
		}
	}
	notesStr, _ := json.Marshal(notes)
	nem := new(models.NCardExtModel).Init()
	nCardExMap := nem.Find(map[string]interface{}{nem.Field.F_ncard_id: nCardID})
	if len(nCardExMap) > 0 {
		if _, updateErr := nem.Update(map[string]interface{}{nem.Field.F_ncard_id: nCardID},
			map[string]interface{}{nem.Field.F_notes: string(notesStr)}); updateErr != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	} else {
		if _, err = nem.Insert(map[string]interface{}{nem.Field.F_ncard_id: nCardID, nem.Field.F_notes: string(notesStr)}); err != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	return
}

//获取NCard描述
func (n *NCardLogic) GetCardExt(nCardID int, notes *[]cards.CardNote) {
	nem := new(models.NCardExtModel).Init()
	dataMap := nem.Find(map[string]interface{}{nem.Field.F_ncard_id: nCardID})
	if len(dataMap) > 0 {
		note := dataMap[nem.Field.F_notes].(string)
		json.Unmarshal([]byte(note), notes)
	}
	return
}

//添加ncard
func (n *NCardLogic) AddNCard(ctx context.Context, busID int, args *cards.ArgsAddNCard) (nCardID int, err error) {

	//验证参数
	err = n.checkNCardData(busID, args.NCardBase, args.IncludeSingles, args.GiveSingles)
	if err != nil {
		return
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
	ncm := new(models.NCardModel).Init()
	nsm := new(models.NCardSingleModel).Init(ncm.Model.GetOrmer())

	//计算包含单项目总次数
	var sum int = 0
	for _, single := range args.IncludeSingles {
		sum += single.Num
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
	//添加基本信息
	ncm.Model.Begin()
	if nCardID, err = ncm.Insert(map[string]interface{}{
		ncm.Field.F_bus_id:                busID,
		ncm.Field.F_price:                 args.Price,
		ncm.Field.F_is_ground:             cards.SINGLE_IS_GROUND_no,
		ncm.Field.F_real_price:            args.RealPrice,
		ncm.Field.F_ctime:                 time.Now().Unix(),
		ncm.Field.F_img_id:                imgID,
		ncm.Field.F_name:                  args.Name,
		ncm.Field.F_sort_desc:             args.ShortDesc,
		ncm.Field.F_bind_id:               getBusMainBindId(ctx, busID),
		ncm.Field.F_has_give_signle:       hasGive,
		ncm.Field.F_service_period:        args.ServicePeriod,
		ncm.Field.F_validcount:            sum,
		ncm.Field.F_is_permanent_validity: ispermanentvalidity,
	}); err != nil {
		ncm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	// 添加单项目，先去单项目规格表中查找是否含有次规格卡(并且没有删除）,入参之前已经检查
	// 添加单项目
	var nCardSingleData []map[string]interface{}
	for _, single := range args.IncludeSingles {
		nCardSingleData = append(nCardSingleData, map[string]interface{}{
			nsm.Field.F_single_id: single.SingleID,
			nsm.Field.F_ncard_id:  nCardID,
			nsm.Field.F_num:       single.Num,
			nsm.Field.F_ssp_id:    single.SspId,
		})
	}
	if err = nsm.InsertAll(nCardSingleData); err != nil {
		ncm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//添加赠送单项目数据
	if len(args.GiveSingles) > 0 {
		ngm := new(models.NCardGiveModel).Init(ncm.Model.GetOrmer())
		var giveSingleData []map[string]interface{}
		for _, single := range args.GiveSingles {
			giveSingleData = append(giveSingleData, map[string]interface{}{
				ngm.Field.F_single_id:          single.SingleID,
				ngm.Field.F_ncard_id:           nCardID,
				ngm.Field.F_num:                single.Num,
				ngm.Field.F_period_of_validity: single.PeriodOfValidity,
			})
		}
		if err = ngm.InsertAll(giveSingleData); err != nil {
			ncm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//添加赠品描述
	if hasGive && len(args.GiveSingleDesc) > 0 {
		ndescm := new(models.NcardGiveDescModel).Init(ncm.Model.GetOrmer())
		giveSingleDescStr, _ := json.Marshal(args.GiveSingleDesc)
		descData := map[string]interface{}{
			ndescm.Field.F_ncard_id: nCardID,
			ndescm.Field.F_desc:     string(giveSingleDescStr),
		}
		if _, err = ndescm.Model.Data(descData).Insert(); err != nil {
			ncm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//添加NCardExt
	if err = n.AddCardExt(ncm, nCardID, args.Notes); err != nil {
		ncm.Model.RollBack()
		return
	}
	ncm.Model.Commit()
	//添加风控统计任务
	new(ItemLogic).AddXCardTask(ctx, nCardID, cards.ITEM_TYPE_ncard)
	return
}

//编辑NCard数据
func (n *NCardLogic) EditNCard(ctx context.Context, busID int, args *cards.ArgsEditNCard) (err error) {

	//验证NCard数据
	ncm := new(models.NCardModel).Init()
	nCardInfo := ncm.GetByNCardID(args.CardID, ncm.Field.F_bus_id)
	if len(nCardInfo) == 0 {
		err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
		return
	}
	if nCardInfo[ncm.Field.F_bus_id] != strconv.Itoa(busID) {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	//验证参数
	err = n.checkNCardData(busID, args.NCardBase, args.IncludeSingles, args.GiveSingles)
	if err != nil {
		return
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
	nsm := new(models.NCardSingleModel).Init(ncm.Model.GetOrmer())
	ngm := new(models.NCardGiveModel).Init(ncm.Model.GetOrmer())

	//计算单项目改动
	includeSingles := nsm.GetByNCardID(args.CardID)
	//需要新增的项目
	var addIncSingles []map[string]interface{}
	//需要修改的单项目
	var updateSingles = map[int]int{} //id:num
	//需要删除的ids
	var delIDs []int
	//获取需要修改的项目以及需要添加的项目

	for _, single := range args.IncludeSingles {
		hasd := 0
		for _, dbSingle := range includeSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[nsm.Field.F_single_id].(string) {
				hasd = 1
				//需要修改次数
				if strconv.Itoa(single.Num) != dbSingle[nsm.Field.F_num].(string) {
					dbID, _ := strconv.Atoi(dbSingle[nsm.Field.F_id].(string))
					updateSingles[dbID] = single.Num
				}

			}
		}

		if hasd == 0 {
			addIncSingles = append(addIncSingles, map[string]interface{}{
				nsm.Field.F_ncard_id:  args.CardID,
				nsm.Field.F_num:       single.Num,
				nsm.Field.F_single_id: single.SingleID,
			})
		}

	}
	//获取需要删除的id
	for _, dbSingle := range includeSingles {
		hasd := 0
		for _, single := range args.IncludeSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[nsm.Field.F_single_id].(string) {
				hasd = 1
				break
			}
		}

		if hasd == 0 {
			sid, _ := strconv.Atoi(dbSingle[nsm.Field.F_id].(string))
			delIDs = append(delIDs, sid)
		}
	}
	//计算赠送项目改动
	giveSingles := ngm.GetByNCardID(args.CardID)
	//if len(giveSingles) == 0 && len(args.GiveSingles) == 0 {
	//	return
	//}
	//需要新增的项目
	var addGiveSingles []map[string]interface{}
	//需要修改的单项目
	var updateGiveSingles = map[int]int{} //id:num
	//需要删除的ids
	var delGiveIDs []int
	//获取需要修改的项目以及需要添加的项目
	for _, single := range args.GiveSingles {
		hasd := 0
		for _, dbSingle := range giveSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[ngm.Field.F_single_id].(string) {
				hasd = 1
				//需要修改次数
				if strconv.Itoa(single.Num) != dbSingle[ngm.Field.F_num].(string) {
					dbID, _ := strconv.Atoi(dbSingle[ngm.Field.F_id].(string))
					updateGiveSingles[dbID] = single.Num
				}
			}
		}
		if hasd == 0 {
			addGiveSingles = append(addGiveSingles, map[string]interface{}{
				ngm.Field.F_ncard_id:           args.CardID,
				ngm.Field.F_num:                single.Num,
				ngm.Field.F_single_id:          single.SingleID,
				ngm.Field.F_period_of_validity: single.PeriodOfValidity,
			})
		}
	}
	//获取需要删除的id

	for _, dbSingle := range giveSingles {
		hasd := 0
		for _, single := range args.GiveSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[ngm.Field.F_single_id].(string) {
				hasd = 1
				break
			}
		}

		if hasd == 0 {
			sid, _ := strconv.Atoi(dbSingle[ngm.Field.F_id].(string))
			delGiveIDs = append(delGiveIDs, sid)
		}
	}

	//计算包含单项目总次数
	var sum int = 0
	for _, single := range args.IncludeSingles {
		sum += single.Num
	}
	// 卡项有效期 1-是；2-否；如果永久有效则service_period字段可忽略，否者必填
	ispermanentvalidity := cards.IS_PERMANENT_YES
	if args.ServicePeriod != 0 {
		ispermanentvalidity = cards.IS_PERMANENT_NO
	}
	//修改主表信息
	ncm.Model.Begin()
	if err = ncm.UpdateByNCardID(args.CardID, map[string]interface{}{
		ncm.Field.F_price:                 args.Price,
		ncm.Field.F_real_price:            args.RealPrice,
		ncm.Field.F_img_id:                imgID,
		ncm.Field.F_name:                  args.Name,
		ncm.Field.F_sort_desc:             args.ShortDesc,
		ncm.Field.F_has_give_signle:       hasGive,
		ncm.Field.F_service_period:        args.ServicePeriod,
		ncm.Field.F_validcount:            sum,
		ncm.Field.F_is_permanent_validity: ispermanentvalidity,
	}); err != nil {
		ncm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	//修改包含的单项
	//1. 修改需要修改的项目
	if len(updateSingles) > 0 {
		for id, num := range updateSingles {
			if err = nsm.UpdateNumById(id, num); err != nil {
				ncm.Model.RollBack()
				err = toolLib.CreateKcErr(_const.DB_ERR)
				return
			}
		}
	}
	//2. 添加需要添加的项目
	if len(addIncSingles) > 0 {
		if err = nsm.InsertAll(addIncSingles); err != nil {
			ncm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//3. 删除需要删除的项目
	if len(delIDs) > 0 {
		if err = nsm.DelByIds(delIDs); err != nil {
			ncm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//修改赠送的单项目
	//1. 修改需要修改的项目
	if len(updateGiveSingles) > 0 {
		for id, num := range updateGiveSingles {
			if err = ngm.UpdateNumById(id, num); err != nil {
				ncm.Model.RollBack()
				err = toolLib.CreateKcErr(_const.DB_ERR)
				return
			}
		}
	}
	//2. 添加需要添加的项目
	if len(addGiveSingles) > 0 {
		if err = ngm.InsertAll(addGiveSingles); err != nil {
			ncm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//3. 删除需要删除的项目
	if len(delGiveIDs) > 0 {
		if err = ngm.DelByIds(delGiveIDs); err != nil {
			ncm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//修改赠品描述
	ndesm := new(models.NcardGiveDescModel).Init(ncm.Model.GetOrmer())
	//1.删除原有赠品描述
	if err = ndesm.DelByNcardId(args.CardID); err != nil {
		ncm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//2.新增赠品描述
	if len(args.GiveSingles) > 0 && len(args.GiveSingleDesc) > 0 {
		giveSingleDescStr, _ := json.Marshal(args.GiveSingleDesc)
		descData := map[string]interface{}{
			ndesm.Field.F_ncard_id: args.CardID,
			ndesm.Field.F_desc:     string(giveSingleDescStr),
		}
		if _, err = ndesm.Model.Data(descData).Insert(); err != nil {
			ncm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//编辑cardExt
	if err = n.EditCardExt(ncm, args.CardID, args.Notes); err != nil {
		ncm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	ncm.Model.Commit()

	return
}

// 检查添加的限次卡规格是否存在
func checkNcardSspId(sspIds, singleIDs []int, incSingles []cards.IncSingle) (err error) {
	//验证规格是否重复
	sspIds = tools.RemoveArrayZero(sspIds)

	//规格ids去重判断长度
	uniqueSspIds := functions.ArrayUniqueInt(sspIds)
	if len(sspIds) != len(uniqueSspIds) {
		err = toolLib.CreateKcErr(_const.SPEC_REPEAT_ERR)
		return
	}
	//刷选有规格的单项目
	var smSingleIds []int
	sm := new(models.SingleModel).Init()
	singlesInfo := sm.Model.Where(map[string]interface{}{
		sm.Field.F_single_id: []interface{}{"IN", singleIDs},
		sm.Field.F_has_spec:  1,
	}).Field([]string{sm.Field.F_single_id}).Select()
	for _, single := range singlesInfo {
		SingleId, _ := strconv.Atoi(single[sm.Field.F_single_id].(string))
		smSingleIds = append(smSingleIds, SingleId)
	}
	//从规格表获取单项目规格对应关系
	sspm := new(models.SingleSpecPriceModel).Init()
	singleSpecPrices := sspm.GetBySspids(sspIds, sspm.Field.F_single_id, sspm.Field.F_ssp_id)
	singleToSpec := make(map[int][]int)
	for _, ssp := range singleSpecPrices {
		sspSingleId, _ := strconv.Atoi(ssp[sspm.Field.F_single_id].(string))
		sspSpecId, _ := strconv.Atoi(ssp[sspm.Field.F_ssp_id].(string))
		singleToSpec[sspSingleId] = append(singleToSpec[sspSingleId], sspSpecId)
	}

	// 规格选择是否正确
	for k, incSingle := range incSingles {
		//有规格的则进入判断
		if functions.InArray(incSingle.SingleID, smSingleIds) {
			//未传规格
			if incSingle.SspId == 0 {
				err = toolLib.CreateKcErr(_const.NO_SPEC_ERR)
				return
			}
			//规格不属于该单项目
			if !functions.InArray(incSingle.SspId, singleToSpec[incSingle.SingleID]) {
				err = toolLib.CreateKcErr(_const.SPEC_DATA_ERR)
				return
			}
		} else {
			if incSingle.SspId != 0 {
				err = toolLib.CreateKcErr(_const.SPEC_HAS_NO_ERR)
				return
			}
			incSingles[k].SspId = 0
		}
	}

	return
}

//检查限次卡的入参数据
func (n *NCardLogic) checkNCardData(busID int, nCardBase cards.NCardBase, incSingles []cards.IncSingle, giveSingles []cards.IncSingle) (err error) {
	if err = cards.VerfiyName(nCardBase.Name); err != nil {
		return
	}
	if err = cards.VerfiyPrice(nCardBase.RealPrice, nCardBase.Price); err != nil {
		return
	}
	if err = cards.VerfiyServicePeriod(nCardBase.ServicePeriod); err != nil {
		return
	}
	if err = cards.VerifySinglesNum(len(incSingles)); err != nil {
		return
	}
	if err = cards.VerifyGiveSinglesNum(len(giveSingles)); err != nil {
		return
	}
	//单项目id集合
	var singleIDs []int
	//sspIds := make([]int,0)
	for _, single := range incSingles {
		if single.Num <= 0 {
			err = toolLib.CreateKcErr(_const.PARAM_ERR)
			return
		}
		singleIDs = append(singleIDs, single.SingleID)
		//if single.SspId > 0 {
		//	sspIds = append(sspIds, single.SspId)
		//}
	}
	if err = checkSingles(busID, singleIDs); err != nil {
		return
	}

	// 检查添加的单项目是否存在
	//sspIds := make([]int, 0)
	//for _, single := range incSingles {
	//	sspIds = append(sspIds, single.SspId)
	//}
	//if err = checkNcardSspId(sspIds, singleIDs, incSingles); err != nil {
	//	return
	//}

	// 赠送项目
	if len(giveSingles) > 0 {
		var giveIDs []int
		for _, giveSingle := range giveSingles {
			if giveSingle.Num <= 0 {
				err = toolLib.CreateKcErr(_const.PARAM_ERR)
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

func getBusInfo(ctx context.Context, busID int, info *cards.BusInfo) (err error) {
	args := new(bus2.ArgsBusInfo)
	args.BusID = busID
	reply := new(bus2.ReplyBusInfo)
	bc := new(bus.Bus).Init()
	defer bc.Close()
	if err = bc.BusInfo(ctx, args, reply); err != nil {
		return
	}
	info.BusCompanyName = reply.CompanyName
	info.BusBrandName = reply.BrandName
	if reply.BrandImg > 0 {
		_, info.BusIcon = getImg(ctx, reply.BrandImg, 0)
	} else {
		info.BusIcon = kcgin.AppConfig.String("public.defaultIconImage")
	}

	return
}

//获取限次卡的详情
//@param int nCardID 限次卡id
//@param int shopID 门店id
//@return  cards.ReplyNCardInfo reply
func (n *NCardLogic) NCardInfo(ctx context.Context, nCardID int, shopID ...int) (reply cards.ReplyNCardInfo, err error) {
	reply.IncludeSingles = make([]cards.IncSingleDetail2, 0)
	reply.GiveSingles = make([]cards.IncSingleDetail2, 0)
	reply.GiveSingleDesc=make([]cards.GiveSingleDesc,0)
	ncm := new(models.NCardModel).Init()
	nCardInfo := ncm.GetByNCardID(nCardID)
	if len(nCardInfo) == 0 {
		err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
		return
	}

	var shopStatus = 0
	var shopSales = 0
	var shopid = 0
	if len(shopID) > 0 && shopID[0] > 0 {
		snm := new(models.ShopNCardModel).Init()
		shopNcInfo := snm.GetByShopIDAndNCardIDs(shopID[0], []int{nCardID})
		if len(shopNcInfo) == 0 {
			err = toolLib.CreateKcErr(_const.NO_IN_SHOP)
			return
		}
		shopStatus, _ = strconv.Atoi(shopNcInfo[0][snm.Field.F_status].(string))
		reply.SsId, _ = strconv.Atoi(shopNcInfo[0][snm.Field.F_id].(string))
		shopSales, _ = strconv.Atoi(shopNcInfo[0][snm.Field.F_sales].(string))
		shopid = shopID[0]
	}
	reply.ShareLink = tools.GetShareLink(nCardID, shopid, cards.ITEM_TYPE_ncard)
	imgId, _ := strconv.Atoi(nCardInfo[ncm.Field.F_img_id].(string))
	imgHash, imgUrl := getImg(ctx, imgId, cards.ITEM_TYPE_ncard)
	reply.ImgHash = imgHash
	reply.ImgUrl = imgUrl
	reply.ShopStatus = shopStatus
	_ = mapstructure.WeakDecode(nCardInfo, &reply.NCardBase)
	reply.CtimeStr = time.Unix(int64(reply.Ctime), 0).Format("2006/01/02 15:04:05")
	_ = mapstructure.WeakDecode(nCardInfo, &reply)
	if len(shopID) > 0 && shopID[0] > 0 {
		reply.Sales = shopSales
	}
	//商户信息
	if err = getBusInfo(ctx, reply.BusID, &reply.BusInfo); err != nil {
		err = toolLib.CreateKcErr(_const.SHOP_INFO_ERR)
		return
	}
	//统计限时限次卡包含项目的总次数
	reply.SingleTotalNum = GetItemCardSingleNum(nCardID, cards.ITEM_TYPE_ncard)
	//获取包含的限次卡和赠送项目信息
	nsm := new(models.NCardSingleModel).Init()
	nCardSingles := nsm.GetByNCardID(nCardID)
	if len(nCardSingles) > 0 && nCardSingles[0][nsm.Field.F_single_id].(string) == "0" {
		reply.IsAllSingle = true
	}
	//singleIds := functions.ArrayValue2Array(nsm.Field.F_single_id, nCardSingles)
	var singleIds []int
	ngm := new(models.NCardGiveModel).Init()
	var nCardGives []map[string]interface{}
	if nCardInfo[ncm.Field.F_has_give_signle].(string) == strconv.Itoa(cards.HAS_GIVE_SINGLE_yes) {
		nCardGives = ngm.GetByNCardID(nCardID)
		giveSingleIds := functions.ArrayValue2Array(ngm.Field.F_single_id, nCardGives)
		singleIds = append(singleIds, giveSingleIds...)
	}

	if len(nCardGives) > 0 {
		allSingles ,_:= getIncSingles2(ctx,shopid,"","", singleIds, nCardGives)
		for _, single := range nCardGives {
			for i := range allSingles {
				if single[ngm.Field.F_single_id].(string) == strconv.Itoa(allSingles[i].SingleID) {
					allSingles[i].Num, _ = strconv.Atoi(single[ngm.Field.F_num].(string))
					allSingles[i].PeriodOfValidity, _ = strconv.Atoi(single[ngm.Field.F_period_of_validity].(string))
					reply.GiveSingles = append(reply.GiveSingles, allSingles[i])
					break
				}
			}
		}

		// 获取赠品描述信息
		ndescm := new(models.NcardGiveDescModel).Init()
		desc, ok := ndescm.GetByNcardId(nCardID)[ndescm.Field.F_desc].(string)
		if ok {
			json.Unmarshal([]byte(desc), reply.GiveSingleDesc)
		}
	}

	//获取CardExt
	n.GetCardExt(nCardID, &reply.Notes)

	// 获取限次卡门店添加详情  15 -- []int {3,4,6}
	busId, _ := strconv.Atoi(nCardInfo[ncm.Field.F_bus_id].(string))
	nCardShopModel := new(models.NCardShopModel).Init()
	nCardShopLists := nCardShopModel.GetByNcardIdAndBusId(nCardID, busId)

	nCardShopIds := make([]int, 0)
	for _, icInfoValue := range nCardShopLists {
		sshopId, _ := strconv.Atoi(icInfoValue[nCardShopModel.Field.F_shop_id].(string))
		nCardShopIds = append(nCardShopIds, sshopId)
	}
	var replyShop []bus2.ReplyShopName
	rLists := make([]cards.ReplyShopName, 0) // 临时存放返回数据
	rpcBus := new(bus.Shop).Init()
	defer rpcBus.Close()
	err = rpcBus.GetShopNameByShopIds(ctx, &nCardShopIds, &replyShop)
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
	_ = redis.RedisGlobMgr.Hincrby(constkey.NCARD_CLIKS, strconv.Itoa(nCardID), 1)
	return
}

//获取商家的限次卡列表
func (n *NCardLogic) GetBusPage(ctx context.Context, busId, shopId, start, limit int, isGround string, filterShopHasAdd bool) (list cards.ReplyNCardPage, err error) {

	if busId <= 0 || start < 0 || limit < 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	ncm := new(models.NCardModel).Init()
	nCards := make([]map[string]interface{}, 0)
	list.List = make([]cards.NCardDesc, 0)
	list.IndexImg = make(map[int]string)
	//子店已添加的卡项
	var shopAddCards []map[string]interface{}
	scModel := new(models.ShopNCardModel).Init()
	where := []base.WhereItem{
		{ncm.Field.F_bus_id, busId},
		{ncm.Field.F_is_del, cards.IS_BUS_DEL_no},
	}
	if shopId > 0 {
		shopAddCards = scModel.SelectRcardsByWherePage([]base.WhereItem{{scModel.Field.F_shop_id, shopId}, {scModel.Field.F_is_del, cards.IS_BUS_DEL_no}}, 0, 0)
		if filterShopHasAdd && len(shopAddCards) > 0 {
			shopHasAddNcardIds := functions.ArrayValue2Array(scModel.Field.F_ncard_id, shopAddCards)
			where = append(where, base.WhereItem{ncm.Field.F_ncard_id, []interface{}{"NOT IN", shopHasAddNcardIds}})
		}
	}

	//获取总数量
	//if isGround == "" {
	nCards = ncm.SelectNCardsByWherePage(where, start, limit)
	list.TotalNum = ncm.GetNumByWhere(where)

	//} else {
	//	isground, _ := strconv.Atoi(isGround)
	//	isground = isground - 1
	//	nCards = ncm.SelectNCardsByWherePage(where,start,limit,isGround)
	//	list.TotalNum = ncm.GetNumByWhere(where)
	//
	//}
	if len(nCards) == 0 {
		return
	}

	list.List = make([]cards.NCardDesc, len(nCards))
	for index, nc := range nCards {
		_ = mapstructure.WeakDecode(nc, &list.List[index].NCardBase)
		for _, shopCard := range shopAddCards {
			if nc[ncm.Field.F_ncard_id].(string) == shopCard[scModel.Field.F_ncard_id].(string) { //表明当前子店已添加该卡项
				nc["ShopItemId"] = shopCard[scModel.Field.F_id].(string)
				nc["ShopStatus"] = shopCard[scModel.Field.F_status].(string)
				nc["ShopHasAdd"] = 1
				nc["ShopDelStatus"] = shopCard[scModel.Field.F_is_del].(string)
				break
			}
		}
		list.List[index].CtimeStr = time.Unix(int64(list.List[index].Ctime), 0).Format("2006/01/02 15:04:05")
		_ = mapstructure.WeakDecode(nc, &list.List[index])
		list.List[index].Clicks, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.NCARD_CLIKS, nc[ncm.Field.F_ncard_id].(string)))
	}
	//获取图片信息
	imgIds := functions.ArrayValue2Array(ncm.Field.F_img_id, nCards)
	list.IndexImg = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_ncard)

	return
}

//设置限次卡的适用门店(废用）
func (n *NCardLogic) SetNCardShop(ctx context.Context, busId int, args *cards.ArgsSetNCardShop) (err error) {
	if len(args.NCardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.CARD_ID_IS_NIL)
		return
	}
	if args.IsAllShop == false && len(args.ShopIDs) == 0 {
		err = toolLib.CreateKcErr(_const.SHOPID_NTL)
		return
	}

	//限次卡id重复提交判断
	realNCardIDs := functions.ArrayUniqueInt(args.NCardIDs)
	if len(realNCardIDs) != len(args.NCardIDs) {
		err = toolLib.CreateKcErr(_const.DATA_REPEAT_ERR)
		return
	}

	//店铺id重复提交判断
	realShopIds := functions.ArrayUniqueInt(args.ShopIDs)
	if args.IsAllShop == false && len(realShopIds) != len(args.ShopIDs) {
		err = toolLib.CreateKcErr(_const.DATA_REPEAT_ERR)
		return
	}

	//检查限次卡是否属于企业
	ncm := new(models.NCardModel).Init()
	dataArr := ncm.GetByNCardIDs(realNCardIDs)
	if len(dataArr) != len(realNCardIDs) {
		err = toolLib.CreateKcErr(_const.PARAM_ERR, "has invalid cardID")
		return
	}
	busIdStr := strconv.Itoa(busId)
	for _, data := range dataArr {
		if busIdStr != data[ncm.Field.F_bus_id].(string) {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
	}

	nsm := new(models.NCardShopModel).Init()
	var insertData []map[string]interface{}
	//走全部适用逻辑
	if args.IsAllShop == true {
		for _, nCardID := range realNCardIDs {
			insertData = append(insertData, map[string]interface{}{
				nsm.Field.F_ncard_id: nCardID,
				nsm.Field.F_bus_id:   busId,
				nsm.Field.F_shop_id:  0,
			})
		}

	} else {
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
		for _, nCardID := range realNCardIDs {
			for _, shopID := range realShopIds {
				insertData = append(insertData, map[string]interface{}{
					nsm.Field.F_ncard_id: nCardID,
					nsm.Field.F_bus_id:   busId,
					nsm.Field.F_shop_id:  shopID,
				})
			}
		}
	}

	//处理规格
	if len(insertData) > 0 {
		nsm.Model.Begin()
		if err = nsm.DelByNCardIDs(realNCardIDs); err != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			nsm.Model.RollBack()
			return
		}
		if err = nsm.InsertAll(insertData); err != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			nsm.Model.RollBack()
			return
		}
		nsm.Model.Commit()
	}

	var disableIds []int
	var downIds []int

	snm := new(models.ShopNCardModel).Init()
	shopNCards := snm.GetByNCardIDs(realNCardIDs)
	if args.IsAllShop {
		for _, hcard := range dataArr {
			hcardStatus, _ := strconv.Atoi(hcard[ncm.Field.F_is_ground].(string)) // 总店中卡的状态
			hcardID, _ := strconv.Atoi(hcard[ncm.Field.F_ncard_id].(string))
			for _, shopHcard := range shopNCards {
				shopStatus, _ := strconv.Atoi(shopHcard[snm.Field.F_status].(string)) // 子店中卡的状态
				id, _ := strconv.Atoi(shopHcard[snm.Field.F_id].(string))
				shopHcardID, _ := strconv.Atoi(shopHcard[snm.Field.F_ncard_id].(string)) // 子店中卡的id
				if hcardID != shopHcardID {
					continue
				}
				if shopStatus == cards.STATUS_DISABLE && hcardStatus == cards.SINGLE_IS_GROUND_yes { // 子店中卡的状态为禁用并且总店卡的状态为上架时,子店才可以上架
					downIds = append(downIds, id)
				}
			}
		}
	} else {
		for _, hcard := range dataArr {
			hcardStatus, _ := strconv.Atoi(hcard[ncm.Field.F_is_ground].(string)) // 总店中卡的状态
			hcardID, _ := strconv.Atoi(hcard[ncm.Field.F_ncard_id].(string))
			for _, shopHcard := range shopNCards {
				shopStatus, _ := strconv.Atoi(shopHcard[snm.Field.F_status].(string)) // 子店中卡的状态
				id, _ := strconv.Atoi(shopHcard[snm.Field.F_id].(string))
				shopID, _ := strconv.Atoi(shopHcard[snm.Field.F_shop_id].(string))       // 已经存在的子店id
				shopHcardID, _ := strconv.Atoi(shopHcard[snm.Field.F_ncard_id].(string)) // 子店中卡的id
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
		_ = snm.UpdateByIDs(disableIds, map[string]interface{}{
			snm.Field.F_status:     cards.STATUS_DISABLE,
			snm.Field.F_under_time: time.Now().Unix(),
		})
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, []int{}, 0, cards.ITEM_TYPE_ncard, disableIds)
	}
	if len(downIds) > 0 {
		_ = snm.UpdateByIDs(downIds, map[string]interface{}{
			snm.Field.F_status: cards.STATUS_OFF_SALE,
		})
	}
	return
}

//总店上下架限次卡
func (n *NCardLogic) DownUpNCard(ctx context.Context, busId int, args *cards.ArgsDownUpNCard) (err error) {
	// 限次卡ID列表交验
	if len(args.NCardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	realNCardIDs := functions.ArrayUniqueInt(args.NCardIDs)
	//交验是否有重复的id
	if len(realNCardIDs) != len(args.NCardIDs) {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//检查限次卡是否属于企业
	ncm := new(models.NCardModel).Init()
	dataArr := ncm.GetByNCardIDs(realNCardIDs, ncm.Field.F_bus_id, ncm.Field.F_ncard_id, ncm.Field.F_is_ground)
	busIdStr := strconv.Itoa(busId)
	for _, data := range dataArr {
		if busIdStr != data[ncm.Field.F_bus_id].(string) {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
	}
	//将id中已经上架的限次卡id和未上架的限次卡id分开处理
	var nCardDescList []struct {
		NCardID  int `mapstructure:"ncard_id"`
		IsGround int `mapstructure:"is_ground"`
	}
	_ = mapstructure.WeakDecode(dataArr, &nCardDescList)
	var downIds, upIds []int
	for _, nCardDesc := range nCardDescList {
		if nCardDesc.IsGround == cards.IS_GROUND_no {
			downIds = append(downIds, nCardDesc.NCardID)
		} else {
			upIds = append(upIds, nCardDesc.NCardID)
		}
	}

	snm := new(models.ShopNCardModel).Init(ncm.Model.GetOrmer())
	//下架操作, 只处理已经上架的限次卡id
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		snm.Model.Begin()
		if err = ncm.UpdateByNCardIDs(upIds, map[string]interface{}{
			ncm.Field.F_is_ground:     cards.IS_GROUND_no,
			ncm.Field.F_under_time:    time.Now().Unix(),
			ncm.Field.F_sale_shop_num: 0,
		}); err != nil {
			snm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//将分店的限次卡设置为总店禁用
		if err = snm.UpdateByNCardIDs(upIds, map[string]interface{}{
			snm.Field.F_status:     cards.STATUS_DISABLE,
			snm.Field.F_under_time: time.Now().Unix(),
		}); err != nil {
			snm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		snm.Model.Commit()
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, upIds, 0, cards.ITEM_TYPE_ncard)
	}

	//上架操作, 只处理未上架的限次卡id
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		snm.Model.Begin()
		if err = ncm.UpdateByNCardIDs(downIds, map[string]interface{}{
			ncm.Field.F_is_ground:  cards.IS_GROUND_yes,
			ncm.Field.F_under_time: 0,
		}); err != nil {
			snm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//将已添加到子店的且子店适用的限次卡解除总店禁用状态
		//2.解除总店禁用状态
		if err = snm.UpdateByNCardIDs(downIds, map[string]interface{}{snm.Field.F_status: cards.STATUS_OFF_SALE}); err != nil {
			snm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		snm.Model.Commit()
	}
	return
}

//子店获取适用本店的限次卡列表
func (n *NCardLogic) ShopGetBusNCardPage(ctx context.Context, busId, shopId, start, limit int) (list cards.ReplyNCardPage, err error) {
	if busId <= 0 || shopId < 0 || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	nsm := new(models.NCardShopModel).Init()
	nCardShops := nsm.GetPageByShopID(busId, shopId, start, limit)
	if len(nCardShops) == 0 {
		return
	}
	//获取限次卡基本信息
	ncm := new(models.NCardModel).Init()
	snm := new(models.ShopNCardModel).Init()

	nCardIds := functions.ArrayValue2Array(nsm.Field.F_ncard_id, nCardShops)
	nCards := ncm.GetByNCardIDs(nCardIds)
	if len(nCards) == 0 {
		return
	}

	list.TotalNum = nsm.GetNumByShopID(busId, shopId)
	list.List = make([]cards.NCardDesc, len(nCards))
	// 店面已添加限次卡列表
	shopNCards := snm.GetByShopIDAndNCardIDs(shopId, nCardIds)
	shopCardIds := functions.ArrayValue2Array(snm.Field.F_ncard_id, shopNCards)
	for index, nCard := range nCards {
		_ = mapstructure.WeakDecode(nCard, &list.List[index].NCardBase)
		list.List[index].CtimeStr = time.Unix(int64(list.List[index].Ctime), 0).Format("2006/01/02 15:04:05")

		_ = mapstructure.WeakDecode(nCard, &list.List[index])
		list.List[index].Clicks, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.NCARD_CLIKS, nCard[ncm.Field.F_ncard_id].(string)))
		cardId, _ := strconv.Atoi(nCard[ncm.Field.F_ncard_id].(string))
		shopHasAdd := 0
		if functions.InArray(cardId, shopCardIds) {
			shopHasAdd = 1
		}
		list.List[index].ShopHasAdd = shopHasAdd
		for _, shopNCard := range shopNCards {
			nCardID, _ := strconv.ParseInt(shopNCard[snm.Field.F_ncard_id].(string), 10, 64)
			if list.List[index].NCardID == int(nCardID) {
				list.List[index].ShopHasAdd = 1
				status, _ := strconv.ParseInt(shopNCard[snm.Field.F_status].(string), 10, 64)
				list.List[index].ShopStatus = int(status)
			}
		}
	}
	//获取图片信息
	imgIds := functions.ArrayValue2Array(ncm.Field.F_img_id, nCards)
	list.IndexImg = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_ncard)

	return
}

//子店添加限次卡到自己的门店
func (n *NCardLogic) ShopAddNCard(ctx context.Context, busId, shopId int, args *cards.ArgsShopAddNCard) (err error) {
	args.NCardIDs = functions.ArrayUniqueInt(args.NCardIDs)
	if len(args.NCardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//检查限次卡id是否适用当前门店
	//nsm := new(models.NCardShopModel).Init()
	//nCardShop := nsm.GetByShopIDAndNCardIDs(busId, shopId, args.NCardIDs)
	//if len(nCardShop) != len(args.NCardIDs) {
	//	err = toolLib.CreateKcErr(_const.NO_IN_SHOP)
	//	return
	//}

	// 提取门店已经添加过的卡项
	snm := new(models.ShopNCardModel).Init()
	shopNCardLists := snm.GetByShopIdByNCardIDs(shopId,args.NCardIDs)
	shopNCardIds := functions.ArrayValue2Array(snm.Field.F_ncard_id, shopNCardLists)

	// 刷选出已经添加过并且删除的数据
	delNcardIdSlice := make([]int, 0)
	for _, hcardMap := range shopNCardLists {
		isDel, _ := strconv.Atoi(hcardMap[snm.Field.F_is_del].(string))
		if isDel == cards.IS_BUS_DEL_yes {
			delHcardId, _ := strconv.Atoi(hcardMap[snm.Field.F_ncard_id].(string))
			delNcardIdSlice = append(delNcardIdSlice, delHcardId)
		}
	}

	// 更新门店之前添加过并删除的数据
	if len(delNcardIdSlice) > 0 {
		// 更新数据删除和上下架状态
		if err  = snm.UpdateShopIdByNCardIDs(delNcardIdSlice, shopId, map[string]interface{}{
			snm.Field.F_is_del: cards.IS_BUS_DEL_no,
			snm.Field.F_status: cards.STATUS_OFF_SALE,
		}); err != nil  {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
		//更新卡项关联表
		sirModel := new(models.ShopItemRelationModel).Init()
		if b := sirModel.UpdateByItemIdsAndShopId(delNcardIdSlice, cards.ITEM_TYPE_ncard, shopId, map[string]interface{}{
			sirModel.Field.F_is_del: cards.IS_BUS_DEL_no,
			sirModel.Field.F_status: cards.STATUS_OFF_SALE,
		}); !b {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}


	//需要添加的NCardID列表
	addNCardIds := make([]int, 0)
	for _, nCardID := range args.NCardIDs {
		if functions.InArray(nCardID, shopNCardIds) == false {
			addNCardIds = append(addNCardIds, nCardID)
		}
	}

	//校验当前门店是否已经将卡项内涉及到的单项目添加到自己的门店内
	allSingle, singleIds, err := new(ItemLogic).getItemCardIncSingleIds(addNCardIds, cards.ITEM_TYPE_ncard)
	if err != nil {
		return
	}
	if err = new(ItemLogic).validShopSingleContainItemCardSingles(shopId, busId, allSingle, singleIds); err != nil {
		return
	}

	// 需要添加的数据
	var addData []map[string]interface{}
	shopItemRelationData := make([]map[string]interface{}, 0)
	shopItemRelationModel := new(models.ShopItemRelationModel).Init()
	for _, ncardId := range addNCardIds {
		status := cards.STATUS_OFF_SALE
		ctime := time.Now().Local().Unix()
		addData = append(addData, map[string]interface{}{
			snm.Field.F_ncard_id: ncardId,
			snm.Field.F_status:   status,
			snm.Field.F_shop_id:  shopId,
			snm.Field.F_ctime:    ctime,
		})
		shopItemRelationData = append(shopItemRelationData, map[string]interface{}{
			shopItemRelationModel.Field.F_item_id:   ncardId,
			shopItemRelationModel.Field.F_item_type: cards.ITEM_TYPE_ncard,
			shopItemRelationModel.Field.F_status:    cards.STATUS_OFF_SALE,
			shopItemRelationModel.Field.F_shop_id:   shopId,
			shopItemRelationModel.Field.F_is_del:    cards.ITEM_IS_DEL_NO,
		})
	}
	// 过滤的数据添加到门店限次卡表
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
	ncShopModel := new(models.NCardShopModel).Init()
	ncshopSms := ncShopModel.GetByNCardIDs(args.NCardIDs)
	ncshopIds := functions.ArrayValue2Array(ncShopModel.Field.F_ncard_id, ncshopSms)
	addNcShopIds := make([]int, 0)
	for _, ncardId := range args.NCardIDs {
		if functions.InArray(ncardId, ncshopIds) == false {
			addNcShopIds = append(addNcShopIds, ncardId)
		}
	}
	if len(addNcShopIds) == 0 {
		return
	}

	var addncShopData []map[string]interface{} // 添加适用门店表的数据
	for _, ncardId := range addNcShopIds {
		addncShopData = append(addncShopData, map[string]interface{}{
			ncShopModel.Field.F_ncard_id: ncardId,
			ncShopModel.Field.F_shop_id:  shopId,
			ncShopModel.Field.F_bus_id:   busId,
		})
	}
	// 过滤的数据添加到适用限次卡表
	if len(addncShopData) > 0 {
		err = ncShopModel.InsertAll(addncShopData)
		if err != nil {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}

	//获取限次卡在总店的上下架状态
	//ncm := new(models.NCardModel).Init()
	//nCards := ncm.GetByNCardIDs(addNCardIds, ncm.Field.F_ncard_id, ncm.Field.F_is_ground)
	//nCardMap := functions.ArrayRebuild(ncm.Field.F_ncard_id, nCards)
	//var addData []map[string]interface{}
	//for _, nCardID := range addNCardIds {
	//	status := cards.STATUS_OFF_SALE
	//	if _, ok := nCardMap[strconv.Itoa(nCardID)]; ok {
	//		if nCard, ok := nCardMap[strconv.Itoa(nCardID)].(map[string]interface{}); ok {
	//			isGround, _ := strconv.Atoi(nCard[ncm.Field.F_is_ground].(string))
	//			if isGround == cards.IS_GROUND_no {
	//				status = cards.STATUS_DISABLE
	//			}
	//		}
	//	}
	//	ctime := time.Now().Local().Unix()
	//	addData = append(addData, map[string]interface{}{
	//		snm.Field.F_ncard_id: nCardID,
	//		snm.Field.F_status:   status,
	//		snm.Field.F_shop_id:  shopId,
	//		snm.Field.F_ctime:    ctime,
	//	})
	//}
	//
	//if len(addData) > 0 {
	//	if _, err = snm.InsertAll(addData); err != nil {
	//		err = toolLib.CreateKcErr(_const.DB_ERR)
	//		return
	//	}
	//}
	return
}

//获取子店的限次卡列表
func (n *NCardLogic) ShopNCardPage(ctx context.Context, shopId, start, limit, status int) (list cards.ReplyNCardPage, err error) {
	if shopId <= 0 || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	list.List = make([]cards.NCardDesc, 0)
	//获取门店的限次卡数据
	snm := new(models.ShopNCardModel).Init()
	shopNCards := snm.GetPageByShopID(shopId, start, limit, status)
	//shopNCards = SortMapByIntField(snm.Field.F_ncard_id, shopNCards)
	nCardIDs := functions.ArrayValue2Array(snm.Field.F_ncard_id, shopNCards)
	//获取限次卡基本信息
	ncm := new(models.NCardModel).Init()
	nCards := ncm.GetByNCardIDs(nCardIDs)
	if len(nCards) == 0 {
		return
	}
	//获取不同卡项-适用单项目的个数和赠送单项目的个数
	gaagsNumMap := GetApplyAndGiveSingleNum(nCardIDs, cards.ITEM_TYPE_ncard)
	list.List = make([]cards.NCardDesc, len(nCards))
	for index, nCard := range nCards {
		_ = mapstructure.WeakDecode(nCard, &list.List[index])
		_ = mapstructure.WeakDecode(nCard, &list.List[index].NCardBase)
		list.List[index].CtimeStr = time.Unix(int64(list.List[index].Ctime), 0).Format("2006/01/02 15:04:05")
		list.List[index].ApplySingleNum = gaagsNumMap[list.List[index].NCardID].ApplySingleNum
		list.List[index].GiveSingleNum = gaagsNumMap[list.List[index].NCardID].GiveSingleNum
		for _, shopNCard := range shopNCards {
			shopNcardID, _ := strconv.Atoi(shopNCard[snm.Field.F_ncard_id].(string))
			if list.List[index].NCardID == shopNcardID { // 已经添加过的限时卡
				list.List[index].ShopHasAdd = 1
				list.List[index].Clicks, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.HNCARD_CLIKS, strconv.Itoa(list.List[index].NCardID)))
				list.List[index].ShopItemId, _ = strconv.Atoi(shopNCards[index][snm.Field.F_id].(string))
				status, _ := strconv.Atoi(shopNCard[snm.Field.F_status].(string))
				list.List[index].ShopStatus = status
			}
		}
	}
	//获取图片信息
	imgIds := functions.ArrayValue2Array(ncm.Field.F_img_id, nCards)
	list.IndexImg = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_ncard)
	//获取数量
	list.TotalNum = snm.GetNumByShopID(shopId, status)

	return
}

//门店上下架限次卡
func (n *NCardLogic) ShopDownUpNCard(ctx context.Context, shopId int, args *cards.ArgsShopDownUpNCard) (err error) {
	args.CardIDs = functions.ArrayUniqueInt(args.CardIDs)
	if len(args.CardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	//获取门店限次卡信息
	snm := new(models.ShopNCardModel).Init()
	snm.Model.Begin()
	shopNCards := snm.GetByShopIDAndNCardIDs(shopId, args.CardIDs)
	var shopNCardStruct []struct {
		Id      int
		ShopId  int
		Status  int
		NcardId int
	}
	var upIds, downIds, ncardIds []int
	_ = mapstructure.WeakDecode(shopNCards, &shopNCardStruct)
	for _, shopNCardDesc := range shopNCardStruct {
		if shopNCardDesc.Status == cards.STATUS_OFF_SALE {
			downIds = append(downIds, shopNCardDesc.Id)
			ncardIds = append(ncardIds, shopNCardDesc.NcardId)
		} else if shopNCardDesc.Status == cards.STATUS_ON_SALE {
			upIds = append(upIds, shopNCardDesc.Id)
			ncardIds = append(ncardIds, shopNCardDesc.NcardId)
		}
	}

	ncardModel := new(models.NCardModel).Init(snm.Model.GetOrmer())
	var decOrInc string
	//限次卡下架
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
		if !shopItemRealtionModel.UpdateStatusByItemIds(ncardIds, cards.ITEM_TYPE_ncard, cards.STATUS_OFF_SALE, shopId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		decOrInc = "dec"
	}
	//限次卡上架
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
		if !shopItemRealtionModel.UpdateStatusByItemIds(ncardIds, cards.ITEM_TYPE_ncard, cards.STATUS_ON_SALE, shopId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		decOrInc = "inc"
	}

	if len(decOrInc) > 0 {
		//	更新总店中对应现时卡的在售门店数量
		if !ncardModel.UpdateSaleShopNum(ncardIds, decOrInc) {
			snm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	snm.Model.Commit()
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, []int{}, shopId, cards.ITEM_TYPE_ncard, upIds)
	}
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, []int{}, shopId, cards.ITEM_TYPE_ncard, downIds)
	}
	return
}

//ShopNcardRpc
func (n *NCardLogic) ShopNcardRpc(ctx context.Context, params *cards.ArgsShopNcardRpc, list *cards.ReplyShopNcardRpc) (err error) {
	if params.ShopId <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//获取门店的限次卡数据
	snm := new(models.ShopNCardModel).Init()
	shopNCards := snm.GetByShopIDAndNCardIDs(params.ShopId, params.NcardIds)
	//shopNCards = SortMapByIntField(snm.Field.F_ncard_id, shopNCards)
	nCardIDs := functions.ArrayValue2Array(snm.Field.F_ncard_id, shopNCards)
	//获取限次卡基本信息
	ncm := new(models.NCardModel).Init()
	nCards := ncm.GetByNCardIDs(nCardIDs)
	if len(nCards) == 0 {
		return
	}
	list.List = make([]cards.NCardDesc, len(nCards))
	for index, nCard := range nCards {
		_ = mapstructure.WeakDecode(nCard, &list.List[index].NCardBase)
		_ = mapstructure.WeakDecode(nCard, &list.List[index])
	}
	for index := 0; index < len(list.List); index++ {
		_ = mapstructure.WeakDecode(shopNCards[index], &list.List[index])
		list.List[index].ShopHasAdd = 1
	}
	return
}

//总店-删除
func (s *NCardLogic) DeleteNCardLogic(ctx context.Context, args *cards.ArgsDeleteNCard, reply *bool) error {
	//实例化模型
	model := new(models.NCardModel).Init()
	model.Model.Begin()
	//修改数据
	data := map[string]interface{}{
		model.Field.F_is_del:   model.DelStatus(),
		model.Field.F_del_time: time.Now().Unix(),
	}
	if err := model.UpdateByNCardIDs(args.NcardIds, data); err != nil {
		model.Model.RollBack()
		return err
	}
	//实例化分店模型
	shopModel := new(models.ShopNCardModel).Init()

	//修改数据
	if err := shopModel.UpdateByNCardIDs(args.NcardIds, data); err != nil {
		model.Model.RollBack()
		return err
	}
	//同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.NcardIds, cards.ITEM_TYPE_ncard) {
		model.Model.RollBack()
		return toolLib.CreateKcErr(_const.DB_ERR)
	}
	model.Model.Commit()
	return nil
}

//分店-删除
func (s *NCardLogic) DeleteShopNCardLogic(ctx context.Context, args *cards.ArgsDeleteShopNCard, reply *bool) error {
	//实例化分店模型
	shopModel := new(models.ShopNCardModel).Init()
	shopId, _ := args.BsToken.GetShopId()
	data := map[string]interface{}{
		shopModel.Field.F_is_del:   shopModel.DelStatus(),
		shopModel.Field.F_del_time: time.Now().Unix(),
	}
	shopModel.Model.Begin()
	if err := shopModel.UpdateShopIdByNCardIDs(args.NcardIds, shopId, data); err != nil {
		shopModel.Model.RollBack()
		return err
	}
	//同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.NcardIds, cards.ITEM_TYPE_ncard, shopId) {
		shopModel.Model.RollBack()
		return toolLib.CreateKcErr(_const.DB_ERR)
	}
	shopModel.Model.Commit()
	return nil
}
