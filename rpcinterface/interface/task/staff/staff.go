/******************************************
@Description:
@Time : 2020/11/13 11:14
@Author :lixiaojun

*******************************************/
package staff

import "context"

type ArgsSetAddLeaveLeDeleteStaff struct {
	StaffId int
	Status int //员工状态：1-新增；2-离职；3-删除
}
type Staff interface {
	//员工 新增/离职/删除
	SetAddLeaveLeDeleteStaff(ctx context.Context,args *ArgsSetAddLeaveLeDeleteStaff,reply *bool)error
}
