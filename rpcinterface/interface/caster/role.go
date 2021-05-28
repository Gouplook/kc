package caster

import (
	"context"
)

type Empty struct {

}

type AddRoleReq struct{
	Name string
	Info string
}

type RoleID struct {
	ID int
}

type RoleStruct struct {
	ID int `mapstructure:"id"`
	Name string `mapstructure:"name"`
	Info string `mapstructure:"info"`
}

type Page struct {
	Page int
	PageSize int
}

type RoleList struct {
	List []RoleStruct
}

type Role interface {
	AddRole(ctx context.Context,  args *AddRoleReq, reply *RoleID)(err error)
	DelRole(ctx context.Context, args *RoleID, reply *Empty)(err error)
	UpdateRole(ctx context.Context, args *RoleStruct, reply *Empty)(err error)
	RoleInfo(ctx context.Context, args *RoleID, reply *RoleStruct)(err error)
	RoleList(ctx context.Context, args *Page, reply *RoleList)(err error)
}
