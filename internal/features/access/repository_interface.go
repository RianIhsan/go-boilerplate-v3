package access

import (
	"context"
	"ams-sentuh/internal/entities"
)

type AccessRepositoryInterface interface {
	Create(ctx context.Context, model entities.Access) (entities.Access, error)
	GetAll(ctx context.Context) ([]entities.Access, error)
	GetByID(ctx context.Context, id uint64) (entities.Access, error)
	Update(ctx context.Context, model entities.Access) (entities.Access, error)
	Delete(ctx context.Context, id uint64) error
}
