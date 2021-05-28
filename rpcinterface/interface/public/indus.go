package public

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//返回领域、行业所有数据
type ReplyIndusAll struct {
	IndustryId      int              //领域ID
	IndustryName    string           //领域名称
	IndustryIconUrl string           //领域图片地址
	Sub             []ReplyIndusInfo //行业
}

//返回领域下的行业数据请求参数
type ArgsIndusInfo struct {
	IndustryId int //领域ID
}

//返回领域下的行业数据
type ReplyIndusInfo struct {
	IndusId   int    //行业ID
	Name      string //行业名称
	IconImgId int    //图片ID
}

//返回所有领域数据
type ReplyIndustryInfo struct {
	IndustryId      int    //领域ID
	IndustryName    string //领域名称
	IconImgId       int    //图片ID
	IndustryIconUrl string //图片icon地址
}

//添加领域入参
type ArgsIndustryAdd struct {
	common.Autoken
	Name string
}

//修改领域入参
type ArgsIndustryUpdate struct {
	common.Autoken
	Id   int
	Name string
}

//删除领域入参
type ArgsIndustryDel struct {
	common.Autoken
	Id int
}

//查询领域入参
type ArgsIndustryGet struct {
}

//查询领域返回
type Industry struct {
	Id   int    `mapstructure:"industry_id"`
	Name string `mapstructure:"industry_name"`
}

type IndusInfo struct {
	Id         int    //行业id
	IndustryId int    //领域id
	Name       string //行业名称
	ImgHash    string //行业封面图片哈希
}

//添加行业入参
type ArgsIndusAdd struct {
	common.Autoken
	IndusInfo
}

//批量添加行业入参
type ArgsIndusAdds struct {
	common.Autoken
	IndustryId int //领域id
	List       []IndusAdd
}

type IndusAdd struct {
	Name    string //行业名称
	ImgHash string //行业封面图片哈希
}

//删除一条行业入参
type ArgsIndusDel struct {
	common.Autoken
	Id int
}

//批量删除行业入参
type ArgsIndusDels struct {
	common.Autoken
	Ids []int
}

//按领域id查询行业入参
type ArgsIndusGet struct {
	IndustryId int
}

//查询行业返回
type IndusRes struct {
	Id   int    `mapstructure:"indus_id"`
	Name string `mapstructure:"name"`
	Img
}

//图片
type Img struct {
	Hash string
	Path string
}

//无参
type Args struct{}

//查询所有领域级下属行业 返回
type Reply struct {
	IndustryInfos []IndustryInfo
}

//领域
type IndustryInfo struct {
	IndustryId   int    `mapstructure:"industry_id"`
	IndustryName string `mapstructure:"industry_name"`
	Sub          []IndusInfos
}

//根据领域IDs获取领域信息入参
type ArgsGetIndustrysByIds struct {
	Ids []int
}

//行业
type IndusInfos struct {
	IndusId int    `mapstructure:"indus_id"`
	Name    string `mapstructure:"name"`
	ImgPath string `mapstructure:"imgPath"`
}

type IndusDetail struct {
	IndusId int    `mapstructure:"indus_id"`
	Name    string `mapstructure:"name"`
	ImgPath string
	ImgHash string
}

type Indus interface {
	//获取领域、行业所有数据
	GetAll(ctx context.Context, reply *[]ReplyIndusAll) error
	//获取领域下的所有行业
	GetIndusByIndustryId(ctx context.Context, industryId *ArgsIndusInfo, reply *[]ReplyIndusInfo) error
	//获取所有领域
	GetIndustry(ctx context.Context, reply *ReplyIndustryInfo) error

	//添加领域
	AddIndustry(ctx context.Context, args *ArgsIndustryAdd, reply *int) error
	//领域改名
	UpdateIndustry(ctx context.Context, args *ArgsIndustryUpdate, reply *bool) error
	//删除领域
	DelIndustry(ctx context.Context, args *ArgsIndustryDel, reply *bool) error
	//查询所有领域
	GetIndustrys(ctx context.Context, args *ArgsIndustryGet, reply *[]Industry) error
	//根据领域IDs获取领域
	GetIndustrysByIds(ctx context.Context, args *ArgsGetIndustrysByIds, reply *[]Industry) error

	//添加行业
	AddIndus(ctx context.Context, args *ArgsIndusAdd, reply *int) error
	//批量添加行业
	AddMoreIndus(ctx context.Context, args *ArgsIndusAdds, reply *int) error
	//修改行业信息
	UpdateIndus(ctx context.Context, args *ArgsIndusAdd, reply *bool) error
	//删除一条行业
	DelIndus(ctx context.Context, args *ArgsIndusDel, reply *bool) error
	//批量删除行业
	DelMoreIndus(ctx context.Context, args *ArgsIndusDels, reply *bool) error
	//根据领域Id查询行业
	GetInduss(ctx context.Context, args *ArgsIndusGet, reply *[]IndusRes) error
	//获取全部领域和行业
	GetAllInfos(ctx context.Context, args *Args, reply *Reply) error
	//根据行业id查询一条行业详情
	GetIndusDetail(ctx context.Context, args *int, reply *IndusDetail) error
}
