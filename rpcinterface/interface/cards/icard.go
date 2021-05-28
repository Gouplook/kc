package cards

import (
	"context"

	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//InputParams 入参
type InputParams struct {
	common.Input
	common.BsToken
}

//OutputReply 出参
type OutputReply struct {
	common.Output
}

//PageOutputReply 出参
type PageOutputReply struct {
	common.PaginationOutput
}

//ICardBase ICardBase
type ICardBase struct {
	Name           string  `mapstructure:"name" form:"name" json:"Name"`                               // 名称
	BusID          int     `mapstructure:"busId" form:"busId" json:"BusId"`                            // 商户ID
	ShortDesc      string  `mapstructure:"shortDesc"  form:"shortDesc" json:"SortDesc"`                // 短描述short
	RealPrice      float64 `mapstructure:"realPrice" form:"realPrice" json:"RealPrice"`                // 现价
	Price          float64 `mapstructure:"price" form:"price" json:"Price"`                            // 标价
	ServicePeriod  int     `mapstructure:"servicePeriod" form:"servicePeriod" json:"ServicePeriod"`    // 保险时间 月
	SaleShopNum    int     `mapstructure:"saleShopNum" form:"saleShopNum" json:"SaleShopNum"`          // 在售门店数量
	ImgID          int     `mapstructure:"imgId" form:"imgId" json:"ImgID"`                            // 图片ID
	Sales          int     `mapstructure:"sales" form:"sales" json:"Sales"`                            // 销量
	Ctime          string  `mapstructure:"ctime" form:"ctime" json:"Ctime"`                            // 发布时间
	Click          int     `mapstructure:"click" form:"click" json:"Click"`                            // 点击量
	SaleShopCount  int     `mapstructure:"saleShopCount" form:"saleShopCount" json:"SaleShopCount"`    // 在售门店数
	HasGiveSignle  int     `mapstructure:"hasGiveSignle" form:"hasGiveSignle" json:"HasGiveSignle"`    // 是否有赠送单服务
	IsGround       int     `mapstructure:"isGround" form:"isGround" json:"IsGround"`                   // 主表状态 是否上架 0=否 1=是
	IcardID        int     `mapstructure:"icardId" form:"icardId" json:"IcardId"`                      // 身份卡id
	Status         int     `mapstructure:"status" form:"status" json:"Status"`                         // 身份卡在门店的状态 1=下架 2=上架 3=被总店禁用
	ShopIsDel      int     `mapstructure:"ShopIsDel" form:"shopIsDel" json:"ShopIsDel"`                //在门店的删除状态
	IsSelfShop     int     `mapstructure:"isSelfShop" form:"isSelfShop" json:"IsSelfShop"`             // 已添加至本店 1=yes 0=no
	IsDel          int     `mapstructure:"isDel" form:"isDel" json:"isDel"`                            //是否删除 0-否；1-是
	SsID           int     `mapstructure:"ssId" form:"ssId" json:"SsId"`                               // 身份卡表和门店关联主键
	DiscountSingle float64 `mapstructure:"discountSingle" form:"discountSingle" json:"DiscountSingle"` // 折扣
	DiscountGoods  float64 `mapstructure:"discountGoods" form:"discountGoods" json:"DiscountGoods"`    // 折扣
	IsAllSingle    bool    `mapstructure:"isAllSingle" form:"isAllSingle" json:"IsAllSingle"`          // 是否全部单项目
	ApplySingleNum int     `mapstructure:"applySingleNum" form:"applySingleNum" json:"applySingleNum"` //适用单项目的个数
	GiveSingleNum  int     `mapstructure:"giveSingleNum" form:"giveSingleNum" json:"giveSingleNum"`    //赠送单项目的个数
	IsAllProduct   bool    `mapstructure:"isAllProduct" form:"isAllProduct" json:"IsAllProduct"`       // 是否全部商品
	ShopItemId     int     `mapstructure:"shopItemId" form:"shopItemId" json:"ShopItemId"`             // 是否全部商品

}

//OutputICardList OutputICardList
type OutputICardList struct {
	common.Pagination
	Lists []ICardBase
}

//InputParamsICardList InputParamsICardList
type InputParamsICardList struct {
	common.PaginationInput
	common.BsToken
	// SearchShopID int `mapstructure:"searchShopId" form:"searchShopId" json:"SearchShopId"` // SearchShopId
	Status           int    `mapstructure:"status" form:"status" json:"Status"`                               // Status 身份卡在门店的状态 1=下架 2=上架 3=被总店禁用
	Ground           int    `mapstructure:"ground" form:"ground" json:"Ground"`                               // Ground 身份卡在门店的状态 1=下架 2=上架 3=被总店禁用
	ShopID           int    `mapstructure:"shopId" form:"shopId" json:"ShopId"`                               // ShopId
	IsDel            string `mapstructure:"isDel" form:"isDel" json:"isDel"`                                  //过滤删除的数据：""-全部； "0"-否；"1"-是
	FilterShopHasAdd bool   `mapstructure:"filterShopHasAdd" form:"filterShopHasAdd" json:"filterShopHasAdd"` //false-获取全部，true-过滤添加过的数据
}

//InputParamsICardSave InputParamsICardSave
type InputParamsICardSave struct {
	common.BsToken
	ICardBase
	Notes           string `mapstructure:"notes" form:"notes" json:"Notes"`                               //温馨提示
	IsGround        int    `mapstructure:"isGround" form:"isGround" json:"IsGround"`                      // 是否上架 0=否 1=是
	IncludeSingles  string `mapstructure:"includeSingles" form:"includeSingles" json:"IncludeSingles"`    // 搭配服务-单项目
	IncludeProducts string `mapstructure:"includeProducts" form:"includeProducts" json:"IncludeProducts"` // 搭配服务-商品
	GiveSingles     string `mapstructure:"giveSingles" form:"giveSingles" json:"GiveSingles"`             // 赠送服务
	ImgHash         string `mapstructure:"imgHash" form:"imgHash" json:"ImgHash"`                         // 图片hash
	GiveSingleDesc  string `mapstructure:"giveSingleDesc" form:"giveSingleDesc" json:"GiveSingleDesc"`    // 赠品描述
	IsSync          int    `mapstructure:"isSync" form:"isSync" json:"isSync"`                            //是否同步身份卡折扣，0-否；1-是
}

//OutputParamsICardSave OutputParamsICardSave
type OutputParamsICardSave struct {
	IcardID int
}

//InputParamsICardInfo InputParamsICardInfo
type InputParamsICardInfo struct {
	common.BsToken
	IcardID int `mapstructure:"icardId" form:"icardId" json:"IcardId"` // 身份卡id
	ShopID  int `mapstructure:"shopId" form:"shopId" json:"ShopId"`    // ShopId
}

//OutputParamsICardInfo OutputParamsICardInfo
type OutputParamsICardInfo struct {
	ICardBase
	ShareLink string //分享链接
	ImgHash   string `json:"ImgHash"` // img hash
	ImgURL    string `json:"ImgUrl"`  // img url
	// GiveSingleDesc       []string               `json:"GiveSingleDesc"` // 赠品描述
	// ICardInfoIcardGive   []ICardInfoIcardGive   // 附赠单项目
	IsAllSingle          bool                   //适用于全部单项目
	IsAllProduct         bool                   //适用于全部商品
	ICardInfoIcardGoods  []ICardInfoIcardGoods  `json:"IncProducts"`    // 服务-商品
	ICardInfoIcardSingle []ICardInfoIcardSingle `json:"IncludeSingles"` // 服务-单项目
	ShopLists            []ReplyShopName        // 总店身份卡店添加信息
	DiscountLists        []float64              // 适用所有=[],适用部分项目=[min,max]
	BusInfo
}

//ICardInfoIcardGive ICardInfoIcardGive
type ICardInfoIcardGive struct {
	Discount float64
	Name     string
	Price    float64
	SingleID int
}

//ICardInfoIcardGoods ICardInfoIcardGoods
type ICardInfoIcardGoods struct {
	Discount float64 `json:"Discount"`
	Name     string  `json:"Name"`
	Price    string  `json:"SpecPrice"` // 标价
	GoodsID  int     `json:"ProductID"`
	ImgURL   string  `json:"ImgUrl"`
	ImgID    int     `json:"ImgId"`
}

//ICardInfoIcardSingle ICardInfoIcardSingle
type ICardInfoIcardSingle struct {
	Discount    float64 `json:"Discount"`
	Name        string  `json:"Name"`
	Price       string  `json:"SpecPrice"` // 标价
	SingleID    int     `json:"SingleId"`
	ImgURL      string  `json:"ImgUrl"`
	ImgID       int     `json:"ImgId"`
	ServiceTime int     `json:"ServiceTime"`
}

//InputParamsICardSetOnOff InputParamsICardSetOnOff
type InputParamsICardSetOnOff struct {
	common.BsToken
	IcardID  int    `mapstructure:"icardId" form:"icardId" json:"IcardId"`    // 身份卡id
	IcardIds string `mapstructure:"icardIds" form:"icardIds" json:"IcardIds"` // 身份卡ids
}

//OutputParamsICardSetOnOff OutputParamsICardSetOnOff
type OutputParamsICardSetOnOff struct {
	IcardIds []int `mapstructure:"icardIds" form:"icardIds" json:"IcardIds"` // 身份卡id
	Rows     int   `mapstructure:"rows" form:"rows" json:"Rows"`             // 受影响数
}

//InputParamsICardPush InputParamsICardPush
type InputParamsICardPush struct {
	common.BsToken
	IcardIds string `mapstructure:"icardIds" form:"icardIds" json:"IcardIds"` // 身份卡id
	ShopIds  string `mapstructure:"shopIds" form:"shopIds" json:"ShopIds"`    // shopIds
}

//OutputParamsICardPush OutputParamsICardPush
type OutputParamsICardPush struct {
	IcardIds []int `mapstructure:"icardIds" form:"icardIds" json:"IcardId"` // 身份卡id
}

//InputParamsICardAddToShop InputParamsICardAddToShop
type InputParamsICardAddToShop struct {
	common.BsToken
	IcardIds string `mapstructure:"icardIds" form:"icardIds" json:"IcardIds"` // 身份卡id
}

//OutputParamsICardAddToShop OutputParamsICardAddToShop
type OutputParamsICardAddToShop struct {
	IcardIds []int `mapstructure:"icardIds" form:"icardIds" json:"IcardId"` // 身份卡id
}

//InputParamsICardCanUse InputParamsICardCanUse
type InputParamsICardCanUse struct {
	common.BsToken
	ShopID    int    `mapstructure:"shopId" form:"shopId" json:"ShopId"`          // ShopId
	UID       int    `mapstructure:"uId" form:"uId" json:"uId"`                   // uid
	GoodsIds  string `mapstructure:"goodsIds" form:"goodsIds" json:"GoodsIds"`    // 商品id
	SingleIds string `mapstructure:"singleIds" form:"singleIds" json:"SingleIds"` // 单项目id
}

//InputParamsICardCanUseForUser InputParamsICardCanUseForUser
type InputParamsICardCanUseForUser struct {
	common.Utoken
	ShopID    int    `mapstructure:"shopId" form:"shopId" json:"ShopId"`          // ShopId
	GoodsIds  string `mapstructure:"goodsIds" form:"goodsIds" json:"GoodsIds"`    // 商品id
	SingleIds string `mapstructure:"singleIds" form:"singleIds" json:"SingleIds"` // 单项目id
}

//CardInfo CardInfo
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

//获取iCard企业基本信息出参-风控统计用
type ReplyGetBusBaseInfoRpc struct {
	BusId  int
	BindId int
}

//删除身份卡-入参
type InputParamsDelete struct {
	common.BsToken
	IcardIds string `mapstructure:"icardIds" form:"icardIds" json:"IcardIds"`
}

//获取身份卡的折扣信息-出参
type ReplyGetIcardDiscountById struct {
	ICardId         int
	ProductDiscount []IcardProductDiscountParams //商品折扣信息
	SingleDiscount  []IcardSingleDiscountParams  //项目折扣信息
}
type IcardProductDiscountParams struct {
	GoodsId  int     //商品id，0代表适用于全部
	Name     string  //商品名字
	DisCount float64 //商品折扣
}
type IcardSingleDiscountParams struct {
	SingleId int     //项目id，0代表适用于全部
	Name     string  //单项目名字
	Discount float64 //项目折扣
}

//获取身份卡备份表中的项目折扣-入参
type ArgsGetICardSingleDiscount struct {
	ICardId     int
	IsSync      string
	RequestType int //类型：0-获取项目折扣；1-获取产品折扣
}
type GetICardSingleDiscountBase struct {
	ICardId    int
	SingleId   int
	SingleName string
	GoodsId    int
	GoodsName  string
	Discount   float64
}

//获取身份卡备份表中的项目折扣-出参
type ReplyGetICardSingleDiscount struct {
	SingleDiscount []GetICardSingleDiscountBase
}

//ICard 折扣卡管理（身份卡）
type ICard interface {
	List(ctx context.Context, args *InputParamsICardList, reply *OutputICardList) error
	Save(ctx context.Context, args *InputParamsICardSave, reply *OutputParamsICardSave) error
	Info(ctx context.Context, args *InputParamsICardInfo, reply *OutputParamsICardInfo) error
	ShopList(ctx context.Context, args *InputParamsICardList, reply *OutputICardList) error
	OurShopList(ctx context.Context, args *InputParamsICardList, reply *OutputICardList) error
	Delete(ctx context.Context, args *InputParamsDelete, reply *bool) error
	Push(ctx context.Context, args *InputParamsICardPush, reply *OutputParamsICardPush) error
	SetOnOff(ctx context.Context, args *InputParamsICardSetOnOff, reply *OutputParamsICardSetOnOff) error
	SetOn(ctx context.Context, args *InputParamsICardSetOnOff, reply *OutputParamsICardSetOnOff) error
	SetOff(ctx context.Context, args *InputParamsICardSetOnOff, reply *OutputParamsICardSetOnOff) error
	ShopSetOnOff(ctx context.Context, args *InputParamsICardSetOnOff, reply *OutputParamsICardSetOnOff) error
	ShopSetOn(ctx context.Context, args *InputParamsICardSetOnOff, reply *OutputParamsICardSetOnOff) error
	ShopSetOff(ctx context.Context, args *InputParamsICardSetOnOff, reply *OutputParamsICardSetOnOff) error
	AddToShop(ctx context.Context, args *InputParamsICardAddToShop, reply *OutputParamsICardAddToShop) error
	UserICardList(ctx context.Context, args *InputParams, reply *OutputReply) error
	CanUseICardList(ctx context.Context, args *InputParamsICardCanUse, reply *OutputParamsICardCanUse) error
	CanUseICardListForUser(ctx context.Context, args *InputParamsICardCanUseForUser, reply *OutputParamsICardCanUse) error
	GetBusBaseInfoRpc(ctx context.Context, iCardId *int, reply *ReplyGetBusBaseInfoRpc) error
	//获取身份卡的折扣信息
	GetIcardDiscountById(ctx context.Context, iCardId *int, reply *ReplyGetIcardDiscountById) error
	//获取身份卡备份表中的项目折扣
	GetICardSingleDiscount(ctx context.Context, args *ArgsGetICardSingleDiscount, reply *ReplyGetICardSingleDiscount) error
}
