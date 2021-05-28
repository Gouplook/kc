package insurance

/**
 * @className consume
 * @author liyang<654516092@qq.com>
 * @date 2020/9/10 14:31
 */

//获取待上传保险公司消费记录入参
type ArgsConsumeTask  struct {
	RunId int  //id
	Limit int  //一次获取数量
}

//返回待上传保险公司消费记录
type ReplyConsumeTask struct {
	Id int
	Status int
}

//保险消费数据
type InsuranceConsumeData struct {
	CardNo string `json:"cardNo"`
	CardName string `json:"cardName"`
	ConsumePrice float64 `json:"consumePrice"`
	RealConsumePrice float64 `json:"realConsumePrice"`
	Ctime int64 `json:"ctime"`
}

//单个消费记录详情
type SingleConsumeLog struct {
	RelationLogId int
	RelationId int
	CardPackageId int
	CardPackageSn string
	CardPackageType int
	JsonData InsuranceConsumeData
	Status int
}
