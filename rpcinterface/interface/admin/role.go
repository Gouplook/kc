package admin

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/interface/common"
)

type ArgsAddUpdateRole struct {
	RoleBase
	RuleIds []int // 权限ID
}

type ArgsGetRoleList struct {
	RoleIds []uint // 角色 Id
	common.Paging
}

type ReplyGetRoleList struct {
	TotalNum int
	List     []RoleBase
}

type RoleBase struct {
	ID   int
	Name string
	Info string
}

type ArgsVerifyPermission struct {
	RoleId int
	Path   string
}

type CasbinRuleBase struct {
	ID       int
	RuleName string `mapstructure:"rule_name"`
	RuleDesc string `mapstructure:"rule_desc"`
	ParentId int    `mapstructure:"parent_id"`
	Type     int    `mapstructure:"type"` //  0:功能;1:菜单
	Enable   bool   // 权限是否启用
}

type ReplyGetCasbinRule struct {
	CasbinRule map[int]CasbinRuleBase
}

type ArgsGetRoleMemus struct {
	RoleId    int // 0: 全部数据
	ParentId  int
	Type      int // 0:功能/1:菜单
	FuncLimit bool
}
type ReplyGetRoleMemus struct {
	List     []CasbinRuleBase
	SubMemus []CasbinRuleBase
	//SubFuncs []CasbinRuleBase
	//SubMemu map[int][]CasbinRuleBase // 菜单
	SubGongN map[int][]CasbinRuleBase // 功能
}

type Role interface {
	// AddCasbinRole 添加角色
	AddCasbinRole(ctx context.Context, args *ArgsAddUpdateRole, reply *bool) error
	// DeleteCasbinRoles 删除角色
	DeleteCasbinRoles(ctx context.Context, roleIds *[]int, reply *bool) error
	// UpdateCasbinRole 更新角色
	UpdateCasbinRole(ctx context.Context, args *ArgsAddUpdateRole, reply *bool) error
	// GetCasbinRoleList 角色列表
	GetCasbinRoleList(ctx context.Context, args *ArgsGetRoleList, reply *ReplyGetRoleList) error
	// GetCasbinRoleInfo 角色详情
	GetCasbinRoleInfo(ctx context.Context, roleId *int, reply *[]RoleBase)
	// GetPolicyForRole 根据role ID 获取策略对应角色的策略
	GetPolicyForRole(ctx context.Context, roleId *int, reply *[]int) error
	// VerifyPermission 验证权限
	VerifyPermission(ctx context.Context, args *ArgsVerifyPermission, reply *bool) error
	// GetCasbinRuleByRoleId 根据角色ID获取所有的权限
	GetCasbinRuleByRoleId(ctx context.Context, roleId *int, reply *ReplyGetCasbinRule) error
	// GetRoleMemus GetRoleMemus
	GetRoleMemus(ctx context.Context, args *ArgsGetRoleMemus, reply *ReplyGetRoleMemus) error
	// GetRuleNamesByRoleId 根据权限id获取该角色下所有的权限名
	GetRuleNamesByRoleId(ctx context.Context, roleId *int, reply *[]string) error
}
