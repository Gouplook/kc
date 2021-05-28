package logics

import (
	"context"
	"encoding/json"
	"fmt"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"strconv"
	"strings"
	"time"

	"git.900sui.cn/kc/kcgin/logs"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/mapstructure"
	"git.900sui.cn/kc/rpcCards/common/models"
	"git.900sui.cn/kc/rpcCards/common/tools"
	"git.900sui.cn/kc/rpcCards/constkey"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/client/bus"
	cardsClient "git.900sui.cn/kc/rpcinterface/client/cards"
	cards4 "git.900sui.cn/kc/rpcinterface/client/elastic/cards"
	"git.900sui.cn/kc/rpcinterface/client/file"
	order2 "git.900sui.cn/kc/rpcinterface/client/order"
	"git.900sui.cn/kc/rpcinterface/client/product"
	"git.900sui.cn/kc/rpcinterface/client/public"
	task2 "git.900sui.cn/kc/rpcinterface/client/task"
	cards2 "git.900sui.cn/kc/rpcinterface/client/task/cards"
	bus2 "git.900sui.cn/kc/rpcinterface/interface/bus"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	file2 "git.900sui.cn/kc/rpcinterface/interface/file"
	"git.900sui.cn/kc/rpcinterface/interface/order"
	product2 "git.900sui.cn/kc/rpcinterface/interface/product"
	public2 "git.900sui.cn/kc/rpcinterface/interface/public"
	task1 "git.900sui.cn/kc/rpcinterface/interface/task"
	cards3 "git.900sui.cn/kc/rpcinterface/interface/task/cards"
)

type ItemLogic struct {
}

//单项目结构
type ItemSingle struct {
	SingleId int
	Num      int
	Name     string
}

//门店-项目的es文档任务添加到rabbitmq任务交换机
//@param itemIds []int  项目id
func setShopItem(ctx context.Context, itemIds []int, shopId int, itemType int, shopItemIds ...[]int) bool {
	rpcTask := new(cards2.ShopItems).Init()
	var reply = false
	switch itemType {
	case cards.ITEM_TYPE_single:
		mShopSingle := new(models.ShopSingleModel).Init()
		var shopSingles []map[string]interface{}
		if len(shopItemIds) > 0 && len(shopItemIds[0]) > 0 {
			shopSingles = mShopSingle.GetBySsids(shopItemIds[0])
		} else {
			if shopId > 0 {
				shopSingles = mShopSingle.GetByShopidAndSingleids(shopId, itemIds, []string{
					mShopSingle.Field.F_ss_id,
					mShopSingle.Field.F_single_id,
					mShopSingle.Field.F_shop_id,
				})
			} else {
				shopSingles = mShopSingle.GetBySingleids(itemIds)
			}
		}

		for _, shopSingle := range shopSingles {
			shopItemId, _ := strconv.Atoi(shopSingle[mShopSingle.Field.F_ss_id].(string))
			itemId, _ := strconv.Atoi(shopSingle[mShopSingle.Field.F_single_id].(string))
			itemShopId, _ := strconv.Atoi(shopSingle[mShopSingle.Field.F_shop_id].(string))
			rpcTask.SetItems(ctx, &cards3.ShopItems{
				ShopItemId: shopItemId,
				ItemId:     itemId,
				ShopId:     itemShopId,
				ItemType:   cards.ITEM_TYPE_single,
			}, &reply)
		}
	case cards.ITEM_TYPE_sm:
		mShopSm := new(models.ShopSmModel).Init()
		var shopSms []map[string]interface{}
		if len(shopItemIds) > 0 && len(shopItemIds[0]) > 0 {
			shopSms = mShopSm.GetByIds(shopItemIds[0])
		} else {
			if shopId > 0 {
				shopSms = mShopSm.GetByShopidAdSmids(shopId, itemIds)
			} else {
				shopSms = mShopSm.GetBySmids(itemIds)
			}
		}

		for _, sm := range shopSms {
			shopSmid, _ := strconv.Atoi(sm[mShopSm.Field.F_id].(string))
			smid, _ := strconv.Atoi(sm[mShopSm.Field.F_sm_id].(string))
			itemShopId, _ := strconv.Atoi(sm[mShopSm.Field.F_shop_id].(string))
			rpcTask.SetItems(ctx, &cards3.ShopItems{
				ShopItemId: shopSmid,
				ItemId:     smid,
				ShopId:     itemShopId,
				ItemType:   cards.ITEM_TYPE_sm,
			}, &reply)
		}
	case cards.ITEM_TYPE_card:
		mShopCard := new(models.ShopCardModel).Init()
		var shopCards []map[string]interface{}
		if len(shopItemIds) > 0 && len(shopItemIds[0]) > 0 {
			shopCards = mShopCard.GetByIDs(shopItemIds[0])
		} else {
			if shopId > 0 {
				shopCards = mShopCard.GetByShopIDAndCardIDs(shopId, itemIds)
			} else {
				shopCards = mShopCard.GetByCardIDs(itemIds)
			}
		}

		for _, card := range shopCards {
			shopCardid, _ := strconv.Atoi(card[mShopCard.Field.F_id].(string))
			cardId, _ := strconv.Atoi(card[mShopCard.Field.F_card_id].(string))
			itemShopId, _ := strconv.Atoi(card[mShopCard.Field.F_shop_id].(string))
			rpcTask.SetItems(ctx, &cards3.ShopItems{
				ShopItemId: shopCardid,
				ItemId:     cardId,
				ShopId:     itemShopId,
				ItemType:   cards.ITEM_TYPE_card,
			}, &reply)
		}
	case cards.ITEM_TYPE_hcard:
		mShopHcard := new(models.ShopHcardModel).Init()
		var shopHcards []map[string]interface{}
		if len(shopItemIds) > 0 && len(shopItemIds[0]) > 0 {
			shopHcards = mShopHcard.GetByIDs(shopItemIds[0])
		} else {
			if shopId > 0 {
				shopHcards = mShopHcard.GetByShopIDAndHcardIDs(shopId, itemIds)
			} else {
				shopHcards = mShopHcard.GetByHcardIDs(itemIds)
			}
		}

		for _, hcard := range shopHcards {
			shopHcardid, _ := strconv.Atoi(hcard[mShopHcard.Field.F_id].(string))
			hcardId, _ := strconv.Atoi(hcard[mShopHcard.Field.F_hcard_id].(string))
			itemShopId, _ := strconv.Atoi(hcard[mShopHcard.Field.F_shop_id].(string))
			rpcTask.SetItems(ctx, &cards3.ShopItems{
				ShopItemId: shopHcardid,
				ItemId:     hcardId,
				ShopId:     itemShopId,
				ItemType:   cards.ITEM_TYPE_hcard,
			}, &reply)
		}
	case cards.ITEM_TYPE_ncard:
		mShopNcard := new(models.ShopNCardModel).Init()
		var shopNcards []map[string]interface{}
		if len(shopItemIds) > 0 && len(shopItemIds[0]) > 0 {
			shopNcards = mShopNcard.GetByIDs(shopItemIds[0])
		} else {
			if shopId > 0 {
				shopNcards = mShopNcard.GetByShopIDAndNCardIDs(shopId, itemIds)
			} else {
				shopNcards = mShopNcard.GetByNCardIDs(itemIds)
			}
		}

		for _, ncard := range shopNcards {
			shopNcardid, _ := strconv.Atoi(ncard[mShopNcard.Field.F_id].(string))
			ncardId, _ := strconv.Atoi(ncard[mShopNcard.Field.F_ncard_id].(string))
			itemShopId, _ := strconv.Atoi(ncard[mShopNcard.Field.F_shop_id].(string))
			rpcTask.SetItems(ctx, &cards3.ShopItems{
				ShopItemId: shopNcardid,
				ItemId:     ncardId,
				ShopId:     itemShopId,
				ItemType:   cards.ITEM_TYPE_ncard,
			}, &reply)
		}
	case cards.ITEM_TYPE_hncard:
		mShopHncard := new(models.ShopHNCardModel).Init()
		var shopHncards []map[string]interface{}
		if len(shopItemIds) > 0 && len(shopItemIds[0]) > 0 {
			shopHncards = mShopHncard.GetByIDs(shopItemIds[0])
		} else {
			if shopId > 0 {
				shopHncards = mShopHncard.GetByShopIDAndHNCardIDs(shopId, itemIds)
			} else {
				shopHncards = mShopHncard.GetByHNCardIDs(itemIds)
			}
		}

		for _, hncard := range shopHncards {
			shopHncardid, _ := strconv.Atoi(hncard[mShopHncard.Field.F_id].(string))
			hncardId, _ := strconv.Atoi(hncard[mShopHncard.Field.F_hncard_id].(string))
			itemShopId, _ := strconv.Atoi(hncard[mShopHncard.Field.F_shop_id].(string))
			rpcTask.SetItems(ctx, &cards3.ShopItems{
				ShopItemId: shopHncardid,
				ItemId:     hncardId,
				ShopId:     itemShopId,
				ItemType:   cards.ITEM_TYPE_hncard,
			}, &reply)
		}
	case cards.ITEM_TYPE_rcard:
		mShopRcard := new(models.ShopRcardModel).Init()
		var shopRcards []map[string]interface{}
		if len(shopItemIds) > 0 && len(shopItemIds[0]) > 0 {
			shopRcards = mShopRcard.GetByIds(shopItemIds[0])
		} else {
			if shopId > 0 {
				shopRcards = mShopRcard.GetByShopidAndRcardids(shopId, itemIds)
			} else {
				shopRcards = mShopRcard.GetByRcardids(itemIds)
			}
		}

		for _, rcard := range shopRcards {
			shopRcardid, _ := strconv.Atoi(rcard[mShopRcard.Field.F_id].(string))
			rcardId, _ := strconv.Atoi(rcard[mShopRcard.Field.F_rcard_id].(string))
			itemShopId, _ := strconv.Atoi(rcard[mShopRcard.Field.F_shop_id].(string))
			rpcTask.SetItems(ctx, &cards3.ShopItems{
				ShopItemId: shopRcardid,
				ItemId:     rcardId,
				ShopId:     itemShopId,
				ItemType:   cards.ITEM_TYPE_rcard,
			}, &reply)
		}
	case cards.ITEM_TYPE_icard:
		mShopIcard := new(models.ShopIcardModel).Init()
		var shopIcards []map[string]interface{}
		if len(shopItemIds) > 0 && len(shopItemIds[0]) > 0 {
			shopIcards = mShopIcard.GetByIds(shopItemIds[0])
		} else {
			if shopId > 0 {
				shopIcards = mShopIcard.GetByShopidAndIcardids(shopId, itemIds)
			} else {
				shopIcards = mShopIcard.GetByIcardids(itemIds)
			}
		}

		for _, icard := range shopIcards {
			shopIcardid, _ := strconv.Atoi(icard[mShopIcard.Field.F_id].(string))
			icardId, _ := strconv.Atoi(icard[mShopIcard.Field.F_icard_id].(string))
			itemShopId, _ := strconv.Atoi(icard[mShopIcard.Field.F_shop_id].(string))
			rpcTask.SetItems(ctx, &cards3.ShopItems{
				ShopItemId: shopIcardid,
				ItemId:     icardId,
				ShopId:     itemShopId,
				ItemType:   cards.ITEM_TYPE_icard,
			}, &reply)
		}
	}
	return true
}

//为es的shop-items文档获取项目信息
func (i *ItemLogic) GetItemInfo4Es(shopId, itemId, itemType int) (reply cards.ReplyItemInfo4Es, err error) {
	reply = cards.ReplyItemInfo4Es{}
	switch itemType {
	case cards.ITEM_TYPE_single:
		mShopItem := new(models.ShopSingleModel).Init()
		shopSingleInfo := mShopItem.GetByShopidAndSingleid(shopId, itemId)
		if len(shopSingleInfo) == 0 {
			return
		}
		mItem := new(models.SingleModel).Init()
		singleInfo := mItem.GetBySingleId(itemId, []string{
			mItem.Field.F_name,
		})
		itemShopPrice, _ := strconv.ParseFloat(shopSingleInfo[mShopItem.Field.F_changed_real_price].(string), 64)
		minPrice, _ := strconv.ParseFloat(shopSingleInfo[mShopItem.Field.F_changed_min_price].(string), 64)
		if minPrice > 0 {
			itemShopPrice = minPrice
		}
		status, _ := strconv.Atoi(shopSingleInfo[mShopItem.Field.F_status].(string))
		reply = cards.ReplyItemInfo4Es{
			ItemName:       singleInfo[mItem.Field.F_name].(string),
			ItemShopPrice:  itemShopPrice,
			ItemShopStatus: status,
		}
		return
	case cards.ITEM_TYPE_sm:
		mShopItem := new(models.ShopSmModel).Init()
		shopSmInfo := mShopItem.GetByShopidAdSmid(shopId, itemId)
		if len(shopSmInfo) == 0 {
			return
		}
		mItem := new(models.SmModel).Init()
		smInfo := mItem.GetBySmid(itemId, []string{
			mItem.Field.F_name,
			mItem.Field.F_real_price,
		})
		itemShopPrice, _ := strconv.ParseFloat(smInfo[mItem.Field.F_real_price].(string), 64)
		status, _ := strconv.Atoi(shopSmInfo[mShopItem.Field.F_status].(string))
		reply = cards.ReplyItemInfo4Es{
			ItemName:       smInfo[mItem.Field.F_name].(string),
			ItemShopPrice:  itemShopPrice,
			ItemShopStatus: status,
		}
		return
	case cards.ITEM_TYPE_card:
		mShopItem := new(models.ShopCardModel).Init()
		shopCardInfo := mShopItem.GetByShopIDAndCardIDs(shopId, []int{itemId})
		if len(shopCardInfo) == 0 {
			return
		}
		mItem := new(models.CardModel).Init()
		cardInfo := mItem.GetByCardID(itemId, mItem.Field.F_real_price, mItem.Field.F_name)
		itemShopPrice, _ := strconv.ParseFloat(cardInfo[mItem.Field.F_real_price].(string), 64)
		status, _ := strconv.Atoi(shopCardInfo[0][mShopItem.Field.F_status].(string))
		reply = cards.ReplyItemInfo4Es{
			ItemName:       cardInfo[mItem.Field.F_name].(string),
			ItemShopPrice:  itemShopPrice,
			ItemShopStatus: status,
		}
		return
	case cards.ITEM_TYPE_hcard:
		mShopItem := new(models.ShopHcardModel).Init()
		shopHcardInfo := mShopItem.GetByShopIDAndHcardIDs(shopId, []int{itemId})
		if len(shopHcardInfo) == 0 {
			return
		}
		mItem := new(models.HcardModel).Init()
		hcardInfo := mItem.GetHcardByID(itemId, mItem.Field.F_name, mItem.Field.F_real_price)
		itemShopPrice, _ := strconv.ParseFloat(hcardInfo[mItem.Field.F_real_price].(string), 64)
		status, _ := strconv.Atoi(shopHcardInfo[0][mShopItem.Field.F_status].(string))
		reply = cards.ReplyItemInfo4Es{
			ItemName:       hcardInfo[mItem.Field.F_name].(string),
			ItemShopPrice:  itemShopPrice,
			ItemShopStatus: status,
		}
		return
	case cards.ITEM_TYPE_ncard:
		mShopItem := new(models.ShopNCardModel).Init()
		shopNcardInfo := mShopItem.GetByShopIDAndNCardIDs(shopId, []int{itemId})
		if len(shopNcardInfo) == 0 {
			return
		}
		mItem := new(models.NCardModel).Init()
		ncardInfo := mItem.GetByNCardID(itemId, mItem.Field.F_real_price, mItem.Field.F_name)
		itemShopPrice, _ := strconv.ParseFloat(ncardInfo[mItem.Field.F_real_price].(string), 64)
		status, _ := strconv.Atoi(shopNcardInfo[0][mShopItem.Field.F_status].(string))
		reply = cards.ReplyItemInfo4Es{
			ItemName:       ncardInfo[mItem.Field.F_name].(string),
			ItemShopPrice:  itemShopPrice,
			ItemShopStatus: status,
		}
		return
	case cards.ITEM_TYPE_hncard:
		mShopItem := new(models.ShopHNCardModel).Init()
		shopHncardInfo := mShopItem.GetByShopIDAndHNCardIDs(shopId, []int{itemId})
		if len(shopHncardInfo) == 0 {
			return
		}
		mItem := new(models.HNCardModel).Init()
		hncardInfo := mItem.GetByHNCardID(itemId, mItem.Field.F_name, mItem.Field.F_real_price)
		itemShopPrice, _ := strconv.ParseFloat(hncardInfo[mItem.Field.F_real_price].(string), 64)
		status, _ := strconv.Atoi(shopHncardInfo[0][mShopItem.Field.F_status].(string))
		reply = cards.ReplyItemInfo4Es{
			ItemName:       hncardInfo[mItem.Field.F_name].(string),
			ItemShopPrice:  itemShopPrice,
			ItemShopStatus: status,
		}
		return
	case cards.ITEM_TYPE_rcard:
		mShopItem := new(models.ShopRcardModel).Init()
		shopRcardInfo := mShopItem.GetByShopidAndRcardid(shopId, itemId)
		if len(shopRcardInfo) == 0 {
			return
		}
		mItem := new(models.RcardModel).Init()
		rcardInfo := mItem.GetByRcardId(itemId, []string{mItem.Field.F_name, mItem.Field.F_real_price})
		itemShopPrice, _ := strconv.ParseFloat(rcardInfo[mItem.Field.F_real_price].(string), 64)
		status, _ := strconv.Atoi(shopRcardInfo[mShopItem.Field.F_status].(string))
		reply = cards.ReplyItemInfo4Es{
			ItemName:       rcardInfo[mItem.Field.F_name].(string),
			ItemShopPrice:  itemShopPrice,
			ItemShopStatus: status,
		}
		return
	case cards.ITEM_TYPE_icard:
		mShopItem := new(models.ShopIcardModel).Init()
		shopRcardInfo := mShopItem.GetByShopidAndRcardid(shopId, itemId)
		if len(shopRcardInfo) == 0 {
			return
		}
		mItem := new(models.IcardModel).Init()
		rcardInfo := mItem.GetByIcardId(itemId, []string{mItem.Field.F_name, mItem.Field.F_real_price})
		itemShopPrice, _ := strconv.ParseFloat(rcardInfo[mItem.Field.F_real_price].(string), 64)
		status, _ := strconv.Atoi(shopRcardInfo[mShopItem.Field.F_status].(string))
		reply = cards.ReplyItemInfo4Es{
			ItemName:       rcardInfo[mItem.Field.F_name].(string),
			ItemShopPrice:  itemShopPrice,
			ItemShopStatus: status,
		}
		return
	}

	return
}

//根据项目的门店id，获取数据
func (i *ItemLogic) GetItemsBySsids(ctx context.Context, args *cards.ArgsGetItemsBySsids) (reply map[cards.SsId]cards.ItemBase, err error) {
	if len(args.SsIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	reply = make(map[cards.SsId]cards.ItemBase)
	switch args.ItemType {
	case cards.ITEM_TYPE_single:
		reply, err = i.getSingls(args.SsIds)
	case cards.ITEM_TYPE_sm:
		reply, err = i.getSms(args.SsIds)
	case cards.ITEM_TYPE_card:
		reply, err = i.getCards(ctx, args.SsIds)
	case cards.ITEM_TYPE_hcard:
		reply, err = i.getHcards(args.SsIds)
	case cards.ITEM_TYPE_ncard:
		reply, err = i.getNcards(args.SsIds)
	case cards.ITEM_TYPE_hncard:
		reply, err = i.getHncards(args.SsIds)
	case cards.ITEM_TYPE_rcard:
		reply, err = i.getRards(ctx, args.SsIds)
	case cards.ITEM_TYPE_icard:
		reply, err = i.getIcards(ctx, args.SsIds)
	}
	return
}

//获取单项目
func (i *ItemLogic) getSingls(ssIds []int) (reply map[cards.SsId]cards.ItemBase, err error) {
	mShopItem := new(models.ShopSingleModel).Init()
	shopSingles := mShopItem.GetBySsids(ssIds)
	if len(shopSingles) == 0 {
		return
	}
	reply = make(map[cards.SsId]cards.ItemBase)
	// 获取单项目IDs
	var smIds []int
	for _, shopSm := range shopSingles {
		singleId, _ := strconv.Atoi(shopSm[mShopItem.Field.F_single_id].(string))
		ssId, _ := strconv.Atoi(shopSm[mShopItem.Field.F_ss_id].(string))
		status, _ := strconv.Atoi(shopSm[mShopItem.Field.F_status].(string))
		shopId, _ := strconv.Atoi(shopSm[mShopItem.Field.F_shop_id].(string))
		shopSales, _ := strconv.Atoi(shopSm[mShopItem.Field.F_sales].(string))
		shopRealPrice, _ := strconv.ParseFloat(shopSm[mShopItem.Field.F_changed_real_price].(string), 64)
		changedMinPrice, _ := strconv.ParseFloat(shopSm[mShopItem.Field.F_changed_min_price].(string), 64)
		if changedMinPrice > 0 {
			shopRealPrice = changedMinPrice
		}
		reply[cards.SsId(ssId)] = cards.ItemBase{
			SsId:          ssId,
			Status:        status,
			ItemId:        singleId,
			ShopId:        shopId,
			ShopSales:     shopSales,
			ShopRealPrice: shopRealPrice,
		}
		smIds = append(smIds, singleId)
	}
	// 获取单项目数据
	mSm := new(models.SingleModel).Init(mShopItem.Model.GetOrmer())
	sms := mSm.GetBySingleids(smIds, []string{
		mSm.Field.F_single_id,
		mSm.Field.F_name,
		mSm.Field.F_img_id,
		mSm.Field.F_price,
		mSm.Field.F_real_price,
		mSm.Field.F_service_time,
	})
	var rebuildSms = make(map[int]struct {
		SmId          int
		Name          string
		ImgId         int
		Price         float64
		RealPrice     float64
		HasGiveSignle int
		ServicePeriod int
	})
	for _, sm := range sms {
		singleId, _ := strconv.Atoi(sm[mSm.Field.F_single_id].(string))
		imgId, _ := strconv.Atoi(sm[mSm.Field.F_img_id].(string))
		price, _ := strconv.ParseFloat(sm[mSm.Field.F_price].(string), 64)
		realPrice, _ := strconv.ParseFloat(sm[mSm.Field.F_real_price].(string), 64)
		servicePeriod, _ := strconv.Atoi(sm[mSm.Field.F_service_time].(string))

		rebuildSms[singleId] = struct {
			SmId          int
			Name          string
			ImgId         int
			Price         float64
			RealPrice     float64
			HasGiveSignle int
			ServicePeriod int
		}{SmId: singleId, Name: sm[mSm.Field.F_name].(string), ImgId: imgId, Price: price, RealPrice: realPrice, ServicePeriod: servicePeriod}
	}
	//获取适用门店

	//获取单项目的名称
	for ssId, v := range reply {
		v.ItemName = rebuildSms[v.ItemId].Name
		v.Price = rebuildSms[v.ItemId].Price
		v.RealPrice = rebuildSms[v.ItemId].RealPrice
		//v.CableShopIds = functions.ArrayUniqueInt(smShopIds[v.ItemId]) // 适用门店ID
		v.ServicePeriod = rebuildSms[v.ItemId].ServicePeriod
		v.ImgId = rebuildSms[v.ItemId].ImgId
		reply[ssId] = v
	}
	return
}

//根据sspIds 获取价格规格名称组合
func (i *ItemLogic) getSpecNames(sspIds []int) map[int]string {
	var sspId2specMap = make(map[int]string)
	if len(sspIds) == 0 {
		return sspId2specMap
	}
	//单项目规格数据
	var singleSpecRes cards.ReplySubServer2
	if err := new(SingleLogic).GetSingleSpecBySspId(&cards.ArgsSubSpecID{SspIds: sspIds}, &singleSpecRes); err != nil {
		return sspId2specMap
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
			for _, specId := range specIds {
				sspId2specMap[sspId] = sspId2specMap[sspId] + subSpecFullName[specId] + "-"
			}
			sspId2specMap[sspId] = sspId2specMap[sspId][:len(sspId2specMap[sspId])-1]
		}
	}
	//sspModel := new(models.SingleSpecPriceModel).Init()
	//sspMaps := sspModel.GetBySspids(sspIds)
	//var specIds []int
	//for _, sspMap := range sspMaps {
	//	specIds = append(specIds, functions.StrExplode2IntArr(sspMap[sspModel.Field.F_spec_ids].(string), ",")...)
	//}
	//var specMap = make(map[string]string)
	//if len(specIds) > 0 {
	//	sbsModel := new(models.SingleBusSpecModel).Init()
	//	sbsMaps := sbsModel.GetBySpecids(specIds)
	//	for _, sbsMap := range sbsMaps {
	//		specMap[sbsMap[sbsModel.Field.F_spec_id].(string)] = sbsMap[sbsModel.Field.F_name].(string)
	//	}
	//}
	//for _, sspMap := range sspMaps {
	//	split := strings.Split(sspMap[sspModel.Field.F_spec_ids].(string), ",")
	//	ss := []string{}
	//	for _, s := range split {
	//		ss = append(ss, specMap[s])
	//	}
	//	sspId, _ := strconv.Atoi(sspMap[sspModel.Field.F_ssp_id].(string))
	//	sspId2specMap[sspId] = strings.Join(ss, "-")
	//}
	return sspId2specMap
}

//获取套餐
func (i *ItemLogic) getSms(ssIds []int) (reply map[cards.SsId]cards.ItemBase, err error) {
	mShopItem := new(models.ShopSmModel).Init()
	shopSms := mShopItem.GetByIds(ssIds)
	if len(shopSms) == 0 {
		return
	}
	reply = make(map[cards.SsId]cards.ItemBase)
	//获取套餐id
	var smIds []int
	for _, shopSm := range shopSms {
		smId, _ := strconv.Atoi(shopSm[mShopItem.Field.F_sm_id].(string))
		ssId, _ := strconv.Atoi(shopSm[mShopItem.Field.F_id].(string))
		status, _ := strconv.Atoi(shopSm[mShopItem.Field.F_status].(string))
		shopId, _ := strconv.Atoi(shopSm[mShopItem.Field.F_shop_id].(string))
		shopSales, _ := strconv.Atoi(shopSm[mShopItem.Field.F_sales].(string))
		reply[cards.SsId(ssId)] = cards.ItemBase{
			SsId:      ssId,
			Status:    status,
			ItemId:    smId,
			ShopId:    shopId,
			ShopSales: shopSales,
		}
		smIds = append(smIds, smId)
	}
	//获取套餐数据
	mSm := new(models.SmModel).Init()
	sms := mSm.GetBySmids(smIds, []string{
		mSm.Field.F_sm_id,
		mSm.Field.F_name,
		mSm.Field.F_img_id,
		mSm.Field.F_price,
		mSm.Field.F_real_price,
		mSm.Field.F_has_give_signle,
		mSm.Field.F_service_period,
		mSm.Field.F_validcount,
	})

	var rebuildSms = make(map[int]struct {
		SmId          int
		Name          string
		ImgId         int
		Price         float64
		RealPrice     float64
		HasGiveSignle int
		ServicePeriod int
		ValidCount    int
	})
	for _, sm := range sms {
		smId, _ := strconv.Atoi(sm[mSm.Field.F_sm_id].(string))
		imgId, _ := strconv.Atoi(sm[mSm.Field.F_img_id].(string))
		price, _ := strconv.ParseFloat(sm[mSm.Field.F_price].(string), 64)
		realPrice, _ := strconv.ParseFloat(sm[mSm.Field.F_real_price].(string), 64)
		hasGiveSignle, _ := strconv.Atoi(sm[mSm.Field.F_has_give_signle].(string))
		servicePeriod, _ := strconv.Atoi(sm[mSm.Field.F_service_period].(string))
		validCount, _ := strconv.Atoi(sm[mSm.Field.F_validcount].(string))

		rebuildSms[smId] = struct {
			SmId          int
			Name          string
			ImgId         int
			Price         float64
			RealPrice     float64
			HasGiveSignle int
			ServicePeriod int
			ValidCount    int
		}{SmId: smId, Name: sm[mSm.Field.F_name].(string), ImgId: imgId, Price: price, RealPrice: realPrice, HasGiveSignle: hasGiveSignle,
			ServicePeriod: servicePeriod, ValidCount: validCount}
	}

	//获取适用门店
	mSmShop := new(models.SmShopModel).Init()
	smShops := mSmShop.GetBySmids(smIds)
	var smShopIds map[int][]int
	smShopIds = make(map[int][]int)
	for _, smShop := range smShops {
		smId, _ := strconv.Atoi(smShop[mSmShop.Field.F_sm_id].(string))
		shopId, _ := strconv.Atoi(smShop[mSmShop.Field.F_shop_id].(string))
		if shopId == 0 {
			smShopIds[smId] = []int{0} //适用所有门店
		} else {
			smShopIds[smId] = append(smShopIds[smId], shopId)
		}
	}

	var singleIds = []int{}
	//获取包含的单项目
	mSmSingle := new(models.SmSingleModel).Init()
	smSingles := mSmSingle.GetBySmids(smIds)
	singleIds = functions.ArrayValue2Array(mSmSingle.Field.F_single_id, smSingles)
	sspIds := functions.ArrayValue2Array(mSmSingle.Field.F_ssp_id, smSingles)
	sspId2specMap := i.getSpecNames(sspIds)

	//获取赠送的单项目数据
	mSmGive := new(models.SmGiveModel).Init()
	giveSingles := mSmGive.GetBySmids(smIds)
	singleIds = append(singleIds, functions.ArrayValue2Array(mSmGive.Field.F_single_id, giveSingles)...)
	singles := i.getSinglesBySingleids(singleIds)
	//获取赠品描述数据
	mSmGiveDesc := new(models.SmGiveDescModel).Init()
	giveSinglesDesc := mSmGiveDesc.GetBySmids(smIds)
	giveDescMap := make(map[int][]cards.GiveSingleDesc)
	for _, dv := range giveSinglesDesc {
		smId, _ := strconv.Atoi(dv[mSmGiveDesc.Field.F_sm_id].(string))
		descStr := dv[mSmGiveDesc.Field.F_desc].(string)
		desc := []cards.GiveSingleDesc{}
		json.Unmarshal([]byte(descStr), &desc)
		giveDescMap[smId] = desc
	}
	//获取单项目数据

	var smSinglesMap = make(map[int][]cards.ItemSingle)
	for _, v := range smSingles {
		singleId, _ := strconv.Atoi(v[mSmSingle.Field.F_single_id].(string))
		smId, _ := strconv.Atoi(v[mSmSingle.Field.F_sm_id].(string))
		num, _ := strconv.Atoi(v[mSmSingle.Field.F_num].(string))
		sspId, _ := strconv.Atoi(v[mSmSingle.Field.F_ssp_id].(string))
		specNames := ""
		if v, ok := sspId2specMap[sspId]; ok {
			specNames = v
		}
		smSinglesMap[smId] = append(smSinglesMap[smId], cards.ItemSingle{
			SingleId:  singleId,
			Num:       num,
			Name:      singles[singleId],
			SspId:     sspId,
			SpecNames: specNames,
		})
	}

	var giveSinglesMap = make(map[int][]cards.ItemSingle)
	for _, v := range giveSingles {
		singleId, _ := strconv.Atoi(v[mSmGive.Field.F_single_id].(string))
		smId, _ := strconv.Atoi(v[mSmGive.Field.F_sm_id].(string))
		num, _ := strconv.Atoi(v[mSmGive.Field.F_num].(string))
		periodOfValidity, _ := strconv.Atoi(v[mSmGive.Field.F_period_of_validity].(string))

		giveSinglesMap[smId] = append(giveSinglesMap[smId], cards.ItemSingle{
			SingleId:         singleId,
			Num:              num,
			Name:             singles[singleId],
			PeriodOfValidity: periodOfValidity,
		})
	}
	//获取单项目的名称
	for ssId, v := range reply {
		v.ItemName = rebuildSms[v.ItemId].Name
		v.ImgId = rebuildSms[v.ItemId].ImgId
		v.Price = rebuildSms[v.ItemId].Price
		v.RealPrice = rebuildSms[v.ItemId].RealPrice
		v.ServicePeriod = rebuildSms[v.ItemId].ServicePeriod
		v.ValidCount = rebuildSms[v.ItemId].ValidCount
		v.CableShopIds = functions.ArrayUniqueInt(smShopIds[v.ItemId])
		v.Singles = smSinglesMap[v.ItemId]
		if rebuildSms[v.ItemId].HasGiveSignle == cards.HAS_GIVE_SINGLE_yes {
			v.Gives = giveSinglesMap[v.ItemId]
			if giveDesc, ok := giveDescMap[v.ItemId]; ok {
				v.GiveSingleDesc = giveDesc
			}
		}
		reply[ssId] = v
	}
	return
}

//获取综合卡
func (i *ItemLogic) getCards(ctx context.Context, ssIds []int) (reply map[cards.SsId]cards.ItemBase, err error) {
	mShopItem := new(models.ShopCardModel).Init()
	shopCards := mShopItem.GetByIDs(ssIds)
	if len(shopCards) == 0 {
		return
	}
	reply = make(map[cards.SsId]cards.ItemBase)
	//获取套餐id
	var cardIds []int
	for _, shopCard := range shopCards {
		cardId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_card_id].(string))
		ssId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_id].(string))
		status, _ := strconv.Atoi(shopCard[mShopItem.Field.F_status].(string))
		shopId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_shop_id].(string))
		shopSales, _ := strconv.Atoi(shopCard[mShopItem.Field.F_sales].(string))
		reply[cards.SsId(ssId)] = cards.ItemBase{
			SsId:      ssId,
			Status:    status,
			ItemId:    cardId,
			ShopId:    shopId,
			ShopSales: shopSales,
		}
		cardIds = append(cardIds, cardId)
	}
	//获取套餐数据
	mCard := new(models.CardModel).Init(mShopItem.Model.GetOrmer())
	items := mCard.GetByCardIDs(cardIds,
		mCard.Field.F_card_id,
		mCard.Field.F_name,
		mCard.Field.F_img_id,
		mCard.Field.F_price,
		mCard.Field.F_real_price,
		mCard.Field.F_has_give_signle,
		mCard.Field.F_service_period,
	)

	var rebuildCards = make(map[int]struct {
		CardId        int
		Name          string
		ImgId         int
		Price         float64
		RealPrice     float64
		HasGiveSignle int
		ServicePeriod int
	})
	for _, card := range items {
		cardId, _ := strconv.Atoi(card[mCard.Field.F_card_id].(string))
		imgId, _ := strconv.Atoi(card[mCard.Field.F_img_id].(string))
		price, _ := strconv.ParseFloat(card[mCard.Field.F_price].(string), 64)
		realPrice, _ := strconv.ParseFloat(card[mCard.Field.F_real_price].(string), 64)
		hasGiveSignle, _ := strconv.Atoi(card[mCard.Field.F_has_give_signle].(string))
		servicePeriod, _ := strconv.Atoi(card[mCard.Field.F_service_period].(string))

		rebuildCards[cardId] = struct {
			CardId        int
			Name          string
			ImgId         int
			Price         float64
			RealPrice     float64
			HasGiveSignle int
			ServicePeriod int
		}{CardId: cardId, Name: card[mCard.Field.F_name].(string), ImgId: imgId, Price: price, RealPrice: realPrice, HasGiveSignle: hasGiveSignle, ServicePeriod: servicePeriod}
	}

	//获取适用门店
	mCardShop := new(models.CardShopModel).Init(mShopItem.Model.GetOrmer())
	cardShops := mCardShop.GetByCardIDs(cardIds)
	var cardShopIds map[int][]int
	cardShopIds = make(map[int][]int)
	for _, cardShop := range cardShops {
		cardId, _ := strconv.Atoi(cardShop[mCardShop.Field.F_card_id].(string))
		shopId, _ := strconv.Atoi(cardShop[mCardShop.Field.F_shop_id].(string))
		if shopId == 0 {
			cardShopIds[cardId] = []int{0} //适用所有门店
		} else {
			cardShopIds[cardId] = append(cardShopIds[cardId], shopId)
		}
	}

	var singleIds = []int{}
	//获取包含的单项目
	mCardSingle := new(models.CardSingleModel).Init(mShopItem.Model.GetOrmer())
	cardSingles := mCardSingle.GetByCardIds(cardIds)
	singleIds = functions.ArrayValue2Array(mCardSingle.Field.F_single_id, cardSingles)
	//获取赠送的单项目数据
	mCardGive := new(models.CardGiveModel).Init(mShopItem.Model.GetOrmer())
	giveSingles := mCardGive.GetByCardIds(cardIds)
	singleIds = append(singleIds, functions.ArrayValue2Array(mCardGive.Field.F_single_id, giveSingles)...)
	singles := i.getSinglesBySingleids(singleIds)

	//获取赠品描述数据
	mGiveDesc := new(models.CardGiveDescModel).Init(mShopItem.Model.GetOrmer())
	giveSinglesDesc := mGiveDesc.GetByCardids(cardIds)
	giveDescMap := make(map[int][]cards.GiveSingleDesc)
	for _, dv := range giveSinglesDesc {
		smId, _ := strconv.Atoi(dv[mGiveDesc.Field.F_card_id].(string))
		descStr := dv[mGiveDesc.Field.F_desc].(string)
		desc := []cards.GiveSingleDesc{}
		json.Unmarshal([]byte(descStr), &desc)
		giveDescMap[smId] = desc
	}
	//获取单项目数据
	var cardSinglesMap = make(map[int][]cards.ItemSingle)
	for _, v := range cardSingles {
		singleId, _ := strconv.Atoi(v[mCardSingle.Field.F_single_id].(string))
		cardId, _ := strconv.Atoi(v[mCardSingle.Field.F_card_id].(string))
		cardSinglesMap[cardId] = append(cardSinglesMap[cardId], cards.ItemSingle{
			SingleId: singleId,
			Num:      0,
			Name:     singles[singleId],
		})
	}

	var giveSinglesMap = make(map[int][]cards.ItemSingle)
	for _, v := range giveSingles {
		singleId, _ := strconv.Atoi(v[mCardGive.Field.F_single_id].(string))
		cardId, _ := strconv.Atoi(v[mCardGive.Field.F_card_id].(string))
		num, _ := strconv.Atoi(v[mCardGive.Field.F_num].(string))
		periodOfValidity, _ := strconv.Atoi(v[mCardGive.Field.F_period_of_validity].(string))

		giveSinglesMap[cardId] = append(giveSinglesMap[cardId], cards.ItemSingle{
			SingleId:         singleId,
			Num:              num,
			Name:             singles[singleId],
			PeriodOfValidity: periodOfValidity,
		})
	}

	//获取包含的产品
	mCardGoods := new(models.CardGoodsModel).Init(mShopItem.Model.GetOrmer())
	cardGoods := mCardGoods.GetByCardIds(cardIds)
	var cardGoodsMap = make(map[int][]cards.ItemGoods)
	if len(cardGoods) > 0 {
		if cardGoods[0][mCardGoods.Field.F_product_id].(string) == "0" { //适用于全部商品
			for _, v := range cardGoods {
				cardId, _ := strconv.Atoi(v[mCardGoods.Field.F_card_id].(string))
				goodsId, _ := strconv.Atoi(v[mCardGoods.Field.F_product_id].(string))
				cardGoodsMap[cardId] = append(cardGoodsMap[cardId], cards.ItemGoods{
					GoodsId: goodsId,
				})
			}
		} else {
			goodsIds := functions.ArrayValue2Array(mCardGoods.Field.F_product_id, cardGoods)
			//获取产品名称
			rpcGoods := new(product.Product).Init()
			argGoods := product2.ArgsProductGetByIds{
				Ids: goodsIds,
			}
			replyGoods := []product2.ReplyProductGetByIds{}
			err = rpcGoods.GetProductByIds(ctx, &argGoods, &replyGoods)
			if err != nil {
				return
			}
			rebuildGoods := map[int]product2.ReplyProductGetByIds{}
			for _, goods := range replyGoods {
				rebuildGoods[goods.Id] = goods
			}

			for _, v := range cardGoods {
				cardId, _ := strconv.Atoi(v[mCardGoods.Field.F_card_id].(string))
				goodsId, _ := strconv.Atoi(v[mCardGoods.Field.F_product_id].(string))
				if rebuildGoods[goodsId].Id == 0 {
					continue
				}
				cardGoodsMap[cardId] = append(cardGoodsMap[cardId], cards.ItemGoods{
					GoodsId: goodsId,
					Name:    rebuildGoods[goodsId].Name,
				})
			}
		}
	}

	for ssId, v := range reply {
		v.ItemName = rebuildCards[v.ItemId].Name
		v.ImgId = rebuildCards[v.ItemId].ImgId
		v.Price = rebuildCards[v.ItemId].Price
		v.RealPrice = rebuildCards[v.ItemId].RealPrice
		v.ServicePeriod = rebuildCards[v.ItemId].ServicePeriod
		v.CableShopIds = functions.ArrayUniqueInt(cardShopIds[v.ItemId])
		v.Singles = cardSinglesMap[v.ItemId]
		if rebuildCards[v.ItemId].HasGiveSignle == cards.HAS_GIVE_SINGLE_yes {
			v.Gives = giveSinglesMap[v.ItemId]
			if giveDesc, ok := giveDescMap[v.ItemId]; ok {
				v.GiveSingleDesc = giveDesc
			}
		}
		if _, ok := cardGoodsMap[v.ItemId]; ok {
			v.Goods = cardGoodsMap[v.ItemId]
		}
		reply[ssId] = v
	}
	return
}

//获取限时卡
func (i *ItemLogic) getHcards(ssIds []int) (reply map[cards.SsId]cards.ItemBase, err error) {
	mShopItem := new(models.ShopHcardModel).Init()
	shopCards := mShopItem.GetByIDs(ssIds)
	if len(shopCards) == 0 {
		return
	}
	reply = make(map[cards.SsId]cards.ItemBase)
	//获取卡项id
	var cardIds []int
	for _, shopCard := range shopCards {
		cardId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_hcard_id].(string))
		ssId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_id].(string))
		status, _ := strconv.Atoi(shopCard[mShopItem.Field.F_status].(string))
		shopId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_shop_id].(string))
		shopSales, _ := strconv.Atoi(shopCard[mShopItem.Field.F_sales].(string))
		reply[cards.SsId(ssId)] = cards.ItemBase{
			SsId:      ssId,
			Status:    status,
			ItemId:    cardId,
			ShopId:    shopId,
			ShopSales: shopSales,
		}
		cardIds = append(cardIds, cardId)
	}
	//获取套餐数据
	mCard := new(models.HcardModel).Init(mShopItem.Model.GetOrmer())
	items := mCard.GetHcardByIDs(cardIds,
		mCard.Field.F_hcard_id,
		mCard.Field.F_name,
		mCard.Field.F_img_id,
		mCard.Field.F_price,
		mCard.Field.F_real_price,
		mCard.Field.F_has_give_signle,
		mCard.Field.F_service_period,
	)

	var rebuildCards = make(map[int]struct {
		CardId        int
		Name          string
		ImgId         int
		Price         float64
		RealPrice     float64
		HasGiveSignle int
		ServicePeriod int
	})
	for _, card := range items {
		cardId, _ := strconv.Atoi(card[mCard.Field.F_hcard_id].(string))
		imgId, _ := strconv.Atoi(card[mCard.Field.F_img_id].(string))
		price, _ := strconv.ParseFloat(card[mCard.Field.F_price].(string), 64)
		realPrice, _ := strconv.ParseFloat(card[mCard.Field.F_real_price].(string), 64)
		hasGiveSignle, _ := strconv.Atoi(card[mCard.Field.F_has_give_signle].(string))
		servicePeriod, _ := strconv.Atoi(card[mCard.Field.F_service_period].(string))

		rebuildCards[cardId] = struct {
			CardId        int
			Name          string
			ImgId         int
			Price         float64
			RealPrice     float64
			HasGiveSignle int
			ServicePeriod int
		}{CardId: cardId, Name: card[mCard.Field.F_name].(string), ImgId: imgId, Price: price, RealPrice: realPrice, HasGiveSignle: hasGiveSignle, ServicePeriod: servicePeriod}
	}

	//获取适用门店
	mCardShop := new(models.HcardShopModel).Init(mShopItem.Model.GetOrmer())
	cardShops := mCardShop.GetByHcardIDs(cardIds)
	var cardShopIds map[int][]int
	cardShopIds = make(map[int][]int)
	for _, cardShop := range cardShops {
		cardId, _ := strconv.Atoi(cardShop[mCardShop.Field.F_hcard_id].(string))
		shopId, _ := strconv.Atoi(cardShop[mCardShop.Field.F_shop_id].(string))
		if shopId == 0 {
			cardShopIds[cardId] = []int{0} //适用所有门店
		} else {
			cardShopIds[cardId] = append(cardShopIds[cardId], shopId)
		}
	}

	var singleIds = []int{}
	//获取包含的单项目
	mCardSingle := new(models.HcardSingleModel).Init(mShopItem.Model.GetOrmer())
	cardSingles := mCardSingle.GetByHcardIds(cardIds)
	singleIds = functions.ArrayValue2Array(mCardSingle.Field.F_single_id, cardSingles)
	//获取赠送的单项目数据
	mCardGive := new(models.HcardGiveModel).Init(mShopItem.Model.GetOrmer())
	giveSingles := mCardGive.GetByHcardIds(cardIds)
	singleIds = append(singleIds, functions.ArrayValue2Array(mCardGive.Field.F_single_id, giveSingles)...)
	singles := i.getSinglesBySingleids(singleIds)

	//获取赠品描述数据
	hcGiveDesc := new(models.HcardGiveDescModel).Init(mShopItem.Model.GetOrmer())
	giveSinglesDesc := hcGiveDesc.GetByHcardids(cardIds)
	giveDescMap := make(map[int][]cards.GiveSingleDesc)
	for _, dv := range giveSinglesDesc {
		smId, _ := strconv.Atoi(dv[hcGiveDesc.Field.F_hcard_id].(string))
		descStr := dv[hcGiveDesc.Field.F_desc].(string)
		desc := []cards.GiveSingleDesc{}
		json.Unmarshal([]byte(descStr), &desc)
		giveDescMap[smId] = desc
	}

	//获取单项目数据
	var cardSinglesMap = make(map[int][]cards.ItemSingle)
	for _, v := range cardSingles {
		singleId, _ := strconv.Atoi(v[mCardSingle.Field.F_single_id].(string))
		cardId, _ := strconv.Atoi(v[mCardSingle.Field.F_hcard_id].(string))
		cardSinglesMap[cardId] = append(cardSinglesMap[cardId], cards.ItemSingle{
			SingleId: singleId,
			Num:      0,
			Name:     singles[singleId],
		})
	}

	var giveSinglesMap = make(map[int][]cards.ItemSingle)
	for _, v := range giveSingles {
		singleId, _ := strconv.Atoi(v[mCardGive.Field.F_single_id].(string))
		cardId, _ := strconv.Atoi(v[mCardGive.Field.F_hcard_id].(string))
		num, _ := strconv.Atoi(v[mCardGive.Field.F_num].(string))
		periodOfValidity, _ := strconv.Atoi(v[mCardGive.Field.F_period_of_validity].(string))

		giveSinglesMap[cardId] = append(giveSinglesMap[cardId], cards.ItemSingle{
			SingleId:         singleId,
			Num:              num,
			Name:             singles[singleId],
			PeriodOfValidity: periodOfValidity,
		})
	}
	//获取单项目的名称
	for ssId, v := range reply {
		v.ItemName = rebuildCards[v.ItemId].Name
		v.ImgId = rebuildCards[v.ItemId].ImgId
		v.Price = rebuildCards[v.ItemId].Price
		v.RealPrice = rebuildCards[v.ItemId].RealPrice
		v.ServicePeriod = rebuildCards[v.ItemId].ServicePeriod
		v.CableShopIds = functions.ArrayUniqueInt(cardShopIds[v.ItemId])
		v.Singles = cardSinglesMap[v.ItemId]
		if rebuildCards[v.ItemId].HasGiveSignle == cards.HAS_GIVE_SINGLE_yes {
			v.Gives = giveSinglesMap[v.ItemId]
			if giveDesc, ok := giveDescMap[v.ItemId]; ok {
				v.GiveSingleDesc = giveDesc
			}
		}
		reply[ssId] = v
	}
	return
}

//获取限次卡
func (i *ItemLogic) getNcards(ssIds []int) (reply map[cards.SsId]cards.ItemBase, err error) {
	mShopItem := new(models.ShopNCardModel).Init()
	shopCards := mShopItem.GetByIDs(ssIds)
	if len(shopCards) == 0 {
		return
	}
	reply = make(map[cards.SsId]cards.ItemBase)
	//获取卡项id
	var cardIds []int
	for _, shopCard := range shopCards {
		cardId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_ncard_id].(string))
		ssId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_id].(string))
		status, _ := strconv.Atoi(shopCard[mShopItem.Field.F_status].(string))
		shopId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_shop_id].(string))
		shopSales, _ := strconv.Atoi(shopCard[mShopItem.Field.F_sales].(string))
		reply[cards.SsId(ssId)] = cards.ItemBase{
			SsId:      ssId,
			Status:    status,
			ItemId:    cardId,
			ShopId:    shopId,
			ShopSales: shopSales,
		}
		cardIds = append(cardIds, cardId)
	}
	//获取卡项数据
	mCard := new(models.NCardModel).Init(mShopItem.Model.GetOrmer())
	items := mCard.GetByNCardIDs(cardIds,
		mCard.Field.F_ncard_id,
		mCard.Field.F_name,
		mCard.Field.F_img_id,
		mCard.Field.F_price,
		mCard.Field.F_real_price,
		mCard.Field.F_has_give_signle,
		mCard.Field.F_service_period,
		mCard.Field.F_validcount,
	)

	var rebuildCards = make(map[int]struct {
		CardId        int
		Name          string
		ImgId         int
		Price         float64
		RealPrice     float64
		HasGiveSignle int
		ServicePeriod int
		ValidCount    int // 包含单项目的总次数
	})
	for _, card := range items {
		cardId, _ := strconv.Atoi(card[mCard.Field.F_ncard_id].(string))
		imgId, _ := strconv.Atoi(card[mCard.Field.F_img_id].(string))
		price, _ := strconv.ParseFloat(card[mCard.Field.F_price].(string), 64)
		realPrice, _ := strconv.ParseFloat(card[mCard.Field.F_real_price].(string), 64)
		hasGiveSignle, _ := strconv.Atoi(card[mCard.Field.F_has_give_signle].(string))
		servicePeriod, _ := strconv.Atoi(card[mCard.Field.F_service_period].(string))
		validCount, _ := strconv.Atoi(card[mCard.Field.F_validcount].(string))

		rebuildCards[cardId] = struct {
			CardId        int
			Name          string
			ImgId         int
			Price         float64
			RealPrice     float64
			HasGiveSignle int
			ServicePeriod int
			ValidCount    int // 包含单项目的总次数
		}{CardId: cardId, Name: card[mCard.Field.F_name].(string), ImgId: imgId, Price: price, RealPrice: realPrice, HasGiveSignle: hasGiveSignle,
			ServicePeriod: servicePeriod, ValidCount: validCount}
	}

	//获取适用门店
	mCardShop := new(models.NCardShopModel).Init(mShopItem.Model.GetOrmer())
	cardShops := mCardShop.GetByNCardIDs(cardIds)
	var cardShopIds map[int][]int
	cardShopIds = make(map[int][]int)
	for _, cardShop := range cardShops {
		cardId, _ := strconv.Atoi(cardShop[mCardShop.Field.F_ncard_id].(string))
		shopId, _ := strconv.Atoi(cardShop[mCardShop.Field.F_shop_id].(string))
		if shopId == 0 {
			cardShopIds[cardId] = []int{0} //适用所有门店
		} else {
			cardShopIds[cardId] = append(cardShopIds[cardId], shopId)
		}
	}

	var singleIds = []int{}
	//获取包含的单项目
	mCardSingle := new(models.NCardSingleModel).Init(mShopItem.Model.GetOrmer())
	cardSingles := mCardSingle.GetByNCardIds(cardIds)
	singleIds = functions.ArrayValue2Array(mCardSingle.Field.F_single_id, cardSingles)
	sspIds := functions.ArrayValue2Array(mCardSingle.Field.F_ssp_id, cardSingles)
	sspId2specMap := i.getSpecNames(sspIds)

	//获取赠送的单项目数据
	mCardGive := new(models.NCardGiveModel).Init(mShopItem.Model.GetOrmer())
	giveSingles := mCardGive.GetByNCardIds(cardIds)
	singleIds = append(singleIds, functions.ArrayValue2Array(mCardGive.Field.F_single_id, giveSingles)...)
	singles := i.getSinglesBySingleids(singleIds)

	//获取赠品描述数据
	ncGiveDesc := new(models.NcardGiveDescModel).Init(mShopItem.Model.GetOrmer())
	giveSinglesDesc := ncGiveDesc.GetByNcardids(cardIds)
	giveDescMap := make(map[int][]cards.GiveSingleDesc)
	for _, dv := range giveSinglesDesc {
		smId, _ := strconv.Atoi(dv[ncGiveDesc.Field.F_ncard_id].(string))
		descStr := dv[ncGiveDesc.Field.F_desc].(string)
		desc := []cards.GiveSingleDesc{}
		json.Unmarshal([]byte(descStr), &desc)
		giveDescMap[smId] = desc
	}

	//获取单项目数据
	var cardSinglesMap = make(map[int][]cards.ItemSingle)
	for _, v := range cardSingles {
		singleId, _ := strconv.Atoi(v[mCardSingle.Field.F_single_id].(string))
		cardId, _ := strconv.Atoi(v[mCardSingle.Field.F_ncard_id].(string))
		num, _ := strconv.Atoi(v[mCardSingle.Field.F_num].(string))
		sspId, _ := strconv.Atoi(v[mCardSingle.Field.F_ssp_id].(string))
		specNames := ""
		if v, ok := sspId2specMap[sspId]; ok {
			specNames = v
		}
		cardSinglesMap[cardId] = append(cardSinglesMap[cardId], cards.ItemSingle{
			SingleId:  singleId,
			Num:       num,
			Name:      singles[singleId],
			SspId:     sspId,
			SpecNames: specNames,
		})
	}

	var giveSinglesMap = make(map[int][]cards.ItemSingle)
	for _, v := range giveSingles {
		singleId, _ := strconv.Atoi(v[mCardGive.Field.F_single_id].(string))
		cardId, _ := strconv.Atoi(v[mCardGive.Field.F_ncard_id].(string))
		num, _ := strconv.Atoi(v[mCardGive.Field.F_num].(string))
		periodOfValidity, _ := strconv.Atoi(v[mCardGive.Field.F_period_of_validity].(string))

		giveSinglesMap[cardId] = append(giveSinglesMap[cardId], cards.ItemSingle{
			SingleId:         singleId,
			Num:              num,
			Name:             singles[singleId],
			PeriodOfValidity: periodOfValidity,
		})
	}
	//获取单项目的名称
	for ssId, v := range reply {
		v.ItemName = rebuildCards[v.ItemId].Name
		v.ImgId = rebuildCards[v.ItemId].ImgId
		v.Price = rebuildCards[v.ItemId].Price
		v.RealPrice = rebuildCards[v.ItemId].RealPrice
		v.ServicePeriod = rebuildCards[v.ItemId].ServicePeriod
		v.ValidCount = rebuildCards[v.ItemId].ValidCount
		v.CableShopIds = functions.ArrayUniqueInt(cardShopIds[v.ItemId])
		v.Singles = cardSinglesMap[v.ItemId]
		if rebuildCards[v.ItemId].HasGiveSignle == cards.HAS_GIVE_SINGLE_yes {
			v.Gives = giveSinglesMap[v.ItemId]
			if giveDesc, ok := giveDescMap[v.ItemId]; ok {
				v.GiveSingleDesc = giveDesc
			}
		}
		reply[ssId] = v
	}
	return
}

//getIcards  获取身份卡数据
func (i *ItemLogic) getIcards(ctx context.Context, ssIds []int) (reply map[cards.SsId]cards.ItemBase, err error) {
	shopIcardModel := new(models.ShopIcardModel).Init()
	shopIcardData := shopIcardModel.GetRcards(map[string]interface{}{
		"id": []interface{}{"IN", ssIds},
	}, 0, len(ssIds))
	if len(shopIcardData) == 0 {
		return
	}
	reply = make(map[cards.SsId]cards.ItemBase)

	var icardIDs []int
	for _, v := range shopIcardData {
		shopIcardID, _ := strconv.Atoi(v["id"].(string))
		status, _ := strconv.Atoi(v["status"].(string))
		shopID, _ := strconv.Atoi(v["shop_id"].(string))
		icardID, _ := strconv.Atoi(v["icard_id"].(string))
		sales, _ := strconv.Atoi(v["sales"].(string))

		reply[cards.SsId(shopIcardID)] = cards.ItemBase{
			ShopId:    shopID,      //门店id
			ItemId:    icardID,     //卡项id
			SsId:      shopIcardID, //卡项在门店的id
			ShopSales: sales,       // 门店销量
			Status:    status,      //卡项在门店的销售状态
		}
		icardIDs = append(icardIDs, icardID)
	}

	//身份卡
	icardModel := new(models.IcardModel).Init()
	icardData := icardModel.GetPaginationData(models.Condition{
		Where: map[string]interface{}{
			"icard_id": []interface{}{"IN", icardIDs},
		},
		Limit: len(icardIDs),
	})
	icardArray := make(map[int]interface{})
	for _, v := range icardData {
		icardID, _ := strconv.Atoi(v["icard_id"].(string))
		icardArray[icardID] = v
	}

	//身份卡赠送的单服务
	icardGiveModel := new(models.IcardGiveModel).Init()
	icardGiveData := icardGiveModel.GetAll(models.Condition{
		Where: map[string]interface{}{
			"icard_id": []interface{}{"IN", icardIDs},
		},
	})

	//获取单项目single_id=>single_name
	singleIds := functions.ArrayValue2Array("single_id", icardGiveData)
	singles := i.getSinglesBySingleids(singleIds)
	icardGivesMap := make(map[int][]cards.ItemSingle)
	for _, v := range icardGiveData {
		icardID, _ := strconv.Atoi(v["icard_id"].(string))
		singleID, _ := strconv.Atoi(v["single_id"].(string))
		num, _ := strconv.Atoi(v["num"].(string))

		icardGivesMap[icardID] = append(icardGivesMap[icardID], cards.ItemSingle{
			SingleId: singleID,
			Num:      num,
			Name:     singles[singleID],
		})
	}

	//获取赠品描述数据
	icGiveDesc := new(models.IcardGiveDescModel).Init()
	giveSinglesDesc := icGiveDesc.GetByIcardids(icardIDs)
	giveDescMap := make(map[int][]cards.GiveSingleDesc)
	for _, dv := range giveSinglesDesc {
		smId, _ := strconv.Atoi(dv[icGiveDesc.Field.F_icard_id].(string))
		descStr := dv[icGiveDesc.Field.F_desc].(string)
		desc := []cards.GiveSingleDesc{}
		json.Unmarshal([]byte(descStr), &desc)
		giveDescMap[smId] = desc
	}

	//身份卡包含的商品
	icardGoodsModel := new(models.IcardGoodsModel).Init()
	icardGoodsData := icardGoodsModel.GetAll(models.Condition{
		Where: map[string]interface{}{
			"icard_id": []interface{}{"IN", icardIDs},
		},
	})

	goodsIDs := functions.ArrayValue2Array("goods_id", icardGoodsData)
	goodsData, _ := getIncProducts(ctx, goodsIDs)
	//TODO  goods select
	icardGoodsMap := make(map[int][]cards.ItemGoods)
	for _, v := range icardGoodsData {
		icardID, _ := strconv.Atoi(v["icard_id"].(string))
		goodsID, _ := strconv.Atoi(v["goods_id"].(string))
		price, _ := strconv.ParseFloat(goodsData[goodsID].SpecPrice, 64)
		discount, _ := strconv.ParseFloat(v["discount"].(string), 64)

		icardGoodsMap[icardID] = append(icardGoodsMap[icardID], cards.ItemGoods{
			GoodsId:  goodsID,
			Name:     goodsData[goodsID].Name,
			Price:    price,
			Discount: discount,
		})
	}

	//身份卡包含的单项目
	icardSingleModel := new(models.IcardSingleModel).Init()
	icardSingleData := icardSingleModel.GetAll(models.Condition{
		Where: map[string]interface{}{
			"icard_id": []interface{}{"IN", icardIDs},
		},
	})

	singleIds1 := functions.ArrayValue2Array("single_id", icardSingleData)
	singles1 := i.getSinglesBySingleids(singleIds1)
	icardSingleMap := make(map[int][]cards.ItemSingle)
	for _, v := range icardSingleData {
		icardID, _ := strconv.Atoi(v["icard_id"].(string))
		singleID, _ := strconv.Atoi(v["single_id"].(string))
		discount, _ := strconv.ParseFloat(v["discount"].(string), 64)

		icardSingleMap[icardID] = append(icardSingleMap[icardID], cards.ItemSingle{
			SingleId: singleID,
			Discount: discount,
			Name:     singles1[singleID],
		})
	}

	//身份卡适用门店
	icardShopModel := new(models.IcardShopModel).Init()
	icardShopData := icardShopModel.GetAll(models.Condition{
		Where: map[string]interface{}{
			"icard_id": []interface{}{"IN", icardIDs},
		},
	})
	icardShopIds := make(map[int][]int)
	for _, v := range icardShopData {
		icardID, _ := strconv.Atoi(v["icard_id"].(string))
		shopID, _ := strconv.Atoi(v["shop_id"].(string))
		icardShopIds[icardID] = append(icardShopIds[icardID], shopID)
	}

	for shopIcardID, v := range reply {
		// ItemName      string       //卡项名称
		// ImgId         int          //封面图片id
		// Price         float64      //面值
		// RealPrice     float64      //真实售价
		// ServicePeriod int          //保险周期 月
		// ValidCount    int          // 包含单项目的总次数 --- N
		// ShopRealPrice float64      //门店真实售价 --- N
		// CableShopIds  []int        //适用门店ids
		// Gives         []ItemSingle //赠送的服务
		// Singles       []ItemSingle //包含的服务
		// Goods         []ItemGoods  //包含的产品
		icardMap := icardArray[v.ItemId].(map[string]interface{})
		v.ItemName = icardMap["name"].(string)
		imgID, _ := strconv.Atoi(icardMap["img_id"].(string))
		v.ImgId = imgID
		price, _ := strconv.ParseFloat(icardMap["price"].(string), 64)
		v.Price = price
		RealPrice, _ := strconv.ParseFloat(icardMap["real_price"].(string), 64)
		v.RealPrice = RealPrice
		ServicePeriod, _ := strconv.Atoi(icardMap["service_period"].(string))
		v.ServicePeriod = ServicePeriod
		v.CableShopIds = icardShopIds[v.ItemId]
		v.Gives = icardGivesMap[v.ItemId]
		v.Singles = icardSingleMap[v.ItemId]
		v.Goods = icardGoodsMap[v.ItemId]
		if giveDesc, ok := giveDescMap[v.ItemId]; ok {
			v.GiveSingleDesc = giveDesc
		}
		reply[shopIcardID] = v
	}

	return
}

//获取限时限次卡
func (i *ItemLogic) getHncards(ssIds []int) (reply map[cards.SsId]cards.ItemBase, err error) {
	mShopItem := new(models.ShopHNCardModel).Init()
	shopCards := mShopItem.GetByIDs(ssIds)
	if len(shopCards) == 0 {
		return
	}
	reply = make(map[cards.SsId]cards.ItemBase)
	//获取卡项id
	var cardIds []int
	for _, shopCard := range shopCards {
		cardId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_hncard_id].(string))
		ssId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_id].(string))
		status, _ := strconv.Atoi(shopCard[mShopItem.Field.F_status].(string))
		shopId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_shop_id].(string))
		shopSales, _ := strconv.Atoi(shopCard[mShopItem.Field.F_sales].(string))
		reply[cards.SsId(ssId)] = cards.ItemBase{
			SsId:      ssId,
			Status:    status,
			ItemId:    cardId,
			ShopId:    shopId,
			ShopSales: shopSales,
		}
		cardIds = append(cardIds, cardId)
	}
	//获取卡项数据
	mCard := new(models.HNCardModel).Init(mShopItem.Model.GetOrmer())
	items := mCard.GetByHNCardIDs(cardIds,
		mCard.Field.F_hncard_id,
		mCard.Field.F_name,
		mCard.Field.F_img_id,
		mCard.Field.F_price,
		mCard.Field.F_real_price,
		mCard.Field.F_has_give_signle,
		mCard.Field.F_service_period,
		mCard.Field.F_validcount,
	)

	var rebuildCards = make(map[int]struct {
		CardId        int
		Name          string
		ImgId         int
		Price         float64
		RealPrice     float64
		HasGiveSignle int
		ServicePeriod int
		ValidCount    int // 包含单项目的总次数
	})
	for _, card := range items {
		cardId, _ := strconv.Atoi(card[mCard.Field.F_hncard_id].(string))
		imgId, _ := strconv.Atoi(card[mCard.Field.F_img_id].(string))
		price, _ := strconv.ParseFloat(card[mCard.Field.F_price].(string), 64)
		realPrice, _ := strconv.ParseFloat(card[mCard.Field.F_real_price].(string), 64)
		hasGiveSignle, _ := strconv.Atoi(card[mCard.Field.F_has_give_signle].(string))
		servicePeriod, _ := strconv.Atoi(card[mCard.Field.F_service_period].(string))
		validCount, _ := strconv.Atoi(card[mCard.Field.F_validcount].(string))
		rebuildCards[cardId] = struct {
			CardId        int
			Name          string
			ImgId         int
			Price         float64
			RealPrice     float64
			HasGiveSignle int
			ServicePeriod int
			ValidCount    int // 包含单项目的总次数
		}{CardId: cardId, Name: card[mCard.Field.F_name].(string), ImgId: imgId, Price: price, RealPrice: realPrice, HasGiveSignle: hasGiveSignle,
			ServicePeriod: servicePeriod, ValidCount: validCount}
	}

	//获取适用门店
	mCardShop := new(models.HNCardShopModel).Init(mShopItem.Model.GetOrmer())
	cardShops := mCardShop.GetByHNCardIDs(cardIds)
	var cardShopIds map[int][]int
	cardShopIds = make(map[int][]int)
	for _, cardShop := range cardShops {
		cardId, _ := strconv.Atoi(cardShop[mCardShop.Field.F_hncard_id].(string))
		shopId, _ := strconv.Atoi(cardShop[mCardShop.Field.F_shop_id].(string))
		if shopId == 0 {
			cardShopIds[cardId] = []int{0} //适用所有门店
		} else {
			cardShopIds[cardId] = append(cardShopIds[cardId], shopId)
		}
	}

	var singleIds = []int{}
	//获取包含的单项目
	mCardSingle := new(models.HNCardSingleModel).Init(mShopItem.Model.GetOrmer())
	cardSingles := mCardSingle.GetByHNCardIds(cardIds)
	singleIds = functions.ArrayValue2Array(mCardSingle.Field.F_single_id, cardSingles)
	sspIds := functions.ArrayValue2Array(mCardSingle.Field.F_ssp_id, cardSingles)
	sspId2specMap := i.getSpecNames(sspIds)
	//获取赠送的单项目数据
	mCardGive := new(models.HNCardGiveModel).Init(mShopItem.Model.GetOrmer())
	giveSingles := mCardGive.GetByHNCardIds(cardIds)
	singleIds = append(singleIds, functions.ArrayValue2Array(mCardGive.Field.F_single_id, giveSingles)...)
	singles := i.getSinglesBySingleids(singleIds)

	//获取赠品描述数据
	hncGiveDesc := new(models.HncardGiveDescModel).Init(mShopItem.Model.GetOrmer())
	giveSinglesDesc := hncGiveDesc.GetByHncardids(cardIds)
	giveDescMap := make(map[int][]cards.GiveSingleDesc)
	for _, dv := range giveSinglesDesc {
		smId, _ := strconv.Atoi(dv[hncGiveDesc.Field.F_hncard_id].(string))
		descStr := dv[hncGiveDesc.Field.F_desc].(string)
		desc := []cards.GiveSingleDesc{}
		json.Unmarshal([]byte(descStr), &desc)
		giveDescMap[smId] = desc
	}
	//获取单项目数据
	var cardSinglesMap = make(map[int][]cards.ItemSingle)
	for _, v := range cardSingles {
		singleId, _ := strconv.Atoi(v[mCardSingle.Field.F_single_id].(string))
		cardId, _ := strconv.Atoi(v[mCardSingle.Field.F_hncard_id].(string))
		num, _ := strconv.Atoi(v[mCardSingle.Field.F_num].(string))
		sspId, _ := strconv.Atoi(v[mCardSingle.Field.F_ssp_id].(string))
		specNames := ""
		if v, ok := sspId2specMap[sspId]; ok {
			specNames = v
		}
		cardSinglesMap[cardId] = append(cardSinglesMap[cardId], cards.ItemSingle{
			SingleId:  singleId,
			Num:       num,
			Name:      singles[singleId],
			SspId:     sspId,
			SpecNames: specNames,
		})
	}

	var giveSinglesMap = make(map[int][]cards.ItemSingle)
	for _, v := range giveSingles {
		singleId, _ := strconv.Atoi(v[mCardGive.Field.F_single_id].(string))
		cardId, _ := strconv.Atoi(v[mCardGive.Field.F_hncard_id].(string))
		num, _ := strconv.Atoi(v[mCardGive.Field.F_num].(string))
		periodOfValidity, _ := strconv.Atoi(v[mCardGive.Field.F_period_of_validity].(string))

		giveSinglesMap[cardId] = append(giveSinglesMap[cardId], cards.ItemSingle{
			SingleId:         singleId,
			Num:              num,
			Name:             singles[singleId],
			PeriodOfValidity: periodOfValidity,
		})
	}
	//获取单项目的名称
	for ssId, v := range reply {
		v.ItemName = rebuildCards[v.ItemId].Name
		v.ImgId = rebuildCards[v.ItemId].ImgId
		v.Price = rebuildCards[v.ItemId].Price
		v.RealPrice = rebuildCards[v.ItemId].RealPrice
		v.ServicePeriod = rebuildCards[v.ItemId].ServicePeriod
		v.ValidCount = rebuildCards[v.ItemId].ValidCount
		v.CableShopIds = functions.ArrayUniqueInt(cardShopIds[v.ItemId])
		v.Singles = cardSinglesMap[v.ItemId]
		if rebuildCards[v.ItemId].HasGiveSignle == cards.HAS_GIVE_SINGLE_yes {
			v.Gives = giveSinglesMap[v.ItemId]
			if giveDesc, ok := giveDescMap[v.ItemId]; ok {
				v.GiveSingleDesc = giveDesc
			}
		}
		reply[ssId] = v
	}
	return
}

func (i *ItemLogic) getSinglesBySingleids(singleIds []int) (singles map[int]string) {
	mSingle := new(models.SingleModel).Init()
	r := mSingle.GetBySingleids(singleIds, []string{
		mSingle.Field.F_single_id,
		mSingle.Field.F_name,
	})
	singles = make(map[int]string)
	for _, v := range r {
		singleId, _ := strconv.Atoi(v[mSingle.Field.F_single_id].(string))
		singles[singleId] = v[mSingle.Field.F_name].(string)
	}

	return
}

//获取充值卡
func (i *ItemLogic) getRards(ctx context.Context, ssIds []int) (reply map[cards.SsId]cards.ItemBase, err error) {
	mShopItem := new(models.ShopRcardModel).Init()
	shopCards := mShopItem.GetByIDs(ssIds)
	if len(shopCards) == 0 {
		return
	}
	reply = make(map[cards.SsId]cards.ItemBase)

	//获取充值卡id
	var cardIds []int
	for _, shopCard := range shopCards {
		cardId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_rcard_id].(string))
		ssId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_id].(string))
		status, _ := strconv.Atoi(shopCard[mShopItem.Field.F_status].(string))
		shopId, _ := strconv.Atoi(shopCard[mShopItem.Field.F_shop_id].(string))
		shopSales, _ := strconv.Atoi(shopCard[mShopItem.Field.F_sales].(string))
		reply[cards.SsId(ssId)] = cards.ItemBase{
			SsId:      ssId,
			Status:    status,
			ItemId:    cardId,
			ShopId:    shopId,
			ShopSales: shopSales,
		}
		cardIds = append(cardIds, cardId)
	}
	//获取充值卡数据
	mCard := new(models.RcardModel).Init(mShopItem.Model.GetOrmer())
	items := mCard.GetByRcardIDs(cardIds,
		mCard.Field.F_rcard_id,
		mCard.Field.F_name,
		mCard.Field.F_img_id,
		mCard.Field.F_price,
		mCard.Field.F_real_price,
		mCard.Field.F_has_give_signle,
		mCard.Field.F_service_period,
		mCard.Field.F_discount_type,
		mCard.Field.F_discount,
	)

	var rebuildCards = make(map[int]struct {
		CardId        int
		Name          string
		ImgId         int
		Price         float64
		RealPrice     float64
		HasGiveSignle int
		ServicePeriod int
		DiscountType  int
		Discount      float64
	})
	for _, card := range items {
		cardId, _ := strconv.Atoi(card[mCard.Field.F_rcard_id].(string))
		imgId, _ := strconv.Atoi(card[mCard.Field.F_img_id].(string))
		price, _ := strconv.ParseFloat(card[mCard.Field.F_price].(string), 64)
		realPrice, _ := strconv.ParseFloat(card[mCard.Field.F_real_price].(string), 64)
		hasGiveSignle, _ := strconv.Atoi(card[mCard.Field.F_has_give_signle].(string))
		servicePeriod, _ := strconv.Atoi(card[mCard.Field.F_service_period].(string))
		discountType, _ := strconv.Atoi(card[mCard.Field.F_discount_type].(string))
		discount, _ := strconv.ParseFloat(card[mCard.Field.F_discount].(string), 64)
		rebuildCards[cardId] = struct {
			CardId        int
			Name          string
			ImgId         int
			Price         float64
			RealPrice     float64
			HasGiveSignle int
			ServicePeriod int
			DiscountType  int
			Discount      float64
		}{CardId: cardId, Name: card[mCard.Field.F_name].(string), ImgId: imgId, Price: price, RealPrice: realPrice,
			HasGiveSignle: hasGiveSignle, ServicePeriod: servicePeriod, DiscountType: discountType, Discount: discount,
		}
	}

	//获取适用门店
	mCardShop := new(models.RcardShopModel).Init(mShopItem.Model.GetOrmer())
	cardShops := mCardShop.GetByRcardIDs(cardIds)
	var cardShopIds map[int][]int
	cardShopIds = make(map[int][]int)
	for _, cardShop := range cardShops {
		cardId, _ := strconv.Atoi(cardShop[mCardShop.Field.F_rcard_id].(string))
		shopId, _ := strconv.Atoi(cardShop[mCardShop.Field.F_shop_id].(string))
		if shopId == 0 {
			cardShopIds[cardId] = []int{0} //适用所有门店
		} else {
			cardShopIds[cardId] = append(cardShopIds[cardId], shopId)
		}
	}

	var singleIds = []int{}
	//获取包含的单项目
	mCardSingle := new(models.RcardSingleModel).Init(mShopItem.Model.GetOrmer())
	cardSingles := mCardSingle.GetByRcardIds(cardIds)
	singleIds = functions.ArrayValue2Array(mCardSingle.Field.F_single_id, cardSingles)
	//获取赠送的单项目数据
	mCardGive := new(models.RcardGiveModel).Init(mShopItem.Model.GetOrmer())
	giveSingles := mCardGive.GetByRcardIds(cardIds)
	singleIds = append(singleIds, functions.ArrayValue2Array(mCardGive.Field.F_single_id, giveSingles)...)
	singles := i.getSinglesBySingleids(singleIds)

	//获取赠品描述数据
	rcGiveDesc := new(models.RcardGiveDescModel).Init()
	giveSinglesDesc := rcGiveDesc.GetByRcardids(cardIds)
	giveDescMap := make(map[int][]cards.GiveSingleDesc)
	for _, dv := range giveSinglesDesc {
		smId, _ := strconv.Atoi(dv[rcGiveDesc.Field.F_rcard_id].(string))
		descStr := dv[rcGiveDesc.Field.F_desc].(string)
		desc := []cards.GiveSingleDesc{}
		json.Unmarshal([]byte(descStr), &desc)
		giveDescMap[smId] = desc
	}
	//获取单项目数据
	var cardSinglesMap = make(map[int][]cards.ItemSingle)
	for _, v := range cardSingles {
		singleId, _ := strconv.Atoi(v[mCardSingle.Field.F_single_id].(string))
		cardId, _ := strconv.Atoi(v[mCardSingle.Field.F_rcard_id].(string))
		discount, _ := strconv.ParseFloat(v[mCardSingle.Field.F_discount].(string), 64)
		cardSinglesMap[cardId] = append(cardSinglesMap[cardId], cards.ItemSingle{
			SingleId: singleId,
			Num:      0,
			Name:     singles[singleId],
			Discount: discount,
		})
	}

	var giveSinglesMap = make(map[int][]cards.ItemSingle)
	for _, v := range giveSingles {
		singleId, _ := strconv.Atoi(v[mCardGive.Field.F_single_id].(string))
		cardId, _ := strconv.Atoi(v[mCardGive.Field.F_rcard_id].(string))
		num, _ := strconv.Atoi(v[mCardGive.Field.F_num].(string))
		periodOfValidity, _ := strconv.Atoi(v[mCardGive.Field.F_period_of_validity].(string))

		giveSinglesMap[cardId] = append(giveSinglesMap[cardId], cards.ItemSingle{
			SingleId:         singleId,
			Num:              num,
			Name:             singles[singleId],
			PeriodOfValidity: periodOfValidity,
		})
	}

	//获取包含的产品
	mCardGoods := new(models.RcardGoodsModel).Init(mShopItem.Model.GetOrmer())
	cardGoods := mCardGoods.GetByRcardIds(cardIds)
	var cardGoodsMap = make(map[int][]cards.ItemGoods)
	if len(cardGoods) > 0 {
		if cardGoods[0][mCardGoods.Field.F_product_id].(string) == "0" { //适用于全部商品
			for _, v := range cardGoods {
				cardId, _ := strconv.Atoi(v[mCardGoods.Field.F_rcard_id].(string))
				goodsId, _ := strconv.Atoi(v[mCardGoods.Field.F_product_id].(string))
				//discount,_:=strconv.ParseFloat(v[mCardSingle.Field.F_discount].(string),64)
				cardGoodsMap[cardId] = append(cardGoodsMap[cardId], cards.ItemGoods{
					GoodsId:  goodsId,
					Discount: 10,
				})
			}
		} else {
			goodsIds := functions.ArrayValue2Array(mCardGoods.Field.F_product_id, cardGoods)
			//获取产品名称
			rpcGoods := new(product.Product).Init()
			argGoods := product2.ArgsProductGetByIds{
				Ids: goodsIds,
			}
			replyGoods := []product2.ReplyProductGetByIds{}
			err = rpcGoods.GetProductByIds(ctx, &argGoods, &replyGoods)
			if err != nil {
				return
			}
			rebuildGoods := map[int]product2.ReplyProductGetByIds{}
			for _, goods := range replyGoods {
				rebuildGoods[goods.Id] = goods
			}

			for _, v := range cardGoods {
				cardId, _ := strconv.Atoi(v[mCardGoods.Field.F_rcard_id].(string))
				goodsId, _ := strconv.Atoi(v[mCardGoods.Field.F_product_id].(string))
				if rebuildGoods[goodsId].Id == 0 {
					continue
				}
				//discount,_:=strconv.ParseFloat(v[mCardSingle.Field.F_discount].(string),64)
				cardGoodsMap[cardId] = append(cardGoodsMap[cardId], cards.ItemGoods{
					GoodsId:  goodsId,
					Name:     rebuildGoods[goodsId].Name,
					Discount: 10,
				})
			}
		}
	}

	for ssId, v := range reply {
		v.ItemName = rebuildCards[v.ItemId].Name
		v.ImgId = rebuildCards[v.ItemId].ImgId
		v.Price = rebuildCards[v.ItemId].Price
		v.RealPrice = rebuildCards[v.ItemId].RealPrice
		v.ServicePeriod = rebuildCards[v.ItemId].ServicePeriod
		v.CableShopIds = functions.ArrayUniqueInt(cardShopIds[v.ItemId])
		v.Singles = cardSinglesMap[v.ItemId]
		v.Discount = rebuildCards[v.ItemId].Discount
		v.DiscountType = rebuildCards[v.ItemId].DiscountType
		if rebuildCards[v.ItemId].HasGiveSignle == cards.HAS_GIVE_SINGLE_yes {
			v.Gives = giveSinglesMap[v.ItemId]
			if giveDesc, ok := giveDescMap[v.ItemId]; ok {
				v.GiveSingleDesc = giveDesc
			}
		}
		if _, ok := cardGoodsMap[v.ItemId]; ok {
			v.Goods = cardGoodsMap[v.ItemId]
		}
		reply[ssId] = v
	}
	return
}

//根据条件查询九百岁
func (i *ItemLogic) GetInfos(ctx context.Context, args *cards.ArgsAppInfos) (*cards.ReplyAppInfo, error) {
	if args.Id <= 0 {
		args.Id = cards.ITEM_TYPE_single
	}
	rpcEs := new(cards4.ShopCards).Init()
	defer rpcEs.Close()

	var replyEs map[string]interface{}
	err := rpcEs.SearchItem(ctx, args, &replyEs)
	if err != nil {
		return nil, err
	}

	if len(replyEs) == 0 {
		return &cards.ReplyAppInfo{Lists: []cards.AppInfo{}}, nil
	}

	//行业id集合
	var bindIds []int
	var districtIds []int
	var shopIds []int
	var shopItemIds []int
	if replyEs["result"] == nil {
		return &cards.ReplyAppInfo{Lists: []cards.AppInfo{}}, nil
	}
	result := replyEs["result"].([]interface{})
	for i := range result {
		source := result[i].(map[string]interface{})
		if source["Ctime"] != "" && source["Ctime"] != 0 &&
			source["Ctime"] != nil {
			source["Ctime"] = functions.TimeToStr(int64(source["Ctime"].(float64)))
		}
		bindIds = append(bindIds, int(source["MainBindId"].(float64)))
		if source["BindId"] != nil {
			for _, v := range source["BindId"].([]interface{}) {
				bindIds = append(bindIds, int(v.(float64)))
			}
		}
		if /*args.Flag == 1 &&*/ source["DistrictId"] != nil && len(source["DistrictId"].([]interface{})) > 0 {
			for _, v := range source["DistrictId"].([]interface{}) {
				districtIds = append(districtIds, int(v.(float64)))
			}
		}
		shopIds = append(shopIds, int(source["ShopId"].(float64)))
		if source["Location"] != nil && len(source["Location"].([]interface{})) == 2 {
			lon := source["Location"].([]interface{})[0].(float64)
			lat := source["Location"].([]interface{})[1].(float64)
			source["Lon"] = lon
			source["Lat"] = lat
			if args.Lon > 0 && args.Lat > 0 {
				source["Distance"] = tools.GetDistance(lon, lat, args.Lon, args.Lat)
			} else {
				source["Distance"] = 0
			}
		}
		if args.Flag == 1 {
			for _, item := range source["Sub"].([]interface{}) {
				shopItemIds = append(shopItemIds, int(item.(map[string]interface{})["ItemId"].(float64)))
			}
		}
	}
	//扩展数据
	type extendData struct {
		ServicePeriod   int     //保险周期，限时月份
		ServiceTime     int     //服务时长
		TotalNum        int     //总次数
		Price           float64 //标价
		ServiceDiscount float64 //服务折扣
		ProductDiscount float64 //商品折扣
	}
	extendDataMap := map[int]extendData{}
	if len(shopItemIds) > 0 {
		var servicePeriods []map[string]interface{}
		switch args.Id {
		case cards.ITEM_TYPE_single:
			model := new(models.SingleModel).Init()
			servicePeriods = model.GetBySingleids(shopItemIds, []string{model.Field.F_single_id, model.Field.F_service_time, model.Field.F_price})
			for _, period := range servicePeriods {
				id, _ := strconv.Atoi(period[model.Field.F_single_id].(string))
				serviceTime, _ := strconv.Atoi(period[model.Field.F_service_time].(string))
				price, _ := strconv.ParseFloat(period[model.Field.F_price].(string), 64)
				extendDataMap[id] = extendData{
					ServicePeriod: 0,
					ServiceTime:   serviceTime,
					TotalNum:      0,
					Price:         price,
				}
			}
		case cards.ITEM_TYPE_sm:
			smModel := new(models.SmModel).Init()
			servicePeriods = smModel.GetBySmids(shopItemIds, []string{smModel.Field.F_service_period, smModel.Field.F_sm_id, smModel.Field.F_price, smModel.Field.F_validcount})
			for _, period := range servicePeriods {
				id, _ := strconv.Atoi(period[smModel.Field.F_sm_id].(string))
				servicePeriod, _ := strconv.Atoi(period[smModel.Field.F_service_period].(string))
				totalNum, _ := strconv.Atoi(period[smModel.Field.F_validcount].(string))
				price, _ := strconv.ParseFloat(period[smModel.Field.F_price].(string), 64)
				extendDataMap[id] = extendData{
					ServicePeriod: servicePeriod,
					ServiceTime:   0,
					TotalNum:      totalNum,
					Price:         price,
				}
			}
		case cards.ITEM_TYPE_card:
			cardModel := new(models.CardModel).Init()
			servicePeriods = cardModel.GetByCardIDs(shopItemIds, cardModel.Field.F_card_id, cardModel.Field.F_service_period, cardModel.Field.F_price)
			for _, period := range servicePeriods {
				id, _ := strconv.Atoi(period[cardModel.Field.F_card_id].(string))
				servicePeriod, _ := strconv.Atoi(period[cardModel.Field.F_service_period].(string))
				price, _ := strconv.ParseFloat(period[cardModel.Field.F_price].(string), 64)
				extendDataMap[id] = extendData{
					ServicePeriod: servicePeriod,
					ServiceTime:   0,
					TotalNum:      0,
					Price:         price,
				}
			}
		case cards.ITEM_TYPE_hcard:
			hcardModel := new(models.HcardModel).Init()
			servicePeriods = hcardModel.GetHcardByIDs(shopItemIds, hcardModel.Field.F_hcard_id, hcardModel.Field.F_service_period, hcardModel.Field.F_price)
			for _, period := range servicePeriods {
				id, _ := strconv.Atoi(period[hcardModel.Field.F_hcard_id].(string))
				servicePeriod, _ := strconv.Atoi(period[hcardModel.Field.F_service_period].(string))
				price, _ := strconv.ParseFloat(period[hcardModel.Field.F_price].(string), 64)
				extendDataMap[id] = extendData{
					ServicePeriod: servicePeriod,
					ServiceTime:   0,
					TotalNum:      0,
					Price:         price,
				}
			}
		case cards.ITEM_TYPE_ncard:
			ncardModel := new(models.NCardModel).Init()
			servicePeriods = ncardModel.GetByNCardIDs(shopItemIds, ncardModel.Field.F_ncard_id, ncardModel.Field.F_service_period, ncardModel.Field.F_price, ncardModel.Field.F_validcount)
			for _, period := range servicePeriods {
				id, _ := strconv.Atoi(period[ncardModel.Field.F_ncard_id].(string))
				servicePeriod, _ := strconv.Atoi(period[ncardModel.Field.F_service_period].(string))
				price, _ := strconv.ParseFloat(period[ncardModel.Field.F_price].(string), 64)
				totalNum, _ := strconv.Atoi(period[ncardModel.Field.F_validcount].(string))
				extendDataMap[id] = extendData{
					ServicePeriod: servicePeriod,
					ServiceTime:   0,
					TotalNum:      totalNum,
					Price:         price,
				}
			}
		case cards.ITEM_TYPE_hncard:
			hncardModel := new(models.HNCardModel).Init()
			servicePeriods = hncardModel.GetByHNCardIDs(shopItemIds, hncardModel.Field.F_hncard_id, hncardModel.Field.F_service_period, hncardModel.Field.F_validcount, hncardModel.Field.F_price)
			for _, period := range servicePeriods {
				id, _ := strconv.Atoi(period[hncardModel.Field.F_hncard_id].(string))
				servicePeriod, _ := strconv.Atoi(period[hncardModel.Field.F_service_period].(string))
				price, _ := strconv.ParseFloat(period[hncardModel.Field.F_price].(string), 64)
				totalNum, _ := strconv.Atoi(period[hncardModel.Field.F_validcount].(string))
				extendDataMap[id] = extendData{
					ServicePeriod: servicePeriod,
					ServiceTime:   0,
					TotalNum:      totalNum,
					Price:         price,
				}
			}
		case cards.ITEM_TYPE_rcard:
			rcardModel := new(models.RcardModel).Init()
			servicePeriods = rcardModel.GetRcardsByRcardIds(shopItemIds, rcardModel.Field.F_rcard_id, rcardModel.Field.F_service_period, rcardModel.Field.F_price)
			for _, period := range servicePeriods {
				id, _ := strconv.Atoi(period[rcardModel.Field.F_rcard_id].(string))
				servicePeriod, _ := strconv.Atoi(period[rcardModel.Field.F_service_period].(string))
				price, _ := strconv.ParseFloat(period[rcardModel.Field.F_price].(string), 64)
				extendDataMap[id] = extendData{
					ServicePeriod: servicePeriod,
					ServiceTime:   0,
					TotalNum:      0,
					Price:         price,
				}
			}
		case cards.ITEM_TYPE_icard:
			icardModel := new(models.IcardModel).Init()
			servicePeriods = icardModel.GetRcardsByIcardIds(shopItemIds, icardModel.Field.F_icard_id, icardModel.Field.F_service_period, icardModel.Field.F_price)
			icardIds := functions.ArrayValue2Array(icardModel.Field.F_icard_id, servicePeriods)
			isModel := new(models.IcardSingleModel).Init()
			isMaps := isModel.GetByIcardIds(icardIds, isModel.Field.F_icard_id, isModel.Field.F_discount)
			isMap := functions.ArrayRebuild(isModel.Field.F_icard_id, isMaps)
			igModel := new(models.IcardGoodsModel).Init()
			igMaps := igModel.GetByIcardIds(icardIds, igModel.Field.F_icard_id, igModel.Field.F_discount)
			igMap := functions.ArrayRebuild(igModel.Field.F_icard_id, igMaps)
			for _, period := range servicePeriods {
				id, _ := strconv.Atoi(period[icardModel.Field.F_icard_id].(string))
				servicePeriod, _ := strconv.Atoi(period[icardModel.Field.F_service_period].(string))
				price, _ := strconv.ParseFloat(period[icardModel.Field.F_price].(string), 64)
				isDis := 0.00
				if v, ok := isMap[strconv.Itoa(id)]; ok {
					isDis, _ = strconv.ParseFloat(v.(map[string]interface{})[isModel.Field.F_discount].(string), 64)
				}
				igDis := 0.00
				if v, ok := igMap[strconv.Itoa(id)]; ok {
					igDis, _ = strconv.ParseFloat(v.(map[string]interface{})[igModel.Field.F_discount].(string), 64)
				}
				extendDataMap[id] = extendData{
					ServicePeriod:   servicePeriod,
					ServiceTime:     0,
					TotalNum:        0,
					Price:           price,
					ServiceDiscount: isDis,
					ProductDiscount: igDis,
				}
			}
		}
	}
	total := replyEs["total"].(float64)
	//去重
	bindIds = functions.ArrayUniqueInt(bindIds)
	shopIds = functions.ArrayUniqueInt(shopIds)

	//获取行业map
	rpcPublic := new(public.Indus).Init()
	defer rpcPublic.Close()
	var replyPublic []public2.ReplyIndusInfo
	if err = rpcPublic.GetIndusByBindIds(ctx, &bindIds, &replyPublic); err != nil {
		return nil, err
	}
	var bindMap = make(map[int]string)
	for _, info := range replyPublic {
		bindMap[info.IndusId] = info.Name
	}

	var regionMap = make(map[int]string)
	var districtMap = make(map[int]string)
	//if args.Flag == 1 {
	//获取区
	rpcPublic2 := new(public.Region).Init()
	defer rpcPublic2.Close()
	var replyPubilc2 []public2.RegionInfo
	if err = rpcPublic2.GetRegion(ctx, &public2.ArgsRegion{args.Cid}, &replyPubilc2); err != nil {
		return nil, err
	}
	for _, info := range replyPubilc2 {
		regionMap[info.RegionId] = info.RegionName
	}

	//获取商圈
	if len(districtIds) > 0 {
		districtIds = functions.ArrayUniqueInt(districtIds)
		var replyDistrictInfo []public2.ReplyDistrictInfo
		rpcPublic2.GetDistrictByDistrictIds(ctx, &public2.ArgsDistricts{districtIds}, &replyDistrictInfo)
		for _, info := range replyDistrictInfo {
			districtMap[info.DistrictId] = info.DistrictName
		}
	}
	//}

	//获取图片和详细地址和电话号码
	rpcShop := new(bus.Shop).Init()
	defer rpcShop.Close()
	replyShopInfos := map[int]map[string]interface{}{}
	if err = rpcShop.GetShopAddressAndImgByIdS(ctx, &shopIds, &replyShopInfos); err != nil {
		return nil, err
	}

	for i := range result {
		source := result[i].(map[string]interface{})
		var bindNameStr string
		if source["MainBindId"] != nil {
			bindNameStr += bindMap[int(source["MainBindId"].(float64))] + ","
		}
		if source["BindId"] != nil {
			for _, v := range source["BindId"].([]interface{}) {
				bindId := int(v.(float64))
				if bindId != int(source["MainBindId"].(float64)) {
					bindNameStr += bindMap[bindId] + ","
				}
			}
		}
		source["BindNames"] = strings.TrimSuffix(bindNameStr, ",")
		//if args.Flag == 1 {
		source["DName"] = regionMap[int(source["Did"].(float64))]
		if source["DistrictId"] != nil {
			if len(source["DistrictId"].([]interface{})) > 0 {
				var DistrictNamesStr string
				for _, v := range source["DistrictId"].([]interface{}) {
					DistrictNamesStr += districtMap[int(v.(float64))] + ","
				}
				source["DistrictNames"] = strings.TrimSuffix(DistrictNamesStr, ",")
			}
		} else {
			source["DistrictId"] = 0
		}
		//}
		source["ShopAddress"] = replyShopInfos[int(source["ShopId"].(float64))]["Address"]
		source["ShopImage"] = replyShopInfos[int(source["ShopId"].(float64))]["Image"]
		source["ShopPhone"] = replyShopInfos[int(source["ShopId"].(float64))]["Phone"]
		if source["Active"] != nil && len(source["Active"].([]interface{})) > 0 {
			var m = make(map[int]string)
			for _, v := range source["Active"].([]interface{}) {
				m[int(v.(float64))] = cards.ActivityScreen[int(v.(float64))]
			}
			source["Active"] = m
		}
		if len(extendDataMap) > 0 {
			for i, item := range source["Sub"].([]interface{}) {
				source["Sub"].([]interface{})[i].(map[string]interface{})["ServicePeriod"] =
					extendDataMap[int(item.(map[string]interface{})["ItemId"].(float64))].ServicePeriod
				source["Sub"].([]interface{})[i].(map[string]interface{})["ServiceTime"] =
					extendDataMap[int(item.(map[string]interface{})["ItemId"].(float64))].ServiceTime
				source["Sub"].([]interface{})[i].(map[string]interface{})["Price"] =
					extendDataMap[int(item.(map[string]interface{})["ItemId"].(float64))].Price
				source["Sub"].([]interface{})[i].(map[string]interface{})["TotalNum"] =
					extendDataMap[int(item.(map[string]interface{})["ItemId"].(float64))].TotalNum
				source["Sub"].([]interface{})[i].(map[string]interface{})["ServiceDiscount"] =
					extendDataMap[int(item.(map[string]interface{})["ItemId"].(float64))].ServiceDiscount
				source["Sub"].([]interface{})[i].(map[string]interface{})["ProductDiscount"] =
					extendDataMap[int(item.(map[string]interface{})["ItemId"].(float64))].ProductDiscount
			}
		}
	}

	var lists []cards.AppInfo
	if err = mapstructure.WeakDecode(result, &lists); err != nil {
		return nil, err
	}

	return &cards.ReplyAppInfo{TotalNum: int(total), Lists: lists}, nil
}

//根据项目查询详情-适用门店详情
func (i *ItemLogic) GetDetailById(ctx context.Context, args *cards.ArgsShopList, reply *[]order.ReplyCableShopInfo) error {
	*reply = []order.ReplyCableShopInfo{}
	if args.ItemId <= 0 || args.ItemType < cards.ITEM_TYPE_single || args.ItemType > cards.ITEM_TYPE_icard {
		return toolLib.CreateKcErr(_const.PARAM_ERR)
	}
	var shopIds []int
	var maps []map[string]interface{}
	var busId int
	switch args.ItemType {
	case cards.ITEM_TYPE_single:
		model := new(models.SingleModel).Init()
		m := model.GetBySingleId(args.ItemId, []string{model.Field.F_bus_id})
		if len(m) > 0 {
			busId, _ = strconv.Atoi(m[model.Field.F_bus_id].(string))
		}
	case cards.ITEM_TYPE_sm:
		model := new(models.SmShopModel).Init()
		maps = model.GetBySmId(args.ItemId)
	case cards.ITEM_TYPE_card:
		model := new(models.CardShopModel).Init()
		maps = model.GetByCardId(args.ItemId)
	case cards.ITEM_TYPE_hcard:
		model := new(models.HcardShopModel).Init()
		maps = model.GetByHcardIDs([]int{args.ItemId})
	case cards.ITEM_TYPE_ncard:
		model := new(models.NCardShopModel).Init()
		maps = model.GetByNCardIDs([]int{args.ItemId})
	case cards.ITEM_TYPE_hncard:
		model := new(models.HNCardShopModel).Init()
		maps = model.GetByHNCardIDs([]int{args.ItemId})
	case cards.ITEM_TYPE_rcard:
		model := new(models.RcardShopModel).Init()
		maps = model.GetByRcardId(args.ItemId)
	case cards.ITEM_TYPE_icard:
		model := new(models.IcardShopModel).Init()
		maps = model.GetByIcardId(args.ItemId)
	}
	var replyShopInfos []bus2.ReplyShopInfos
	rpcShop := new(bus.Shop).Init()
	defer rpcShop.Close()
	var flag bool
	if busId > 0 {
		flag = true
		if err := rpcShop.GetAvailableShopByBusId(ctx, &bus2.ArgsAvaBusId{
			BusId: busId, Lng: args.Lng, Lat: args.Lat,
		}, &replyShopInfos); err != nil {
			return err
		}
	}
	for _, m := range maps {
		if m["shop_id"].(string) == "0" {
			flag = true
			busId, _ := strconv.Atoi(m["bus_id"].(string))
			if err := rpcShop.GetAvailableShopByBusId(ctx, &bus2.ArgsAvaBusId{
				BusId: busId, Lng: args.Lng, Lat: args.Lat,
			}, &replyShopInfos); err != nil {
				return err
			}
			break
		}
	}
	if !flag && len(maps) > 0 {
		shopIds = functions.ArrayValue2Array("shop_id", maps)
		if err := rpcShop.GetAvailableShopByShopids(ctx, &bus2.ArgsAvaGetShops{
			ShopIds: shopIds, Lng: args.Lng, Lat: args.Lat,
		}, &replyShopInfos); err != nil {
			return nil
		}
	}

	var bindIds []int
	var imgIds []int
	for _, info := range replyShopInfos {
		bindIds = append(bindIds, info.MainBindId)
		bindIds = append(bindIds, functions.StrExplode2IntArr(info.BindId, ",")...)
		imgIds = append(imgIds, info.ShopImg)
	}
	//获取行业名称
	bindIds = functions.ArrayUniqueInt(bindIds)
	var indusMap = make(map[int]string)
	if len(bindIds) > 0 {
		rpcPub := new(public.Indus).Init()
		defer rpcPub.Close()
		var replyPublic []public2.ReplyIndusInfo
		if err := rpcPub.GetIndusByBindIds(ctx, &bindIds, &replyPublic); err != nil {
			return err
		}
		for _, v := range replyPublic {
			indusMap[v.IndusId] = v.Name
		}
	}
	//获取照片
	var replyFiles map[int]file2.ReplyFileInfo
	if len(imgIds) > 0 {
		rpcFile := new(file.Upload).Init()
		defer rpcFile.Close()
		if err := rpcFile.GetImageByIds(ctx, imgIds, &replyFiles); err != nil {
			return nil
		}
	}
	var r order.ReplyCableShopInfo
	for _, info := range replyShopInfos {
		r.ShopId = info.ShopId
		r.ShopName = info.ShopName
		r.BranchName = info.BranchName
		r.BindId = info.BindId
		r.MainBindId = info.MainBindId
		r.Longitude = info.Longitude
		r.Latitude = info.Latitude
		r.Distance = info.Distance
		r.Address = info.Address
		r.ContactCall = info.ContactCall
		r.Contact = info.Contact
		r.BusinessHours = info.BusinessHours
		if v, ok := replyFiles[info.ShopImg]; ok {
			r.ShopImgUrl = v.Path
			r.ShopImgHash = v.Hash
		}
		r.CompanyName = info.CompanyName
		r.IndustryId = info.IndustryId
		r.Status = info.Status
		r.Ctime = info.Ctime
		r.ShopImg = info.ShopImg

		r.IndusName = indusMap[info.MainBindId] + ","
		arr := functions.StrExplode2IntArr(info.BindId, ",")
		for _, i := range arr {
			r.IndusName += indusMap[i] + ","
		}
		r.IndusName = strings.TrimSuffix(r.IndusName, ",")
		*reply = append(*reply, r)
	}
	return nil
}

// 门店详情-门店拥有的项目-上架的项目
func (i *ItemLogic) GetItemsByShopId(ctx context.Context, args *cards.ArgsGetItemsByShopId, reply *cards.ReplyGetItemsByShopId) (err error) {
	shopId, itemType := args.ShopId, args.ItemType
	reply.Lists = []cards.ReplyGetItemsByShopIdBase{}
	reply.IndexImg = map[int]file2.ReplyFileInfo{}
	if shopId == 0 {
		return
	}

	sirModel := new(models.ShopItemRelationModel).Init()
	start, limit := args.GetStart(), args.GetPageSize()
	where := []base.WhereItem{
		{sirModel.Field.F_shop_id, shopId},
		{sirModel.Field.F_is_del, cards.RELATION_IS_DEL_NO},
		{sirModel.Field.F_status, cards.STATUS_ON_SALE},
	}
	if itemType > 0 {
		where = append(where, base.WhereItem{sirModel.Field.F_item_type, itemType})
	} else {
		where = append(where, base.WhereItem{sirModel.Field.F_item_type, []interface{}{"gt", cards.ITEM_TYPE_single}}) //过滤掉单项目
	}
	sirMaps := sirModel.SelectByWhereByPage(where, start, limit)
	if len(sirMaps) == 0 {
		return
	}
	reply.TotalNum = sirModel.GetTotalNumByWhere(where)
	shopItemIdMap := make(map[int][]int) //key:卡项类型，value:卡项id
	for _, v := range sirMaps {
		typ, _ := strconv.Atoi(v[sirModel.Field.F_item_type].(string))
		itmeId, _ := strconv.Atoi(v[sirModel.Field.F_item_id].(string))
		if _, ok := shopItemIdMap[typ]; !ok {
			shopItemIdMap[typ] = make([]int, 0)
		}
		shopItemIdMap[typ] = append(shopItemIdMap[typ], itmeId)
	}

	var (
		imgIds []int
	)
	for cardType, itemIds := range shopItemIdMap {
		switch cardType {
		case cards.ITEM_TYPE_single:
			// 分店项目
			shopSingleModel := new(models.ShopSingleModel).Init()
			ssRes := shopSingleModel.GetByShopidAndSingleids(shopId, itemIds)
			if len(ssRes) == 0 {
				continue
			}
			var shopSingleStruct = []struct {
				SsId             int
				SingleId         int
				ChangedRealPrice float64
				ChangedMinPrice  float64
				ChangedMaxPrice  float64
				Status           int
				Sales            int
			}{}
			_ = mapstructure.WeakDecode(ssRes, &shopSingleStruct)
			var realPrice float64
			list := make([]cards.ReplyGetItemsByShopIdBase, 0)
			for _, v := range shopSingleStruct {
				if v.ChangedMinPrice > 0 {
					realPrice = v.ChangedMinPrice
				} else {
					realPrice = v.ChangedRealPrice
				}
				list = append(list, cards.ReplyGetItemsByShopIdBase{
					ItemType:  cardType,
					ItemId:    v.SingleId,
					Name:      "",
					ImgId:     0,
					RealPrice: realPrice,
					Price:     0,
					Sales:     v.Sales,
					ShopId:    shopId,
					SsId:      v.SsId,
					Discount:  "10",
				})
			}
			singleIds := functions.ArrayValue2Array(shopSingleModel.Field.F_single_id, ssRes)
			singleIds = functions.ArrayUniqueInt(singleIds)
			//	单项目数据
			singleModel := new(models.SingleModel).Init()
			singles := singleModel.GetBySingleids(singleIds)
			for index, v := range list {
				if v.ItemType == cards.ITEM_TYPE_single {
					for _, single := range singles {
						if strconv.Itoa(v.ItemId) == single[singleModel.Field.F_single_id].(string) {
							list[index].Name = single[singleModel.Field.F_name].(string)
							list[index].Price, _ = strconv.ParseFloat(single[singleModel.Field.F_price].(string), 64)
							imgId, _ := strconv.Atoi(single[singleModel.Field.F_img_id].(string))
							imgIds = append(imgIds, imgId)
							list[index].ImgId = imgId
							list[index].ServicePeriod, _ = strconv.Atoi(single[singleModel.Field.F_service_time].(string))
							break
						}
					}
				}
			}
			reply.Lists = append(reply.Lists, list...)
		case cards.ITEM_TYPE_sm:
			shopSmModel := new(models.ShopSmModel).Init()
			ssRes := shopSmModel.GetByShopidAdSmids(shopId, itemIds)
			if len(ssRes) == 0 {
				continue
			}
			list := make([]cards.ReplyGetItemsByShopIdBase, 0)
			for _, v := range ssRes {
				smId, _ := strconv.Atoi(v[shopSmModel.Field.F_sm_id].(string))
				sales, _ := strconv.Atoi(v[shopSmModel.Field.F_sales].(string))
				id, _ := strconv.Atoi(v[shopSmModel.Field.F_id].(string))
				list = append(list, cards.ReplyGetItemsByShopIdBase{
					ItemType:  cardType,
					ItemId:    smId,
					Name:      "",
					ImgId:     0,
					RealPrice: 0,
					Price:     0,
					Sales:     sales,
					ShopId:    shopId,
					SsId:      id,
					Discount:  "10",
				})
			}
			//	sm
			smIds := functions.ArrayValue2Array(shopSmModel.Field.F_sm_id, ssRes)
			smIds = functions.ArrayUniqueInt(smIds)
			smModel := new(models.SmModel).Init()
			sms := smModel.GetBySmids(smIds)
			for index, v := range list {
				if v.ItemType == cards.ITEM_TYPE_sm {
					for _, card := range sms {
						if strconv.Itoa(v.ItemId) == card[smModel.Field.F_sm_id].(string) {
							imgId, _ := strconv.Atoi(card[smModel.Field.F_img_id].(string))
							list[index].ImgId = imgId
							imgIds = append(imgIds, imgId)
							list[index].Name = card[smModel.Field.F_name].(string)
							list[index].RealPrice, _ = strconv.ParseFloat(card[smModel.Field.F_real_price].(string), 64)
							list[index].Price, _ = strconv.ParseFloat(card[smModel.Field.F_price].(string), 64)
							list[index].ServicePeriod, _ = strconv.Atoi(card[smModel.Field.F_service_period].(string))
							list[index].ValidCount, _ = strconv.Atoi(card[smModel.Field.F_validcount].(string))
							break
						}
					}
				}
			}
			reply.Lists = append(reply.Lists, list...)
		case cards.ITEM_TYPE_card:
			shopCardModel := new(models.ShopCardModel).Init()

			scRes := shopCardModel.GetByShopIDAndCardIDs(shopId, itemIds)
			if len(scRes) == 0 {
				continue
			}
			list := make([]cards.ReplyGetItemsByShopIdBase, 0)
			for _, v := range scRes {
				cardId, _ := strconv.Atoi(v[shopCardModel.Field.F_card_id].(string))
				sales, _ := strconv.Atoi(v[shopCardModel.Field.F_sales].(string))
				id, _ := strconv.Atoi(v[shopCardModel.Field.F_id].(string))
				list = append(list, cards.ReplyGetItemsByShopIdBase{
					ItemType:  cardType,
					ItemId:    cardId,
					Name:      "",
					ImgId:     0,
					RealPrice: 0,
					Price:     0,
					Sales:     sales,
					ShopId:    shopId,
					SsId:      id,
					Discount:  "10",
				})
			}
			cardIds := functions.ArrayValue2Array(shopCardModel.Field.F_card_id, scRes)
			cardIds = functions.ArrayUniqueInt(cardIds)

			cardModel := new(models.CardModel).Init()
			cards2 := cardModel.GetByCardIDs(cardIds)
			for index, v := range list {
				if v.ItemType == cards.ITEM_TYPE_card {
					for _, card := range cards2 {
						if strconv.Itoa(v.ItemId) == card[cardModel.Field.F_card_id].(string) {
							imgId, _ := strconv.Atoi(card[cardModel.Field.F_img_id].(string))
							imgIds = append(imgIds, imgId)
							list[index].ImgId = imgId
							list[index].Name = card[cardModel.Field.F_name].(string)
							list[index].RealPrice, _ = strconv.ParseFloat(card[cardModel.Field.F_real_price].(string), 64)
							list[index].Price, _ = strconv.ParseFloat(card[cardModel.Field.F_price].(string), 64)
							list[index].ServicePeriod, _ = strconv.Atoi(card[cardModel.Field.F_service_period].(string))
							break
						}
					}
				}
			}
			reply.Lists = append(reply.Lists, list...)
		case cards.ITEM_TYPE_hcard:
			shopHcardModel := new(models.ShopHcardModel).Init()
			shRes := shopHcardModel.GetByShopIDAndHcardIDs(shopId, itemIds)
			if len(shRes) == 0 {
				continue
			}
			list := make([]cards.ReplyGetItemsByShopIdBase, 0)
			for _, v := range shRes {
				hcardId, _ := strconv.Atoi(v[shopHcardModel.Field.F_hcard_id].(string))
				sales, _ := strconv.Atoi(v[shopHcardModel.Field.F_sales].(string))
				id, _ := strconv.Atoi(v[shopHcardModel.Field.F_id].(string))
				list = append(list, cards.ReplyGetItemsByShopIdBase{
					ItemType:  cardType,
					ItemId:    hcardId,
					Name:      "",
					ImgId:     0,
					RealPrice: 0,
					Price:     0,
					Sales:     sales,
					ShopId:    shopId,
					SsId:      id,
					Discount:  "10",
				})
			}
			hcardIds := functions.ArrayValue2Array(shopHcardModel.Field.F_hcard_id, shRes)
			hcardIds = functions.ArrayUniqueInt(hcardIds)
			//
			hcardModel := new(models.HcardModel).Init()
			hcards := hcardModel.GetByHcardIDsAndGround(hcardIds)
			for index, v := range list {
				if v.ItemType == cards.ITEM_TYPE_hcard {
					for _, card := range hcards {
						if strconv.Itoa(v.ItemId) == card[hcardModel.Field.F_hcard_id].(string) {
							imgId, _ := strconv.Atoi(card[hcardModel.Field.F_img_id].(string))
							imgIds = append(imgIds, imgId)
							list[index].ImgId = imgId
							list[index].Name = card[hcardModel.Field.F_name].(string)
							list[index].RealPrice, _ = strconv.ParseFloat(card[hcardModel.Field.F_real_price].(string), 64)
							list[index].Price, _ = strconv.ParseFloat(card[hcardModel.Field.F_price].(string), 64)
							list[index].ServicePeriod, _ = strconv.Atoi(card[hcardModel.Field.F_service_period].(string))
							break
						}
					}
				}
			}
			reply.Lists = append(reply.Lists, list...)
		case cards.ITEM_TYPE_ncard:
			shopNCardModel := new(models.ShopNCardModel).Init()
			sncRes := shopNCardModel.GetByShopIDAndNCardIDs(shopId, itemIds)
			if len(sncRes) == 0 {
				continue
			}
			list := make([]cards.ReplyGetItemsByShopIdBase, 0)
			for _, v := range sncRes {
				hncardId, _ := strconv.Atoi(v[shopNCardModel.Field.F_ncard_id].(string))
				sales, _ := strconv.Atoi(v[shopNCardModel.Field.F_sales].(string))
				id, _ := strconv.Atoi(v[shopNCardModel.Field.F_id].(string))
				list = append(list, cards.ReplyGetItemsByShopIdBase{
					ItemType:  cardType,
					ItemId:    hncardId,
					Name:      "",
					ImgId:     0,
					RealPrice: 0,
					Price:     0,
					Sales:     sales,
					ShopId:    shopId,
					SsId:      id,
					Discount:  "10",
				})
			}
			ncardIds := functions.ArrayValue2Array(shopNCardModel.Field.F_ncard_id, sncRes)
			ncardIds = functions.ArrayUniqueInt(ncardIds)
			//
			nCardModel := new(models.NCardModel).Init()
			ncards := nCardModel.GetByNCardIDs(ncardIds)
			for index, v := range list {
				if v.ItemType == cards.ITEM_TYPE_ncard {
					for _, card := range ncards {
						if strconv.Itoa(v.ItemId) == card[nCardModel.Field.F_ncard_id].(string) {
							imgId, _ := strconv.Atoi(card[nCardModel.Field.F_img_id].(string))
							imgIds = append(imgIds, imgId)
							list[index].ImgId = imgId
							list[index].Name = card[nCardModel.Field.F_name].(string)
							list[index].RealPrice, _ = strconv.ParseFloat(card[nCardModel.Field.F_real_price].(string), 64)
							list[index].Price, _ = strconv.ParseFloat(card[nCardModel.Field.F_price].(string), 64)
							list[index].ServicePeriod, _ = strconv.Atoi(card[nCardModel.Field.F_service_period].(string))
							list[index].ValidCount, _ = strconv.Atoi(card[nCardModel.Field.F_validcount].(string))
							break
						}
					}
				}
			}
			reply.Lists = append(reply.Lists, list...)
		case cards.ITEM_TYPE_hncard:
			shopHNCardModel := new(models.ShopHNCardModel).Init()
			shnRes := shopHNCardModel.GetByShopIDAndHNCardIDs(shopId, itemIds)
			if len(shnRes) == 0 {
				continue
			}
			list := make([]cards.ReplyGetItemsByShopIdBase, 0)
			for _, v := range shnRes {
				hncardId, _ := strconv.Atoi(v[shopHNCardModel.Field.F_hncard_id].(string))
				sales, _ := strconv.Atoi(v[shopHNCardModel.Field.F_sales].(string))
				id, _ := strconv.Atoi(v[shopHNCardModel.Field.F_id].(string))
				list = append(list, cards.ReplyGetItemsByShopIdBase{
					ItemType:  cardType,
					ItemId:    hncardId,
					Name:      "",
					ImgId:     0,
					RealPrice: 0,
					Price:     0,
					Sales:     sales,
					ShopId:    shopId,
					SsId:      id,
					Discount:  "10",
				})
			}
			hncardIds := functions.ArrayValue2Array(shopHNCardModel.Field.F_hncard_id, shnRes)
			hncardIds = functions.ArrayUniqueInt(hncardIds)

			//
			hNCardModel := new(models.HNCardModel).Init()
			hnCards := hNCardModel.GetByHNCardIDs(hncardIds)
			for index, v := range list {
				if v.ItemType == cards.ITEM_TYPE_hncard {
					for _, card := range hnCards {
						if strconv.Itoa(v.ItemId) == card[hNCardModel.Field.F_hncard_id].(string) {
							imgId, _ := strconv.Atoi(card[hNCardModel.Field.F_img_id].(string))
							imgIds = append(imgIds, imgId)
							list[index].ImgId = imgId
							list[index].Name = card[hNCardModel.Field.F_name].(string)
							list[index].RealPrice, _ = strconv.ParseFloat(card[hNCardModel.Field.F_real_price].(string), 64)
							list[index].Price, _ = strconv.ParseFloat(card[hNCardModel.Field.F_price].(string), 64)
							list[index].ServicePeriod, _ = strconv.Atoi(card[hNCardModel.Field.F_service_period].(string))
							list[index].ValidCount, _ = strconv.Atoi(card[hNCardModel.Field.F_validcount].(string))
							break
						}
					}
				}
			}
			reply.Lists = append(reply.Lists, list...)
		case cards.ITEM_TYPE_rcard:
			shopRcardModel := new(models.ShopRcardModel).Init()
			srRes := shopRcardModel.GetByShopIDAndCardIDs(shopId, itemIds)
			if len(srRes) == 0 {
				continue
			}
			list := make([]cards.ReplyGetItemsByShopIdBase, 0)
			for _, v := range srRes {
				cardId, _ := strconv.Atoi(v[shopRcardModel.Field.F_rcard_id].(string))
				sales, _ := strconv.Atoi(v[shopRcardModel.Field.F_sales].(string))
				id, _ := strconv.Atoi(v[shopRcardModel.Field.F_id].(string))
				list = append(list, cards.ReplyGetItemsByShopIdBase{
					ItemType:  cardType,
					ItemId:    cardId,
					Name:      "",
					ImgId:     0,
					RealPrice: 0,
					Price:     0,
					Sales:     sales,
					ShopId:    shopId,
					SsId:      id,
					Discount:  "10",
				})
			}
			rcardIds := functions.ArrayValue2Array(shopRcardModel.Field.F_rcard_id, srRes)
			rcardIds = functions.ArrayUniqueInt(rcardIds)
			rsModel := new(models.RcardSingleModel).Init()
			rsMaps := rsModel.GetByRcardIds(rcardIds)

			rcardModel := new(models.RcardModel).Init()
			rcards := rcardModel.GetRcardsByRcardIds(rcardIds)
			for index, v := range list {
				if v.ItemType == cards.ITEM_TYPE_rcard {
					for _, card := range rcards {
						if strconv.Itoa(v.ItemId) == card[rcardModel.Field.F_rcard_id].(string) {
							imgId, _ := strconv.Atoi(card[rcardModel.Field.F_img_id].(string))
							imgIds = append(imgIds, imgId)
							list[index].ImgId = imgId
							list[index].Name = card[rcardModel.Field.F_name].(string)
							list[index].RealPrice, _ = strconv.ParseFloat(card[rcardModel.Field.F_real_price].(string), 64)
							list[index].Price, _ = strconv.ParseFloat(card[rcardModel.Field.F_price].(string), 64)
							list[index].ServicePeriod, _ = strconv.Atoi(card[rcardModel.Field.F_service_period].(string))
							break
						}
					}
					for _, info := range rsMaps {
						if strconv.Itoa(v.ItemId) == info[rsModel.Field.F_rcard_id].(string) {
							list[index].Discount = info[rsModel.Field.F_discount].(string)
							break
						}
					}
				}
			}
			reply.Lists = append(reply.Lists, list...)
		case cards.ITEM_TYPE_icard:
			shopIcardModel := new(models.ShopIcardModel).Init()

			siRes := shopIcardModel.GetByShopIDAndHNCardIDs(shopId, itemIds)
			if len(siRes) == 0 {
				continue
			}
			list := make([]cards.ReplyGetItemsByShopIdBase, 0)
			icardIds := functions.ArrayValue2Array(shopIcardModel.Field.F_icard_id, siRes)
			icardIds = functions.ArrayUniqueInt(icardIds)
			for _, v := range siRes {
				cardId, _ := strconv.Atoi(v[shopIcardModel.Field.F_icard_id].(string))
				sales, _ := strconv.Atoi(v[shopIcardModel.Field.F_sales].(string))
				id, _ := strconv.Atoi(v[shopIcardModel.Field.F_id].(string))
				list = append(list, cards.ReplyGetItemsByShopIdBase{
					ItemType:  cardType,
					ItemId:    cardId,
					Name:      "",
					ImgId:     0,
					RealPrice: 0,
					Price:     0,
					Sales:     sales,
					ShopId:    shopId,
					SsId:      id,
					Discount:  "10",
				})
			}
			//身份卡折扣
			icModel := new(models.IcardSingleModel).Init()
			icardLists := icModel.GetByIcardIds(icardIds)
			icardDiscountMap := make(map[string][]map[string]interface{})
			for _, icardList := range icardLists {
				icardId := icardList[icModel.Field.F_icard_id].(string)
				if _, ok := icardDiscountMap[icardId]; !ok {
					icardDiscountMap[icardId] = make([]map[string]interface{}, 0)
				}
				icardDiscountMap[icardId] = append(icardDiscountMap[icardId], icardList)
			}

			icardModel := new(models.IcardModel).Init()
			icards := icardModel.GetRcardsByIcardIds(icardIds)
			for index, v := range list {
				if v.ItemType == cards.ITEM_TYPE_icard {
					for _, card := range icards {
						if strconv.Itoa(v.ItemId) == card[icardModel.Field.F_icard_id].(string) {
							imgId, _ := strconv.Atoi(card[icardModel.Field.F_img_id].(string))
							imgIds = append(imgIds, imgId)
							list[index].ImgId = imgId
							list[index].Name = card[icardModel.Field.F_name].(string)
							list[index].RealPrice, _ = strconv.ParseFloat(card[icardModel.Field.F_real_price].(string), 64)
							list[index].Price, _ = strconv.ParseFloat(card[icardModel.Field.F_price].(string), 64)
							list[index].ServicePeriod, _ = strconv.Atoi(card[icardModel.Field.F_service_period].(string))
							break
						}
					}
					for icardIdStr, icardList := range icardDiscountMap {
						if strconv.Itoa(v.ItemId) == icardIdStr {
							minDiscount := icardList[0][icModel.Field.F_discount].(string)
							if len(icardList) == 1 {
								list[index].Discount = minDiscount
							} else {
								maxDiscount := icardList[len(icardList)-1][icModel.Field.F_discount].(string)
								list[index].Discount = fmt.Sprintf("%s~%s", minDiscount, maxDiscount)
							}
						}
					}
				}
			}
			reply.Lists = append(reply.Lists, list...)
		}
	}

	imgIds = functions.ArrayUniqueInt(imgIds)
	reply.DefaultImgs = make(map[int]file2.ReplyFileInfo)
	// rpc file
	rpcFile := new(file.Upload).Init()
	var respFile map[int]file2.ReplyFileInfo
	if err = rpcFile.GetImageByIds(ctx, imgIds, &respFile); err != nil {
		return
	}
	if itemType > 0 {
		respFile[0] = file2.ReplyFileInfo{
			Id:   0,
			Hash: "",
			Path: constkey.CardsSmallDefaultPics[args.ItemType],
		}
	} else {
		//获取全部卡项的默认图片
		for itype, path := range constkey.CardsSmallDefaultPics {
			reply.DefaultImgs[itype] = file2.ReplyFileInfo{
				Id:   0,
				Hash: "",
				Path: path,
			}
		}
	}
	reply.IndexImg = respFile

	return
}

//购买成功卡项的销量
func (i *ItemLogic) IncrItemSales(ctx context.Context, orderSn string, reply *bool) (err error) {
	if len(orderSn) == 0 {
		toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	prefixSn := fmt.Sprintf("%s%s", string(orderSn[0]), string(orderSn[1]))
	if prefixSn == order.PREFIX_GOODS {
		return
	}
	//获取订单详情
	rpcOrder := new(order2.ItemOrder).Init()
	replyOrder := order.ReplyGetOrderInfoByOrderSnRpc{}
	err = rpcOrder.GetOrderInfoByOrderSnRpc(ctx, &orderSn, &replyOrder)
	if err != nil {
		return
	}
	switch replyOrder.OrderType {
	case order.PAY_ORDER_TYPE_SINGLE:
		//单项目
		mSingle := new(models.SingleModel).Init()
		mSSP := new(models.SingleSpecPriceModel).Init()
		mShopSingle := new(models.ShopSingleModel).Init()
		mSSSP := new(models.ShopSingleSpecPriceModel).Init()
		for _, orderItem := range replyOrder.OrderItems {
			//修改总的销售量
			mSingle.IncrSalesBySingleid(orderItem.ItemId, orderItem.Num)
			//修改门店的销售量
			mShopSingle.IncrSalesBySingleidAndShopid(orderItem.ItemId, replyOrder.ShopId, orderItem.Num)
			//修改规格的销量
			if orderItem.SspId > 0 {
				mSSP.IncrSalesBySspid(orderItem.SspId, orderItem.Num)
				mSSSP.IncrSalesByShopidAndSspid(replyOrder.ShopId, orderItem.SspId, orderItem.Num)
			}
		}
	case order.PAY_ORDER_TYPE_CARD:
		//卡项
		switch replyOrder.CardOrderType {
		case order.ORDER_TYPE_SM:
			//套餐
			mSm := new(models.SmModel).Init()
			mShopSm := new(models.ShopSmModel).Init(mSm.Model.GetOrmer())
			for _, orderItem := range replyOrder.OrderItems {
				mSm.IncrSalesBySmid(orderItem.ItemId, orderItem.Num)
				mShopSm.IncrSalesByShopidAndSmid(replyOrder.ShopId, orderItem.ItemId, orderItem.Num)
			}
		case order.ORDER_TYPE_CARD:
			mCard := new(models.CardModel).Init()
			mShopCard := new(models.ShopCardModel).Init(mCard.Model.GetOrmer())
			for _, orderItem := range replyOrder.OrderItems {
				mCard.IncrSalesByCardID(orderItem.ItemId, orderItem.Num)
				mShopCard.IncrSalesByShopidAndCardid(replyOrder.ShopId, orderItem.ItemId, orderItem.Num)
			}
		case order.ORDER_TYPE_HCARD:
			mHcard := new(models.HcardModel).Init()
			mShopHcard := new(models.ShopHcardModel).Init(mHcard.Model.GetOrmer())
			for _, orderItem := range replyOrder.OrderItems {
				mHcard.IncrSalesByHcardID(orderItem.ItemId, orderItem.Num)
				mShopHcard.IncrSalesByShopidAndCardid(replyOrder.ShopId, orderItem.ItemId, orderItem.Num)
			}
		case order.ORDER_TYPE_NCARD:
			mNcard := new(models.NCardModel).Init()
			mShopNcard := new(models.ShopNCardModel).Init(mNcard.Model.GetOrmer())
			for _, orderItem := range replyOrder.OrderItems {
				mNcard.IncrSalesByNCardID(orderItem.ItemId, orderItem.Num)
				mShopNcard.IncrSalesByShopidAndCardid(replyOrder.ShopId, orderItem.ItemId, orderItem.Num)
			}
		case order.ORDER_TYPE_HNCARD:
			mHncard := new(models.HNCardModel).Init()
			mShopHncard := new(models.ShopHNCardModel).Init(mHncard.Model.GetOrmer())
			for _, orderItem := range replyOrder.OrderItems {
				mHncard.IncrSalesByHNCardID(orderItem.ItemId, orderItem.Num)
				mShopHncard.IncrSalesByShopidAndCardid(replyOrder.ShopId, orderItem.ItemId, orderItem.Num)
			}
		case order.ORDER_TYPE_RCARD:
			mHncard := new(models.RcardModel).Init()
			mShopHncard := new(models.ShopRcardModel).Init(mHncard.Model.GetOrmer())
			for _, orderItem := range replyOrder.OrderItems {
				mHncard.IncrSalesByRcardID(orderItem.ItemId, orderItem.Num)
				mShopHncard.IncrSalesByShopidAndRcardid(replyOrder.ShopId, orderItem.ItemId, orderItem.Num)
			}
		case order.ORDER_TYPE_ICARD:
			mHncard := new(models.IcardModel).Init()
			mShopHncard := new(models.ShopIcardModel).Init(mHncard.Model.GetOrmer())
			for _, orderItem := range replyOrder.OrderItems {
				mHncard.IncrSalesByIcardID(orderItem.ItemId, orderItem.Num)
				mShopHncard.IncrSalesByShopidAndIcardid(replyOrder.ShopId, orderItem.ItemId, orderItem.Num)
			}
		}
	}
	return
}

//卡项收藏入参
func (i *ItemLogic) CollectItems(ctx context.Context, args *cards.ArgsCollectItems, reply *bool) (err error) {
	if args.ItemType < cards.ITEM_TYPE_single || args.ItemType > cards.ITEM_TYPE_icard {
		return toolLib.CreateKcErr(_const.PARAM_ERR)
	}
	// 检查数据是否存在
	items, e := i.GetItemsBySsids(ctx, &cards.ArgsGetItemsBySsids{SsIds: []int{args.SsId}, ItemType: args.ItemType})
	if e != nil {
		err = e
		return
	}
	if len(items) == 0 {
		return toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
	}
	item := items[cards.SsId(args.SsId)]
	shopId := item.ShopId
	if item.Status != cards.STATUS_ON_SALE {
		return toolLib.CreateKcErr(_const.SHOP_SINGLE_OFF_SALE)
	}
	// 检查店铺是否存在
	rpcBus := new(bus.Shop).Init()
	defer rpcBus.Close()
	var busShopRes bus2.ReplyShopInfo
	if err = rpcBus.GetShopByShopid(ctx, &bus2.ArgsGetShop{ShopId: shopId}, &busShopRes); err != nil {
		return
	}

	ccm := new(models.CardsCollectModel).Init()
	collectInfo := ccm.Find(map[string]interface{}{ccm.Field.F_uid: args.Uid, ccm.Field.F_item_id: args.ItemId, ccm.Field.F_item_type: args.ItemType,
		ccm.Field.F_shop_id: shopId}, ccm.Field.F_id)

	if len(collectInfo) == 0 { // 未收藏，执行收藏逻辑
		if ccm.Insert(map[string]interface{}{
			ccm.Field.F_shop_id: shopId, ccm.Field.F_item_type: args.ItemType, ccm.Field.F_item_id: args.ItemId, ccm.Field.F_ss_id: args.SsId,
			ccm.Field.F_uid: args.Uid, ccm.Field.F_bus_id: busShopRes.BusId, ccm.Field.F_ctime: time.Now().Unix(),
		}) == 0 {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	} else { // 已收藏,执行取消逻辑
		if !ccm.Delete(map[string]interface{}{ccm.Field.F_id: collectInfo[ccm.Field.F_id]}) {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}
	*reply = true
	return
}

//卡项收藏状态
func (i *ItemLogic) GetCollectStatus(ctx context.Context, args *cards.ArgsCollectStatus, reply *cards.ReplyCollectStatus) (err error) {

	if args.ItemType < cards.ITEM_TYPE_single || args.ItemType > cards.ITEM_TYPE_icard {
		return toolLib.CreateKcErr(_const.PARAM_ERR)
	}
	ccm := new(models.CardsCollectModel).Init()
	collectInfo := ccm.Find(map[string]interface{}{ccm.Field.F_uid: args.Uid, ccm.Field.F_item_id: args.ItemId, ccm.Field.F_item_type: args.ItemType,
		ccm.Field.F_ss_id: args.SsId}, ccm.Field.F_id)

	if len(collectInfo) == 0 { // 未收藏
		reply.CollectStatus = 0
	} else { // 已收藏
		reply.CollectStatus = 1
	}

	return
}

//获取用户收藏的卡项入参
func (i *ItemLogic) GetCollectItems(ctx context.Context, args *cards.ArgsGetCollectItems, reply *cards.ReplyGetCollectItems) (err error) {
	where := make(map[string]interface{})
	ccm := new(models.CardsCollectModel).Init()
	where[ccm.Field.F_uid] = args.Uid
	if args.ShopId > 0 {
		where[ccm.Field.F_shop_id] = args.ShopId
	}
	if args.ItemType > 0 {
		where[ccm.Field.F_item_type] = args.ItemType
	}
	collectMaps := ccm.SelectByPage(where, args.GetStart(), args.GetPageSize())
	if len(collectMaps) == 0 {
		return
	}
	var collectItemMap = make(map[int][]map[string]interface{})
	var shopBusIdMap = make(map[int]int)
	for _, v := range collectMaps {
		itemType, _ := strconv.Atoi(v[ccm.Field.F_item_type].(string))
		_, ok := collectItemMap[itemType]
		if !ok {
			collectItemMap[itemType] = make([]map[string]interface{}, 0)
		}
		collectItemMap[itemType] = append(collectItemMap[itemType], v)
		shopId, _ := strconv.Atoi(v[ccm.Field.F_shop_id].(string))
		if _, ok := shopBusIdMap[shopId]; !ok {
			busId, _ := strconv.Atoi(v[ccm.Field.F_bus_id].(string))
			shopBusIdMap[shopId] = busId
		}
	}
	var collectStruct = make(map[int][]int)
	for index, v := range collectItemMap {
		collectStruct[index] = functions.ArrayUniqueInt(functions.ArrayValue2Array(ccm.Field.F_ss_id, v))
	}
	var list []cards.ReplyGetItemsByShopIdBase
	var imgIds []int
	var realPrice float64
	for itemType, ssIds := range collectStruct {
		if len(ssIds) == 0 {
			continue
		}
		items, e := i.GetItemsBySsids(ctx, &cards.ArgsGetItemsBySsids{SsIds: ssIds, ItemType: itemType})
		if e != nil {
			err = e
			return
		}

		for _, item := range items {
			if item.ShopRealPrice > 0 {
				realPrice = item.ShopRealPrice
			} else {
				realPrice = item.RealPrice
			}
			list = append(list, cards.ReplyGetItemsByShopIdBase{
				ItemId: item.ItemId, Name: item.ItemName, ImgId: item.ImgId, RealPrice: realPrice, Price: item.Price,
				Sales: item.ShopSales, SsId: item.SsId, ServicePeriod: item.ServicePeriod, ValidCount: item.ValidCount,
				ItemType: itemType, ShopId: item.ShopId, BusId: shopBusIdMap[item.ShopId],
			})
			imgIds = append(imgIds, item.ImgId)
		}
	}

	// rpcFile
	imgIds = functions.ArrayUniqueInt(imgIds)
	if len(imgIds) > 0 {
		rpcFile := new(file.Upload).Init()
		defer rpcFile.Close()
		var rpcFileRes map[int]file2.ReplyFileInfo
		if err = rpcFile.GetImageByIds(ctx, imgIds, &rpcFileRes); err != nil {
			return
		}
		reply.IndexImg = rpcFileRes
	}
	reply.Lists = list
	reply.TotalNum = ccm.GetTotalNum(where)
	return
}

func (i *ItemLogic) GetItemAllXCradsNumByShopId(ctx context.Context, shopId int, reply *cards.ReplyGetItemAllXCradsNumByShopId) (err error) {
	rpcShop := new(bus.Shop).Init()
	defer rpcShop.Close()
	var rpcShopRes bus2.ReplyShopInfo
	if err = rpcShop.GetShopByShopid(ctx, &bus2.ArgsGetShop{ShopId: shopId}, &rpcShopRes); err != nil {
		return
	}
	nowUnix := time.Now().Unix()
	firstMonth, lastMonth := tools.GetFirstAndLastOfMonth(nowUnix)
	busId := rpcShopRes.BusId
	reply.BusId = busId
	reply.DateTime = nowUnix

	////单项目
	//siM := new(models.SingleModel).Init()
	//simMaps := siM.GetList(map[string]interface{}{siM.Field.F_bus_id: busId}, []string{siM.Field.F_single_id})
	//if len(simMaps) > 0 {
	//	singleIds := functions.ArrayValue2Array(siM.Field.F_single_id, simMaps)
	//	ssm := new(models.ShopSingleModel).Init()
	//	simWhere := map[string]interface{}{ssm.Field.F_single_id: []interface{}{"IN", singleIds}}
	//	simNum := ssm.GetTotal(simWhere)
	//	simWhere[ssm.Field.F_under_time] = []interface{}{"between", []int64{firstMonth, lastMonth}}
	//	simMonthUnderNum := ssm.GetTotal(simWhere) //月下架量
	//	reply.AllItemXCardNum = reply.AllItemXCardNum + simNum
	//	reply.AllUnderItemXCardNum = reply.AllUnderItemXCardNum + simMonthUnderNum
	//}

	//套餐
	sM := new(models.SmModel).Init()
	sMMaps := sM.Select(map[string]interface{}{sM.Field.F_bus_id: busId}, sM.Field.F_sm_id)
	if len(sMMaps) > 0 {
		smIds := functions.ArrayValue2Array(sM.Field.F_sm_id, sMMaps)
		ssM := new(models.ShopSmModel).Init() //子店
		smWhere := map[string]interface{}{ssM.Field.F_sm_id: []interface{}{"IN", smIds}}
		smNum := ssM.GetTotalNum(smWhere)
		smWhere[ssM.Field.F_under_time] = []interface{}{"between", []int64{firstMonth, lastMonth}}
		smMonthUnderNum := ssM.GetTotalNum(smWhere) //月下架量
		reply.AllItemXCardNum = reply.AllItemXCardNum + smNum
		reply.AllUnderItemXCardNum = reply.AllUnderItemXCardNum + smMonthUnderNum
	}

	//综合卡
	cM := new(models.CardModel).Init()
	cMMaps := cM.Select(map[string]interface{}{cM.Field.F_bus_id: busId}, cM.Field.F_card_id)
	if len(cMMaps) > 0 {
		cardIds := functions.ArrayValue2Array(cM.Field.F_card_id, cMMaps)
		scM := new(models.ShopCardModel).Init() //子店
		scardWhere := map[string]interface{}{scM.Field.F_card_id: []interface{}{"IN", cardIds}}
		scardNum := scM.GetTotalNum(scardWhere)
		scardWhere[scM.Field.F_under_time] = []interface{}{"between", []int64{firstMonth, lastMonth}}
		scardMothUnderNum := scM.GetTotalNum(scardWhere) //月下架量
		reply.AllItemXCardNum = reply.AllItemXCardNum + scardNum
		reply.AllUnderItemXCardNum = reply.AllUnderItemXCardNum + scardMothUnderNum
	}

	//限时卡
	hM := new(models.HcardModel).Init()
	hMMaps := hM.Select(map[string]interface{}{hM.Field.F_bus_id: busId}, hM.Field.F_hcard_id)
	if len(hMMaps) > 0 {
		hcardIds := functions.ArrayValue2Array(hM.Field.F_hcard_id, hMMaps)
		shM := new(models.ShopHcardModel).Init() //子店
		shcardWhere := map[string]interface{}{shM.Field.F_hcard_id: []interface{}{"IN", hcardIds}}
		shcardNum := shM.GetTotalNum(shcardWhere)
		shcardWhere[shM.Field.F_under_time] = []interface{}{"between", []int64{firstMonth, lastMonth}}
		shcardMothUnderNum := shM.GetTotalNum(shcardWhere) //月下架量
		reply.AllItemXCardNum = reply.AllItemXCardNum + shcardNum
		reply.AllUnderItemXCardNum = reply.AllUnderItemXCardNum + shcardMothUnderNum
	}

	//限次卡
	hcM := new(models.NCardModel).Init()
	hcMMaps := hcM.Select(map[string]interface{}{hcM.Field.F_bus_id: busId}, hcM.Field.F_ncard_id)
	if len(hcMMaps) > 0 {
		ncardIds := functions.ArrayValue2Array(hcM.Field.F_ncard_id, hcMMaps)
		sncM := new(models.ShopNCardModel).Init() //子店
		sncardWhere := map[string]interface{}{sncM.Field.F_ncard_id: []interface{}{"IN", ncardIds}}
		sncardNum := sncM.GetTotalNum(sncardWhere)
		sncardWhere[sncM.Field.F_under_time] = []interface{}{"between", []int64{firstMonth, lastMonth}}
		sncardMonthUnderNum := sncM.GetTotalNum(sncardWhere) //月下架量
		reply.AllItemXCardNum = reply.AllItemXCardNum + sncardNum
		reply.AllUnderItemXCardNum = reply.AllUnderItemXCardNum + sncardMonthUnderNum
	}

	//限时限次卡
	hncM := new(models.HNCardModel).Init()
	hncMMaps := hncM.Select(map[string]interface{}{hncM.Field.F_bus_id: busId}, hncM.Field.F_hncard_id)
	if len(hncMMaps) > 0 {
		hncardIds := functions.ArrayValue2Array(hncM.Field.F_hncard_id, hncMMaps)
		shncM := new(models.ShopHNCardModel).Init() //子店
		shncMWhere := map[string]interface{}{shncM.Field.F_hncard_id: []interface{}{"IN", hncardIds}}
		shncardNum := shncM.GetTotalNum(shncMWhere)
		shncMWhere[shncM.Field.F_under_time] = []interface{}{"between", []int64{firstMonth, lastMonth}}
		shncardMonthUnderNum := shncM.GetTotalNum(shncMWhere) //月下架量
		reply.AllItemXCardNum = reply.AllItemXCardNum + shncardNum
		reply.AllUnderItemXCardNum = reply.AllUnderItemXCardNum + shncardMonthUnderNum
	}

	//充值卡
	rM := new(models.RcardModel).Init()
	rMMaps := rM.Select(map[string]interface{}{rM.Field.F_bus_id: busId}, rM.Field.F_rcard_id)
	if len(rMMaps) > 0 {
		rcardIds := functions.ArrayValue2Array(rM.Field.F_rcard_id, rMMaps)
		srM := new(models.ShopRcardModel).Init() //子店
		srMWhere := map[string]interface{}{srM.Field.F_rcard_id: []interface{}{"IN", rcardIds}}
		srcardNum := srM.GetTotalNum(srMWhere)
		srMWhere[srM.Field.F_under_time] = []interface{}{"between", []int64{firstMonth, lastMonth}}
		srcardMonthUnderNum := srM.GetTotalNum(srMWhere) //月下架量
		reply.AllItemXCardNum = reply.AllItemXCardNum + srcardNum
		reply.AllUnderItemXCardNum = reply.AllUnderItemXCardNum + srcardMonthUnderNum
	}

	//身份卡
	iM := new(models.IcardModel).Init()
	iMaps := iM.Select(map[string]interface{}{iM.Field.F_bus_id: busId}, iM.Field.F_icard_id)
	if len(iMaps) > 0 {
		icardIds := functions.ArrayValue2Array(iM.Field.F_icard_id, iMaps)
		siM := new(models.ShopIcardModel).Init()
		siMWhere := map[string]interface{}{siM.Field.F_icard_id: []interface{}{"IN", icardIds}}
		sicardNum := siM.GetTotalNum(siMWhere)
		siMWhere[siM.Field.F_under_time] = []interface{}{"between", []int64{firstMonth, lastMonth}}
		sicardMonthUnderNum := siM.GetTotalNum(siMWhere) //月下架量
		reply.AllItemXCardNum = reply.AllItemXCardNum + sicardNum
		reply.AllUnderItemXCardNum = reply.AllUnderItemXCardNum + sicardMonthUnderNum
	}

	return
}

//预付卡包含的单项目
func (i *ItemLogic) GetItemIncludeSingles(ctx context.Context, args *cards.ArgsGetCardProductsSinglesInfo, reply *cards.ReplyGetItemIncludeSingles) (err error) {
	start, limit := args.GetStart(), args.GetPageSize()
	shopId, itemId := args.ShopId, args.ItemId
	reply.Lists = make([]cards.IncSingleDetail2, 0)
	// 只取上架的数据
	switch args.ItemType {
	case cards.ITEM_TYPE_sm: //套餐
		sm := new(models.SmModel).Init()
		smInfo := sm.GetBySmid(itemId)
		if len(smInfo) == 0 {
			err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
			return
		}
		//获取包含的套餐和赠送项目信息
		mSmSingle := new(models.SmSingleModel).Init()
		smSingles := mSmSingle.SelectByPage(map[string]interface{}{mSmSingle.Field.F_sm_id: itemId}, start, limit)
		//reply.TotalNum = mSmSingle.Count([]base.WhereItem{{mSmSingle.Field.F_sm_id, itemId}})
		singleIds := functions.ArrayValue2Array(mSmSingle.Field.F_single_id, smSingles)
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, smSingles)
		reply.TotalNum = totalNum
		for _, single := range smSingles {
			if single[mSmSingle.Field.F_ssp_id].(string) != "0" {
				for index, allSingle := range allSingles {
					if single[mSmSingle.Field.F_ssp_id].(string) == strconv.Itoa(allSingle.SspId) {
						allSingles[index].Num, _ = strconv.Atoi(single[mSmSingle.Field.F_num].(string))
						reply.Lists = append(reply.Lists, allSingles[index])
						break
					}
				}
			} else {
				for index, allSingle := range allSingles {
					if single[mSmSingle.Field.F_single_id].(string) == strconv.Itoa(allSingle.SingleID) {
						allSingles[index].Num, _ = strconv.Atoi(single[mSmSingle.Field.F_num].(string))
						reply.Lists = append(reply.Lists, allSingles[index])
						break
					}
				}
			}
		}
	case cards.ITEM_TYPE_card: //综合卡
		ncm := new(models.CardModel).Init()
		cardInfo := ncm.GetByCardID(itemId)
		if len(cardInfo) == 0 {
			err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
			return
		}
		busId, _ := strconv.Atoi(cardInfo[ncm.Field.F_bus_id].(string))
		//获取包含的综合卡
		csm := new(models.CardSingleModel).Init()
		where := map[string]interface{}{csm.Field.F_card_id: itemId}
		cardSingles := csm.SelectByPage(where, start, limit)
		if len(cardSingles) == 0 {
			return
		}

		//reply.TotalNum = csm.GetTotalNum(where)
		//是否有适用全部的单项目
		where2 := where
		where2[csm.Field.F_single_id] = 0
		if len(csm.Find(where2)) > 0 { //适用全部
			cardSingles, reply.TotalNum, err = i.getBusOrShopSingleByPage(ctx, args.Paging, busId, shopId)
		}
		singleIds := functions.ArrayValue2Array(csm.Field.F_single_id, cardSingles)
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, cardSingles)
		reply.TotalNum = totalNum
		for _, single := range cardSingles {
			for index, allSingle := range allSingles {
				if single[csm.Field.F_single_id].(string) == strconv.Itoa(allSingle.SingleID) {
					reply.Lists = append(reply.Lists, allSingles[index])
					break
				}
			}
		}
	case cards.ITEM_TYPE_hcard: //限时卡
		hcardModel := new(models.HcardModel).Init()
		hcardInfo := hcardModel.GetHcardByID(itemId)
		if len(hcardInfo) == 0 {
			err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
			return
		}
		// 获取包含的搭配项目和赠送项目
		hcardSingleModel := new(models.HcardSingleModel).Init()
		where := map[string]interface{}{hcardSingleModel.Field.F_hcard_id: itemId}
		hcardSingles := hcardSingleModel.SelectByPage(where, start, limit)
		if len(hcardSingles) == 0 {
			return
		}
		//reply.TotalNum = hcardSingleModel.GetTotalNum(where)
		singleIds := functions.ArrayValue2Array(hcardSingleModel.Field.F_single_id, hcardSingles) // 所有搭配的项目ids
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, hcardSingles)
		reply.TotalNum = totalNum
		for _, single := range hcardSingles {
			for index, allSingle := range allSingles {
				if single[hcardSingleModel.Field.F_single_id].(string) == strconv.Itoa(allSingle.SingleID) {
					reply.Lists = append(reply.Lists, allSingles[index])
					break
				}
			}
		}
	case cards.ITEM_TYPE_hncard: //限时限次卡
		hNcm := new(models.HNCardModel).Init()
		hHNCardInfo := hNcm.GetByHNCardID(itemId)
		if len(hHNCardInfo) == 0 {
			err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
			return
		}
		//获取包含的套餐和赠送项目信息
		hncardSingleModel := new(models.HNCardSingleModel).Init()
		hncardSingles := hncardSingleModel.SelectByPage(map[string]interface{}{hncardSingleModel.Field.F_hncard_id: itemId}, start, limit)
		if len(hncardSingles) == 0 {
			return
		}
		//reply.TotalNum = hncardSingleModel.Count([]base.WhereItem{{hncardSingleModel.Field.F_hncard_id, itemId}})
		singleIds := functions.ArrayValue2Array(hncardSingleModel.Field.F_single_id, hncardSingles)
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, hncardSingles)
		reply.TotalNum = totalNum
		for _, single := range hncardSingles {
			if single[hncardSingleModel.Field.F_ssp_id].(string) != "0" {
				for index, allSingle := range allSingles {
					if single[hncardSingleModel.Field.F_ssp_id].(string) == strconv.Itoa(allSingle.SspId) {
						allSingles[index].Num, _ = strconv.Atoi(single[hncardSingleModel.Field.F_num].(string))
						reply.Lists = append(reply.Lists, allSingles[index])
						break
					}
				}
			} else {
				for index, allSingle := range allSingles {
					if single[hncardSingleModel.Field.F_single_id].(string) == strconv.Itoa(allSingle.SingleID) {
						allSingles[index].Num, _ = strconv.Atoi(single[hncardSingleModel.Field.F_num].(string))
						reply.Lists = append(reply.Lists, allSingles[index])
						break
					}
				}
			}
		}
	case cards.ITEM_TYPE_ncard: //限次卡
		ncm := new(models.NCardModel).Init()
		nCardInfo := ncm.GetByNCardID(itemId)
		if len(nCardInfo) == 0 {
			err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
			return
		}
		//获取包含的套餐和赠送项目信息
		ncardSingleModel := new(models.NCardSingleModel).Init()
		nCardSingles := ncardSingleModel.SelectByPage(map[string]interface{}{ncardSingleModel.Field.F_ncard_id: itemId}, start, limit)
		if len(nCardSingles) == 0 {
			return
		}
		//reply.TotalNum = ncardSingleModel.Count([]base.WhereItem{{ncardSingleModel.Field.F_ncard_id, itemId}})
		singleIds := functions.ArrayValue2Array(ncardSingleModel.Field.F_single_id, nCardSingles)
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, nCardSingles)
		reply.TotalNum = totalNum
		for _, single := range nCardSingles {
			if single[ncardSingleModel.Field.F_ssp_id].(string) != "0" {
				for index, allSingle := range allSingles {
					if single[ncardSingleModel.Field.F_ssp_id].(string) == strconv.Itoa(allSingle.SspId) {
						allSingles[index].Num, _ = strconv.Atoi(single[ncardSingleModel.Field.F_num].(string))
						reply.Lists = append(reply.Lists, allSingles[index])
						break
					}
				}
			} else {
				for index, allSingle := range allSingles {
					if single[ncardSingleModel.Field.F_single_id].(string) == strconv.Itoa(allSingle.SingleID) {
						allSingles[index].Num, _ = strconv.Atoi(single[ncardSingleModel.Field.F_num].(string))
						reply.Lists = append(reply.Lists, allSingles[index])
						break
					}
				}
			}
		}
	case cards.ITEM_TYPE_rcard: //充值卡
		mRcard := new(models.RcardModel).Init()
		rcardInfo := mRcard.GetByRcardId(itemId)
		if len(rcardInfo) == 0 {
			err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
			return
		}
		busId, _ := strconv.Atoi(rcardInfo[mRcard.Field.F_bus_id].(string))
		//获取包含的综合卡
		rsm := new(models.RcardSingleModel).Init()
		where := map[string]interface{}{mRcard.Field.F_rcard_id: itemId}
		rcardSingles := rsm.SelectByPage(where, start, limit)
		if len(rcardSingles) == 0 {
			return
		}
		reply.TotalNum = rsm.GetTotalNum(where)
		//是否有适用全部的单项目
		where2 := where
		where2[rsm.Field.F_single_id] = 0
		if len(rsm.Find(where2)) > 0 { //适用全部
			rcardSingles, reply.TotalNum, err = i.getBusOrShopSingleByPage(ctx, args.Paging, busId, shopId)
		}
		singleIds := functions.ArrayValue2Array(rsm.Field.F_single_id, rcardSingles)
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, rcardSingles)
		reply.TotalNum = totalNum
		for _, single := range rcardSingles {
			for index, allSingle := range allSingles {
				if single["single_id"].(string) == strconv.Itoa(allSingle.SingleID) {
					reply.Lists = append(reply.Lists, allSingles[index])
					break
				}
			}
		}
	case cards.ITEM_TYPE_icard: //身份卡
		icardModel := new(models.IcardModel).Init()
		icardInfo := icardModel.GetOneByPK(itemId)
		if len(icardInfo) == 0 {
			return toolLib.CreateKcErr(_const.ICARD_NOT_FOUND_ERROR)
		}
		busId, _ := strconv.Atoi(icardInfo["bus_id"].(string))
		icardSingleModel := new(models.IcardSingleModel).Init()
		where := map[string]interface{}{"icard_id": itemId}
		icardSingles := icardSingleModel.SelectByPage(where, start, limit)
		if len(icardSingles) == 0 {
			return
		}
		//reply.TotalNum = icardSingleModel.GetTotalNum(where)
		//是否有适用全部的单项目
		where2 := where
		where2["single_id"] = 0
		var discount float64
		allIcardSingel := icardSingleModel.Find(where2)
		if len(allIcardSingel) > 0 { //适用全部
			discount, _ = strconv.ParseFloat(allIcardSingel["discount"].(string), 64)
			icardSingles, reply.TotalNum, err = i.getBusOrShopSingleByPage(ctx, args.Paging, busId, shopId)
		}
		singleIds := functions.ArrayValue2Array("single_id", icardSingles)
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, icardSingles)
		reply.TotalNum = totalNum
		for _, single := range icardSingles {
			for index, allSingle := range allSingles {
				if discount > 0 { //全部
					allSingles[index].Discount = discount
				}
				if single["single_id"].(string) == strconv.Itoa(allSingle.SingleID) { //部分
					if single["discount"] != nil {
						allSingles[index].Discount, _ = strconv.ParseFloat(single["discount"].(string), 64)
					}
					reply.Lists = append(reply.Lists, allSingles[index])
				}
			}
		}
	}
	return
}

//预付卡赠送的单项目
func (i *ItemLogic) GetItemGiveSingles(ctx context.Context, args *cards.ArgsGetCardProductsSinglesInfo, reply *cards.ReplyGetItemGiveSingles) (err error) {
	start, limit := args.GetStart(), args.GetPageSize()
	shopId, itemId := args.ShopId, args.ItemId
	reply.Lists = make([]cards.IncSingleDetail2, 0)
	// 只取上架的数据
	switch args.ItemType {
	case cards.ITEM_TYPE_sm: //套餐
		sm := new(models.SmModel).Init()
		smInfo := sm.GetBySmid(itemId)
		if len(smInfo) == 0 {
			err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
			return
		}
		//获取包含的套餐和赠送项目信息
		mSmSingle := new(models.SmGiveModel).Init()
		where := []base.WhereItem{{mSmSingle.Field.F_sm_id, itemId}}
		smSingles := mSmSingle.SelectByPage(where, start, limit)
		singleIds := functions.ArrayValue2Array(mSmSingle.Field.F_single_id, smSingles)
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, smSingles)
		reply.TotalNum = totalNum
		for _, single := range smSingles {
			for index, allSingle := range allSingles {
				if single[mSmSingle.Field.F_single_id].(string) == strconv.Itoa(allSingle.SingleID) {
					allSingles[index].PeriodOfValidity, _ = strconv.Atoi(single[mSmSingle.Field.F_period_of_validity].(string))
					allSingles[index].Num, _ = strconv.Atoi(single[mSmSingle.Field.F_num].(string))
					reply.Lists = append(reply.Lists, allSingles[index])
					break
				}
			}
		}
	case cards.ITEM_TYPE_card: //综合卡
		ncm := new(models.CardModel).Init()
		cardInfo := ncm.GetByCardID(itemId)
		if len(cardInfo) == 0 {
			err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
			return
		}
		//获取包含的综合卡
		csm := new(models.CardGiveModel).Init()
		where := []base.WhereItem{{csm.Field.F_card_id, itemId}}
		cardSingles := csm.SelectByPage(where, start, limit)
		if len(cardSingles) == 0 {
			return
		}
		singleIds := functions.ArrayValue2Array(csm.Field.F_single_id, cardSingles)
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, cardSingles)
		reply.TotalNum = totalNum
		for _, single := range cardSingles {
			for index, allSingle := range allSingles {
				if single[csm.Field.F_single_id].(string) == strconv.Itoa(allSingle.SingleID) {
					allSingles[index].PeriodOfValidity, _ = strconv.Atoi(single[csm.Field.F_period_of_validity].(string))
					allSingles[index].Num, _ = strconv.Atoi(single[csm.Field.F_num].(string))
					reply.Lists = append(reply.Lists, allSingles[index])
					break
				}
			}
		}
	case cards.ITEM_TYPE_hcard: //限时卡
		hcardModel := new(models.HcardModel).Init()
		hcardInfo := hcardModel.GetHcardByID(itemId)
		if len(hcardInfo) == 0 {
			err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
			return
		}
		// 获取包含的搭配项目和赠送项目
		hcardSingleModel := new(models.HcardGiveModel).Init()
		where := []base.WhereItem{{hcardSingleModel.Field.F_hcard_id, itemId}}
		hcardSingles := hcardSingleModel.SelectByPage(where, start, limit)
		if len(hcardSingles) == 0 {
			return
		}
		singleIds := functions.ArrayValue2Array(hcardSingleModel.Field.F_single_id, hcardSingles) // 所有搭配的项目ids
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, hcardSingles)
		reply.TotalNum = totalNum
		for _, single := range hcardSingles {
			for index, allSingle := range allSingles {
				if single[hcardSingleModel.Field.F_single_id].(string) == strconv.Itoa(allSingle.SingleID) {
					allSingles[index].PeriodOfValidity, _ = strconv.Atoi(single[hcardSingleModel.Field.F_period_of_validity].(string))
					allSingles[index].Num, _ = strconv.Atoi(single[hcardSingleModel.Field.F_num].(string))
					reply.Lists = append(reply.Lists, allSingles[index])
					break
				}
			}
		}
	case cards.ITEM_TYPE_hncard: //限时限次卡
		hNcm := new(models.HNCardModel).Init()
		hHNCardInfo := hNcm.GetByHNCardID(itemId)
		if len(hHNCardInfo) == 0 {
			err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
			return
		}
		//获取包含的套餐和赠送项目信息
		hncardSingleModel := new(models.HNCardGiveModel).Init()
		where := []base.WhereItem{{hncardSingleModel.Field.F_hncard_id, itemId}}
		hncardSingles := hncardSingleModel.SelectByPage(where, start, limit)
		if len(hncardSingles) == 0 {
			return
		}
		singleIds := functions.ArrayValue2Array(hncardSingleModel.Field.F_single_id, hncardSingles)
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, hncardSingles)
		reply.TotalNum = totalNum
		for _, single := range hncardSingles {
			for index, allSingle := range allSingles {
				if single[hncardSingleModel.Field.F_single_id].(string) == strconv.Itoa(allSingle.SingleID) {
					allSingles[index].PeriodOfValidity, _ = strconv.Atoi(single[hncardSingleModel.Field.F_period_of_validity].(string))
					allSingles[index].Num, _ = strconv.Atoi(single[hncardSingleModel.Field.F_num].(string))
					reply.Lists = append(reply.Lists, allSingles[index])
					break
				}
			}
		}
	case cards.ITEM_TYPE_ncard: //限次卡
		ncm := new(models.NCardModel).Init()
		nCardInfo := ncm.GetByNCardID(itemId)
		if len(nCardInfo) == 0 {
			err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
			return
		}
		//获取包含的套餐和赠送项目信息
		ncardSingleModel := new(models.NCardGiveModel).Init()
		where := []base.WhereItem{{ncardSingleModel.Field.F_ncard_id, itemId}}
		nCardSingles := ncardSingleModel.SelectByPage(where, start, limit)
		if len(nCardSingles) == 0 {
			return
		}
		singleIds := functions.ArrayValue2Array(ncardSingleModel.Field.F_single_id, nCardSingles)
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, nCardSingles)
		reply.TotalNum = totalNum
		for _, single := range nCardSingles {
			for index, allSingle := range allSingles {
				if single[ncardSingleModel.Field.F_single_id].(string) == strconv.Itoa(allSingle.SingleID) {
					allSingles[index].PeriodOfValidity, _ = strconv.Atoi(single[ncardSingleModel.Field.F_period_of_validity].(string))
					allSingles[index].Num, _ = strconv.Atoi(single[ncardSingleModel.Field.F_num].(string))
					reply.Lists = append(reply.Lists, allSingles[index])
					break
				}
			}
		}
	case cards.ITEM_TYPE_rcard: //充值卡
		mRcard := new(models.RcardModel).Init()
		rcardInfo := mRcard.GetByRcardId(itemId)
		if len(rcardInfo) == 0 {
			err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
			return
		}
		//获取包含的综合卡
		rsm := new(models.RcardGiveModel).Init()
		where := []base.WhereItem{{mRcard.Field.F_rcard_id, itemId}}
		rcardSingles := rsm.SelectByPage(where, start, limit)
		if len(rcardSingles) == 0 {
			return
		}

		singleIds := functions.ArrayValue2Array(rsm.Field.F_single_id, rcardSingles)
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, rcardSingles)
		reply.TotalNum = totalNum
		for _, single := range rcardSingles {
			for index, allSingle := range allSingles {
				if single["single_id"].(string) == strconv.Itoa(allSingle.SingleID) {
					allSingles[index].PeriodOfValidity, _ = strconv.Atoi(single[rsm.Field.F_period_of_validity].(string))
					allSingles[index].Num, _ = strconv.Atoi(single[rsm.Field.F_num].(string))
					reply.Lists = append(reply.Lists, allSingles[index])
					break
				}
			}
		}
	case cards.ITEM_TYPE_icard: //身份卡
		icardModel := new(models.IcardModel).Init()
		icardInfo := icardModel.GetOneByPK(itemId)
		if len(icardInfo) == 0 {
			return toolLib.CreateKcErr(_const.ICARD_NOT_FOUND_ERROR)
		}
		icardSingleModel := new(models.IcardGiveModel).Init()
		where := []base.WhereItem{{icardSingleModel.Field.F_icard_id, itemId}}
		icardSingles := icardSingleModel.SelectByPage(where, start, limit)
		if len(icardSingles) == 0 {
			return
		}
		singleIds := functions.ArrayValue2Array("single_id", icardSingles)
		allSingles, totalNum := getIncSingles2(ctx, shopId, args.ShopStatus, args.IsDel, singleIds, icardSingles)
		reply.TotalNum = totalNum
		for _, single := range icardSingles {
			for index, allSingle := range allSingles {
				if single["single_id"].(string) == strconv.Itoa(allSingle.SingleID) { //部分
					allSingles[index].Num, _ = strconv.Atoi(single[icardSingleModel.Field.F_num].(string))
					reply.Lists = append(reply.Lists, allSingles[index])
				}
			}
		}
	}
	return
}

//获取总店或字段的单项目列表
func (i *ItemLogic) getBusOrShopSingleByPage(ctx context.Context, paging common.Paging, busId, shopId int) (singles []map[string]interface{}, totalNum int, err error) {
	rpcSingle := new(cardsClient.Single).Init()
	defer rpcSingle.Close()
	start, limit := paging.GetStart(), paging.GetPageSize()
	if shopId == 0 { //查询总店上架的单项目
		mSingle := new(models.SingleModel).Init()
		//singles = mSingle.GetPageByBusId(busId, start, limit, cards.IS_GROUND_yes)
		singles = mSingle.GetPageByBusId(busId, start, limit, "0", cards.IS_DEL_NO)
		//totalNum = mSingle.GetNumByBusId(busId, cards.IS_GROUND_yes)
		totalNum = mSingle.GetNumByBusId(busId, "0", cards.IS_DEL_NO)
	} else { //查询门店上架的单项目
		mShopSingle := new(models.ShopSingleModel).Init()
		singles = mShopSingle.GetByShopid(shopId, start, limit, fmt.Sprint(cards.STATUS_ON_SALE), "0", []int{})
		totalNum = mShopSingle.GetNumByShopid(shopId, fmt.Sprint(cards.STATUS_ON_SALE), "0", []int{})
	}
	return
}

//新增卡片时风控统计
func (i *ItemLogic) AddXCardTask(ctx context.Context, cardId, cardType int) {
	args := &task1.ArgsAddDelGoods{
		Ids:      []int{cardId},
		CardType: cardType,
		Type:     task1.DEL_ADD_IS_CRADS,
	}
	publicRpc := new(task2.Public).Init()
	defer publicRpc.Close()
	reply := false
	err := publicRpc.AddDelGoods(ctx, args, &reply)
	if err != nil {
		logs.Error("rpcCard AddXCardTask 方法调用RMQ错误: %s\n", err.Error())
		return
	}
}

// 综合卡、充值卡、身份卡 商品详情
func (i *ItemLogic) GetCardProductsInfo(ctx context.Context, args *cards.ArgsGetCardProductsSinglesInfo, reply *cards.ReplyGetCardProductsInfo) (err error) {
	mCardGoods := new(models.CardGoodsModel).Init()
	mRcardGoods := new(models.RcardGoodsModel).Init()
	mIcardGoods := new(models.IcardGoodsModel).Init()
	reply.Lists = make([]cards.IncProductDetail2, 0)

	page, pageSize := args.GetStart(), args.GetPageSize()
	var busId int
	var productsInfo []map[string]interface{}
	var productIDs []int
	var isAllProduct bool
	var discount float64
	//类型判断
	switch args.ItemType {
	case cards.ITEM_TYPE_card: //综合卡
		firstProduct := mCardGoods.Find(map[string]interface{}{mCardGoods.Field.F_card_id: args.ItemId})
		if productId, ok := firstProduct[mCardGoods.Field.F_product_id]; ok {
			if productId == "0" {
				isAllProduct = true
				cardModel := new(models.CardModel).Init()
				busId, _ = strconv.Atoi(cardModel.GetByCardID(args.ItemId, cardModel.Field.F_bus_id)[cardModel.Field.F_bus_id].(string))
				break
			}
		}
		productsInfo = mCardGoods.GetByCardIdPage(args.ItemId, page, pageSize)
		if len(productsInfo) > 0 {
			productIDs = functions.ArrayValue2Array(mCardGoods.Field.F_product_id, productsInfo)
			reply.TotalNum = mCardGoods.CountByCardId(args.ItemId)
		}
	case cards.ITEM_TYPE_rcard: //充值卡
		firstProduct := mRcardGoods.Find(map[string]interface{}{mRcardGoods.Field.F_rcard_id: args.ItemId})
		if productId, ok := firstProduct[mRcardGoods.Field.F_product_id]; ok {
			if productId == "0" {
				isAllProduct = true
				rcardModel := new(models.RcardModel).Init()
				busId, _ = strconv.Atoi(rcardModel.GetByRcardId(args.ItemId, []string{rcardModel.Field.F_bus_id})[rcardModel.Field.F_bus_id].(string))
				break
			}
		}
		productsInfo = mRcardGoods.GetByRcardIdPage(args.ItemId, page, pageSize)
		if len(productsInfo) > 0 {
			productIDs = functions.ArrayValue2Array(mRcardGoods.Field.F_product_id, productsInfo)
			reply.TotalNum = mRcardGoods.CountByCardId(args.ItemId)
		}
	case cards.ITEM_TYPE_icard: //身份卡
		firstProduct := mIcardGoods.Model.Where(map[string]interface{}{mIcardGoods.Field.F_icard_id: args.ItemId}).Find()
		if productId, ok := firstProduct["goods_id"]; ok {
			if productId == "0" {
				isAllProduct = true
				icardModel := new(models.IcardModel).Init()
				icardMaps := icardModel.GetRcardsByIcardIds([]int{args.ItemId}, icardModel.Field.F_bus_id)
				busId, _ = strconv.Atoi(icardMaps[0][icardModel.Field.F_bus_id].(string))
				discount, _ = strconv.ParseFloat(firstProduct["discount"].(string), 64)
				break
			}
		}
		productsInfo = mIcardGoods.GetAll(models.Condition{
			Where: map[string]interface{}{
				"icard_id": args.ItemId,
			},
			Offset: page, Limit: pageSize,
			Order: "id",
		})
		if len(productsInfo) > 0 {
			for index, v := range productsInfo {
				productsInfo[index]["product_id"] = v["goods_id"]
			}
			productIDs = functions.ArrayValue2Array("goods_id", productsInfo)
			reply.TotalNum = mIcardGoods.GetTotalNum(map[string]interface{}{"icard_id": args.ItemId})
		}
	}
	//若是查询全部
	if isAllProduct {
		productRpc := new(product.Product).Init()
		defer productRpc.Close()
		proArgs := &product2.ArgsGetProductIds{
			BusId:  busId,
			ShopId: args.ShopId,
			Paging: args.Paging,
		}
		proReply := product2.ReplyGetProductIds{}
		if err = productRpc.GetProductIds(ctx, proArgs, &proReply); err != nil {
			return err
		}
		reply.TotalNum = proReply.TotalNum
		productIDs = proReply.ProductIds
	}
	var allProducts map[int]cards.IncProductDetail2
	if allProducts, err = getIncProducts2(ctx, productIDs); err != nil {
		return
	}
	if isAllProduct {
		for _, cardProduct := range allProducts {
			cardProduct.Discount = discount
			reply.Lists = append(reply.Lists, cardProduct)
		}
	} else {
		for _, product := range productsInfo {
			for goodsId, cardProduct := range allProducts {
				if discount > 0 { //全部
					cardProduct.Discount = discount
				}
				if product["product_id"].(string) == strconv.Itoa(goodsId) { //部分
					if product["discount"] != nil {
						cardProduct.Discount, _ = strconv.ParseFloat(product["discount"].(string), 64)
					}
					reply.Lists = append(reply.Lists, cardProduct)
					break
				}
			}
		}
	}
	return nil
}

//获取预付卡包含和赠送的单项目合集 allSingle=true时，singleIds可以为空
func (i *ItemLogic) getItemCardIncSingleIds(itemsIds /*预付卡id*/ []int, itemType /*预付卡类型*/ int) (allSingle bool, singleIds []int, err error) {
	allSingle = false
	switch itemType {
	case cards.ITEM_TYPE_sm: //套餐
		smSingleM := new(models.SmSingleModel).Init()
		smSingleMaps := smSingleM.GetBySmids(itemsIds)
		if len(smSingleMaps) > 0 {
			singleIds = functions.ArrayValue2Array(smSingleM.Field.F_single_id, smSingleMaps)
		}
		smGiveM := new(models.SmGiveModel).Init()
		smGiveMaps := smGiveM.GetBySmids(itemsIds)
		if len(smGiveMaps) > 0 {
			singleIds = append(singleIds, functions.ArrayValue2Array(smGiveM.Field.F_single_id, smGiveMaps)...)
		}
	case cards.ITEM_TYPE_card: //综合卡
		cardSingleM := new(models.CardSingleModel).Init()
		cardSingleMaps := cardSingleM.GetByCardIds(itemsIds)
		for _, cardSingleMap := range cardSingleMaps {
			if cardSingleMap[cardSingleM.Field.F_single_id].(string) == "0" { //包含全部
				allSingle = true
				return
			}
		}
		if len(cardSingleMaps) > 0 {
			singleIds = functions.ArrayValue2Array(cardSingleM.Field.F_single_id, cardSingleMaps)
		}
		cardGiveM := new(models.CardGiveModel).Init()
		cardGiveMaps := cardGiveM.GetByCardIds(itemsIds)
		if len(cardGiveMaps) > 0 {
			singleIds = append(singleIds, functions.ArrayValue2Array(cardGiveM.Field.F_single_id, cardGiveMaps)...)
		}
	case cards.ITEM_TYPE_hcard: //限时卡
		hcardSingleM := new(models.HcardSingleModel).Init()
		hcardSingleMaps := hcardSingleM.GetByHcardIds(itemsIds)
		if len(hcardSingleMaps) > 0 {
			singleIds = functions.ArrayValue2Array(hcardSingleM.Field.F_single_id, hcardSingleMaps)
		}
		hcardGiveM := new(models.HcardGiveModel).Init()
		hcardGiveMaps := hcardGiveM.GetByHcardIds(itemsIds)
		if len(hcardGiveMaps) > 0 {
			singleIds = append(singleIds, functions.ArrayValue2Array(hcardGiveM.Field.F_single_id, hcardGiveMaps)...)
		}
	case cards.ITEM_TYPE_hncard: //限时限次卡
		hncardSingleM := new(models.HNCardSingleModel).Init()
		hncardSingleMaps := hncardSingleM.GetByHNCardIds(itemsIds)
		if len(hncardSingleMaps) > 0 {
			singleIds = functions.ArrayValue2Array(hncardSingleM.Field.F_single_id, hncardSingleMaps)
		}
		hncardGiveM := new(models.HNCardGiveModel).Init()
		hncardGiveMaps := hncardGiveM.GetByHNCardIds(itemsIds)
		if len(hncardGiveMaps) > 0 {
			singleIds = append(singleIds, functions.ArrayValue2Array(hncardGiveM.Field.F_single_id, hncardGiveMaps)...)
		}
	case cards.ITEM_TYPE_ncard: //限次卡
		ncardSingleM := new(models.NCardSingleModel).Init()
		ncardSingleMaps := ncardSingleM.GetByNCardIds(itemsIds)
		if len(ncardSingleMaps) > 0 {
			singleIds = functions.ArrayValue2Array(ncardSingleM.Field.F_single_id, ncardSingleMaps)
		}
		ncardGiveM := new(models.NCardGiveModel).Init()
		ncardGiveMaps := ncardGiveM.GetByNCardIds(itemsIds)
		if len(ncardGiveMaps) > 0 {
			singleIds = append(singleIds, functions.ArrayValue2Array(ncardGiveM.Field.F_single_id, ncardGiveMaps)...)
		}
	case cards.ITEM_TYPE_rcard: //充值卡
		rcardSingleM := new(models.RcardSingleModel).Init()
		rcardSingleMaps := rcardSingleM.GetByRcardids(itemsIds)
		for _, rcardSingleMap := range rcardSingleMaps {
			if rcardSingleMap[rcardSingleM.Field.F_single_id].(string) == "0" { //包含全部
				allSingle = true
				return
			}
		}
		if len(rcardSingleMaps) > 0 {
			singleIds = functions.ArrayValue2Array(rcardSingleM.Field.F_single_id, rcardSingleMaps)
		}
		rcardGiveM := new(models.RcardGiveModel).Init()
		rcardGiveMaps := rcardGiveM.GetByRcardids(itemsIds)
		if len(rcardGiveMaps) > 0 {
			singleIds = append(singleIds, functions.ArrayValue2Array(rcardGiveM.Field.F_single_id, rcardGiveMaps)...)
		}
	case cards.ITEM_TYPE_icard: //身份卡
		icardSingleM := new(models.IcardSingleModel).Init()
		icardSingleMaps := icardSingleM.GetAll(models.Condition{
			Where: map[string]interface{}{"single_id": []interface{}{"IN", itemsIds}},
		})
		for _, icardSingleMap := range icardSingleMaps {
			if icardSingleMap["single_id"].(string) == "0" { //包含全部
				allSingle = true
				return
			}
		}
		if len(icardSingleMaps) > 0 {
			singleIds = functions.ArrayValue2Array("single_id", icardSingleMaps)
		}
		icardGiveM := new(models.IcardGiveModel).Init()
		icardGiveMaps := icardGiveM.GetAll(models.Condition{
			Where: map[string]interface{}{"single_id": []interface{}{"IN", itemsIds}},
		})
		if len(icardGiveMaps) > 0 {
			singleIds = append(singleIds, functions.ArrayValue2Array("single_id", icardGiveMaps)...)
		}
	default:
		return
	}
	singleIds = functions.ArrayUniqueInt(singleIds)
	return
}

//校验卡项包含和赠送的单项目是否已经添加到门店中
func (i *ItemLogic) validShopSingleContainItemCardSingles(shopId, busId int, allSingle bool, singleIds []int) (err error) {
	shopSingleM := new(models.ShopSingleModel).Init()

	if allSingle { //卡项至少有一个是包含全部单项目
		//singleM:=new(models.SingleModel).Init()
		//singleMaps:=singleM.GetList(map[string]interface{}{singleM.Field.F_bus_id:busId},[]string{singleM.Field.F_single_id})
		//busSingleIds:=functions.ArrayValue2Array(singleM.Field.F_single_id,singleMaps)
		//shopSingleMaps:=shopSingleM.GetByShopIdAndSingleIds(shopId,busSingleIds)
		//if len(shopSingleMaps)<len(busSingleIds){
		//	return toolLib.CreateKcErr(_const.SHOP_SINGLE_NOT_CONTAIN_BUS_SINGLE)
		//}
		return
	}
	//卡项包含部分单项目
	if len(singleIds) == 0 {
		return
	}
	shopSingleMaps := shopSingleM.GetByShopIdAndSingleIds(shopId, singleIds)
	if len(shopSingleMaps) < len(singleIds) {
		return toolLib.CreateKcErr(_const.SHOP_SINGLE_NOT_CONTAIN_BUS_SINGLE)
	}
	return
}

//添加卡项时查询店铺下是否包含项目
func (i *ItemLogic) CheckProductInShop(ctx context.Context, busId, shopId, ItemType int, ItemIds []int) (err error, reply bool) {
	err, isAll, productIds := i.GetItemIncProductsId(ctx, ItemType, ItemIds)
	if err != nil {
		return err, false
	}
	if isAll {
		reply = true
		//productRpc := new(product.Product).Init()
		//defer productRpc.Close()
		//proArgs := &product2.ArgsIsShopProductEqBus{
		//	BusId: busId,
		//	ShopId: shopId,
		//}
		//err = productRpc.IsShopProductEqBus(ctx,proArgs,&reply)
		return
	} else {
		productRpc := new(product.Product).Init()
		defer productRpc.Close()
		proArgs := &product2.ArgsIsShopIncProducts{
			ShopId:     shopId,
			ProductIds: productIds,
		}
		err = productRpc.IsShopIncProducts(ctx, proArgs, &reply)
		return
	}
}

//根据卡项id和类型获取包含商品id
func (i *ItemLogic) GetItemIncProductsId(ctx context.Context, ItemType int, ItemIds []int) (err error, isAll bool, productIds []int) {
	switch ItemType {
	case cards.ITEM_TYPE_card:
		mcardGoods := new(models.CardGoodsModel).Init()
		goodsDataMap := mcardGoods.GetByCardIds(ItemIds)
		for _, v := range goodsDataMap {
			productId, _ := strconv.Atoi(v[mcardGoods.Field.F_product_id].(string))
			if productId == 0 {
				isAll = true
				return
			}
			productIds = append(productIds, productId)
		}
	case cards.ITEM_TYPE_rcard:
		mRcardGoods := new(models.RcardGoodsModel).Init()
		goodsDataMap := mRcardGoods.GetByRcardids(ItemIds)
		for _, v := range goodsDataMap {
			productId, _ := strconv.Atoi(v[mRcardGoods.Field.F_product_id].(string))
			if productId == 0 {
				isAll = true
				return
			}
			productIds = append(productIds, productId)
		}
	case cards.ITEM_TYPE_icard:
		mIcardGoods := new(models.IcardGoodsModel).Init()
		goodsDataMap := mIcardGoods.GetAll(models.Condition{
			Where: map[string]interface{}{
				"icard_id": []interface{}{"IN", ItemIds},
			},
		})
		for _, v := range goodsDataMap {
			productId, _ := strconv.Atoi(v["goods_id"].(string))
			if productId == 0 {
				isAll = true
				return
			}
			productIds = append(productIds, productId)
		}
	}
	productIds = functions.ArrayUniqueInt(productIds)
	return
}

//预付卡默认图片
func (i *ItemLogic) GetItemDefaultImgs(ctx context.Context, args *cards.ArgsGetItemDefaultImgs, reply *cards.ReplyGetItemDefaultImgs) (err error) {
	reply.CardsSmallDefaultPics, reply.CardsDefaultPics = map[int]string{}, map[int]string{}
	imgType, itemType := args.ImgType, args.ItemType
	if imgType == 1 {
		reply.CardsSmallDefaultPics = constkey.CardsSmallDefaultPics
		if itemType > 0 {
			reply.CardsSmallDefaultPics = map[int]string{itemType: constkey.CardsSmallDefaultPics[itemType]}
		}
	} else if imgType == 2 {
		reply.CardsDefaultPics = constkey.CardsDefaultPics
		if itemType > 0 {
			reply.CardsDefaultPics = map[int]string{itemType: constkey.CardsDefaultPics[itemType]}
		}
	} else {
		reply.CardsSmallDefaultPics = constkey.CardsSmallDefaultPics
		reply.CardsDefaultPics = constkey.CardsDefaultPics
		if itemType > 0 {
			reply.CardsSmallDefaultPics = map[int]string{itemType: constkey.CardsSmallDefaultPics[itemType]}
			reply.CardsDefaultPics = map[int]string{itemType: constkey.CardsDefaultPics[itemType]}
		}
	}
	return
}
