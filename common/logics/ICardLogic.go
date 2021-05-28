package logics

import (
	"context"
	"encoding/json"
	"fmt"
	"git.900sui.cn/kc/rpcCards/common/tools"
	"git.900sui.cn/kc/rpcCards/constkey"
	"git.900sui.cn/kc/rpcinterface/client/bus"
	"git.900sui.cn/kc/rpcinterface/client/product"
	bus2 "git.900sui.cn/kc/rpcinterface/interface/bus"
	product2 "git.900sui.cn/kc/rpcinterface/interface/product"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/kcgin/logs"
	"git.900sui.cn/kc/mapstructure"
	"git.900sui.cn/kc/rpcCards/common/models"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	order2 "git.900sui.cn/kc/rpcinterface/client/order"
	cardsTask "git.900sui.cn/kc/rpcinterface/client/task/cards"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/order"
	"github.com/smallnest/rpcx/log"
)

//ICardLogic 折扣卡管理（身份卡）
type ICardLogic struct {
	SaleShopCount int
}

//List 列表
func (s *ICardLogic) List(ctx context.Context, busID, shopId int, args cards.InputParamsICardList, reply *cards.OutputICardList) error {

	list := []cards.ICardBase{}
	status := args.Status
	status = args.Ground

	imgIDs, totalNum := s.getICardList(busID, shopId, args.Page, args.PageSize, status, []int{}, &list, args.IsDel, args.FilterShopHasAdd)

	reply.TotalNum = totalNum
	reply.Lists = list
	reply.IndexImg = getImgsByImgIds(ctx, imgIDs, cards.ITEM_TYPE_icard)
	//子店已添加的卡项
	var shopAddCards []map[string]interface{}
	scModel := new(models.ShopSmModel).Init()
	if shopId > 0 {
		cardIds := make([]int, 0)
		for _, v := range list {
			cardIds = append(cardIds, v.IcardID)
		}
		shopAddCards = scModel.GetByShopidAdSmids(shopId, cardIds)
		for index, list := range reply.Lists {
			for _, shopCard := range shopAddCards {
				if strconv.Itoa(list.IcardID) == shopCard[scModel.Field.F_sm_id].(string) {
					reply.Lists[index].ShopItemId, _ = strconv.Atoi(shopCard[scModel.Field.F_id].(string))
					reply.Lists[index].IsSelfShop = 1
					reply.Lists[index].Status, _ = strconv.Atoi(shopCard[scModel.Field.F_status].(string))
					reply.Lists[index].ShopIsDel, _ = strconv.Atoi(shopCard[scModel.Field.F_id].(string))
					break
				}
			}
		}
	}

	return nil
}

func (s *ICardLogic) getICardList(busID, shopId, page, pageSize, ground int, icardIDs []int, list *[]cards.ICardBase, isDel string, filterShopHasAdd bool) (imgIDs []int, totalNum int) {
	where := make(map[string]interface{})
	if busID == 0 {
		return
	}
	where["bus_id"] = busID
	if len(icardIDs) != 0 {
		where["icard_id"] = []interface{}{"IN", icardIDs}
	}
	//TODO 默认只取未删除的数据
	isDel = strconv.Itoa(cards.ITEM_IS_DEL_NO)
	if isDel != "" {
		isDelInt, _ := strconv.Atoi(isDel)
		where["is_del"] = isDelInt
	}

	var whereGround []int
	if ground == 0 {
		whereGround = []int{
			cards.IS_GROUND_no,
			cards.IS_GROUND_yes,
		}
	} else {
		ground = ground - 1
		whereGround = []int{ground}
	}
	model := new(models.IcardModel).Init()
	where["is_ground"] = []interface{}{"IN", whereGround}
	if shopId > 0 && filterShopHasAdd {
		siMdel := new(models.ShopIcardModel).Init()
		siMaps := siMdel.GetAll(models.Condition{
			Where:  map[string]interface{}{siMdel.Field.F_shop_id: shopId,siMdel.Field.F_is_del: cards.IS_BUS_DEL_no},
			Offset: 0,
			Limit:  0,
			Order:  siMdel.Field.F_id,
		})
		if len(siMaps) > 0 {
			hasAddIcardIds := functions.ArrayValue2Array(siMdel.Field.F_icard_id, siMaps)
			where[model.Field.F_icard_id] = []interface{}{"NOT IN ", hasAddIcardIds}
		}
	}

	condition := models.Condition{
		Where:  where,
		Offset: (page - 1) * pageSize,
		Limit:  pageSize,
		Order:  "ctime desc",
	}
	data := model.GetPaginationData(condition, "icard_id", "is_del", "bus_id", "name", "real_price", "price", "service_period", "sales", "ctime", "img_id", "is_ground", "sale_shop_num")
	totalNum = model.GetCount(condition)

	//存储门店上架数管道
	total := len(data)
	//intChan := make(chan string, total)
	// wt.Add(total)

	stringKeyResult := make(map[int]*cards.ICardBase, total)

	var allIcardIds []int
	//shopIcardModel := new(models.ShopIcardModel).Init()
	for _, v := range data {
		imgID, _ := strconv.Atoi(v["img_id"].(string))
		isGround, _ := strconv.Atoi(v["is_ground"].(string))
		imgIDs = append(imgIDs, imgID)
		ctime64, _ := strconv.ParseInt(v["ctime"].(string), 10, 64)
		tm := time.Unix(ctime64, 0)
		ctime := tm.Format("2006-01-02 15:04:05")

		servicePeriod, _ := strconv.Atoi(v["service_period"].(string))
		realPrice, _ := strconv.ParseFloat(v["real_price"].(string), 64)
		icardID, _ := strconv.Atoi(v["icard_id"].(string))
		price, _ := strconv.ParseFloat(v["price"].(string), 64)
		sales, _ := strconv.Atoi(v["sales"].(string))
		saleShopNum, _ := strconv.Atoi(v["sale_shop_num"].(string))
		isDel, _ := strconv.Atoi(v["is_del"].(string))
		//多协程处理上架门店数
		//go getSaleShopCount(shopIcardModel, intChan, v["icard_id"])
		allIcardIds = append(allIcardIds, icardID)

		stringKeyResult[icardID] = &cards.ICardBase{
			BusID:         busID,
			Ctime:         ctime,
			Click:         0,
			SaleShopCount: 0,
			SaleShopNum:   saleShopNum,
			ServicePeriod: servicePeriod,
			RealPrice:     realPrice,
			IcardID:       icardID,
			Name:          v["name"].(string),
			Price:         price,
			Sales:         sales,
			ImgID:         imgID,
			IsGround:      isGround,
			IsDel:         isDel,
		}
	}

	//从管道中获取结果回写数据
	//go setSaleShopCount(intChan, &stringKeyResult)
	//wt.Wait()

	//sort 排序
	mlen := len(stringKeyResult)
	ss := make([]int, 0, mlen)
	for _, v := range stringKeyResult {
		ss = append(ss, v.IcardID)
	}
	sort.Ints(ss)
	sort.Sort(sort.Reverse(sort.IntSlice(ss)))

	for _, s := range ss {
		*list = append(*list, *stringKeyResult[s])
	}

	return
}

var wt sync.WaitGroup

func setSaleShopCount(c chan string, stringKeyResult *map[int]*cards.ICardBase) {
	for v := range c {
		strArr := strings.Split(v, "-")
		cardID, _ := strconv.Atoi(strArr[0])
		saleShopCount := strArr[1]
		if s, ok := (*stringKeyResult)[cardID]; ok {
			saleShopCountInt, _ := strconv.Atoi(saleShopCount)
			s.SaleShopCount = saleShopCountInt
		}
		wt.Done()
	}
}

func getSaleShopCount(shopIcardModel *models.ShopIcardModel, c chan string, icardID interface{}) {
	count := shopIcardModel.GetTotalNum(map[string]interface{}{
		shopIcardModel.Field.F_icard_id: icardID,
		shopIcardModel.Field.F_status:   cards.STATUS_ON_SALE,
	})

	str := icardID.(string) + "-" + strconv.Itoa(count)
	c <- str
}

//AddCardExt 添加ICard描述
func (i *ICardLogic) AddCardExt(icm *models.IcardModel, icardID int, notesString string) (err error) {
	var notes []cards.CardNote
	json.Unmarshal([]byte(notesString), &notes)
	if len(notes) > constkey.CardNotesMaxNum {
		return toolLib.CreateKcErr(_const.NOTES_NUM_MAX_10)
	}
	for _, note := range notes {
		if functions.Mb4Strlen(note.Notes) > constkey.CardNotesSimpleMaxLength {
			return toolLib.CreateKcErr(_const.NOTES_LEN_MAX_30)
		}
	}
	notesStr, _ := json.Marshal(notes)
	icardExtModel := new(models.IcardExtModel).Init(icm.Model.GetOrmer())
	icardExtModel.Delete(models.Condition{
		Where: map[string]interface{}{
			icardExtModel.Field.F_icard_id: icardID,
		}})
	if _, err = icardExtModel.Insert(map[string]interface{}{
		icardExtModel.Field.F_icard_id: icardID,
		icardExtModel.Field.F_notes:    string(notesStr)}); err != nil {
		err = toolLib.CreateKcErr(_const.DB_ERR)
	}
	return
}

//Save 总店添加/编辑身份卡
func (s *ICardLogic) Save(ctx context.Context, busID int, args cards.InputParamsICardSave, reply *cards.OutputParamsICardSave) (err error) {
	bindID := getBusMainBindId(ctx, busID)

	imgID, err := checkImg(ctx, args.ImgHash)
	if err != nil {
		return toolLib.CreateKcErr(_const.IMG_ERROR)
	}

	IcardModel := new(models.IcardModel).Init()
	IcardModel.Model.Begin()
	// 主表
	IcardData := map[string]interface{}{
		IcardModel.Field.F_bus_id:         busID,
		IcardModel.Field.F_bind_id:        bindID,
		IcardModel.Field.F_name:           args.Name,
		IcardModel.Field.F_sort_desc:      args.ShortDesc,
		IcardModel.Field.F_real_price:     args.RealPrice,
		IcardModel.Field.F_price:          args.Price,
		IcardModel.Field.F_service_period: args.ServicePeriod,
		IcardModel.Field.F_img_id:         imgID,
		IcardModel.Field.F_ctime:          time.Now().Unix(),
	}

	if args.IcardID != 0 {
		//校验卡项是否删除
		icardMap := IcardModel.GetByIcardId(args.IcardID, []string{IcardModel.Field.F_is_del})
		isDel, _ := strconv.Atoi(icardMap[IcardModel.Field.F_is_del].(string))
		if isDel == cards.ITEM_IS_DEL_YES {
			return toolLib.CreateKcErr(_const.ITEM_IS_DEL)
		}
		IcardData[IcardModel.Field.F_icard_id] = args.IcardID
	} else {
		IcardData[IcardModel.Field.F_is_ground] = cards.IS_GROUND_yes //发布身份卡时默认上架
		args.IsSync = cards.IS_SYNC_NO                                //第一次发卡默认是不同步折扣
	}
	icardID := IcardModel.CreateOrUpdate(IcardData)
	//卡项服务-商品
	if args.IncludeProducts != "" {
		var mapResult []map[string]interface{}
		if err := json.Unmarshal([]byte(args.IncludeProducts), &mapResult); err != nil {
			log.Info("IncludeProducts json string error:", err)
			IcardModel.Model.RollBack()
			return toolLib.CreateKcErr(_const.ICARD_INCLUDEPRODUCTS_ERROR)
		}
		icardGoodsModel := new(models.IcardGoodsModel).Init()
		var icardGoodsData []map[string]interface{}
		igbModel := new(models.IcardGoodsBackupModel).Init()
		var icardGoodsBackupData []map[string]interface{}
		repeatSave := make(map[float64]bool)
		for _, v := range mapResult {
			if _, ok := repeatSave[v["productID"].(float64)]; ok {
				continue
			}
			repeatSave[v["productID"].(float64)] = true
			icardGoodsData = append(icardGoodsData, map[string]interface{}{
				icardGoodsModel.Field.F_icard_id: icardID,
				icardGoodsModel.Field.F_goods_id: v["productID"],
				icardGoodsModel.Field.F_discount: v["discount"],
			})
			icardGoodsBackupData = append(icardGoodsBackupData, map[string]interface{}{
				igbModel.Field.F_icard_id:   icardID,
				igbModel.Field.F_is_sync:    args.IsSync,
				igbModel.Field.F_goods_id:   v["productID"],
				igbModel.Field.F_discount:   v["discount"],
				igbModel.Field.F_backup_num: igbModel.GetLastBackUumByIcardId(true, "", icardID),
			})
		}
		//备份商品折扣信息
		igbModel.InsertAll(icardGoodsBackupData)
		icardGoodsModel.Delete(models.Condition{
			Where: map[string]interface{}{
				icardGoodsModel.Field.F_icard_id: icardID,
			}})
		icardGoodsModel.InsertAll(icardGoodsData)
	}

	//卡项服务-单项目
	if args.IncludeSingles != "" {
		var mapResult []map[string]interface{}
		if err := json.Unmarshal([]byte(args.IncludeSingles), &mapResult); err != nil {
			log.Info("IncludeSingles json string error:", err)
			IcardModel.Model.RollBack()
			return toolLib.CreateKcErr(_const.ICARD_INCLUDESINGLES_ERROR)
		}

		icardSingleModel := new(models.IcardSingleModel).Init()
		var icardSingleData []map[string]interface{}
		//备份商品折扣信息
		isbModel := new(models.IcardSingleBackupModel).Init()
		var icardGoodsBackupData []map[string]interface{}
		repeatSave := make(map[float64]bool)
		for _, v := range mapResult {
			if _, ok := repeatSave[v["singleID"].(float64)]; ok {
				continue
			}
			repeatSave[v["singleID"].(float64)] = true
			icardSingleData = append(icardSingleData, map[string]interface{}{
				icardSingleModel.Field.F_icard_id:  icardID,
				icardSingleModel.Field.F_single_id: v["singleID"],
				icardSingleModel.Field.F_discount:  v["discount"],
			})
			icardGoodsBackupData = append(icardGoodsBackupData, map[string]interface{}{
				isbModel.Field.F_icard_id:   icardID,
				isbModel.Field.F_is_sync:    args.IsSync,
				isbModel.Field.F_single_id:  v["singleID"],
				isbModel.Field.F_discount:   v["discount"],
				isbModel.Field.F_backup_num: isbModel.GetLastBackUumByIcardId(true, "", icardID),
			})
		}
		//备份项目折扣信息
		isbModel.InsertAll(icardGoodsBackupData)
		icardSingleModel.Delete(models.Condition{
			Where: map[string]interface{}{
				icardSingleModel.Field.F_icard_id: icardID,
			}})
		icardSingleModel.InsertAll(icardSingleData)
	}

	//添加NCardExt
	if err = s.AddCardExt(IcardModel, icardID, args.Notes); err != nil {
		IcardModel.Model.RollBack()
		return toolLib.CreateKcErr(_const.ICARD_EXT_ERROR)
	}

	IcardModel.Model.Commit()
	reply.IcardID = icardID

	//添加风控统计任务
	if args.IcardID == 0 {
		new(ItemLogic).AddXCardTask(ctx, icardID, cards.ITEM_TYPE_icard)
	}
	//是否同步身份卡折扣
	if args.IsSync == cards.IS_SYNC_YES {
		s.SetIcardDiscountTask(ctx, icardID)
	}
	return
}

//Info 详情
func (s *ICardLogic) Info(ctx context.Context, args cards.InputParamsICardInfo, reply *cards.OutputParamsICardInfo) (err error) {
	icardID := args.IcardID
	shopID := args.ShopID

	IcardModel := new(models.IcardModel).Init()
	IcardData := IcardModel.GetOneByPK(icardID)
	if len(IcardData) == 0 {
		return toolLib.CreateKcErr(_const.ICARD_NOT_FOUND_ERROR)
	}

	var imgHash, imgURL string
	if IcardData[IcardModel.Field.F_img_id] != nil {
		imgID, _ := strconv.Atoi(IcardData[IcardModel.Field.F_img_id].(string))
		imgHash, imgURL = getImg(ctx, imgID, cards.ITEM_TYPE_icard)
	}

	//服务-单项目
	icardSingleModel := new(models.IcardSingleModel).Init()
	icardSingleData := icardSingleModel.GetAll(models.Condition{
		Where: map[string]interface{}{
			IcardModel.Field.F_icard_id: icardID,
		},
		Limit: 1,
	})

	//服务-商品
	icardGoodsModel := new(models.IcardGoodsModel).Init()
	icardGoodsData := icardGoodsModel.GetAll(models.Condition{
		Where: map[string]interface{}{
			IcardModel.Field.F_icard_id: icardID,
		},
		Limit: 1,
	})

	discountSingle := 0.0
	var isAllSingle, isAllGoods bool
	if len(icardSingleData) > 0 {
		discountSingle, _ = strconv.ParseFloat(icardSingleData[0][icardSingleModel.Field.F_discount].(string), 64)
		if singleID, _ := strconv.Atoi(icardSingleData[0][icardSingleModel.Field.F_single_id].(string)); singleID == 0 {
			isAllSingle = true
		}
	}

	discountGoods := 0.0
	if len(icardGoodsData) > 0 {
		discountGoods, _ = strconv.ParseFloat(icardGoodsData[0][icardGoodsModel.Field.F_discount].(string), 64)
		if goodsID, _ := strconv.Atoi(icardGoodsData[0][icardGoodsModel.Field.F_goods_id].(string)); goodsID == 0 {
			isAllGoods = true
		}
	}

	// 获取折扣，根据身份卡包含的单项目id，0=适用所有项目
	// 适用所有，折扣返回一个固定值
	// 适用部分项目，折扣返回一个区间值[min,max]
	icModel := new(models.IcardSingleModel).Init()
	discountTemp := make([]float64, 0)
	icardLists := icModel.GetByIcardIds([]int{icardID})
	for _, icardList := range icardLists {
		discount, _ := strconv.ParseFloat(icardList[icModel.Field.F_discount].(string), 64)
		discountTemp = append(discountTemp, discount)
	}

	sort.Float64s(discountTemp)
	min, max := discountTemp[0], discountTemp[len(discountTemp)-1]
	replyDiscount := make([]float64, 0)
	if isAllSingle == true {
		replyDiscount = append(replyDiscount, min)
	} else {
		replyDiscount = append(replyDiscount, min)
		replyDiscount = append(replyDiscount, max)
	}
	reply.DiscountLists = replyDiscount

	//var icardSingleIds []int
	//discounts := make(map[int]interface{})
	//for _, v := range icardSingleData {
	//	singleID, _ := strconv.Atoi(v["single_id"].(string))
	//	icardSingleIds = append(icardSingleIds, singleID)
	//	discounts[singleID] = v["discount"]
	//}

	//getIncInfSinglesData := getIncInfSingles(ctx, icardSingleIds)
	//var icardSingle []cards.ICardInfoIcardSingle
	//for _, v := range getIncInfSinglesData {
	//	price, _ := strconv.ParseFloat(v.Price, 64)
	//	discount := 0.0
	//	if discounts[v.SingleID] != nil {
	//		discount, _ = strconv.ParseFloat(discounts[v.SingleID].(string), 64)
	//	}
	//
	//	icardSingle = append(icardSingle, cards.ICardInfoIcardSingle{
	//		SingleID:    v.SingleID,
	//		Name:        v.Name,
	//		Price:       price,
	//		Discount:    discount,
	//		ImgID:       v.ImgId,
	//		ImgURL:      v.ImgUrl,
	//		RealPrice:   v.RealPrice,
	//		ServiceTime: v.ServiceTime,
	//	})
	//}

	//服务-商品
	// icardGoodsModel := new(models.IcardGoodsModel).Init()
	// icardGoodsData := icardGoodsModel.GetAll(models.Condition{
	// 	Where: map[string]interface{}{
	// 		"icard_id": icardID,
	// 	},
	// })

	// var icardGoodsIds []int
	// goodDiscounts := make(map[int]interface{})
	// for _, v := range icardGoodsData {
	// 	goodsStr := v["goods_id"].(string)
	// 	goodsID, _ := strconv.Atoi(goodsStr)
	// 	icardGoodsIds = append(icardGoodsIds, goodsID)
	// 	goodDiscounts[goodsID] = v["discount"]
	// }

	//包含的商品
	//goodsData, _ := getIncProducts(ctx, icardGoodsIds)
	//var icardGoods []cards.ICardInfoIcardGoods
	//for _, v := range goodsData {
	//	discount := 0.0
	//	if goodDiscounts[v.ProductID] != nil {
	//		discount, _ = strconv.ParseFloat(goodDiscounts[v.ProductID].(string), 64)
	//	}
	//
	//	icardGoods = append(icardGoods, cards.ICardInfoIcardGoods{
	//		GoodsID:  v.ProductID,
	//		Name:     v.Name,
	//		Price:    v.SpecPrice,
	//		ImgURL:   v.ImgUrl,
	//		Discount: discount,
	//		ImgID:    v.ImgId,
	//	})
	//}
	//
	//fmt.Println(icardGoods)

	// 获取身份卡门店添加详情  15 -- []int {3,4,6}
	busId, _ := strconv.Atoi(IcardData[IcardModel.Field.F_bus_id].(string))
	iCardShopModel := new(models.IcardShopModel).Init()
	iCardShopLists := iCardShopModel.GetByIcardIdAndBusId(icardID, busId)

	iCardShopIds := make([]int, 0)
	for _, icInfoValue := range iCardShopLists {
		sshopId, _ := strconv.Atoi(icInfoValue[iCardShopModel.Field.F_shop_id].(string))
		iCardShopIds = append(iCardShopIds, sshopId)
	}
	var replyShop []bus2.ReplyShopName
	rLists := make([]cards.ReplyShopName, 0) // 临时存放返回数据
	rpcBus := new(bus.Shop).Init()
	defer rpcBus.Close()
	err = rpcBus.GetShopNameByShopIds(ctx, &iCardShopIds, &replyShop)
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

	realPrice, _ := strconv.ParseFloat(IcardData[IcardModel.Field.F_real_price].(string), 64)
	price, _ := strconv.ParseFloat(IcardData[IcardModel.Field.F_price].(string), 64)
	servicePeriod, _ := strconv.Atoi(IcardData[IcardModel.Field.F_service_period].(string))
	hasGiveSignle, _ := strconv.Atoi(IcardData[IcardModel.Field.F_has_give_signle].(string))
	isGround, _ := strconv.Atoi(IcardData[IcardModel.Field.F_is_ground].(string))

	//商户信息
	if err = getBusInfo(ctx, busId, &reply.BusInfo); err != nil {
		err = toolLib.CreateKcErr(_const.SHOP_INFO_ERR)
		return
	}

	//已添加到门店的身份卡
	shopIcardModel := new(models.ShopIcardModel).Init()
	shopIcardData := shopIcardModel.GetByShopidAndRcardid(shopID, icardID)

	var ssid int
	if len(shopIcardData) != 0 {
		ssid, _ = strconv.Atoi(shopIcardData[shopIcardModel.Field.F_id].(string))
	}

	reply.ICardBase = cards.ICardBase{
		IcardID:        icardID,
		Name:           IcardData[IcardModel.Field.F_name].(string),
		ShortDesc:      IcardData[IcardModel.Field.F_sort_desc].(string),
		RealPrice:      realPrice,
		Price:          price,
		ServicePeriod:  servicePeriod,
		HasGiveSignle:  hasGiveSignle,
		IsGround:       isGround,
		SsID:           ssid,
		DiscountSingle: discountSingle,
		DiscountGoods:  discountGoods,
	}

	reply.IsAllProduct = isAllGoods
	reply.IsAllSingle = isAllSingle
	reply.ImgHash = imgHash
	reply.ImgURL = imgURL
	//reply.ICardInfoIcardSingle = icardSingle
	//reply.ICardInfoIcardGoods = icardGoods
	reply.ICardInfoIcardSingle = []cards.ICardInfoIcardSingle{}
	reply.ICardInfoIcardGoods = []cards.ICardInfoIcardGoods{}
	reply.ShareLink = tools.GetShareLink(icardID, shopID, cards.ITEM_TYPE_icard)

	return
}

//Delete 总店/分店删除身份卡
func (s *ICardLogic) Delete(ctx context.Context, busId, shopId int, icardIdsStr string) (reply bool, err error) {
	if icardIdsStr == "" {
		err = toolLib.CreateKcErr(_const.CARD_ID_IS_NIL)
		return
	}
	var icardIds []int
	if err = json.Unmarshal([]byte(icardIdsStr), &icardIds); err != nil {
		err = toolLib.CreateKcErr(_const.ICARD_PARAMS_ERROR)
		return
	}
	shopIcardModel := new(models.ShopIcardModel).Init()
	shopIcardModelWhere := map[string]interface{}{
		shopIcardModel.Field.F_icard_id: []interface{}{"IN", icardIds},
		shopIcardModel.Field.F_is_del:   cards.ITEM_IS_DEL_NO,
	}
	if shopId == 0 {
		//总店身份卡-软删除
		icardModel := new(models.IcardModel).Init()
		icardModel.UpdateAll(map[string]interface{}{
			icardModel.Field.F_is_del:       cards.ITEM_IS_DEL_YES,
			shopIcardModel.Field.F_del_time: time.Now().Unix(),
		}, models.Condition{
			Where: map[string]interface{}{
				icardModel.Field.F_icard_id: []interface{}{"IN", icardIds},
				icardModel.Field.F_is_del:   cards.ITEM_IS_DEL_NO,
				icardModel.Field.F_bus_id:   busId,
			},
			Offset: 0,
			Limit:  0,
			Order:  "",
		})
	} else {
		shopIcardModelWhere[shopIcardModel.Field.F_shop_id] = shopId
	}
	//分店身份卡-软删除
	shopIcardModel.UpdateAll(map[string]interface{}{
		shopIcardModel.Field.F_is_del:   cards.ITEM_IS_DEL_YES,
		shopIcardModel.Field.F_del_time: time.Now().Unix(),
	}, shopIcardModelWhere)

	//同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(icardIds, cards.ITEM_TYPE_icard, shopId) {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	reply = true
	return
}

//Push 推送门店（废用）
func (s *ICardLogic) Push(ctx context.Context, args cards.InputParamsICardPush, reply *cards.OutputParamsICardPush) (err error) {
	busID, _ := args.BsToken.GetBusId()

	if args.IcardIds == "" || args.ShopIds == "" {
		return toolLib.CreateKcErr(_const.ICARD_PARAMS_ERROR)
	}

	var icardIds, shopIds []int
	err1 := json.Unmarshal([]byte(args.IcardIds), &icardIds)
	err2 := json.Unmarshal([]byte(args.ShopIds), &shopIds)
	if err1 != nil || err2 != nil {
		return toolLib.CreateKcErr(_const.ICARD_PARAMS_ERROR)
	}
	//过滤掉删除的id
	icardModel := new(models.IcardModel).Init()
	icardMaps := icardModel.GetAll(models.Condition{
		Where: map[string]interface{}{
			icardModel.Field.F_icard_id: []interface{}{"IN", icardIds},
			icardModel.Field.F_is_del:   cards.ITEM_IS_DEL_NO,
		},
		Offset: 0,
		Limit:  0,
		Order:  "",
	})
	if len(icardMaps) == 0 {
		return toolLib.CreateKcErr(_const.ITEM_IS_DEL)
	}
	icardIds = functions.ArrayValue2Array(icardModel.Field.F_icard_id, icardMaps)
	var data []map[string]interface{}
	icardShopModel := new(models.IcardShopModel).Init()
	for _, icardID := range icardIds {
		for _, shopID := range shopIds {
			data = append(data, map[string]interface{}{
				icardShopModel.Field.F_shop_id:  shopID,
				icardShopModel.Field.F_bus_id:   busID,
				icardShopModel.Field.F_icard_id: icardID,
			})
		}
	}

	icardShopModel.Delete(models.Condition{
		Where: map[string]interface{}{
			icardShopModel.Field.F_icard_id: []interface{}{"IN", icardIds},
		}})
	icardShopModel.InsertAll(data)

	reply.IcardIds = icardIds
	return
}

//OurShopList 添加到门店的身份卡列表
func (s *ICardLogic) OurShopList(ctx context.Context, args cards.InputParamsICardList, reply *cards.OutputICardList) error {
	busID, _ := args.BsToken.GetBusId()
	page := args.Page
	pageSize := args.PageSize

	shopID := args.ShopID
	status := args.Status
	ground := args.Ground

	if shopID == 0 {
		return toolLib.CreateKcErr(_const.ICARD_PARAMS_ERROR)
	}

	var whereStatus []int
	if status == 0 {
		whereStatus = []int{
			cards.STATUS_OFF_SALE,
			cards.STATUS_ON_SALE,
			cards.STATUS_DISABLE,
		}
	} else {
		whereStatus = []int{status}
	}

	shopIcardModel := new(models.ShopIcardModel).Init()
	shopIcardData := shopIcardModel.GetAll(models.Condition{
		Where: map[string]interface{}{
			shopIcardModel.Field.F_shop_id: shopID,
			shopIcardModel.Field.F_status:  []interface{}{"IN", whereStatus},
			shopIcardModel.Field.F_is_del:  cards.ITEM_IS_DEL_NO, //暂时只取未删除的卡项
		},
		Offset: (page - 1) * pageSize,
		Limit:  pageSize,
	})

	if len(shopIcardData) == 0 {
		reply.Lists = []cards.ICardBase{}
		reply.IndexImg = map[int]string{}
		return nil
	}

	var icardIds []int
	shopStatusMap := make(map[int]int)
	shopItemIdMap := make(map[int]int)
	for _, v := range shopIcardData {
		icardIDInt, _ := strconv.Atoi(v[shopIcardModel.Field.F_icard_id].(string))
		icardIds = append(icardIds, icardIDInt)
		shopStatus, _ := strconv.Atoi(v[shopIcardModel.Field.F_status].(string))
		shopItemId, _ := strconv.Atoi(v[shopIcardModel.Field.F_id].(string))
		shopStatusMap[icardIDInt] = shopStatus
		shopItemIdMap[icardIDInt] = shopItemId
	}

	list := []cards.ICardBase{}
	imgIds, totalNum := s.getICardList(busID, shopID, page, pageSize, ground, icardIds, &list, args.IsDel, args.FilterShopHasAdd)

	for k, v := range list {
		list[k].Status = shopStatusMap[v.IcardID]
		list[k].ShopItemId = shopItemIdMap[v.IcardID]
		list[k].SsID = shopItemIdMap[v.IcardID]
	}

	reply.Lists = list
	reply.TotalNum = totalNum
	reply.IndexImg = getImgsByImgIds(ctx, imgIds, cards.ITEM_TYPE_icard)
	return nil
}

//ShopList 门店列表
func (s *ICardLogic) ShopList(ctx context.Context, args cards.InputParamsICardList, reply *cards.OutputICardList) error {
	busID, _ := args.BsToken.GetBusId()
	page := args.Page
	pageSize := args.PageSize

	ShopID := args.ShopID
	if ShopID == 0 {
		return toolLib.CreateKcErr(_const.ICARD_PARAMS_ERROR)
	}

	shopIDs := []int{ShopID, models.AllShopID}

	IcardShop := new(models.IcardShopModel).Init()
	IcardShopData := IcardShop.GetAll(models.Condition{
		Where: map[string]interface{}{
			IcardShop.Field.F_shop_id: []interface{}{"IN", shopIDs},
		},
		// Offset: (page - 1) * pageSize,
		// Limit:  pageSize,
	})

	if len(IcardShopData) == 0 {
		reply.Lists = []cards.ICardBase{}
		reply.IndexImg = map[int]string{}
		return nil
	}

	var icardIDs []int
	for _, v := range IcardShopData {
		icardIDStr := v[IcardShop.Field.F_icard_id].(string)
		icardIDInt, _ := strconv.Atoi(icardIDStr)
		icardIDs = append(icardIDs, icardIDInt)
	}

	list := []cards.ICardBase{}
	imgIDs, totalNum := s.getICardList(busID, ShopID, page, pageSize, 0, icardIDs, &list, args.IsDel, args.FilterShopHasAdd)
	var ids []int
	for _, v := range list {
		ids = append(ids, v.IcardID)
	}
	gaagsNumMap := make(map[int]ApplyAndGiveSingleNum)
	if len(ids) > 0 {
		//获取不同卡项-适用单项目的个数和赠送单项目的个数
		gaagsNumMap = GetApplyAndGiveSingleNum(ids, cards.ITEM_TYPE_icard)
	}
	//添加状态
	shopIcardModel := new(models.ShopIcardModel).Init()
	shopIcardData := shopIcardModel.GetAll(models.Condition{
		Where: map[string]interface{}{
			shopIcardModel.Field.F_shop_id:  []interface{}{"IN", shopIDs},
			shopIcardModel.Field.F_icard_id: []interface{}{"IN", icardIDs},
		},
	})

	shopIcardStatus := make(map[int]interface{})
	for _, v := range shopIcardData {
		subIcardID, _ := strconv.Atoi(v[shopIcardModel.Field.F_icard_id].(string))
		subStatus, _ := strconv.Atoi(v[shopIcardModel.Field.F_status].(string))
		shopIcardStatus[subIcardID] = subStatus
	}

	for k, v := range list {
		if _, ok := shopIcardStatus[v.IcardID]; ok {
			list[k].IsSelfShop = 1
		}
		list[k].IsAllSingle = gaagsNumMap[v.IcardID].IsAllSingle
		list[k].ApplySingleNum = gaagsNumMap[v.IcardID].ApplySingleNum
		list[k].GiveSingleNum = gaagsNumMap[v.IcardID].GiveSingleNum
	}

	reply.Lists = list
	reply.TotalNum = totalNum
	reply.IndexImg = getImgsByImgIds(ctx, imgIDs, cards.ITEM_TYPE_icard)
	return nil
}

//SetOn 上架
func (s *ICardLogic) SetOn(ctx context.Context, args cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) (err error) {
	var icardIds []int
	if err := json.Unmarshal([]byte(args.IcardIds), &icardIds); err != nil {
		return toolLib.CreateKcErr(_const.ICARD_PARAMS_ERROR)
	}

	icardModel := new(models.IcardModel).Init()
	res := icardModel.UpdateAll(map[string]interface{}{
		icardModel.Field.F_is_ground:  cards.IS_GROUND_yes,
		icardModel.Field.F_under_time: 0,
	}, models.Condition{
		Where: map[string]interface{}{
			icardModel.Field.F_icard_id: []interface{}{"IN", icardIds},
		},
	})

	reply.IcardIds = icardIds
	reply.Rows = res

	//将已添加到子店的身份卡解除总店禁用状态 需要将不适用的门店套餐id 过滤掉
	mShopIcard := new(models.ShopIcardModel).Init()
	shopIcards := mShopIcard.GetByIcardids(icardIds)
	if len(shopIcards) == 0 {
		return
	}

	mIcardShop := new(models.IcardShopModel).Init()
	icardShops := mIcardShop.GetByIcardIds(icardIds)

	var icardShopArr = []string{} //smid_shopid
	for _, smshop := range icardShops {
		icardShopArr = append(icardShopArr, fmt.Sprintf("%s_%s", smshop[mIcardShop.Field.F_icard_id].(string), smshop[mIcardShop.Field.F_shop_id].(string)))
	}
	var unDisableIds = []int{}
	for _, shopicard := range shopIcards {
		smidShopidStr := fmt.Sprintf("%s_%s", shopicard[mShopIcard.Field.F_icard_id].(string), shopicard[mShopIcard.Field.F_shop_id].(string))
		smidAllStr := fmt.Sprintf("%s_0", shopicard[mShopIcard.Field.F_icard_id].(string))

		if functions.InArray(smidAllStr, icardShopArr) || functions.InArray(smidShopidStr, icardShopArr) {
			shopicardId, _ := strconv.Atoi(shopicard[mShopIcard.Field.F_id].(string))
			unDisableIds = append(unDisableIds, shopicardId)
		}
	}
	if len(unDisableIds) > 0 {
		mShopIcard.UpdateByIds(unDisableIds, map[string]interface{}{
			mShopIcard.Field.F_status: cards.STATUS_OFF_SALE,
		})
	}
	return
}

//SetOff 下架
func (s *ICardLogic) SetOff(ctx context.Context, args cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) (err error) {
	var icardIds []int
	if err := json.Unmarshal([]byte(args.IcardIds), &icardIds); err != nil {
		return toolLib.CreateKcErr(_const.ICARD_PARAMS_ERROR)
	}

	icardModel := new(models.IcardModel).Init()
	res := icardModel.UpdateAll(map[string]interface{}{
		icardModel.Field.F_is_ground:     cards.IS_GROUND_no,
		icardModel.Field.F_under_time:    time.Now().Unix(),
		icardModel.Field.F_sale_shop_num: 0,
	}, models.Condition{
		Where: map[string]interface{}{
			icardModel.Field.F_icard_id: []interface{}{"IN", icardIds},
		},
	})

	//门店所有此身份卡全部禁用
	shopIcardModel := new(models.ShopIcardModel).Init()
	shopIcardModel.UpdateAll(map[string]interface{}{
		shopIcardModel.Field.F_status: cards.STATUS_DISABLE,
	}, map[string]interface{}{
		shopIcardModel.Field.F_icard_id: []interface{}{"IN", icardIds},
	})

	reply.IcardIds = icardIds
	reply.Rows = res
	setShopItem(ctx, icardIds, 0, cards.ITEM_TYPE_icard)
	return
}

//SetOnOff 上下架
func (s *ICardLogic) SetOnOff(ctx context.Context, args cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) (err error) {
	icardID := args.IcardID
	icardModel := new(models.IcardModel).Init()
	icardData := icardModel.GetOneByPK(icardID)

	if len(icardData) == 0 {
		return toolLib.CreateKcErr(_const.ICARD_NOT_FOUND_ERROR)
	}

	currentIsGround, _ := strconv.Atoi(icardData[icardModel.Field.F_is_ground].(string))
	isGround := cards.IS_GROUND_no
	var underTime int64
	if currentIsGround != cards.IS_GROUND_yes {
		isGround = cards.IS_GROUND_no
	} else {
		underTime = time.Now().Unix()
	}

	icardModel.Update(map[string]interface{}{
		icardModel.Field.F_is_ground:  isGround,
		icardModel.Field.F_under_time: underTime,
	}, icardID)

	reply.IcardIds = []int{icardID}
	shopIcardModel := new(models.ShopIcardModel).Init()

	updateStatus := cards.STATUS_DISABLE
	if isGround != cards.IS_GROUND_no {
		updateStatus = cards.STATUS_OFF_SALE
	}

	//门店所有此身份卡全部禁用&上架
	shopIcardModel.UpdateAll(map[string]interface{}{
		shopIcardModel.Field.F_status: updateStatus,
	}, map[string]interface{}{
		shopIcardModel.Field.F_icard_id: icardID,
	})

	setShopItem(ctx, []int{icardID}, 0, cards.ITEM_TYPE_icard)
	return
}

//ShopSetOnOff 门店上下架
func (s *ICardLogic) ShopSetOnOff(ctx context.Context, args cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) (err error) {
	shopID, _ := args.BsToken.GetShopId()
	icardID := args.IcardID

	shopIcardModel := new(models.ShopIcardModel).Init()
	shopIcardData := shopIcardModel.GetAll(models.Condition{
		Where: map[string]interface{}{
			shopIcardModel.Field.F_shop_id:  shopID,
			shopIcardModel.Field.F_icard_id: icardID,
		},
	})

	if len(shopIcardData) == 0 {
		return toolLib.CreateKcErr(_const.ICARD_NOT_FOUND_ERROR)
	}
	//同步上架门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	for _, v := range shopIcardData {
		s := "dec"
		status := cards.STATUS_OFF_SALE
		currentStatus, _ := strconv.Atoi(v[shopIcardModel.Field.F_status].(string))
		if currentStatus != cards.STATUS_ON_SALE {
			status = cards.STATUS_ON_SALE
			s = "inc"
		}
		var underTime int64
		if status == cards.STATUS_OFF_SALE {
			underTime = time.Now().Unix()
		}
		shopIcardModel.Update(map[string]interface{}{
			shopIcardModel.Field.F_status:     status,
			shopIcardModel.Field.F_ctime:      time.Now().Unix(),
			shopIcardModel.Field.F_under_time: underTime,
		}, v[shopIcardModel.Field.F_id])

		icardModel := new(models.IcardModel).Init()
		icardModel.UpdateSaleShopNum([]int{icardID}, s)
		if !shopItemRealtionModel.UpdateStatusByItemIds([]int{icardID}, cards.ITEM_TYPE_icard, status, shopID) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	reply.IcardIds = []int{icardID}
	setShopItem(ctx, []int{icardID}, shopID, cards.ITEM_TYPE_icard)
	return
}

//ShopSetOn 门店上架
func (s *ICardLogic) ShopSetOn(ctx context.Context, args cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) (err error) {
	shopID, _ := args.BsToken.GetShopId()
	if shopID <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var icardIds []int
	if err := json.Unmarshal([]byte(args.IcardIds), &icardIds); err != nil {
		return toolLib.CreateKcErr(_const.ICARD_PARAMS_ERROR)
	}

	shopIcardModel := new(models.ShopIcardModel).Init()
	shopIcardData := shopIcardModel.GetAll(models.Condition{
		Where: map[string]interface{}{
			shopIcardModel.Field.F_shop_id:  shopID,
			shopIcardModel.Field.F_icard_id: []interface{}{"IN", icardIds},
		},
	})

	if len(shopIcardData) == 0 {
		return toolLib.CreateKcErr(_const.ICARD_NOT_FOUND_ERROR)
	}

	res := shopIcardModel.UpdateAll(map[string]interface{}{
		shopIcardModel.Field.F_status: cards.STATUS_ON_SALE,
	}, map[string]interface{}{
		shopIcardModel.Field.F_shop_id:  shopID,
		shopIcardModel.Field.F_icard_id: []interface{}{"IN", icardIds},
	})

	icardModel := new(models.IcardModel).Init()
	icardModel.UpdateSaleShopNum(icardIds, "inc")
	//同步上架门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.UpdateStatusByItemIds(icardIds, cards.ITEM_TYPE_icard, cards.STATUS_ON_SALE, shopID) {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	reply.IcardIds = icardIds
	reply.Rows = res
	setShopItem(ctx, icardIds, shopID, cards.ITEM_TYPE_icard)
	return
}

//ShopSetOff 门店下架
func (s *ICardLogic) ShopSetOff(ctx context.Context, args cards.InputParamsICardSetOnOff, reply *cards.OutputParamsICardSetOnOff) (err error) {
	shopID, _ := args.BsToken.GetShopId()
	if shopID <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	var icardIds []int
	if err := json.Unmarshal([]byte(args.IcardIds), &icardIds); err != nil {
		return toolLib.CreateKcErr(_const.ICARD_PARAMS_ERROR)
	}

	shopIcardModel := new(models.ShopIcardModel).Init()
	shopIcardData := shopIcardModel.GetAll(models.Condition{
		Where: map[string]interface{}{
			shopIcardModel.Field.F_shop_id:  shopID,
			shopIcardModel.Field.F_icard_id: []interface{}{"IN", icardIds},
		},
	})

	if len(shopIcardData) == 0 {
		return toolLib.CreateKcErr(_const.ICARD_NOT_FOUND_ERROR)
	}

	res := shopIcardModel.UpdateAll(map[string]interface{}{
		shopIcardModel.Field.F_status:     cards.STATUS_OFF_SALE,
		shopIcardModel.Field.F_under_time: time.Now().Unix(),
	}, map[string]interface{}{
		shopIcardModel.Field.F_shop_id:  shopID,
		shopIcardModel.Field.F_icard_id: []interface{}{"IN", icardIds},
	})

	icardModel := new(models.IcardModel).Init()
	icardModel.UpdateSaleShopNum(icardIds, "dec")
	//同步下架门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.UpdateStatusByItemIds(icardIds, cards.ITEM_TYPE_icard, cards.STATUS_OFF_SALE, shopID) {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	reply.IcardIds = icardIds
	reply.Rows = res
	setShopItem(ctx, icardIds, shopID, cards.ITEM_TYPE_icard)
	return
}

//AddToShop 添加至本店
func (s *ICardLogic) AddToShop(ctx context.Context, busID, shopID int, args cards.InputParamsICardAddToShop, reply *cards.OutputParamsICardAddToShop) (err error) {
	var icardIds []int
	if err1 := json.Unmarshal([]byte(args.IcardIds), &icardIds); err1 != nil {
		return toolLib.CreateKcErr(_const.ICARD_PARAMS_ERROR)
	}
	//icardModel := new(models.IcardModel).Init()
	//icardCount := icardModel.GetCount(models.Condition{
	//	Where: map[string]interface{}{
	//		icardModel.Field.F_icard_id:  []interface{}{"IN", icardIds},
	//		icardModel.Field.F_is_ground: models.IS_GROUND_OFF,
	//	},
	//})
	//
	//if icardCount > 0 {
	//	return toolLib.CreateKcErr(_const.ICARD_ALREADY_OFF)
	//}
	//判断本店是否存在包含商品
	//busId, _ := args.BsToken.GetBusId()
	err, checkSuccess := new(ItemLogic).CheckProductInShop(ctx, busID, shopID, cards.ITEM_TYPE_icard, icardIds)
	if err != nil {
		return err
	}
	if !checkSuccess {
		return toolLib.CreateKcErr(_const.SHOP_PRODUCT_NOT_CONTAIN_BUS_PRODUCT)
	}
	//校验当前门店是否已经将卡项内涉及到的单项目添加到自己的门店内
	allSingle, singleIds, err := new(ItemLogic).getItemCardIncSingleIds(icardIds, cards.ITEM_TYPE_icard)
	if err != nil {
		return
	}
	if err = new(ItemLogic).validShopSingleContainItemCardSingles(shopID, busID, allSingle, singleIds); err != nil {
		return
	}

	// 先进行检查，[1,2,3] ,其中2已经添加过，需要过滤掉，-- [1,3]
	// 过滤掉已经添加了的身份卡id
	shopIcardModel := new(models.ShopIcardModel).Init()
	shopIcardLists := shopIcardModel.GetByShopIDAndHNCardIDs(shopID, icardIds)
	shopIcIds := functions.ArrayValue2Array(shopIcardModel.Field.F_icard_id, shopIcardLists)

	// 刷选出已经添加过并且删除的数据
	delHcardIdSlice := make([]int, 0)
	for _, hcardMap := range shopIcardLists {
		isDel, _ := strconv.Atoi(hcardMap[shopIcardModel.Field.F_is_del].(string))
		if isDel == cards.IS_BUS_DEL_yes {
			delHcardId, _ := strconv.Atoi(hcardMap[shopIcardModel.Field.F_icard_id].(string))
			delHcardIdSlice = append(delHcardIdSlice, delHcardId)
		}
	}

	// 更新门店之前添加过并删除的数据
	if len(delHcardIdSlice) > 0 {
		// 更新数据删除和上下架状态
		if updateBool := shopIcardModel.UpdateShopIdBySmids(delHcardIdSlice, shopID, map[string]interface{}{
			shopIcardModel.Field.F_is_del: cards.IS_BUS_DEL_no,
			shopIcardModel.Field.F_status: cards.STATUS_OFF_SALE,
		}); !updateBool {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
		//更新卡项关联表
		sirModel := new(models.ShopItemRelationModel).Init()
		if b := sirModel.UpdateByItemIdsAndShopId(delHcardIdSlice, cards.ITEM_TYPE_icard, shopID, map[string]interface{}{
			sirModel.Field.F_is_del: cards.IS_BUS_DEL_no,
			sirModel.Field.F_status: cards.STATUS_OFF_SALE,
		}); !b {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}


	addIcardIds := make([]int, 0)
	for _, shopIcId := range icardIds {
		if functions.InArray(shopIcId, shopIcIds) == false {
			addIcardIds = append(addIcardIds, shopIcId)
		}
	}

	// 需要添加的数据
	var shopIcardData []map[string]interface{}
	shopItemRelationData := make([]map[string]interface{}, 0)
	shopItemRelationModel := new(models.ShopItemRelationModel).Init()
	for _, icardId := range addIcardIds {
		status := cards.STATUS_OFF_SALE
		ctime := time.Now().Local().Unix()
		shopIcardData = append(shopIcardData, map[string]interface{}{
			shopIcardModel.Field.F_icard_id: icardId,
			shopIcardModel.Field.F_shop_id:  shopID,
			shopIcardModel.Field.F_ctime:    ctime,
			shopIcardModel.Field.F_status:   status,
		})
		shopItemRelationData = append(shopItemRelationData, map[string]interface{}{
			shopItemRelationModel.Field.F_item_id:   icardId,
			shopItemRelationModel.Field.F_item_type: cards.ITEM_TYPE_icard,
			shopItemRelationModel.Field.F_status:    cards.STATUS_OFF_SALE,
			shopItemRelationModel.Field.F_shop_id:   shopID,
			shopItemRelationModel.Field.F_is_del:    cards.ITEM_IS_DEL_NO,
		})
	}
	// 过滤的数据添加到门店身份卡表
	if len(shopIcardData) > 0 {
		shopIcardModel.InsertAll(shopIcardData)
	}
	//门店卡项关联表数据插入
	if len(shopItemRelationData) > 0 {
		if shopItemRelationModel.InsertAll(shopItemRelationData) <= 0 {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	//shopIcardModel.DeleteAll(map[string]interface{}{
	//	shopIcardModel.Field.F_shop_id:  shopID,
	//	shopIcardModel.Field.F_icard_id: []interface{}{"IN", icardIds},
	//})

	// 初始化适用门店表模型
	icShopModel := new(models.IcardShopModel).Init()
	icShopLists := icShopModel.GetByIcardIds(icardIds)
	icshopIds := functions.ArrayValue2Array(icShopModel.Field.F_icard_id, icShopLists)
	addIcShopIds := make([]int, 0)
	for _, icardId := range icardIds {
		if functions.InArray(icardId, icshopIds) == false {
			addIcShopIds = append(addIcShopIds, icardId)
		}
	}
	if len(addIcShopIds) == 0 {
		return
	}

	var addIcShopData []map[string]interface{} // 添加适用门店表的数据
	for _, icardId := range addIcShopIds {
		addIcShopData = append(addIcShopData, map[string]interface{}{
			icShopModel.Field.F_icard_id: icardId,
			icShopModel.Field.F_shop_id:  shopID,
			icShopModel.Field.F_bus_id:   busID,
		})
	}
	// 过滤的数据添加到适用身份卡
	if len(addIcShopData) > 0 {
		icShopModel.InsertAll(addIcShopData)
	}

	reply.IcardIds = icardIds
	return
}

//CanUseICardList 查看用户权益身份卡列表
func (s *ICardLogic) CanUseICardList(ctx context.Context, args cards.InputParamsICardCanUse, reply *cards.OutputParamsICardCanUse) (err error) {
	shopID := args.ShopID
	UID := args.UID
	goods := args.GoodsIds
	singles := args.SingleIds

	reply.GoodsIcardList = []cards.CardInfo{}
	reply.SingleIcardList = []cards.CardInfo{}
	goodsCardInfo, singlesCardInfo, err := s.getCanUseICardList(ctx, UID, shopID, goods, singles)
	if err != nil {
		return
	}

	if len(goodsCardInfo) > 0 {
		reply.GoodsIcardList = goodsCardInfo
	}

	if len(singlesCardInfo) > 0 {
		reply.SingleIcardList = singlesCardInfo
	}
	return
}

//CanUseICardListForUser 查看用户权益身份卡列表
func (s *ICardLogic) CanUseICardListForUser(ctx context.Context, args cards.InputParamsICardCanUseForUser, reply *cards.OutputParamsICardCanUse) (err error) {
	UID, _ := args.Utoken.GetUid()
	shopID := args.ShopID
	goods := args.GoodsIds
	singles := args.SingleIds

	logs.Info("CanUseICardListForUser UID:", UID, "shopID:", shopID, "goods:", goods, "singles:", singles)
	reply.GoodsIcardList = []cards.CardInfo{}
	reply.SingleIcardList = []cards.CardInfo{}
	goodsCardInfo, singlesCardInfo, err := s.getCanUseICardList(ctx, UID, shopID, goods, singles)
	if err != nil {
		return
	}

	if len(goodsCardInfo) > 0 {
		reply.GoodsIcardList = goodsCardInfo
	}

	if len(singlesCardInfo) > 0 {
		reply.SingleIcardList = singlesCardInfo
	}
	return
}

func (s *ICardLogic) getCanUseICardList(ctx context.Context, UID, shopID int, goods, singles string) (goodsCardInfo, singlesCardInfo []cards.CardInfo, err error) {
	icardOrderClient := new(order2.IcardOrderClient).Init()
	defer icardOrderClient.Close()
	orderArgs := &order.InputParamsICardCanUse{
		UID:       UID,
		ShopID:    shopID,
		GoodsIds:  goods,
		SingleIds: singles,
	}

	orderReply := &order.OutputParamsICardCanUse{}
	if err = icardOrderClient.GetIcardListByUserID(context.Background(), orderArgs, orderReply); err != nil {
		return
	}

	// getImg(ctx, imgID, cards.ITEM_TYPE_icard)
	// var goodsCardInfo []cards.CardInfo
	for _, v := range orderReply.GoodsIcardList {
		_, imgURL := getImg(ctx, v.ImgID, cards.ITEM_TYPE_icard)
		goodsCardInfo = append(goodsCardInfo, cards.CardInfo{
			Name:          v.Name,
			CardPackageID: v.CardPackageID,
			Discount:      v.Discount,
			ImgID:         v.ImgID,
			ImgURL:        imgURL,
			IcardID:       v.IcardID,
			ExpireTime:    v.ExpireTime,
			ExpireSurDay:  v.ExpireSurDay,
			ServicePeriod: v.ServicePeriod,
		})
	}

	// var singlesCardInfo []cards.CardInfo
	for _, v := range orderReply.SingleIcardList {
		_, imgURL := getImg(ctx, v.ImgID, cards.ITEM_TYPE_icard)
		singlesCardInfo = append(singlesCardInfo, cards.CardInfo{
			Name:          v.Name,
			CardPackageID: v.CardPackageID,
			Discount:      v.Discount,
			ImgID:         v.ImgID,
			ImgURL:        imgURL,
			IcardID:       v.IcardID,
			ExpireTime:    v.ExpireTime,
			ExpireSurDay:  v.ExpireSurDay,
			ServicePeriod: v.ServicePeriod,
		})
	}

	return
}

//UserICardList 查看用户身份卡列表
func (s *ICardLogic) UserICardList(ctx context.Context, args cards.InputParams) (reply cards.PageOutputReply) {
	return
}

//获取iCard企业基本信息-风控统计用
func (i *ICardLogic) GetBusBaseInfoRpc(iCardId int) (reply cards.ReplyGetBusBaseInfoRpc, err error) {
	iCardmodel := new(models.IcardModel).Init()
	resMap := iCardmodel.GetOneByPK(iCardId)
	_ = mapstructure.WeakDecode(resMap, &reply)
	return
}

//获取身份卡的折扣信息
func (i *ICardLogic) GetICardDiscountById(ctx context.Context, iCardId int, reply *cards.ReplyGetIcardDiscountById) (err error) {
	//身份卡包含的商品
	igModel := new(models.IcardGoodsModel).Init()
	igMaps := igModel.GetByIcardIds([]int{iCardId})
	_ = mapstructure.WeakDecode(igMaps, &reply.ProductDiscount)
	//身份卡包含的项目
	isModel := new(models.IcardSingleModel).Init()
	isMaps := isModel.GetByIcardIds([]int{iCardId})
	_ = mapstructure.WeakDecode(isMaps, &reply.SingleDiscount)
	reply.ICardId = iCardId
	//查询商品信息
	if len(reply.ProductDiscount) > 0 && reply.ProductDiscount[0].GoodsId > 0 {
		goodIds := functions.ArrayValue2Array(igModel.Field.F_goods_id, igMaps)
		rpcProduct := new(product.Product).Init()
		defer rpcProduct.Close()
		var rpcProductRes []product2.ReplyProductGetByIds
		if err = rpcProduct.GetProductByIds(ctx, &product2.ArgsProductGetByIds{Ids: goodIds}, &rpcProductRes); err != nil {
			return
		}
		for index, v := range reply.ProductDiscount {
			for _, v2 := range rpcProductRes {
				if v.GoodsId == v2.Id {
					reply.ProductDiscount[index].Name = v2.Name
					break
				}
			}
		}
	}
	//查询单项目信息
	if len(reply.SingleDiscount) > 0 && reply.SingleDiscount[0].SingleId > 0 {
		singleIds := functions.ArrayValue2Array(isModel.Field.F_single_id, isMaps)
		singleModel := new(models.SingleModel).Init()
		singleMaps := new(SingleLogic).GetSimpleSingleInfos(ctx, singleIds)
		for index, v := range reply.SingleDiscount {
			for _, v2 := range *singleMaps {
				v2SingleId, _ := strconv.Atoi(v2[singleModel.Field.F_single_id].(string))
				if v.SingleId == v2SingleId {
					reply.SingleDiscount[index].Name = v2[singleModel.Field.F_name].(string)
					break
				}
			}
		}
	}
	return
}

//获取身份卡备份表中的项目折扣
func (i *ICardLogic) GetICardSingleDiscount(ctx context.Context, args *cards.ArgsGetICardSingleDiscount, reply *cards.ReplyGetICardSingleDiscount) (err error) {
	iCardId, isSync := args.ICardId, args.IsSync
	discountMaps := make([]map[string]interface{}, 0)
	if args.RequestType == 0 {
		isbModel := new(models.IcardSingleBackupModel).Init()
		lastBackupNum := isbModel.GetLastBackUumByIcardId(false, isSync, iCardId)
		isbModel2 := new(models.IcardSingleBackupModel).Init()
		isbMaps := isbModel2.FindByIcardIds([]int{iCardId}, lastBackupNum)
		if len(isbMaps) == 0 {
			return
		}
		singleIds := make([]int, 0)
		for _, v := range isbMaps {
			singleId, _ := strconv.Atoi(v[isbModel2.Field.F_single_id].(string))
			if singleId > 0 {
				singleIds = append(singleIds, singleId)
			}
		}
		if len(singleIds) > 0 {
			//获取单项目name
			res := new(SingleLogic).GetSimpleSingleInfos(ctx, singleIds)
			resRebuild := functions.ArrayRebuild(isbModel2.Field.F_single_id, *res)
			for index, v := range isbMaps {
				singleId, _ := v[isbModel2.Field.F_single_id].(string)
				if _, ok := resRebuild[singleId]; ok {
					isbMaps[index]["SingleName"] = resRebuild[singleId].(map[string]interface{})["name"].(string)
				}
			}
		}
		discountMaps = isbMaps
	} else {
		igbModel := new(models.IcardGoodsBackupModel).Init()
		lastBackupNum := igbModel.GetLastBackUumByIcardId(false, isSync, iCardId)
		isbModel2 := new(models.IcardGoodsBackupModel).Init()
		isbMaps := isbModel2.FindByIcardIds([]int{iCardId}, lastBackupNum)
		if len(isbMaps) == 0 {
			return
		}
		goodIds := make([]int, 0)
		for _, v := range isbMaps {
			goodsId, _ := strconv.Atoi(v[isbModel2.Field.F_goods_id].(string))
			if goodsId > 0 {
				goodIds = append(goodIds, goodsId)
			}
		}
		if len(goodIds) > 0 {
			//产品
			rpcProduct := new(product.Product).Init()
			defer rpcProduct.Close()
			var rpcProductRes []product2.ReplyProductGetByIds
			if err = rpcProduct.GetProductByIds(ctx, &product2.ArgsProductGetByIds{Ids: goodIds}, &rpcProductRes); err != nil {
				return
			}
			if len(rpcProductRes) == 0 {
				return
			}
			for index, v := range isbMaps {
				goodsId, _ := strconv.Atoi(v[isbModel2.Field.F_goods_id].(string))
				for _, v2 := range rpcProductRes {
					if v2.Id == goodsId {
						isbMaps[index]["GoodsName"] = v2.Name
						break
					}
				}
			}
		}
		discountMaps = isbMaps
	}

	_ = mapstructure.WeakDecode(discountMaps, &reply.SingleDiscount)
	return
}

//同步身份卡折扣到卡包
func (i *ICardLogic) SetIcardDiscountTask(ctx context.Context, iCardId int) {
	rpcTask := new(cardsTask.ShopItems).Init()
	defer rpcTask.Close()
	var reply bool
	rpcTask.SetIcardDiscount(ctx, &iCardId, &reply)
}
