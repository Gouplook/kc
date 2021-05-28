package order

type CardIcardRechargeLog struct {
	Id           int
	Price        float64
	ConsumePrice float64
}

//获取卡包详情-RPC
type ArgsGetRpcIcard struct {
	Id int  //卡包ID
}

//卡包详情
type CardIcardInfo struct {
	Id                 int     //卡包ID
	CardSn             string  //编号
	BusId              int     //企业/商户ID
	ShopId             int     //分店ID
	Uid                int     //下单用户ID
	RealPrice          float64 //实际金额
	Price              float64 //面值金额
	ConsumePrice       float64 //消费面值金额
	ActualConsumePrice float64 //实际消费金额
	ServicePeriod      int     //保险周期
	DisaccountType     int     //折扣方式
	Disaccount         float64 //折扣率
	CardId             int     //卡ID
	CardName           string  //卡名称
	Status             int     //状态
	PayTime            int64   //付款时间
	FirstConsumeTime   int64   //第一次消费时间
	ConsumeingTime     int64   //最近一次消费时间
	ConsumeCompTime    int64   //消费完成时间
	PayChannel         int     //支付渠道
	InsuranceChannel   int     //保险渠道
	Deleted            int     //是否正常显示
	Ctime              int     //生成时间
}

//用户卡包详情
type CardIcardUserInfo struct {
	RealPrice          float64 //实际金额
	Price              float64 //面值金额
	ConsumePrice       float64 //消费面值金额
	ActualConsumePrice float64 //实际消费金额
	DisaccountType     int     //1=购卡面值打折 2=项目打折
	Disaccount         float64 //折扣率
	InsuranceChannel   int     //保险渠道
	CardName           string  //卡名称
	CardId             int     //卡ID
	ImgId              int     //卡封面
	PayTime            int64   //付款日期时间戳
	PayTimeStr         string  //付款日期
	ServicePeriod      int     //周期
	Discount           float64 //身份卡折扣
	ExpireDate         int64   //过期日期期时间戳
	ExpireDateStr      string  // 过期日期
	ExpireSurDay       int     // 剩余天数
	DiscountSingle     float64 // 单项目折扣
	DiscountGoods      float64 // 商品折扣
}

//用户卡包中包含的单项目
type CardIcardUserSingle struct {
	SingleId   int    //单项目ID
	SingleName string `mapstructure:"name"` //单项目名称
	ImgId      int
	Price      string
	canUse     int //0=未知 1=不可用 2=可用
}
