package admin

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	//是否删除
	IS_DEL_YES = 1 //已删除
	IS_DEL_NO  = 0 //未删除
)

type AdvSpaInfo struct {
	Id      int
	Name    string //名称
	ImgWide int    `mapstructure:"img_wide"` //图片宽
	ImgHigh int    `mapstructure:"img_high"` //图片高
}

//广告位增加 和 修改
type ArgsAdvSpaAdd struct {
	common.Autoken
	AdvSpaInfo
}

//按分页 查询 广告位
type ArgsAdvSpaGet struct {
	common.Paging
}

//按id查询一条
type ArgsAdvSpaGetOne struct {
	Id int
}

//广告位 返回
type ReplyAdvSpaPage struct {
	TotalNum int
	List     []AdvSpaInfo
}

//广告位删除
type ArgsAdvSpaDel struct {
	common.Autoken
	Id int
}

//广告 新增 和 修改
type ArgsAdvAdd struct {
	common.Autoken
	AdvAdd
}
type AdvAdd struct {
	Id        int
	Title     string
	SpaId     int
	CityId    int
	Url       string
	UrlType  int
	Sort      int
	ShortDesc string
	Price     float64
	ImgHash   string
}

//广告查询
type ArgsAdvGet struct {
	common.Paging
	SpaId int //广告类型id
}

//查询返回
type ReplyAdvPage struct {
	TotalNum int
	List     []AdvRes
}
type AdvRes struct {
	Id     int
	Title  string
	SpaId  int `mapstructure:"spa_id"`
	CityId int `mapstructure:"city_id"`
	Url    string
	Sort   int
	UrlType int  `mapstructure:"url_type"`
	//Desc	string
	//Price	float64
}

//广告查询一条
type ArgsAdvGetOne struct {
	Id int
}

//查询一条返回
type ReplyAdvInfo struct {
	ShortDesc string `mapstructure:"short_desc"`
	Price     float64
	Img
}

//图片
type Img struct {
	Hash string
	Path string
}

//删除一条广告
type ArgsAdvDel struct {
	common.Autoken
	Id int
}

type Args struct{}

//根据广告位Id 查询所属广告
type ArgsBySpaId struct {
	PageSize int
	SpaId    int
}

//查询全部广告位和广告 返回
type ReplyBySpaId struct {
	Lists       []AdvInfo
	ImgPathMap map[int]string
}
type AdvInfo struct {
	Id    int    `mapstructure:"id"`
	Title string `mapstructure:"title"`
	Cid  int `mapstructure:"city_id"`
	Cname string
	Url   string `mapstructure:"url"`
	ImgId int    `mapstructure:"img_id"`
	Sort  int    `mapstructure:"sort"`
	UrlType int `mapstructure:"url_type"`
	ShortDesc string
}

type Adv interface {

	/*//添加广告位
	AddAdvSpa(context.Context, *ArgsAdvSpaAdd, *int) error*/
	//修改广告位
	UpdateAdvSpa(context.Context, *ArgsAdvSpaAdd, *bool) error
	//删除广告位
	DelAdvSpa(context.Context, *ArgsAdvSpaDel, *bool) error
	//按分页查询广告位
	GetAdvSpaByPage(context.Context, *ArgsAdvSpaGet, *ReplyAdvSpaPage) error
	//无条件查询广告位
	GetAdvSpas(context.Context, *Args, *map[string]string) error
	//按id查询一条广告位
	GetAdvSpaOne(context.Context, *ArgsAdvSpaGetOne, *AdvSpaInfo) error

	//添加 广告
	AddAdv(context.Context, *ArgsAdvAdd, *int) error
	//修改 广告
	UpdateAdv(context.Context, *ArgsAdvAdd, *bool) error
	//删除 广告
	DelAdv(context.Context, *ArgsAdvDel, *bool) error
	//查询广告
	GetAdvs(context.Context, *ArgsAdvGet, *ReplyAdvPage) error
	//查询一条广告
	GetAdv(context.Context, *ArgsAdvGetOne, *ReplyAdvInfo) error

	//前台查询
	//根据广告位Id查询所属广告
	GetAdvBySpaId(context.Context, *ArgsBySpaId, *ReplyBySpaId) error
}
