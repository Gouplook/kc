package order

import "git.900sui.cn/kc/rpcinterface/interface/common"

/*
	-------------选择卡项目下的权益及充值项目结构体--------------------
*/

//入参
type ArgsGetChooseMarket struct {
	common.BsToken
	CustomerId   int
	SingleItemId int
}

//返回
type ResponseChooseMarketList struct {
	BusId      int          //商户id
	CustomerId int          //客户id
	SingleId   int          //单项目id，根据前端要求
	EquityList []EquityList //权益包
	ICardList  []ICardList  //身份卡包
	SmList     []SmList     //套餐卡包
	SingleList []SingleList //单项目卡包
	NCardList  []NCardList  //限次卡包
	HCardList  []HCardList  //限次卡包
	HNcardList []HNcardList //限次卡包
	RCardList  []RCardList  //充值卡包
	CardList   []CardList   //综合卡包
}

type HNcardList struct {
	Id             int     //主键id
	CardId         int     //卡包id
	CardName       string  //卡包名称
	RelationId     int     //关系id
	BuyShopId      int     //购买店铺id
	SingleItemId   int     //单项目id
	SingleItemName string  //单项目名称
	DiscountRate   float64 //折扣率
	Count          int     //数量
	ExpiredTime    string  //过期时间
	RemainingDay   int     //剩余天数
	Status         int     //状态
}

type HCardList struct {
	Id             int    //主键id
	CardId         int    //卡包id
	CardName       string //卡包名称
	RelationId     int    //关系id
	BuyShopId      int    //购买店铺id
	SingleItemId   int    //单项目id
	SingleItemName string //单项目名称
	ExpiredTime    string //过期时间
	RemainingDay   int    //剩余天数
	Status         int    //状态
}

type NCardList struct {
	Id             int    //主键id
	CardId         int    //卡包id
	CardName       string //卡包名称
	RelationId     int    //关系id
	BuyShopId      int    //购买店铺id
	SingleItemId   int    //单项目id
	SingleItemName string //单项目名称
	Count          int    //剩余数量
	ExpiredTime    string //过期时间
	RemainingDay   int    //剩余天数
	Status         int    //状态
}

type SingleList struct {
	Id             int     //主键id
	BuyShopId      int     //购买店铺id
	RelationId     int     //关系id
	SingleItemId   int     //单项目id
	SingleItemName string  //单项目名称
	Price          float64 //市场价
	RealPrice      float64 //实际价格
	Status         int     //状态
}

type SmList struct {
	Id             int    //主键id
	CardId         int    //卡包id
	CardName       string //卡包名称
	RelationId     int    //关系id
	BuyShopId      int    //购买店铺id
	SingleItemId   int    //单项目id
	SingleItemName string //单项目名称
	Count          int    //剩余数量
	ExpiredTime    string //过期时间
	RemainingDay   int    //剩余天数
	Status         int    //状态
}

type ICardList struct {
	Id             int     //主键id
	CardId         int     //卡包id
	CardName       string  //卡包名称
	RelationId     int     //关系id
	BuyShopId      int     //购买店铺id
	SingleItemId   int     //单项目id
	SingleItemName string  //单项目名称
	DiscountRate   float64 //折扣率
	ExpiredTime    string  //过期时间
	RemainingDay   int     //剩余天数
	Status         int     //状态
}

//权益-身份
type EquityList struct {
	Id             int    //主键id
	CardId         int    //卡包id
	CardName       string //卡包名称
	RelationId     int    //关系id
	BuyShopId      int    //购买店铺id
	SingleItemId   int    //单项目id
	SingleItemName string //单项目名称
	Count          int    //剩余次数
	ExpiredTime    string //过期时间
	RemainingDay   int    //剩余天数
	Status         int    //状态
}

//充值卡
type RCardList struct {
	Id           int     //卡id
	RelationId   int     //关系id
	CardName     string  //卡名称
	BuyShopId    int     //购买店铺id
	Balance      float64 //余额
	DiscountRate float64 //折扣率
	ExpiredTime  string  //过期时间
	RemainingDay int     //剩余天数
	Status       int     //状态
}
type CardList struct {
	Id           int     //卡id
	RelationId   int     //关系id
	CardName     string  //卡名称
	BuyShopId    int     //购买店铺id
	Balance      float64 //余额
	DiscountRate float64 //折扣率
	ExpiredTime  string  //过期时间
	RemainingDay int     //剩余天数
	Status       int     //状态
}
