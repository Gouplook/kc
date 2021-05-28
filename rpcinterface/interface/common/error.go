/**
 * @Author: Gosin
 * @Date: 2020/4/9 16:20
 */
package common

import "git.900sui.cn/kc/base/common/toolLib"

const (
	// 公共服务验证
	ENCODE_IS_NIL       = "1000001"
	ENCODE_ERR          = "1000002"
	ENCODE_DATA_TIMEOUT = "1000003"
	PERMISSION_ERR      = "1000004"
	PHONE_VERIFY_ERR    = "1000005"

	// 用户服务验证
	CHANNEL_ERR = "1100001"
	DEVICE_ERR  = "1100002"

	// 商户服务验证
	MEMBER_SOURCE_ERROR    = "1200001"
	MEMBER_GENDER_ERROR    = "1200002"
	MEMBER_CARD_NUM_ERROR  = "1200003"
	MEMBER_DATA_TYPE_ERROR = "1200004"
	MEMBER_NAME_IS_NIL     = "1200005"

	//卡项服务
	CARDS_PRICE_ERR            = "1300001"
	CARDS_PERIOD_ERR           = "1300002"
	CARDS_NAME_EMP_ERR         = "1300003"
	CARDS_INC_SINGLES_EMP_ERR  = "1300004"
	CARDS_INC_SINGLES_MAX_ERR  = "1300005"
	CARDS_INC_PRODUCTS_EMP_ERR = "1300006"
	CARDS_INC_PRODUCTS_MAX_ERR = "1300007"
	CARDS_STATUS_OFF_ERR       = "1300008"

	//充值卡
	RCARD_DISCOUNT_TYPE_ERR  = "1400001"
	RCARD_DISCOUNT_PRICE_ERR = "1400002"
	RCARD_DISCOUNT_ITEM_ERR  = "1400003"
	RCARD_DISCOUNT_ERR       = "1400004"

	//二维码
	QRCODE_EXPIRED = "1500001"
)

var (
	errMsg = map[string]string{
		// 公共服务验证
		ENCODE_IS_NIL:       "EncodeStr数据为空",
		ENCODE_ERR:          "解密失败",
		ENCODE_DATA_TIMEOUT: "数据已过期",
		PERMISSION_ERR:      "没有操作权限",
		PHONE_VERIFY_ERR:    "手机号错误",

		// 用户服务验证
		CHANNEL_ERR: "非法渠道",
		DEVICE_ERR:  "非法设备",

		// 商户服务验证
		MEMBER_SOURCE_ERROR:    "会员来源错误",
		MEMBER_GENDER_ERROR:    "会员性别错误",
		MEMBER_CARD_NUM_ERROR:  "证件号码错误",
		MEMBER_DATA_TYPE_ERROR: "错误的数据类型",
		MEMBER_NAME_IS_NIL:     "会员名称不能为空",

		//卡项服务
		CARDS_PRICE_ERR:            "价格设置错误",
		CARDS_PERIOD_ERR:           "服务期限设置错误",
		CARDS_NAME_EMP_ERR:         "名称错误",
		CARDS_INC_SINGLES_EMP_ERR:  "未添加包含的单项目",
		CARDS_INC_SINGLES_MAX_ERR:  "超过最多数量限制",
		CARDS_INC_PRODUCTS_EMP_ERR: "未添加包含的商品",
		CARDS_INC_PRODUCTS_MAX_ERR: "商品超过最多数量限制",

		CARDS_STATUS_OFF_ERR: "存在已下架的项目",

		RCARD_DISCOUNT_TYPE_ERR:  "充值卡折扣类型不对",
		RCARD_DISCOUNT_PRICE_ERR: "面值折扣类型售价必须小于面值",
		RCARD_DISCOUNT_ITEM_ERR:  "项目折扣类型售价必须等于面值",
		RCARD_DISCOUNT_ERR:       "折扣必须在1-10之间",

		//二维码
		//QRCODE_EXPIRED:"当前二维码信息已过期",
		QRCODE_EXPIRED: "二维码信息有误，请重新扫描",
	}
)

// 获取错误信息
func GetInterfaceError(code string) error {
	if val, ok := errMsg[code]; ok {
		return toolLib.CreateKcErr(code, val)
	}
	return toolLib.CreateKcErr(code)
}
