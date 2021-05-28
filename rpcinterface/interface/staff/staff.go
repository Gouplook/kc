package staff

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//定义员工相关接口
// @author liyang<654516092@qq.com>
// @date  2020/4/1 9:09

//添加员工入参
type ArgsAddStaff struct {
	common.Utoken              //用户信息
	common.BsToken             //企业/商户/分店ID
	Name           string      //员工姓名
	NickName       string      //员工昵称
	Sex            int         //性别 1=女 2=男
	Phone          string      //员工手机号 -自动注册为九百岁账号
	PostId         int         //员工所属岗位
	IsCraftMan     int         //员工是否为手艺人 1=否 2=是
	BirthDay       string      //员工生日 格式如：2019-02-02
	EmergencyPhone string      //员工紧急联系电话
	CardNo         string      //员工身份证号
	EntryDate      string      //入职日期 格式如：2019-01-20
	BasicSalary    float64     //员工底薪工资
	Guaranteed     float64     //员工保底工资
	Allowance      float64     //员工福利津贴/月
	Education      int         //员工学历 0=无学历 1=小学 2=初中 3=高中 4=大专 5=本科及以上
	ImgHash        string      //员工头像hash
	Detail         StaffDetail //员工详情结构体
}

//修改员工入参
type ArgsEditStaff struct {
	common.Utoken              //用户信息
	common.BsToken             //企业/商户/分店ID
	StaffId        int         //员工ID
	Name           string      //员工姓名
	NickName       string      //员工昵称
	Sex            int         //性别 1=女 2=男
	Phone          string      //员工手机号
	PostId         int         //员工所属岗位
	IsCraftMan     int         //员工是否为手艺人 1=否 2=是
	WorkStatus     int         //在职状态 1=在职 2=离职
	BirthDay       string      //员工生日 格式如：2019-02-02
	EmergencyPhone string      //员工紧急联系电话
	CardNo         string      //员工身份证号
	DepartureDate  string      //离职日期 格式如：2019-01-20
	EntryDate      string      //入职日期 格式如：2019-01-20
	BasicSalary    float64     //员工底薪工资
	Guaranteed     float64     //员工保底工资
	Allowance      float64     //员工福利津贴/月
	Education      int         //员工学历 0=无学历 1=小学 2=初中 3=高中 4=大专 5=本科及以上
	ImgHash        string      //员工头像hash
	Detail         StaffDetail //员工详情结构体
}

//员工删除入参
type ArgsDelStaff struct {
	common.Utoken      //用户信息
	common.BsToken     //企业/商户/分店ID
	StaffId        int //员工ID
}

//员工详情入参
type StaffDetail struct {
	HealthCode        string //健康证件码
	WorkMomentImgHash string //工作瞬间，多个以","号隔开
	CertImgHash       string //相关证件，多个以","号隔开
	Introduction      string //个人简介
	Address           string //详细地址
}

//返回添加员工信息
type ReplyStaff struct {
	StaffId int //员工ID
}

//获取员工详情入参
type ArgsStaffInfo struct {
	common.Utoken      //用户信息
	common.BsToken     //企业/商户/分店ID
	StaffId        int //员工ID
}

//获取员工详情返回信息
type ReplyStaffInfo struct {
	StaffId        int     //员工ID
	Name           string  //员工姓名
	NickName       string  //员工昵称
	Sex            int     //性别 0=女 1=男
	Phone          string  //员工手机号
	PostId         int     //员工所属岗位
	PostName       string  //岗位
	IsCraftMan     int     //员工是否为手艺人 1=否 2=是
	WorkStatus     int     //在职状态 1=在职 2=离职
	BirthDay       string  //员工生日 格式如：2019-02-02
	EmergencyPhone string  //员工紧急联系电话
	BusId          int     //员工所属企业/商户ID
	ShopId         int     //员工所属分店ID
	CardNo         string  //员工身份证号
	Uid            int     //员工系统账号UID
	ImgId          int     //员工头像ID
	DepartureDate  string  //离职日期 格式如：2019-01-20
	EntryDate      string  //入职日期 格式如：2019-01-20
	BasicSalary    float64 //员工底薪工资
	Guaranteed     float64 //员工保底工资
	Allowance      float64 //员工福利津贴/月
	Education      int     //员工学历 0=无学历 1=小学 2=初中 3=高中 4=大专 5=本科及以上
	Deleted        int     //是否删除 0=正常 1=已删除
	Detail         Detail  //员工详情结构体
}

//员工详情
type Detail struct {
	HealthCode    string //健康证件码
	WorkMomentImg string //工作瞬间，多个以","号隔开
	CertImg       string //相关证件，多个以","号隔开
	Address       string //详细地址
	Introduction  string //员工详情
}

//获取员工列表-企业/商户/分店
type ArgsGetStaffList struct {
	common.Paging         //分页信息
	common.Utoken         //用户信息
	common.BsToken        //企业/商户ID
	Keyword        string //关键字
	IsCraftMan int        //是否为手艺人
}

//获取员工列表-前端
type ArgsStaffList struct{
	common.Paging   //分页信息
	ShopId int
	IsCraftMan int        //是否为手艺人
}

//返回员工列表信息
type ReplyGetStaffList struct {
	StaffId    int    //员工ID
	IsCraftMan int //员工是否为手艺人 1=否 2=是
	Name       string //员工姓名
	NickName   string //员工昵称
	SexName    string //女/男
	PostName   string //岗位名称
	Phone      string //手机号
	BusId      int    //员工所属企业/商户ID
	ShopId     int    //员工所属分店ID
	ShopName   string //分店名称
	BranchName string //分店名
	EntryDate  string //员工入职日期 字符串
	ImgId      int //员工头像ID
}

//返回公共总数量
type ReplyStaffList struct {
	Lists     []ReplyGetStaffList //数据
	TotalNum int                 //总数量
}
// 简单的员工信息
type SimpleStaff struct {
	StaffId        int     //员工ID
	Name           string  //员工姓名
	Score   int     // 技师评分
	NickName       string  //员工昵称
	ImgId          int     //员工头像ID
	ImgUrl string
}

//===ES搜索===========/
//搜索入参
type ArgsSearchWhere struct {
	common.Paging        //分页
	BusId         int    //企业/商户ID
	ShopId        int    //分店ID
	Keywords      string //搜索关键词 姓名/昵称/手机号
	PostId        int    //岗位ID
	IsCraftMan    int    //员工是否为手艺人
}

//搜索返回信息
type ReplySearch struct {
	Data     []ReplySearchData //搜索员工数据
	TotalNum int               //员工数量
}

//搜索返回信息-Data节点
type ReplySearchData struct {
	StaffId  int    //员工ID
	BusId    int    //企业/商户ID
	ShopId   int    //分店ID
	PostId   int    //岗位ID
	NickName string //昵称
	Name     string //姓名
	Phone    string //联系电话
}

// 查找员工信息内部rpc
type ArgsGetListByStaffIds struct {
	ShopId   int
	StaffIds []int
}

//查找员工信息内部rpc
type ReplyGetListByStaffIds struct {
	StaffId int
	Name    string
}

//查找员工信息内部rpc
type ReplyGetListByStaffIds2 struct {
	StaffId int
	Name    string
	NickeName string `mapstructure:"nick_name"`
}

// 员工列表-无需验证-分店
type ArgsGetStaffListByShopId struct {
	common.Paging
	ShopId int
	IsCraftMan int
}
type ReplyGetStaffListByShopIdBase struct {
	StaffId    int    //员工ID
	Name       string //员工姓名
	NickName   string //员工昵称
	SexName    string //女/男
	PostName   string //岗位名称
	ImgId int
	ImgUrl string
	AvaStar float64 // 员工评分

}
//员工列表-无需验证-分店
type ReplyGetStaffListByShopId struct {
	Lists []ReplyGetStaffListByShopIdBase
	TotalNum int
}
type ReplyGetStaffInfoById struct {
	ReplyStaffInfo
	Rank 		map[string]RankInfo//员工评分详情
}

type ReplyStaffNameAndPost struct {
	StaffId int
	Name  string
	NickName string
	PostId int
	PostName string
	Phone string
}
//获取当月员工新增或离职率-入参
type ArgsGetStaffAddMinus struct {
	StaffId int
}
//获取当月员工新增或离职率-出参
type ReplyGetStaffAddMinus struct {
	TotalWorkStaffNum int//总的在职员工数
	CurMonthAddStaffNum int //本月新增员工数
	CurMonthStaffAddRate float64 //本月员工新增率
	CurMonthMinusStaffNum int //本月离职数
	CurMonthStaffMinusRate float64 //本月员工离职率
	WorkStatus int//1-在职 2-离职
	Ctime int64 //新增员工创建时间
	DepartureDate int64 //离职日期
}

const (
	//员工状态
	STAFF_INSERT = 1
	STAFF_QUIT = 2
	STAFF_DELETE = 3
	//#员工是否为手艺人
	//否
	STAFF_craft_man_no  = 1
	//是
	STAFF_craft_man_yes = 2
)
//根据id获取员工状态（新增，离职，删除）
type ArgsGetStaffStatusById struct {
	StaffId int
}
type ReplyGetStaffStatusById struct {
	BusId int  //员工所属企业/商户ID
	Status int //1-新增 2-离职 3-删除
}
//  根据busId 获取在在职员工总数量
type ArgsGetStaffByBusIdNum struct {
	BusId      int    // 员工所属企业/商户ID
}
type ReplyGetStaffByBusIdNum struct {
	StaffNum int // 店铺员工在职总数量
}

//定义员工接口
type Staff interface {
	//添加员工
	AddStaff(ctx context.Context, args *ArgsAddStaff, reply *ReplyStaff) error
	//编辑员工
	EditStaff(ctx context.Context, args *ArgsEditStaff, reply *ReplyStaff) error
	//删除员工-逻辑删除
	DeleteStaff(ctx context.Context, args *ArgsDelStaff, reply *ReplyStaff) error
	//获取员工详情-企业/商户
	GetStaffInfoByBus(ctx context.Context, args *ArgsStaffInfo, reply *ReplyStaffInfo) error
	//获取员工详情-分店
	GetStaffInfoByShop(ctx context.Context, args *ArgsStaffInfo, reply *ReplyStaffInfo) error
	//获取员工列表-企业/商户
	GetStaffByBusId(ctx context.Context, args *ArgsGetStaffList, reply *ReplyStaffList) error
	//获取员工列表-分店
	GetStaffByShopId(ctx context.Context, args *ArgsGetStaffList, reply *ReplyStaffList) error
	//获取员工详情-rpc内部调用
	GetStaffDetail(ctx context.Context, args *ArgsStaffInfo, reply *ReplyStaffInfo) error
	//获取员工列表-rpc内部调用
	GetListByStaffIds(ctx context.Context, args *ArgsGetListByStaffIds, reply *[]ReplyGetListByStaffIds) error
	//获取员工列表-rpc内部调用
	GetListByStaffIds2(ctx context.Context, staffIds *[]int, reply *[]ReplyGetListByStaffIds2) error

	//员工列表-无需验证-分店
	GetStaffListByShopId(ctx context.Context,args *ArgsGetStaffListByShopId,reply *ReplyGetStaffListByShopId)error
	//员工详情-无需验证-分店
	GetStaffInfoById(ctx context.Context,args *int,reply *ReplyGetStaffInfoById)error
	//根据多个员工id 查询员工名称和对应岗位信息 rpc 调用
	GetStaffNameAndPostByIds(ctx context.Context,args *[]int,reply *[]ReplyStaffNameAndPost)error
	//获取当月员工新增或离职率-rpc
	GetStaffAddMinusByBusIdRpc(ctx context.Context,args *ArgsGetStaffAddMinus,reply *ReplyGetStaffAddMinus)error
	//根据id获取员工状态（新增，离职，删除）
	GetStaffStatusById(ctx context.Context,args *ArgsGetStaffStatusById,reply *ReplyGetStaffStatusById)error
	// 根据busId 获取在在职员工总数量
	GetStaffByBusIdNum(ctx context.Context, args *ArgsGetStaffByBusIdNum, reply *ReplyGetStaffByBusIdNum) error
}
