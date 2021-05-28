package order

import "git.900sui.cn/kc/rpcinterface/interface/common"

/**
 * @className cardPackageQrcode
 * @author liyang<654516092@qq.com>
 * @date 2020/9/18 13:39
 */



//卡包二维码监测入参
type ArgsCardPackageQrcodeCheck struct {
	common.Utoken //用户信息
	common.BsToken //企业/商户信息
	RelationId int //卡包关联ID
}

//卡包二维码返回信息
type ReplyCardPackageQrcodeCheck struct {
	RelationId int
}

//卡包消费码查询结果入参
type ArgsCardPackageQrcode struct {
	common.Utoken //用户信息
	common.BsToken //企业/商户信息
	ConsumeCode string
}

//卡包消费码查询返回结果
type ReplyCardPackageQrcode struct {
	Uid int
	RelationId int
	CardPackageType int
}