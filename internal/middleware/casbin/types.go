package casbin

type CasbinService interface {
	RegisterUser(user UserCasbin) error
	UpdateRole(user UserCasbin) error
	DeleteRole(user UserCasbin) error
	AddRole(user UserCasbin) error
	AddPolicy(role RolePolicy) error
	RemovePolicy(role RolePolicy) error
	UpdatePolicy(role RolePolicy) error
	UpdatePathForPolicies(oldPath, newPath, method string) error
}

type UserCasbin struct {
	ID       uint
	RoleId   uint
	LastRole uint
}

type PolicyPath struct {
	Path   string
	Method string
}

type RolePolicy struct {
	ID         uint
	Policy     []PolicyPath
	LastPolicy []PolicyPath
}
