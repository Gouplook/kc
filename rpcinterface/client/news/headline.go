package news

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/news"
)

type Headline struct {
	client.Baseclient
}

func (h *Headline) Init() *Headline {
	h.ServiceName = "rpc_news"
	h.ServicePath = "Headline"
	return h
}

//后台－添加
func (h *Headline) AddHeadline(ctx context.Context, req *news.AddHeadlineReq, resp *news.HeadlineID) (err error) {
	return h.Call(ctx, "AddHeadline", req, resp)
}

//后台－删除
func (h *Headline) DeleteHeadline(ctx context.Context, req *news.HeadlineID, resp *news.Empty) (err error) {
	return h.Call(ctx, "DeleteHeadline", req, resp)
}

//后台－更新
func (h *Headline) UpdateHeadline(ctx context.Context, req *news.HeadlineReq, resp *news.Empty) (err error) {
	return h.Call(ctx, "UpdateHeadline", req, resp)
}

//后台－详情
func (h *Headline) HeadlineInfo(ctx context.Context, req *news.HeadlineID, resp *news.HeadlineResp) (err error) {
	return h.Call(ctx, "HeadlineInfo", req, resp)
}

//后台-列表
func (h *Headline) HeadlineList(ctx context.Context, req *news.PageReq, resp *news.HeadlineListResp) (err error) {
	return h.Call(ctx, "HeadlineList", req, resp)
}

//前台-详情
func (h *Headline) HeadlineInfo2(ctx context.Context, req *news.HeadlineID, resp *news.HeadlineResp) (err error) {
	return h.Call(ctx, "HeadlineInfo2", req, resp)
}

//前台－推荐阅读列表
func (h *Headline) RecommendationReadingList(ctx context.Context, req *news.RecommendationReadingReq, resp *news.RecommendationReadingResp) (err error) {
	return h.Call(ctx, "RecommendationReadingList", req, resp)
}

//前台头条列表
func (h *Headline) WebHeadlineList(ctx context.Context, req *news.ArgsWebHeadlineList, resp *news.ReplyWebHeadlineList) (err error) {
	return h.Call(ctx, "WebHeadlineList", req, resp)
}

// 前台滚动的推荐头条
func (h *Headline) RollRecommendHeadline(ctx context.Context, req *news.ArgsRollRecommendHeadline, resp *[]news.ReplyRollRecommendHeadline) (err error) {
	return h.Call(ctx, "RollRecommendHeadline", req, resp)
}
