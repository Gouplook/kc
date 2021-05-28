package order

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//InputParamsICardCanUse InputParamsICardCanUse
type InputParamsICardCanUse struct {
	common.BsToken
	ShopID    int    `mapstructure:"shopId" form:"shopId" json:"ShopId"`          // ShopId
	UID       int    `mapstructure:"uid" form:"uid" json:"uid"`                   // uid
	GoodsIds  string `mapstructure:"goodsIds" form:"goodsIds" json:"GoodsIds"`    // 商品id
	SingleIds string `mapstructure:"singleIds" form:"singleIds" json:"SingleIds"` // 单项目id
}

//CardInfo CardInfo1
type CardInfo struct {
	CardPackageID int     `mapstructure:"cardPackageID" form:"cardPackageID" json:"CardPackageId"` // card_package_id
	Name          string  `mapstructure:"name" form:"name" json:"Name"`                            // 名称
	Discount      float64 `mapstructure:"discount" form:"discount" json:"Discount"`                // 折扣
	IcardID       int     `mapstructure:"icardId" form:"icardId" json:"IcardId"`                   // 身份卡ID
	ImgID         int     `mapstructure:"imgId" form:"imgId" json:"ImgId"`                         // 卡包封面ID
	ImgURL        string  `mapstructure:"imgUrl" form:"imgUrl" json:"ImgUrl"`                      // 卡包封面URL
	ExpireTime    int64   `mapstructure:"expireTime" form:"expireTime" json:"ExpireTime"`          // 到期时间
	ExpireSurDay  int64   `mapstructure:"expireSurDay" form:"expireSurDay" json:"ExpireSurDay"`    // 剩余天数
	ServicePeriod int     `mapstructure:"servicePeriod" form:"servicePeriod" json:"ServicePeriod"` // 有效期（单位：月）
}

//OutputParamsICardCanUse OutputParamsICardCanUse
type OutputParamsICardCanUse struct {
	SingleIcardList []CardInfo
	GoodsIcardList  []CardInfo
}

//IcardOrder IcardOrder
type IcardOrder interface {
	//根据用户id查询身份卡列表
	GetIcardListByUserID(ctx context.Context, args *InputParamsICardCanUse, reply *OutputParamsICardCanUse) error
}
