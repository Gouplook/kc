package news

import (
	"context"
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	// 收藏动作
	// 未收藏
	COLLECT_NO = 0
	// 已收藏
	COLLECT_YES = 1

	// 收藏类型
	// 攻略
	Collect_Raider = 1
	// 探店
	Collect_Exploration = 2
	// 单项目
	Collect_Signle = 3
	// 服务套餐
	Collect_Sm = 4
	// 综合卡
	Collect_Card = 5
	// 限时卡
	Collect_Hcard = 6
	// 限次卡
	Collect_Ncard = 7
	// 限时限次卡
	Collect_Hncard = 8
	// 产品
	Collect_Product = 9
	// 门店
	Collect_Shop = 10
)

func getColletActionList() []int {
	return []int{
		COLLECT_NO,
		COLLECT_YES,
	}
}

// VerifyColletAction 验证收藏动作
func VerifyColletAction(action int) bool {
	return functions.InArray(action, getColletActionList())
}

func getCollectTypeList() []int {
	return []int{
		Collect_Raider, Collect_Exploration,
		//Collect_Signle, Collect_Sm, Collect_Card,
		//Collect_Hcard, Collect_Ncard, Collect_Hncard, Collect_Product, Collect_Shop,
	}
}

// VerifyCollectType 验证收藏的类型
func VerifyCollectType(collectType int) bool {
	return functions.InArray(collectType, getCollectTypeList())
}

// ArgsUserCollect 前台收藏入参
type ArgsUserCollect struct {
	common.Utoken
	CollectId int
	Type      int
}

// CollectBase 收藏基础结构体
type CollectBase struct {
	Id        int `mapstructure:id`
	CollectId int `mapstructure:"collect_id"`
	Type      int `mapstructure:"type"`
	Uid       int `mapstructure:"uid"`
	Attention int `mapstructure:"attention"`
}

// ArgsUserCollectList 用户收藏的文章:攻略/探店
type ArgsUserCollectList struct {
	common.Paging
	common.Utoken
}
type UserCollectListBase struct {
	CollectId int  `mapstructure:"collect_id"`
	Title            string `mapstructure:"title"`
	ImageID          int    `mapstructure:"image_id"` // 封面图片ID
	ImageUrl         string
	AuthorID         int    `mapstructure:"author_id"`
	Author           string `mapstructure:"author"` // 探员
	Introduction     string // 探员简介
	AvatarImgUrl     string // 探员头像
	BackgroundImgUrl string // 探员背景
	CollectionCount  int    `mapstructure:"collection_count"` // 收藏量
	View             int
	Type             int `mapstructure:"type"` // 1-攻略;2-探店
}
type ReplyUserCollectList struct {
	TotalNum int
	Lists     []UserCollectListBase
}

type Collect interface {
	// UserCollect 用户收藏
	UserCollect(ctx context.Context, args *ArgsUserCollect, reply *bool) error
	// UserCollectList 用户收藏的攻略和探店
	UserCollectList(ctx context.Context, args *ArgsUserCollectList, reply *ReplyUserCollectList) error
}
