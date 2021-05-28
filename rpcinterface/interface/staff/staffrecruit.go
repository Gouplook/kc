package staff

// 定义员工招聘相关接口
// @author yinjinlin<yinjinlin_uplook@163.com>
// @date  2020/10/13 15:15

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	//---薪资报酬类型----
	//1000元/月以下
	STAFFRECRUIT_salary_type_1 = 1
	//1000-2000元/月
	STAFFRECRUIT_salary_type_2 = 2
	//2001-4000元/月
	STAFFRECRUIT_salary_type_3 = 3
	//4001-6000元/月
	STAFFRECRUIT_salary_type_4 = 4
	//6001-8000元/月
	STAFFRECRUIT_salary_type_5 = 5
	//8001-10000元/月
	STAFFRECRUIT_salary_type_6 = 6
	//10001-15000元/月
	STAFFRECRUIT_salary_type_7 = 7
	//15001-25000元/月
	STAFFRECRUIT_salary_type_8 = 8
	//25000元/月以上
	STAFFRECRUIT_salary_type_9 = 9

	//---工作经验----
	//少于1年
	STAFFRECRUIT_undergo_type_1 = 1
	//1-2年
	STAFFRECRUIT_undergo_type_2 = 2
	//3-5年
	STAFFRECRUIT_undergo_type_3 = 3
	//6-7年
	STAFFRECRUIT_undergo_type_4 = 4
	//8-10年
	STAFFRECRUIT_undergo_type_5 = 5
	//10年以上
	STAFFRECRUIT_undergo_type_6 = 6

	//---学历要求类型---
	//不限
	STAFFRECRUIT_degrees_type_1 = 1
	//小学及以上
	STAFFRECRUIT_degrees_type_2 = 2
	//初中及以上
	STAFFRECRUIT_degrees_type_3 = 3
	//高中及以上
	STAFFRECRUIT_degrees_type_4 = 4
	//大专及以上
	STAFFRECRUIT_degrees_type_5 = 5
	//本科及以上
	STAFFRECRUIT_degrees_type_6 = 6

	//---工作性质类型---
	//全职
	STAFFRECRUIT_work_type_1 = 1
	//兼职
	STAFFRECRUIT_work_type_2 = 2
	//学徒
	STAFFRECRUIT_work_type_3 = 3

	//---性别----
	//男
	STAFFRECRUIT_sex_type_1 = 1
	//女
	STAFFRECRUIT_sex_type_2 = 2
	//不限
	STAFFRECRUIT_sex_type_3 = 3

	//---职位福利---
	//五险一金
	STAFFRECRUIT_material_type_1 = 1
	//包吃
	STAFFRECRUIT_material_type_2 = 2
	//包住
	STAFFRECRUIT_material_type_3 = 3
	//每周双休
	STAFFRECRUIT_material_type_4 = 4
	//年底双薪
	STAFFRECRUIT_material_type_5 = 5
	//房补
	STAFFRECRUIT_material_type_6 = 6
	//话补
	STAFFRECRUIT_material_type_7 = 7
	//交补
	STAFFRECRUIT_material_type_8 = 8
	//饭补
	STAFFRECRUIT_material_type_9 = 9
	//加班补助
	STAFFRECRUIT_material_type_10 = 10

	//---招聘信息状态----
	//已发布
	STAFFRECRUIT_status_type_2 = 2
	//暂停
	STAFFRECRUIT_status_type_3 = 3

	//---数据库删除状态----
	//正常
	STAFFRECRUIT_deleted_no = 0
	//已删除
	STAFFRECRUIT_deleted_yes = 1
)

//枚举薪资报酬类型
func GetEnumSalaryType() map[int]interface{} {
	return map[int]interface{}{
		STAFFRECRUIT_salary_type_1: "1000元/月以下",
		STAFFRECRUIT_salary_type_2: "1000-2000元/月",
		STAFFRECRUIT_salary_type_3: "2001-4000元/月",
		STAFFRECRUIT_salary_type_4: "4001-6000元/月",
		STAFFRECRUIT_salary_type_5: "6001-8000元/月",
		STAFFRECRUIT_salary_type_6: "8001-10000元/月",
		STAFFRECRUIT_salary_type_7: "10001-15000元/月",
		STAFFRECRUIT_salary_type_8: "15001-25000元/月",
		STAFFRECRUIT_salary_type_9: "25000元/月以上",
	}
}

//枚举工作经验
func GetEnumUndergoType() map[int]interface{} {
	return map[int]interface{}{
		STAFFRECRUIT_undergo_type_1: "少于1年",
		STAFFRECRUIT_undergo_type_2: "1-2年",
		STAFFRECRUIT_undergo_type_3: "3-5年",
		STAFFRECRUIT_undergo_type_4: "6-7年",
		STAFFRECRUIT_undergo_type_5: "8-10年",
		STAFFRECRUIT_undergo_type_6: "10年以上",
	}
}

//枚举学历要求类型
func GetEnumDegreesType() map[int]interface{} {
	return map[int]interface{}{
		STAFFRECRUIT_degrees_type_1: "不限",
		STAFFRECRUIT_degrees_type_2: "小学及以上",
		STAFFRECRUIT_degrees_type_3: "初中及以上",
		STAFFRECRUIT_degrees_type_4: "高中及以上",
		STAFFRECRUIT_degrees_type_5: "大专及以上",
		STAFFRECRUIT_degrees_type_6: "本科及以上",
	}
}

//枚举工作性质类型
func GetEnumWorkType() map[int]interface{} {
	return map[int]interface{}{
		STAFFRECRUIT_work_type_1: "全职",
		STAFFRECRUIT_work_type_2: "兼职",
		STAFFRECRUIT_work_type_3: "学徒",
	}
}

//枚举性别
func GetSexType() map[int]interface{} {
	return map[int]interface{}{
		STAFFRECRUIT_sex_type_1: "男",
		STAFFRECRUIT_sex_type_2: "女",
		STAFFRECRUIT_sex_type_3: "不限",
	}
}

//枚举职位福利
func GetEnumMaterialType() map[int]interface{} {
	return map[int]interface{}{
		STAFFRECRUIT_material_type_1:  "五险一金",
		STAFFRECRUIT_material_type_2:  "包吃",
		STAFFRECRUIT_material_type_3:  "包住",
		STAFFRECRUIT_material_type_4:  "每周双休",
		STAFFRECRUIT_material_type_5:  "年底双薪",
		STAFFRECRUIT_material_type_6:  "房补",
		STAFFRECRUIT_material_type_7:  "话补",
		STAFFRECRUIT_material_type_8:  "交补",
		STAFFRECRUIT_material_type_9:  "饭补",
		STAFFRECRUIT_material_type_10: "加班补助",
	}
}

//枚举招聘信息状态
func GetEnumStatus() map[int]interface{} {
	return map[int]interface{}{
		STAFFRECRUIT_status_type_2: "发布中",
		STAFFRECRUIT_status_type_3: "暂停中",
	}
}

type ArgsEmpty struct {
}

//添加招聘信息基本信息结构
type RecruitInfoBase struct {
	ReId           int    // 招聘信息编号，自增ID
	Title          string // 招聘职位
	BusId          int    // 商家Uid
	PositionId     int    // 职位名称
	HeadCount      int    // 招聘人数
	Salary         int    // 薪资报酬
	WorkExperience int    // 工作经验
	Sex            int    // 性别
	Degrees        int    // 学历
	WorkType       int    // 工作性质
	ProvinceId     int    // 地址省id
	CityId         int    // 地址市id
	CountyId       int    // 地址县id
	Address        string // 详细地址
	MaterialId     string // 职位福利,多个id以'',''号隔开
	Phone          string // 联系人手机号
	RealName       string // 联系人姓名
	SendTime       string // 发布时间
	Status         int    // 发布状态   2 = 已发布  3 = 暂停
	Ctime          string // 创建时间
	Clicks         int64  // 阅读次数
	DeliveryNum    int64  // 投递数量
	PostFact       string // 岗位职责
	PostAsk        string // 岗位要求
}

//添加招聘信息入参数
type ArgsAddRecruitInfo struct {
	common.BsToken //企业/商户/分店信息
	RecruitInfoBase
}

//返回添加招聘信息编号
type ReplyStaffRecruit struct {
	Reid int //招聘信息编号
}

//获取招聘信息列表入参数
type ArgsGetStaffRecruitList struct {
	common.Paging
	common.BsToken //企业/商户/分店信息
}

//招聘信息列表基本信息结构体
type StaffRecruitListInfo struct {
	ReId         int    // 招聘信息编号
	PositionName string // 职位名称
	Title        string // 招聘职位
	HeadCount    int    // 招聘人数
	Salary       string // 薪资报酬
	Degrees      string // 学历
	Wtype        string // 工作性质
	Status       string // 发布状态   2 = 已发布  3 = 暂停
	SendTime     string // 发布时间
	Clicks       int64  // 阅读次数
}

//招聘信息返回列表
type ReplyGetStaffRecruitList struct {
	TotalNum int
	Lists    []StaffRecruitListInfo
}

//编辑招聘信息入参数
type ArgsEditStaffRecruit struct {
	common.BsToken //企业/商户/分店信息
	RecruitInfoBase
}

//获取招聘信息详情入参
type ArgsStaffRecruitInfo struct {
	common.BsToken     // 企业/商户/分店信息
	RedId          int // 招聘信息编号
}

// 招聘信息详情返回
type ReplyStaffRecruitInfo struct {
	ReId         int    // 招聘信息编号，自增ID
	Title        string // 招聘职位
	BusId        int    // 商家Uid
	PositionName string // 职位名称
	PositionId   int    // 职位ID
	HeadCount    int    // 招聘人数
	Salary       int    // 薪资报酬
	Undergo      int    // 工作经验
	Sex          int    // 性别
	Degrees      int    // 学历
	Wtype        int    // 工作性质
	PId          int    // 地址省id
	TId          int    // 地址市id
	CId          int    // 地址县id
	Address      string // 详细地址
	MaterialId   string // 职位福利,多个id以'',''号隔开
	Phone        string // 联系人手机号
	RealName     string // 联系人姓名
	SendTime     string // 发布时间   格式如：2020-10-10
	Status       int    // 发布状态   2 = 已发布  3 = 暂停
	Ctime        string // 创建时间   格式如：2020-10-10
	Clicks       int64  // 阅读次数
	DeliveryNum  int64  // 投递数量
	Fact         string // 岗位职责
	Ask          string // 岗位要求
	Deleted      int    // 删除数据  删除状态 0=正常 1=已删除
}

//删除招聘信息入参数
type ArgsDelStaffRecruit struct {
	common.BsToken     // 企业/商户/分店信息
	ReId           int // 招聘信息编号，自增ID
}

//批量删除招聘信息
type ArgsBatchDelStaffRecruit struct {
	common.BsToken
	ReId []int
}

//暂停/发布 招聘信息入参数
type ArgsSuspendReleaseStaffRecruit struct {
	common.BsToken
	ReId int
}

// 返回暂停/发布状态
type ReplySuspendReleaseStaffRecuit struct {
	Status int // 发布状态   2 = 已发布  3 = 暂停
}

//批量暂停/发布 招聘信息
type ArgsBatchSusReltaffRecruit struct {
	common.BsToken
	ReId []int
}

//获取招聘职位-出参
type GetRecruitPositionsBase struct {
	PositionId int    //职位id
	Name       string //职位名
}
type GetRecruitPositionsBase2 struct {
	PositionId   int    //职位id
	Name         string //职位类别
	SubPositions []GetRecruitPositionsBase
}
type ReplyGetRecruitPositions struct {
	Lists []GetRecruitPositionsBase2
}

//定义员工招聘接口
type StaffRecruit interface {
	//添加招聘信息
	AddStaffRecruit(ctx context.Context, args *ArgsAddRecruitInfo, reply *ReplyStaffRecruit) error
	//获取招聘信息列表
	GetStaffRecruitList(ctx context.Context, args *ArgsGetStaffRecruitList, reply *ReplyGetStaffRecruitList) error
	//编辑招聘信息
	EditStaffRecruit(ctx context.Context, args *ArgsEditStaffRecruit, reply *ReplyStaffRecruit) error
	//获取招聘信息详情
	GetStaffRecruitInfo(ctx context.Context, args *ArgsStaffRecruitInfo, reply *ReplyStaffRecruitInfo) error
	//删除招聘信息
	DelStaffRecruit(ctx context.Context, args *ArgsDelStaffRecruit, reply *ReplyStaffRecruit) error
	//批量删除招聘信息
	BatchDelStaffRecruit(ctx context.Context, args *ArgsBatchDelStaffRecruit, reply *bool) error
	//暂停招聘信息
	SuspendStaffRecruit(ctx context.Context, args *ArgsSuspendReleaseStaffRecruit, reply *ReplySuspendReleaseStaffRecuit) error
	//批量暂停招聘信息
	BatchSuspendStaffRecruit(ctx context.Context, args *ArgsBatchSusReltaffRecruit, reply *bool) error
	//发布招聘信息
	ReleaseStaffRecruit(ctx context.Context, args *ArgsSuspendReleaseStaffRecruit, reply *ReplySuspendReleaseStaffRecuit) error
	//批量发布招聘信息
	BatchReleaseStaffRecruit(ctx context.Context, args *ArgsBatchSusReltaffRecruit, reply *bool) error
	//获取招聘职位
	GetRecruitPositions(ctx context.Context, args *ArgsEmpty, reply *ReplyGetRecruitPositions) error
}
