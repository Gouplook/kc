package poplar

import (
    "context"
    "git.900sui.cn/kc/rpcinterface/client"
    "git.900sui.cn/kc/rpcinterface/interface/poplar"
)

type Student struct {
    client.Baseclient
}

func (student *Student)Init() *Student {
    student.ServiceName = "rpc_poplar"
    student.ServicePath = "Student"
    return student
}

func (student *Student)GetStudentByName(ctx context.Context, args *poplar.ArgsGetStudentByName, reply *poplar.ReplyStudent) error {
    return student.Call(ctx, "GetStudentByName", args, reply)
}
func (student *Student)GetStudentsByAge(ctx context.Context, args *poplar.ArgsGetStudentsByAge, reply *[]poplar.ReplyStudent) error {
    return student.Call(ctx, "GetStudentsByAge", args, reply)
}
func (student *Student)GetStudents(ctx context.Context, args *poplar.ArgsGetStudents, reply *[]poplar.ReplyStudent) error {
    return student.Call(ctx, "GetStudents", args, reply)
}
