package news

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

//　AuthorBase　AuthorBase
type AuthorBase struct {
	Id              int    `mapstructure:"id"`
	AuthorId        int    `mapstructure:"author_id"`
	Type            int    `mapstructure:"type"`
	Recommend       int    `mapstructure:"recommend"`         // 是否推荐 0-不推荐 1-推荐
	AvatarImgId     int    `mapstructure:"avatar_img_id"`     // 探员头像id
	Introduction    string `mapstructure:"introduction"`      // 探员简介
	BackgroundImgId int    `mapstructure:"background_img_id"` // 探员背景图id
	Ctime           int64  `mapstructure:"ctime"`
	CtimeStr        string
}

// ArgsAddAuthor　添加作者入参
type ArgsAddAuthor struct {
	AuthorId          int
	Type              int
	AvatarImgHash     string // 探员头像hash
	Introduction      string // 探员简介
	BackgroundImgHash string // 探员背景图hash
}

// ArgsRecommendAuthor 是否推荐作者入参
type ArgsRecommendAuthor struct {
	Id        int
	Recommend int
}

// ArgsAuthorList 作者列表入参
type ArgsAuthorList struct {
	common.Paging
	Phone string
	Nick  string
}

// AuthorListBase AuthorListBase
type AuthorListBase struct {
	Id           int    `mapstructure:"id"`
	AuthorId     int    `mapstructure:"author_id"`
	Type         int    `mapstructure:"type"`
	Recommend    int    `mapstructure:"recommend"`    // 是否推荐 0-不推荐 1-推荐
	Introduction string `mapstructure:"introduction"` // 探员简介
	Ctime        int64  `mapstructure:"ctime"`
	CtimeStr     string
	Nick         string
}

// ReplyAuthorList 作者列表出参
type ReplyAuthorList struct {
	TotalNum int
	Lists     []AuthorListBase
}

// ArgsUpdateAuthor 更新探员信息
type ArgsUpdateAuthor struct {
	Id                int
	AvatarImgHash     string // 探员头像hash
	Introduction      string // 探员简介
	BackgroundImgHash string // 探员背景图hash
}

// ReplyAuthorInfo ReplyAuthorInfo
type ReplyAuthorInfo struct {
	AuthorBase
	AvatarImgHash     string // 探员头像hash
	BackgroundImgHash string // 探员背景图hash
	AvatarImgUrl      string // 探员头像url
	BackgroundImgUrl  string // 探员背景图url
}

// ArgsGetAuthors 获取作者入参
type ArgsGetAuthors struct {
	Type int
}

// ReplyGetAuthors 获取作者出参
type ReplyGetAuthors struct {
	AuthorId int `mapstructure:"author_id"`
	Nick     string
}

// ArgsUserAuthorList 探员列表入参
type ArgsUserAuthorList struct {
	common.Paging
	Nick string
	Type int
	GetAll bool // 全部数据,为true时,type无效
}

// UserAuthorListBase UserAuthorListBase
type UserAuthorListBase struct {
	AuthorId         int    `mapstructure:"author_id"`
	Nick             string // 昵称
	Introduction     string `mapstructure:"introduction"`  // 探员简介
	AvatarImgId      int    `mapstructure:"avatar_img_id"` // 探员头像id
	AvatarImgUrl     string // 探员头像
	Recommend        int    `mapstructure:"recommend"`         // 是否推荐
	BackgroundImgId  int    `mapstructure:"background_img_id"` // 探员背景ID
	BackgroundImgUrl string // 探员背景
	Type             int    `mapstructure:"type"` // 类型
}

// ReplyUserAuthorList 探员列表出参
type ReplyUserAuthorList struct {
	TotalNum int
	Lists     []UserAuthorListBase
}

// ArgsRecommendAuthorList 推荐探员/作者
type ArgsRecommendAuthorList struct {
	Type int
	GetAll bool // 全部推荐,为true时,type无效
	common.Paging
}
type ReplyRecommendAuthorList struct {
	TotalNum int
	Lists     []UserAuthorListBase
}
type Author interface {
	// AddAuthor　添加作者
	AddAuthor(ctx context.Context, args *ArgsAddAuthor, reply *bool) error
	// DelAuthor 删除作者
	DelAuthor(ctx context.Context, id *int, reply *bool) error
	// RecommendAuthor 是否推荐作者
	RecommendAuthor(ctx context.Context, args *ArgsRecommendAuthor, reply *bool) error
	// AuthorList 作者列表
	AuthorList(ctx context.Context, args *ArgsAuthorList, reply *ReplyAuthorList) error
	// AuthorInfo 作者详情
	AuthorInfo(ctx context.Context, id *int, reply *ReplyAuthorInfo) error
	// UpdateAuthor 更新探员作者
	UpdateAuthor(ctx context.Context, args *ArgsUpdateAuthor, reply *bool) error
	// GetAuthors  获取作者,只获取Nick和UID
	GetAuthors(ctx context.Context, args *ArgsGetAuthors, reply *[]ReplyGetAuthors) error

	// UserAuthorList 前台探员列表
	UserAuthorList(ctx context.Context, args *ArgsUserAuthorList, reply *ReplyUserAuthorList) error
	// RecommendAuthorList 前台探员/作者推荐
	RecommendAuthorList(ctx context.Context, args *ArgsRecommendAuthorList, reply *ReplyRecommendAuthorList) error
}
