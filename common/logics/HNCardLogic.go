//限时限次卡业务处理
//@author loop
//@date 2020/4/15 14:55
package logics

import (
	"context"
	"encoding/json"
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
	"strconv"
	"time"
)

type HNCardLogic struct {
}

//添加HNCard描述
func (h *HNCardLogic) AddCardExt(hNcm *models.HNCardModel, hHNCardID int, notes []cards.CardNote) (err error) {
	if len(notes) > constkey.CardNotesMaxNum {
		return toolLib.CreateKcErr(_const.NOTES_NUM_MAX_10)
	}
	for _, note := range notes {
		if functions.Mb4Strlen(note.Notes) > constkey.CardNotesSimpleMaxLength {
			return toolLib.CreateKcErr(_const.NOTES_LEN_MAX_30)
		}
	}
	notesStr, _ := json.Marshal(notes)
	nem := new(models.HNCardExtModel).Init(hNcm.Model.GetOrmer())
	if _, err = nem.Insert(map[string]interface{}{nem.Field.F_hncard_id: hHNCardID, nem.Field.F_notes: string(notesStr)}); err != nil {
		err = toolLib.CreateKcErr(_const.DB_ERR)
	}
	return
}

//编辑HNCard描述
func (h *HNCardLogic) EditCardExt(hNcm *models.HNCardModel, hHNCardID int, notes []cards.CardNote) (err error) {
	if len(notes) > constkey.CardNotesMaxNum {
		return toolLib.CreateKcErr(_const.NOTES_NUM_MAX_10)
	}
	for _, note := range notes {
		if functions.Mb4Strlen(note.Notes) > constkey.CardNotesSimpleMaxLength {
			return toolLib.CreateKcErr(_const.NOTES_LEN_MAX_30)
		}
	}
	notesStr, _ := json.Marshal(notes)
	nem := new(models.HNCardExtModel).Init(hNcm.Model.GetOrmer())

	hHNCardExMap := nem.Find(map[string]interface{}{nem.Field.F_hncard_id: hHNCardID})
	if len(hHNCardExMap) > 0 {
		if _, updateErr := nem.Update(map[string]interface{}{nem.Field.F_hncard_id: hHNCardID},
			map[string]interface{}{nem.Field.F_notes: string(notesStr)}); updateErr != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	} else {
		if _, err = nem.Insert(map[string]interface{}{nem.Field.F_hncard_id: hHNCardID, nem.Field.F_notes: string(notesStr)}); err != nil {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	return
}

//获取HNCard描述
func (h *HNCardLogic) GetCardExt(hHNCardID int, notes *[]cards.CardNote) {
	nem := new(models.HNCardExtModel).Init()
	dataMap := nem.Find(map[string]interface{}{nem.Field.F_hncard_id: hHNCardID})
	if len(dataMap) > 0 {
		data := dataMap[nem.Field.F_notes].(string)
		json.Unmarshal([]byte(data), notes)
	}
	return
}

//添加hncard
func (h *HNCardLogic) AddHNCard(ctx context.Context, busID int, args *cards.ArgsAddHNCard) (hHNCardID int, err error) {
	//验证参数
	err = h.checkHNCardData(busID, args.HNCardBase, args.IncludeSingles, args.GiveSingles)
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
	hNcm := new(models.HNCardModel).Init()
	hNsm := new(models.HNCardSingleModel).Init(hNcm.Model.GetOrmer())

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
	//添加基本信息
	hNcm.Model.Begin()
	if hHNCardID, err = hNcm.Insert(map[string]interface{}{
		hNcm.Field.F_bus_id:          busID,
		hNcm.Field.F_price:           args.Price,
		hNcm.Field.F_is_ground:       cards.SINGLE_IS_GROUND_no,
		hNcm.Field.F_real_price:      args.RealPrice,
		hNcm.Field.F_ctime:           time.Now().Unix(),
		hNcm.Field.F_img_id:          imgID,
		hNcm.Field.F_name:            args.Name,
		hNcm.Field.F_sort_desc:       args.ShortDesc,
		hNcm.Field.F_bind_id:         getBusMainBindId(ctx, busID),
		hNcm.Field.F_has_give_signle: hasGive,
		hNcm.Field.F_service_period:  args.ServicePeriod,
		hNcm.Field.F_validcount:      sum,
	}); err != nil {
		hNcm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//添加单项目
	var hHNCardSingleData []map[string]interface{}
	for _, single := range args.IncludeSingles {
		hHNCardSingleData = append(hHNCardSingleData, map[string]interface{}{
			hNsm.Field.F_single_id: single.SingleID,
			hNsm.Field.F_hncard_id: hHNCardID,
			hNsm.Field.F_ssp_id:    single.SspId,
			hNsm.Field.F_num:       single.Num,
		})
	}
	if err = hNsm.InsertAll(hHNCardSingleData); err != nil {
		hNcm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//添加赠送单项目数据
	if len(args.GiveSingles) > 0 {
		hNgm := new(models.HNCardGiveModel).Init(hNcm.Model.GetOrmer())
		var giveSingleData []map[string]interface{}
		for _, single := range args.GiveSingles {
			giveSingleData = append(giveSingleData, map[string]interface{}{
				hNgm.Field.F_single_id:          single.SingleID,
				hNgm.Field.F_hncard_id:          hHNCardID,
				hNgm.Field.F_num:                single.Num,
				hNgm.Field.F_period_of_validity: single.PeriodOfValidity,
			})
		}
		if err = hNgm.InsertAll(giveSingleData); err != nil {
			hNcm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//添加赠品描述
	if hasGive && len(args.GiveSingleDesc) > 0 {
		hndescm := new(models.HncardGiveDescModel).Init(hNcm.Model.GetOrmer())
		giveSingleDescStr, _ := json.Marshal(args.GiveSingleDesc)
		descData := map[string]interface{}{
			hndescm.Field.F_hncard_id: hHNCardID,
			hndescm.Field.F_desc:      string(giveSingleDescStr),
		}
		if _, err = hndescm.Model.Data(descData).Insert(); err != nil {
			hNcm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//添加HNCardExt
	if err = h.AddCardExt(hNcm, hHNCardID, args.Notes); err != nil {
		hNcm.Model.RollBack()
		return
	}
	hNcm.Model.Commit()
	//添加风控统计任务
	new(ItemLogic).AddXCardTask(ctx, hHNCardID, cards.ITEM_TYPE_hncard)
	return
}

//编辑HNCard数据
func (h *HNCardLogic) EditHNCard(ctx context.Context, busID int, args *cards.ArgsEditHNCard) (err error) {

	nm := new(models.HNCardModel).Init()
	hNsm := new(models.HNCardSingleModel).Init(nm.Model.GetOrmer())

	//验证HNCard数据
	hHNCardInfo := nm.GetByHNCardID(args.CardID, nm.Field.F_bus_id)
	if len(hHNCardInfo) == 0 {
		err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
		return
	}
	if hHNCardInfo[nm.Field.F_bus_id] != strconv.Itoa(busID) {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}

	//验证参数
	err = h.checkHNCardData(busID, args.HNCardBase, args.IncludeSingles, args.GiveSingles)
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
	//计算单项目改动
	includeSingles := hNsm.GetByHNCardID(args.CardID)
	//需要新增的项目
	var addIncSingles []map[string]interface{}
	//需要修改的单项目
	var updateSingles = map[int]int{} //id:num
	//需要删除的ids
	var delIDs []int
	//获取需要修改的项目以及需要添加的项目
out1:
	for _, single := range args.IncludeSingles {
		for _, dbSingle := range includeSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[hNsm.Field.F_single_id].(string) {
				//需要修改次数
				if strconv.Itoa(single.Num) != dbSingle[hNsm.Field.F_num].(string) {
					dbID, _ := strconv.Atoi(dbSingle[hNsm.Field.F_id].(string))
					updateSingles[dbID] = single.Num
				}
				continue out1
			}
		}

		addIncSingles = append(addIncSingles, map[string]interface{}{
			hNsm.Field.F_hncard_id: args.CardID,
			hNsm.Field.F_num:       single.Num,
			hNsm.Field.F_single_id: single.SingleID,
		})
	}
	//获取需要删除的id
out2:
	for _, dbSingle := range includeSingles {
		for _, single := range args.IncludeSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[hNsm.Field.F_single_id].(string) {
				continue out2
			}
		}
		sid, _ := strconv.Atoi(dbSingle[hNsm.Field.F_id].(string))
		delIDs = append(delIDs, sid)
	}

	//计算赠送项目改动
	hNgm := new(models.HNCardGiveModel).Init(nm.Model.GetOrmer())
	giveSingles := hNgm.GetByHNCardID(args.CardID)
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
out3:
	for _, single := range args.GiveSingles {
		for _, dbSingle := range giveSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[hNgm.Field.F_single_id].(string) {
				//需要修改次数
				if strconv.Itoa(single.Num) != dbSingle[hNgm.Field.F_num].(string) {
					dbID, _ := strconv.Atoi(dbSingle[hNgm.Field.F_id].(string))
					updateGiveSingles[dbID] = single.Num
				}
				continue out3
			}
		}

		addGiveSingles = append(addGiveSingles, map[string]interface{}{
			hNgm.Field.F_hncard_id:          args.CardID,
			hNgm.Field.F_num:                single.Num,
			hNgm.Field.F_single_id:          single.SingleID,
			hNgm.Field.F_period_of_validity: single.PeriodOfValidity,
		})
	}
	//获取需要删除的id
out4:
	for _, dbSingle := range giveSingles {
		for _, single := range args.GiveSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[hNgm.Field.F_single_id].(string) {
				continue out4
			}
		}
		sid, _ := strconv.Atoi(dbSingle[hNgm.Field.F_id].(string))
		delGiveIDs = append(delGiveIDs, sid)
	}

	//计算包含单项目总次数
	var sum int = 0
	for _, single := range args.IncludeSingles {
		sum += single.Num
	}
	//修改主表信息
	nm.Model.Begin()
	if err = nm.UpdateByHNCardID(args.CardID, map[string]interface{}{
		nm.Field.F_price:           args.Price,
		nm.Field.F_real_price:      args.RealPrice,
		nm.Field.F_img_id:          imgID,
		nm.Field.F_name:            args.Name,
		nm.Field.F_sort_desc:       args.ShortDesc,
		nm.Field.F_has_give_signle: hasGive,
		nm.Field.F_service_period:  args.ServicePeriod,
		nm.Field.F_validcount:      sum,
	}); err != nil {
		nm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	//修改包含的单项
	//1. 修改需要修改的项目
	if len(updateSingles) > 0 {
		for id, num := range updateSingles {
			if err = hNsm.UpdateNumById(id, num); err != nil {
				nm.Model.RollBack()
				err = toolLib.CreateKcErr(_const.DB_ERR)
				return
			}
		}
	}
	//2. 添加需要添加的项目
	if len(addIncSingles) > 0 {
		if err = hNsm.InsertAll(addIncSingles); err != nil {
			nm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//3. 删除需要删除的项目
	if len(delIDs) > 0 {
		if err = hNsm.DelByIds(delIDs); err != nil {
			nm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//修改赠送的单项目
	//1. 修改需要修改的项目
	if len(updateGiveSingles) > 0 {
		for id, num := range updateGiveSingles {
			if err = hNgm.UpdateNumById(id, num); err != nil {
				nm.Model.RollBack()
				err = toolLib.CreateKcErr(_const.DB_ERR)
				return
			}
		}
	}
	//2. 添加需要添加的项目
	if len(addGiveSingles) > 0 {
		if err = hNgm.InsertAll(addGiveSingles); err != nil {
			nm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//3. 删除需要删除的项目
	if len(delGiveIDs) > 0 {
		if err = hNgm.DelByIds(delGiveIDs); err != nil {
			nm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//修改赠品描述
	hngdescm := new(models.HncardGiveDescModel).Init(nm.Model.GetOrmer())
	//1.删除原有赠品描述
	if err = hngdescm.DelByHncardId(args.CardID); err != nil {
		nm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//2.新增赠品描述
	if len(args.GiveSingles) > 0 && len(args.GiveSingleDesc) > 0 {
		giveSingleDescStr, _ := json.Marshal(args.GiveSingleDesc)
		descData := map[string]interface{}{
			hngdescm.Field.F_hncard_id: args.CardID,
			hngdescm.Field.F_desc:      string(giveSingleDescStr),
		}
		if _, err = hngdescm.Model.Data(descData).Insert(); err != nil {
			nm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	// update cardExt
	if err = h.EditCardExt(nm, args.CardID, args.Notes); err != nil {
		nm.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	nm.Model.Commit()

	return
}

//检查限时限次卡的入参数据
func (h *HNCardLogic) checkHNCardData(busID int, hHNCardBase cards.HNCardBase, incSingles []cards.IncSingle, giveSingles []cards.IncSingle) (err error) {
	if err = cards.VerfiyName(hHNCardBase.Name); err != nil {
		return
	}
	if err = cards.VerfiyPrice(hHNCardBase.RealPrice, hHNCardBase.Price); err != nil {
		return
	}
	if err = cards.VerfiyServicePeriod(hHNCardBase.ServicePeriod); err != nil {
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
	for _, single := range incSingles {
		if single.Num <= 0 {
			err = toolLib.CreateKcErr(_const.PARAM_ERR)
			return
		}
		singleIDs = append(singleIDs, single.SingleID)
	}
	if err = checkSingles(busID, singleIDs); err != nil {
		return
	}

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

	//验证规格
	//var sspIds []int
	//for _, single := range incSingles {
	//	sspIds = append(sspIds, single.SspId)
	//}
	//if err = h.CheckSspIds(sspIds, singleIDs, incSingles); err != nil {
	//	return
	//}

	return
}

//获取限时限次卡的详情
//@param int hHNCardID 限时限次卡id
//@param int shopID 门店id
//@return  cards.ReplyHNCardInfo reply
func (h *HNCardLogic) HNCardInfo(ctx context.Context, hHNCardID int, shopID ...int) (reply cards.ReplyHNCardInfo, err error) {
	reply.IncludeSingles = make([]cards.IncSingleDetail2, 0)
	reply.GiveSingles = make([]cards.IncSingleDetail2, 0)
	reply.GiveSingleDesc=make([]cards.GiveSingleDesc,0)
	hNcm := new(models.HNCardModel).Init()
	hHNCardInfo := hNcm.GetByHNCardID(hHNCardID)
	if len(hHNCardInfo) == 0 {
		err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
		return
	}
	var shopStatus = 0
	var shopSales = 0
	var shopid = 0
	var shopIsDel = 0
	if len(shopID) > 0 && shopID[0] > 0 {
		shm := new(models.ShopHNCardModel).Init()
		shopNcInfo := shm.GetByShopIDAndHNCardIDs(shopID[0], []int{hHNCardID})
		if len(shopNcInfo) == 0 {
			err = toolLib.CreateKcErr(_const.NO_IN_SHOP)
			return
		}
		shopStatus, _ = strconv.Atoi(shopNcInfo[0][shm.Field.F_status].(string))
		reply.SsId, _ = strconv.Atoi(shopNcInfo[0][shm.Field.F_id].(string))
		shopSales, _ = strconv.Atoi(shopNcInfo[0][shm.Field.F_sales].(string))
		shopIsDel, _ = strconv.Atoi(shopNcInfo[0][shm.Field.F_is_del].(string))
		shopid = shopID[0]
	}
	imgId, _ := strconv.Atoi(hHNCardInfo[hNcm.Field.F_img_id].(string))
	imgHash, imgUrl := getImg(ctx, imgId, cards.ITEM_TYPE_hncard)
	reply.ImgHash = imgHash
	reply.ImgUrl = imgUrl
	reply.ShopStatus = shopStatus
	reply.ShareLink = tools.GetShareLink(hHNCardID, shopid, cards.ITEM_TYPE_hncard)
	_ = mapstructure.WeakDecode(hHNCardInfo, &reply.HNCardBase)
	reply.CtimeStr = time.Unix(int64(reply.Ctime), 0).Format("2006/01/02 15:04:05")
	_ = mapstructure.WeakDecode(hHNCardInfo, &reply)
	if len(shopID) > 0 && shopID[0] > 0 {
		reply.Sales = shopSales
		reply.IsDel = shopIsDel
	}

	//商户信息
	if err = getBusInfo(ctx, reply.BusID, &reply.BusInfo); err != nil {
		err = toolLib.CreateKcErr(_const.SHOP_INFO_ERR)
		return
	}

	//.....
	//统计限时限次卡包含项目的总次数
	reply.SingleTotalNum = GetItemCardSingleNum(hHNCardID, cards.ITEM_TYPE_hncard)

	//获取包含的限时限次卡和赠送项目信息
	hNsm := new(models.HNCardSingleModel).Init()
	hHNCardSingles := hNsm.GetByHNCardID(hHNCardID)
	if len(hHNCardSingles) > 0 && hHNCardSingles[0][hNsm.Field.F_single_id].(string) == "0" {
		reply.IsAllSingle = true
	}
	//singleIds := functions.ArrayValue2Array(hNsm.Field.F_single_id, hHNCardSingles)
	var singleIds []int
	hNgm := new(models.HNCardGiveModel).Init()
	var hHNCardGives []map[string]interface{}
	if hHNCardInfo[hNcm.Field.F_has_give_signle].(string) == strconv.Itoa(cards.HAS_GIVE_SINGLE_yes) {
		hHNCardGives = hNgm.GetByHNCardID(hHNCardID)
		giveSingleIds := functions.ArrayValue2Array(hNgm.Field.F_single_id, hHNCardGives)
		singleIds = append(singleIds, giveSingleIds...)
	}
	allSingles,_ := getIncSingles2(ctx,shopid,"","", singleIds, hHNCardGives)

	if len(hHNCardGives) > 0 {
		for _, single := range hHNCardGives {
			for i := range allSingles {
				if single[hNgm.Field.F_single_id].(string) == strconv.Itoa(allSingles[i].SingleID) {
					allSingles[i].Num, _ = strconv.Atoi(single[hNgm.Field.F_num].(string))
					allSingles[i].PeriodOfValidity, _ = strconv.Atoi(single[hNgm.Field.F_period_of_validity].(string))
					reply.GiveSingles = append(reply.GiveSingles, allSingles[i])
					break
				}
			}
		}
		//获取赠品描述信息
		hngdescm := new(models.HncardGiveDescModel).Init()
		desc, ok := hngdescm.GetByHncardId(hHNCardID)[hngdescm.Field.F_desc].(string)
		if ok {
			json.Unmarshal([]byte(desc), &reply.GiveSingleDesc)
		}
	}

	//获取CardExt
	h.GetCardExt(hHNCardID, &reply.Notes)

	// 获取限时限次卡门店添加详情  15 -- []int {3,4,6}
	busId, _ := strconv.Atoi(hHNCardInfo[hNcm.Field.F_bus_id].(string))
	hNModel := new(models.HNCardShopModel).Init()
	hNLists := hNModel.GetByHNCardIdAndBusId(hHNCardID, busId)

	hNShopIds := make([]int, 0)
	for _, hNInfoValue := range hNLists {
		sshopId, _ := strconv.Atoi(hNInfoValue[hNModel.Field.F_shop_id].(string))
		hNShopIds = append(hNShopIds, sshopId)
	}
	var replyShop []bus2.ReplyShopName
	rLists := make([]cards.ReplyShopName, 0)
	rpcBus := new(bus.Shop).Init()
	defer rpcBus.Close()
	err = rpcBus.GetShopNameByShopIds(ctx, &hNShopIds, &replyShop)
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
	_ = redis.RedisGlobMgr.Hincrby(constkey.HNCARD_CLIKS, strconv.Itoa(hHNCardID), 1)
	return
}

//获取商家的限时限次卡列表
func (h *HNCardLogic) GetBusPage(ctx context.Context, busId, shopId, start, limit int, isGround string, filterShopHasAdd bool) (list cards.ReplyHNCardPage, err error) {

	if busId <= 0 || start < 0 || limit < 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	hNcm := new(models.HNCardModel).Init()
	hHNCards := make([]map[string]interface{}, 0)
	list.List = make([]cards.HNCardDesc, 0)
	list.IndexImg = make(map[int]string)
	////获取总数量 (
	//if isGround == "" {
	//	hHNCards = hNcm.GetPageByBusID(busId, start, limit)
	//	list.TotalNum = hNcm.GetNumByBusID(busId)
	//} else {
	//	isground, _ := strconv.Atoi(isGround)
	//	isground = isground - 1
	//	hHNCards = hNcm.GetPageByBusID(busId, start, limit, isground)
	//	list.TotalNum = hNcm.GetNumByBusID(busId, isground)
	//}
	//子店已添加的卡项
	var shopAddCards []map[string]interface{}
	scModel := new(models.ShopHNCardModel).Init()
	where := make([]base.WhereItem, 0)
	where = append(where, base.WhereItem{hNcm.Field.F_bus_id, busId})
	where = append(where, base.WhereItem{hNcm.Field.F_is_del, cards.IS_BUS_DEL_no})

	if shopId > 0 {
		shopAddCards = scModel.SelectRcardsByWherePage([]base.WhereItem{{scModel.Field.F_shop_id, shopId}, {scModel.Field.F_is_del, cards.IS_BUS_DEL_no}}, 0, 0)
		if filterShopHasAdd && len(shopAddCards) > 0 {
			shopHasAddHncardIds := functions.ArrayValue2Array(scModel.Field.F_hncard_id, shopAddCards)
			where = append(where, base.WhereItem{hNcm.Field.F_hncard_id, []interface{}{"NOT IN", shopHasAddHncardIds}})
		}
	}

	////获取总数量
	hHNCards = hNcm.SelectHNCardsByWherePage(where, start, limit)
	list.TotalNum = hNcm.GetNumByWhere(where)

	if len(hHNCards) == 0 {
		return
	}

	list.List = make([]cards.HNCardDesc, len(hHNCards))
	for index, nc := range hHNCards {
		_ = mapstructure.WeakDecode(nc, &list.List[index].HNCardBase)
		for _, shopCard := range shopAddCards {
			if nc[hNcm.Field.F_hncard_id].(string) == shopCard[scModel.Field.F_hncard_id].(string) { //表明当前子店已添加该卡项
				nc["ShopItemId"] = shopCard[scModel.Field.F_id].(string)
				nc["ShopStatus"] = shopCard[scModel.Field.F_status].(string)
				nc["ShopHasAdd"] = 1
				nc["ShopDelStatus"] = shopCard[scModel.Field.F_is_del].(string)
				break
			}
		}
		list.List[index].CtimeStr = time.Unix(int64(list.List[index].Ctime), 0).Format("2006/01/02 15:04:05")
		_ = mapstructure.WeakDecode(nc, &list.List[index])
		list.List[index].Clicks, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.HNCARD_CLIKS, nc[hNcm.Field.F_hncard_id].(string)))
	}
	//获取图片信息
	imgIds := functions.ArrayValue2Array(hNcm.Field.F_img_id, hHNCards)
	list.IndexImg = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_hncard)

	return
}

//设置限时限次卡的适用门店（废用）
func (h *HNCardLogic) SetHNCardShop(ctx context.Context, busId int, args *cards.ArgsSetHNCardShop) (err error) {
	if len(args.HNCardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.CARD_ID_IS_NIL)
		return
	}
	if args.IsAllShop == false && len(args.ShopIDs) == 0 {
		err = toolLib.CreateKcErr(_const.SHOPID_NTL)
		return
	}

	//限次卡id重复提交判断
	realHNCardIDs := functions.ArrayUniqueInt(args.HNCardIDs)
	if len(realHNCardIDs) != len(args.HNCardIDs) {
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
	hcm := new(models.HNCardModel).Init()
	dataArr := hcm.GetByHNCardIDs(realHNCardIDs)
	if len(dataArr) != len(realHNCardIDs) {
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
	// 适用门店
	hsm := new(models.HNCardShopModel).Init()
	var insertData []map[string]interface{}
	//走全部适用逻辑
	if args.IsAllShop == true {
		for _, hHNCardID := range realHNCardIDs {
			insertData = append(insertData, map[string]interface{}{
				hsm.Field.F_hncard_id: hHNCardID,
				hsm.Field.F_bus_id:    busId,
				hsm.Field.F_shop_id:   0,
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
		for _, hHNCardID := range realHNCardIDs {
			for _, shopID := range realShopIds {
				insertData = append(insertData, map[string]interface{}{
					hsm.Field.F_hncard_id: hHNCardID,
					hsm.Field.F_bus_id:    busId,
					hsm.Field.F_shop_id:   shopID,
				})
			}
		}
	}

	//处理数据
	if len(insertData) > 0 {
		hsm.Model.Begin()
		if err = hsm.DelByHNCardIDs(realHNCardIDs); err != nil {
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
	var downIDs []int

	// 更新添加到门店表
	shm := new(models.ShopHNCardModel).Init()
	shopHNCards := shm.GetByHNCardIDs(realHNCardIDs)
	if args.IsAllShop {
		for _, hcard := range dataArr {
			//hcardStatus, _ := strconv.Atoi(hcard[hcm.Field.F_is_ground].(string)) // 总店中卡的状态
			isDel, _ := strconv.Atoi(hcard[hcm.Field.F_is_del].(string))
			hcardID, _ := strconv.Atoi(hcard[hcm.Field.F_hncard_id].(string))
			for _, shopHcard := range shopHNCards {
				shopStatus, _ := strconv.Atoi(shopHcard[shm.Field.F_status].(string)) // 子店中卡的状态
				id, _ := strconv.Atoi(shopHcard[shm.Field.F_id].(string))
				shopHcardID, _ := strconv.Atoi(shopHcard[shm.Field.F_hncard_id].(string)) // 子店中卡的id
				if hcardID != shopHcardID {
					continue
				}
				//if shopStatus == cards.STATUS_DISABLE && hcardStatus == cards.SINGLE_IS_GROUND_yes { // 子店中卡的状态为禁用并且总店卡的状态为上架时,子店才可以上架
				//	downIDs = append(downIDs, id)
				//}
				// 一期优化----yinjinlin-2021-04-08
				// 全部适用：之前门店已经添加过并且： 1>.上架/或下架 状态 ：都不需要更新
				//                              2>.被总店禁用 状态 ：需要更新为下架状态
				//分店中没有被删除，子店可以选择上架
				if isDel == cards.IS_BUS_DEL_no && shopStatus == cards.STATUS_DISABLE {
					downIDs = append(downIDs, id)
				}
			}
		}
	} else {
		for _, hcard := range dataArr {
			//hcardStatus, _ := strconv.Atoi(hcard[hcm.Field.F_is_ground].(string)) // 总店中卡的状态
			isDel, _ := strconv.Atoi(hcard[hcm.Field.F_is_del].(string))
			hcardID, _ := strconv.Atoi(hcard[hcm.Field.F_hncard_id].(string))
			for _, shopHcard := range shopHNCards {
				shopStatus, _ := strconv.Atoi(shopHcard[shm.Field.F_status].(string)) // 子店中卡的状态
				id, _ := strconv.Atoi(shopHcard[shm.Field.F_id].(string))
				//shopID, _ := strconv.Atoi(shopHcard[shm.Field.F_shop_id].(string))        // 已经存在的子店id
				shopHcardID, _ := strconv.Atoi(shopHcard[shm.Field.F_hncard_id].(string)) // 子店中卡的id
				if hcardID != shopHcardID {
					continue
				}
				//if functions.InArray(shopID, realShopIds) { // 以前已经添加过的子店并且status="总店禁用";现在适用时需要将status恢复为下架;前提为总店卡的状态为上架时
				//	if hcardStatus == cards.SINGLE_IS_GROUND_yes && shopStatus == cards.STATUS_DISABLE {
				//		downIDs = append(downIDs, id)
				//	}
				//} else { // 以前已经添加过的子店并且status为"上架"或"下架";如果现在不再适用现时卡时，status需更改为"总店禁用"
				//	if shopStatus != cards.STATUS_DISABLE {
				//		disableIds = append(disableIds, id)
				//	}
				//}

				// 部分适用
				// 以前添加过到门店，并且门店现在是上架/下架 状态，总店设置不适用了，需要将门店状态改为：被总店禁用
				if shopStatus != cards.STATUS_DISABLE && isDel == cards.IS_BUS_DEL_no {
					disableIds = append(disableIds, id)
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
		setShopItem(ctx, []int{}, 0, cards.ITEM_TYPE_hncard, disableIds)
	}
	if len(downIDs) > 0 {
		_ = shm.UpdateByIDs(downIDs, map[string]interface{}{
			shm.Field.F_status: cards.STATUS_OFF_SALE,
		})
	}
	return
}

//总店上下架限时限次卡 (一期优化废用）
func (h *HNCardLogic) DownUpHNCard(ctx context.Context, busId int, args *cards.ArgsDownUpHNCard) (err error) {

	// 限时限次卡ID列表交验
	if len(args.HNCardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	realHNCardIDs := functions.ArrayUniqueInt(args.HNCardIDs)
	//交验是否有重复的id
	if len(realHNCardIDs) != len(args.HNCardIDs) {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//检查限时限次卡是否属于企业
	hNcm := new(models.HNCardModel).Init()
	dataArr := hNcm.GetByHNCardIDs(realHNCardIDs, hNcm.Field.F_bus_id, hNcm.Field.F_hncard_id, hNcm.Field.F_is_ground)
	busIdStr := strconv.Itoa(busId)
	for _, data := range dataArr {
		if busIdStr != data[hNcm.Field.F_bus_id].(string) {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
	}
	//将id中已经上架的限时限次卡id和未上架的限时限次卡id分开处理
	var hHNCardDescList []struct {
		HNCardID int `mapstructure:"hncard_id"`
		IsGround int `mapstructure:"is_ground"`
	}
	_ = mapstructure.WeakDecode(dataArr, &hHNCardDescList)
	var downIds, upIds []int
	for _, hHNCardDesc := range hHNCardDescList {
		if hHNCardDesc.IsGround == cards.IS_GROUND_no {
			downIds = append(downIds, hHNCardDesc.HNCardID)
		} else {
			upIds = append(upIds, hHNCardDesc.HNCardID)
		}
	}

	shm := new(models.ShopHNCardModel).Init(hNcm.Model.GetOrmer())
	//下架操作, 只处理已经上架的限时限次卡id
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		shm.Model.Begin()
		if err = hNcm.UpdateByHNCardIDs(upIds, map[string]interface{}{
			hNcm.Field.F_is_ground:     cards.IS_GROUND_no,
			hNcm.Field.F_under_time:    time.Now().Unix(),
			hNcm.Field.F_sale_shop_num: 0,
		}); err != nil {
			shm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//将分店的限时限次卡设置为总店禁用
		if err = shm.UpdateByHNCardIDs(upIds, map[string]interface{}{
			shm.Field.F_status:     cards.STATUS_DISABLE,
			shm.Field.F_under_time: time.Now().Unix(),
		}); err != nil {
			shm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		shm.Model.Commit()
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, upIds, 0, cards.ITEM_TYPE_hncard)
	}

	//上架操作, 只处理未上架的限时限次卡id
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		shm.Model.Begin()
		if err = hNcm.UpdateByHNCardIDs(downIds, map[string]interface{}{
			hNcm.Field.F_is_ground:  cards.IS_GROUND_yes,
			hNcm.Field.F_under_time: 0,
		}); err != nil {
			shm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//将已添加到子店的且子店适用的限时限次卡解除总店禁用状态
		//2.解除总店禁用状态
		if err = shm.UpdateByHNCardIDs(downIds, map[string]interface{}{
			shm.Field.F_status: cards.STATUS_OFF_SALE,
		}); err != nil {
			shm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		shm.Model.Commit()
	}
	return
}

//子店获取适用本店的限时限次卡列表
func (h *HNCardLogic) ShopGetBusHNCardPage(ctx context.Context, busId, shopId, start, limit int) (list cards.ReplyHNCardPage, err error) {
	if busId <= 0 || shopId < 0 || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	hNsm := new(models.HNCardShopModel).Init()
	hHNCardShops := hNsm.GetPageByShopID(busId, shopId, start, limit)
	if len(hHNCardShops) == 0 {
		return
	}
	//获取限时限次卡基本信息
	hNcm := new(models.HNCardModel).Init()
	shm := new(models.ShopHNCardModel).Init()

	hHNCardIds := functions.ArrayValue2Array(hNsm.Field.F_hncard_id, hHNCardShops)
	hHNCards := hNcm.GetByHNCardIDs(hHNCardIds)
	if len(hHNCards) == 0 {
		return
	}

	list.TotalNum = hNsm.GetNumByShopID(busId, shopId)
	list.List = make([]cards.HNCardDesc, len(hHNCards))
	// 店面已添加限时限次卡列表
	shopHNCards := shm.GetByShopIDAndHNCardIDs(shopId, hHNCardIds)
	shopCardIds := functions.ArrayValue2Array(shm.Field.F_hncard_id, shopHNCards)
	for index, hHNCard := range hHNCards {
		_ = mapstructure.WeakDecode(hHNCard, &list.List[index].HNCardBase)
		list.List[index].CtimeStr = time.Unix(int64(list.List[index].Ctime), 0).Format("2006/01/02 15:04:05")
		_ = mapstructure.WeakDecode(hHNCard, &list.List[index])
		list.List[index].Clicks, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.HNCARD_CLIKS, hHNCard[hNcm.Field.F_hncard_id].(string)))
		cardId, _ := strconv.Atoi(hHNCard[hNcm.Field.F_hncard_id].(string))
		shopHasAdd := 0
		if functions.InArray(cardId, shopCardIds) {
			shopHasAdd = 1
		}
		list.List[index].ShopHasAdd = shopHasAdd
		for _, shopHNCard := range shopHNCards {
			hHNCardID, _ := strconv.ParseInt(shopHNCard[shm.Field.F_hncard_id].(string), 10, 64)
			if list.List[index].HNCardID == int(hHNCardID) {
				list.List[index].ShopHasAdd = 1
				status, _ := strconv.ParseInt(shopHNCard[shm.Field.F_status].(string), 10, 64)
				list.List[index].ShopStatus = int(status)
			}
		}
	}
	//获取图片信息
	imgIds := functions.ArrayValue2Array(hNcm.Field.F_img_id, hHNCards)
	list.IndexImg = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_hncard)

	return
}

//子店添加限时限次卡到自己的门店
func (h *HNCardLogic) ShopAddHNCard(ctx context.Context, busId, shopId int, args *cards.ArgsShopAddHNCard) (err error) {
	args.HNCardIDs = functions.ArrayUniqueInt(args.HNCardIDs)
	if len(args.HNCardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//检查限时限次卡id是否适用当前门店
	//hNsm := new(models.HNCardShopModel).Init()
	//hHNCardShop := hNsm.GetByShopIDAndHNCardIDs(busId, shopId, args.HNCardIDs)
	//if len(hHNCardShop) != len(args.HNCardIDs) {
	//	err = toolLib.CreateKcErr(_const.NO_IN_SHOP)
	//	return
	//}

	//提取门店已经添加过的卡项
	shm := new(models.ShopHNCardModel).Init()
	shopHNCardLists := shm.GetByShopIDAndHNCardIDs(shopId, args.HNCardIDs)
	shopHNCardIds := functions.ArrayValue2Array(shm.Field.F_hncard_id, shopHNCardLists)

	// 刷选出已经添加过并且删除的数据
	delHNcardIdSlice := make([]int, 0)
	for _, hcardMap := range shopHNCardLists {
		isDel, _ := strconv.Atoi(hcardMap[shm.Field.F_is_del].(string))
		if isDel == cards.IS_BUS_DEL_yes {
			delHncardId, _ := strconv.Atoi(hcardMap[shm.Field.F_hncard_id].(string))
			delHNcardIdSlice = append(delHNcardIdSlice, delHncardId)
		}
	}

	// 更新门店之前添加过并删除的数据
	if len(delHNcardIdSlice) > 0 {
		// 更新数据删除和上下架状态
		err = shm.UpdateByHNCardIDsAndShopId(delHNcardIdSlice,shopId, map[string]interface{}{
			shm.Field.F_is_del: cards.IS_BUS_DEL_no,
			shm.Field.F_status: cards.STATUS_OFF_SALE,
		})
		if err != nil  {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
		// 更新卡项关联表
		sirModel := new(models.ShopItemRelationModel).Init()
		if b := sirModel.UpdateByItemIdsAndShopId(delHNcardIdSlice, cards.ITEM_TYPE_hncard, shopId, map[string]interface{}{
			sirModel.Field.F_is_del: cards.IS_BUS_DEL_no,
			sirModel.Field.F_status: cards.STATUS_OFF_SALE,
		}); !b {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}

	//需要添加的HNCardID列表
	addHNCardIds := make([]int, 0)
	for _, hHNCardID := range args.HNCardIDs {
		if functions.InArray(hHNCardID, shopHNCardIds) == false {
			addHNCardIds = append(addHNCardIds, hHNCardID)
		}
	}

	//校验当前门店是否已经将卡项内涉及到的单项目添加到自己的门店内
	allSingle, singleIds, err := new(ItemLogic).getItemCardIncSingleIds(addHNCardIds, cards.ITEM_TYPE_hncard)
	if err != nil {
		return
	}
	if err = new(ItemLogic).validShopSingleContainItemCardSingles(shopId, busId, allSingle, singleIds); err != nil {
		return
	}

	var addData []map[string]interface{}
	shopItemRelationData := make([]map[string]interface{}, 0)
	shopItemRelationModel := new(models.ShopItemRelationModel).Init()
	for _, hncardId := range addHNCardIds {
		status := cards.STATUS_OFF_SALE
		ctime := time.Now().Local().Unix()
		addData = append(addData, map[string]interface{}{
			shm.Field.F_hncard_id: hncardId,
			shm.Field.F_status:    status,
			shm.Field.F_shop_id:   shopId,
			shm.Field.F_ctime:     ctime,
		})
		shopItemRelationData = append(shopItemRelationData, map[string]interface{}{
			shopItemRelationModel.Field.F_item_id:   hncardId,
			shopItemRelationModel.Field.F_item_type: cards.ITEM_TYPE_hncard,
			shopItemRelationModel.Field.F_status:    cards.STATUS_OFF_SALE,
			shopItemRelationModel.Field.F_shop_id:   shopId,
			shopItemRelationModel.Field.F_is_del:    cards.ITEM_IS_DEL_NO,
		})
	}
	// 过滤的数据添加到门店限时限次卡表
	if len(addData) > 0 {
		id, _ := shm.InsertAll(addData)
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
	hnShopModel := new(models.HcardShopModel).Init()
	hncardShopLists := hnShopModel.GetByHcardIDs(args.HNCardIDs)
	hncardShopIds := functions.ArrayValue2Array(hnShopModel.Field.F_hcard_id, hncardShopLists)
	addhncardShopIds := make([]int, 0)
	for _, hcardId := range args.HNCardIDs {
		if functions.InArray(hcardId, hncardShopIds) == false {
			addhncardShopIds = append(addhncardShopIds, hcardId)
		}
	}

	var addhnShopData []map[string]interface{} // 添加适用门店表的数据
	for _, hncardId := range addhncardShopIds {
		addhnShopData = append(addhnShopData, map[string]interface{}{
			hnShopModel.Field.F_hcard_id: hncardId,
			hnShopModel.Field.F_shop_id:  shopId,
			hnShopModel.Field.F_bus_id:   busId,
		})
	}
	// 过滤的数据添加到适用限次卡表
	if len(addhnShopData) > 0 {
		if id := hnShopModel.InsertAll(addhnShopData); id < 0 {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}

	//获取限时限次卡在总店的上下架状态
	//hNcm := new(models.HNCardModel).Init()
	////hHNCards := hNcm.GetByHNCardIDs(addHNCardIds, hNcm.Field.F_hncard_id, hNcm.Field.F_is_ground)
	//hHNCards := hNcm.GetByHNCardIDs(addHNCardIds, hNcm.Field.F_hncard_id, hNcm.Field.F_is_del)
	//hHNCardMap := functions.ArrayRebuild(hNcm.Field.F_hncard_id, hHNCards)
	//var addData []map[string]interface{}
	//for _, hHNCardID := range addHNCardIds {
	//	status := cards.STATUS_OFF_SALE
	//	if _, ok := hHNCardMap[strconv.Itoa(hHNCardID)]; ok {
	//		if hHNCard, ok := hHNCardMap[strconv.Itoa(hHNCardID)].(map[string]interface{}); ok {
	//			//isGround, _ := strconv.Atoi(hHNCard[hNcm.Field.F_is_ground].(string))
	//			//if isGround == cards.IS_GROUND_no {
	//			//	status = cards.STATUS_DISABLE
	//			//}
	//			isDel, _ := strconv.Atoi(hHNCard[hNcm.Field.F_is_del].(string))
	//			if isDel == cards.IS_BUS_DEL_no {
	//				status = cards.STATUS_OFF_SALE
	//			}
	//		}
	//	}
	//	ctime := time.Now().Local().Unix()
	//	addData = append(addData, map[string]interface{}{
	//		shm.Field.F_hncard_id: hHNCardID,
	//		shm.Field.F_status:    status,
	//		shm.Field.F_shop_id:   shopId,
	//		shm.Field.F_ctime:     ctime,
	//	})
	//}
	//
	//if len(addData) > 0 {
	//	if _, err = shm.InsertAll(addData); err != nil {
	//		err = toolLib.CreateKcErr(_const.DB_ERR)
	//		return
	//	}
	//}

	return
}

func SortMapByIntField(field string, maps []map[string]interface{}) []map[string]interface{} {
	for i := 1; i < len(maps); i++ {
		for j := 0; j < len(maps)-i; j++ {
			prev, _ := strconv.Atoi(maps[j][field].(string))
			next, _ := strconv.Atoi(maps[j+1][field].(string))
			if prev > next {
				temp := maps[j][field]
				maps[j][field] = maps[j+1][field]
				maps[j+1][field] = temp
			}
		}
	}
	return maps
}

//获取子店的限时限次卡列表
func (h *HNCardLogic) ShopHNCardPage(ctx context.Context, shopId, start, limit, status int) (list cards.ReplyHNCardPage, err error) {
	if shopId <= 0 || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	//获取门店的限时限次卡数据 (已经过滤掉删除的限时限次卡）
	shm := new(models.ShopHNCardModel).Init()
	shopHNCards := shm.GetPageByShopID(shopId, start, limit, status)
	//shopHNCards = SortMapByIntField(shm.Field.F_hncard_id, shopHNCards)
	hHNCardIDs := functions.ArrayValue2Array(shm.Field.F_hncard_id, shopHNCards)
	list.List = make([]cards.HNCardDesc, 0)
	//获取限时限次卡基本信息
	hNcm := new(models.HNCardModel).Init()
	hHNCards := hNcm.GetByHNCardIDs(hHNCardIDs)
	if len(hHNCards) == 0 {
		return
	}
	//获取不同卡项-适用单项目的个数和赠送单项目的个数
	gaagsNumMap := GetApplyAndGiveSingleNum(hHNCardIDs, cards.ITEM_TYPE_hncard)
	list.List = make([]cards.HNCardDesc, len(hHNCards))
	for index, hHNCard := range hHNCards {
		_ = mapstructure.WeakDecode(hHNCard, &list.List[index])
		_ = mapstructure.WeakDecode(hHNCard, &list.List[index].HNCardBase)
		list.List[index].CtimeStr = time.Unix(int64(list.List[index].Ctime), 0).Format("2006/01/02 15:04:05")
		list.List[index].ApplySingleNum = gaagsNumMap[list.List[index].HNCardID].ApplySingleNum
		list.List[index].GiveSingleNum = gaagsNumMap[list.List[index].HNCardID].GiveSingleNum
		for _, shopHNCard := range shopHNCards {
			shopHcardID, _ := strconv.Atoi(shopHNCard[shm.Field.F_hncard_id].(string))
			if list.List[index].HNCardID == shopHcardID { // 已经添加过的限时卡
				list.List[index].ShopHasAdd = 1
				list.List[index].Clicks, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.HNCARD_CLIKS, strconv.Itoa(list.List[index].HNCardID)))
				list.List[index].ShopItemId, _ = strconv.Atoi(shopHNCards[index][shm.Field.F_id].(string))
				status, _ := strconv.Atoi(shopHNCard[shm.Field.F_status].(string))
				list.List[index].ShopStatus = status
			}
		}
	}
	//获取图片信息
	imgIds := functions.ArrayValue2Array(hNcm.Field.F_img_id, hHNCards)
	list.IndexImg = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_hncard)
	//获取数量(已经过滤掉删除的限时限次卡)
	list.TotalNum = shm.GetNumByShopID(shopId, status)

	return
}

//门店上下架限时限次卡
func (h *HNCardLogic) ShopDownUpHNCard(ctx context.Context, shopId int, args *cards.ArgsShopDownUpHNCard) (err error) {

	args.CardIDs = functions.ArrayUniqueInt(args.CardIDs)
	if len(args.CardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	//获取门店限时限次卡信息
	shm := new(models.ShopHNCardModel).Init()
	shm.Model.Begin()
	shopHNCards := shm.GetByShopIDAndHNCardIDs(shopId, args.CardIDs)
	var shopHNCardStruct []struct {
		Id       int
		ShopId   int
		Status   int
		HncardId int
	}
	var upIds, downIds, hncardIds []int
	_ = mapstructure.WeakDecode(shopHNCards, &shopHNCardStruct)
	for _, shopHNCardDesc := range shopHNCardStruct {
		if shopHNCardDesc.Status == cards.STATUS_OFF_SALE {
			downIds = append(downIds, shopHNCardDesc.Id)
			hncardIds = append(hncardIds, shopHNCardDesc.HncardId)
		} else if shopHNCardDesc.Status == cards.STATUS_ON_SALE {
			upIds = append(upIds, shopHNCardDesc.Id)
			hncardIds = append(hncardIds, shopHNCardDesc.HncardId)
		}
	}

	hncardModel := new(models.HNCardModel).Init(shm.Model.GetOrmer())
	var decOrInc string
	//限时限次卡下架
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		if err = shm.UpdateByIDs(upIds, map[string]interface{}{
			shm.Field.F_status:     cards.STATUS_OFF_SALE,
			shm.Field.F_under_time: time.Now().Unix(),
		}); err != nil {
			shm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
		}
		//同步下架门店卡项关联表数据
		shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
		if !shopItemRealtionModel.UpdateStatusByItemIds(hncardIds, cards.ITEM_TYPE_hncard, cards.STATUS_OFF_SALE, shopId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		decOrInc = "dec"
	}
	//限时限次卡上架
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		if err = shm.UpdateByIDs(downIds, map[string]interface{}{
			shm.Field.F_status:     cards.STATUS_ON_SALE,
			shm.Field.F_under_time: 0,
		}); err != nil {
			shm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
		}
		//同步上架门店卡项关联表数据
		shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
		if !shopItemRealtionModel.UpdateStatusByItemIds(hncardIds, cards.ITEM_TYPE_hncard, cards.STATUS_ON_SALE, shopId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		decOrInc = "inc"
	}
	if len(decOrInc) > 0 {
		//	更新总店中对应现时卡的在售门店数量
		if !hncardModel.UpdateSaleShopNum(hncardIds, decOrInc) {
			shm.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	shm.Model.Commit()

	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, []int{}, shopId, cards.ITEM_TYPE_hncard, upIds)
	}
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, []int{}, shopId, cards.ITEM_TYPE_hncard, downIds)
	}

	return
}

// 子店限时限次卡-rpc
func (n *HNCardLogic) ShopHncardRpc(ctx context.Context, params *cards.ArgsShopHncardRpc, list *cards.ReplyShopHncardRpc) (err error) {
	if params.ShopId <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//获取门店的限时限次卡数据
	shm := new(models.ShopHNCardModel).Init()
	shopHNCards := shm.GetByShopIDAndHNCardIDs(params.ShopId, params.HNCardIds)
	//shopHNCards = SortMapByIntField(shm.Field.F_hncard_id, shopHNCards)
	hHNCardIDs := functions.ArrayValue2Array(shm.Field.F_hncard_id, shopHNCards)

	//获取限时限次卡基本信息
	hNcm := new(models.HNCardModel).Init()
	hHNCards := hNcm.GetByHNCardIDs(hHNCardIDs)
	if len(hHNCards) == 0 {
		return
	}
	list.List = make([]cards.HNCardDesc, len(hHNCards))
	for index, hHNCard := range hHNCards {
		_ = mapstructure.WeakDecode(hHNCard, &list.List[index].HNCardBase)
		_ = mapstructure.WeakDecode(hHNCard, &list.List[index])
	}
	for index := 0; index < len(list.List); index++ {
		_ = mapstructure.WeakDecode(shopHNCards[index], &list.List[index])
		list.List[index].ShopHasAdd = 1
	}
	return
}

//验证限时限次卡规格信息
func (h *HNCardLogic) CheckSspIds(sspIds, singleIDs []int, incSingles []cards.IncSingle) (err error) {
	//1.验证规格是否重复
	//1.1去除规格ids零值
	sspIds = tools.RemoveArrayZero(sspIds)
	//1.2规格ids去重判断长度
	uniqueSspIds := functions.ArrayUniqueInt(sspIds)
	if len(sspIds) != len(uniqueSspIds) {
		err = toolLib.CreateKcErr(_const.SPEC_REPEAT_ERR)
		return
	}

	//2.判断有规格的单项目，规格选择是否正确
	//2.1获取有规格的单项目id
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

	//2.2从规格表获取单项目规格对应关系
	sspm := new(models.SingleSpecPriceModel).Init()
	singleSpecPrices := sspm.GetBySspids(sspIds, sspm.Field.F_single_id, sspm.Field.F_ssp_id)
	singleToSpec := make(map[int][]int)
	for _, ssp := range singleSpecPrices {
		sspSingleId, _ := strconv.Atoi(ssp[sspm.Field.F_single_id].(string))
		sspSpecId, _ := strconv.Atoi(ssp[sspm.Field.F_ssp_id].(string))
		singleToSpec[sspSingleId] = append(singleToSpec[sspSingleId], sspSpecId)
	}
	//2.3判断
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

// 总店删除限时限次卡
func (h *HNCardLogic) DeleteHNCard(ctx context.Context, args *cards.ArgsDelHNCard) (err error) {
	busId, err := checkBus(args.BsToken, true)
	if err != nil {
		return
	}
	hncardModel := new(models.HNCardModel).Init()
	hncardList := hncardModel.FindHNcardIdsAndBusId(args.HNCardIds, busId, hncardModel.Field.F_bus_id)
	// 检查需要删除的卡是否属于商户
	if len(hncardList) != len(args.HNCardIds) {
		err = toolLib.CreateKcErr(_const.NO_IN_BUS)
		return
	}

	// 删除总店 并且同步分店
	r := hncardModel.UpdateByHNCardIDsAndBusId(args.HNCardIds, busId, map[string]interface{}{
		hncardModel.Field.F_is_del:        cards.IS_BUS_DEL_yes,
		hncardModel.Field.F_del_time:      time.Now().Unix(),
		hncardModel.Field.F_sale_shop_num: 0,
	})
	if r != nil {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	// 同步分店
	shopHncardModel := new(models.ShopHNCardModel).Init()
	err = shopHncardModel.UpdateByHNCardIDs(args.HNCardIds, map[string]interface{}{
		shopHncardModel.Field.F_is_del:   cards.IS_BUS_DEL_yes,
		shopHncardModel.Field.F_del_time: time.Now().Unix(),
	})
	if err != nil {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.HNCardIds, cards.ITEM_TYPE_hncard) {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	return
}

// 子店删除限时限次卡
func (h *HNCardLogic) ShopDeleteHNCard(ctx context.Context, args *cards.ArgsDelHNCard) (err error) {
	shopId, _ := checkShop(args.BsToken)
	shopHncardModel := new(models.ShopHNCardModel).Init()
	shopHncardlist := shopHncardModel.FindHNcardIdsAndBusId(args.HNCardIds, shopId, shopHncardModel.Field.F_shop_id)
	// 检测选择删除的限时限次卡是否存在
	if len(args.HNCardIds) != len(shopHncardlist) {
		err = toolLib.CreateKcErr(_const.NO_IN_BUS)
		return
	}
	err = shopHncardModel.UpdateByHNCardIDsAndShopId(args.HNCardIds, shopId, map[string]interface{}{
		shopHncardModel.Field.F_is_del:   cards.IS_BUS_DEL_yes,
		shopHncardModel.Field.F_del_time: time.Now().Unix(),
	})
	if err != nil {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.HNCardIds, cards.ITEM_TYPE_hncard, shopId) {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	return
}
