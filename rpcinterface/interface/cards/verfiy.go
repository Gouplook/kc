//参数统一验证文件
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/15 15:05
package cards

import (
	"git.900sui.cn/kc/kcgin"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	//是否上架 0=否 1=是
	SINGLE_IS_GROUND_no  = 0
	SINGLE_IS_GROUND_yes = 1
	//图片类型 1=相册 2=服务效果图片 3=仪器设备图片
	IMG_TYPE_album  = 1
	IMG_TYPE_effect = 2
	IMG_TYPE_tool   = 3
	//规格价格是否删除
	SPEC_PRICE_IS_DEL_no  = 0
	SPEC_PRICE_IS_DEL_yes = 1
	//单项目是否存在多规格 0=否 1=是
	SINGLE_HAS_SPEC_no  = 0
	SINGLE_HAS_SPEC_yes = 1

	//上下架操作 1=上架 2=下架
	OPT_UP   uint8 = 1
	OPT_DOWN uint8 = 2

	//是否有赠送服务 0=没有 1=有
	HAS_GIVE_SINGLE_no  int = 0
	HAS_GIVE_SINGLE_yes int = 1

	//项目在总店的上下架状态 0=下架 1=上架
	IS_GROUND_no  = 0
	IS_GROUND_yes = 1

	// 项目在总店的状态 0=否，1=是
	IS_BUS_DEL_no  = 0
	IS_BUS_DEL_yes = 1
	BUS_STATUS_DEL = 3 // 被总店已删除

	//项目在分店的状态
	STATUS_OFF_SALE = 1 // 下架
	STATUS_ON_SALE  = 2 // 上架
	STATUS_DISABLE  = 3 // 禁用
	STATUS_DELETE   = 4 // 删除

	//充值卡折扣类型 1=面值折扣 2=卡项折扣
	DISCOUNT_TYPE_price = 1
	DISCOUNT_TYPE_item  = 2

	//充值类型 0=购卡充值 1=复充值
	Recharge_Type_Buy_Card       = 0
	Recharge_Type_Dupli_Recharge = 1

	// 充值卡规则是否删除，0=否 1=是
	RECHARGE_TYPE_NO  = 0
	RECHARGE_TYPE_YES = 1

	// 添加充值卡规则最大条数
	RCARD_MAX_NUM = 10

	//折扣率
	DISCOUNTMAX = 20
	DICOUNTMIN  = 1

	// 项目折扣 1-无折扣；2-有折扣
	SINGLE_DISCOUNT_no  = 1
	SINGLE_DISCOUNT_yes = 2

	// 购卡或充值须100起
	BUY_CRARD_MIN_AMOUNT = 100

	//是否删除
	IS_DEL_NO  = 0 //否
	IS_DEL_YES = 1 //是

	// 门店卡项关联是否删除
	RELATION_IS_DEL_NO = 0  //否
	RELATION_IS_DEL_YES = 1 //是

	//卡项是否删除
	ITEM_IS_DEL_NO  = 0 //否
	ITEM_IS_DEL_YES = 1 //是
	//是否同步身份卡折扣
	IS_SYNC_NO  = 0 //否
	IS_SYNC_YES = 1 //是

	// 是否适用所有门店 0=否，1=是
	IS_USE_SHOP_NO  = 0
	IS_USE_SHOP_YES = 1

	// 是否永久有效：1-是；2-否
	IS_PERMANENT_YES = 1
	IS_PERMANENT_NO  = 2

	// 单项目是否享有折扣 1-无折扣；2-有折扣 折扣值为10.0 表示折扣
	IS_HAVE_NO  = 1
	IS_HAVE_YES = 2
	IS_HAVE_NO_DISCOUNT = 10.0
)

//包含的单项目
type IncSingle struct {
	SingleID         int `mapstructure:"single_id" json:"SingleId"` //单项目id
	Num              int `mapstructure:"num"`                       //单项目次数
	SspId            int `mapstructure:"ssp_id"`                    //规格id
	PeriodOfValidity int `mapstructure:"period_of_validity"`        //有效期，单位天(只有赠送项目有效)
}

//赠品描述
type GiveSingleDesc struct {
	Desc string `mapstructure:"desc" json:"desc"`
}

//包含的单项目详情
type IncSingleDetail struct {
	IncSingle
	Name        string //套餐名称
	Price       string //标价
	RealPrice   string //售价
	ImgId       int    //图片id
	ImgUrl      string //图片地址
	ServiceTime int    //服务时长
}

//包含的单项目详情
type IncSingleDetail2 struct {
	IncSingle
	Name        string  //套餐名称
	Price       string  //标价
	RealPrice   string  //售价
	ImgId       int     //图片id
	ImgUrl      string  //图片地址
	ServiceTime int     //服务时长
	SpecNames   string  //规格组合名称
	Discount    float64 //身份卡才有该值
}

//包含的单项目(无限次)
type IncInfSingle struct {
	SingleID int `mapstructure:"single_id" json:"SingleId"` //单项目ID
}

//包含的单项目详情(无限次)
type IncInfSingleDetail struct {
	IncInfSingle
	Name        string //套餐名称
	Price       string //标价
	RealPrice   string //售价
	ImgId       int    //图片id
	ImgUrl      string //图片地址
	ServiceTime int    //服务时长
}

//包含的商品
type IncProduct struct {
	ProductID int `mapstructure:"product_id"` //商品ID
}

//包含的商品详情
type IncProductDetail struct {
	IncProduct
	Name      string `mapstructure:"name"`       //商品名称
	SpecPrice string `mapstructure:"spec_price"` //商品价格
	ImgId     int    //图片id
	ImgUrl    string //图片地址
}

//包含的商品详情2
type IncProductDetail2 struct {
	IncProduct
	Name      string  `mapstructure:"name"`       //商品名称
	SpecPrice string  `mapstructure:"spec_price"` //商品价格
	Discount  float64 //折扣价格
	ImgId     int     //图片id
	ImgUrl    string  //图片地址
}

//验证售价，标价
func VerfiyPrice(realPrice float64, price float64) error {
	if realPrice <= 0 {
		return common.GetInterfaceError(common.CARDS_PRICE_ERR)
	}

	return nil
}

//验证保险期限
func VerfiyServicePeriod(servicePeriod int) error {
	maxServicePeriod, err := kcgin.AppConfig.Int("card.maxServicePeriod")
	if err != nil {
		maxServicePeriod = 48
	}
	if servicePeriod < 0 || servicePeriod > maxServicePeriod {
		return common.GetInterfaceError(common.CARDS_PERIOD_ERR)
	}
	return nil
}

//验证项目名称
func VerfiyName(name string) error {
	if len(name) == 0 {
		return common.GetInterfaceError(common.CARDS_NAME_EMP_ERR)
	}
	return nil
}

//验证包含单项目的数量
func VerifySinglesNum(singleNum int) error {
	if singleNum == 0 {
		//return nil
		return common.GetInterfaceError(common.CARDS_INC_SINGLES_EMP_ERR)
	}
	maxSingleNum, err := kcgin.AppConfig.Int("card.maxSingleNum")
	if err != nil {
		maxSingleNum = 20
	}
	if singleNum > maxSingleNum {
		return common.GetInterfaceError(common.CARDS_INC_SINGLES_MAX_ERR)
	}
	return nil
}

//验证包含商品的数量
func VerifyProductsNum(productNum int) error {
	if productNum == 0 {
		//return nil
		return common.GetInterfaceError(common.CARDS_INC_PRODUCTS_EMP_ERR)
	}
	maxProductNum, err := kcgin.AppConfig.Int("card.maxProductNum")
	if err != nil {
		maxProductNum = 20
	}
	if productNum > maxProductNum {
		return common.GetInterfaceError(common.CARDS_INC_PRODUCTS_MAX_ERR)
	}
	return nil
}

//验证赠送的项目数量
func VerifyGiveSinglesNum(singleNum int) error {
	if singleNum == 0 {
		return nil
	}
	maxSingleNum, err := kcgin.AppConfig.Int("card.maxGiveSingleNum")
	if err != nil {
		maxSingleNum = 10
	}
	if singleNum > maxSingleNum {
		return common.GetInterfaceError(common.CARDS_INC_SINGLES_MAX_ERR)
	}
	return nil
}

// 验证分店是否允许
func VerifyStatus(status int) error {
	if status != STATUS_ON_SALE {
		return common.GetInterfaceError(common.CARDS_STATUS_OFF_ERR)
	}
	return nil
}

// 验证分店是否上架
func VerifyGround(ground int) error {
	if ground != IS_GROUND_yes {
		return common.GetInterfaceError(common.CARDS_STATUS_OFF_ERR)
	}
	return nil
}

//验证充值卡的折扣信息
func VerifyRcardDiscount(discountType int, price float64, realPrice float64, discount float64) error {
	//不同折扣类型的折扣方式验证
	if discountType < DISCOUNT_TYPE_price || discountType > DISCOUNT_TYPE_item {
		return common.GetInterfaceError(common.RCARD_DISCOUNT_TYPE_ERR)
	}
	//面值折扣
	if discountType == DISCOUNT_TYPE_price {
		//实际金额不能大于等于面值
		if realPrice > price+realPrice {
			return common.GetInterfaceError(common.RCARD_DISCOUNT_PRICE_ERR)
		}
	} else { //项目折扣
		//if decimal.NewFromFloat(price).Cmp(decimal.NewFromFloat(realPrice)) != 0 {
		//	return tools.GetInterfaceError(tools.RCARD_DISCOUNT_ITEM_ERR)
		//}
		if discount < 1 && discount > 10 {
			return common.GetInterfaceError(common.RCARD_DISCOUNT_ERR)
		}
	}

	return nil
}
