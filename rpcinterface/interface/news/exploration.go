package news

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

type ExplorationItemResp struct {
	Content   string `mapstructure:"content"`
	ImageID   int    `mapstructure:"image_id"`
	ImageUrl  string
	ImageHash string
}

type ExplorationItemReq struct {
	Content   string
	ImageHash string
}

type AddExplorationReq struct {
	Title                   string
	Area                    string // 多个地区以","隔开
	AuthorID                int
	Recommendation          bool
	RecommendationStartTime int64
	RecommendationEndTime   int64
	ImageHash               string
	VideoUrl                string
	ShopID                  int
	Description             string
	Content                 string

	Items []ExplorationItemReq
}

type ExplorationID struct {
	ID int
}

type ExplorationReq struct {
	ID                      int
	Title                   string
	Area                    string
	AuthorID                int
	Recommendation          bool
	RecommendationStartTime int64
	RecommendationEndTime   int64
	ImageHash               string

	VideoUrl    string
	ShopID      int
	Description string
	Content     string

	Items []ExplorationItemReq
}

type ExplorationResp struct {
	ID                      int    `mapstructure:"id"`
	Title                   string `mapstructure:"title"`                     // 标题
	Area                    string `mapstructure:"area"`                      // 地区
	AuthorID                int    `mapstructure:"author_id"`                 // 探员ID
	Author                  string `mapstructure:"author"`                    // 探员
	Recommendation          bool   `mapstructure:"recommendation"`            //是否推荐
	RecommendationStartTime int64  `mapstructure:"recommendation_start_time"` // 推荐开始时间
	RecommendationEndTime   int64  `mapstructure:"recommendation_end_time"`   // 推荐结束时间
	ImageID                 int    `mapstructure:"image_id"`                  // 封面图片ID
	ImageUrl                string
	ImageHash               string
	VideoUrl                string `mapstructure:"video_url"`   // 关联视频连接
	ShopID                  int    `mapstructure:"shop_id"`     // 关联shopID
	Description             string `mapstructure:"description"` // 描述
	Content                 string `mapstructure:"content"`     // 正文
	View                    int

	Items []ExplorationItemResp
}

type ExplorationDescribe struct {
	ID              int    `mapstructure:"id"`
	Title           string `mapstructure:"title"`
	Area            string `mapstructure:"area"`
	AuthorID        int    `mapstructure:"author_id"`
	Author          string `mapstructure:"author"`
	Admin           bool   `mapstructure:"admin"`
	Recommendation  bool   `mapstructure:"recommendation"`
	ShopID          int    `mapstructure:"shop_id"`
	CommentCount    int    `mapstructure:"comment_count"`
	CollectionCount int    `mapstructure:"collection_count"`
	View            int
}

type ListReq struct {
	Page           int
	PageSize       int
	Source         uint8 // 0:全部, 1:后台, 2:用户
	Recommendation uint8 //0:全部, 1:推荐, 2:不推荐
}
type ExplorationList struct {
	TotalNum int
	Lists     []ExplorationDescribe
}

// ArgsExplorationListByAuthId 获取探员下所有的探店
type ArgsExplorationListByAuthId struct {
	common.Utoken
	common.Paging
	AuthorId int
}
type ExplorationListByAuthIdBase struct {
	ID              int
	Title           string // 标题
	CollectionCount int    `mapstructure:"collection_count"` // 收藏量
	View            int    // 浏览量
	ImageID         int    `mapstructure:"image_id"` // 封面图片ID
	ImageUrl        string
	ImageHash       string
	Ctime           int64
	CtimeStr        string
	Collection      bool // 用户是否收藏
}
type ReplyExplorationListByAuthId struct {
	TotalNum int
	Lists     []ExplorationListByAuthIdBase
}

type ArgsExplorationInfo struct {
	ID int
}
type ExplorationInfo struct {
	ID                      int    `mapstructure:"id"`
	Title                   string `mapstructure:"title"`                     // 标题
	Area                    string `mapstructure:"area"`                      // 地区
	AuthorID                int    `mapstructure:"author_id"`                 // 探员ID
	AvatarUrl string //用户头像
	Author                  string `mapstructure:"author"`                    // 探员
	Recommendation          bool   `mapstructure:"recommendation"`            //是否推荐
	RecommendationStartTime int64  `mapstructure:"recommendation_start_time"` // 推荐开始时间
	RecommendationEndTime   int64  `mapstructure:"recommendation_end_time"`   // 推荐结束时间
	ImageID                 int    `mapstructure:"image_id"`                  // 封面图片ID
	ImageUrl                string
	ImageHash               string
	VideoUrl                string `mapstructure:"video_url"`   // 关联视频连接
	ShopID                  int    `mapstructure:"shop_id"`     // 关联shopID
	ShopInfo			ShopInfo
	Description             string `mapstructure:"description"` // 描述
	Content                 string `mapstructure:"content"`     // 正文
	View                    int    // 浏览量
	Ctime                   int64
	CtimeStr                string
	Collection              bool // 用户是否收藏
}

type ShopInfo struct {
	BusId         int            //商户信息
	ShopId        int            //门店id
	ShopName      string         //门店名称
	BranchName    string         //分店名
	ShopAddress   string         //门店地址
	ShopImage     string         //门店图片
	ShopPhone     string         //门店电话
	IndustryId    int            //领域id
	MainBindId    int            //主行业id
	BindId        []int          //行业id集合
	BindNames     string         //行业名称
	Pid           int            //省id
	Cid           int            //市id
	Did           int            //区id
	DName         string         //区名称
	DistrictId    []int          //商圈id
	DistrictNames string         //商圈名称
	Lon           float64        //经度
	Lat           float64        //纬度
	CommentScore  float64        //平均评分
	Distance      int            //距离 m
	SafeCode	int			//商家安全码颜色 1=黑色 2=红色 3=黄色  4=绿色
}

type ReplyExplorationInfo struct {
	ExplorationInfo
	ShareLink string //分享链接
	Items []ExplorationItemResp
}

type UserExplorationListBase struct {
	ID               int    `mapstructure:"id"`
	Title            string `mapstructure:"title"`
	ImageID          int    `mapstructure:"image_id"` // 封面图片ID
	ImageUrl         string
	ImageHash        string
	AuthorID         int    `mapstructure:"author_id"`
	Author           string `mapstructure:"author"` // 探员
	Introduction     string `mapstructure:"introduction"` // 探员简介
	AvatarImgUrl     string // 探员头像
	BackgroundImgUrl string // 探员背景
	//Admin           bool   `mapstructure:"admin"`
	//Recommendation  bool   `mapstructure:"recommendation"`
	//ShopID          int    `mapstructure:"shop_id"`
	CommentCount    int `mapstructure:"comment_count"`    // 评论量
	CollectionCount int `mapstructure:"collection_count"` // 收藏量
	View            int
	Collection      bool // 用户是否收藏
}

// 前台探店列表入参
type ArgsUserExplorationList struct {
	Recommend bool
	common.Utoken
	common.Paging
}

// 前台探店列表出参
type ReplyUserExplorationList struct {
	TotalNum int
	Lists     []UserExplorationListBase
}

// 前台探店详情入参
type ArgsUserExplorationInfo struct {
	common.Utoken
	ID int
	Lat float64
	Lnt float64
}

// 前台探店推荐入参
type ArgsExplorationRecommend struct {
	common.Utoken
	common.Paging
}

// 前台探店推荐出参
type ReplyExplorationRecommend struct {
	TotalNum int
	Lists     []UserExplorationListBase
}

type Exploration interface {
	// 后台添加探店
	AddExploration(ctx context.Context, req *AddExplorationReq, resp *ExplorationID) (err error)
	// 后台更新探店
	UpdateExploration(ctx context.Context, req *ExplorationReq, resp *Empty) (err error)
	// 后台探店详情
	ExplorationInfo(ctx context.Context, req *ExplorationID, resp *ExplorationResp) (err error)
	ExplorationInfo2(ctx context.Context, req *ExplorationID, resp *ExplorationResp) (err error)
	// 后台探店列表
	ExplorationList(ctx context.Context, req *ListReq, resp *ExplorationList) (err error)

	// ExlorationListByAuthId 前台获取探员下的探店
	ExplorationListByAuthId(ctx context.Context, args *ArgsExplorationListByAuthId, reply *ReplyUserExplorationList) (err error)
	// UserExplorationInfo 前台探店详情
	UserExplorationInfo(ctx context.Context, args *ArgsUserExplorationInfo, reply *ReplyExplorationInfo) (err error)
	// UserExplorationList 前台探店列表
	UserExplorationList(ctx context.Context, args *ArgsUserExplorationList, reply *ReplyUserExplorationList) (err error)
	// ExplorationRecommend 前台探店推荐
	ExplorationRecommend(ctx context.Context, args *ArgsExplorationRecommend, reply *ReplyUserExplorationList) (err error)
}
