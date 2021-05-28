//套餐业务处理
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/15 14:55
package logics

import (
	"context"
	"encoding/json"
	"fmt"
	"git.900sui.cn/kc/base/common/models/base"
	"strconv"
	"strings"
	"time"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/mapstructure"
	"git.900sui.cn/kc/redis"
	"git.900sui.cn/kc/rpcCards/common/models"
	"git.900sui.cn/kc/rpcCards/common/tools"
	"git.900sui.cn/kc/rpcCards/constkey"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/client/bus"
	"git.900sui.cn/kc/rpcinterface/client/file"
	bus2 "git.900sui.cn/kc/rpcinterface/interface/bus"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	file2 "git.900sui.cn/kc/rpcinterface/interface/file"
	redis2 "github.com/gomodule/redigo/redis"
)

type SmLogic struct {
}

//验证商户
//@param common.BsToken bsToken 商家token
//@param bool mustTopBus 要求是否必须是总店
func checkBus(bsToken common.BsToken, mustTopBus ...bool) (busId int, err error) {
	busId, err = bsToken.GetBusId()
	if err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if len(mustTopBus) > 0 && mustTopBus[0] == true {
		r := false
		r, err = bsToken.GetBusAcc()
		if !r || err != nil {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
	}
	return
}

//验证门店
func checkShop(bsToken common.BsToken) (shopId int, err error) {
	shopId, err = bsToken.GetShopId()
	if err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	return
}

//AddOrEditSmExt 新增/修改套餐拓展信息数据
func (s *SmLogic) AddOrEditSmExt(smModel *models.SmModel, smId int, notes []cards.CardNote, isAdd bool /*true:新增;false:修改*/) (err error) {
	if smId == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	if len(notes) > constkey.CardNotesMaxNum {
		return toolLib.CreateKcErr(_const.NOTES_NUM_MAX_10)
	}
	for _, note := range notes {
		if functions.Mb4Strlen(note.Notes) > constkey.CardNotesSimpleMaxLength {
			return toolLib.CreateKcErr(_const.NOTES_LEN_MAX_30)
		}
	}
	notesStr, _ := json.Marshal(notes)
	cardExt := new(models.SmExtModel).Init()
	mapParams := map[string]interface{}{
		cardExt.Field.F_sm_id: smId,
		cardExt.Field.F_notes: string(notesStr),
	}
	if isAdd && cardExt.Insert(mapParams) <= 0 {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	if !isAdd {
		smExMap := cardExt.Find(map[string]interface{}{cardExt.Field.F_sm_id: smId})
		if len(smExMap) > 0 && !cardExt.UpdateBySmID(smId, mapParams) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		} else if len(smExMap) == 0 && cardExt.Insert(mapParams) <= 0 {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	return
}

//GetSmExitBySmID 获取套餐拓展信息详情
func (s *SmLogic) GetSmExitBySmID(smId int, reply *[]cards.CardNote) {
	if smId == 0 {
		return
	}
	cardExt := new(models.SmExtModel).Init()
	dataMap := cardExt.GetBySmID(smId)
	if len(dataMap) > 0 {
		data := dataMap[cardExt.Field.F_notes].(string)
		json.Unmarshal([]byte(data), reply)
	}
	return
}

//添加套餐
func (s *SmLogic) AddSm(ctx context.Context, sm *cards.ArgsAddSm) (smId int, err error) {
	busId, err := checkBus(sm.BsToken, true)
	if err != nil {
		return
	}
	//验证参数
	err = s.checkSmData(busId, sm.SmBase, sm.IncludeSingles, sm.GiveSingles)
	if err != nil {
		return
	}
	//验证图片
	imgId, err := checkImg(ctx, sm.ImgHash)
	if err != nil {
		return
	}
	var hasGive uint8 = 0
	if len(sm.GiveSingles) > 0 {
		hasGive = 1
	}
	//获取包含的项目的总数量
	totalSingleNum := 0
	for _, sm := range sm.IncludeSingles {
		totalSingleNum += sm.Num
	}
	// 购买或者充值须100起
	if tools.RunMode == "prod" {
		if sm.RealPrice < cards.BUY_CRARD_MIN_AMOUNT {
			err = toolLib.CreateKcErr(_const.BUY_CRARD_MIN_AMOUNT_ERR)
			return
		}
	}
	//暂时做成永久有效
	servicePeriod := sm.ServicePeriod

	if sm.IsPermanentValidity == 1 || sm.IsPermanentValidity == 2 {
		servicePeriod = 0
	}
	//添加基本信息
	mSm := new(models.SmModel).Init()
	smId = mSm.Insert(map[string]interface{}{
		mSm.Field.F_bus_id:                busId,
		mSm.Field.F_price:                 sm.Price,
		mSm.Field.F_is_ground:             cards.SINGLE_IS_GROUND_yes,
		mSm.Field.F_real_price:            sm.RealPrice,
		mSm.Field.F_ctime:                 time.Now().Local().Unix(),
		mSm.Field.F_img_id:                imgId,
		mSm.Field.F_name:                  sm.Name,
		mSm.Field.F_sort_desc:             sm.SortDesc,
		mSm.Field.F_bind_id:               getBusMainBindId(ctx, busId),
		mSm.Field.F_has_give_signle:       hasGive,
		mSm.Field.F_service_period:        servicePeriod,
		mSm.Field.F_validcount:            totalSingleNum,
		mSm.Field.F_is_permanent_validity: cards.IS_PERMANENT_YES, // 暂时做成永久有效
	})
	if smId <= 0 {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//添加单项目
	mSmSingle := new(models.SmSingleModel).Init()
	smSingleData := []map[string]interface{}{}
	for _, single := range sm.IncludeSingles {
		smSingleData = append(smSingleData, map[string]interface{}{
			mSmSingle.Field.F_single_id: single.SingleID,
			mSmSingle.Field.F_sm_id:     smId,
			mSmSingle.Field.F_num:       single.Num,
			mSmSingle.Field.F_ssp_id:    single.SspId,
		})
	}
	if len(smSingleData) > 0 {
		mSmSingle.InsertAll(smSingleData)
	}
	//添加赠送单项目数据
	if len(sm.GiveSingles) > 0 {
		mSmGive := new(models.SmGiveModel).Init()
		giveSingleData := []map[string]interface{}{}
		for _, single := range sm.GiveSingles {
			giveSingleData = append(giveSingleData, map[string]interface{}{
				mSmGive.Field.F_single_id:          single.SingleID,
				mSmGive.Field.F_sm_id:              smId,
				mSmGive.Field.F_num:                single.Num,
				mSmGive.Field.F_period_of_validity: single.PeriodOfValidity,
			})
		}
		mSmGive.InsertAll(giveSingleData)

		//添加赠品描述
		if len(sm.GiveSingleDesc) > 0 {
			giveSingleDesc, _ := json.Marshal(sm.GiveSingleDesc)
			smdescm := new(models.SmGiveDescModel).Init()
			descData := map[string]interface{}{
				smdescm.Field.F_sm_id: smId,
				smdescm.Field.F_desc:  string(giveSingleDesc),
			}
			smdescm.Model.Data(descData).Insert()
		}
	}
	// 添加现时卡拓展信息数据
	if err = s.AddOrEditSmExt(mSm, smId, sm.Notes, true); err != nil {
		return
	}
	//添加风控统计任务
	new(ItemLogic).AddXCardTask(ctx, smId, cards.ITEM_TYPE_sm)
	return
}

//编辑套餐数据
func (s *SmLogic) EditSm(ctx context.Context, sm *cards.ArgsEditSm) (err error) {
	busId, err := checkBus(sm.BsToken, true)
	if err != nil {
		return
	}
	//验证套餐数据
	mSm := new(models.SmModel).Init()
	smInfo := mSm.GetBySmid(sm.SmId)
	if len(smInfo) == 0 {
		err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
		return
	}
	if smInfo[mSm.Field.F_bus_id] != strconv.Itoa(busId) {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	//验证参数
	err = s.checkSmData(busId, sm.SmBase, sm.IncludeSingles, sm.GiveSingles)
	if err != nil {
		return
	}
	imgId, err := checkImg(ctx, sm.ImgHash)
	if err != nil {
		return
	}
	var hasGive uint8 = 0
	if len(sm.GiveSingles) > 0 {
		hasGive = 1
	}
	//获取包含的项目的总数量
	totalSingleNum := 0
	for _, sm := range sm.IncludeSingles {
		totalSingleNum += sm.Num
	}

	// 购买或者充值须100起
	if tools.RunMode == "prod" {
		if sm.RealPrice < cards.BUY_CRARD_MIN_AMOUNT {
			err = toolLib.CreateKcErr(_const.BUY_CRARD_MIN_AMOUNT_ERR)
			return
		}
	}
	//暂时做成永久有效
	servicePeriod := sm.ServicePeriod
	if sm.IsPermanentValidity == 1 || sm.IsPermanentValidity == 2 {
		servicePeriod = 0
	}

	//修改主表信息
	r := mSm.UpdateBySmid(sm.SmId, map[string]interface{}{
		mSm.Field.F_price:           sm.Price,
		mSm.Field.F_real_price:      sm.RealPrice,
		mSm.Field.F_img_id:          imgId,
		mSm.Field.F_name:            sm.Name,
		mSm.Field.F_sort_desc:       sm.SortDesc,
		mSm.Field.F_has_give_signle: hasGive,
		mSm.Field.F_service_period:  servicePeriod,
		mSm.Field.F_validcount:      totalSingleNum,
		mSm.Field.F_is_permanent_validity:cards.IS_PERMANENT_YES,

	})

	if r == false {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	//修改包含的单项
	mSmSingle := new(models.SmSingleModel).Init()
	incSingles := mSmSingle.GetBySmid(sm.SmId)
	//定义需要新增的项目
	var addIncSingles []map[string]interface{}
	//定义需要修改的单项目
	var updateSingles = map[int]int{} //id=>num
	//定义需要删除的ids
	var delIds = []int{}
	for _, single := range sm.IncludeSingles {
		hasd := 0
		for _, dbSingle := range incSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[mSmSingle.Field.F_single_id].(string) {
				hasd = 1
				//需要修改次数
				if strconv.Itoa(single.Num) != dbSingle[mSmSingle.Field.F_num].(string) {
					dbId, _ := strconv.Atoi(dbSingle[mSmSingle.Field.F_id].(string))
					updateSingles[dbId] = single.Num
				}
			}
		}

		// 新增单项目
		if hasd == 0 {
			addIncSingles = append(addIncSingles, map[string]interface{}{
				mSmSingle.Field.F_sm_id:     sm.SmId,
				mSmSingle.Field.F_num:       single.Num,
				mSmSingle.Field.F_single_id: single.SingleID,
			})
		}
	}

	//计算需要删除的id
	for _, dbSingle := range incSingles {
		hasd := 0
		for _, single := range sm.IncludeSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[mSmSingle.Field.F_single_id].(string) {
				hasd = 1
				break
			}
		}
		if hasd == 0 {
			sid, _ := strconv.Atoi(dbSingle[mSmSingle.Field.F_id].(string))
			delIds = append(delIds, sid)
		}
	}
	if len(addIncSingles) > 0 {
		mSmSingle.InsertAll(addIncSingles)
	}
	if len(updateSingles) > 0 {
		for id, num := range updateSingles {
			mSmSingle.UpdateNumById(id, num)
		}
	}
	if len(delIds) > 0 {
		mSmSingle.DelByIds(delIds)
	}

	//修改包含的单项
	mSmGive := new(models.SmGiveModel).Init()
	gives := mSmGive.GetBySmid(sm.SmId)
	if len(sm.GiveSingles) == 0 && len(gives) == 0 {
		return
	}

	//定义需要新增的项目
	var addGiveSingles []map[string]interface{}
	//定义需要修改的单项目
	var updateGiveSingles = map[int]int{} //id=>num
	//定义需要删除的ids
	var delGiveIds = []int{}
	for _, single := range sm.GiveSingles {
		hasd := 0
		for _, dbSingle := range gives {
			if strconv.Itoa(single.SingleID) == dbSingle[mSmGive.Field.F_single_id].(string) {
				hasd = 1
				//需要修改次数
				if strconv.Itoa(single.Num) != dbSingle[mSmGive.Field.F_num].(string) {
					dbId, _ := strconv.Atoi(dbSingle[mSmGive.Field.F_id].(string))
					updateGiveSingles[dbId] = single.Num
				}
				// 更新有效期天数
				if strconv.Itoa(single.PeriodOfValidity) != dbSingle[mSmGive.Field.F_period_of_validity].(string){
					dbId, _ := strconv.Atoi(dbSingle[mSmGive.Field.F_id].(string))
					mSmGive.UpdateValidityById(dbId,single.PeriodOfValidity)
				}
			}
		}

		if hasd == 0 {
			addGiveSingles = append(addGiveSingles, map[string]interface{}{
				mSmGive.Field.F_sm_id:     sm.SmId,
				mSmGive.Field.F_num:       single.Num,
				mSmGive.Field.F_single_id: single.SingleID,
				mSmGive.Field.F_period_of_validity:single.PeriodOfValidity,
			})
		}
	}
	//计算需要删除的id
	for _, dbSingle := range gives {
		hasd := 0
		for _, single := range sm.GiveSingles {
			if strconv.Itoa(single.SingleID) == dbSingle[mSmSingle.Field.F_single_id].(string) {
				hasd = 1
				break
			}
		}
		if hasd == 0 {
			sid, _ := strconv.Atoi(dbSingle[mSmGive.Field.F_id].(string))
			delGiveIds = append(delGiveIds, sid)
		}
	}

	if len(addGiveSingles) > 0 {
		mSmGive.InsertAll(addGiveSingles)
	}
	if len(updateGiveSingles) > 0 {
		for id, num := range updateGiveSingles {
			mSmGive.UpdateNumById(id, num)
		}
	}
	if len(delGiveIds) > 0 {
		mSmGive.DelByIds(delGiveIds)
	}

	//修改赠品描述
	smdesm := new(models.SmGiveDescModel).Init()
	if len(smdesm.GetBySmId(sm.SmId))>0{
		//1.删除原有赠品描述
		smdesm.DelBySmId(sm.SmId)
	}
	//2.新增赠品描述
	if len(sm.GiveSingles) >0 && len(sm.GiveSingleDesc)>0 {
		giveSingleDesc ,_:=json.Marshal(sm.GiveSingleDesc)
		descData := map[string]interface{}{
			smdesm.Field.F_sm_id : sm.SmId,
			smdesm.Field.F_desc : string(giveSingleDesc),
		}
		smdesm.Insert(descData)
	}

	// 添加套餐拓展信息数据
	_=s.AddOrEditSmExt(mSm,sm.SmId, sm.Notes, false)

	return
}

//获取套餐的详情
//@param int smId 套餐id
//@param int shopId 门店id
//@return  cards.ReplySmInfo reply
func (s *SmLogic) SmInfo(ctx context.Context, smId int, shopId ...int) (reply cards.ReplySmInfo, err error) {
	mSm := new(models.SmModel).Init()
	sm := mSm.GetBySmid(smId)
	if len(sm) == 0 {
		err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
		return
	}

	var shopStatus int = 0
	var shopSales int = 0
	var shopid = 0
	if len(shopId) > 0 && shopId[0] > 0 {
		mShopSm := new(models.ShopSmModel).Init()
		shopSm := mShopSm.GetByShopidAdSmid(shopId[0], smId)
		if len(shopSm) == 0 {
			err = toolLib.CreateKcErr(_const.NO_IN_SHOP)
			return
		}
		shopStatus, _ = strconv.Atoi(shopSm[mShopSm.Field.F_status].(string))
		reply.SsId, _ = strconv.Atoi(shopSm[mShopSm.Field.F_id].(string))
		shopSales, _ = strconv.Atoi(shopSm[mShopSm.Field.F_sales].(string))
		shopid = shopId[0]
	}
	reply.ShareLink = tools.GetShareLink(smId, shopid, cards.ITEM_TYPE_sm)
	smBase := cards.SmBase{}
	mapstructure.WeakDecode(sm, &smBase)
	imgId, _ := strconv.Atoi(sm[mSm.Field.F_img_id].(string))
	isGround, _ := strconv.Atoi(sm[mSm.Field.F_is_ground].(string))
	imgHash, imgUrl := getImg(ctx, imgId, cards.ITEM_TYPE_sm)
	if len(shopId) > 0 && shopId[0] > 0 {
		smBase.Sales = shopSales
	}
	reply.SmBase = smBase
	reply.SmId = smId
	reply.ImgHash = imgHash
	reply.ImgUrl = imgUrl
	reply.IsGround = isGround
	reply.ShopStatus = shopStatus

	//商户信息
	if err = getBusInfo(ctx, reply.BusID, &reply.BusInfo); err != nil {
		err = toolLib.CreateKcErr(_const.SHOP_INFO_ERR)
		return
	}

	// 获取套餐门店添加详情  15 -- []int {3,4,6}
	busId, _ := strconv.Atoi(sm[mSm.Field.F_bus_id].(string))
	smShopModel := new(models.SmShopModel).Init()
	smShopInfo := smShopModel.GetBySmIdByBusId(smId, busId)

	smShopIds := make([]int, 0)
	for _, smShopIdValue := range smShopInfo {
		sshopId, _ := strconv.Atoi(smShopIdValue[smShopModel.Field.F_shop_id].(string))
		smShopIds = append(smShopIds, sshopId)
	}

	var replyShop []bus2.ReplyShopName
	rLists := make([]cards.ReplyShopName, 0)
	rpcBus := new(bus.Shop).Init()
	defer rpcBus.Close()
	err = rpcBus.GetShopNameByShopIds(ctx, &smShopIds, &replyShop)
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

	//获取包含的套餐和赠送项目信息
	mSmSingle := new(models.SmSingleModel).Init()
	smSingles := mSmSingle.GetBySmid(smId)
	singleIds := functions.ArrayValue2Array(mSmSingle.Field.F_single_id, smSingles)

	giveSingles := []map[string]interface{}{}
	mSmGive := new(models.SmGiveModel).Init()
	if sm[mSm.Field.F_has_give_signle].(string) == strconv.Itoa(cards.HAS_GIVE_SINGLE_yes) {
		giveSingles = mSmGive.GetBySmid(smId)
		giveSingleIds := functions.ArrayValue2Array(mSmGive.Field.F_single_id, giveSingles)
		singleIds = append(singleIds, giveSingleIds...)
	}

	if len(giveSingles) > 0 {
		allSingles, _ := getIncSingles2(ctx, shopid, "", "", singleIds, smSingles)
		for _, single := range giveSingles {
			for i := range allSingles {
				if single[mSmGive.Field.F_single_id].(string) == strconv.Itoa(allSingles[i].SingleID) {
					allSingles[i].Num, _ = strconv.Atoi(single[mSmGive.Field.F_num].(string))
					allSingles[i].PeriodOfValidity, _ = strconv.Atoi(single[mSmGive.Field.F_period_of_validity].(string))
					reply.GiveSingles = append(reply.GiveSingles, allSingles[i])
					break
				}
			}
		}

		//获取赠品描述信息
		smdescm := new(models.SmGiveDescModel).Init()
		desc, ok := smdescm.GetBySmId(smId)[smdescm.Field.F_desc].(string)
		if ok {
			json.Unmarshal([]byte(desc), &reply.GiveSingleDesc)
		}
	}
	s.GetSmExitBySmID(smId, &reply.Notes)
	//浏览次数加1
	redis.RedisGlobMgr.Hincrby(constkey.SM_CLICKS, strconv.Itoa(smId), 1)
	return
}

//获取商家的套餐列表
func (s *SmLogic) GetBusPage(ctx context.Context, busId, shopId, start, limit int, isGround string, filterShopHasAdd bool) (list cards.ReplySmPage, err error) {
	list = cards.ReplySmPage{
		TotalNum:  0,
		List:      []cards.ListSm{},
		IndexImgs: map[int]string{},
	}
	if busId <= 0 || start < 0 || limit < 0 {
		return
	}

	mSm := new(models.SmModel).Init()
	sms := make([]map[string]interface{}, 0)
	list.IndexImgs = make(map[int]string)

	//子店已添加的卡项
	var shopAddCards []map[string]interface{}
	scModel := new(models.ShopSmModel).Init()
	where := make([]base.WhereItem, 0)
	where = append(where, base.WhereItem{mSm.Field.F_bus_id, busId})
	where = append(where, base.WhereItem{mSm.Field.F_is_del, cards.IS_BUS_DEL_no})
	if shopId > 0 {
		shopAddCards = scModel.SelectRcardsByWherePage([]base.WhereItem{{scModel.Field.F_shop_id, shopId},{scModel.Field.F_is_del, cards.IS_BUS_DEL_no}}, 0, 0)
		if filterShopHasAdd && len(shopAddCards) > 0 {
			shopHasAddSmIds := functions.ArrayValue2Array(scModel.Field.F_sm_id, shopAddCards)
			where = append(where, base.WhereItem{mSm.Field.F_sm_id, []interface{}{"NOT IN", shopHasAddSmIds}})
		}
	}

	//获取总数量
	//if isGround == "" {
	sms = mSm.SelectSmsByWherePage(where, start, limit)
	list.TotalNum = mSm.GetNumByWhere(where)
	//} else {
	//	isground, _ := strconv.Atoi(isGround)
	//	isground = isground - 1
	//	sms = mSm.SelectSmsByWherePage(where,start,limit,isGround)
	//	list.TotalNum = mSm.GetNumByWhere(where)
	//}

	if len(sms) == 0 {
		return
	}
	smsArr := []cards.SmBase{}
	mapstructure.WeakDecode(sms, &smsArr)
	for k, sm := range sms {
		smId, _ := strconv.Atoi(sm[mSm.Field.F_sm_id].(string))
		sales, _ := strconv.Atoi(sm[mSm.Field.F_sales].(string))
		//isGround, _ := strconv.Atoi(sm[mSm.Field.F_is_ground].(string))
		clicks, _ := redis2.Int(redis.RedisGlobMgr.Hget(constkey.SM_CLICKS, sm[mSm.Field.F_sm_id].(string)))
		imgId, _ := strconv.Atoi(sm[mSm.Field.F_img_id].(string))
		shopStatus, shopHasAdd, shopItemId, shopDelStatus := 0, 0, 0, 0
		for _, shopCard := range shopAddCards {
			if sm[mSm.Field.F_sm_id].(string) == shopCard[scModel.Field.F_sm_id].(string) { //表明当前子店已添加该卡项
				shopItemId, _ = strconv.Atoi(shopCard[scModel.Field.F_id].(string))
				shopStatus, _ = strconv.Atoi(shopCard[scModel.Field.F_status].(string))
				shopHasAdd = 1
				shopDelStatus, _ = strconv.Atoi(shopCard[scModel.Field.F_is_del].(string))
				break
			}
		}
		list.List = append(list.List, cards.ListSm{
			SmId:     smId,
			SmBase:   smsArr[k],
			CtimeStr: functions.TimeToStr(int64(smsArr[k].Ctime)),
			Clicks:   clicks,
			Sales:    sales,
			//IsGround:      isGround,
			ShopStatus:    shopStatus,
			ShopHasAdd:    shopHasAdd,
			ShopDelStatus: shopDelStatus,
			ShopItemId:    shopItemId,
			ImgId:         imgId,
		})
	}
	//获取图片信息
	imgIds := functions.ArrayValue2Array(mSm.Field.F_img_id, sms)
	list.IndexImgs = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_sm)
	return
}

//设置套餐的适用门店(一期优化，废用）
func (s *SmLogic) SetSmShop(ctx context.Context, args *cards.ArgsSetSmShop) (err error) {
	busId, err := checkBus(args.BsToken, true)
	if err != nil {
		return
	}
	if len(args.SmIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	if args.IsAllShop == false && len(args.ShopIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//套餐id重复提交判断
	realSmIds := functions.ArrayUniqueInt(args.SmIds)
	if len(realSmIds) != len(args.SmIds) {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//店铺id重复提交判断
	realShopIds := functions.ArrayUniqueInt(args.ShopIds)
	if len(realShopIds) != len(args.ShopIds) && args.IsAllShop == false {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	//检查套餐是否属于企业
	busSms, err := s.checkSmids(busId, realSmIds)
	if err != nil {
		return
	}

	mSmShop := new(models.SmShopModel).Init()
	var insertData = []map[string]interface{}{}
	//走全部适用逻辑
	if args.IsAllShop == true {
		for _, smId := range realSmIds {
			insertData = append(insertData, map[string]interface{}{
				mSmShop.Field.F_sm_id:   smId,
				mSmShop.Field.F_bus_id:  busId,
				mSmShop.Field.F_shop_id: 0,
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
		for _, smId := range realSmIds {
			for _, shopid := range realShopIds {
				insertData = append(insertData, map[string]interface{}{
					mSmShop.Field.F_sm_id:   smId,
					mSmShop.Field.F_bus_id:  busId,
					mSmShop.Field.F_shop_id: shopid,
				})
			}
		}
	}

	//处理规格
	if len(insertData) > 0 {
		mSmShop.DelBySmIds(realSmIds)
		mSmShop.InsertAll(insertData)
	}

	//提取目前是上架状态的套餐id
	var onGroundSmids []int
	for _, sm := range busSms {
		if sm["is_ground"].(string) == strconv.Itoa(cards.IS_GROUND_yes) {
			tmpSmid, _ := strconv.Atoi(sm["sm_id"].(string))
			onGroundSmids = append(onGroundSmids, tmpSmid)
		}
	}

	mShopSm := new(models.ShopSmModel).Init()
	shopSms := mShopSm.GetBySmids(realSmIds)
	var downIds = []int{}
	//需要把已添加到门店的但是状态为总店禁用的套餐状态改成下架状态
	for _, shopsm := range shopSms {
		shopStatus, _ := strconv.Atoi(shopsm[mShopSm.Field.F_status].(string))
		shopSmId, _ := strconv.Atoi(shopsm[mShopSm.Field.F_id].(string))
		shopid, _ := strconv.Atoi(shopsm[mShopSm.Field.F_shop_id].(string))
		smId, _ := strconv.Atoi(shopsm[mShopSm.Field.F_sm_id].(string))
		if shopStatus == cards.STATUS_DISABLE && functions.InArray(smId, onGroundSmids) {
			if args.IsAllShop == true || functions.InArray(shopid, realShopIds) {
				downIds = append(downIds, shopSmId)
			}
		}
	}

	if len(downIds) > 0 {
		mShopSm.UpdateByIds(downIds, map[string]interface{}{
			mShopSm.Field.F_status: cards.STATUS_OFF_SALE,
		})
	}

	//如果是设置成部分门店适用，需要把在门店已上架的但是现在改为不适用的门店套餐改为总店警用
	if args.IsAllShop == false {
		var disableIds = []int{}
		for _, shopsm := range shopSms {
			shopStatus, _ := strconv.Atoi(shopsm[mShopSm.Field.F_status].(string))
			if shopStatus == cards.STATUS_DISABLE {
				continue
			}
			shopid, _ := strconv.Atoi(shopsm[mShopSm.Field.F_shop_id].(string))
			if functions.InArray(shopid, realShopIds) == false {
				disableId, _ := strconv.Atoi(shopsm[mShopSm.Field.F_id].(string))
				disableIds = append(disableIds, disableId)
			}
		}
		if len(disableIds) > 0 {
			mShopSm.UpdateByIds(disableIds, map[string]interface{}{
				mShopSm.Field.F_status:     cards.STATUS_DISABLE,
				mShopSm.Field.F_under_time: time.Now().Unix(),
			})
			//添加维护es的shop-item文档的任务
			setShopItem(ctx, []int{}, 0, cards.ITEM_TYPE_sm, disableIds)
		}
	}

	return nil
}

//总店上下架套餐
func (s *SmLogic) DownUpSm(ctx context.Context, args *cards.ArgsDownUpSm) (err error) {
	busId, err := checkBus(args.BsToken, true)
	if err != nil {
		return
	}

	//检查套餐是否属于企业
	sms, err := s.checkSmids(busId, args.SmIds)
	if err != nil {
		return
	}
	var smsStruct []struct {
		SmId     int
		IsGround int
	}
	mapstructure.WeakDecode(sms, &smsStruct)
	var downIds, upIds []int
	for _, sm := range smsStruct {
		if sm.IsGround == cards.IS_GROUND_no {
			downIds = append(downIds, sm.SmId)
		} else {
			upIds = append(upIds, sm.SmId)
		}
	}
	mSm := new(models.SmModel).Init()
	mShopSm := new(models.ShopSmModel).Init()
	//下架操作
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		r := mSm.UpdateBySmids(upIds, map[string]interface{}{
			mSm.Field.F_is_ground:     cards.IS_GROUND_no,
			mSm.Field.F_under_time:    time.Now().Unix(),
			mSm.Field.F_sale_shop_num: 0,
		})
		if r == false {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//将分店的套餐设置为总店禁用
		mShopSm.UpdateBySmids(upIds, map[string]interface{}{
			mShopSm.Field.F_status:     cards.STATUS_DISABLE,
			mShopSm.Field.F_under_time: time.Now().Unix(),
		})
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, upIds, 0, cards.ITEM_TYPE_sm)
	}

	//上架操作
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		r := mSm.UpdateBySmids(downIds, map[string]interface{}{
			mSm.Field.F_is_ground:  cards.IS_GROUND_yes,
			mSm.Field.F_under_time: 0,
		})
		if r == false {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//将已添加到子店的套餐解除总店禁用状态 需要将不适用的门店套餐id 过滤掉
		shopSms := mShopSm.GetBySmids(downIds)
		if len(shopSms) == 0 {
			return
		}
		mSmShop := new(models.SmShopModel).Init()
		smShops := mSmShop.GetBySmids(downIds)

		var smShopArr = []string{} //smid_shopid
		for _, smshop := range smShops {
			smShopArr = append(smShopArr, fmt.Sprintf("%s_%s", smshop[mSmShop.Field.F_sm_id].(string), smshop[mSmShop.Field.F_shop_id].(string)))
		}
		var unDisableIds = []int{}
		for _, shopsm := range shopSms {
			smidShopidStr := fmt.Sprintf("%s_%s", shopsm[mShopSm.Field.F_sm_id].(string), shopsm[mShopSm.Field.F_shop_id].(string))
			smidAllStr := fmt.Sprintf("%s_0", shopsm[mShopSm.Field.F_sm_id].(string))

			if functions.InArray(smidAllStr, smShopArr) || functions.InArray(smidShopidStr, smShopArr) {
				shopsmId, _ := strconv.Atoi(shopsm[mShopSm.Field.F_id].(string))
				unDisableIds = append(unDisableIds, shopsmId)
			}
		}
		if len(unDisableIds) > 0 {
			mShopSm.UpdateByIds(unDisableIds, map[string]interface{}{
				mShopSm.Field.F_status: cards.STATUS_OFF_SALE,
			})
		}
	}

	return
}

//子店获取适用本店的套餐列表
func (s *SmLogic) ShopGetBusSmPage(ctx context.Context, busId, shopId, start, limit int) (reply cards.ReplyShopGetBusSmPage, err error) {
	if busId <= 0 || shopId <= 0 || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	reply = cards.ReplyShopGetBusSmPage{
		ReplySmPage: cards.ReplySmPage{
			TotalNum:  0,
			List:      []cards.ListSm{},
			IndexImgs: map[int]string{},
		},
	}
	mSmShop := new(models.SmShopModel).Init()
	smShops := mSmShop.GetPageByShopId(busId, shopId, start, limit)
	if len(smShops) == 0 {
		return
	}
	//获取套餐基本信息
	smIds := functions.ArrayValue2Array(mSmShop.Field.F_sm_id, smShops)
	mSm := new(models.SmModel).Init()
	sms := mSm.GetBySmids(smIds)
	if len(sms) == 0 {
		return
	}
	var smsArr []cards.SmBase
	_ = mapstructure.WeakDecode(sms, &smsArr)
	smsMap := map[string]cards.ListSm{}
	//获取门店添加状态
	mShopSmModel := new(models.ShopSmModel).Init()
	shopSms := mShopSmModel.GetByShopidAdSmids(shopId, smIds)
	shopSmids := functions.ArrayValue2Array(mShopSmModel.Field.F_sm_id, shopSms)

	for k, sm := range sms {
		smId, _ := strconv.Atoi(sm[mSm.Field.F_sm_id].(string))
		sales, _ := strconv.Atoi(sm[mSm.Field.F_sales].(string))
		isGround, _ := strconv.Atoi(sm[mSm.Field.F_is_ground].(string))
		clicks, _ := redis2.Int(redis.RedisGlobMgr.Hget(constkey.SM_CLICKS, sm[mSm.Field.F_sm_id].(string)))
		imgId, _ := strconv.Atoi(sm[mSm.Field.F_img_id].(string))
		shopHasAdd := 0
		if functions.InArray(smId, shopSmids) {
			shopHasAdd = 1
		}

		smsMap[sm[mSm.Field.F_sm_id].(string)] = cards.ListSm{
			SmId:       smId,
			SmBase:     smsArr[k],
			CtimeStr:   functions.TimeToStr(int64(smsArr[k].Ctime)),
			Clicks:     clicks,
			Sales:      sales,
			IsGround:   isGround,
			ShopStatus: 0,
			ShopHasAdd: shopHasAdd,
			ImgId:      imgId,
		}
	}

	for _, smshop := range smShops {
		reply.List = append(reply.List, smsMap[smshop[mSmShop.Field.F_sm_id].(string)])
	}
	//获取图片信息
	imgIds := functions.ArrayValue2Array(mSm.Field.F_img_id, sms)
	reply.IndexImgs = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_sm)
	//获取数量
	reply.TotalNum = mSmShop.GetNumByShopId(busId, shopId)

	return

}

//子店添加套餐到自己的门店
func (s *SmLogic) ShopAddSm(args *cards.ArgsShopAddSm) (err error) {
	args.SmIds = functions.ArrayUniqueInt(args.SmIds)
	if len(args.SmIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	var busId, shopId int
	if busId, err = checkBus(args.BsToken); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopId, err = checkShop(args.BsToken); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopId <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}

	//检查套餐id是否适用当前门店
	mSmShop := new(models.SmShopModel).Init()
	smShops := mSmShop.GetByShopIdAndSmids(busId, shopId, args.SmIds)
	if len(smShops) != len(args.SmIds) {
		err = toolLib.CreateKcErr(_const.NO_IN_SHOP)
		return
	}
	//过滤掉已经添加了的套餐id
	mShopSm := new(models.ShopSmModel).Init()
	shopSms := mShopSm.GetByShopidAdSmids(shopId, args.SmIds)
	shopSmIds := functions.ArrayValue2Array(mShopSm.Field.F_sm_id, shopSms)


	// 实际需要添加卡项
	var addSmids = []int{}
	for _, smid := range args.SmIds {
		if functions.InArray(smid, shopSmIds) == false {
			addSmids = append(addSmids, smid)
		}
	}
	if len(addSmids) == 0 {
		return
	}

	//校验当前门店是否已经将卡项内涉及到的单项目添加到自己的门店内
	allSingle, singleIds, err := new(ItemLogic).getItemCardIncSingleIds(addSmids, cards.ITEM_TYPE_sm)
	if err != nil {
		return
	}
	if err = new(ItemLogic).validShopSingleContainItemCardSingles(shopId, busId, allSingle, singleIds); err != nil {
		return
	}

	//获取套餐在总店的上下架状态
	mSm := new(models.SmModel).Init()
	sms := mSm.GetBySmids(addSmids, []string{
		mSm.Field.F_sm_id,
		mSm.Field.F_is_ground,
	})
	smsMap := functions.ArrayRebuild(mSm.Field.F_sm_id, sms)
	var addData []map[string]interface{}
	shopItemRelationData := make([]map[string]interface{}, 0)
	shopItemRelationModel := new(models.ShopItemRelationModel).Init()
	for _, addSmid := range addSmids {
		status := cards.STATUS_OFF_SALE
		if _, ok := smsMap[strconv.Itoa(addSmid)]; ok {
			if newsmmap, ok2 := smsMap[strconv.Itoa(addSmid)].(map[string]interface{}); ok2 {
				isground, _ := strconv.Atoi(newsmmap[mSm.Field.F_is_ground].(string))
				if isground == cards.IS_GROUND_no {
					status = cards.STATUS_DISABLE
				}
			}
		}

		ctime := time.Now().Local().Unix()
		addData = append(addData, map[string]interface{}{
			mShopSm.Field.F_sm_id:   addSmid,
			mShopSm.Field.F_status:  status,
			mShopSm.Field.F_shop_id: shopId,
			mShopSm.Field.F_ctime:   ctime,
		})

		shopItemRelationData = append(shopItemRelationData, map[string]interface{}{
			shopItemRelationModel.Field.F_item_id:   addSmid,
			shopItemRelationModel.Field.F_item_type: cards.ITEM_TYPE_sm,
			shopItemRelationModel.Field.F_status:    cards.STATUS_OFF_SALE,
			shopItemRelationModel.Field.F_shop_id:   shopId,
			shopItemRelationModel.Field.F_is_del:    cards.ITEM_IS_DEL_NO,
		})
	}

	if len(addData) > 0 {
		mShopSm.InsertAll(addData)
	}
	//门店卡项关联表数据插入
	if len(shopItemRelationData) > 0 {
		if shopItemRelationModel.InsertAll(shopItemRelationData) <= 0 {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	return
}

// 子店添加套餐（一期优化，去掉总店推送，改为门店自动拉取）
func (s *SmLogic) ShopAddToSm(ctx context.Context, args *cards.ArgsShopAddSm) (err error) {
	shopId, _ := args.GetShopId()
	busId, _ := args.GetBusId()
	// 先进行检查，[1,2,3] ,其中2已经添加过，需要过滤掉，-- [1,3]
	// 过滤掉已经添加了的套餐id
	mShopSm := new(models.ShopSmModel).Init()
	// shopSmLists := mShopSm.GetByShopidAdSmids(shopId, args.SmIds)
	shopSmLists := mShopSm.GetByShopIdBySmids(shopId, args.SmIds)
	shopSmIds := functions.ArrayValue2Array(mShopSm.Field.F_sm_id, shopSmLists)

	// 刷选出已经添加过并且删除的数据
	delHcardIdSlice := make([]int, 0)
	for _, hcardMap := range shopSmLists {
		isDel, _ := strconv.Atoi(hcardMap[mShopSm.Field.F_is_del].(string))
		if isDel == cards.IS_BUS_DEL_yes {
			delHcardId, _ := strconv.Atoi(hcardMap[mShopSm.Field.F_sm_id].(string))
			delHcardIdSlice = append(delHcardIdSlice, delHcardId)
		}
	}

	// 更新门店之前添加过并删除的数据
	if len(delHcardIdSlice) > 0 {
		// 更新数据删除和上下架状态
		if updateBool := mShopSm.UpdateShopIdBySmids(delHcardIdSlice, shopId, map[string]interface{}{
			mShopSm.Field.F_is_del: cards.IS_BUS_DEL_no,
			mShopSm.Field.F_status: cards.STATUS_OFF_SALE,
		}); !updateBool {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
		//更新卡项关联表
		sirModel := new(models.ShopItemRelationModel).Init()
		if b := sirModel.UpdateByItemIdsAndShopId(delHcardIdSlice, cards.ITEM_TYPE_sm, shopId, map[string]interface{}{
			sirModel.Field.F_is_del: cards.IS_BUS_DEL_no,
			sirModel.Field.F_status: cards.STATUS_OFF_SALE,
		}); !b {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}


	addSmids := make([]int, 0)
	for _, smid := range args.SmIds {
		if functions.InArray(smid, shopSmIds) == false {
			addSmids = append(addSmids, smid)
		}
	}

	//校验当前门店是否已经将卡项内涉及到的单项目添加到自己的门店内
	allSingle, singleIds, err := new(ItemLogic).getItemCardIncSingleIds(addSmids, cards.ITEM_TYPE_sm)
	if err != nil {
		return
	}
	if err = new(ItemLogic).validShopSingleContainItemCardSingles(shopId, busId, allSingle, singleIds); err != nil {
		return
	}

	// 需要添加数据
	var addData []map[string]interface{}
	for _, addSmid := range addSmids {
		status := cards.STATUS_OFF_SALE
		ctime := time.Now().Local().Unix()
		addData = append(addData, map[string]interface{}{
			mShopSm.Field.F_sm_id:   addSmid,
			mShopSm.Field.F_status:  status,
			mShopSm.Field.F_shop_id: shopId,
			mShopSm.Field.F_ctime:   ctime,
		})
	}

	// 过滤的数据添加到门店套餐表
	if len(addData) > 0 {
		id := mShopSm.InsertAll(addData)
		if id < 0 {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}

	// 初始化适用门店模型
	smShopModel := new(models.SmShopModel).Init()
	smshopSms := smShopModel.GetBySmids(args.SmIds)
	smshopIds := functions.ArrayValue2Array(smShopModel.Field.F_sm_id, smshopSms)
	addSmShopIds := make([]int, 0)
	for _, smid := range args.SmIds {
		if functions.InArray(smid, smshopIds) == false {
			addSmShopIds = append(addSmShopIds, smid)
		}
	}
	if len(addSmShopIds) == 0 {
		return
	}

	var addSmShopData []map[string]interface{} // 添加适用门店表的数据
	for _, addSmShopId := range addSmShopIds {
		addSmShopData = append(addSmShopData, map[string]interface{}{
			smShopModel.Field.F_sm_id:   addSmShopId,
			smShopModel.Field.F_shop_id: shopId,
			smShopModel.Field.F_bus_id:  busId,
		})
	}
	// 过滤的数据添加到适用套餐表
	if len(addSmShopData) > 0 {
		id := smShopModel.InsertAll(addSmShopData)
		if id < 0 {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}

	mSm := new(models.SmModel).Init()
	sms := mSm.GetBySmids(addSmids, []string{
		mSm.Field.F_sm_id,
		mSm.Field.F_is_ground,
	})
	smsMap := functions.ArrayRebuild(mSm.Field.F_sm_id, sms)
	shopItemRelationData := make([]map[string]interface{}, 0)
	shopItemRelationModel := new(models.ShopItemRelationModel).Init()
	for _, addSmid := range addSmids {
		status := cards.STATUS_OFF_SALE
		if _, ok := smsMap[strconv.Itoa(addSmid)]; ok {
			if newsmmap, ok2 := smsMap[strconv.Itoa(addSmid)].(map[string]interface{}); ok2 {
				isground, _ := strconv.Atoi(newsmmap[mSm.Field.F_is_ground].(string))
				if isground == cards.IS_GROUND_no {
					status = cards.STATUS_DISABLE
				}
			}
		}

		ctime := time.Now().Local().Unix()
		addData = append(addData, map[string]interface{}{
			mShopSm.Field.F_sm_id:   addSmid,
			mShopSm.Field.F_status:  status,
			mShopSm.Field.F_shop_id: shopId,
			mShopSm.Field.F_ctime:   ctime,
		})

		shopItemRelationData = append(shopItemRelationData, map[string]interface{}{
			shopItemRelationModel.Field.F_item_id:   addSmid,
			shopItemRelationModel.Field.F_item_type: cards.ITEM_TYPE_sm,
			shopItemRelationModel.Field.F_status:    cards.STATUS_OFF_SALE,
			shopItemRelationModel.Field.F_shop_id:   shopId,
			shopItemRelationModel.Field.F_is_del:    cards.ITEM_IS_DEL_NO,
		})
	}
	//门店卡项关联表数据插入
	if len(shopItemRelationData) > 0 {
		if shopItemRelationModel.InsertAll(shopItemRelationData) <= 0 {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	return
}

//获取子店的套餐列表
func (s *SmLogic) ShopSmPage(ctx context.Context, shopId, start, limit, status int) (reply cards.ReplyShopSmPage, err error) {
	if shopId <= 0 || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	reply = cards.ReplyShopSmPage{
		TotalNum:  0,
		Lists:     []cards.ListShopSm{},
		IndexImgs: map[int]string{},
	}
	reply.Lists = make([]cards.ListShopSm, 0)
	//获取门店的套餐数据
	mShopSm := new(models.ShopSmModel).Init()
	shopSms := mShopSm.GetPageByShopId(shopId, start, limit, status)
	smIds := functions.ArrayValue2Array(mShopSm.Field.F_sm_id, shopSms)
	//获取套餐基本信息
	mSm := new(models.SmModel).Init()
	sms := mSm.GetBySmids(smIds)
	if len(sms) == 0 {
		return
	}
	var smsArr []cards.SmBase
	_ = mapstructure.WeakDecode(sms, &smsArr)
	smsMap := map[string]cards.ListShopSm{}
	//获取不同卡项-适用单项目的个数和赠送单项目的个数
	gaagsNumMap := GetApplyAndGiveSingleNum(smIds, cards.ITEM_TYPE_sm)

	for k, sm := range sms {
		smId, _ := strconv.Atoi(sm[mSm.Field.F_sm_id].(string))
		isGround, _ := strconv.Atoi(sm[mSm.Field.F_is_ground].(string))
		clicks, _ := redis2.Int(redis.RedisGlobMgr.Hget(constkey.SM_CLICKS, sm[mSm.Field.F_sm_id].(string)))
		imgId, _ := strconv.Atoi(sm[mSm.Field.F_img_id].(string))

		smsMap[sm[mSm.Field.F_sm_id].(string)] = cards.ListShopSm{
			ShopSmId: 0,
			ListSm: cards.ListSm{
				SmId:           smId,
				SmBase:         smsArr[k],
				CtimeStr:       functions.TimeToStr(int64(smsArr[k].Ctime)),
				Clicks:         clicks,
				Sales:          0,
				IsGround:       isGround,
				ShopStatus:     0,
				ShopHasAdd:     1,
				ImgId:          imgId,
				ApplySingleNum: gaagsNumMap[smId].ApplySingleNum,
				GiveSingleNum:  gaagsNumMap[smId].GiveSingleNum,
			},
		}
	}

	for _, shopsm := range shopSms {
		listShopSm := smsMap[shopsm[mShopSm.Field.F_sm_id].(string)]
		listShopSm.ShopSmId, _ = strconv.Atoi(shopsm[mShopSm.Field.F_id].(string))
		listShopSm.ShopItemId = listShopSm.ShopSmId
		listShopSm.ShopStatus, _ = strconv.Atoi(shopsm[mShopSm.Field.F_status].(string))
		listShopSm.Sales, _ = strconv.Atoi(shopsm[mShopSm.Field.F_sales].(string))
		reply.Lists = append(reply.Lists, listShopSm)
	}
	//获取图片信息
	imgIds := functions.ArrayValue2Array(mSm.Field.F_img_id, sms)
	reply.IndexImgs = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_sm)
	//获取数量
	reply.TotalNum = mShopSm.GetNumByShopId(shopId, status)

	return
}

//门店上下架套餐
func (s *SmLogic) ShopDownUpSm(ctx context.Context, args *cards.ArgsShopDownUpSm) (err error) {
	var shopId int
	if shopId, err = checkShop(args.BsToken); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopId <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	args.ShopSmIds = functions.ArrayUniqueInt(args.ShopSmIds)
	if len(args.ShopSmIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	//获取门店套餐信息
	mShopSm := new(models.ShopSmModel).Init()
	shopSms := mShopSm.GetByIds(args.ShopSmIds)
	var shopSmsStruct []struct {
		Id     int
		ShopId int
		Status int
		SmId   int
	}
	var upIds, downIds, upSmids, downSmids []int
	_ = mapstructure.WeakDecode(shopSms, &shopSmsStruct)
	for _, shopSm := range shopSmsStruct {
		if shopId != shopSm.ShopId {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
		if shopSm.Status == cards.STATUS_OFF_SALE {
			downIds = append(downIds, shopSm.Id)
			downSmids = append(downSmids, shopSm.SmId)
		} else if shopSm.Status == cards.STATUS_ON_SALE {
			upIds = append(upIds, shopSm.Id)
			upSmids = append(upSmids, shopSm.SmId)
		}
	}
	mSm := new(models.SmModel).Init()
	//套餐下架
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		mShopSm.UpdateByIds(upIds, map[string]interface{}{
			mShopSm.Field.F_status: cards.STATUS_OFF_SALE,
			mSm.Field.F_under_time: time.Now().Unix(),
		})
		for _, smId := range upSmids {
			mSm.DecrSaleShopNumBySmid(smId, 1)
		}
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, upSmids, shopId, cards.ITEM_TYPE_sm)
		//同步下架门店卡项关联表数据
		shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
		if !shopItemRealtionModel.UpdateStatusByItemIds(upSmids, cards.ITEM_TYPE_sm, cards.STATUS_OFF_SALE, shopId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	//套餐上架
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		mShopSm.UpdateByIds(downIds, map[string]interface{}{
			mShopSm.Field.F_status: cards.STATUS_ON_SALE,
		})
		for _, smId := range downSmids {
			mSm.IncrSaleShopNumBySmid(smId, 1)
		}
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, downSmids, shopId, cards.ITEM_TYPE_sm)
		//同步上架门店卡项关联表数据
		shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
		if !shopItemRealtionModel.UpdateStatusByItemIds(downSmids, cards.ITEM_TYPE_sm, cards.STATUS_ON_SALE, shopId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	return
}

//检查套餐是否属于商家
func (s *SmLogic) checkSmids(busId int, smIds []int) (sms []map[string]interface{}, err error) {
	//检查套餐是否属于企业
	mSm := new(models.SmModel).Init()
	sms = mSm.GetBySmids(smIds)
	busIdStr := strconv.Itoa(busId)
	for _, sm := range sms {
		if sm[mSm.Field.F_bus_id].(string) != busIdStr {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
	}
	return
}

//检查套餐的入参数据
func (s *SmLogic) checkSmData(busId int, smBase cards.SmBase, incSingles []cards.IncSingle, giveSingles []cards.IncSingle) (err error) {
	if err = cards.VerfiyName(smBase.Name); err != nil {
		return
	}
	if err = cards.VerfiyPrice(smBase.RealPrice, smBase.Price); err != nil {
		return
	}
	//if err = cards.VerfiyServicePeriod(smBase.ServicePeriod); err != nil {
	//	return
	//}
	if err = cards.VerifySinglesNum(len(incSingles)); err != nil {
		return
	}
	if err = cards.VerifyGiveSinglesNum(len(giveSingles)); err != nil {
		return
	}
	if smBase.IsPermanentValidity == cards.IS_PERMANENT_NO && smBase.ServicePeriod == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	//单项目id集合
	var singleIds []int
	var m = make(map[string]bool)
	for _, single := range incSingles {
		if single.Num <= 0 {
			err = toolLib.CreateKcErr(_const.PARAM_ERR)
			return
		}
		if _, ok := m[fmt.Sprintf("%d,%d", single.SingleID, single.SspId)]; ok {
			err = toolLib.CreateKcErr(_const.SINGLE_REPEAT_ERR)
			return
		} else {
			m[fmt.Sprintf("%d,%d", single.SingleID, single.SspId)] = false
		}
		singleIds = append(singleIds, single.SingleID)
	}
	//验证包含的单项目有没有传入规格
	//var sspIds []int
	//for _, single := range incSingles {
	//	sspIds = append(sspIds, single.SspId)
	//}
	//if err = new(HNCardLogic).CheckSspIds(sspIds, singleIds, incSingles); err != nil {
	//	return
	//}
	/*model := new(models.SingleModel).Init()
	maps := model.GetBySingleids(singleIds, []string{model.Field.F_single_id, model.Field.F_has_spec})
	for _, m := range maps {
		for _, single := range incSingles {
			if m[model.Field.F_single_id].(string) == strconv.Itoa(single.SingleID) {
				if m[model.Field.F_has_spec].(string) == "1" && single.SspId == 0 {
					return toolLib.CreateKcErr(_const.SINGLE_EXIST_SPEC)
				}
				break
			}
		}
	}*/

	var giveIds []int
	if len(giveSingles) > 0 {
		for _, gsingle := range giveSingles {
			if gsingle.Num <= 0 {
				err = toolLib.CreateKcErr(_const.PARAM_ERR)
				return
			}
			if gsingle.PeriodOfValidity == 0 {
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

	if err = checkSingles(busId, singleIds); err != nil {
		return
	}

	return nil
}

//验证单项目是否属于商家
func checkSingles(busId int, singleIds []int) error {
	mSingle := new(models.SingleModel).Init()
	singles := mSingle.GetBySingleids(singleIds, []string{
		mSingle.Field.F_bus_id,
		mSingle.Field.F_single_id,
	})

	//if len(singleIds) != len(singles) {
	//	return toolLib.CreateKcErr(_const.SHOP_SINGLE_NOEXIST)
	//}

	busIdStr := strconv.Itoa(busId)
	for _, single := range singles {
		if busIdStr != single[mSingle.Field.F_bus_id].(string) {
			return toolLib.CreateKcErr(_const.OPT_OTHER_BUS_ITEM)
		}
	}

	return nil
}

//检查封面图片
func checkImg(ctx context.Context, imgHash string) (imgId int, err error) {
	imgId = 0
	if len(strings.Trim(imgHash, "")) == 0 {
		return
	}
	rpcImg := new(file.Upload).Init()
	defer rpcImg.Close()
	var replyImg = file2.ReplyFileInfo{}
	err = rpcImg.GetImageByHash(ctx, imgHash, &replyImg)
	if err != nil {
		err = toolLib.CreateKcErr(_const.PICTURES_EMPTY)
		return
	}
	imgId = replyImg.Id
	return
}

//获取封面图片
func getImg(ctx context.Context, imgId int, cardType int) (imgHash, imgUrl string) {
	imgHash = ""
	imgUrl = ""
	if imgId <= 0 {
		if cardType > 0 {
			imgUrl = constkey.CardsDefaultPics[cardType]
		}
		return
	}
	rpcImg := new(file.Upload).Init()
	defer rpcImg.Close()
	var replyImg = file2.ReplyFileInfo{}
	err := rpcImg.GetImageById(ctx, imgId, &replyImg)
	if err != nil {
		return
	}
	imgHash = replyImg.Hash
	imgUrl = replyImg.Path
	return
}

//根据图片ids 获取图片数据
func getImgsByImgIds(ctx context.Context, imgIds []int, cardType int) map[int]string {
	pics := make(map[int]string)
	pics[0] = ""
	if cardType > 0 {
		pics[0] = constkey.CardsSmallDefaultPics[cardType]
	}

	rpcImg := new(file.Upload).Init()
	defer rpcImg.Close()
	var replyImgs = map[int]file2.ReplyFileInfo{}
	err := rpcImg.GetImageByIds(ctx, imgIds, &replyImgs)
	if err != nil {
		return pics
	}
	for _, v := range replyImgs {
		pics[v.Id] = v.Path
	}
	return pics
}

//获取商家的主营id
func getBusMainBindId(ctx context.Context, busid int) (mainBindId int) {
	rpcBus := new(bus.Bus).Init()
	defer rpcBus.Close()
	args := bus2.ArgsSingleBus{BusId: busid}
	reply := &bus2.ReplySingleBus{}
	err := rpcBus.GetByBusid(ctx, &args, reply)
	if err != nil {
		return 0
	}

	mainBindId = reply.MainBindId
	return
}

//获取包含的单项目
func getIncSingles(ctx context.Context, singleIds []int) (reply map[int]cards.IncSingleDetail) {
	reply = map[int]cards.IncSingleDetail{}
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
		minp, _ := strconv.ParseFloat(single.MinPrice, 64)
		maxp, _ := strconv.ParseFloat(single.MaxPrice, 64)
		if minp > 0 || maxp > 0 {
			realPrice = fmt.Sprintf("%s-%s", single.MinPrice, single.MaxPrice)
		}
		reply[single.SingleId] = cards.IncSingleDetail{
			IncSingle: cards.IncSingle{
				SingleID: single.SingleId,
				Num:      0,
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
func getIncSingles2(ctx context.Context, shopId int, status, isDel string, singleIds []int, maps []map[string]interface{}) (reply []cards.IncSingleDetail2, singleNum int) {
	reply = []cards.IncSingleDetail2{}
	sspIds := functions.ArrayValue2Array("ssp_id", maps)
	singleIds = functions.ArrayUniqueInt(singleIds)
	var sspId2specNames = make(map[string]string) //规格组合名
	if len(sspIds) > 0 {
		//单项目规格数据
		var singleSpecRes cards.ReplySubServer2
		if err := new(SingleLogic).GetSingleSpecBySspId(&cards.ArgsSubSpecID{SspIds: sspIds}, &singleSpecRes); err != nil {
			return
		}
		if singleSpecRes.SubServer != nil {
			parentSpecMap := make(map[int][]int, 0)
			for _, v := range *singleSpecRes.SubServer {
				parentSpecMap[v.SspId] = functions.StrExplode2IntArr(v.SubServerIds, ",")
			}
			subSpecFullName := make(map[int]string) //子规格全名 格式：map[1]="技师等级-初级" map[2]="技师等级-中级"
			if singleSpecRes.Specs != nil {         //具体规格，格式：一级规格包含的二级规格
				singleSpecResSpecs := *singleSpecRes.Specs
				for _, parentSpec := range singleSpecResSpecs {
					for _, subSpec := range parentSpec.Sub {
						subSpecFullName[subSpec.Id] = parentSpec.Name + ":" + subSpec.Name
					}
				}
			}
			for sspId, specIds := range parentSpecMap {
				sspIdStr := strconv.Itoa(sspId)
				for _, specId := range specIds {
					sspId2specNames[sspIdStr] = sspId2specNames[sspIdStr] + subSpecFullName[specId] + "-"
				}
				sspId2specNames[sspIdStr] = sspId2specNames[sspIdStr][:len(sspId2specNames[sspIdStr])-1]
			}
		}
	}

	var ssmMapRebuild = make(map[string]interface{})
	ssm := new(models.ShopSingleModel).Init()
	if shopId > 0 { //获取门店未删除和上架的单项目
		shopSingleWhere := []base.WhereItem{
			{ssm.Field.F_shop_id, shopId},
		}
		if len(singleIds) > 0 {
			shopSingleWhere = append(shopSingleWhere, base.WhereItem{ssm.Field.F_single_id, []interface{}{"IN", singleIds}})
		}
		if status != "" {
			statusInt, _ := strconv.Atoi(status)
			shopSingleWhere = append(shopSingleWhere, base.WhereItem{ssm.Field.F_status, statusInt})
		}
		if isDel != "" {
			isDelInt, _ := strconv.Atoi(isDel)
			shopSingleWhere = append(shopSingleWhere, base.WhereItem{ssm.Field.F_is_del, isDelInt})
		}
		ssmMaps := ssm.SelectRcardsByWherePage(shopSingleWhere, 0, 0, []string{
			ssm.Field.F_single_id, ssm.Field.F_name, ssm.Field.F_changed_real_price,
			ssm.Field.F_changed_min_price, ssm.Field.F_changed_max_price,
		})
		ssmMapRebuild = functions.ArrayRebuild(ssm.Field.F_single_id, ssmMaps)
		singleIds = functions.ArrayValue2Array(ssm.Field.F_single_id, ssmMaps)
		singleNum = ssm.GetTotalWhere(shopSingleWhere)
	}

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
	var replyMap = make(map[int]cards.IncSingleDetail2)

	for _, single := range singlesStruct {
		realPrice := single.RealPrice
		minp, _ := strconv.ParseFloat(single.MinPrice, 64)
		maxp, _ := strconv.ParseFloat(single.MaxPrice, 64)
		if minp > 0 || maxp > 0 { //总店单项目的最大和最小价格
			realPrice = fmt.Sprintf("%v-%v", single.MinPrice, single.MaxPrice)
		}
		if shopId > 0 && len(ssmMapRebuild) > 0 {
			ssmSingleMap := ssmMapRebuild[strconv.Itoa(single.SingleId)].(map[string]interface{})
			if ssmSingleMap[ssm.Field.F_name].(string) != "" {
				single.Name = ssmSingleMap[ssm.Field.F_name].(string)
			}
			shopMixPrice, _ := strconv.ParseFloat(ssmSingleMap[ssm.Field.F_changed_min_price].(string), 64)
			shopMaxPrice, _ := strconv.ParseFloat(ssmSingleMap[ssm.Field.F_changed_max_price].(string), 64)
			if shopMixPrice > 0 || shopMaxPrice > 0 {
				realPrice = fmt.Sprintf("%v-%v", shopMixPrice, shopMaxPrice)
			}
		}
		r := cards.IncSingleDetail2{
			IncSingle: cards.IncSingle{
				SingleID: single.SingleId,
				Num:      0,
			},
			Name:        single.Name,
			Price:       single.Price,
			RealPrice:   realPrice,
			ImgId:       single.ImgId,
			ImgUrl:      imgs[single.ImgId],
			ServiceTime: single.ServiceTime,
		}
		replyMap[single.SingleId] = r
		reply = append(reply, r)
	}
	for _, m := range maps {
		if m["ssp_id"] != nil && m["ssp_id"].(string) != "0" {
			singleId, _ := strconv.Atoi(m["single_id"].(string))
			sspId, _ := strconv.Atoi(m["ssp_id"].(string))
			r := replyMap[singleId]
			r.SspId = sspId
			r.SpecNames = sspId2specNames[m["ssp_id"].(string)]
			reply = append(reply, r)
		}
	}
	return
}

/*
if m["ssp_id"].(string) != "0" {
			sspId, _ := strconv.Atoi(m["ssp_id"].(string))
			incSingle2 := reply[single.SingleId]
			incSingle2.SspId = sspId
		}
*/

//总店-删除
func (s *SmLogic) DeleteSmLogic(ctx context.Context, args *cards.ArgsDeleteSm, reply *bool) bool {
	//实例化模型
	model := new(models.SmModel).Init()
	model.Model.Begin()
	//修改数据
	data := map[string]interface{}{
		model.Field.F_is_del:   model.DelStatus(),
		model.Field.F_del_time: time.Now().Unix(),
	}
	if b := model.UpdateBySmids(args.SmIds, data); !b {
		model.Model.RollBack()
		return b
	}
	//实例化分店模型
	shopModel := new(models.ShopSmModel).Init()

	//修改数据
	if b := shopModel.UpdateBySmids(args.SmIds, data); !b {
		model.Model.RollBack()
		return b
	}
	//同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.SmIds, cards.ITEM_TYPE_sm) {
		model.Model.RollBack()
		return false
	}
	model.Model.Commit()
	return true
}

//分店-删除
func (s *SmLogic) DeleteShopSmLogic(ctx context.Context, args *cards.ArgsDeleteShopSm, reply *bool) bool {
	//实例化分店模型
	shopModel := new(models.ShopSmModel).Init()
	shopId, _ := args.BsToken.GetShopId()
	data := map[string]interface{}{
		shopModel.Field.F_is_del:   shopModel.DelStatus(),
		shopModel.Field.F_del_time: time.Now().Unix(),
	}
	shopModel.Model.Begin()
	if b := shopModel.UpdateByShopSmids(args.SmIds, shopId, data); !b {
		shopModel.Model.RollBack()
		return b
	}
	//同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.SmIds, cards.ITEM_TYPE_sm, shopId) {
		shopModel.Model.RollBack()
		return false
	}
	shopModel.Model.Commit()
	return true
}
