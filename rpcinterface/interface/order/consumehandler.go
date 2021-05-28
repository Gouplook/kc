package order

// 确认消费业务逻辑当中使用的结构体
// @author liyang<654516092@qq.com>
// @date  2020/8/10 16:17

//套餐包含的单项目
type SmSingleHandler struct {
	Id  int
	CardPackageId int
	SingleId int
	Name string
	Count int
	TransferNum int
}

//综合卡包含的单项目
type CardSingleHandler struct {
	Id  int
	CardPackageId int
	SingleId int
	Name string
}

//综合卡包含的产品
type CardGoodHandler struct {
	Id  int
	CardPackageId int
	GoodsId int
	Name string
}

//限时卡包含的单项目
type HcardSingleHandler struct {
	Id  int
	CardPackageId int
	SingleId int
	Name string
}

//限次卡包含的产品
type NcardSingleHandler struct {
	Id  int
	CardPackageId int
	SingleId int
	Name string
	Count int
	TransferNum int
}

//限时限次卡的单项目
type HncardSingleHandler struct {
	Id  int
	CardPackageId int
	SingleId int
	Name string
	Count int
	TransferNum int
}



//单项目、套餐、限时卡、限次卡、限时限次卡存储内容
type CommonConfData struct {
	SingleId int
	SingleName string
	SspId int
	SspName string
	Num int
	Price float64
	StaffId []int
}

//综合卡存储内容
type CommonCardConfData struct {
	SingleId int
	SingleName string
	SspId int
	SspName string
	Type int
	Num int
	Price float64
	StaffId []int
}



