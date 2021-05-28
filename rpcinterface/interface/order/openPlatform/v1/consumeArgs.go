/*
 * @Author: your name
 * @Date: 2021-05-19 15:03:44
 * @LastEditTime: 2021-05-19 17:17:50
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \rpcOrder\rpcinterface\interface\order\openPlatform\v1\consumeArgs.go
 */
package v1

//确认消费需要的参数格式
type OPV1RcardConsumeSrvRequest struct {
	ShopId       int     //门店id
	UId          int     //用户id
	RelationId   int     //卡包id
	ConsumePrice float64 //消费金额
	ConfirmType  int     //消费确认类型 1=短信验证 2=动态码验证 3=无验证
	Captcha      string  //短语验证码或者动态码
}
