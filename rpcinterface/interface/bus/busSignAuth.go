package bus

/**
 * @className 企业/商户绑定账户
 * @author liyang<654516092@qq.com>
 * @date 2020/11/9 15:04
 */

//绑定账户入参
type ArgsSupplyAccount struct {
	Phone string //手机号
	Sign  string //签名字符串
}

//验证账户入参
type ArgsAuthAccount struct {
	Phone string //手机号
}

//返回
type ReplyAccount struct {
	Phone string
}

//sass绑定账户入参
type ArgsSassSupplyAccount struct {
	CompanyName string //企业名称
	SignCode    string //信息对接识别码
	Phone       string //手机号
	Captcha     string //短信验证码
}

//sass查询绑定账户入参
type ArgsQuerySassAccount struct {
	CompanyName string //企业名称
	SignCode    string //信息对接识别码
}

//sass查询绑定账户出参
type ReplyQuerySassAccount struct {
	Phone string //手机号
}
