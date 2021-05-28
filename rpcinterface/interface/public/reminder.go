//不同行业的温馨提示
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/3 18:43
package public

import "context"

const (
	Type_checkbox = 1 // 复选
	Type_select   = 2 // 单选
	Type_input    = 3 // 输入

	Require_true  = 1 // 必选
	Require_false = 0 // 非必选
)

type ReminderInfo struct {
	Name    string   // 名称
	Depict  string   // 描述
	Type    int      // 类型
	Require int      // 是否必传 1=必传 0=非必传
	Key     string   // 键名
	Value   []string // 数据
}

type Reminder interface {
	// 获取行业的温馨提示
	GetReminderInfo(ctx context.Context, indusId *int, reply *[]ReminderInfo)
}
