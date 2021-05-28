package news

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

type Empty struct {
}

type AddHeadlineReq struct {
	Title          string
	Author         string
	Source         string
	Link           string
	ImgHash        string
	Recommendation bool
	ReleaseTime    int64
	Synopsis       string
	Content        string
}

type HeadlineID struct {
	ID int
}

type HeadlineReq struct {
	ID     int
	Title  string
	Author string
	Source string
	Link   string

	ImgHash string

	Recommendation bool
	ReleaseTime    int64
	Synopsis       string
	Content        string
}

type HeadlineResp struct {
	ID       int    `mapstructure:"id"`
	Title    string `mapstructure:"title"`
	Author   string `mapstructure:"author"`
	Source   string `mapstructure:"source"`
	Link     string `mapstructure:"link"`
	ImgHash  string
	ImageUrl string

	Recommendation bool   `mapstructure:"recommendation"`
	ReleaseTime    int64  `mapstructure:"release_time"`
	ReleaseTimeStr string // 发布时间str
	Ctime          int64  `mapstructure:"ctime"`
	CtimeStr       string
	Synopsis       string `mapstructure:"synopsis"`
	Content        string `mapstructure:"content"`

	View int
	ShareLink string //分享链接
}

type PageReq struct {
	Page     int
	PageSize int
}

type RecommendationReadingReq struct {
	common.Paging
	common.Utoken
	Type  int // 1:最新头条(时间最新);2:最热头条数据(浏览量最多)
	Count int // 推荐阅读显示量
}

type HeadlineListBase struct {
	ID             int    `mapstructure:"id"`
	Title          string `mapstructure:"title"`
	Author         string `mapstructure:"author"`
	ReleaseTime    int64  `mapstructure:"release_time"`
	ReleaseTimeStr string
	Ctime          int64 `mapstructure:"ctime"`
	CtimeStr       string
	View           int
}

type HeadlineListResp struct {
	TotalNum int
	Lists     []HeadlineListBase
}
type RecommendationReadingBase struct {
	ID             int    `mapstructure:"id"`
	Title          string `mapstructure:"title"`
	Author         string `mapstructure:"author"`
	ReleaseTime    int64  `mapstructure:"release_time"`
	ReleaseTimeStr string
	Synopsis       string `mapstructure:"synopsis"`
	Ctime          int64  `mapstructure:"ctime"`
	CtimeStr       string
	ImgUrl         string `mapstructure:"img_url"`
	ImgId          int    `mapstructure:"img_id"`
	View           int
}

type RecommendationReadingResp struct {
	TotalNum int
	Lists     []RecommendationReadingBase
}

type RecomendReadBase []RecommendationReadingBase

func (slice RecomendReadBase) Len() int {
	return len(slice)
}

func (slice RecomendReadBase) Less(i, j int) bool {
	return slice[i].View > slice[j].View
}

func (slice RecomendReadBase) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// 前台头条列表入参
type ArgsWebHeadlineList struct {
	Count int
	common.Paging
	common.Utoken
}

// 前台头条列表出参
type ReplyWebHeadlineList struct {
	TotalNum int
	Lists     []HeadlineListBase
}

type ArgsRollRecommendHeadline struct {
}

// 前台滚动的推荐头条入参
type ReplyRollRecommendHeadline struct {
	ID    int    `mapstructure:"id"`
	Title string `mapstructure:"title"`
}

type Headline interface {
	// 后台添加头条
	AddHeadline(ctx context.Context, req *AddHeadlineReq, resp *HeadlineID) (err error)
	// 后台删除头条
	DeleteHeadline(ctx context.Context, req *HeadlineID, resp *Empty) (err error)
	// 后台更新头条
	UpdateHeadline(ctx context.Context, req *HeadlineReq, resp *Empty) (err error)
	// 后台头条详情
	HeadlineInfo(ctx context.Context, req *HeadlineID, resp *HeadlineResp) (err error)
	// 后台头条列表
	HeadlineList(ctx context.Context, req *PageReq, resp *HeadlineListResp) (err error)
	//　前台头条详情
	HeadlineInfo2(ctx context.Context, req *HeadlineID, resp *HeadlineResp) (err error)
	// 前台头条推荐列表
	RecommendationReadingList(ctx context.Context, req *RecommendationReadingReq, resp *RecommendationReadingResp) (err error)
	//　前台头条列表
	WebHeadlineList(ctx context.Context, req *ArgsWebHeadlineList, resp *ReplyWebHeadlineList) (err error)
	// 前台滚动的推荐头条
	RollRecommendHeadline(ctx context.Context, req *ArgsRollRecommendHeadline, resp *[]ReplyRollRecommendHeadline) (err error)
}
