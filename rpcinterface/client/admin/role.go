package admin

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/admin"
)

type Role struct {
	client.Baseclient
}

func (r *Role) Init() *Role {
	r.ServiceName = "rpc_admin"
	r.ServicePath = "Role"
	return r
}

// AddCasbinRole 添加角色
func (r *Role) AddCasbinRole(ctx context.Context, args *admin.ArgsAddUpdateRole, reply *bool) error {
	return r.Call(ctx, "AddCasbinRole", args, reply)
}

// DeleteCasbinRoles 删除角色
func (r *Role) DeleteCasbinRoles(ctx context.Context, roleIds *[]int, reply *bool) error {
	return r.Call(ctx, "DeleteCasbinRoles", roleIds, reply)
}

// UpdateCasbinRole 更新角色
func (r *Role) UpdateCasbinRole(ctx context.Context, args *admin.ArgsAddUpdateRole, reply *bool) error {
	return r.Call(ctx, "UpdateCasbinRole", args, reply)
}

// GetCasbinRoleList 角色列表
func (r *Role) GetCasbinRoleList(ctx context.Context, args *admin.ArgsGetRoleList, reply *admin.ReplyGetRoleList) error {
	return r.Call(ctx, "GetCasbinRoleList", args, reply)
}

// GetCasbinRoleInfo 角色详情
func (r *Role) GetCasbinRoleInfo(ctx context.Context, roleId *int, reply *admin.RoleBase) error {
	return r.Call(ctx, "GetCasbinRoleInfo", roleId, reply)
}

// GetPolicyForRole 根据role ID 获取策略对应角色的策略
func (r *Role) GetPolicyForRole(ctx context.Context, roleId *int, reply *[]int) error {
	return r.Call(ctx, "GetPolicyForRole", roleId, reply)
}

// VerifyPermission 验证权限
func (r *Role) VerifyPermission(ctx context.Context, args *admin.ArgsVerifyPermission, reply *bool) error {
	return r.Call(ctx, "VerifyPermission", args, reply)
}

// GetCasbinRuleByRoleId 根据角色ID获取所有的权限
func (r *Role) GetCasbinRuleByRoleId(ctx context.Context, roleId *int, reply *admin.ReplyGetCasbinRule) error {
	return r.Call(ctx, "GetCasbinRuleByRoleId", roleId, reply)
}

// GetRoleMemus GetRoleMemus
func (r *Role) GetRoleMemus(ctx context.Context, args *admin.ArgsGetRoleMemus, reply *admin.ReplyGetRoleMemus) error {
	return r.Call(ctx, "GetRoleMemus", args, reply)
}

// GetRuleNamesByRoleId 根据权限id获取该角色下所有的权限名
func (r *Role) GetRuleNamesByRoleId(ctx context.Context, roleId *int, reply *[]string) error {
	return r.Call(ctx, "GetRuleNamesByRoleId", roleId, reply)
}
