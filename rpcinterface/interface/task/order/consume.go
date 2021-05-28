//消费相关任务入队列
//@author yangzhiwu<578154898@qq.com>
//@date 2020/8/17 10:32

package order

import "context"

type Consume interface {
	//将消费记录索引主表的主键id加入交换机
	SetLogRelationId(ctx context.Context, logRelationId *int, reply *bool ) error
}
