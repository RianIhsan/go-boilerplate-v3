package service

import (
	"context"
	"errors"
	"ams-sentuh/config"
	"ams-sentuh/internal/entities"
	"ams-sentuh/internal/features/permission"
	"ams-sentuh/internal/features/permission/dto"
)

type permissionService struct {
	cfg            *config.Config
	permissionRepo permission.PermissionRepositoryInterface
}

func NewPermissionService(cfg *ServiceConfig) permission.PermissionServiceInterface {
	return &permissionService{
		cfg:            cfg.Config,
		permissionRepo: cfg.PermissionRepoInterface,
	}
}

func (p permissionService) AddPermission(ctx context.Context, register dto.PermissionRegister) (dto.Permission, error) {
	entity := entities.Permission{
		Name:     register.Name,
		Path:     register.Path,
		Method:   register.Method,
		AccessID: &register.AccessId,
		Type:     register.Type,
	}

	created, err := p.permissionRepo.Create(ctx, entity)
	if err != nil {
		return dto.Permission{}, err
	}

	return dto.Permission{
		ID:        created.ID,
		Name:      created.Name,
		Path:      created.Path,
		Method:    created.Method,
		AccessId:  *created.AccessID,
		Type:      created.Type,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
	}, nil
}

func (p permissionService) GetListPermission(ctx context.Context, accessID *uint64) ([]dto.Permission, error) {
	permissions, err := p.permissionRepo.GetPermissions(ctx, accessID)
	if err != nil {
		return nil, err
	}

	var results []dto.Permission
	for _, perm := range permissions {
		results = append(results, dto.Permission{
			ID:        perm.ID,
			Name:      perm.Name,
			Path:      perm.Path,
			Method:    perm.Method,
			AccessId:  *perm.AccessID,
			Type:      perm.Type,
			CreatedAt: perm.CreatedAt,
			UpdatedAt: perm.UpdatedAt,
		})
	}

	return results, nil
}

func (p permissionService) GetPermission(ctx context.Context, id uint64) (dto.Permission, error) {
	perm, err := p.permissionRepo.GetPermission(ctx, id)
	if err != nil {
		return dto.Permission{}, err
	}

	return dto.Permission{
		ID:       perm.ID,
		Name:     perm.Name,
		Path:     perm.Path,
		Method:   perm.Method,
		AccessId: *perm.AccessID,
		Type:     perm.Type,
	}, nil
}

func (p permissionService) UpdatePermission(ctx context.Context, id uint64, update dto.PermissionUpdate) error {
	_, err := p.permissionRepo.UpdatePermission(ctx, entities.Permission{
		ID:       id,
		Name:     update.Name,
		Path:     update.Path,
		Method:   update.Method,
		AccessID: &update.AccessId,
		Type:     update.Type,
	})
	if err != nil {
		return errors.New("failed to update permission: " + err.Error())
	}
	return nil
}

func (p permissionService) DeletePermission(ctx context.Context, id uint64) error {
	err := p.permissionRepo.DeletePermission(ctx, id)
	if err != nil {
		return errors.New("failed to delete permission: " + err.Error())
	}
	return nil
}
