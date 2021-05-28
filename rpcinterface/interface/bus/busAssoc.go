package bus

/**
 * @className busAssoc
 * @author liyang<654516092@qq.com>
 * @date 2020/9/4 13:50
 */
const (
	//#协会会员类型
	//会员单位
	BUSASSOCJOIN_assoc_type_member = 1
	//理事单位
	BUSASSOCJOIN_assoc_type_director = 2
	//副会长单位
	BUSASSOCJOIN_assoc_type_vicepresident = 3
)

type ReplyBusAssocInfo struct {
	//卡协会员
	IsAssoc int     //是否为协会会员 0=否 1=是
	ExpireTime int64
	ExpireTimeStr string
}

