package news

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

const (
	IS_DEL_YES = 1
	IS_DEL_NO  = 0

	MY_QUIZ    = 1
	MY_REPLY   = 2
	MY_COLLECT = 3

	CHECK_STATUS_PASS = 2
)

//添加一级分类入参
type ArgsCateAdd struct {
	common.Autoken
	Name string //一级分类名称
}

//添加子分类
type ArgsSubCateAdd struct {
	common.Autoken
	Id      int    //一级分类id
	Name    string //二级分类名称
	ImgHash string //分类图片
}

//无参
type Args struct{}

//查询所有一级和子类分类
type ReplyCateAll struct {
	CateInfo   []CateInfo
	ImgPathMap map[int]string
}
type CateInfo struct {
	Id   int    `mapstructure:"id"`
	Name string `mapstructure:"name"`
	Sub  []SubCateInfo
}
type SubCateInfo struct {
	Id    int    `mapstructure:"id"`
	Name  string `mapstructure:"name"`
	ImgId string `mapstructure:"img_id"`
}

//按分页 推荐 审核 查询提问 后台查询
type ArgsQuizGetP struct {
	common.Paging
	IsRec  int //1不推荐  2推荐
	Status int //审核状态 1待审核 2已通过 3未通过
}

//后台查询提问返回
type ReplyQuizPage struct {
	TotalNum int
	Lists     []QuizRes
}
type QuizRes struct {
	Id        int    `mapstructure:"id"`      //提问id
	Context   string `mapstructure:"content"` //问题内容
	CateId    int    `mapstructure:"cate_id"` //所属分类
	CateName  string
	GiveALike int    `mapstructure:"give_a_like"` //点赞 数量
	Answer    int    `mapstructure:"answer"`      //回答 数量
	Collect   int    `mapstructure:"collect"`     //关注 数量
	Time      string `mapstructure:"create_time"` //时间
	Status    int    `mapstructure:"status"`      //审核状态默认2  1待审核 2已通过 3未通过
	IsRec     int    `mapstructure:"is_rec"`      //是否推荐首页默认1  1不推荐 2推荐
	Source    int    `mapstructure:"source"`      //来源默认0  0是用户 1是后台
}

//添加提问
type ArgsQuizAdd struct {
	common.Autoken
	Context  string           //问题内容
	Id       int              //问题Id
	AuthorId int              //作者id
	CateId   int              //二级分类id
	Source   int              //0是用户 1是平台
	Status   int              //审核状态 默认 1待审核 2已通过 3未通过   //后台用
	IsRec    int              //是否推荐 默认 1不推荐 2推荐   //后台用
	Imgs     []ImgHashExplain //图片和说明
	RecTime  string           //推荐日期 起始时间-结束时间
}
type ImgHashExplain struct {
	ImgHash string
	Explain string
}

//查询一条问题详情
type ReplyQuizOne struct {
	Context   string   `mapstructure:"content"`   //问题内容
	Id        int      `mapstructure:"id"`        //问题Id
	AuthorId  int      `mapstructure:"author_id"` //作者id
	CateId    int      `mapstructure:"cate_id"`   //二级分类id
	Status    int      `mapstructure:"status"`    //审核状态 默认 1待审核 2已通过 3未通过   //后台用
	IsRec     int      `mapstructure:"is_rec"`    //是否推荐 默认 1不推荐 2推荐   //后台用
	Imgs      []ImgExp //图片和说明
	StartTime string   `mapstructure:"start_time"` //推荐日期 起始时间
	EndTime   string   `mapstructure:"end_time"`   //推荐日期 结束时间
}

type ImgExp struct {
	ImgHash string //图片hash
	ImgExp  string //文章
	ImgUrl  string //图片url
}

//审核
type ArgsQuizCheck struct {
	common.Autoken
	Id     int //问题id
	Status int //2已通过 3未通过
}

//根据用户id查询问题
type ArgsQuizGet struct {
	Uid int
}

//根据用户id查询问题返回 //app 我的提问  我的回复   我的收藏
type ReplyQuiz struct {
	MyQuiz    []MyQuiz                 `mapstructure:"myQuiz"`    //我的提问数据
	MyReply   []MyReply                `mapstructure:"myReply"`   //我的回复
	MyCollect []MyCollect              `mapstructure:"myCollect"` //我的收藏
	ReplyNum  []map[string]interface{} //回复数量
}
type MyQuiz struct {
	Id      int    `mapstructure:"quiz_id"` //提问id
	Context string `mapstructure:"content"` //提问内容
	//Answer  int    //回答 数量
}
type MyReply struct {
	Id      int    `mapstructure:"quiz_id"` //提问id
	Context string `mapstructure:"content"` //回复内容
	//Answer  int    //回答 数量
}
type MyCollect struct {
	Id      int    `mapstructure:"quiz_id"` //提问id
	Context string `mapstructure:"content"` //提问内容
	//Answer  int    //回答 数量
}

//APP端查询问题入参
type ArgsQuizGetA struct {
	common.Paging
	Uid int
}

//APP端查询问题返回
type ReplyQuizGetA struct {
	List    []QuizInfo
	QuizIds []int
}

//问题查询返回
type QuizInfo struct {
	Id         int          `mapstructure:"id"`          //提问id
	Uid        int          `mapstructure:"author_id"`   //用户id
	Context    string       `mapstructure:"content"`     //问题内容
	CateId     int          `mapstructure:"cate_id"`     //所属分类
	GiveALike  int          `mapstructure:"give_a_like"` //点赞 数量
	Answer     int          `mapstructure:"answer"`      //回答 数量
	Collect    int          `mapstructure:"collect"`     //关注 数量
	CreateTime string       `mapstructure:"create_time"` //时间
	List       []ContextImg //图片和说明的集合
}

//问题图片加说明
type ContextImg struct {
	ImgPath string
	Explain string
}

//根据问题id查询 答案和评论
type ArgsQuizD struct {
	Uid int
	Id  int
}

//根据问题id查询 答案和评论 //列表进入查询返回  //答案或者评论回复
type ReplyInfo struct {
	Uid        int    `mapstructure:"uid"`         //用户id
	Id         int    `mapstructure:"id"`          //答案id
	Reply      int    `mapstructure:"reply"`       //回复数量
	GiveALike  int    `mapstructure:"give_a_like"` //点赞数量
	Content    string `mapstructure:"content"`     //内容
	CreateTime string `mapstructure:"create_time"` //回复时间
	//ParentId   int    //0代表 是答案  有值代表 是评论一条答案
	Sub []CommentInfo `mapstructure:"sub"`
}
type CommentInfo struct {
	Id      int    `mapstructure:"id"`
	Context string `mapstructure:"content"`
	Uid     int    `mapstructure:"uid"`
}

//根据问题id查询 答案和评论 //我的问答 进入查询返回
type ReplyQuizB struct {
	QuizInfo  QuizInfo    //提问主题信息
	ReplyInfo []ReplyInfo //回复评论信息
	IsCollect bool
}

//按 分页 审核 查询答案或者评论
type ArgsAnswerGetP struct {
	common.Paging
	Status int //审核状态 默认 0待审核 1已通过 2未通过
}

//按 分页 审核 查询答案
type ReplyAnswerPage struct {
	TotalNum int
	List     []Answer
}
type Answer struct {
	Id         int    `mapstructure:"id"`          //答案id
	Content    string `mapstructure:"content"`     //回答内容
	GiveALike  int    `mapstructure:"give_a_like"` //点赞数量
	Reply      int    `mapstructure:"reply"`       //回复数量
	CreateTime string `mapstructure:"create_time"` //回答时间
	Status     int    `mapstructure:"status"`      //审核状态
}

//通过审核 或者 拒绝审核 答案 or 评论
type ArgsAnswerCheck struct {
	common.Autoken
	Id     int //回复 id
	Status int //2通过 3未通过
}

//按 分页 审核 查询评论
type ReplyCommentPage struct {
	TotalNum int
	List     []Comment
}
type Comment struct {
	Id         int    `mapstructure:"id"`          //评论id
	Content    string `mapstructure:"content"`     //评论内容
	Status     int    `mapstructure:"status"`      //审核状态
	CreateTime string `mapstructure:"create_time"` //评论时间
}

//添加 关键字 或者修改
type ArgsKeyWordAdd struct {
	common.Autoken
	Id      int    //关键字 id
	ModId   int    //关键字 所属模块id
	KeyWord string //关键字
}

//删除关键字
type ArgsKeyWordDel struct {
	common.Autoken
	Id int
}

//用户 添加 回复
type ArgsAnswerAdd struct {
	common.Autoken
	QuizId   int    //提问id
	AnswerId int    //答案id  //如果是 评论 必须要传答案id
	Content  string //回复内容
}

//用户收藏提问
type ArgsAnswerCollect struct {
	common.Autoken
	QuizId  int
	Content string
}

//用户点赞
type ArgsGiveALike struct {
	common.Autoken
	QuizId   int //哪个有值 就是给哪个 点赞
	AnswerId int
}

//按模块查询关键字
type ArgsKeyWordGet struct {
	common.Paging
	ModId int //  可传 可不传
}

//查询所有关键字返回
type ReplyKeyWord struct {
	TotalNum int
	List     []KeyWordInfo
}
type KeyWordInfo struct {
	Id      int
	KeyWord string
	ModId   int `mapstructure:"mod_id"`
}

//用户添加提问
type ArgsQuizAdd2 struct {
	common.Autoken
	CateId  int
	Content string
	Imgs    []ImgHashExplain //图片和说明
}

//用户取消收藏
type ArgsCollectCancel struct {
	common.Autoken
	QuizId int
}

//根据二级分类id查询二级分类详情
type ReplyCateDetail struct {
	Name    string
	ImgHash string
	ImgUrl  string
}

type Quiz interface {

	//提问 三步骤     其他权益
	//  1 提问  	    点赞 收藏 回复
	//  2 回答	    点赞 回复
	//  3 回复答案    无

	//后台

	//添加一级分类
	AddCate(ctx context.Context, args *ArgsCateAdd, reply *int) error
	//添加二级分类
	AddSubCate(ctx context.Context, args *ArgsSubCateAdd, reply *int) error
	//修改二级分类
	UpdateSubCate(ctx context.Context, args *ArgsSubCateAdd, reply *bool) error
	//查询所有一级分类一级下属二级分类
	GetCateAll(ctx context.Context, args *Args, reply *ReplyCateAll) error
	//按分页 推荐 审核 查询提问 后台查询
	GetQuizByPage(ctx context.Context, args *ArgsQuizGetP, reply *ReplyQuizPage) error
	//添加提问
	AddQuiz(ctx context.Context, args *ArgsQuizAdd, reply *int) error
	//修改提问
	UpdateQuiz(ctx context.Context, args *ArgsQuizAdd, reply *bool) error
	//通过审核 或者 拒绝审核
	CheckQuiz(ctx context.Context, args *ArgsQuizCheck, reply *bool) error
	//按 分页 审核 查询答案 后台用
	GetAnswerByPage(ctx context.Context, args *ArgsAnswerGetP, reply *ReplyAnswerPage) error
	//通过审核 或者 拒绝审核 答案 or 评论
	CheckAnswer(ctx context.Context, args *ArgsAnswerCheck, reply *bool) error
	//按 分页 审核 查询评论
	GetCommentByPage(ctx context.Context, args *ArgsAnswerGetP, reply *ReplyCommentPage) error
	//添加 关键字
	AddKeyWord(ctx context.Context, args *ArgsKeyWordAdd, reply *int) error
	//修改 关键字
	UpdateKeyWord(ctx context.Context, args *ArgsKeyWordAdd, reply *bool) error
	//删除 关键字
	DelKeyWord(ctx context.Context, args *ArgsKeyWordDel, reply *bool) error
	//根据提问id查询详情
	GetQuizOne(ctx context.Context, args *int, reply *ReplyQuizOne) error
	//根据2级分类查询1级分类id
	GetCateBySubId(ctx context.Context, args *int, reply *int) error
	//根据二级分类id查询详情
	GetCateDetailById(ctx context.Context, args *int, reply *ReplyCateDetail) error

	//前后台 共用

	//按模块查询关键字
	GetKeyWordByModId(ctx context.Context, args *ArgsKeyWordGet, reply *ReplyKeyWord) error

	//前台

	//添加提问
	AddQuiz2(ctx context.Context, args *ArgsQuizAdd2, reply *int) error
	//根据用户Id查询 我的问答 数据
	GetQuizByUid(ctx context.Context, args *ArgsQuizGet, reply *ReplyQuiz) error
	//App端查询提问
	GetQuizByApp(ctx context.Context, args *ArgsQuizGetA, reply *ReplyQuizGetA) error
	//app端根据问题id查询答案和评论  //从问答列表 进入
	GetQuizDetail(ctx context.Context, args *ArgsQuizD, reply *[]ReplyInfo) error
	//app端根据问题id查询答案和评论  //从我的问答 问题id 查询详情
	GetQuizDetailB(ctx context.Context, args *ArgsQuizD, reply *ReplyQuizB) error
	//用户 添加 回复
	AddAnswer(ctx context.Context, args *ArgsAnswerAdd, reply *int) error
	//用户 收藏 提问
	CollectAnswer(ctx context.Context, args *ArgsAnswerCollect, reply *int) error
	//用户 取消 收藏提问
	CancelCollect(ctx context.Context, args *ArgsCollectCancel, reply *bool) error
	//用户点赞    包括 提问点赞 答案点赞
	AddGiveALike(ctx context.Context, args *ArgsGiveALike, reply *bool) error
}
