/******************************************
@Description:
@Time : 2020/11/17 10:06
@Author :lixiaojun

*******************************************/
package constkey

const (
	//卡项温馨提示最多10条
	CardNotesMaxNum = 10
	//卡项温馨提示单条字数最大字数<=30
	CardNotesSimpleMaxLength = 30
)

var (
	TestShareLink = "https://m.900sui.cn/index/#/serve/serveDet?singleId=%d&shopId=%d&storeItemId=%d" //测试环境分享链接
	ProdShareLink = "https://m.900sui.com/#/serve/serveDet?singleId=%d&shopId=%d&storeItemId=%d"      //生产环境分享链接
	ReguInfoDesc  = "本店已与政府监管平台信息对接，已开通资金存管/预付卡保险。消费者购卡时请查收资金存管信息或保险单。"                                       //店铺监管信息描述
)