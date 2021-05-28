package news

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/news"
)

type Author struct {
	client.Baseclient
}

func (a *Author) Init() *Author {
	a.ServiceName = "rpc_news"
	a.ServicePath = "Author"
	return a
}

// AddAuthor　添加作者
func (a *Author) AddAuthor(ctx context.Context, args *news.ArgsAddAuthor, reply *bool) error {
	return a.Call(ctx, "AddAuthor", args, reply)
}

// DelAuthor 删除作者
func (a *Author) DelAuthor(ctx context.Context, id *int, reply *bool) error {
	return a.Call(ctx, "DelAuthor", id, reply)
}

// RecommendAuthor 是否推荐作者
func (a *Author) RecommendAuthor(ctx context.Context, args *news.ArgsRecommendAuthor, reply *bool) error {
	return a.Call(ctx, "RecommendAuthor", args, reply)
}

// AuthorList 作者列表
func (a *Author) AuthorList(ctx context.Context, args *news.ArgsAuthorList, reply *news.ReplyAuthorList) error {
	return a.Call(ctx, "AuthorList", args, reply)
}

// AuthorInfo 作者详情
func (a *Author) AuthorInfo(ctx context.Context, id *int, reply *news.ReplyAuthorInfo) error {
	return a.Call(ctx, "AuthorInfo", id, reply)
}

// UpdateAuthor 更新探员作者
func (a *Author) UpdateAuthor(ctx context.Context, args *news.ArgsUpdateAuthor, reply *bool) error {
	return a.Call(ctx, "UpdateAuthor", args, reply)
}

// GetAuthors  获取作者,只获取Nick和UID
func (a *Author) GetAuthors(ctx context.Context, args *news.ArgsGetAuthors, reply *[]news.ReplyGetAuthors) error {
	return a.Call(ctx, "GetAuthors", args, reply)
}

// UserAuthorList 前台探员列表
func (a *Author) UserAuthorList(ctx context.Context, args *news.ArgsUserAuthorList, reply *news.ReplyUserAuthorList) error {
	return a.Call(ctx, "UserAuthorList", args, reply)
}

// RecommendAuthorList 前台探员/作者推荐
func (a *Author) RecommendAuthorList(ctx context.Context, args *news.ArgsRecommendAuthorList, reply *news.ReplyRecommendAuthorList) error {
	return a.Call(ctx, "RecommendAuthorList", args, reply)
}
