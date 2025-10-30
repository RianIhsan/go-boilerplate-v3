package service

import (
	"context"
	"github.com/pkg/errors"
	"ams-sentuh/config"
	"ams-sentuh/internal/entities"
	"ams-sentuh/internal/features/access"
	"ams-sentuh/internal/features/access/dto"
)

type accessService struct {
	cfg        *config.Config
	accessRepo access.AccessRepositoryInterface
}

func NewAccessService(cfg *ServiceConfig) access.AccessServiceInterface {
	return &accessService{
		cfg:        cfg.Config,
		accessRepo: cfg.AccessRepoInterface,
	}
}

func (as accessService) RegisterAccess(ctx context.Context, req dto.AccessRegisterRequest) (*dto.AccessRegisterResponse, error) {
	createdAccess, err := as.accessRepo.Create(ctx, dto.ConvertToAccessEntity(req))
	if err != nil {
		return nil, errors.New("failed to create access: " + err.Error())
	}

	return dto.ConvertToAccessRegisterResponse(createdAccess), nil
}

func (as accessService) GetAllAccess(ctx context.Context) ([]dto.GetAllAccessResponse, error) {
	dataAccess, err := as.accessRepo.GetAll(ctx)
	if err != nil {
		return nil, errors.New("failed to get all access")
	}

	return dto.ConvertToGetAllAccessResponseList(dataAccess), nil
}

func (as accessService) GetAccess(ctx context.Context, id uint64) (dto.GetAllAccessResponse, error) {
	data, err := as.accessRepo.GetByID(ctx, id)
	if err != nil {
		return dto.GetAllAccessResponse{}, errors.New("failed to get access")
	}
	return dto.ConvertToGetAllAccessResponse(data), nil
}

func (as accessService) UpdateAccess(ctx context.Context, id uint64, req dto.UpdateAccessRequest) error {
	_, err := as.accessRepo.Update(ctx, entities.Access{
		ID:       id,
		Name:     req.Name,
		Link:     req.Link,
		Priority: req.Priority,
	})
	if err != nil {
		return errors.New("failed to update access: " + err.Error())
	}

	return nil
}

func (as accessService) DeleteAccess(ctx context.Context, id uint64) error {
	err := as.accessRepo.Delete(ctx, id)
	if err != nil {
		return errors.New("failed to delete access: " + err.Error())
	}
	return nil
}
