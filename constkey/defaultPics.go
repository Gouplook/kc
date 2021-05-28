package constkey

import (
	"git.900sui.cn/kc/kcgin"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

var CardsDefaultPics map[int]string
var CardsSmallDefaultPics map[int]string

func init() {
	CardsDefaultPics = map[int]string{
		cards.ITEM_TYPE_single: kcgin.AppConfig.String("single.defaultPic"),
		cards.ITEM_TYPE_sm:     kcgin.AppConfig.String("sm.defaultPic"),
		cards.ITEM_TYPE_card:   kcgin.AppConfig.String("card.defaultPic"),
		cards.ITEM_TYPE_hcard:  kcgin.AppConfig.String("hcard.defaultPic"),
		cards.ITEM_TYPE_ncard:  kcgin.AppConfig.String("ncard.defaultPic"),
		cards.ITEM_TYPE_hncard: kcgin.AppConfig.String("hncard.defaultPic"),
		cards.ITEM_TYPE_rcard:  kcgin.AppConfig.String("rcard.defaultPic"),
		cards.ITEM_TYPE_icard:  kcgin.AppConfig.String("icard.defaultPic"),
	}

	CardsSmallDefaultPics = map[int]string{
		cards.ITEM_TYPE_single: kcgin.AppConfig.String("single.defaultSamllPic"),
		cards.ITEM_TYPE_sm:     kcgin.AppConfig.String("sm.defaultSamllPic"),
		cards.ITEM_TYPE_card:   kcgin.AppConfig.String("card.defaultSamllPic"),
		cards.ITEM_TYPE_hcard:  kcgin.AppConfig.String("hcard.defaultSamllPic"),
		cards.ITEM_TYPE_ncard:  kcgin.AppConfig.String("ncard.defaultSamllPic"),
		cards.ITEM_TYPE_hncard: kcgin.AppConfig.String("hncard.defaultSamllPic"),
		cards.ITEM_TYPE_rcard:  kcgin.AppConfig.String("rcard.defaultSamllPic"),
		cards.ITEM_TYPE_icard:  kcgin.AppConfig.String("icard.defaultSamllPic"),
	}
}
