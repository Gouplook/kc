package logics

// 限时卡业务处理
import (
	"context"
	"encoding/json"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/rpcCards/common/tools"
	"git.900sui.cn/kc/rpcCards/constkey"
	"strconv"
	"time"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/mapstructure"
	"git.900sui.cn/kc/redis"
	"git.900sui.cn/kc/rpcCards/common/models"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/client/bus"
	bus2 "git.900sui.cn/kc/rpcinterface/interface/bus"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	redis2 "github.com/gomodule/redigo/redis"
)

// HcardLogic HcardLogic
type HcardLogic struct{}

// AddHcard 添加限时卡数据
func (h *HcardLogic) AddHcard(ctx context.Context, args *cards.ArgsAddHcard) (hcardID int, err error) {

	// 验证商户
	busID, err := checkBus(args.BsToken, true)
	if err != nil {
		return
	}
	// 验证参数
	if err = h.checkHcardData(busID, args.HcardBase, args.IncludeSingles, args.GiveSingles); err != nil {
		return
	}
	// 验证图片
	imgID, err := checkImg(ctx, args.ImgHash)
	if err != nil {
		return
	}
	// 是否包含赠送服务
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
	// 添加限时卡基本信息数据
	hcardModel := new(models.HcardModel).Init()
	hcardModel.Model.Begin()
	hcardID = hcardModel.Insert(map[string]interface{}{
		hcardModel.Field.F_bus_id:          busID,
		hcardModel.Field.F_price:           args.Price,
		hcardModel.Field.F_is_ground:       cards.SINGLE_IS_GROUND_no,
		hcardModel.Field.F_real_price:      args.RealPrice,
		hcardModel.Field.F_ctime:           time.Now().Unix(),
		hcardModel.Field.F_img_id:          imgID,
		hcardModel.Field.F_name:            args.Name,
		hcardModel.Field.F_sort_desc:       args.SortDesc,
		hcardModel.Field.F_bind_id:         getBusMainBindId(ctx, busID),
		hcardModel.Field.F_has_give_signle: hasGive,
		hcardModel.Field.F_service_period:  args.ServicePeriod,
	})
	if hcardID <= 0 { // 添加失败
		hcardModel.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	// 添加单项目
	hcardSingleModel := new(models.HcardSingleModel).Init(hcardModel.Model.GetOrmer())
	var hcardSingleData []map[string]interface{}
	for _, signle := range args.IncludeSingles {
		hcardSingleData = append(hcardSingleData, map[string]interface{}{
			hcardSingleModel.Field.F_hcard_id:  hcardID,
			hcardSingleModel.Field.F_single_id: signle.SingleID,
		})
	}
	hcardSingID := hcardSingleModel.InsertAll(hcardSingleData)
	if hcardSingID <= 0 { // 添加失败
		hcardModel.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	// 添加赠送项目
	if hasGive {
		hgm := new(models.HcardGiveModel).Init(hcardModel.Model.GetOrmer())
		var giveSingleData []map[string]interface{}
		for _, signle := range args.GiveSingles {
			giveSingleData = append(giveSingleData, map[string]interface{}{
				hgm.Field.F_hcard_id:           hcardID,
				hgm.Field.F_single_id:          signle.SingleID,
				hgm.Field.F_num:                signle.Num,
				hgm.Field.F_period_of_validity: signle.PeriodOfValidity,
			})
		}
		giveSingleID := hgm.InsertAll(giveSingleData)
		if giveSingleID <= 0 { // 添加失败
			hcardModel.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	// 添加赠品描述
	if hasGive && len(args.GiveSingleDesc) > 0 {
		hdescm := new(models.HcardGiveDescModel).Init(hcardModel.Model.GetOrmer())
		giveSingleDescStr, _ := json.Marshal(args.GiveSingleDesc)
		descData := map[string]interface{}{
			hdescm.Field.F_hcard_id: hcardID,
			hdescm.Field.F_desc:     string(giveSingleDescStr),
		}
		if _, err = hdescm.Model.Data(descData).Insert(); err != nil {
			hcardModel.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	// 添加现时卡拓展信息数据
	if err = h.AddOrEditHcardExt(hcardModel, hcardID, args.Notes, true); err != nil {
		hcardModel.Model.RollBack()
		return
	}
	hcardModel.Model.Commit()
	// 添加风控统计任务
	new(ItemLogic).AddXCardTask(ctx, hcardID, cards.ITEM_TYPE_hcard)
	return
}

// EditHcard 修改限时卡数据
func (h *HcardLogic) EditHcard(ctx context.Context, args *cards.ArgsEditHcard) (err error) {
	// 检查商户
	busID, err := checkBus(args.BsToken, true)
	if err != nil {
		return
	}
	// 根据busid获取商户busID
	hcardModel := new(models.HcardModel).Init()
	hcardInfo := hcardModel.GetHcardByID(args.HcardID, hcardModel.Field.F_bus_id)
	if len(hcardInfo) == 0 { // 未找到数据
		err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
		return
	}
	// 当前商户id和修改的限时卡所属商户id不匹配
	if hcardInfo[hcardModel.Field.F_bus_id] != strconv.Itoa(busID) {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	// 检查入参数据
	if err = h.checkHcardData(busID, args.HcardBase, args.IncludeSingles, args.GiveSingles); err != nil {
		return
	}
	// 检查封面图片
	imgID, err := checkImg(ctx, args.ImgHash)
	if err != nil {
		return
	}
	// 是否包含赠送的单项目
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
	// 限时卡主体信息更新
	hcardModel.Model.Begin()
	if !hcardModel.UpdateByHcardID(args.HcardID, map[string]interface{}{
		hcardModel.Field.F_price:           args.Price,
		hcardModel.Field.F_real_price:      args.RealPrice,
		hcardModel.Field.F_img_id:          imgID,
		hcardModel.Field.F_name:            args.Name,
		hcardModel.Field.F_sort_desc:       args.SortDesc,
		hcardModel.Field.F_has_give_signle: hasGive,
		hcardModel.Field.F_service_period:  args.ServicePeriod}) {
		hcardModel.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	// ------------------修改限时卡包含的单项目操作(单项目无需更新因为没有num字段)------------------
	// 表中已经存在的所有单项目
	hcardSingleModel := new(models.HcardSingleModel).Init(hcardModel.Model.GetOrmer())
	oldHcardSingles := hcardSingleModel.GetByHcardID(args.HcardID)

	var (
		// 新增的限时卡单项目
		newAddHcardSingles []map[string]interface{}
		// 需要删除的限时卡单项目
		delHcardSingleIDs []int
	)
	// 收集需要新增单项目的数据
begin1:
	for _, includeSingle := range args.IncludeSingles { // 新的单项目
		for _, oldHcardSingle := range oldHcardSingles { // 旧的单项目
			oldHcardSingleIDStr := oldHcardSingle[hcardSingleModel.Field.F_single_id].(string)
			if strconv.Itoa(includeSingle.SingleID) == oldHcardSingleIDStr {
				// 不需要记录需要更新的数据因没更新字段，后期如果有的话可以在此处归集
				continue begin1
			}
		}
		// 需要新增的数据
		newAddHcardSingles = append(newAddHcardSingles, map[string]interface{}{
			hcardSingleModel.Field.F_hcard_id:  args.HcardID,
			hcardSingleModel.Field.F_single_id: includeSingle.SingleID,
		})
	}
	// 收集需要删除的旧项目id
begin2:
	for _, oldHcardSingle := range oldHcardSingles { // 旧的单项目
		for _, includeSingle := range args.IncludeSingles { // 新的单项目
			oldHcardSingleIDStr := oldHcardSingle[hcardSingleModel.Field.F_single_id].(string)
			if strconv.Itoa(includeSingle.SingleID) == oldHcardSingleIDStr {
				continue begin2
			}
			// 需要删除的限时单项目id(此处收集的是旧单项目id)
			delID, _ := strconv.Atoi(oldHcardSingle[hcardSingleModel.Field.F_id].(string))
			delHcardSingleIDs = append(delHcardSingleIDs, delID)
		}
	}
	if len(delHcardSingleIDs) > 0 && !hcardSingleModel.DeleteByIDs(delHcardSingleIDs) {
		hcardSingleModel.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	if len(newAddHcardSingles) > 0 && (hcardSingleModel.InsertAll(newAddHcardSingles)) <= 0 {
		hcardSingleModel.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	// ------------------修改赠送的单项目操作------------------
	hcardGiveModel := new(models.HcardGiveModel).Init(hcardModel.Model.GetOrmer())
	// 表中已经存在的赠送项目
	oldHcardGives := hcardGiveModel.GetByHcardID(args.HcardID)
	var (
		// 需要新增的赠送项目
		newAddHcardGives []map[string]interface{}
		// 需要删除的赠送项目
		delHcardGiveIDs []int
		// 需要更新的赠送项目
		updateHcardGives = map[int]int{}
	)
	// 收集需要新增和更新的赠送项目数据
begin3:
	for _, giveSingle := range args.GiveSingles {
		for _, oldHcardGive := range oldHcardGives {
			oldHcardGiveIDStr := oldHcardGive[hcardGiveModel.Field.F_single_id].(string)
			if strconv.Itoa(giveSingle.SingleID) == oldHcardGiveIDStr {
				// 如果数据新旧不一致旧记录下数据后面批量更新
				if strconv.Itoa(giveSingle.Num) != oldHcardGive[hcardGiveModel.Field.F_num].(string) {
					giveID, _ := strconv.Atoi(oldHcardGive[hcardGiveModel.Field.F_id].(string))
					updateHcardGives[giveID] = giveSingle.Num
				}
				continue begin3
			}
		}
		newAddHcardGives = append(newAddHcardGives, map[string]interface{}{
			hcardGiveModel.Field.F_single_id:          giveSingle.SingleID,
			hcardGiveModel.Field.F_hcard_id:           args.HcardID,
			hcardGiveModel.Field.F_num:                giveSingle.Num,
			hcardGiveModel.Field.F_period_of_validity: giveSingle.PeriodOfValidity,
		})
	}
	// 收集需要删除的赠送项目
begin4:
	for _, oldHcardGive := range oldHcardGives {
		for _, giveSingle := range args.GiveSingles {
			if strconv.Itoa(giveSingle.SingleID) == oldHcardGive[hcardGiveModel.Field.F_single_id].(string) {
				continue begin4
			}
		}
		delID, _ := strconv.Atoi(oldHcardGive[hcardGiveModel.Field.F_id].(string))
		delHcardGiveIDs = append(delHcardGiveIDs, delID)
	}
	if len(delHcardGiveIDs) > 0 && !hcardGiveModel.DeleteByIDs(delHcardGiveIDs) {
		hcardGiveModel.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	if len(newAddHcardGives) > 0 && (hcardGiveModel.InsertAll(newAddHcardGives)) <= 0 {
		hcardGiveModel.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	if len(updateHcardGives) > 0 {
		for id, num := range updateHcardGives {
			if !hcardGiveModel.UpdateNumByID(id, num) {
				hcardGiveModel.Model.RollBack()
				err = toolLib.CreateKcErr(_const.DB_ERR)
				return
			}
		}
	}

	// 修改赠品描述
	hgdescm := new(models.HcardGiveDescModel).Init(hcardModel.Model.GetOrmer())
	// 1.删除原有赠品描述
	if err = hgdescm.DelByHcardId(args.HcardID); err != nil {
		hcardModel.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	// 2.新增赠品描述
	if len(args.GiveSingles) > 0 && len(args.GiveSingleDesc) > 0 {
		giveSingleDescStr, _ := json.Marshal(args.GiveSingleDesc)
		descData := map[string]interface{}{
			hgdescm.Field.F_hcard_id: args.HcardID,
			hgdescm.Field.F_desc:     string(giveSingleDescStr),
		}
		if _, err = hgdescm.Model.Data(descData).Insert(); err != nil {
			hcardModel.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	// 添加现时卡拓展信息数据
	if err = h.AddOrEditHcardExt(hcardModel, args.HcardID, args.Notes, false); err != nil {
		hcardModel.Model.RollBack()
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	hcardModel.Model.Commit()

	return
}

// DeleteHcard 总店删除限时卡
func (h *HcardLogic) DeleteHcard(ctx context.Context, args *cards.ArgsDeleteHcard) (err error) {
	if len(args.HcardIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	// 检查商户
	busID, err := checkBus(args.BsToken, true)
	if err != nil {
		return
	}
	hcardModel := new(models.HcardModel).Init()
	hcardList := hcardModel.FindHcardByHcardIDAndBusID(args.HcardIds, busID, hcardModel.Field.F_bus_id)

	// 检查需要删除的卡是否属于商户
	if len(hcardList) != len(args.HcardIds) {
		err = toolLib.CreateKcErr(_const.NO_IN_BUS)
		return
	}
	// 检查子店是否已经添加过该卡项(不用检测，由于总店没有上下架功能）
	shopHcardModel := new(models.ShopHcardModel).Init()

	// 删除总店 并且同步分店
	r := hcardModel.UpdateDelByHcardIds(args.HcardIds, busID, map[string]interface{}{
		hcardModel.Field.F_is_del:        cards.IS_BUS_DEL_yes,
		hcardModel.Field.F_del_time:      time.Now().Unix(),
		hcardModel.Field.F_sale_shop_num: 0,
	})
	if r == false {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	// 同步分店
	r2 := shopHcardModel.UpdateDelByHcardIds(args.HcardIds, map[string]interface{}{
		shopHcardModel.Field.F_is_del:   cards.IS_BUS_DEL_yes,
		shopHcardModel.Field.F_del_time: time.Now().Unix(),
	})
	if r2 == false {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	// 同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.HcardIds, cards.ITEM_TYPE_hncard) {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	return
}

// HcardInfo 获取限时卡详情数据
func (h *HcardLogic) HcardInfo(ctx context.Context, hcardID int, shopID ...int) (reply cards.ReplyHcardInfo, err error) {
	reply.IncludeSingles = make([]cards.IncInfSingleDetail, 0)
	reply.GiveSingles = make([]cards.IncSingleDetail, 0)
	reply.GiveSingleDesc = make([]cards.GiveSingleDesc, 0)
	hcardModel := new(models.HcardModel).Init()
	hcardInfo := hcardModel.GetHcardByID(hcardID)
	if len(hcardInfo) == 0 {
		err = toolLib.CreateKcErr(_const.INFO_NO_EXISTS)
		return
	}
	// 查询子店的信息(当shopId存在时)
	shopStatus := 0
	shopSales := 0
	shopid := 0
	if len(shopID) > 0 && shopID[0] > 0 {
		shopHcardModel := new(models.ShopHcardModel).Init()
		shopHcInfo := shopHcardModel.GetByShopIDAndHcardIDs(shopID[0], []int{hcardID})
		if len(shopHcInfo) == 0 {
			err = toolLib.CreateKcErr(_const.NO_IN_SHOP)
			return
		}
		shopStatus, _ = strconv.Atoi(shopHcInfo[0][shopHcardModel.Field.F_status].(string))
		reply.SsId, _ = strconv.Atoi(shopHcInfo[0][shopHcardModel.Field.F_id].(string))
		shopSales, _ = strconv.Atoi(shopHcInfo[0][shopHcardModel.Field.F_sales].(string))
		shopid = shopID[0]
	}
	reply.ShareLink = tools.GetShareLink(hcardID, shopid, cards.ITEM_TYPE_hcard)

	imgID, _ := strconv.Atoi(hcardInfo[hcardModel.Field.F_img_id].(string))
	imgHash, imgURL := getImg(ctx, imgID, cards.ITEM_TYPE_hcard)
	reply.ImgHash = imgHash
	reply.ImgUrl = imgURL
	reply.ShopStatus = shopStatus
	_ = mapstructure.WeakDecode(hcardInfo, &reply.HcardBase)
	reply.CtimeStr = time.Unix(int64(reply.Ctime), 0).Format("2006/01/02 15:04:05")
	_ = mapstructure.WeakDecode(hcardInfo, &reply)
	if len(shopID) > 0 && shopID[0] > 0 {
		reply.Sales = shopSales
	}
	// 商户信息
	if err = getBusInfo(ctx, reply.BusID, &reply.BusInfo); err != nil {
		err = toolLib.CreateKcErr(_const.SHOP_INFO_ERR)
		return
	}

	// 获取包含的搭配项目
	hcardSingleModel := new(models.HcardSingleModel).Init()
	hcardSingles := hcardSingleModel.GetByHcardID(hcardID)
	if len(hcardSingles) > 0 && hcardSingles[0][hcardSingleModel.Field.F_single_id].(string) == "0" {
		reply.IsAllSingle = true
	}
	// 获取赠送项目
	var allGiveSingles map[int]cards.IncSingleDetail
	hcardGiveModel := new(models.HcardGiveModel).Init()
	var hcardGiveSingles []map[string]interface{}
	// 如果限时卡详情中包含赠送项目就将所有的赠送项目找出来
	if hcardInfo[hcardModel.Field.F_has_give_signle].(string) == strconv.Itoa(cards.HAS_GIVE_SINGLE_yes) {
		hcardGiveSingles = hcardGiveModel.GetByHcardID(hcardID)
		hcardGiveSingIDs := functions.ArrayValue2Array(hcardGiveModel.Field.F_single_id, hcardGiveSingles) // 赠送的项目ids
		allGiveSingles = getIncSingles(ctx, hcardGiveSingIDs)
	}

	if len(hcardGiveSingles) > 0 { // 详情包含的赠送项目
		for _, giveSignle := range hcardGiveSingles {
			singleID, _ := strconv.Atoi(giveSignle[hcardGiveModel.Field.F_single_id].(string))
			sInfo := allGiveSingles[singleID]
			sInfo.Num, _ = strconv.Atoi(giveSignle[hcardGiveModel.Field.F_num].(string))
			sInfo.PeriodOfValidity, _ = strconv.Atoi(giveSignle[hcardGiveModel.Field.F_period_of_validity].(string))
			reply.GiveSingles = append(reply.GiveSingles, sInfo)
		}

		// 获取赠品描述信息
		hgdescm := new(models.HcardGiveDescModel).Init()
		desc, ok := hgdescm.GetByHcardId(hcardID)[hgdescm.Field.F_desc].(string)
		if ok {
			json.Unmarshal([]byte(desc), &reply.GiveSingleDesc)
		}
	}
	//	获取现时卡拓展信息数据
	h.GetHcardExitByHcardID(hcardID, &reply.Notes)

	// 获取限时卡门店添加详情  15 -- []int {3,4,6}
	busId, _ := strconv.Atoi(hcardInfo[hcardModel.Field.F_bus_id].(string))
	hSModel := new(models.HcardShopModel).Init()
	hsInfo := hSModel.GetByHcardIDsByBusId(hcardID, busId)

	hsShopIds := make([]int, 0)
	for _, hsInfoValue := range hsInfo {
		sshopId, _ := strconv.Atoi(hsInfoValue[hSModel.Field.F_shop_id].(string))
		hsShopIds = append(hsShopIds, sshopId)
	}

	var replyShop []bus2.ReplyShopName
	rLists := make([]cards.ReplyShopName, 0)
	rpcBus := new(bus.Shop).Init()
	defer rpcBus.Close()
	err = rpcBus.GetShopNameByShopIds(ctx, &hsShopIds, &replyShop)
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

	// 浏览次数加1
	_ = redis.RedisGlobMgr.Hincrby(constkey.HCARD_CLIKS, strconv.Itoa(hcardID), 1)
	return
}

// BusHcardPage 获取总店的限时卡列表
func (h *HcardLogic) BusHcardPage(ctx context.Context, busID, shopId, start, limit int, isGround string, filterShopHasAdd bool) (list cards.ReplyHcardPage, err error) {
	if busID <= 0 || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	// 获取总店的限时卡信息
	hcardModel := new(models.HcardModel).Init()
	hcards := make([]map[string]interface{}, 0)
	list.List = make([]cards.HcardDesc, 0)
	list.IndexImg = make(map[int]string)

	// //获取总数量
	// if isGround == "" {
	//	hcards = hcardModel.GetPageByBusID(busID, start, limit)
	//	list.TotalNum = hcardModel.GetNumByBusID(busID)
	// } else {
	//	isground, _ := strconv.Atoi(isGround)
	//	isground = isground - 1
	//	hcards = hcardModel.GetPageByBusID(busID, start, limit, isground)
	//	list.TotalNum = hcardModel.GetNumByBusID(busID, isground)
	// }

	// 子店已添加的卡项
	var shopAddCards []map[string]interface{}
	scModel := new(models.ShopHcardModel).Init()
	where := make([]base.WhereItem, 0)
	where = append(where, base.WhereItem{hcardModel.Field.F_bus_id, busID})
	where = append(where, base.WhereItem{hcardModel.Field.F_is_del, cards.IS_BUS_DEL_no})
	if shopId > 0 {
		shopAddCards = scModel.SelectRcardsByWherePage([]base.WhereItem{{scModel.Field.F_shop_id, shopId}, {scModel.Field.F_is_del, cards.IS_BUS_DEL_no}}, 0, 0)
		if filterShopHasAdd && len(shopAddCards) > 0 {
			shopHasAddhcardIds := functions.ArrayValue2Array(scModel.Field.F_hcard_id, shopAddCards)
			where = append(where, base.WhereItem{hcardModel.Field.F_hcard_id, []interface{}{"NOT IN", shopHasAddhcardIds}})
		}
	}

	// 获取总数量 过滤删除的限时卡(
	hcards = hcardModel.SelectHcardsByWherePage(where, start, limit)
	list.TotalNum = hcardModel.GetNumByWhere(where)

	if len(hcards) == 0 {
		return
	}

	list.List = make([]cards.HcardDesc, len(hcards))
	list.IndexImg = make(map[int]string)
	for index, hcard := range hcards {
		_ = mapstructure.WeakDecode(hcard, &list.List[index].HcardBase)
		for _, shopCard := range shopAddCards {
			if hcard[hcardModel.Field.F_hcard_id].(string) == shopCard[scModel.Field.F_hcard_id].(string) { // 表明当前子店已添加该卡项
				hcard["ShopItemId"] = shopCard[scModel.Field.F_id].(string)
				hcard["ShopStatus"] = shopCard[scModel.Field.F_status].(string)
				hcard["ShopHasAdd"] = 1
				hcard["ShopDelStatus"] = shopCard[scModel.Field.F_is_del].(string)
				break
			}
		}
		_ = mapstructure.WeakDecode(hcard, &list.List[index])
		list.List[index].CtimeStr = time.Unix(int64(list.List[index].Ctime), 0).Format("2006/01/02 15:04:05")
		list.List[index].Clicks, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.HCARD_CLIKS, hcard[hcardModel.Field.F_hcard_id].(string)))
	}
	// 获取图片信息
	imgIDs := functions.ArrayValue2Array(hcardModel.Field.F_img_id, hcards)
	list.IndexImg = getImgsByImgIds(ctx, imgIDs, cards.ITEM_TYPE_hcard)
	return
}

// SetHcardShop 总店设置限时卡的适用门店（废用）
func (h *HcardLogic) SetHcardShop(ctx context.Context, args *cards.ArgsSetHcardShop) (err error) {
	var busID int
	busID, err = checkBus(args.BsToken, true)
	if err != nil {
		return
	}
	if len(args.HcardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.CARD_ID_IS_NIL)
		return
	}
	if args.IsAllShop == false && len(args.ShopIDs) == 0 {
		err = toolLib.CreateKcErr(_const.SHOPID_NTL)
		return
	}

	// 限次卡id重复提交判断
	realHcardIDs := functions.ArrayUniqueInt(args.HcardIDs)
	if len(realHcardIDs) != len(args.HcardIDs) {
		err = toolLib.CreateKcErr(_const.DATA_REPEAT_ERR)
		return
	}

	// 店铺id重复提交判断
	realShopIds := functions.ArrayUniqueInt(args.ShopIDs)
	if args.IsAllShop == false && len(realShopIds) != len(args.ShopIDs) {
		err = toolLib.CreateKcErr(_const.DATA_REPEAT_ERR)
		return
	}
	// 检查限时卡是否属于企业;是否上架
	hcardModel := new(models.HcardModel).Init()
	// 1.该方法适用于全部状态的卡项，
	dataArr := hcardModel.GetHcardByIDs(realHcardIDs) // 根据卡id获取全部数据
	// 1.这部分是只允许上架的卡项适用门店
	// dataArr := hcardModel.GetByHcardIDsAndGround(realHcardIDs, hcardModel.Field.F_bus_id) // 根据需要适用的卡id获取上架的限时卡数据
	if len(dataArr) != len(realHcardIDs) { // 校验是否有不属于总店的现时卡(如果全部为总店的卡时，两者应该相等)
		err = toolLib.CreateKcErr(_const.PARAM_ERR, "无效限时卡id")
		return
	}
	busIDStr := strconv.Itoa(busID)
	for _, data := range dataArr { // 数据中的busID是否和token中的busID一致
		if busIDStr != data[hcardModel.Field.F_bus_id].(string) {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
	}
	hcardShopModel := new(models.HcardShopModel).Init()
	var insertData []map[string]interface{}
	// 走全部适用逻辑
	if args.IsAllShop == true {
		for _, hcardID := range realHcardIDs {
			insertData = append(insertData, map[string]interface{}{
				hcardShopModel.Field.F_hcard_id: hcardID,
				hcardShopModel.Field.F_bus_id:   busID,
				hcardShopModel.Field.F_shop_id:  0, // 0:代表全部适用子店
			})
		}
	} else { // 部分门店适用
		// 调用门店RPC服务，检查门店id是否合法
		rpcShop := new(bus.Shop).Init()
		defer rpcShop.Close()
		checkArgs := &bus2.ArgsCheckShop{
			BusId:   busID,
			ShopIds: realShopIds,
		}
		checkReply := &bus2.ReplyCheckShop{}
		err = rpcShop.CheckBusShop(ctx, checkArgs, checkReply)
		if err != nil {
			return
		}
		for _, hcardID := range realHcardIDs {
			for _, shopID := range realShopIds {
				insertData = append(insertData, map[string]interface{}{
					hcardShopModel.Field.F_hcard_id: hcardID,
					hcardShopModel.Field.F_bus_id:   busID,
					hcardShopModel.Field.F_shop_id:  shopID,
				})
			}
		}
	}
	// 将旧数据删除重新插入
	if len(insertData) > 0 {
		hcardShopModel.Model.Begin()
		// 将之前的旧数据删除(例如:之前hcardID=1适用1号门店,但是现在hcardID=1适用于2号门店不再适用于1号门店,此时就需要将之前的数据删除,插入新的数据)
		if !hcardShopModel.DelByHcardIDs(realHcardIDs) {
			hcardShopModel.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		if result := hcardShopModel.InsertAll(insertData); result <= 0 {
			hcardShopModel.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		hcardShopModel.Model.Commit()
	}

	// 这部分负责更新子店中已经存在的卡项状态

	// 一期优化----yinjinlin-2021-04-08
	// 某张卡，之前是适用的门店并且门店已经添加，后被总店修改为不适用，
	// 需要更新shopHcardModel中stauts为：被总店禁用

	var disableIDs, downIDs []int
	shopHcardModel := new(models.ShopHcardModel).Init()
	shopHcards := shopHcardModel.GetByHcardIDs(realHcardIDs) // 根据限时卡id找到所有已经添加到门店的数据
	// 如果是设置成部分门店适用，需要把在门店"已上架"的但是现在改为不适用的门店限次卡改为"总店禁用"
	if args.IsAllShop {
		for _, hcard := range dataArr {
			// hcardStatus, _ := strconv.Atoi(hcard[hcardModel.Field.F_is_ground].(string)) // 总店中卡的状态
			isDel, _ := strconv.Atoi(hcard[hcardModel.Field.F_is_del].(string)) // 总店中没有被删除
			hcardID, _ := strconv.Atoi(hcard[hcardModel.Field.F_hcard_id].(string))
			for _, shopHcard := range shopHcards {
				shopStatus, _ := strconv.Atoi(shopHcard[shopHcardModel.Field.F_status].(string)) // 子店中卡的状态
				id, _ := strconv.Atoi(shopHcard[shopHcardModel.Field.F_id].(string))
				shopHcardID, _ := strconv.Atoi(shopHcard[shopHcardModel.Field.F_hcard_id].(string)) // 子店中卡的id
				if hcardID != shopHcardID {
					continue
				}
				// if shopStatus == cards.STATUS_DISABLE && hcardStatus == cards.SINGLE_IS_GROUND_yes { // 子店中卡的状态为禁用并且总店卡的状态为上架时,子店才可以上架
				//	downIDs = append(downIDs, id)
				// }

				// 一期优化----yinjinlin-2021-04-08
				// 全部适用：之前门店已经添加过并且分： 1>.上架/或下架 状态 ：都不需要更新
				//                                2>.被总店禁用 状态 ：需要更新为下架状态
				if isDel == cards.IS_BUS_DEL_no && shopStatus == cards.STATUS_DISABLE {
					downIDs = append(downIDs, id)
				}
			}
		}
	} else {
		for _, hcard := range dataArr {
			// hcardStatus, _ := strconv.Atoi(hcard[hcardModel.Field.F_is_ground].(string)) // 总店中卡的状态
			isDel, _ := strconv.Atoi(hcard[hcardModel.Field.F_is_del].(string))
			hcardID, _ := strconv.Atoi(hcard[hcardModel.Field.F_hcard_id].(string))
			for _, shopHcard := range shopHcards {
				shopStatus, _ := strconv.Atoi(shopHcard[shopHcardModel.Field.F_status].(string)) // 子店中卡的状态
				id, _ := strconv.Atoi(shopHcard[shopHcardModel.Field.F_id].(string))
				// shopID, _ := strconv.Atoi(shopHcard[shopHcardModel.Field.F_shop_id].(string))       // 已经存在的子店id
				shopHcardID, _ := strconv.Atoi(shopHcard[shopHcardModel.Field.F_hcard_id].(string)) // 子店中卡的id
				if hcardID != shopHcardID {
					continue
				}
				// if functions.InArray(shopID, realShopIds) { // 以前已经添加过的子店并且status="总店禁用";现在适用时需要将status恢复为下架;前提为总店卡的状态为上架时
				//	if hcardStatus == cards.SINGLE_IS_GROUND_yes && shopStatus == cards.STATUS_DISABLE {
				//		downIDs = append(downIDs, id)
				//	}
				//
				//
				// } else { // 以前已经添加过的子店并且status为"上架"或"下架";如果现在不再适用现时卡时，status需更改为"总店禁用"
				//	if shopStatus != cards.STATUS_DISABLE {
				//		disableIDs = append(disableIDs, id)
				//	}
				//
				// }

				// 部分适用
				// 以前添加过到门店，并且门店现在是上架/下架 状态，总店设置不适用了，需要将门店状态改为：被总店禁用
				if shopStatus != cards.STATUS_DISABLE && isDel == cards.IS_BUS_DEL_no {
					disableIDs = append(disableIDs, id)
				}
			}
		}
	}
	if len(disableIDs) > 0 {
		_ = shopHcardModel.UpdateByIDs(disableIDs, map[string]interface{}{
			shopHcardModel.Field.F_status:     cards.STATUS_DISABLE,
			shopHcardModel.Field.F_under_time: time.Now().Unix(),
		})
		// 添加维护es的shop-item文档的任务
		setShopItem(ctx, []int{}, 0, cards.ITEM_TYPE_hcard, disableIDs)
	}
	if len(downIDs) > 0 {
		_ = shopHcardModel.UpdateByIDs(downIDs, map[string]interface{}{
			shopHcardModel.Field.F_status: cards.STATUS_OFF_SALE,
		})
	}
	return
}

// DownUpHcard 总店上下架限时卡 (一期优化废用）
func (h *HcardLogic) DownUpHcard(ctx context.Context, args *cards.ArgsDownUpHcard) (err error) {
	// 检查商户信息
	busID := 0
	busID, err = checkBus(args.BsToken, true)
	if err != nil {
		return
	}
	// 检查限时卡ids是否为空
	if len(args.HcardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	// 检查限时卡是否有重复数据
	realHcardIDs := functions.ArrayUniqueInt(args.HcardIDs)
	if len(realHcardIDs) != len(args.HcardIDs) {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	// 遍历检查所选限时卡是否属于当前的商户所拥有
	hcardModel := new(models.HcardModel).Init()
	hcardDatas := hcardModel.GetHcardByIDs(realHcardIDs, hcardModel.Field.F_bus_id, hcardModel.Field.F_hcard_id, hcardModel.Field.F_is_ground)
	busStr := strconv.Itoa(busID)
	if len(hcardDatas) != len(realHcardIDs) { // 校验是否有不属于总店的现时卡
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	for _, hcardData := range hcardDatas {
		if busStr != hcardData[hcardModel.Field.F_bus_id].(string) {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
	}
	// 将id中已经上架的限时卡id和未上架的限次卡id分开处理
	var hcardDescList []struct {
		HcardID  int `mapstructure:"hcard_id"`
		IsGround int `mapstructure:"is_ground"`
	}
	var downIDs, upIDs []int
	_ = mapstructure.WeakDecode(hcardDatas, &hcardDescList)
	for _, hcardDesc := range hcardDescList {
		if hcardDesc.IsGround == cards.IS_GROUND_no { // 已经下架的限时卡ids
			downIDs = append(downIDs, hcardDesc.HcardID)
		} else {
			upIDs = append(upIDs, hcardDesc.HcardID)
		}
	}
	shopHcardModel := new(models.ShopHcardModel).Init()
	// 处理限时卡下架操作：总店status下架对应的门店status设置为禁用
	if args.OptType == cards.STATUS_OFF_SALE && len(upIDs) > 0 {
		if result := hcardModel.UpdateByHcardIDs(upIDs, map[string]interface{}{
			hcardModel.Field.F_is_ground:     cards.IS_GROUND_no,
			hcardModel.Field.F_under_time:    time.Now().Unix(),
			hcardModel.Field.F_sale_shop_num: 0,
		}); !result { // 更新总店的限时卡的status
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		if result := shopHcardModel.UpdateByHcardIDs(upIDs, map[string]interface{}{
			shopHcardModel.Field.F_status:     cards.STATUS_DISABLE,
			shopHcardModel.Field.F_under_time: time.Now().Unix(),
		}); !result { // 更新已经已添加到门店的限时卡的status
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		// 添加维护es的shop-item文档的任务
		setShopItem(ctx, []int{}, 0, cards.ITEM_TYPE_hcard, upIDs)
	}

	// 总店status上架，门店status由禁用->下架
	if args.OptType == cards.STATUS_ON_SALE && len(downIDs) > 0 {
		// 总店限时卡上架
		if !hcardModel.UpdateByHcardIDs(downIDs, map[string]interface{}{
			hcardModel.Field.F_is_ground:  cards.IS_GROUND_yes,
			hcardModel.Field.F_under_time: 0,
		}) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		// 将已添加到子店的且子店适用的限时卡解除总店禁用状态
		// 更新子店的禁用status
		if !shopHcardModel.UpdateByHcardIDs(downIDs, map[string]interface{}{
			shopHcardModel.Field.F_status: cards.STATUS_OFF_SALE,
		}) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	return
}

// ShopGetBusHcardPage 子店获取适用本店的限时卡列表
func (h *HcardLogic) ShopGetBusHcardPage(ctx context.Context, busID, shopID, start, limit int) (list cards.ReplyShopGetBusHcardPage, err error) {
	if busID <= 0 || shopID < 0 || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}

	// 获取适用于本店的限时卡数据
	hcardShopModel := new(models.HcardShopModel).Init()
	hcardShops := hcardShopModel.GetPageByShopID(busID, shopID, start, limit)
	if len(hcardShops) == 0 {
		return
	}
	hcardIDs := functions.ArrayValue2Array(hcardShopModel.Field.F_hcard_id, hcardShops) // 适用店中的限时卡ids

	// 获取限时卡数据
	hcardModel := new(models.HcardModel).Init()
	hcards := hcardModel.GetHcardByIDs(hcardIDs)
	if len(hcards) == 0 {
		return
	}
	// 获取子店已添加的限时卡数据方便取子店的状态
	shopHcardModel := new(models.ShopHcardModel).Init()
	shopHcards := shopHcardModel.GetByShopIDAndHcardIDs(shopID, hcardIDs)
	shopCardIds := functions.ArrayValue2Array(shopHcardModel.Field.F_hcard_id, shopHcards)
	list.List = make([]cards.HcardDesc, len(hcards))
	list.TotalNum = hcardShopModel.GetNumByBusIDAndShopID(busID, shopID) // 总量
	// 返回值拼装
	for index, hcard := range hcards {
		_ = mapstructure.WeakDecode(hcard, &list.List[index].HcardBase)
		_ = mapstructure.WeakDecode(hcard, &list.List[index])
		list.List[index].CtimeStr = time.Unix(int64(list.List[index].Ctime), 0).Format("2006/01/02 15:04:05")
		list.List[index].Clicks, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.HCARD_CLIKS, hcard[hcardModel.Field.F_hcard_id].(string)))
		cardId, _ := strconv.Atoi(hcard[hcardModel.Field.F_hcard_id].(string))
		shopHasAdd := 0
		if functions.InArray(cardId, shopCardIds) {
			shopHasAdd = 1
		}
		list.List[index].ShopHasAdd = shopHasAdd
		for _, shopHcard := range shopHcards {
			shopHcardID, _ := strconv.Atoi(shopHcard[shopHcardModel.Field.F_hcard_id].(string))
			if list.List[index].HcardID == shopHcardID { // 已经添加过的限时卡
				list.List[index].ShopHasAdd = 1
				status, _ := strconv.Atoi(shopHcard[shopHcardModel.Field.F_status].(string))
				list.List[index].ShopStatus = status
			}
		}
	}
	// 获取图片信息
	imageIDs := functions.ArrayValue2Array(hcardModel.Field.F_img_id, hcards)
	list.IndexImg = getImgsByImgIds(ctx, imageIDs, cards.ITEM_TYPE_hcard)
	return
}

// ShopHcardPage 子店限时卡列表
func (h *HcardLogic) ShopHcardPage(ctx context.Context, args *cards.ArgsShopHcardPage) (list cards.ReplyHcardPage, err error) {
	// 参数检查
	start, limit, shopID := args.GetStart(), args.GetPageSize(), args.ShopID
	if (args.ShopID <= 0 && args.ShopCall) || start < 0 || limit <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	// 根据子店铺id获取已添加到门店的限时卡 (已经过滤掉删除的限时限次卡)
	shopHcardModel := new(models.ShopHcardModel).Init()
	shopHcards := shopHcardModel.GetPageByShopID(args.Status, shopID, start, limit)
	hcardIDs := functions.ArrayValue2Array(shopHcardModel.Field.F_hcard_id, shopHcards) // 限时卡ids
	// 根据上面获取的限时卡id查询对应的限时卡数据
	hcardModel := new(models.HcardModel).Init()
	hcards := hcardModel.GetHcardByIDs(hcardIDs)
	list.List = make([]cards.HcardDesc, 0)
	if len(hcards) <= 0 {
		return
	}
	// 获取不同卡项-适用单项目的个数和赠送单项目的个数
	gaagsNumMap := GetApplyAndGiveSingleNum(hcardIDs, cards.ITEM_TYPE_hcard)

	list.List = make([]cards.HcardDesc, len(hcards))
	for index, hcard := range hcards {
		_ = mapstructure.WeakDecode(hcard, &list.List[index])
		_ = mapstructure.WeakDecode(hcard, &list.List[index].HcardBase)
		list.List[index].CtimeStr = time.Unix(int64(list.List[index].Ctime), 0).Format("2006/01/02 15:04:05")
		list.List[index].Clicks, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.HCARD_CLIKS, hcard[hcardModel.Field.F_hcard_id].(string)))
		list.List[index].ApplySingleNum = gaagsNumMap[list.List[index].HcardID].ApplySingleNum
		list.List[index].GiveSingleNum = gaagsNumMap[list.List[index].HcardID].GiveSingleNum
		for _, shopHcard := range shopHcards {
			shopHcardID, _ := strconv.Atoi(shopHcard[shopHcardModel.Field.F_hcard_id].(string))
			if list.List[index].HcardID == shopHcardID { // 已经添加过的限时卡

				list.List[index].ShopItemId, _ = strconv.Atoi(shopHcard[shopHcardModel.Field.F_id].(string))
				list.List[index].ShopHasAdd = 1
				status, _ := strconv.Atoi(shopHcard[shopHcardModel.Field.F_status].(string))
				list.List[index].ShopStatus = status
			}
		}

	}
	// 获取图片信息
	imgIDs := functions.ArrayValue2Array(hcardModel.Field.F_img_id, hcards)
	list.IndexImg = getImgsByImgIds(ctx, imgIDs, cards.ITEM_TYPE_hcard)
	// 获取数量(已经过滤掉删除的限时限次卡)
	list.TotalNum = shopHcardModel.GetNumByShopID(shopID, args.Status)
	return
}

// ShopAddHcard 子店添加总部限时卡到门店 （一期优化，去掉总店推送，改为门店自动拉取）
func (h *HcardLogic) ShopAddHcard(ctx context.Context, args *cards.ArgsShopAddHcard) (err error) {
	// hcardID去重
	args.HcardIDs = functions.ArrayUniqueInt(args.HcardIDs)
	if len(args.HcardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	// checkBus/Shop
	var busID, shopID int
	if busID, err = checkBus(args.BsToken); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if shopID, err = checkShop(args.BsToken); err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	//
	if shopID <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}

	// 检查限时卡是否适用当前子店
	// hcardShopModel := new(models.HcardShopModel).Init() // 限时卡的适用子店
	// hcardShopArr := hcardShopModel.GetByShopIDAndHcardIDs(busID, shopID, args.HcardIDs)
	// if len(hcardShopArr) != len(args.HcardIDs) { // 适用子店的量要和添加的量一致
	//	err = toolLib.CreateKcErr(_const.NO_IN_SHOP)
	//	return
	// }

	// 提取门店已经添加过的卡项
	shopHcardModel := new(models.ShopHcardModel).Init()
	shopHcardLists := shopHcardModel.GetByShopIDAndHcardIDs(shopID, args.HcardIDs)
	shopHcardExistIDs := functions.ArrayValue2Array(shopHcardModel.Field.F_hcard_id, shopHcardLists) // 已经添加过的ids

	// 刷选出已经添加过并且删除的数据
	delHcardIdSlice := make([]int, 0)
	for _, hcardMap := range shopHcardLists {
		isDel, _ := strconv.Atoi(hcardMap[shopHcardModel.Field.F_is_del].(string))
		if isDel == cards.IS_BUS_DEL_yes {
			delHcardId, _ := strconv.Atoi(hcardMap[shopHcardModel.Field.F_hcard_id].(string))
			delHcardIdSlice = append(delHcardIdSlice, delHcardId)
		}
	}

	// 更新门店之前添加过并删除的数据
	if len(delHcardIdSlice) > 0 {
		// 更新数据删除和上下架状态
		if updateBool := shopHcardModel.UpdateDelByHcardIdsAndshopId(delHcardIdSlice, shopID, map[string]interface{}{
			shopHcardModel.Field.F_is_del: cards.IS_BUS_DEL_no,
			shopHcardModel.Field.F_status: cards.STATUS_OFF_SALE,
		}); !updateBool {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
		//更新卡项关联表
		sirModel := new(models.ShopItemRelationModel).Init()
		if b := sirModel.UpdateByItemIdsAndShopId(delHcardIdSlice, cards.ITEM_TYPE_hcard, shopID, map[string]interface{}{
			sirModel.Field.F_is_del: cards.IS_BUS_DEL_no,
			sirModel.Field.F_status: cards.STATUS_OFF_SALE,
		}); !b {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}

	// 实际要添加的限时卡id
	realAddIDs := make([]int, 0)
	for _, hcardID := range args.HcardIDs {
		if !functions.InArray(hcardID, shopHcardExistIDs) {
			realAddIDs = append(realAddIDs, hcardID)
		}
	}

	// 校验当前门店是否已经将卡项内涉及到的单项目添加到自己的门店内
	allSingle, singleIds, err := new(ItemLogic).getItemCardIncSingleIds(realAddIDs, cards.ITEM_TYPE_hcard)
	if err != nil {
		return
	}
	if err = new(ItemLogic).validShopSingleContainItemCardSingles(shopID, busID, allSingle, singleIds); err != nil {
		return
	}

	// 插入门店选择总店的限时卡数据
	var addShData []map[string]interface{}
	shopItemRelationData := make([]map[string]interface{}, 0)
	shopItemRelationModel := new(models.ShopItemRelationModel).Init()
	for _, realHacadId := range realAddIDs {
		status := cards.STATUS_OFF_SALE // 添加到门店，默认下架状态
		ctime := time.Now().Local().Unix()
		addShData = append(addShData, map[string]interface{}{
			shopHcardModel.Field.F_hcard_id: realHacadId,
			shopHcardModel.Field.F_status:   status,
			shopHcardModel.Field.F_shop_id:  shopID,
			shopHcardModel.Field.F_ctime:    ctime,
		})
		shopItemRelationData = append(shopItemRelationData, map[string]interface{}{
			shopItemRelationModel.Field.F_item_id:   realHacadId,
			shopItemRelationModel.Field.F_item_type: cards.ITEM_TYPE_hcard,
			shopItemRelationModel.Field.F_status:    cards.STATUS_OFF_SALE,
			shopItemRelationModel.Field.F_shop_id:   shopID,
			shopItemRelationModel.Field.F_is_del:    cards.ITEM_IS_DEL_NO,
		})
	}
	//
	if len(addShData) > 0 {
		if id := shopHcardModel.InsertAll(addShData); id < 0 {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}
	// 门店卡项关联表数据插入
	if len(shopItemRelationData) > 0 {
		if shopItemRelationModel.InsertAll(shopItemRelationData) <= 0 {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	// 初始化适用门店模型
	hsModel := new(models.HcardShopModel).Init()
	hsInfo := hsModel.GetByHcardIDs(args.HcardIDs)
	hcardIds := functions.ArrayValue2Array(hsModel.Field.F_hcard_id, hsInfo)
	// 过滤掉已经添加适用门店数据(实际添加）
	addHsHcardIds := make([]int, 0)
	for _, hcardId := range args.HcardIDs {
		if functions.InArray(hcardId, hcardIds) == false {
			addHsHcardIds = append(addHsHcardIds, hcardId)
		}
	}

	var addhcardShopData []map[string]interface{} // 添加适用门店表的数据
	for _, addHcardId := range addHsHcardIds {
		addhcardShopData = append(addhcardShopData, map[string]interface{}{
			hsModel.Field.F_hcard_id: addHcardId,
			hsModel.Field.F_shop_id:  shopID,
			hsModel.Field.F_bus_id:   busID,
		})
	}
	// 过滤的数据添加到适用限时卡表
	if len(addhcardShopData) > 0 {
		if id := hsModel.InsertAll(addhcardShopData); id < 0 {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}

	// 获取总店限时卡上下架状态
	// hcardModel := new(models.HcardModel).Init()
	// //hcardArrs := hcardModel.GetHcardByIDs(realAddIDs, hcardModel.Field.F_hcard_id, hcardModel.Field.F_is_ground)
	// hcardArrs := hcardModel.GetHcardByIDs(realAddIDs, hcardModel.Field.F_hcard_id, hcardModel.Field.F_is_del)
	// hcardMap := functions.ArrayRebuild(hcardModel.Field.F_hcard_id, hcardArrs)
	// var addData []map[string]interface{}
	// for _, realAddID := range realAddIDs {
	//	status := cards.STATUS_OFF_SALE
	//	if _, ok := hcardMap[strconv.Itoa(realAddID)]; ok {
	//		if hcard, ok := hcardMap[strconv.Itoa(realAddID)].(map[string]interface{}); ok {
	//			//isGround, _ := strconv.Atoi(hcard[hcardModel.Field.F_is_ground].(string))
	//			//if isGround == cards.IS_GROUND_no {
	//			//	status = cards.STATUS_DISABLE
	//			//}
	//			isDel, _ := strconv.Atoi(hcard[hcardModel.Field.F_is_del].(string))
	//			if isDel == cards.IS_BUS_DEL_no {
	//				status = cards.STATUS_OFF_SALE
	//			}
	//		}
	//	}
	//	ctime := time.Now().Local().Unix()
	//	addData = append(addData, map[string]interface{}{
	//		shopHcardModel.Field.F_hcard_id: realAddID,
	//		shopHcardModel.Field.F_status:   status,
	//		shopHcardModel.Field.F_shop_id:  shopID,
	//		shopHcardModel.Field.F_ctime:    ctime,
	//	})
	// }
	//
	// if len(addData) > 0 && (shopHcardModel.InsertAll(addData)) <= 0 {
	//	err = toolLib.CreateKcErr(_const.DB_ERR)
	//	return
	// }
	return
}

// 子店删除限时卡
func (h *HcardLogic) ShopDeleteHcard(ctx context.Context, args *cards.ArgsDeleteHcard) (err error) {
	shopId, _ := checkShop(args.BsToken)
	shopHcardModle := new(models.ShopHcardModel).Init()
	shopHcardlist := shopHcardModle.FindHcardIdsAndBusId(args.HcardIds, shopId, shopHcardModle.Field.F_shop_id)
	// 检测选择删除的限时卡是否存在
	if len(args.HcardIds) != len(shopHcardlist) {
		err = toolLib.CreateKcErr(_const.NO_IN_BUS)
		return
	}
	r := shopHcardModle.UpdateDelByHcardIdsAndshopId(args.HcardIds, shopId, map[string]interface{}{
		shopHcardModle.Field.F_is_del:   cards.IS_BUS_DEL_yes,
		shopHcardModle.Field.F_del_time: time.Now().Unix(),
	})
	if r == false {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	// 同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.HcardIds, cards.ITEM_TYPE_hcard, shopId) {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	return
}

// ShopDownUpHcard 子店上下架自己店铺中的限时卡
func (h *HcardLogic) ShopDownUpHcard(ctx context.Context, args *cards.ArgsShopDownUpHcard) (err error) {
	// 验证门店
	shopID, err := checkShop(args.BsToken)
	if err != nil || shopID <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	args.ShopHcardIDs = functions.ArrayUniqueInt(args.ShopHcardIDs)
	if len(args.ShopHcardIDs) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	// 获取门店限时卡信息
	shopHcardModel := new(models.ShopHcardModel).Init()
	shopHcardModel.Model.Begin()
	shopHcards := shopHcardModel.GetByShopIDAndHcardIDs(shopID, args.ShopHcardIDs)
	var shopHcardStruct []struct {
		ID      int
		ShopID  int
		Status  int
		HcardId int
	}
	_ = mapstructure.WeakDecode(shopHcards, &shopHcardStruct)
	var downIDs, upIDs, hcardIds []int
	for _, shopHcardDesc := range shopHcardStruct {
		if shopID != shopHcardDesc.ShopID {
			err = toolLib.CreateKcErr(_const.PARAM_ERR)
			return
		}
		if shopHcardDesc.Status == cards.STATUS_OFF_SALE {
			downIDs = append(downIDs, shopHcardDesc.ID)
			hcardIds = append(hcardIds, shopHcardDesc.HcardId)
		} else {
			upIDs = append(upIDs, shopHcardDesc.ID)
			hcardIds = append(hcardIds, shopHcardDesc.HcardId)
		}
	}

	hcardModel := new(models.HcardModel).Init(shopHcardModel.Model.GetOrmer())
	var decOrInc string
	// 下架
	if args.OptType == cards.STATUS_OFF_SALE && len(upIDs) > 0 {
		if !shopHcardModel.UpdateByIDs(upIDs, map[string]interface{}{
			shopHcardModel.Field.F_status: cards.STATUS_OFF_SALE}) {
			shopHcardModel.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		// 同步下架门店卡项关联表数据
		shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
		if !shopItemRealtionModel.UpdateStatusByItemIds(hcardIds, cards.ITEM_TYPE_hcard, cards.STATUS_OFF_SALE, shopID) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		decOrInc = "dec"
	}
	// 上架
	if args.OptType == cards.STATUS_ON_SALE && len(downIDs) > 0 {
		if !shopHcardModel.UpdateByIDs(downIDs, map[string]interface{}{
			shopHcardModel.Field.F_status: cards.STATUS_ON_SALE}) {
			shopHcardModel.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		// 同步上架门店卡项关联表数据
		shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
		if !shopItemRealtionModel.UpdateStatusByItemIds(hcardIds, cards.ITEM_TYPE_hcard, cards.STATUS_ON_SALE, shopID) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		decOrInc = "inc"
	}
	if len(decOrInc) > 0 {
		//	更新总店中对应现时卡的在售门店数量
		if !hcardModel.UpdateSaleShopNum(hcardIds, decOrInc) {
			shopHcardModel.Model.RollBack()
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	shopHcardModel.Model.Commit()
	if args.OptType == cards.STATUS_OFF_SALE && len(upIDs) > 0 {
		// 添加维护es的shop-item文档的任务
		setShopItem(ctx, []int{}, shopID, cards.ITEM_TYPE_hcard, upIDs)
	}
	if args.OptType == cards.STATUS_ON_SALE && len(downIDs) > 0 {
		// 添加维护es的shop-item文档的任务
		setShopItem(ctx, []int{}, shopID, cards.ITEM_TYPE_hcard, downIDs)
	}

	return
}

// AddOrEditHcardExt 新增/修改现时卡拓展信息数据
func (h *HcardLogic) AddOrEditHcardExt(hcardModel *models.HcardModel, hcardID int, notes []cards.CardNote, isAdd bool /*true:新增;false:修改*/) (err error) {
	if hcardID == 0 {
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
	cardExt := new(models.HcardExtModel).Init(hcardModel.Model.GetOrmer())
	mapParams := map[string]interface{}{
		cardExt.Field.F_hcard_id: hcardID,
		cardExt.Field.F_notes:    string(notesStr),
	}
	if isAdd && cardExt.Insert(mapParams) <= 0 {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	if !isAdd {
		smExMap := cardExt.GetByHcardID(hcardID)
		if len(smExMap) > 0 && !cardExt.UpdateByHcardID(hcardID, mapParams) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		} else if len(smExMap) == 0 && cardExt.Insert(mapParams) <= 0 {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	return
}

// GetHcardExitByHcardID 获取现时卡拓展信息详情
func (h *HcardLogic) GetHcardExitByHcardID(hcardID int, reply *[]cards.CardNote) {
	if hcardID == 0 {
		return
	}
	cardExt := new(models.HcardExtModel).Init()
	dataMap := cardExt.GetByHcardID(hcardID)
	if len(dataMap) > 0 {
		data := dataMap[cardExt.Field.F_notes].(string)
		json.Unmarshal([]byte(data), reply)
	}
	return
}

// ShopHcardListRpc 子店现时卡列表rpc内部调用
func (h *HcardLogic) ShopHcardListRpc(ctx context.Context, args *cards.ArgsShopHcardListRpc, reply *cards.ReplyShopHcardListRpc) (err error) {
	hcardShopModel := new(models.ShopHcardModel).Init()
	listMap := hcardShopModel.GetByShopIDAndHcardIDs(args.ShopId, args.HcardIds)
	if len(listMap) == 0 {
		return
	}
	hcardIDs := functions.ArrayValue2Array(hcardShopModel.Field.F_hcard_id, listMap) // 限时卡ids
	// 根据上面获取的限时卡id查询对应的限时卡数据
	hcardModel := new(models.HcardModel).Init()
	hcards := hcardModel.GetHcardByIDs(hcardIDs)
	reply.List = make([]cards.HcardDesc, len(hcards))
	for index, hcard := range hcards {
		_ = mapstructure.WeakDecode(hcard, &reply.List[index])
		_ = mapstructure.WeakDecode(hcard, &reply.List[index].HcardBase)
		for _, shopHcard := range listMap {
			shopHcardID, _ := strconv.Atoi(shopHcard[hcardShopModel.Field.F_hcard_id].(string))
			if reply.List[index].HcardID == shopHcardID { // 已经添加过的限时卡
				reply.List[index].ShopHasAdd = 1
				status, _ := strconv.Atoi(shopHcard[hcardShopModel.Field.F_status].(string))
				reply.List[index].ShopStatus = status
			}
		}
	}
	return
}

// checkHcardData 检查限时卡的入参数据
func (h *HcardLogic) checkHcardData(busID int, hcardBase cards.HcardBase, incSingles []cards.IncInfSingle, giveSingles []cards.HcardSingle) (err error) {
	// 验证限时卡名称
	if err = cards.VerfiyName(hcardBase.Name); err != nil {
		return
	}
	// 验证限时卡的原价和现价
	if err = cards.VerfiyPrice(hcardBase.RealPrice, hcardBase.Price); err != nil {
		return
	}
	// 验证限时卡的有效期
	if err = cards.VerfiyServicePeriod(hcardBase.ServicePeriod); err != nil {
		return
	}
	// 验证限时卡包含的单项目的数量
	if err = cards.VerifySinglesNum(len(incSingles)); err != nil {
		return
	}
	// 验证限时卡赠送的单项目的数量
	if err = cards.VerifyGiveSinglesNum(len(giveSingles)); err != nil {
		return
	}
	return
}

// 获取套餐/限次/限时/限时限次卡中包含单项目的总次数
func GetItemCardSingleNum(cardId int, cardType int) (singleTotalNum int) {
	cardMaps := make([]map[string]interface{}, 0)
	switch cardType {
	case cards.ITEM_TYPE_sm:
		ssModel := new(models.SmSingleModel).Init()
		cardMaps = ssModel.GetBySmid(cardId)
	// case cards.ITEM_TYPE_hcard:
	//	hsModel := new(models.HcardSingleModel).Init()
	//	cardMaps = hsModel.GetByHcardID(cardId)
	case cards.ITEM_TYPE_ncard:
		ncModel := new(models.NCardSingleModel).Init()
		cardMaps = ncModel.GetByNCardID(cardId)
	case cards.ITEM_TYPE_hncard:
		hncModel := new(models.HNCardSingleModel).Init()
		cardMaps = hncModel.GetByHNCardID(cardId)
	}
	for _, v := range cardMaps {
		num, _ := strconv.Atoi(v["num"].(string))
		singleTotalNum += num
	}
	return
}
