package access

import (
	"context"
	"ams-sentuh/internal/features/access/dto"
)

type AccessServiceInterface interface {
	RegisterAccess(ctx context.Context, request dto.AccessRegisterRequest) (*dto.AccessRegisterResponse, error)
	GetAllAccess(ctx context.Context) ([]dto.GetAllAccessResponse, error)
	GetAccess(ctx context.Context, id uint64) (dto.GetAllAccessResponse, error)
	UpdateAccess(ctx context.Context, id uint64, req dto.UpdateAccessRequest) error
	DeleteAccess(ctx context.Context, id uint64) error
}
