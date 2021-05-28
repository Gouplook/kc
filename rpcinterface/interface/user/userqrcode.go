package user

import "git.900sui.cn/kc/rpcinterface/interface/common"

// @author liyang<654516092@qq.com>
// @date  2020/8/7 11:29

type ArgsUserQrcode struct {
	common.Utoken
}

type ReplyUserQrcode struct {
	QrcodeStr string //二维码信息
}