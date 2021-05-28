package bus

// 商户会员
import (
	"context"
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"git.900sui.cn/kc/rpcinterface/interface/user"
	"strings"
)

const (
	Source_online  = 1 // 线上会员
	Source_offline = 2 // 线下会员

	// 性别
	Gender_unknow = 0 // 未知
	Gender_male   = 1 // 男
	Gender_female = 2 // 女

	HideTypeText  = 1 // 文本
	HideTypeImage = 2 // 图片
	HideTypeDate  = 3 // 日期
)

// 来源数组
func GetSources() []int {
	return []int{
		Source_online,
		Source_offline,
	}
}

// 性别数组
func GetGenders() []int {
	return []int{
		Gender_unknow,
		Gender_male,
		Gender_female,
	}
}

// 隐私数据类型
func GetHideTypes() []int {
	return []int{
		HideTypeText,
		HideTypeImage,
		HideTypeDate,
	}
}

type Source = int // 来源 GetSources()

// 验证来源
func VerfiySource(source Source) error {
	if functions.InArray(source, GetSources()) {
		return nil
	}
	return common.GetInterfaceError(common.MEMBER_SOURCE_ERROR)
}

// 验证性别
func VerfiyGender(gender int) error {
	if functions.InArray(gender, GetGenders()) {
		return nil
	}
	return common.GetInterfaceError(common.MEMBER_GENDER_ERROR)
}

// 验证隐私数据类型
func VerfiyHideType(hideType []HideData) error {
	for _, v := range hideType {
		if !functions.InArray(v.Type, GetHideTypes()) {
			return common.GetInterfaceError(common.MEMBER_DATA_TYPE_ERROR)
		}
	}
	return nil
}

// 验证姓名
func VerfiyName(name string) error {
	if strings.TrimSpace(name) == "" {
		return common.GetInterfaceError(common.MEMBER_NAME_IS_NIL)
	}
	return nil
}

type ArgsAddTag struct {
	common.BsToken

	Name string // 标签名称
}

type ArgsRemoveTag struct {
	common.BsToken

	TagId int // 标签id
}

type ReplyTag struct {
	Id   int    // 标签id
	Name string // 标签名称
}

type ArgsAddMember struct {
	Phone  string
	Name   string
	Gender int
}

type HideData struct {
	Key   string // 标签名
	Value string // 存储数据
	Type  int    // 数据类型
}

type ReplyMemberDetail struct {
	MemberId int    // 会员id
	BusId    int    // 店铺/商户id
	RiskBusId int //风控系统busId
	ShopId   int    // 门店id
	ShopName string // 门店名称
	Uid      int    // 900岁会员id

	Name          string // 姓名
	Phone         string // 手机号
	Gender        int    // 性别
	CardNum       string // 身份证
	CreateTime    int    // 成为会员时间
	CreateTimeStr string // 成为会员时间
	Birth         int    // 生日 YYYY-mm-dd
	Source
	Level              int    // 会员等级
	LevelName          string // 会员等级
	ConsumeAmount      string // 累计消费金额
	ConsumeCount       int    // 累计消费次数
	LastConsumeTime    int    // 最后消费时间
	LastConsumeTimeStr string // 最后消费时间
	AllCard            int    // 累计卡数量
	UseCard            int    // 当前可用卡数量
	AllRights          int    // 累计权益
	UseRights          int    // 当前可用权益
	AllCoupon          int    // 所有优惠卷
	UseCoupon          int    // 当前可用优惠卷
}

type ArgsMemberBase struct {
	common.BsToken
	user.Channel
	user.Device

	MemberId int        // 店铺会员id
	Name     string     // 姓名
	Gender   int        // 性别
	CardNum  string     // 身份证
	Birth    string     // 生日 YYYYmmdd
	Pid      int        // 省id
	Cid      int        // 市id
	Did      int        // 县区id
	Address  string     // 详细地址
	Remark   string     // 备注
	HideData []HideData // 隐藏数据
	Tag      []int      // 会员标签
}

type ReplyMemberBase struct {
	MemberId int // 会员id
	BusId    int // 店铺/商户id
	ShopId   int // 门店id
	Uid      int // 900岁会员id

	Name          string // 姓名
	Phone         string // 手机号
	Gender        int    // 性别
	CardNum       string // 身份证
	CreateTime    int    // 成为会员时间
	CreateTimeStr string // 成为会员时间
	Birth         int    // 生日 YYYY-mm-dd
	Source
	Level    int        // 会员等级
	Pid      int        // 省id
	Cid      int        // 市id
	Did      int        // 县区id
	Address  string     // 详细地址
	Remark   string     // 备注
	HideData []HideData // 隐藏数据
	Tag      []Tag      // 会员标签
}

type Tag struct {
	Id   int    // 标签id
	Name string // 标签名称
}

type ReplyMemberInfo struct {
	Id     int // 会员id
	BusId  int // 店铺/商户id
	ShopId int // 门店id
	Uid    int // 900岁会员id

	Name       string // 姓名
	Phone      string // 手机号
	Gender     int    // 性别
	CardNum    string // 身份证
	CreateTime int    // 身份证
	Birth      int    // 生日 YYYY-mm-dd
	Source
	Level    int        // 会员等级
	Pid      int        // 省id
	Cid      int        // 市id
	Did      int        // 县区id
	Address  string     // 详细地址
	Remark   string     // 备注
	HideData []HideData // 隐藏数据
	Tag      []int      // 会员标签
}

type ArgsMemberInfo struct {
	common.BsToken
	user.Channel
	user.Device

	Uid     int    //平台会员id
	Name    string // 姓名
	Phone   string // 手机号
	Gender  int    // 性别
	CardNum string // 身份证
	Birth   string // 生日 YYYY-mm-dd
	Source
	Level    int        // 会员等级
	Pid      int        // 省id
	Cid      int        // 市id
	Did      int        // 县区id
	Address  string     // 详细地址
	Remark   string     // 备注
	HideData []HideData // 隐藏数据
	Tag      []int      // 会员标签
}

type ArgsMemberParam struct {
	common.BsToken

	common.Paging // 分页

	Keyword      string // 检索关键字 姓名/手机号
	Phone        string // 检索手机号
	ConsumeMonth int    // 消费频次 月数
	ConsumeCount int    // 消费次数
	ShopId       int    // 门店id
	Level        int    // 会员等级
	Tag          int    // 会员标签
	BirthStart   int    // 生日开始时间 例: 1月1日  101
	BirthEnd     int    // 生日截止时间 例: 12月31日  1231
	Source       Source // 来源
}

type MemberItem struct {
	MemberId int    // 会员id
	Name     string // 姓名
	Phone    string // 手机号
	Gender   int    // 性别
	Birth    int    // 生日 YYYYmmdd
	Uid      int    // 会员平台id
	Source
	Level           int    // 会员等级
	ShopId          int    // 所属门店
	ConsumeAmount   string // 累计消费金额
	ConsumeCount    string // 累计消费次数
	LastConsumeTime string // 最后消费时间
	AllCard         string // 总持卡数量
	CreateTime      string // 会员创建时间
}

type MemberList struct {
	Lists []MemberItem // 会员列表

	TotalNum int // 总条数
}

type ArgsMemberRebate struct {
	Busid int //商家id
	Uid   int // 用户uid
}

type ArgsUserToMember struct {
	common.BsToken

	user.ReplyUserinfo // 会员平台信息
}

//根据手机号获取店铺会员用户
type ArgsGetUserInfoByPhone struct {
	Phone  string
	ShopId int
	BusId  int
}
type ReplyGetUserInfoByPhone struct {
	Uid  int
	Name string
}

type ArgsGetByUid struct {
	common.BsToken
	Uid int //用户ID
}

//发送会员短信入参
type ArgsSendMemberSms struct {
	common.BsToken
	Uids       string
	SmsMessage string
}

//添加店铺会员表数据统计入参
type ArgsAddMemberConsumeCount struct {
	MemberConsumeCountBase
}
type MemberConsumeCountBase struct {
	MemberId        int     // 会员ID
	ConsumeAmount   float64 // 累计消费金额
	ConsumeCount    int     // 累计消费次数
	LastConsumeTime int64   // 最后消费时间
	AllCard         int     // 总持卡数量
}
//添加店铺会员表数据统计出参
type ReplyAddMemberConsumeCount struct {
	MemberId int
}
//获取店铺会员消费数据出参
type ReplyGetMemberConsumeCount struct {
	MemberConsumeCountBase
}
// 门店总数 入参数
type ArgsRiskMemberNum struct {
	BusId int // 商铺ID
}

// 获取风控系统商铺ID 门店总数返回值
type ReplyRiskMemberNum struct {
	RiskBusId int   // 风控系统商铺的ID
	UserNum int     // 门店总数量
}

//根据指定店铺的多个uid查询会员信息 入参
type ArgsGetMemberInfo struct {
	Uids []int
	BusId int
}

type ReplyGetMemberInfo struct {
	Uid int
	Name string
	Phone string
}

//根据手机号模糊匹配门店会员-入参
type ArgsGetUserInfoByPhoneMatch struct {
	MatchPhone string //手机号开头
	BusId int
}
//根据手机号模糊匹配门店会员-出参
type ReplyGetUserInfoByPhoneMatch struct {
	Lists []ReplyGetMemberInfo
}

type Member interface {
	// 添加会员标签
	AddTag(ctx context.Context, args *ArgsAddTag, reply *bool) error
	// 移除会员标签
	RemoveTag(ctx context.Context, args *ArgsRemoveTag, reply *bool) error
	// 获取全部会员标签
	GetTagList(ctx context.Context, args *common.BsToken, reply *[]ReplyTag) error
	// 添加会员
	AddMember(ctx context.Context, args *ArgsMemberInfo, reply *bool) error
	// 获取会员列表
	GetMemberList(ctx context.Context, args *ArgsMemberParam, reply *MemberList) error
	// 获取会员信息接口
	GetMemberInfo(ctx context.Context, memberId *int, reply *ReplyMemberInfo) error
	// 获取会员详情接口
	GetMemberDetail(ctx context.Context, memberId *int, reply *ReplyMemberDetail) error
	// 获取会员基础信息档案接口
	GetMemberBase(ctx context.Context, memberId *int, reply *ReplyMemberBase) error
	// 修改会员基础信息档案接口
	EditMemberBase(ctx context.Context, memberInfo *ArgsMemberBase, reply *bool) error
	// 平台会员快速成为门店会员
	UserToMember(ctx context.Context, args *ArgsUserToMember, memberId *int) error
	// 根据uid获取店铺会员折扣信息
	GetMemberRebateByUid(ctx context.Context, args *ArgsMemberRebate, reply *BusLevelDetail) error
	//根据手机号获取店铺会员信息-rpc
	GetUserInfoByPhoneRpc(ctx context.Context, args *ArgsGetUserInfoByPhone, reply *ReplyGetUserInfoByPhone) error
	//发送会员短信入参
	SendMemberSms(ctx context.Context, args *ArgsSendMemberSms, reply *bool) error
	//确认消费完成,添加店铺会员消费数据统计
	AddMemberConsumeCount(ctx context.Context, consumeLogId *int, reply *ReplyAddMemberConsumeCount) error
	//获取店铺会员消费数据
	GetMemberConsumeCount(ctx context.Context, memberId *int, reply *ReplyGetMemberConsumeCount) error
	//购买卡包后，更新会员持卡的量,如果会员不存在则增加一条数据
	AddMemberCardNum(ctx context.Context,relationId *int,reply *ReplyAddMemberConsumeCount)error
	//支付成功自动成为商家会员
	PayUserJoin(ctx context.Context, orderSn *string, reply *bool ) error
	// 获取会员总人数
	GetMemberNum(ctx context.Context, args *ArgsRiskMemberNum, reply *ReplyRiskMemberNum) error
	//根据指定店铺的多个uid查询会员信息  -rpc
	GetMemberInfos(ctx context.Context, args *ArgsGetMemberInfo, reply *[]ReplyGetMemberInfo) error
	//根据手机号模糊匹配门店会员
	GetUserInfoByPhoneMatch(ctx context.Context,args *ArgsGetUserInfoByPhoneMatch,reply *ReplyGetUserInfoByPhoneMatch)error
}
