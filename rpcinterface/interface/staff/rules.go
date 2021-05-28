package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

// 定义员工管理制度规则相关接口
// @author liyang<654516092@qq.com>
// @date  2020/4/7 11:42

//获取企业/商户考勤规则入参
type ArgsAttendanceRules struct {
	common.Utoken   //用户信息
	common.BsToken  //企业/商户信息
}

//获取企业/商户考勤规则返回信息
type ReplyAttendanceRules struct {
	Id int  //规则ID
	AssessmentItem string //考核项名称
	AssessmentStandard int //考核标准
	AssessmentType int    //考核项类型 1=迟到 2=早退 3=请假 4=旷工 5=全勤
	AssessmentMeasures int //考核项措施 1=惩罚 2=奖励
	AssessmentMeasuresName string //考核项措施名称
	RewardAmount float64   //奖罚金额
	RewardMlc  int         //奖罚MLC
}

//考核项规则更新入参
type ArgsSetRules struct {
	common.Utoken   //用户信息
	common.BsToken  //企业/商户信息
	Id  int //考核规则ID
	AssessmentStandard int //考核标准
	RewardAmount float64 //奖罚金额
	RewardMlc  int       //奖罚MLC
}

//考核项规则更新返回信息
type ReplySetRules struct {
	Id int //考核项规则ID
}

//添加行为规范入参
type ArgsAddConductRules struct {
	common.Utoken  //用户信息
	common.BsToken //企业/商户ID
	AssessmentItem string //考核项
	AssessmentContent string //考核内容
	AssessmentMeasures int //考核措施 1=惩罚 2=奖励
	RewardAmount float64   //奖罚金额
	RewardMlc  int         //奖罚MLC
	Status int  //状态 1=开启 2=禁用
}
//更新行为规范入参
type ArgsSetConductRules struct {
	common.Utoken  //用户信息
	common.BsToken //企业/商户ID
	Id  int        //行为规范ID
	AssessmentItem string    //考核项
	AssessmentContent string //考核内容
	RewardAmount float64    //奖罚金额
	RewardMlc  int          //奖罚MLC
	Status     int  //状态 1=开启 2=禁用
}

//返回行为规范信息
type ReplyConductRules struct {
	Id int //行为规范ID
}

//获取行为规范入参
type ArgsGetConductRules struct {
	common.Utoken  //用户信息
	common.BsToken //企业/商户ID
}

//返回行为规范信息
type ReplyConductRulesInfo struct{
	Id int  //规则ID
	AssessmentItem string  //考核名称
	AssessmentContent string //考核内容
	AssessmentMeasures int //考核项措施 1=惩罚 2=奖励
	AssessmentMeasuresName string //考核项措施名称
	RewardAmount float64   //奖罚金额
	RewardMlc  int         //奖罚MLC
	Status  int  //状态 1=开启 2=禁用
	StatusName string //状态名称
}

//删除行为规范入参
type ArgsDelConductRules struct {
	common.Utoken
	common.BsToken
	Id int //行为规范ID
}

//添加行为规范记录
type ArgsAddConductRulesLog struct {
	common.Utoken  //用户信息
	common.BsToken //企业/商户/分店信息
	StaffId int //员工ID
	ConductRuleId int //行为规则ID
	AssessmentTime string //考核时间 格式如：2020-03-04 12:20:01
}
//返回添加行为规范记录信息
type ReplyAddConductRulesLog struct {
	LogId int //记录ID
}

//获取行为规范统计数据入参
type ArgsGetConductRuleslogData struct {
	common.Utoken    //用户信息
	common.BsToken   //企业/商户/分店信息
	common.Paging    //分页信息
	DateYm int   //格式如201911
}
//获取行为规范统计数据返回信息
type ReplyGetConductRuleslogData struct {
	Data  []ReplyGetConductRulesData  //返回数据
	TotalNum int   //返回总数量
}

//获取行为规范统计数据返回信息Data节点
type ReplyGetConductRulesData struct {
	StaffId   int   //员工ID
	StaffName string  //员工姓名
	Rules     map[string]ReplyRulesData //map["规则ID"]map[string]interface{"ruleId":1,"Num":2,"RewardAmount":1000,"RewardMlc":2000}
}
//统计数据返回信息data节点中的data节点
type ReplyRulesData struct {
	RuleId int  //规则ID
	Num int  //规则数量
	RewardAmount float64 //金额
	RewardMlc int  //mlc
}

//获取服务表现规则入参
type ArgsGetServiceRules struct {
	common.Utoken   //用户信息
	common.BsToken  //企业/商户/分店ID
}

//返回服务表现返回信息
type ReplyGetServiceRules struct{
	 RuleId    int     //规则ID
	 RuleName  string  //规则名称
 	 RuleScore string  //规则分数
 	 AssessmentMeasures int //考核项措施 1=惩罚 2=奖励
	 AssessmentMeasuresName string //考核项措施名称
	 RewardAmount float64   //奖罚金额
	 RewardMlc  int         //奖罚MLC
}

//更新服务表现规则入参
type ArgsSetServiceRules struct{
	common.Utoken  //用户信息
	common.BsToken //企业/商户ID
	RuleId int //服务表现考核项ID
	AssessmentMeasures int //考核项措施 1=惩罚 2=奖励
	RewardAmount float64   //奖罚金额
	RewardMlc  int         //奖罚MLC
}

//返回更新服务表现规则信息
type ReplySetServiceRules struct {
	 RuleId  int  //规则ID
}

//执行服务表现统计入参
type ArgsServiceComment struct{
	StaffId int //员工ID
	StartNum  float64 //评价数量
}
//返回执行服务表现统计信息
type ReplyServiceComment struct {
	StaffId  int //员工ID
}


//获取服务表现统计数据入参
type ArgsGetServiceRuleslogData struct {
	common.Utoken    //用户信息
	common.BsToken   //企业/商户/分店信息
	common.Paging    //分页信息
	DateYm int   //格式如201911
}

//获取服务表现统计数据返回信息
type ReplyGetServiceRuleslogData struct {
	Data  []ReplyGetServiceRulesData  //返回数据
	TotalNum int   //返回总数量
}

//获取行服务表现统计数据返回信息Data节点
type ReplyGetServiceRulesData struct {
	StaffId   int   //员工ID
	StaffName string  //员工姓名
	Rules     map[string]ReplyServiceRulesData //map["规则ID"]map[string]interface{"ruleId":1,"Num":2,"RewardAmount":1000,"RewardMlc":2000}
}

//统计数据返回信息data节点中的data节点
type ReplyServiceRulesData struct {
	RuleId int  //规则ID
	Num int  //规则数量
	RewardAmount float64 //金额
	RewardMlc int  //mlc
}

type Rules interface {
	//获取考勤考核项
	/*GetAttendanceRules(ctx context.Context,args *ArgsAttendanceRules,reply *[]ReplyAttendanceRules) error
	//更新考核项
	SetAttendanceRules(ctx context.Context,args *ArgsSetRules,reply *ReplySetRules) error*/
	//新增行为规范考核项
	AddConductRules(ctx context.Context,args *ArgsAddConductRules,reply *ReplyConductRules) error
	//更新行为规范考核项
	SetConductRules(ctx context.Context,args *ArgsSetConductRules,reply *ReplyConductRules) error
	//获取行为规范考核项
	GetConductRules(ctx context.Context,args *ArgsGetConductRules,reply *[]ReplyConductRulesInfo) error
	//删除行为规范考核项
	DeleteConductRules(ctx context.Context,args *ArgsDelConductRules,reply *ReplyConductRules) error
	//新增行为规范记录
	AddConductRulesLog (ctx context.Context,args *ArgsAddConductRulesLog,reply *ReplyAddConductRulesLog) error
	//获取行为规范记录统计-企业/商户
	GetConductRulesLogDatasForBus(ctx context.Context,args *ArgsGetConductRuleslogData,reply *ReplyGetConductRuleslogData) error
	//获取行为规范记录统计-分店
	GetConductRulesLogDatasForShop(ctx context.Context,args *ArgsGetConductRuleslogData,reply *ReplyGetConductRuleslogData) error
	//获取服务表现考核项
	GetServiceRules(ctx context.Context,args *ArgsGetServiceRules,reply *[]ReplyGetServiceRules) error
	//设置服务表现考核项
	SetServiceRules(ctx context.Context,args *ArgsSetServiceRules,reply *ReplySetServiceRules) error
	//执行服务表现统计
	ServiceComment(ctx context.Context,args *ArgsServiceComment,reply *ReplyServiceComment) error
	//获取服务表现记录统计-企业/商户
	GetServiceRulesLogDatasForBus(ctx context.Context,args *ArgsGetServiceRuleslogData,reply *ReplyGetServiceRuleslogData) error
	//获取服务表现记录统计-分店
	GetServiceRulesLogDatasForShop(ctx context.Context,args *ArgsGetServiceRuleslogData,reply *ReplyGetServiceRuleslogData) error
}

