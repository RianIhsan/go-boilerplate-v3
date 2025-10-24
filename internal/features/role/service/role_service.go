package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"ams-sentuh/config"
	"ams-sentuh/internal/entities"
	"ams-sentuh/internal/features/role"
	"ams-sentuh/internal/features/role/dto"
	"ams-sentuh/internal/middleware/casbin"
)

type roleService struct {
	cfg           *config.Config
	roleRepo      role.RoleRepositoryInterface
	casbinService casbin.CasbinService
}

func NewRoleService(cfg *ServiceConfig) role.RoleServiceInterface {
	return &roleService{
		cfg:           cfg.Config,
		roleRepo:      cfg.RoleRepoInterface,
		casbinService: cfg.Casbin,
	}
}

func (r *roleService) AddRole(ctx context.Context, req dto.RegisterRoleRequest) (dto.RegisterRoleResponse, error) {
	createdRole, err := r.roleRepo.Create(ctx, dto.ConvertToEntityRoleRequest(req))
	if err != nil {
		return dto.RegisterRoleResponse{}, err
	}

	casbinData := casbin.RolePolicy{
		ID: uint(createdRole.ID),
	}

	for _, permission := range req.Permissions {
		obj := casbin.PolicyPath{
			Path:   permission.Path,
			Method: permission.Method,
		}
		casbinData.Policy = append(casbinData.Policy, obj)
	}

	err = r.casbinService.AddPolicy(casbinData)
	if err != nil {
		return dto.RegisterRoleResponse{}, err
	}
	return dto.ConvertToRegisterRoleResponse(createdRole), nil
}

func (r *roleService) GetAll(ctx context.Context) ([]dto.RoleResponse, error) {
	roles, err := r.roleRepo.GetAll(ctx)
	if err != nil {
		return []dto.RoleResponse{}, err
	}

	return dto.ConvertToGetListRoleResponse(roles), nil
}

func (r *roleService) ModifyRolePermission(ctx context.Context, data dto.RolePermissions) (entities.Role, error) {
	if data.RoleId == 0 {
		return entities.Role{}, errors.New("invalid role id")
	}

	var toCreate []entities.RolePermission
	var toDelete []uint64

	casbinData := casbin.RolePolicy{
		ID: data.RoleId,
	}

	var permissionsIds []uint64
	for _, v := range data.ListAction {
		permissionsIds = append(permissionsIds, v.ID)
	}

	permissions, err := r.roleRepo.BatchGetPermissions(ctx, permissionsIds)
	if err != nil {
		return entities.Role{}, err
	}

	existingPermissions, err := r.roleRepo.GetGroupedPermissions(ctx, uint64(data.RoleId))
	if err != nil {
		return entities.Role{}, err
	}

	permissionsMap := make(map[uint64]entities.Permission)
	for _, v := range permissions {
		permissionsMap[v.ID] = v
	}

	existingPermissionsMap := make(map[uint]bool)
	for _, p := range existingPermissions {
		existingPermissionsMap[uint(p.PermissionID)] = true
	}

	for _, v := range data.ListAction {
		permission, exists := permissionsMap[v.ID]
		if !exists {
			return entities.Role{}, fmt.Errorf("permission not found for id: %d", v.ID)
		}

		objCasbin := casbin.PolicyPath{
			Path:   permission.Path,
			Method: permission.Method,
		}
		casbinData.Policy = append(casbinData.Policy, objCasbin)
		casbinData.LastPolicy = append(casbinData.LastPolicy, objCasbin)

		if v.Action == 1 {
			toDelete = append(toDelete, v.ID)
		} else if v.Action == 0 {
			if existingPermissionsMap[uint(v.ID)] {
				continue
			}
			toCreate = append(toCreate, entities.RolePermission{
				RoleID:       uint64(data.RoleId),
				PermissionID: v.ID,
				AccessID:     permission.AccessID,
			})
		}
	}

	if len(toDelete) > 0 {
		if err := r.roleRepo.DeleteRolePermission(ctx, uint64(data.RoleId), toDelete); err != nil {
			return entities.Role{}, err
		}
		if err := r.casbinService.RemovePolicy(casbinData); err != nil {
			return entities.Role{}, err
		}
	}

	if len(toCreate) > 0 {
		if err := r.roleRepo.CreateRolePermission(ctx, toCreate); err != nil {
			return entities.Role{}, err
		}
		if err := r.casbinService.AddPolicy(casbinData); err != nil {
			return entities.Role{}, err
		}
	}

	groupedPerms, err := r.roleRepo.GetGroupedPermissions(ctx, uint64(data.RoleId))
	if err != nil {
		return entities.Role{}, err
	}

	existingAccess, err := r.roleRepo.GetRoleAccess(ctx, uint64(data.RoleId))
	if err != nil {
		return entities.Role{}, err
	}

	permAccessMap := map[uint64]bool{}
	roleAccessMap := map[uint64]entities.RoleAccess{}

	for _, p := range groupedPerms {
		permAccessMap[*p.AccessID] = true
	}
	for _, r := range existingAccess {
		roleAccessMap[r.AccessID] = r
	}

	var toCreateAccess []entities.RoleAccess
	var toDeleteAccess []uint64

	for accessID := range permAccessMap {
		if _, exists := roleAccessMap[accessID]; !exists {
			toCreateAccess = append(toCreateAccess, entities.RoleAccess{
				RoleID:   uint64(data.RoleId),
				AccessID: accessID,
			})
		}
	}

	for accessID := range roleAccessMap {
		if _, exists := permAccessMap[accessID]; !exists {
			toDeleteAccess = append(toDeleteAccess, accessID)
		}
	}

	if len(toCreateAccess) > 0 {
		if err := r.roleRepo.CreateRoleAccess(ctx, toCreateAccess); err != nil {
			return entities.Role{}, err
		}
	}
	if len(toDeleteAccess) > 0 {
		if err := r.roleRepo.DeleteRoleAccess(ctx, uint64(data.RoleId), toDeleteAccess); err != nil {
			return entities.Role{}, err
		}
	}

	roles, err := r.roleRepo.GetByID(ctx, data.RoleId)
	if err != nil {
		return entities.Role{}, err
	}
	return *roles, nil
}

func (r *roleService) GetByID(ctx context.Context, roleId uint64) (dto.RoleResponse, error) {
	role, err := r.roleRepo.GetByID(ctx, uint(roleId))
	if err != nil {
		return dto.RoleResponse{}, errors.New("role not found")
	}
	return dto.ConvertToRoleResponse(*role), nil
}

func (r *roleService) UpdateRole(ctx context.Context, id uint64, req dto.UpdateRoleRequest) error {
	_, err := r.roleRepo.GetByID(ctx, uint(id))
	if err != nil {
		return errors.New("role not found")
	}
	err = r.roleRepo.UpdateRole(ctx, &entities.Role{
		ID:   id,
		Name: req.Name,
	})
	if err != nil {
		return errors.New("failed to update role")
	}
	return nil
}

func (r *roleService) DeleteRole(ctx context.Context, roleId uint64) error {
	_, err := r.roleRepo.GetByID(ctx, uint(roleId))
	if err != nil {
		return errors.New("role not found")
	}
	err = r.roleRepo.DeleteRole(ctx, roleId)
	if err != nil {
		return errors.New("failed to delete role")
	}
	return nil
}
