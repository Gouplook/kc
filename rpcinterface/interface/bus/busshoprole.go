package bus

import "git.900sui.cn/kc/rpcinterface/interface/common"

/**
 * @className busshoprole
 * @author liyang<654516092@qq.com>
 * @date 2020/9/2 10:42
 */


//获取分店角色入参
type ArgsShopRole struct {
	common.Utoken
	common.BsToken
}

//返回分店角色信息
type ReplyShopRole struct {
	RoleId     int
	RoleName   string
	RoleDesc   string
	IsReserved int //是否为固化字段 0=普通 1=固化角色(系统角色)
	Ctime int64
	CtimeStr string
}

//获取分店角色用户入参
type ArgsShopRoleUser struct {
	common.Utoken
	common.BsToken
	common.Paging
}

//返回分店角色用户信息
type ReplyShopRoleUser struct {
	Lists []ReplyShopRoleUserInfo
	TotalNum int
}

//返回分店角色用户信息详情
type ReplyShopRoleUserInfo struct {
	Id int //用户ID
	StaffId int //员工ID
	Name string //员工名称
	NickName string //员工昵称
	RoleId int  //角色ID
	RoleName string //角色名称
	Ctime int64 //创建时间
	CtimeStr string //创建日期字符串
}

//分店分配用户权限入参
type ArgsAddShopUser struct {
	common.Utoken
	common.BsToken
	StaffId int //员工ID
	RoleId int  //员工角色ID
}

//分店修改用户权限入参
type ArgsEditShopUser struct {
	common.Utoken
	common.BsToken
	Id    int   //用户ID
	StaffId int //员工ID
	RoleId int  //员工角色ID
}

//分店删除用户权限
type ArgsDelShopUser struct {
	common.Utoken
	common.BsToken
	Id    int   //用户ID
}

//分店分配用户权限返回信息
type ReplyShopUser struct {
	Id int //用户ID
}

//rpc查询入参
type ArgsCheckShopUser struct {
	ShopId  int
	StaffId int
}

//rpc查询返回参数
type ReplyCheckShopUser struct {
	Id int      //用户ID
	ShopId  int //分店ID
	BusId   int //企业/商户ID
	StaffId int //员工ID
	RoleId  int  //角色ID
	Ctime   int64 //创建时间
}


const (
	//#角色类型
	//普通角色
	BUSSHOPROLE_reserved_normal = 0
	//系统角色[固化角色]
	BUSSHOPROLE_reserved_system = 1
)




