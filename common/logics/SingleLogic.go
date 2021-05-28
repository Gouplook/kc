package logics

import (
	"context"
	"encoding/json"
	"fmt"
	"git.900sui.cn/kc/kcgin/logs"
	"git.900sui.cn/kc/rpcCards/common/tools"
	"sort"
	"strconv"
	"strings"
	"time"

	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/mapstructure"
	"git.900sui.cn/kc/redis"
	"git.900sui.cn/kc/rpcCards/common/models"
	"git.900sui.cn/kc/rpcCards/constkey"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/client/bus"
	cards2 "git.900sui.cn/kc/rpcinterface/client/elastic/cards"
	"git.900sui.cn/kc/rpcinterface/client/file"
	"git.900sui.cn/kc/rpcinterface/client/staff"
	bus2 "git.900sui.cn/kc/rpcinterface/interface/bus"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	file2 "git.900sui.cn/kc/rpcinterface/interface/file"
	staff2 "git.900sui.cn/kc/rpcinterface/interface/staff"
	redis2 "github.com/gomodule/redigo/redis"
	"github.com/shopspring/decimal"
)

type SingleLogic struct {
}

const (
	SINGLE_SET_TYPE_ADD  uint8 = 1
	SINGLE_SET_TYPE_EDIT uint8 = 2
)

type singleImg struct {
	IndexImgId   int   // 封面图片id
	SubImgIds    []int // 相册图片id
	EffectImgIds []int // 服务效果图片id
	ToolImgIds   []int // 仪器设备图片id
}

type specPrice struct {
	MinPrice float64
	MaxPrice float64
}

// 获取单个单项目-提供给rpc
func (s *SingleLogic) GetSimpleSingleInfo(ctx context.Context, singleId int) *map[string]interface{} {
	mSingle := new(models.SingleModel).Init()
	singleInfo := mSingle.GetBySingleId(singleId)
	if len(singleInfo) == 0 {
		data := make(map[string]interface{})
		return &data
	}
	return &singleInfo
}

// 获取单个单项目-提供给rpc
func (s *SingleLogic) GetSimpleSingleInfos(ctx context.Context, singleIds []int) *[]map[string]interface{} {
	var data []map[string]interface{}
	mSingle := new(models.SingleModel).Init()
	singleInfos := mSingle.GetBySingleids(singleIds, []string{
		mSingle.Field.F_single_id,
		mSingle.Field.F_name,
		mSingle.Field.F_has_spec,
	})
	if len(singleInfos) == 0 {
		return &data
	}
	return &singleInfos
}

// 获取单项目门店价格
// @param ssId  单项目在门店的id@kc_shop_single.ss_id
func (s *SingleLogic) GetShopSinglePrice(ssId int, sspId int) (float64, error) {
	shopSingleModel := new(models.ShopSingleModel).Init()
	shopSingleInfo := shopSingleModel.GetShopSingleById(ssId)
	if len(shopSingleInfo) == 0 {
		// 单项目不存在
		return 0, toolLib.CreateKcErr(_const.SHOP_SINGLE_NOEXIST)
	}
	status, err := strconv.Atoi(shopSingleInfo[shopSingleModel.Field.F_status].(string))
	if err != nil {
		return 0, err
	}
	if status != models.STATUS_ON_SALE {
		// 项目已下架
		return 0, toolLib.CreateKcErr(_const.SHOP_SINGLE_OFF_SALE)
	}

	if sspId > 0 {
		mSSP := new(models.SingleSpecPriceModel).Init()
		singleSpecPrice := mSSP.GetSingleSpecPriceById(sspId)
		if len(singleSpecPrice) == 0 {
			return 0, toolLib.CreateKcErr(_const.SHOP_SINGLE_SPEC_NOEXIST)
		}

		if singleSpecPrice[mSSP.Field.F_is_del].(string) == strconv.Itoa(cards.SPEC_PRICE_IS_DEL_yes) {
			return 0, toolLib.CreateKcErr(_const.SHOP_SINGLE_SPEC_NOEXIST)
		}
		// 获取子店铺的价格
		shopSingleSpecPriceModel := new(models.ShopSingleSpecPriceModel).Init()
		shopSingleSpecPriceInfo := shopSingleSpecPriceModel.GetBySsidAndSspid(ssId, sspId)

		if len(shopSingleSpecPriceInfo) != 0 {
			shopPrice, _ := strconv.ParseFloat(shopSingleSpecPriceInfo[shopSingleSpecPriceModel.Field.F_price].(string), 64)
			return shopPrice, nil
		}
		// 规格不存在
		return 0, toolLib.CreateKcErr(_const.SHOP_SINGLE_SPEC_NOEXIST)
	}
	shopPrice, _ := strconv.ParseFloat(shopSingleInfo[shopSingleModel.Field.F_changed_real_price].(string), 64)
	return shopPrice, nil
}

// 添加单项目
func (s *SingleLogic) AddSingle(ctx context.Context, singleInfo *cards.ArgsAddSingle) (singleId int, err error) {
	singleId = 0
	// 获取商户id
	busId, err := checkBus(singleInfo.BsToken, true)
	if err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}

	singImg, spPrice, err := s.checkSingleInfo(ctx, &singleInfo.SingleBase, busId)
	if err != nil {
		return
	}

	var hasSpec int = cards.SINGLE_HAS_SPEC_no
	if len(singleInfo.SpecIds) > 0 {
		hasSpec = cards.SINGLE_HAS_SPEC_yes
	}
	// 添加单项目主表数据
	mSingle := new(models.SingleModel).Init()
	singleId = mSingle.Insert(map[string]interface{}{
		mSingle.Field.F_bind_id:      singleInfo.BindId,
		mSingle.Field.F_bus_id:       busId,
		mSingle.Field.F_name:         singleInfo.Name,
		mSingle.Field.F_tag_ids:      functions.Implode(",", singleInfo.TagIds),
		mSingle.Field.F_sort_desc:    singleInfo.SortDesc,
		mSingle.Field.F_service_time: singleInfo.ServiceTime,
		mSingle.Field.F_ctime:        time.Now().Local().Unix(),
		mSingle.Field.F_real_price:   singleInfo.RealPrice,
		mSingle.Field.F_price:        singleInfo.Price,
		mSingle.Field.F_has_spec:     hasSpec,
		mSingle.Field.F_img_id:       singImg.IndexImgId,
		// mSingle.Field.F_is_ground:    cards.SINGLE_IS_GROUND_yes,
		mSingle.Field.F_min_price: spPrice.MinPrice,
		mSingle.Field.F_max_price: spPrice.MaxPrice,
		mSingle.Field.F_subscribe: singleInfo.Subscribe,
	})
	if singleId == 0 {
		err = toolLib.CreateKcErr(_const.DB_ERR)
	}
	// 添加扩展信息表 （一期优化，属性标签删除）
	mSE := new(models.SingleExtModel).Init()
	reminder, _ := json.Marshal(singleInfo.Reminder)
	contextStr, _ := json.Marshal(singleInfo.SingleContext)
	mSE.Insert(map[string]interface{}{
		mSE.Field.F_single_id: singleId,
		// mSE.Field.F_age_bracket:      functions.Implode(",", singleInfo.AgeBracket),
		// mSE.Field.F_sex:              singleInfo.Sex,
		// mSE.Field.F_tailor_indus:     singleInfo.TailorIndus,
		// mSE.Field.F_tailor_sub_indus: functions.Implode(",", singleInfo.TailorSubIndus),

		mSE.Field.F_service_reminder: string(reminder),
		mSE.Field.F_single_context:   string(contextStr),
	})

	// 处理图片数据
	s.setSinglePic(singImg, singleId, SINGLE_SET_TYPE_ADD)
	// 处理规格
	s.setSingleSpec(singleInfo.SpecIds, singleInfo.SpecPrices, singleId, SINGLE_SET_TYPE_ADD)
	// 添加风控统计任务
	new(ItemLogic).AddXCardTask(ctx, singleId, cards.ITEM_TYPE_single)
	return
}

// 获取单项目信息
func (s *SingleLogic) SingleInfo(ctx context.Context, singleId int, uid int, shopId ...int) (sinfo cards.SingleDetail, err error) {
	// 获取单项目基本信息
	mSingle := new(models.SingleModel).Init()
	singleInfo := mSingle.GetBySingleId(singleId)
	if len(singleInfo) == 0 {
		err = toolLib.CreateKcErr(_const.SINGLE_NO_INFO)
		return
	}
	baseInfo := struct {
		BusId       int
		Sales       int
		Name        string
		SortDesc    string
		BindId      int
		TagIds      string
		RealPrice   float64
		Price       float64
		ServiceTime int
		MinPrice    float64
		MaxPrice    float64
		ImgId       int
		HasSpec     int
		Subscribe   string
	}{}
	mapstructure.WeakDecode(singleInfo, &baseInfo)
	// 获取扩展字段
	mSE := new(models.SingleExtModel).Init()
	extInfo := mSE.GetBySingleid(singleId)
	var reminder []cards.ReminderInfo
	var singleContext []cards.SingleContextInfo
	// extIntInfo := []struct {
	// 	//Sex             int
	// 	//TailorIndus     int
	// 	ServiceReminder string
	// 	SingleContext   string
	// 	//AgeBracket      string
	// 	//TailorSubIndus  string [{}.{}]
	// }{}
	// mapstructure.WeakDecode(extInfo, &extIntInfo)
	// reminder := map[string]interface{}{}
	// singleContext := map[string]interface{}{}

	if err := json.Unmarshal([]byte(extInfo[mSE.Field.F_service_reminder].(string)), &reminder); err != nil {
	}
	if err := json.Unmarshal([]byte(extInfo[mSE.Field.F_single_context].(string)), &singleContext); err != nil {
	}
	// if _, ok := reminder["use_instrument"]; ok {
	//	reminder["use_instrument"] = []string{reminder["use_instrument"].(string)}
	// }

	// 获取图片信息
	var imgIds = []int{}
	mSI := new(models.SingleImgModel).Init()
	imgs := mSI.GetBySingleId(singleId)
	imgIds = functions.ArrayValue2Array(mSI.Field.F_img_id, imgs)
	imgIds = append(imgIds, baseInfo.ImgId)
	rpcFile := new(file.Upload).Init()
	defer rpcFile.Close()
	imgsRep := map[int]file2.ReplyFileInfo{}
	err = rpcFile.GetImageByIds(ctx, imgIds, &imgsRep)
	var singleImg = cards.SingleImg{
		IndexPic:   cards.Img{},
		AlbumPics:  []cards.Img{},
		EffectPics: []cards.Img{},
		ToolsPics:  []cards.Img{},
	}
	if err == nil && len(imgsRep) > 0 {
		if _, ok := imgsRep[baseInfo.ImgId]; ok {
			singleImg.IndexPic = cards.Img{
				Hash: imgsRep[baseInfo.ImgId].Hash,
				Url:  imgsRep[baseInfo.ImgId].Path,
			}
		}
		// 默认封面图片
		if baseInfo.ImgId == 0 {
			singleImg.IndexPic = cards.Img{
				Hash: "",
				Url:  constkey.CardsSmallDefaultPics[cards.ITEM_TYPE_single],
			}
		}
		var imgsStruct = []struct {
			ImgId int
			Type  int
		}{}
		mapstructure.WeakDecode(imgs, &imgsStruct)
		for _, img := range imgsStruct {
			if _, ok := imgsRep[img.ImgId]; ok {
				switch img.Type {
				case cards.IMG_TYPE_album:
					singleImg.AlbumPics = append(singleImg.AlbumPics, cards.Img{
						Hash: imgsRep[img.ImgId].Hash,
						Url:  imgsRep[img.ImgId].Path,
					})
				case cards.IMG_TYPE_effect:
					singleImg.EffectPics = append(singleImg.EffectPics, cards.Img{
						Hash: imgsRep[img.ImgId].Hash,
						Url:  imgsRep[img.ImgId].Path,
					})
				case cards.IMG_TYPE_tool:
					singleImg.ToolsPics = append(singleImg.ToolsPics, cards.Img{
						Hash: imgsRep[img.ImgId].Hash,
						Url:  imgsRep[img.ImgId].Path,
					})
				}
			}
		}
	}
	// 标签属性 (一期优化，标签属性删除）
	// ageBracketIds := functions.StrExplode2IntArr(extIntInfo.AgeBracket, ",")
	var ageBracketStr []string
	ageBracketStr = []string{}
	// for _, ageid := range ageBracketIds {
	//	ageBracketStr = append(ageBracketStr, cards.AgeBracket[ageid])
	// }
	// 订制子类
	// tailorSubIds := functions.StrExplode2IntArr(extIntInfo.TailorSubIndus, ",")
	var tailorSubStr []string
	tailorSubStr = []string{}
	// for _, tsubid := range tailorSubIds {
	//	tailorSubStr = append(tailorSubStr, cards.TailorSub[extIntInfo.TailorIndus][tsubid])
	// }

	// 获取规格信息
	var specIds = []cards.SingleSpecIds{}
	singleSpecName := map[int]cards.SingleSpec{}

	specPrices := []cards.SingleSpecPrice{}
	var ssId int = 0
	var onSale int = 0
	var sspStruct = []struct {
		SpecIds string
		Price   float64
		SspId   int
	}{}
	if baseInfo.HasSpec == cards.SINGLE_HAS_SPEC_yes {
		mSS := new(models.SingleSpecModel).Init()
		spec := mSS.GetBySingleid(singleId)
		if len(spec) > 0 {
			json.Unmarshal([]byte(spec[mSS.Field.F_spec_info].(string)), &specIds)
		}
		allSpecId := []int{}
		if len(specIds) > 0 {
			for _, v := range specIds {
				allSpecId = append(allSpecId, v.SpecId)
				allSpecId = append(allSpecId, v.Sub...)
			}
		}
		if len(allSpecId) > 0 {
			mSBS := new(models.SingleBusSpecModel).Init()
			allSpec := mSBS.GetBySpecids(allSpecId)
			allSpecStruct := []cards.SingleSpec{}
			mapstructure.WeakDecode(allSpec, &allSpecStruct)
			for _, v := range allSpecStruct {
				singleSpecName[v.SpecId] = v
			}
		}

		// 获取不同规格的价格
		mSSP := new(models.SingleSpecPriceModel).Init()
		ssp := mSSP.GetBySingleid(singleId)
		mapstructure.WeakDecode(ssp, &sspStruct)
	}
	// 如果有shopid，使用店铺价格替换商家设置的价格
	var shopSpecPriceMap = map[int]float64{}
	var shareLink string
	var shopid = 0
	if len(shopId) > 0 && shopId[0] > 0 {
		mShopSingle := new(models.ShopSingleModel).Init()
		shopSingle := mShopSingle.GetByShopidAndSingleid(shopId[0], singleId)
		if len(shopSingle) == 0 {
			err = toolLib.CreateKcErr(_const.SINGLE_NO_INFO)
			return
		}
		if len(shopSingle) > 0 {
			shopSingleStruct := struct {
				SsId             int
				Status           int
				Sales            int
				ChangedRealPrice float64
				ChangedMinPrice  float64
				ChangedMaxPrice  float64
			}{}
			mapstructure.WeakDecode(shopSingle, &shopSingleStruct)
			baseInfo.RealPrice = shopSingleStruct.ChangedRealPrice
			baseInfo.MinPrice = shopSingleStruct.ChangedMinPrice
			baseInfo.MaxPrice = shopSingleStruct.ChangedMaxPrice
			ssId = shopSingleStruct.SsId
			if shopSingleStruct.Status == models.STATUS_ON_SALE {
				onSale = 1
			}
			baseInfo.Sales = shopSingleStruct.Sales
			mSSSP := new(models.ShopSingleSpecPriceModel).Init()
			shopSpecPrice := mSSSP.GetBySsid(shopSingleStruct.SsId)
			shopSpecPriceStruct := []struct {
				SspId int
				Price float64
			}{}
			mapstructure.WeakDecode(shopSpecPrice, &shopSpecPriceStruct)
			for _, v := range shopSpecPriceStruct {
				shopSpecPriceMap[v.SspId] = v.Price
			}
		}
		shopid = shopId[0]
	}

	for _, v := range sspStruct {
		sspPrice := v.Price
		if _, ok := shopSpecPriceMap[v.SspId]; ok {
			sspPrice = shopSpecPriceMap[v.SspId]
		}
		specPrices = append(specPrices, cards.SingleSpecPrice{
			SpecIds: functions.StrExplode2IntArr(v.SpecIds, ","),
			Price:   sspPrice,
			SspId:   v.SspId,
		})
	}

	// 标签名称
	var tagsStr []cards.TagInfo
	tagIds := functions.StrExplode2IntArr(baseInfo.TagIds, ",")
	if len(tagIds) > 0 {
		mTag := new(models.TagModel).Init()
		tagr := mTag.GetByTagids(tagIds)
		mapstructure.WeakDecode(tagr, &tagsStr)
	}

	// 获取单项目的会员折扣信息
	var memberLevel = cards.MemberLev{
		Level:     0,
		LevelName: "",
		Rebate:    0, // 无折扣
	}

	if uid > 0 {
		rpcMember := new(bus.Member).Init()
		margs := bus2.ArgsMemberRebate{
			Busid: baseInfo.BusId,
			Uid:   uid,
		}
		var userLevel = bus2.BusLevelDetail{}
		_ = rpcMember.GetMemberRebateByUid(ctx, &margs, &userLevel)

		if userLevel.Level > 0 {
			memberLevel.Level = userLevel.Level
			memberLevel.LevelName = userLevel.Name

			if userLevel.ServiceType == bus2.RebateTypeAll {
				// 全部
				memberLevel.Rebate, _ = decimal.NewFromInt(int64(userLevel.ServiceRebate)).Div(decimal.NewFromInt(10)).Truncate(2).Float64()
			} else if userLevel.ServiceType == bus2.RebateTypePart {
				// 部分
				for _, rs := range userLevel.RebateServices {
					if rs.ServiceId == singleId {
						memberLevel.Rebate, _ = decimal.NewFromInt(int64(rs.Rebate)).Div(decimal.NewFromInt(10)).Truncate(2).Float64()
						break
					}
				}
			}
		}
	}

	shareLink = tools.GetShareLink(singleId, shopid, cards.ITEM_TYPE_single)
	// 点击数量加1
	redis.RedisGlobMgr.Hincrby(constkey.SINGLE_CLICKS, strconv.Itoa(singleId), 1)

	sinfo = cards.SingleDetail{
		SingleId: singleId,
		SingleBase: cards.SingleBase{
			Name:        baseInfo.Name,
			BusId:       baseInfo.BusId,
			Sales:       baseInfo.Sales,
			SortDesc:    baseInfo.SortDesc,
			BindId:      baseInfo.BindId,
			TagIds:      tagIds,
			RealPrice:   baseInfo.RealPrice,
			Price:       baseInfo.Price,
			MinPrice:    baseInfo.MinPrice,
			MaxPrice:    baseInfo.MaxPrice,
			ServiceTime: baseInfo.ServiceTime,
			Subscribe:   baseInfo.Subscribe,
			// Sex:            extIntInfo.Sex,
			// AgeBracket:     ageBracketIds,
			// TailorIndus:    extIntInfo.TailorIndus,
			// TailorSubIndus: tailorSubIds,SpecPrices
			Reminder:      reminder,
			SingleContext: singleContext,
			SpecIds:       specIds,
			SpecPrices:    specPrices, // 不同规格的价格
			Pictures:      []string{},
			EffectImgs:    []string{},
			ToolsImgs:     []string{},
		},
		ShareLink: shareLink,
		SingleImg: singleImg,
		// SexStr:        cards.SexBracket[extIntInfo.Sex],
		AgeBracketStr: ageBracketStr,
		// TailorStr:     cards.Tailor[extIntInfo.TailorIndus],
		TailorSubStr: tailorSubStr,
		Specs:        singleSpecName,
		TagsStr:      tagsStr,
		SsId:         ssId,
		OnSale:       onSale,
		MemberLev:    memberLevel,
		IsShop:       cards.IS_USE_SHOP_YES,
	}

	return
}

// 编辑单项目信息
func (s *SingleLogic) EditSingle(ctx context.Context, singleInfo *cards.ArgsEditSingle) (err error) {
	// 获取商户id
	busId, err := checkBus(singleInfo.BsToken, true)
	if err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	// 检查项目信息是否存在
	mSingle := new(models.SingleModel).Init()
	single := mSingle.GetBySingleId(singleInfo.SingleId, []string{mSingle.Field.F_bus_id})
	if len(single) == 0 {
		err = toolLib.CreateKcErr(_const.SINGLE_NO_INFO)
		return
	}
	// 检查项目是否属于商家
	singleBusId, _ := strconv.Atoi(single[mSingle.Field.F_bus_id].(string))
	if singleBusId != busId {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	singImg, spPrice, err := s.checkSingleInfo(ctx, &singleInfo.SingleBase, busId)
	if err != nil {
		return
	}
	var hasSpec int = cards.SINGLE_HAS_SPEC_no
	if len(singleInfo.SpecIds) > 0 {
		hasSpec = cards.SINGLE_HAS_SPEC_yes
	}
	mSSP := new(models.SingleSpecPriceModel).Init()
	mSSSP := new(models.ShopSingleSpecPriceModel).Init()
	var specPricesMap []map[string]interface{}
	var hashs = map[int]string{}
	var specIdsMap = map[int]string{}
	for k, sp := range singleInfo.SpecPrices {
		sort.Ints(sp.SpecIds)
		specIds := functions.Implode(",", sp.SpecIds)
		hashStr := functions.HashMd5(fmt.Sprintf("%s|%d", specIds, singleInfo.SingleId))
		hashs[k] = hashStr
		specIdsMap[k] = specIds
		specPricesMap = append(specPricesMap, map[string]interface{}{
			mSSP.Field.F_single_id: singleInfo.SingleId,
			mSSP.Field.F_price:     sp.Price,
			mSSP.Field.F_spec_ids:  specIds,
			mSSP.Field.F_hash:      hashStr,
		})
	}

	prices := mSSP.GetBySingleid(singleInfo.SingleId)
	pricesStruct := []struct {
		SspId int
		Hash  string
		Price float64
	}{}
	mapstructure.WeakDecode(prices, &pricesStruct)

	// 需要设置为删除的sspid
	var delSspids []int
	for _, v := range pricesStruct {
		hasd := 0
		for k, _ := range singleInfo.SpecPrices {
			if v.Hash == hashs[k] {
				hasd = 1
				break
			}
		}
		if hasd == 0 {
			delSspids = append(delSspids, v.SspId)
		}
	}

	if len(delSspids) > 0 {
		ssModel := new(models.SmSingleModel).Init()
		if count := ssModel.Count([]base.WhereItem{{ssModel.Field.F_ssp_id, []interface{}{"IN", delSspids}}}); count > 0 {
			return toolLib.CreateKcErr(_const.DEL_SINGLE_SPEC_ERROR)
		}
		ncsModel := new(models.NCardSingleModel).Init()
		if count := ncsModel.Count([]base.WhereItem{{ncsModel.Field.F_ssp_id, []interface{}{"IN", delSspids}}}); count > 0 {
			return toolLib.CreateKcErr(_const.DEL_SINGLE_SPEC_ERROR)
		}
		hncsModel := new(models.HNCardSingleModel).Init()
		if count := hncsModel.Count([]base.WhereItem{{hncsModel.Field.F_ssp_id, []interface{}{"IN", delSspids}}}); count > 0 {
			return toolLib.CreateKcErr(_const.DEL_SINGLE_SPEC_ERROR)
		}

		mSSP.UpdateBySspids(delSspids, map[string]interface{}{
			mSSP.Field.F_is_del: cards.SPEC_PRICE_IS_DEL_yes,
		})

		mSSSP.UpdateBySspids(delSspids, map[string]interface{}{
			mSSSP.Field.F_is_del: cards.SPEC_PRICE_IS_DEL_yes,
		})
	}

	// 修改基本信息
	mSingle.UpdateBySingleid(singleInfo.SingleId, map[string]interface{}{
		mSingle.Field.F_bind_id:      singleInfo.BindId,
		mSingle.Field.F_bus_id:       busId,
		mSingle.Field.F_name:         singleInfo.Name,
		mSingle.Field.F_tag_ids:      functions.Implode(",", singleInfo.TagIds),
		mSingle.Field.F_sort_desc:    singleInfo.SortDesc,
		mSingle.Field.F_service_time: singleInfo.ServiceTime,
		mSingle.Field.F_real_price:   singleInfo.RealPrice,
		mSingle.Field.F_price:        singleInfo.Price,
		mSingle.Field.F_has_spec:     hasSpec,
		mSingle.Field.F_img_id:       singImg.IndexImgId,
		// mSingle.Field.F_is_ground:    cards.SINGLE_IS_GROUND_yes,
		mSingle.Field.F_min_price: spPrice.MinPrice,
		mSingle.Field.F_max_price: spPrice.MaxPrice,
		mSingle.Field.F_subscribe: singleInfo.Subscribe,
	})
	// 扩展字段 （一期优化，属性标签删除）
	mSE := new(models.SingleExtModel).Init()
	reminder, _ := json.Marshal(singleInfo.Reminder)
	contextByte, _ := json.Marshal(singleInfo.SingleContext)
	mSE.UpdateBySingleid(singleInfo.SingleId, map[string]interface{}{
		// mSE.Field.F_age_bracket:      functions.Implode(",", singleInfo.AgeBracket),
		mSE.Field.F_service_reminder: string(reminder),
		mSE.Field.F_single_context:   string(contextByte),
		// mSE.Field.F_sex:              singleInfo.Sex,
		// mSE.Field.F_tailor_indus:     singleInfo.TailorIndus,
		// mSE.Field.F_tailor_sub_indus: functions.Implode(",", singleInfo.TailorSubIndus),
	})

	// 处理图片数据
	s.setSinglePic(singImg, singleInfo.SingleId, SINGLE_SET_TYPE_EDIT)

	mSS := new(models.SingleSpecModel).Init()
	ssm := new(models.ShopSingleModel).Init()
	if len(singleInfo.SpecIds) == 0 {
		// 修改 规格都去除了，删除规格和规格价格
		mSS.DeleteBySingleid(singleInfo.SingleId)
		mSSP.UpdateBySingleid(singleInfo.SingleId, map[string]interface{}{
			mSSP.Field.F_is_del: cards.SPEC_PRICE_IS_DEL_yes,
		})

		// 门店删除规格
		shopInfos := ssm.GetBySingleid(singleInfo.SingleId)
		ssIds := functions.ArrayValue2Array(ssm.Field.F_ss_id, shopInfos)
		if len(ssIds) > 0 {
			if b := mSSSP.DelBySsids(ssIds); !b {
				return toolLib.CreateKcErr(_const.DB_ERR)
			}
		}
	}
	specInfo, _ := json.Marshal(singleInfo.SpecIds)
	ssInfo := mSS.GetBySingleid(singleInfo.SingleId)
	if len(ssInfo) > 0 {
		ssid, _ := strconv.Atoi(ssInfo[mSS.Field.F_id].(string))
		mSS.UpdateById(ssid, map[string]interface{}{
			mSS.Field.F_spec_info: string(specInfo),
		})
	} else {
		mSS.Insert(map[string]interface{}{
			mSS.Field.F_single_id: singleInfo.SingleId,
			mSS.Field.F_spec_info: string(specInfo),
		})
	}

	type updateDataStruct struct {
		SspId int
		Price float64
	}
	// 需要同步到子店的新增规格的组合id极价格 sspid=>price
	var addSspids = map[int]float64{}
	var updateData []updateDataStruct
	for k, sp := range singleInfo.SpecPrices {
		hasd := 0
		for _, v := range pricesStruct {
			if hashs[k] == v.Hash {
				hasd = 1
				if v.Price != sp.Price {
					updateData = append(updateData, updateDataStruct{
						SspId: v.SspId,
						Price: sp.Price,
					})
				}
			}
		}
		if hasd == 0 {
			sspid := mSSP.Insert(map[string]interface{}{
				mSSP.Field.F_single_id: singleInfo.SingleId,
				mSSP.Field.F_price:     sp.Price,
				mSSP.Field.F_spec_ids:  specIdsMap[k],
				mSSP.Field.F_hash:      hashs[k],
			})
			addSspids[sspid] = sp.Price
		}
	}

	if len(addSspids) > 0 {
		// 同步规格价格到子店
		// 先获取已添加此单项目的子店的ssid
		// mShopSingle := new(models.ShopSingleModel).Init()
		shopSingles := ssm.GetBySingleid(singleInfo.SingleId)

		if len(shopSingles) > 0 {
			// mSSSP := new(models.ShopSingleSpecPriceModel).Init()
			var addShopSpecPriceData = []map[string]interface{}{}
			for _, shopSingle := range shopSingles {
				for k, addprice := range addSspids {
					addShopSpecPriceData = append(addShopSpecPriceData, map[string]interface{}{
						mSSSP.Field.F_ss_id:   shopSingle[ssm.Field.F_ss_id].(string),
						mSSSP.Field.F_ssp_id:  k,
						mSSSP.Field.F_price:   addprice,
						mSSSP.Field.F_shop_id: shopSingle[ssm.Field.F_shop_id].(string),
					})
				}
			}
			if len(addShopSpecPriceData) > 0 {
				mSSSP.InsertAll(addShopSpecPriceData)
			}
		}
	}
	if len(updateData) > 0 {
		for _, v := range updateData {
			mSSP.UpdateBySspid(v.SspId, map[string]interface{}{
				mSSP.Field.F_price: v.Price,
			})
		}
	}

	// 处理规格
	// s.setSingleSpec(singleInfo.SpecIds, singleInfo.SpecPrices, singleInfo.SingleId, SINGLE_SET_TYPE_EDIT)

	return
}

// 获取商家的单项目列表
func (s *SingleLogic) GetBusSingles(ctx context.Context, busId, shopId, start, limit int, isGround, isDel string, filterShopHasAdd bool) (replys cards.ReplyBusSinglePage, err error) {
	mSingle := new(models.SingleModel).Init()
	where := []base.WhereItem{
		{mSingle.Field.F_bus_id, busId},
	}
	if isGround != "" {
		isground, _ := strconv.Atoi(isGround)
		isground = isground - 1
		where = append(where, base.WhereItem{mSingle.Field.F_is_ground, isGround})
	}
	if isDel != "" {
		isdel, _ := strconv.Atoi(isDel)
		where = append(where, base.WhereItem{mSingle.Field.F_is_del, isdel})
	}

	// 子店已添加的卡项
	var shopAddSingles []map[string]interface{}
	scModel := new(models.ShopSingleModel).Init()
	if shopId > 0 {
		shopAddSingles = scModel.SelectRcardsByWherePage([]base.WhereItem{{scModel.Field.F_shop_id, shopId},{scModel.Field.F_is_del, cards.IS_BUS_DEL_no}}, 0, 0)
		if filterShopHasAdd && len(shopAddSingles) > 0 {
			shopHasAddSingleIds := functions.ArrayValue2Array(scModel.Field.F_single_id, shopAddSingles)
			where = append(where, base.WhereItem{mSingle.Field.F_single_id, []interface{}{"NOT IN", shopHasAddSingleIds}})
		}
	}
	singles := mSingle.SelectSinglesByWherePage(where, start, limit)
	totalNum := mSingle.GetNumByWhere(where)

	singlesStruct := []cards.ListSingle{}
	for index, single := range singles {
		for _, shopCard := range shopAddSingles {
			if single[mSingle.Field.F_single_id].(string) == shopCard[scModel.Field.F_single_id].(string) { // 表明当前子店已添加该卡项
				singles[index]["ShopItemId"], _ = strconv.Atoi(shopCard[scModel.Field.F_ss_id].(string))
				singles[index]["ShopStatus"], _ = strconv.Atoi(shopCard[scModel.Field.F_status].(string))
				singles[index]["ShopHasAdd"] = 1
				singles[index]["ShopDelStatus"], _ = strconv.Atoi(shopCard[scModel.Field.F_is_del].(string))
				break
			}
		}
	}
	mapstructure.WeakDecode(singles, &singlesStruct)
	for k, v := range singlesStruct {
		singlesStruct[k].Click, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.SINGLE_CLICKS, strconv.Itoa(v.SingleId)))
		singlesStruct[k].CtimeStr = functions.TimeToStr(int64(singlesStruct[k].Ctime))
	}
	replys = cards.ReplyBusSinglePage{
		TotalNum:  totalNum,
		List:      singlesStruct,
		TagNames:  map[int]string{},
		IndexImgs: map[int]string{},
	}

	// 获取标签和图片数据
	replys.TagNames, replys.IndexImgs = s.getSinglesTagsAndPics(ctx, singlesStruct)

	return
}

// 获取子店的单项目
func (s *SingleLogic) GetShopSingles(ctx context.Context, shopId, start, limit int, singIds []int, status, isDel string) (replys cards.ReplyShopSinglePage, err error) {
	mShopSingle := new(models.ShopSingleModel).Init()
	shopSingles := mShopSingle.GetByShopid(shopId, start, limit, status, isDel, singIds)
	replys = cards.ReplyShopSinglePage{
		TotalNum:  0,
		List:      []cards.ListSingle{},
		TagNames:  map[int]string{},
		IndexImgs: map[int]string{},
	}

	if len(shopSingles) == 0 {
		return
	}
	var list = []cards.ListSingle{}
	var shopSingleStruct = []struct {
		SsId             int
		SingleId         int
		ChangedRealPrice float64
		ChangedMinPrice  float64
		ChangedMaxPrice  float64
		Status           int
		IsDel            int
		Sales            int
		Name             string
	}{}
	mapstructure.WeakDecode(shopSingles, &shopSingleStruct)

	singleIds := functions.ArrayValue2Array(mShopSingle.Field.F_single_id, shopSingles)
	// 获取单项目数据
	mSingle := new(models.SingleModel).Init()
	singles := mSingle.GetBySingleids(singleIds)
	singlesStruct := []cards.ListSingle{}
	mapstructure.WeakDecode(singles, &singlesStruct)
	singlesStructMap := map[int]cards.ListSingle{}
	for _, v := range singlesStruct {
		singlesStructMap[v.SingleId] = v
	}

	for k, v := range shopSingleStruct {
		list = append(list, singlesStructMap[v.SingleId])
		list[k].Click, _ = redis2.Int(redis.RedisGlobMgr.Hget(constkey.SINGLE_CLICKS, strconv.Itoa(v.SingleId)))
		list[k].SsId = v.SsId
		list[k].IsDel = v.IsDel
		list[k].ShopStatus = v.Status
		list[k].ShopDelStatus = v.IsDel
		if len(v.Name) > 0 {
			list[k].Name = v.Name
		}
		list[k].RealPrice = v.ChangedRealPrice
		list[k].MinPrice = v.ChangedMinPrice
		list[k].MaxPrice = v.ChangedMaxPrice
		list[k].Sales = v.Sales
		list[k].CtimeStr = functions.TimeToStr(int64(list[k].Ctime))
	}

	replys.TagNames, replys.IndexImgs = s.getSinglesTagsAndPics(ctx, list)
	replys.List = list
	// 获取总数量
	replys.TotalNum = 0
	if len(replys.List) > 0 {
		replys.TotalNum = mShopSingle.GetNumByShopid(shopId, status, isDel, singIds)
	}
	return
}

// 获取单项目列表的数据
func (s *SingleLogic) getSinglesTagsAndPics(ctx context.Context, singlesStruct []cards.ListSingle) (tags map[int]string, pics map[int]string) {

	var tagIds, imgIds []int
	for k, v := range singlesStruct {
		if len(v.TagIds) > 0 {
			singlesStruct[k].TagIdsArr = functions.StrExplode2IntArr(v.TagIds, ",")
			tagIds = append(tagIds, singlesStruct[k].TagIdsArr...)
		}
		if v.ImgId > 0 {
			imgIds = append(imgIds, v.ImgId)
		}
	}

	tags = make(map[int]string)
	pics = make(map[int]string)

	// 获取标签
	if len(tagIds) > 0 {
		mTag := new(models.TagModel).Init()
		tagsInfo := mTag.GetByTagids(tagIds)
		tagsStruct := []struct {
			TagId int
			Name  string
		}{}
		mapstructure.WeakDecode(tagsInfo, &tagsStruct)
		for _, v := range tagsStruct {
			tags[v.TagId] = v.Name
		}
	}
	// 默认封面图片
	pics[0] = constkey.CardsSmallDefaultPics[cards.ITEM_TYPE_single]
	// 获取封面图片
	if len(imgIds) > 0 {
		rpcImg := new(file.Upload).Init()
		defer rpcImg.Close()
		var replyImgs = map[int]file2.ReplyFileInfo{}
		err := rpcImg.GetImageByIds(ctx, imgIds, &replyImgs)
		if err != nil {
			return
		}
		for _, v := range replyImgs {
			pics[v.Id] = v.Path
		}
	}
	return
}

// 子店添加单项目
func (s *SingleLogic) ShopAddSingle(singleInfo *cards.ArgsShopAddSingle) (err error) {
	shopId, err := singleInfo.GetShopId()
	if err != nil || shopId <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if len(singleInfo.SingleId) <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	busId, _ := singleInfo.GetBusId()
	// 获取单项目信息
	mSingle := new(models.SingleModel).Init()
	single := mSingle.GetBySingleids(singleInfo.SingleId, []string{
		mSingle.Field.F_single_id,
		mSingle.Field.F_bus_id,
		mSingle.Field.F_name,
		mSingle.Field.F_is_ground,
		mSingle.Field.F_is_del,
		mSingle.Field.F_has_spec,
		mSingle.Field.F_real_price,
		mSingle.Field.F_min_price,
		mSingle.Field.F_max_price,
	})
	if len(single) == 0 {
		err = toolLib.CreateKcErr(_const.SINGLE_NO_INFO)
		return
	}
	var singleStruct = []struct {
		SingleId  int
		BusId     int
		IsGround  int
		HasSpec   int
		RealPrice float64
		MinPrice  float64
		MaxPrice  float64
		Name      string
	}{}
	mapstructure.WeakDecode(single, &singleStruct)
	for _, tmp := range singleStruct {
		if tmp.BusId != busId {
			err = toolLib.CreateKcErr(_const.POWER_ERR)
			return
		}
		/*if tmp.IsGround == cards.SINGLE_IS_GROUND_no {
			err = toolLib.CreateKcErr(_const.GROUND_NO)
			return
		}*/
	}

	// 检查子店是否已经添加过
	mShopSingle := new(models.ShopSingleModel).Init()
	shopSingles := mShopSingle.GetByShopidAndSingleids(shopId, singleInfo.SingleId)
	shopHcardExistIDs := functions.ArrayValue2Array(mShopSingle.Field.F_single_id, shopSingles) // 已经添加过的ids
	delHcardIdSlice := make([]int, 0)
	if len(shopSingles) > 0 {
		// err = toolLib.CreateKcErr(_const.SHOP_SINGLE_HASD)
		// return

		// 刷选出已经添加过并且删除的数据
		for _, hcardMap := range shopSingles {
			isDel, _ := strconv.Atoi(hcardMap[mShopSingle.Field.F_is_del].(string))
			if isDel == cards.IS_BUS_DEL_yes {
				delHcardId, _ := strconv.Atoi(hcardMap[mShopSingle.Field.F_single_id].(string))
				delHcardIdSlice = append(delHcardIdSlice, delHcardId)
			}
		}
	}


	// 更新门店之前添加过并删除的数据
	if len(delHcardIdSlice) > 0 {
		// 更新数据删除和上下架状态
		if updateBool := mShopSingle.UpDateBySingleidsAndShopId(delHcardIdSlice, shopId, map[string]interface{}{
			mShopSingle.Field.F_is_del: cards.IS_BUS_DEL_no,
			mShopSingle.Field.F_status: cards.STATUS_OFF_SALE,
		}); !updateBool {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
		//更新卡项关联表
		sirModel := new(models.ShopItemRelationModel).Init()
		if b := sirModel.UpdateByItemIdsAndShopId(delHcardIdSlice, cards.ITEM_TYPE_single, shopId, map[string]interface{}{
			sirModel.Field.F_is_del: cards.IS_BUS_DEL_no,
			sirModel.Field.F_status: cards.STATUS_OFF_SALE,
		}); !b {
			return toolLib.CreateKcErr(_const.DB_ERR)
		}
	}


	var specSingleIds []int
	var singleIdSsid = make(map[int]int)
	shopItemRelationData := make([]map[string]interface{}, 0)
	shopItemRelationModel := new(models.ShopItemRelationModel).Init()
	for _, tmp := range singleStruct {
		if functions.InArray(tmp.SingleId, shopHcardExistIDs) {
			continue
		}
		ssid := mShopSingle.Insert(map[string]interface{}{
			mShopSingle.Field.F_single_id:          tmp.SingleId,
			mShopSingle.Field.F_shop_id:            shopId,
			mShopSingle.Field.F_ctime:              time.Now().Local().Unix(),
			mShopSingle.Field.F_changed_real_price: tmp.RealPrice,
			mShopSingle.Field.F_changed_min_price:  tmp.MinPrice,
			mShopSingle.Field.F_changed_max_price:  tmp.MaxPrice,
			mShopSingle.Field.F_status:             cards.STATUS_OFF_SALE,
			mShopSingle.Field.F_name:               tmp.Name,
		})

		if tmp.HasSpec == cards.SINGLE_HAS_SPEC_yes {
			specSingleIds = append(specSingleIds, tmp.SingleId)
			singleIdSsid[tmp.SingleId] = ssid
		}
		shopItemRelationData = append(shopItemRelationData, map[string]interface{}{
			shopItemRelationModel.Field.F_item_id:   tmp.SingleId,
			shopItemRelationModel.Field.F_item_type: cards.ITEM_TYPE_single,
			shopItemRelationModel.Field.F_status:    cards.STATUS_OFF_SALE,
			shopItemRelationModel.Field.F_shop_id:   shopId,
			shopItemRelationModel.Field.F_is_del:    cards.ITEM_IS_DEL_NO,
		})
	}

	// 门店卡项关联表数据插入
	if len(shopItemRelationData) > 0 {
		if shopItemRelationModel.InsertAll(shopItemRelationData) <= 0 {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	// 同步多属性价格数据
	if len(specSingleIds) > 0 {
		mSSP := new(models.SingleSpecPriceModel).Init()
		specPrices := mSSP.GetBySingleids(specSingleIds)
		if len(specPrices) == 0 {
			return
		}
		var specPricesStruct = []struct {
			SingleId int
			SspId    int
			Price    float64
		}{}

		mapstructure.WeakDecode(specPrices, &specPricesStruct)
		mSSSP := new(models.ShopSingleSpecPriceModel).Init()
		var shopSpecPrices = []map[string]interface{}{}
		for _, v := range specPricesStruct {
			shopSpecPrices = append(shopSpecPrices, map[string]interface{}{
				mSSSP.Field.F_price:   v.Price,
				mSSSP.Field.F_ssp_id:  v.SspId,
				mSSSP.Field.F_ss_id:   singleIdSsid[v.SingleId],
				mSSSP.Field.F_shop_id: shopId,
			})
		}
		if len(shopSpecPrices) > 0 {
			mSSSP.InsertAll(shopSpecPrices)
		}
	}
	return
}

// 子店铺设置单项目价格
func (s *SingleLogic) ShopChangePrice(args *cards.ArgsShopChangePrice) (err error) {
	// 获取shopId
	shopId, err := args.GetShopId()
	if err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if args.SingleId <= 0 || args.RealPrice <= 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	busId, _ := args.GetBusId()
	// 获取单项目信息
	mSingle := new(models.SingleModel).Init()
	single := mSingle.GetBySingleId(args.SingleId, []string{
		mSingle.Field.F_bus_id,
		mSingle.Field.F_is_ground,
		mSingle.Field.F_is_del,
		mSingle.Field.F_has_spec,
	})
	if len(single) == 0 { // 单项目不存在
		err = toolLib.CreateKcErr(_const.SINGLE_NO_INFO)
		return
	}
	singleStruct := struct { // 单项目信息转结构体
		BusId    int
		IsGround int
		HasSpec  int
	}{}
	mapstructure.WeakDecode(single, &singleStruct)
	// 单项目不是当前商家的
	if singleStruct.BusId != busId {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	// 单项目已被总店下架，子店不能修改价格
	if singleStruct.IsGround == cards.SINGLE_IS_GROUND_no {
		err = toolLib.CreateKcErr(_const.GROUND_NO)
		return
	}
	// 检查单项目是否已被子店添加
	mShopSingle := new(models.ShopSingleModel).Init()
	shopSingle := mShopSingle.GetByShopidAndSingleid(shopId, args.SingleId)
	if len(shopSingle) == 0 { // 子店未添加
		err = toolLib.CreateKcErr(_const.SHOP_NOT_ADD)
		return
	}
	shopSingleStruct := struct {
		SsId            int
		Status          int
		ChangedMinPrice float64
		ChangedMaxPrice float64
	}{}
	// 子店单项目信息转结构体
	mapstructure.WeakDecode(shopSingle, &shopSingleStruct)
	// 项目被总店禁用 不能修改价格
	if shopSingleStruct.Status == models.STATUS_DISABLE {
		err = toolLib.CreateKcErr(_const.GROUND_NO)
		return
	}
	// 子店的最大价格最小价格
	var minP, maxP float64 = 0, 0

	// 单项目存在规格
	if singleStruct.HasSpec == cards.SINGLE_HAS_SPEC_yes {
		// 开始处理不同规格价格
		// 参数的规格数据为空
		if len(args.SpecPrice) == 0 {
			err = toolLib.CreateKcErr(_const.SPEC_PRICE_LOSE)
			return
		}
		// 获取数据库单项目的规格
		mSSP := new(models.SingleSpecPriceModel).Init()
		specPrice := mSSP.GetBySingleid(args.SingleId)
		// 判断参数的规格数量和数据库的规格数量是否相等
		if len(specPrice) != len(args.SpecPrice) {
			err = toolLib.CreateKcErr(_const.SPEC_PRICE_LOSE)
			return
		}
		// 整理出单项目在数据库所有规定的id
		allSspids := functions.ArrayValue2Array(mSSP.Field.F_ssp_id, specPrice)
		// 判断参数提供的规格价格不为0 以及 参数提供的规格id必须包含在allSspids里面
		for _, v := range args.SpecPrice {
			if v.Price <= 0 || functions.InArray(v.SspId, allSspids) == false {
				err = toolLib.CreateKcErr(_const.SPEC_PRICE_LOSE)
				return
			}
		}
		// 获取单项目在店铺的规格价格数据
		mSSSP := new(models.ShopSingleSpecPriceModel).Init()
		shopSpecPrice := mSSSP.GetBySsid(shopSingleStruct.SsId)
		type shopSpecPriceStruct struct {
			Id    int
			Price float64
		}
		shopSpecPriceMap := map[int]shopSpecPriceStruct{}
		for _, v := range shopSpecPrice {
			vid, _ := strconv.Atoi(v[mSSSP.Field.F_id].(string))
			vsspid, _ := strconv.Atoi(v[mSSSP.Field.F_ssp_id].(string))
			vprice, _ := strconv.ParseFloat(v[mSSSP.Field.F_price].(string), 64)
			shopSpecPriceMap[vsspid] = shopSpecPriceStruct{
				Id:    vid,
				Price: vprice,
			}
		}

		// 需要新增的规格
		var addShopSpec = []map[string]interface{}{}
		// 需要修改的规格价格
		var updateShopSpec = map[int]float64{}
		// 对比计算出新增规格和修改规格
		for _, v := range args.SpecPrice {
			if v.Price < minP || minP == 0 {
				minP = v.Price
			}
			if v.Price > maxP || maxP == 0 {
				maxP = v.Price
			}
			if _, ok := shopSpecPriceMap[v.SspId]; ok {
				if shopSpecPriceMap[v.SspId].Price != v.Price {
					updateShopSpec[shopSpecPriceMap[v.SspId].Id] = v.Price
				}
			} else {
				addShopSpec = append(addShopSpec, map[string]interface{}{
					mSSSP.Field.F_price:   v.Price,
					mSSSP.Field.F_ssp_id:  v.SspId,
					mSSSP.Field.F_ss_id:   shopSingleStruct.SsId,
					mSSSP.Field.F_shop_id: shopId,
				})
			}
		}

		if len(addShopSpec) > 0 {
			mSSSP.InsertAll(addShopSpec)
		}

		if len(updateShopSpec) > 0 {
			for id, sprice := range updateShopSpec {
				mSSSP.UpdateById(id, map[string]interface{}{
					mSSSP.Field.F_price: sprice,
				})
			}
		}
	}

	updateData := map[string]interface{}{
		mShopSingle.Field.F_changed_real_price: args.RealPrice,
		mShopSingle.Field.F_changed_min_price:  minP,
		mShopSingle.Field.F_changed_max_price:  maxP,
	}
	args.Name = strings.TrimSpace(args.Name)
	if len(args.Name) > 0 {
		updateData[mShopSingle.Field.F_name] = args.Name
	}
	// 修改子店单项目主表
	mShopSingle.UpDateBySsid(shopSingleStruct.SsId, updateData)

	return
}

/*//总店设置单项目的上下架
func (s *SingleLogic) BusUpDownSingle(ctx context.Context, args *cards.ArgsDownUpSingle) (err error) {
	busId, err := checkBus(args.BsToken, true)
	if err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if len(args.SingleIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	//获取单项目信息
	mSingle := new(models.SingleModel).Init()
	singles := mSingle.GetBySingleids(args.SingleIds, []string{
		mSingle.Field.F_single_id,
		mSingle.Field.F_bus_id,
		mSingle.Field.F_is_ground,
	})
	singlesStruct := []struct {
		SingleId int
		BusId    int
		IsGround int
	}{}
	mapstructure.WeakDecode(singles, &singlesStruct)
	//判断参数中的单项目是否是此商家的
	var downIds, upIds []int
	for _, v := range singlesStruct {
		if v.BusId != busId {
			err = toolLib.CreateKcErr(_const.OPT_OTHER_BUS_ITEM)
			return
		}
		if v.IsGround == cards.SINGLE_IS_GROUND_yes {
			upIds = append(upIds, v.SingleId)
		} else {
			downIds = append(downIds, v.SingleId)
		}
	}

	mShopSingle := new(models.ShopSingleModel).Init()

	//批量下架
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		r := mSingle.UpdateBySingleids(upIds, map[string]interface{}{
			mSingle.Field.F_is_ground:     cards.SINGLE_IS_GROUND_no,
			mSingle.Field.F_under_time:    time.Now().Unix(),
			mSingle.Field.F_sale_shop_num: 0,
		})
		if r == false {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//设置子店铺的单项目状态为总店禁用
		mShopSingle.UpDateBySingleids(upIds, map[string]interface{}{
			mShopSingle.Field.F_status:     models.STATUS_DISABLE,
			mShopSingle.Field.F_under_time: time.Now().Unix(),
		})
		//添加维护es的shop-item文档的任务
		setShopItem(ctx, upIds, 0, cards.ITEM_TYPE_single)
	}

	//批量上架
	if args.OptType == cards.STATUS_ON_SALE && len(downIds) > 0 {
		r := mSingle.UpdateBySingleids(downIds, map[string]interface{}{
			mSingle.Field.F_is_ground:  cards.SINGLE_IS_GROUND_yes,
			mSingle.Field.F_under_time: 0,
		})
		if r == false {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
		//设置子店铺的单项目状态为下架,解除总店禁用状态
		mShopSingle.UpDateBySingleids(downIds, map[string]interface{}{
			mShopSingle.Field.F_status: models.STATUS_OFF_SALE,
		})
	}
	return
}*/

// 总店删除单项目
func (s *SingleLogic) DelSingle(ctx context.Context, args *cards.ArgsDelSingle) (err error) {
	if len(args.SingleIds) < 0 {
		return
	}
	// 初始化模型
	delSingleModel := new(models.SingleModel).Init()
	r := delSingleModel.UpdateBySingleids(args.SingleIds, map[string]interface{}{
		delSingleModel.Field.F_is_del:   cards.IS_BUS_DEL_yes,
		delSingleModel.Field.F_del_time: time.Now().Unix(),
	})
	if r == false {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	// 初始化分店模型
	shopSingleModel := new(models.ShopSingleModel).Init()
	r2 := shopSingleModel.UpDateBySingleids(args.SingleIds, map[string]interface{}{
		shopSingleModel.Field.F_is_del:   cards.IS_BUS_DEL_yes,
		shopSingleModel.Field.F_del_time: time.Now().Unix(),
	})
	if r2 == false {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	// 同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.SingleIds, cards.ITEM_TYPE_single) {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}

	return
}

// 分店 删除项目
func (s *SingleLogic) DelShopSingle(ctx context.Context, shopId int, args *cards.ArgsDelSingle) (err error) {
	// 初始化分店模型
	shopSingleModel := new(models.ShopSingleModel).Init()
	r := shopSingleModel.UpDateBySingleidsAndShopId(args.SingleIds, shopId, map[string]interface{}{
		shopSingleModel.Field.F_is_del:   cards.IS_BUS_DEL_yes,
		shopSingleModel.Field.F_del_time: time.Now().Unix(),
	})
	if r == false {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	// 同步删除门店卡项关联表数据
	shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
	if !shopItemRealtionModel.DelByItemIds(args.SingleIds, cards.ITEM_TYPE_single, shopId) {
		err = toolLib.CreateKcErr(_const.DB_ERR)
		return
	}
	return err
}

// 子店铺设置单项目上下架
func (s *SingleLogic) ShopDownUpSingle(ctx context.Context, args *cards.ArgsShopDownUpSingle) (err error) {
	shopId, err := args.GetShopId()
	if err != nil || shopId <= 0 {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	if len(args.SsIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	mShopSingle := new(models.ShopSingleModel).Init()
	shopSingles := mShopSingle.GetBySsids(args.SsIds)
	shopSinglesStruct := []struct {
		SsId     int
		ShopId   int
		Status   int
		SingleId int
	}{}
	mapstructure.WeakDecode(shopSingles, &shopSinglesStruct)
	var upIds, dowmIds, upSingleIds, downSingleIds []int
	for _, v := range shopSinglesStruct {
		if v.ShopId != shopId {
			err = toolLib.CreateKcErr(_const.OPT_OTHER_BUS_ITEM)
			return
		}
		if v.Status == models.STATUS_OFF_SALE {
			dowmIds = append(dowmIds, v.SsId)
			downSingleIds = append(downSingleIds, v.SingleId)
		}
		if v.Status == models.STATUS_ON_SALE {
			upIds = append(upIds, v.SsId)
			upSingleIds = append(upSingleIds, v.SingleId)
		}
	}
	// 优化  分店设置单项目上下架、删除则只影响该分店该项目状态
	// mSingle := new(models.SingleModel).Init()
	// 下架操作
	if args.OptType == cards.STATUS_OFF_SALE && len(upIds) > 0 {
		mShopSingle.UpDateBySsids(upIds, map[string]interface{}{
			mShopSingle.Field.F_status:     cards.STATUS_OFF_SALE,
			mShopSingle.Field.F_under_time: time.Now().Unix(),
		})
		// 优化  分店设置单项目上下架、删除则只影响该分店该项目状态
		// for _, singleId := range upSingleIds {
		//	mSingle.DecrSaleshopnumBySingleid(singleId, 1)
		// }
		// 添加维护es的shop-item文档的任务
		setShopItem(ctx, upSingleIds, shopId, cards.ITEM_TYPE_single)
		// 同步下架门店卡项关联表数据
		shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
		if !shopItemRealtionModel.UpdateStatusByItemIds(upSingleIds, cards.ITEM_TYPE_single, cards.STATUS_OFF_SALE, shopId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}
	// 上架操作
	if args.OptType == cards.STATUS_ON_SALE && len(dowmIds) > 0 {
		mShopSingle.UpDateBySsids(dowmIds, map[string]interface{}{
			mShopSingle.Field.F_status:     cards.STATUS_ON_SALE,
			mShopSingle.Field.F_under_time: 0,
		})
		// 优化  分店设置单项目上下架、删除则只影响该分店该项目状态
		/*for _, singleId := range downSingleIds {
			mSingle.IncrSaleshopnumBySingleid(singleId, 1)
		}*/
		// 添加维护es的shop-item文档的任务
		setShopItem(ctx, downSingleIds, shopId, cards.ITEM_TYPE_single)
		// 同步上架门店卡项关联表数据
		shopItemRealtionModel := new(models.ShopItemRelationModel).Init()
		if !shopItemRealtionModel.UpdateStatusByItemIds(downSingleIds, cards.ITEM_TYPE_single, cards.STATUS_ON_SALE, shopId) {
			err = toolLib.CreateKcErr(_const.DB_ERR)
			return
		}
	}

	return
}

// 获取单项目的所有属性标签
func (s *SingleLogic) GetAttrs() cards.ReplyGetAttrs {
	data := cards.ReplyGetAttrs{
		SexBracket: []cards.Attr{},
		AgeBracket: []cards.Attr{},
		Tailor:     []cards.Attr{},
		TailorSub:  map[int][]cards.Attr{},
	}
	for k, v := range cards.SexBracket {
		if len(v) == 0 {
			continue
		}
		data.SexBracket = append(data.SexBracket, cards.Attr{
			Id:   k,
			Name: v,
		})
	}

	for k, v := range cards.AgeBracket {
		if len(v) == 0 {
			continue
		}
		data.AgeBracket = append(data.AgeBracket, cards.Attr{
			Id:   k,
			Name: v,
		})
	}

	for k, v := range cards.Tailor {
		if len(v) == 0 {
			continue
		}
		data.Tailor = append(data.Tailor, cards.Attr{
			Id:   k,
			Name: v,
		})
	}

	for pid, tailors := range cards.TailorSub {
		if len(tailors) == 0 {
			continue
		}
		for k, v := range tailors {
			if len(v) == 0 {
				continue
			}
			data.TailorSub[pid] = append(data.TailorSub[pid], cards.Attr{
				Id:   k,
				Name: v,
			})
		}
	}

	return data
}

// 检查单项目的数据信息
func (s *SingleLogic) checkSingleInfo(ctx context.Context, singleInfo *cards.SingleBase, busId int) (singImg singleImg, spPrice specPrice, err error) {
	// 检查价格
	if singleInfo.Price <= 0 || singleInfo.RealPrice <= 0 {
		err = toolLib.CreateKcErr(_const.PRICE_EMPTY)
		return
	}
	if singleInfo.Price < singleInfo.RealPrice {
		err = toolLib.CreateKcErr(_const.PRICE_LT_REALPRICE)
		return
	}
	spPrice.MaxPrice = 0
	spPrice.MinPrice = 0
	// 检查所传的行业id是否是商家包含的行业id
	rpcBus := new(bus.Bus).Init()
	defer rpcBus.Close()
	checkArgs := &bus2.ArgsCheckBindid{
		BusId:  busId,
		BindId: singleInfo.BindId,
	}
	var checkReply = false
	rpcBus.CheckBindid(ctx, checkArgs, &checkReply)
	if checkReply == false {
		err = toolLib.CreateKcErr(_const.BINDID_NOT_IN_BUS)
		return
	}
	// 检查适用年龄
	// if len(singleInfo.AgeBracket) == 0 {
	//	err = toolLib.CreateKcErr(_const.AGE_BRACKET_EMPTY)
	//	return
	// }
	// if s.CheckAgeBracket(singleInfo.AgeBracket) == false {
	//	err = toolLib.CreateKcErr(_const.AGE_BRACKET_ERR)
	//	return
	// }
	// 检查订制分类
	// if s.CheckTailor(singleInfo.TailorIndus) == false {
	//	err = toolLib.CreateKcErr(_const.TAILOR_ERRR)
	//	return
	// }
	// 检查订制分类子项
	// if len(singleInfo.TailorSubIndus) == 0 {
	//	err = toolLib.CreateKcErr(_const.TAILOR_SUB_EMPTY)
	//	return
	// }
	// if s.CheckTailorSub(singleInfo.TailorIndus, singleInfo.TailorSubIndus) == false {
	//	err = toolLib.CreateKcErr(_const.TAILOR_SUB_ERR)
	//	return
	// }
	// 检查标签是否属于商家
	if len(singleInfo.TagIds) > 0 {
		err = s.CheckTagids(busId, singleInfo.TagIds)
		if err != nil {
			return
		}
	}

	// 检查温馨信息  （一期优化，温馨提示只保留提示名称和提示内容）
	// rpcPublic := new(public.Reminder).Init()
	// defer rpcPublic.Close()
	// var bindId = singleInfo.BindId
	// var reminderInfos = []public2.ReminderInfo{}
	// err = rpcPublic.GetReminderInfo(ctx, &bindId, &reminderInfos)
	// if err != nil {
	//	return
	// }

	// reminder := map[string]interface{}{}
	// for _, value := range reminderInfos {
	//	if value.Require == public2.Require_true {
	//		if val, ok := singleInfo.Reminder[value.Key]; ok {
	//			if value.Type == public2.Type_input {
	//				if len(val.(string)) > 0 {
	//					reminder[value.Key] = val
	//				} else {
	//					err = toolLib.CreateKcErr(_const.REMINDER_EMPTY, lang.GetLang(_const.REMINDER_EMPTY)+value.Name)
	//					return
	//				}
	//			} else if value.Type == public2.Type_checkbox || value.Type == public2.Type_select {
	//				if len(val.([]interface{})) > 0 {
	//					reminder[value.Key] = val
	//				} else {
	//					err = toolLib.CreateKcErr(_const.REMINDER_EMPTY, lang.GetLang(_const.REMINDER_EMPTY)+value.Name)
	//					return
	//				}
	//			}
	//		} else {
	//			err = toolLib.CreateKcErr(_const.REMINDER_EMPTY, lang.GetLang(_const.REMINDER_EMPTY)+value.Name)
	//			return
	//		}
	//	} else {
	//		if val, ok := singleInfo.Reminder[value.Key]; ok {
	//			if value.Type == public2.Type_input {
	//				if len(val.(string)) > 0 {
	//					reminder[value.Key] = val
	//				}
	//			} else if value.Type == public2.Type_checkbox || value.Type == public2.Type_select {
	//				if len(val.([]interface{})) > 0 {
	//					reminder[value.Key] = val
	//				}
	//			}
	//		}
	//	}
	// }
	// singleInfo.Reminder = reminder

	var imgHashs []string

	// 检查封面图片信息
	if len(singleInfo.ImgHash) == 0 {
		err = toolLib.CreateKcErr(_const.INDEX_IMG_EMPTY)
		return
	}
	// 检查相册图片
	if len(singleInfo.Pictures) == 0 {
		err = toolLib.CreateKcErr(_const.PICTURES_EMPTY)
		return
	}
	singleInfo.Pictures = functions.ArrayUniqueString(singleInfo.Pictures)
	imgHashs = append(imgHashs, singleInfo.ImgHash)
	imgHashs = append(imgHashs, singleInfo.Pictures...)
	// 服务效果图片 (一期优化，删除）
	// if len(singleInfo.EffectImgs) > 3 || len(singleInfo.ToolsImgs) > 3 || len(singleInfo.Pictures) > 3 {
	//	err = toolLib.CreateKcErr(_const.PIC_MAX_ERR) //最多三张
	//	return
	// }
	// if len(singleInfo.EffectImgs) > 0 {
	//	singleInfo.EffectImgs = functions.ArrayUniqueString(singleInfo.EffectImgs)
	//	imgHashs = append(imgHashs, singleInfo.EffectImgs...)
	// }
	// //仪器设备图片
	// if len(singleInfo.ToolsImgs) > 0 {
	//	singleInfo.ToolsImgs = functions.ArrayUniqueString(singleInfo.ToolsImgs)
	//	imgHashs = append(imgHashs, singleInfo.ToolsImgs...)
	// }

	rpcImg := new(file.Upload).Init()
	defer rpcImg.Close()
	var replyImgs = map[string]file2.ReplyFileInfo{}
	err = rpcImg.GetImageByHashs(ctx, imgHashs, &replyImgs)
	if err != nil {
		err = toolLib.CreateKcErr(_const.PICTURES_EMPTY)
		return
	}
	if indexImg, ok := replyImgs[singleInfo.ImgHash]; ok {
		singImg.IndexImgId = indexImg.Id
	} else {
		err = toolLib.CreateKcErr(_const.INDEX_IMG_EMPTY)
		return
	}
	for _, v := range singleInfo.Pictures {
		if _, ok := replyImgs[v]; ok {
			singImg.SubImgIds = append(singImg.SubImgIds, replyImgs[v].Id)
		}
	}
	logs.Info("SubImgIds: ", singImg.SubImgIds)
	// 服务效果图
	// if len(singleInfo.EffectImgs) > 0 {
	//	for _, v := range singleInfo.EffectImgs {
	//		if _, ok := replyImgs[v]; ok {
	//			singImg.EffectImgIds = append(singImg.EffectImgIds, replyImgs[v].Id)
	//		}
	//	}
	// }
	// 仪器设备图
	// if len(singleInfo.ToolsImgs) > 0 {
	//	for _, v := range singleInfo.ToolsImgs {
	//		if _, ok := replyImgs[v]; ok {
	//			singImg.ToolImgIds = append(singImg.ToolImgIds, replyImgs[v].Id)
	//		}
	//	}
	// }

	// 检查规格
	if len(singleInfo.SpecIds) > 0 {
		var specIds []int
		for _, v := range singleInfo.SpecIds {
			if len(v.Sub) == 0 {
				err = toolLib.CreateKcErr(_const.SPEC_SUB_EMPTY)
				return
			}
			specIds = append(specIds, v.SpecId)
		}
		mBusSpec := new(models.SingleBusSpecModel).Init()
		specs := mBusSpec.GetBySpecids(specIds)
		if len(specs) == 0 {
			err = toolLib.CreateKcErr(_const.SPEC_NOT_IN_BUS)
			return
		}
		var specStruct = []struct {
			SpecId int
			BusId  int
		}{}
		mapstructure.WeakDecode(specs, &specStruct)
		for _, sp := range specStruct {
			if sp.BusId != busId {
				err = toolLib.CreateKcErr(_const.SPEC_NOT_IN_BUS)
				return
			}
		}
		// 检查子规格组合数据
		if len(singleInfo.SpecPrices) == 0 {
			err = toolLib.CreateKcErr(_const.SPEC_PRICE_EMPTY)
			return
		}

		var sonSpecIds []int
		for _, sp := range singleInfo.SpecPrices {
			if len(sp.SpecIds) == 0 {
				err = toolLib.CreateKcErr(_const.SPEC_SUB_EMPTY)
				return
			}
			if sp.Price <= 0 {
				err = toolLib.CreateKcErr(_const.SPEC_PRICE_EMPTY)
				return
			}
			if sp.Price > spPrice.MaxPrice {
				spPrice.MaxPrice = sp.Price
			}
			if spPrice.MinPrice == 0 || sp.Price < spPrice.MinPrice {
				spPrice.MinPrice = sp.Price
			}
			sonSpecIds = append(sonSpecIds, sp.SpecIds...)
		}

		if len(sonSpecIds) > 0 {
			sonSpecs := mBusSpec.GetBySpecids(sonSpecIds)
			pspecIds := functions.ArrayUniqueInt(functions.ArrayValue2Array(mBusSpec.Field.F_p_spec_id, sonSpecs))
			if len(pspecIds) != len(functions.ArrayUniqueInt(specIds)) {
				err = toolLib.CreateKcErr(_const.SPEC_PRICE_NOT_MATCH)
				return
			}
		}
	}
	return
}

// 检查标签id是否属于商家
func (s *SingleLogic) CheckTagids(busId int, tagIds []int) error {
	mTags := new(models.TagModel).Init()
	r := mTags.GetByTagids(tagIds)
	var tags = []struct {
		BusId int
		TagId int
	}{}
	mapstructure.WeakDecode(r, &tags)
	for _, tag := range tags {
		if tag.BusId != busId {
			return toolLib.CreateKcErr(_const.TAG_NOT_IN_BUS)
		}
	}

	return nil
}

// 检查适用年龄
func (s *SingleLogic) CheckAgeBracket(ageBracket []int) bool {
	allowAge := []int{}
	for k, v := range cards.AgeBracket {
		if len(v) == 0 {
			continue
		}
		allowAge = append(allowAge, k)
	}
	for _, v := range ageBracket {
		if functions.InArray(v, allowAge) == false {
			return false
		}
	}
	return true
}

// 检查订制分类
func (s *SingleLogic) CheckTailor(tailorId int) bool {
	allowTailorIds := []int{}
	for k, v := range cards.Tailor {
		if len(v) == 0 {
			continue
		}
		allowTailorIds = append(allowTailorIds, k)
	}
	return functions.InArray(tailorId, allowTailorIds)
}

// 检查订制子分类
func (s *SingleLogic) CheckTailorSub(tailorId int, subIds []int) bool {
	allowIds := []int{}
	tsub := cards.TailorSub[tailorId]
	for k, v := range tsub {
		if len(v) == 0 {
			continue
		}

		allowIds = append(allowIds, k)
	}

	for _, v := range subIds {
		if functions.InArray(v, allowIds) == false {
			return false
		}
	}
	return true
}

// 处理单项目的图片
// @param singleImg singImg 单项目的相册图片，效果图片和仪器设备图片
// @param int singleId 单项目id
// @param uint8 setType 操作类型 1=添加 2=修改
func (s *SingleLogic) setSinglePic(singImg singleImg, singleId int, setType uint8) {
	mSI := new(models.SingleImgModel).Init()
	var imgs []map[string]interface{}
	// 相册图片
	if len(singImg.SubImgIds) > 0 {
		for _, imgId := range singImg.SubImgIds {
			imgs = append(imgs, map[string]interface{}{
				mSI.Field.F_single_id: singleId,
				mSI.Field.F_img_id:    imgId,
				mSI.Field.F_type:      cards.IMG_TYPE_album,
			})
		}
	}
	// 一期优化去除
	// if len(singImg.EffectImgIds) > 0 {
	//	for _, imgId := range singImg.EffectImgIds {
	//		imgs = append(imgs, map[string]interface{}{
	//			mSI.Field.F_single_id: singleId,
	//			mSI.Field.F_img_id:    imgId,
	//			mSI.Field.F_type:      cards.IMG_TYPE_effect,
	//		})
	//	}
	// }
	// if len(singImg.ToolImgIds) > 0 {
	//	for _, imgId := range singImg.ToolImgIds {
	//		imgs = append(imgs, map[string]interface{}{
	//			mSI.Field.F_single_id: singleId,
	//			mSI.Field.F_img_id:    imgId,
	//			mSI.Field.F_type:      cards.IMG_TYPE_tool,
	//		})
	//	}
	// }
	// 添加
	if len(imgs) > 0 && setType == SINGLE_SET_TYPE_ADD {
		mSI.InsertAll(imgs)
		return
	}
	// 修改
	if len(imgs) > 0 && setType == SINGLE_SET_TYPE_EDIT {
		imgsStruct := []struct {
			ImgId int
			Type  int
		}{}
		mapstructure.WeakDecode(imgs, &imgsStruct)
		// 获取数据库当前前的图片id
		dbImgs := mSI.GetBySingleId(singleId)
		dbImgsStruct := []struct {
			Id    int
			ImgId int
			Type  int
		}{}
		mapstructure.WeakDecode(dbImgs, &dbImgsStruct)

		// 统计出需要新增和删除的图片数据
		var addImgs []map[string]interface{}
		var delImgIds []int
		for _, v := range imgsStruct {
			hasd := 0
			for _, v2 := range dbImgsStruct {
				if v.ImgId == v2.ImgId && v.Type == v2.Type {
					hasd = 1
					break
				}
			}
			// 新添加的图片不在数据库里面，需要新增
			if hasd == 0 {
				addImgs = append(addImgs, map[string]interface{}{
					mSI.Field.F_single_id: singleId,
					mSI.Field.F_img_id:    v.ImgId,
					mSI.Field.F_type:      v.Type,
				})
			}
		}

		for _, v := range dbImgsStruct {
			hasd := 0
			for _, v2 := range imgsStruct {
				if v.ImgId == v2.ImgId && v.Type == v2.Type {
					hasd = 1
					break
				}
			}
			if hasd == 0 {
				delImgIds = append(delImgIds, v.Id)
			}
		}
		if len(addImgs) > 0 {
			mSI.InsertAll(addImgs)
		}
		if len(delImgIds) > 0 {
			mSI.DelByIds(delImgIds)
		}
	}
}

// 设置规格
func (s *SingleLogic) setSingleSpec(specIds []cards.SingleSpecIds, specPrices []cards.SingleSpecPrice, singleId int, setType uint8) {
	mSS := new(models.SingleSpecModel).Init()
	mSSP := new(models.SingleSpecPriceModel).Init()
	if len(specIds) == 0 {
		if setType == SINGLE_SET_TYPE_ADD {
			return
		}
		if setType == SINGLE_SET_TYPE_EDIT {
			// 修改 规格都去除了，删除规格和规格价格
			mSS.DeleteBySingleid(singleId)
			mSSP.UpdateBySingleid(singleId, map[string]interface{}{
				mSSP.Field.F_is_del: cards.SPEC_PRICE_IS_DEL_yes,
			})
		}
	}
	specInfo, _ := json.Marshal(specIds)
	ssInfo := mSS.GetBySingleid(singleId)
	if len(ssInfo) > 0 {
		ssid, _ := strconv.Atoi(ssInfo[mSS.Field.F_id].(string))
		mSS.UpdateById(ssid, map[string]interface{}{
			mSS.Field.F_spec_info: string(specInfo),
		})
	} else {
		mSS.Insert(map[string]interface{}{
			mSS.Field.F_single_id: singleId,
			mSS.Field.F_spec_info: string(specInfo),
		})
	}

	var specPricesMap []map[string]interface{}
	var hashs = map[int]string{}
	var specIdsMap = map[int]string{}
	for k, sp := range specPrices {
		sort.Ints(sp.SpecIds)
		specIds := functions.Implode(",", sp.SpecIds)
		hashStr := functions.HashMd5(fmt.Sprintf("%s|%d", specIds, singleId))
		hashs[k] = hashStr
		specIdsMap[k] = specIds
		specPricesMap = append(specPricesMap, map[string]interface{}{
			mSSP.Field.F_single_id: singleId,
			mSSP.Field.F_price:     sp.Price,
			mSSP.Field.F_spec_ids:  specIds,
			mSSP.Field.F_hash:      hashStr,
		})
	}
	// 添加单项目
	if setType == SINGLE_SET_TYPE_ADD {
		// 多规格价格处理
		mSSP.InsertAll(specPricesMap)
		return
	}

	// 修改单项目
	if setType == SINGLE_SET_TYPE_EDIT {
		prices := mSSP.GetBySingleid(singleId)
		pricesStruct := []struct {
			SspId int
			Hash  string
			Price float64
		}{}
		mapstructure.WeakDecode(prices, &pricesStruct)

		type updateDataStruct struct {
			SspId int
			Price float64
		}
		// 需要同步到子店的新增规格的组合id极价格 sspid=>price
		var addSspids = map[int]float64{}
		var updateData []updateDataStruct
		for k, sp := range specPrices {
			hasd := 0
			for _, v := range pricesStruct {
				if hashs[k] == v.Hash {
					hasd = 1
					if v.Price != sp.Price {
						updateData = append(updateData, updateDataStruct{
							SspId: v.SspId,
							Price: sp.Price,
						})
					}
				}
			}
			if hasd == 0 {
				sspid := mSSP.Insert(map[string]interface{}{
					mSSP.Field.F_single_id: singleId,
					mSSP.Field.F_price:     sp.Price,
					mSSP.Field.F_spec_ids:  specIdsMap[k],
					mSSP.Field.F_hash:      hashs[k],
				})
				addSspids[sspid] = sp.Price
			}
		}

		// 需要设置为删除的sspid
		var delSspids []int
		for _, v := range pricesStruct {
			hasd := 0
			for k, _ := range specPrices {
				if v.Hash == hashs[k] {
					hasd = 1
					break
				}
			}
			if hasd == 0 {
				delSspids = append(delSspids, v.SspId)
			}
		}

		if len(addSspids) > 0 {
			// 同步规格价格到子店
			// 先获取已添加此单项目的子店的ssid
			mShopSingle := new(models.ShopSingleModel).Init()
			shopSingles := mShopSingle.GetBySingleid(singleId)

			if len(shopSingles) > 0 {
				mSSSP := new(models.ShopSingleSpecPriceModel).Init()
				var addShopSpecPriceData = []map[string]interface{}{}
				for _, shopSingle := range shopSingles {
					for k, addprice := range addSspids {
						addShopSpecPriceData = append(addShopSpecPriceData, map[string]interface{}{
							mSSSP.Field.F_ss_id:   shopSingle[mShopSingle.Field.F_ss_id].(string),
							mSSSP.Field.F_ssp_id:  k,
							mSSSP.Field.F_price:   addprice,
							mSSSP.Field.F_shop_id: shopSingle[mShopSingle.Field.F_shop_id].(string),
						})
					}
				}
				if len(addShopSpecPriceData) > 0 {
					mSSSP.InsertAll(addShopSpecPriceData)
				}
			}
		}

		if len(delSspids) > 0 {
			mSSP.UpdateBySspids(delSspids, map[string]interface{}{
				mSSP.Field.F_is_del: cards.SPEC_PRICE_IS_DEL_yes,
			})
		}
		if len(updateData) > 0 {
			for _, v := range updateData {
				mSSP.UpdateBySspid(v.SspId, map[string]interface{}{
					mSSP.Field.F_price: v.Price,
				})
			}
		}
	}
}

// 根据单项目id批量获取基础价格信息
func (s *SingleLogic) GetSinglePriceListsBySingleIds(singleIds []int) (result map[int]cards.SinglePriceInfo) {
	singleModel := new(models.SingleModel).Init()
	lists := singleModel.GetBySingleids(singleIds, []string{
		singleModel.Field.F_single_id,
		singleModel.Field.F_name,
		singleModel.Field.F_img_id,
		singleModel.Field.F_price,
		singleModel.Field.F_real_price,
		singleModel.Field.F_min_price,
		singleModel.Field.F_max_price,
	})
	if err := mapstructure.WeakDecode(functions.ArrayRebuild(singleModel.Field.F_single_id, lists), &result); err != nil {
		return result
	}
	return result
}

// 根据手艺人ID获取关联的单项目
func (s *SingleLogic) GetSignlesByStaffId(ctx context.Context, args *cards.ArgsGetSignlesByStaffID, reply *cards.ReplyGetSignlesByStaffID) (err error) {
	shopId, staffId := args.ShopId, args.StaffId
	if shopId == 0 || staffId == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	reply.StaffId = staffId
	// 员工服务,是否是全部服务/指定服务
	rpcStaff := new(staff.StaffService).Init()
	var staffReply staff2.ReplyGetServiceByStaffId
	if err = rpcStaff.GetServiceByStaffId(ctx, &staff2.ArgsGetServiceByStaffId{StaffId: staffId}, &staffReply); err != nil {
		return
	}
	if len(staffReply.SingleIds) == 0 {
		args.First = true
	}
	args.GetAll = staffReply.GetAll
	var shopSspIds []int // 已选中服务的规格
	if !args.GetAll && !args.First {
		args.SingleIds = staffReply.SingleIds
		args.SspIds = staffReply.SspIds
		if len(args.SingleIds) == 0 || len(args.SspIds) == 0 {
			err = toolLib.CreateKcErr(_const.PARAM_ERR)
			return
		}
		shopSspIds = append(shopSspIds, args.SspIds...)
	}
	// 根据单项目ID查找门店单项目
	start, limit := args.GetStart(), args.GetPageSize()
	shopSingleModel := new(models.ShopSingleModel).Init()

	// ------
	// 所有子店的单项目(只选择上架的单项目)
	allShopSinglesMap := shopSingleModel.GetSingles(map[string]interface{}{
		shopSingleModel.Field.F_shop_id: shopId,
		shopSingleModel.Field.F_status:  models.STATUS_ON_SALE,
	}, start, limit)
	var list []cards.SignleStaBase
	var shopSingleStruct = []struct {
		SsId             int
		SingleId         int
		ChangedRealPrice float64
		ChangedMinPrice  float64
		ChangedMaxPrice  float64
		Status           int
		Sales            int
		Name             string
	}{}
	_ = mapstructure.WeakDecode(allShopSinglesMap, &shopSingleStruct)
	allShopSinIds := functions.ArrayValue2Array(shopSingleModel.Field.F_single_id, allShopSinglesMap)

	// 获取单项目数据
	mSingle := new(models.SingleModel).Init()
	where := map[string]interface{}{
		// mSingle.Field.F_bus_id:busId,
		mSingle.Field.F_single_id: []interface{}{"IN", allShopSinIds},
	}
	if args.BindId > 0 || len(args.Name) > 0 {
		if args.BindId > 0 {
			where[mSingle.Field.F_bind_id] = args.BindId
		} else {
			where[mSingle.Field.F_name] = []interface{}{"like", "%" + args.Name + "%"}
		}
	}
	singles := mSingle.GetList(where)
	singlesStruct := []cards.SignleStaBase{}
	mapstructure.WeakDecode(singles, &singlesStruct)
	singlesStructMap := map[int]cards.SignleStaBase{} // 单项目数据

	for _, v := range singlesStruct {
		singlesStructMap[v.SingleId] = v
	}
	var ssids []int // 子店单项目ID
	for index, v := range shopSingleStruct {
		ssids = append(ssids, v.SsId)
		list = append(list, singlesStructMap[v.SingleId])
		list[index].MaxPrice = v.ChangedMaxPrice
		list[index].MinPrice = v.ChangedMinPrice
		list[index].RealPrice = v.ChangedRealPrice
		if len(v.Name) > 0 {
			list[index].Name = v.Name
		}
		list[index].Status = v.Status
		list[index].CtimeStr = time.Unix(list[index].Ctime, 0).Format("2006-01-02 15:04:05")
	}

	// 根据子店单项目IDs获取所有子店规格数据(包含删除的数据)
	var specDetailList []cards.SpecPricesBase
	var shopSingleSpecPriceStruct = []struct {
		Id    int
		SsId  int
		SspId int
		Price float64
	}{}
	shopSingleSpecPriceModel := new(models.ShopSingleSpecPriceModel).Init()
	ssspMap := shopSingleSpecPriceModel.GetBySsids(ssids)
	// 单项目不同规格的价格表主键,对应kc_single_spec_price表中的id
	sspIds := functions.ArrayValue2Array(shopSingleSpecPriceModel.Field.F_ssp_id, ssspMap)
	_ = mapstructure.WeakDecode(ssspMap, &shopSingleSpecPriceStruct)

	// 根据sspIds获取所有单项目规格的价格数据(价格等)
	singleSpecModel := new(models.SingleSpecPriceModel).Init()
	singleSpecMap := singleSpecModel.GetBySspids(sspIds) // 未删除的规格数据
	noDelSSpIds := functions.ArrayValue2Array(singleSpecModel.Field.F_ssp_id, singleSpecMap)
	// 未删除的单项目子规格数据
	var noDelShopSingleSpecPriceStruct = []struct {
		Id    int
		SsId  int
		SspId int
		Price float64
	}{}
	for _, v := range shopSingleSpecPriceStruct {
		if functions.InArray(v.SspId, noDelSSpIds) {
			noDelShopSingleSpecPriceStruct = append(noDelShopSingleSpecPriceStruct, v)
		}
	}
	var singleSpecStruct []cards.SpecPricesBase
	_ = mapstructure.WeakDecode(singleSpecMap, &singleSpecStruct)
	singleSpecData := map[int]cards.SpecPricesBase{}
	for _, v := range singleSpecStruct {
		singleSpecData[v.SspId] = v
	}

	//
	for index, v := range noDelShopSingleSpecPriceStruct {
		if singleSpecData[v.SspId].SspId > 0 {
			specDetailList = append(specDetailList, singleSpecData[v.SspId])
			specDetailList[index].Price = v.Price
		}
	}
	// 规格是否被选中
	for index, v := range specDetailList {
		if args.GetAll {
			specDetailList[index].Selected = true
		} else {
			if functions.InArray(v.SspId, shopSspIds) {
				specDetailList[index].Selected = true
			}
		}
	}
	// fmt.Printf("list: %#v\n", list)
	// fmt.Printf("specDetailList: %#v\n", specDetailList)
	// 获取单项目选中的各种规格
	var specIds, specIdd []cards.SingleSpecIds
	singleSpecName := map[int]cards.SingleSpec{}

	// if baseInfo.HasSpec == cards.SINGLE_HAS_SPEC_yes {
	mSS := new(models.SingleSpecModel).Init()
	specs := mSS.GetBySingleids(allShopSinIds)
	for _, spec := range specs {
		if err = json.Unmarshal([]byte(spec[mSS.Field.F_spec_info].(string)), &specIdd); err != nil {
			continue
		}
		specIds = append(specIds, specIdd...)
	}
	var allSpecId []int
	if len(specIds) > 0 {
		for _, v := range specIds {
			allSpecId = append(allSpecId, v.SpecId)
			allSpecId = append(allSpecId, v.Sub...)
		}
	}
	if len(allSpecId) > 0 {
		mSBS := new(models.SingleBusSpecModel).Init()
		allSpec := mSBS.GetBySpecids(allSpecId)
		var allSpecStruct []cards.SingleSpec
		mapstructure.WeakDecode(allSpec, &allSpecStruct)
		for _, v := range allSpecStruct {
			singleSpecName[v.SpecId] = v
		}
	}
	reply.Specs = singleSpecName
	reply.SpecIds = specIds
	// fmt.Printf("Specs:%#v\n", singleSpecName)
	// fmt.Printf("SpecIds:%#v\n", specIds)

	// 单项目封面图片数据
	singleImageModel := new(models.SingleImgModel).Init()
	singleWhere := map[string]interface{}{
		singleImageModel.Field.F_single_id: []interface{}{"IN", allShopSinIds},
		singleImageModel.Field.F_type:      1,
	}
	singleImageMap := singleImageModel.GetBySingleIds(singleWhere)
	singleImgIds := functions.ArrayValue2Array(singleImageModel.Field.F_img_id, singleImageMap)
	// rpc File
	rpcFile := new(file.Upload).Init()
	defer rpcFile.Close()
	var respFile = map[int]file2.ReplyFileInfo{}
	if err = rpcFile.GetImageByIds(ctx, singleImgIds, &respFile); err != nil {
		return
	}
	for index, v := range list {
		list[index].ImgUrl = respFile[v.ImgId].Path
	}

	reply.List = list
	reply.SpecPrices = specDetailList
	reply.TotalNum = shopSingleModel.GetNumByShopid(shopId, "2", "0", []int{})
	reply.AssignService = args.GetAll

	return
}

// 获取门店的单项目-rpc内部调用
func (s *SingleLogic) GetShopSingleBySingleIdsRpc(ctx context.Context, args *cards.ArgsGetShopSingleBySingleIdsRpc, reply *cards.ReplyGetShopSingleBySingleIdsRpc) (err error) {
	shopSingleModel := new(models.ShopSingleModel).Init()
	shopSingleMap := shopSingleModel.GetByShopidAndSingleids(args.ShopId, args.SingleIds, []string{shopSingleModel.Field.F_single_id,
		shopSingleModel.Field.F_name})
	singleIds := functions.ArrayValue2Array(shopSingleModel.Field.F_single_id, shopSingleMap)
	singleModel := new(models.SingleModel).Init()
	singleMap := singleModel.GetList(map[string]interface{}{singleModel.Field.F_single_id: []interface{}{"IN", singleIds}}, []string{singleModel.Field.F_single_id, singleModel.Field.F_name})
	for _, ssm := range shopSingleMap {
		if len(ssm[shopSingleModel.Field.F_name].(string)) > 0 {
			for i, m := range singleMap {
				if m[singleModel.Field.F_single_id].(string) == ssm[shopSingleModel.Field.F_single_id].(string) {
					singleMap[i][singleModel.Field.F_name] = ssm[shopSingleModel.Field.F_name]
				}
				break
			}
		}
	}
	_ = mapstructure.WeakDecode(singleMap, &reply.List)
	return
}

// 根据门店id查询单项目服务
func (s *SingleLogic) GetSingleByShopIdAndTagId(ctx context.Context, args *cards.ArgsShopSingleByPage, reply *cards.ReplyShopSingle) error {
	if args.ShopId <= 0 {
		return toolLib.CreateKcErr(_const.SHOPID_NTL)
	}
	model := new(models.ShopSingleModel).Init()
	shopInfoMaps := model.GetByShopid(args.ShopId, args.GetStart(), args.GetPageSize(), args.Status, args.IsDel, args.SingleIds)
	if len(shopInfoMaps) == 0 {
		return nil
	}
	singleIds := functions.ArrayValue2Array(model.Field.F_single_id, shopInfoMaps)
	singleModel := new(models.SingleModel).Init()
	singleMaps := singleModel.GetBySingleids(singleIds, []string{singleModel.Field.F_real_price, singleModel.Field.F_price,
		singleModel.Field.F_service_time, singleModel.Field.F_single_id, singleModel.Field.F_name, singleModel.Field.F_img_id})
	var imgIds []int
	for _, singleMap := range singleMaps {
		for _, infoMap := range shopInfoMaps {
			if singleMap[singleModel.Field.F_single_id].(string) == infoMap[model.Field.F_single_id].(string) {
				if infoMap[model.Field.F_changed_real_price].(string) != "0.00" {
					singleMap[singleModel.Field.F_real_price] = infoMap[model.Field.F_changed_real_price]
				}
				if len(infoMap[model.Field.F_name].(string)) > 0 {
					singleMap[singleModel.Field.F_name] = infoMap[model.Field.F_name]
				}
				singleMap["SsId"] = infoMap[model.Field.F_ss_id]
				break
			}
		}
		imgId, _ := strconv.Atoi(singleMap[singleModel.Field.F_img_id].(string))
		imgIds = append(imgIds, imgId)
	}
	imgIds = functions.ArrayUniqueInt(imgIds)
	var list []cards.SingleInfo
	if err := mapstructure.WeakDecode(singleMaps, &list); err != nil {
		return err
	}
	totalNum := model.GetNumByShopid(args.ShopId, args.Status, args.IsDel, []int{})
	reply.Lists = &list
	reply.TotalNum = totalNum
	rpcFile := new(file.Upload).Init()
	defer rpcFile.Close()
	var replyFile map[int]file2.ReplyFileInfo
	if err := rpcFile.GetImageByIds(ctx, imgIds, &replyFile); err != nil {
		return err
	}
	// 默认图片
	replyFile[0] = file2.ReplyFileInfo{
		Id:   0,
		Hash: "",
		Path: constkey.CardsSmallDefaultPics[cards.ITEM_TYPE_single],
	}
	reply.SingleImg = replyFile
	return nil
}

// 根据单项目id获取子规格服务
func (s *SingleLogic) GetSubServerBySingleId(args *cards.ArgsSubServer, reply *cards.ReplySubServer) error {
	reply.SubServer = []cards.SubServer{}
	reply.Specs = []cards.Specs{}
	if args.SingleId <= 0 {
		return toolLib.CreateKcErr(_const.SINGLEID_NTL)
	}

	// 查询单项目的规格信息
	specModel := new(models.SingleSpecModel).Init()
	specMap := specModel.GetBySingleid(args.SingleId, specModel.Field.F_spec_info)
	// 如果单项目没有规格
	if len(specMap) == 0 {
		return nil
	}
	var specInfos []cards.SingleSpecIds
	err := json.Unmarshal([]byte(specMap[specModel.Field.F_spec_info].(string)), &specInfos)
	if err != nil {
		return err
	}
	// 获取规格id数组
	var specIds []int
	for _, info := range specInfos {
		specIds = append(specIds, info.SpecId)
		for _, i := range info.Sub {
			specIds = append(specIds, i)
		}
	}
	// 根据 多个规格id 查询规格名称
	sbsModel := new(models.SingleBusSpecModel).Init()
	sbsMaps := sbsModel.GetBySpecids(specIds)
	// 存入规格map
	var sMap = make(map[int]string)
	var id int
	for _, sbsMap := range sbsMaps {
		id, _ = strconv.Atoi(sbsMap[sbsModel.Field.F_spec_id].(string))
		sMap[id] = sbsMap[sbsModel.Field.F_name].(string)
	}

	// 封装 规格信息
	for _, info := range specInfos {
		var spec cards.Specs
		spec.Id = info.SpecId
		spec.Name = sMap[info.SpecId]
		for _, sub := range info.Sub {
			spec.Sub = append(spec.Sub, cards.SubSpecs{
				Id: sub, Name: sMap[sub],
			})
		}
		reply.Specs = append(reply.Specs, spec)
	}

	// 根据 sspIds 查询规格信息
	sspModel := new(models.SingleSpecPriceModel).Init()
	sspMaps := sspModel.GetBySingleid(args.SingleId)
	if err = mapstructure.WeakDecode(sspMaps, &reply.SubServer); err != nil {
		return err
	}
	if args.ShopId > 0 {
		// 根据singleId  shopId 查询单项目在门店的id
		ssModel := new(models.ShopSingleModel).Init()
		ssMap := ssModel.GetByShopidAndSingleid(args.ShopId, args.SingleId, []string{ssModel.Field.F_ss_id})
		if len(ssMap) == 0 {
			return nil
		}
		// 根据 ssid 查询 sspId对应的价格  数组
		ssId, _ := strconv.Atoi(ssMap[ssModel.Field.F_ss_id].(string))
		model := new(models.ShopSingleSpecPriceModel).Init()
		maps := model.GetBySsid(ssId, model.Field.F_ssp_id, model.Field.F_price)
		if len(maps) == 0 {
			return nil
		}

		rebuildMaps := map[string]string{}
		for _, v := range maps {
			rebuildMaps[v[model.Field.F_ssp_id].(string)] = v[model.Field.F_price].(string)
		}

		for k, m := range reply.SubServer {
			m.Price, _ = strconv.ParseFloat(rebuildMaps[strconv.Itoa(m.SspId)], 64)
			reply.SubServer[k] = m
		}
	}
	return nil
}

// 根据门店id和单项目Id获取单项目数据
func (s *SingleLogic) GetSingleByShopIdAndSingleIds(ctx context.Context, args *cards.ArgsGetSingleByShopIdAndSingleIds, reply *cards.ReplyShopSingle) (err error) {
	if args.ShopId <= 0 {
		return toolLib.CreateKcErr(_const.SHOPID_NTL)
	}
	model := new(models.ShopSingleModel).Init()
	shopInfoMaps := model.GetByShopIdAndSingleIds(args.ShopId, args.SingleIds)
	if len(shopInfoMaps) == 0 {
		return nil
	}
	singleIds := functions.ArrayValue2Array(model.Field.F_single_id, shopInfoMaps)
	singleModel := new(models.SingleModel).Init()
	singleMaps := singleModel.GetBySingleids(singleIds)
	var imgIds []int
	for index, singleMap := range singleMaps {
		for _, infoMap := range shopInfoMaps {
			if singleMap[singleModel.Field.F_single_id].(string) == infoMap[model.Field.F_single_id].(string) {
				if infoMap[model.Field.F_changed_real_price].(string) != "0.00" {
					singleMaps[index][singleModel.Field.F_real_price] = infoMap[model.Field.F_changed_real_price]
				}
				singleMaps[index]["SsId"] = infoMap[model.Field.F_ss_id]
				status, _ := strconv.Atoi(infoMap[model.Field.F_status].(string))
				if status == 1 || status == 2 {
					singleMaps[index]["ShopStatus"] = status
				}
				if status == 3 {
					singleMaps[index]["ShopDelStatus"] = 1
				} else {
					singleMaps[index]["ShopDelStatus"] = infoMap[model.Field.F_is_del]
				}
				break
			}
		}
		imgId, _ := strconv.Atoi(singleMap[singleModel.Field.F_img_id].(string))
		imgIds = append(imgIds, imgId)
	}
	imgIds = functions.ArrayUniqueInt(imgIds)
	var list []cards.SingleInfo
	if err := mapstructure.WeakDecode(singleMaps, &list); err != nil {
		return err
	}
	totalNum := model.GetNumByShopid(args.ShopId, "", "0", args.SingleIds)
	reply.Lists = &list
	reply.TotalNum = totalNum
	rpcFile := new(file.Upload).Init()
	defer rpcFile.Close()
	var replyFile map[int]file2.ReplyFileInfo
	if err := rpcFile.GetImageByIds(ctx, imgIds, &replyFile); err != nil {
		return err
	}
	// 默认图片
	replyFile[0] = file2.ReplyFileInfo{
		Id:   0,
		Hash: "",
		Path: constkey.CardsSmallDefaultPics[cards.ITEM_TYPE_single],
	}
	reply.SingleImg = replyFile
	return
}

// 根据门店ids ，获取数据
func (s *SingleLogic) GetBySsidsRpc(ssIds *[]int, reply *[]cards.ReplyGetBySsidsRpc) (err error) {
	if len(*ssIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	shopSingleModel := new(models.ShopSingleModel).Init()
	r := shopSingleModel.GetBySsids(*ssIds)
	_ = mapstructure.WeakDecode(r, reply)
	return
}

// 获取门店的指定规格数据
func (s *SingleLogic) GetShopSpecs(args *cards.ArgsGetShopSpecs, reply *[]cards.ReplyGetShopSpecs) (err error) {
	if args.ShopId <= 0 || len(args.SspIds) == 0 {
		err = toolLib.CreateKcErr(_const.PARAM_ERR)
		return
	}
	mSSSP := new(models.ShopSingleSpecPriceModel).Init()
	r := mSSSP.GetByShopidAndSspids(args.ShopId, args.SspIds)
	ssids := functions.ArrayValue2Array(mSSSP.Field.F_ss_id, r)
	// 获取规格的最高价和最低价
	mSS := new(models.ShopSingleModel).Init()
	mSSMaps := mSS.GetBySsids(ssids)
	mSSRebuild := functions.ArrayRebuild(mSS.Field.F_ss_id, mSSMaps)
	// 获取规格名称
	mSSP := new(models.SingleSpecPriceModel).Init()
	sspr := mSSP.GetBySspids(args.SspIds)
	singleIds := functions.ArrayValue2Array(mSSP.Field.F_single_id, sspr)
	sspDatas := map[int][]int{}
	var specIds []int
	// ---获取规格id
	for _, v := range sspr {
		sspid, _ := strconv.Atoi(v[mSSP.Field.F_ssp_id].(string))
		sspDatas[sspid] = functions.StrExplode2IntArr(v[mSSP.Field.F_spec_ids].(string), ",")
		specIds = append(specIds, sspDatas[sspid]...)
		for index, v2 := range r {
			if v[mSSP.Field.F_ssp_id].(string) == v2[mSSSP.Field.F_ssp_id].(string) {
				singleId, _ := strconv.Atoi(v[mSSP.Field.F_single_id].(string))
				r[index]["SingleId"] = singleId
			}
		}
	}
	// ---获取规格名称
	mSBS := new(models.SingleBusSpecModel).Init()
	sbsr := mSBS.GetBySpecids(specIds)
	sbsDatas := map[int]string{}
	for _, v := range sbsr {
		specId, _ := strconv.Atoi(v[mSBS.Field.F_spec_id].(string))
		sbsDatas[specId] = v[mSBS.Field.F_name].(string)
	}
	mapstructure.WeakDecode(r, reply)
	// 单项目原价
	mSM := new(models.SingleModel).Init()
	mSMMaps := mSM.GetBySingleids(singleIds, []string{mSM.Field.F_single_id, mSM.Field.F_price})
	mSMRebuild := functions.ArrayRebuild(mSM.Field.F_single_id, mSMMaps)

	for k, v := range *reply {
		for _, tspecId := range sspDatas[v.SspId] {
			v.SspName = fmt.Sprintf("%s,%s", v.SspName, sbsDatas[tspecId])
		}
		(*reply)[k].SspName = strings.TrimLeft(v.SspName, ",")
		minPrice := mSSRebuild[strconv.Itoa(v.SsId)].(map[string]interface{})[mSS.Field.F_changed_min_price].(string)
		(*reply)[k].MinPrice, _ = strconv.ParseFloat(minPrice, 64)
		maxPrice := mSSRebuild[strconv.Itoa(v.SsId)].(map[string]interface{})[mSS.Field.F_changed_max_price].(string)
		(*reply)[k].MaxPrice, _ = strconv.ParseFloat(maxPrice, 64)
		price := mSMRebuild[strconv.Itoa(v.SingleId)].(map[string]interface{})[mSM.Field.F_price].(string) // 单项目原价
		(*reply)[k].OriginalPrice, _ = strconv.ParseFloat(price, 64)
	}
	return
}

// 根据规格ID获取规格数据
func (s *SingleLogic) GetSingleSpecBySspId(args *cards.ArgsSubSpecID, reply *cards.ReplySubServer2) (err error) {
	if len(args.SspIds) == 0 {
		return
	}
	model := new(models.ShopSingleSpecPriceModel).Init()
	maps := model.GetBySspids(args.SspIds)
	sspIds := functions.ArrayValue2Array(model.Field.F_ssp_id, maps)
	sspModel := new(models.SingleSpecPriceModel).Init()
	sspMaps := sspModel.GetBySspids(sspIds)
	var specIds []int
	for _, sspMap := range sspMaps {
		specIds = append(specIds, functions.StrExplode2IntArr(sspMap[sspModel.Field.F_spec_ids].(string), ",")...)
	}
	specIds = functions.ArrayUniqueInt(specIds)
	sbsModel := new(models.SingleBusSpecModel).Init()
	sbsMaps := sbsModel.GetBySpecids(specIds)
	for _, sbsMap := range sbsMaps {
		parentId, _ := strconv.Atoi(sbsMap[sbsModel.Field.F_p_spec_id].(string))
		specIds = append(specIds, parentId)
	}
	specIds = functions.ArrayUniqueInt(specIds)
	sbsMaps = sbsModel.GetBySpecids(specIds)
	var specs []cards.Specs
	for _, m := range sbsMaps {
		if m[sbsModel.Field.F_p_spec_id].(string) == "0" {
			var spec cards.Specs
			id, _ := strconv.Atoi(m[sbsModel.Field.F_spec_id].(string))
			spec.Id = id
			spec.Name = m[sbsModel.Field.F_name].(string)
			var sub []cards.SubSpecs
			for _, sbsMap := range sbsMaps {
				if sbsMap[sbsModel.Field.F_p_spec_id].(string) == m[sbsModel.Field.F_spec_id].(string) {
					var item cards.SubSpecs
					id, _ := strconv.Atoi(sbsMap[sbsModel.Field.F_spec_id].(string))
					item.Id = id
					item.Name = sbsMap[sbsModel.Field.F_name].(string)
					sub = append(sub, item)
				}
			}
			spec.Sub = sub
			specs = append(specs, spec)
		}
	}
	var list []cards.SubServer
	for _, sspMap := range sspMaps {
		for _, m := range maps {
			if sspMap[sspModel.Field.F_ssp_id].(string) == m[model.Field.F_ssp_id].(string) {
				price, _ := strconv.ParseFloat(m[model.Field.F_price].(string), 64)
				sspid, _ := strconv.Atoi(sspMap[sspModel.Field.F_ssp_id].(string))
				list = append(list, cards.SubServer{
					SubServerIds: sspMap[sspModel.Field.F_spec_ids].(string),
					Price:        price,
					SspId:        sspid,
				})
				break
			}
		}
	}
	reply.SubServer = &list
	reply.Specs = &specs
	return
}

//
func (s *SingleLogic) GetBySspids(sspIds *[]int, reply *[]cards.ReplyGetBySspids) (err error) {
	if len(*sspIds) == 0 {
		return nil
	}

	mSSP := new(models.SingleSpecPriceModel).Init()
	r := mSSP.GetBySspids(*sspIds)
	mapstructure.WeakDecode(r, reply)
	return
}

// 根据shopId和批量组合规格查询-rpc确认消费
func (s *SingleLogic) GetByShopSspIds(shopId int, sspId []int) (reply []map[string]interface{}) {
	reply = make([]map[string]interface{}, 0)
	if shopId <= 0 {
		return
	}
	spSpechPriceModel := new(models.ShopSingleSpecPriceModel).Init()
	reply = spSpechPriceModel.GetByShopSspIds(shopId, sspId)
	return
}

// 根据singleids批量获取门店单项目-rpc确认消费
func (s *SingleLogic) GetByShopSingle(shopId int, singleIds []int, status int) (reply []map[string]interface{}) {
	reply = make([]map[string]interface{}, 0)
	if len(singleIds) <= 0 {
		return
	}
	shopSingleModel := new(models.ShopSingleModel).Init()
	if status > 0 {
		reply = shopSingleModel.GetByShopIdAndSingleIds(shopId, singleIds, status)
	} else {
		reply = shopSingleModel.GetByShopIdAndSingleIds(shopId, singleIds)
	}

	return
}

// 根据singleids批量获取门店单项目-rpc确认消费
func (s *SingleLogic) GetPriceByShopIdAndSsspId(shopId int, sspIds []int) (ssspmMaps []map[string]interface{}) {
	if len(sspIds) <= 0 {
		return
	}
	ssspm := new(models.ShopSingleSpecPriceModel).Init()
	ssspmMaps = ssspm.GetByShopidAndSspids(shopId, sspIds)
	if len(ssspmMaps) == 0 {
		return
	}
	ssids := functions.ArrayValue2Array(ssspm.Field.F_ss_id, ssspmMaps)
	shopSingleModel := new(models.ShopSingleModel).Init()
	shopSingleMaps := shopSingleModel.GetBySsids(ssids)
	for index, shopSingleSpecPrice := range ssspmMaps {
		for _, shopSingle := range shopSingleMaps {
			if shopSingleSpecPrice[ssspm.Field.F_ss_id].(string) == shopSingle[shopSingleModel.Field.F_ss_id].(string) {
				ssspmMaps[index]["SingleId"] = shopSingle[shopSingleModel.Field.F_single_id].(string)
				break
			}
		}
	}
	return
}

// 根据singleids批量获取单项目-rpc确认消费
func (s *SingleLogic) GetBySingle(singleIds []int) (reply []map[string]interface{}) {
	reply = make([]map[string]interface{}, 0)
	if len(singleIds) <= 0 {
		return
	}
	singleModel := new(models.SingleModel).Init()
	reply = singleModel.GetBySingleids(singleIds)
	return
}

// 九百岁首页精选服务
func (s *SingleLogic) GetSelectServices(ctx context.Context, args *cards.ArgsGetSelectServices, reply *[]cards.ReplyGetSelectServices) error {
	*reply = []cards.ReplyGetSelectServices{}
	if args.Cid <= 0 {
		return nil
	}
	rpcEs := new(cards2.ShopCards).Init()
	defer rpcEs.Close()
	var replyEs map[string]interface{}
	if err := rpcEs.SearchItem(ctx, &cards.ArgsAppInfos{
		Paging: common.Paging{Page: 0, PageSize: 1}, Cid: args.Cid, Lon: args.Lng, Lat: args.Lat, Num: args.Num, Flag: 1,
	}, &replyEs); err != nil {
		return err
	}
	if len(replyEs) == 0 || replyEs["result"] == nil {
		return nil
	}
	result := replyEs["result"].([]interface{})
	var ssids []int
	var singleIds []int
	for i := range result {
		source := result[i].(map[string]interface{})
		for _, item := range source["Sub"].([]interface{}) {
			it := item.(map[string]interface{})
			ssids = append(ssids, int(it["ShopItemId"].(float64)))
			singleIds = append(singleIds, int(it["ItemId"].(float64)))
			*reply = append(*reply, cards.ReplyGetSelectServices{
				SingleId:   int(it["ItemId"].(float64)),
				Ssid:       int(it["ShopItemId"].(float64)),
				SingleName: it["ItemName"].(string),
				RealPrice:  it["ItemPrice"].(float64),
				ShopId:     int(source["ShopId"].(float64)),
			})
		}
	}
	if len(ssids) == 0 || len(singleIds) == 0 {
		return nil
	}
	ssModel := new(models.ShopSingleModel).Init()
	ssMaps := ssModel.GetBySsids(ssids)
	if len(ssMaps) == 0 {
		return nil
	}
	for _, ssMap := range ssMaps {
		for i, service := range *reply {
			if ssMap[ssModel.Field.F_ss_id].(string) == strconv.Itoa(service.Ssid) {
				sales, _ := strconv.Atoi(ssMap[ssModel.Field.F_sales].(string))
				(*reply)[i].Sales = sales
				if len(ssMap[ssModel.Field.F_name].(string)) > 0 {
					(*reply)[i].SingleName = ssMap[ssModel.Field.F_name].(string)
				}
				break
			}
		}
	}
	sModel := new(models.SingleModel).Init()
	sMaps := sModel.GetBySingleids(singleIds, []string{sModel.Field.F_single_id, sModel.Field.F_img_id, sModel.Field.F_price})
	var imgIds = make([]int, 0, len(sMaps))
	for _, sMap := range sMaps {
		imgId, _ := strconv.Atoi(sMap[sModel.Field.F_img_id].(string))
		imgIds = append(imgIds, imgId)
	}
	rpcFile := new(file.Upload).Init()
	defer rpcFile.Close()
	var replyFile map[int]file2.ReplyFileInfo
	rpcFile.GetImageByIds(ctx, imgIds, &replyFile)
	// 默认图片
	replyFile[0] = file2.ReplyFileInfo{
		Id:   0,
		Hash: "",
		Path: constkey.CardsSmallDefaultPics[cards.ITEM_TYPE_single],
	}

	sMap := functions.ArrayRebuild(sModel.Field.F_single_id, sMaps)
	for i, service := range *reply {
		if v, ok := sMap[strconv.Itoa(service.SingleId)]; ok {
			imgId, _ := strconv.Atoi(v.(map[string]interface{})[sModel.Field.F_img_id].(string))
			(*reply)[i].ImgId = imgId
			(*reply)[i].ImgPath = replyFile[imgId].Path
			price, _ := strconv.ParseFloat(v.(map[string]interface{})[sModel.Field.F_price].(string), 64)
			(*reply)[i].Price = price
		}
	}
	return nil
}
