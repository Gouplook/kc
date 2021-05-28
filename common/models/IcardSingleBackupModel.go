//IcardSingleBackupModel
//2021-04-07 13:31:52

package models

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/base/common/models/base"
	"git.900sui.cn/kc/kcgin/orm"
	"strconv"
)

//表结构体
type IcardSingleBackupModel struct {
	Model *base.Model
	Field IcardSingleBackupModelField
}

//表字段
type IcardSingleBackupModelField struct {
	T_table      string `default:"icard_single_backup"`
	F_id         string `default:"id"`
	F_icard_id   string `default:"icard_id"`
	F_is_sync    string `default:"is_sync"`
	F_discount   string `default:"discount"`
	F_single_id  string `default:"single_id"`
	F_backup_num string `default:"backup_num"`
}

//初始化
func (i *IcardSingleBackupModel) Init(ormer ...orm.Ormer) *IcardSingleBackupModel {
	functions.ReflectModel(&i.Field)
	i.Model = base.NewModel(i.Field.T_table, ormer...)
	return i
}

//新增数据
func (i *IcardSingleBackupModel) Insert(data map[string]interface{}) int {
	result, _ := i.Model.Data(data).Insert()
	return result
}

//InsertAll 批量新增数据
func (i *IcardSingleBackupModel) InsertAll(data []map[string]interface{}) int {
	result, _ := i.Model.InsertAll(data)
	return result
}
func (i *IcardSingleBackupModel) FindByIcardIds(icardIds []int, backupNum ...int) []map[string]interface{} {
	if len(icardIds) == 0 {
		return make([]map[string]interface{}, 0)
	}
	where := []base.WhereItem{
		{i.Field.F_icard_id, []interface{}{"IN", icardIds}},
	}
	if len(backupNum) > 0 {
		where = append(where, base.WhereItem{i.Field.F_backup_num, backupNum[0]})
	}
	return i.Model.Where(where).Select()
}

func (i *IcardSingleBackupModel) DeleteByIcardIds(icardIds []int) bool {
	if len(icardIds) == 0 {
		return false
	}
	_, err := i.Model.Where([]base.WhereItem{
		{i.Field.F_icard_id, []interface{}{"IN", icardIds}},
	}).Delete()
	if err != nil {
		return false
	}
	return true
}

//查询最后一条记录的“备份次数”
//isAddOne:插入数据前查询需要传true,其它情况false
//isSync:插入数据前传入”“，其它情况按需传入
func (i *IcardSingleBackupModel) GetLastBackUumByIcardId(isAddOne bool, isSync string, icardId ...int) int {
	backupNum := 1
	where := make([]base.WhereItem, 0)
	if len(icardId) > 0 {
		where = append(where, base.WhereItem{
			Field: i.Field.F_icard_id,
			Value: icardId,
		})
	}
	if len(isSync) > 0 {
		where = append(where, base.WhereItem{
			Field: i.Field.F_is_sync,
			Value: isSync,
		})
	}
	lastInfo := i.Model.Where(where).OrderBy(i.Field.F_id+" desc ").Limit(0, 1).Select()

	if len(lastInfo) > 0 {
		num, _ := strconv.Atoi(lastInfo[0][i.Field.F_backup_num].(string))
		if isAddOne {
			backupNum += num
		}else{
			backupNum = num
		}
	}
	return backupNum
}
