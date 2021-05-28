//单项目标签表
//@author yangzhiwu<578154898@qq.com>
//@date 2020-03-27 09:37

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
)

//表结构体
type TagModel struct {
	Model *base.Model
	Field TagModelField
}

//表字段
type TagModelField struct {
	T_table  string `default:"tag"`
	F_tag_id string `default:"tag_id"`
	F_bus_id string `default:"bus_id"`
	F_name   string `default:"name"`
	F_is_del string `default:"id_del"`
	F_ctime  string `default:"ctime"`
}

//是否被删除 0=否 1=是
const (
	TAG_IS_DEL_no  = 0
	TAG_IS_DEL_yes = 1
)

//初始化
func (t *TagModel) Init() *TagModel {
	functions.ReflectModel(&t.Field)
	t.Model = base.NewModel(t.Field.T_table)
	return t
}

//新增数据
func (t *TagModel) Insert(data map[string]interface{}) int {
	result, _ := t.Model.Data(data).Insert()
	return result
}

//根据tagids获取标签信息
//@param tagIds []int 标签id数组
//@return map[string]interface{}
func (t *TagModel) GetByTagids(tagIds []int) []map[string]interface{} {
	if len(tagIds) <= 0 {
		return []map[string]interface{}{}
	}
	r := t.Model.Where(map[string]interface{}{
		t.Field.F_tag_id: []interface{}{"IN", tagIds},
	}).Select()

	return r
}

//获取企业/商户下面的未被删除的所有标签
func (t *TagModel) GetByBusid(busId int) []map[string]interface{} {
	if busId <= 0 {
		return []map[string]interface{}{}
	}
	maxNum, _ := kcgin.AppConfig.Int("tag.maxNum")
	if maxNum == 0 {
		maxNum = 100 //没有取到默认100
	}
	return t.Model.Where(map[string]interface{}{
		t.Field.F_bus_id: busId,
		t.Field.F_is_del: TAG_IS_DEL_no,
	}).Limit(0, maxNum).OrderBy(t.Field.F_tag_id + " DESC ").Select()
}

//获取商家标签数量
//@param busId int 商家id
//@return int 数量
func (t *TagModel) GetBusTagNum(busId int) int {
	if busId <= 0 {
		return 0
	}

	return t.Model.Where(map[string]interface{}{
		t.Field.F_bus_id: busId,
		t.Field.F_is_del: TAG_IS_DEL_no,
	}).Count(t.Field.F_tag_id)
}

//根据tagid删除标签
//@param tagId int 标签id
//return true|false
func (t *TagModel) DelByTagid(tagId int) bool {
	if tagId <= 0 {
		return false
	}

	_, err := t.Model.Where(map[string]interface{}{
		t.Field.F_tag_id: tagId,
	}).Data(map[string]interface{}{
		t.Field.F_is_del: cards.TAG_IS_DEL_yes,
	}).Update()
	if err != nil {
		return false
	}

	return true
}

//修改标签数据
//@param tagId int 标签id
//@param data map[string]interface{} 要修改的数据
//return true|false
func (t *TagModel) UpdateByTagid(tagId int, data map[string]interface{}) bool {
	if tagId <= 0 || len(data) == 0 {
		return false
	}

	_, err := t.Model.Where(map[string]interface{}{
		t.Field.F_tag_id: tagId,
	}).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}
