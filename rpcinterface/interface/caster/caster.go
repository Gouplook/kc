package caster

import (
	"context"
)

type AddNodeReq struct{
	RoleID int
	ParentID int
	Enable bool
	Visible bool
	Sign string
	Name string
}

type  AddNodeResp struct{
	ID int
}

type UpdateReq struct {
	ID int
	RoleID int
	Sign string
	Name string
}

type NodeID struct{
	ID int
	RoleID int
}

type Node struct{
	ID int `mapstructure:"id"`
	RoleID int `mapstructure:"role_id"`
	ParentID int `mapstructure:"parent_id"`
	Enable bool `mapstructure:"enable"`
	Sign string `mapstructure:"sign"`
	Name string `mapstructure:"name"`
	Visible bool `mapstructure:"visible"`
}

type SubListReq struct {
	RoleID int
	ParentID int
}
type SubListResp struct {
	List []Node
}

type VerifyReq struct{
	RoleID int
	Path string
}

type VerifyResp struct{
	Result bool
}

type Caster interface {
	Verify(ctx context.Context, args *VerifyReq, reply *VerifyResp)(err error)
	AddNode(ctx context.Context, args *AddNodeReq, reply *AddNodeResp)(err error)
	DelNode(ctx context.Context, args *NodeID, reply *Empty)(err error)
	EnableNode(ctx context.Context,  args *NodeID, resp *Empty) (err  error)
	DisableNode(ctx context.Context,  args *NodeID, resp *Empty) (err  error)
	UpdateNode(ctx context.Context, args *UpdateReq, resp *Empty)(err error)
	SubList(ctx context.Context, args *SubListReq, reply *SubListResp )(err error)
	SubEnableList(ctx context.Context, args *SubListReq, reply *SubListResp )(err error)
}