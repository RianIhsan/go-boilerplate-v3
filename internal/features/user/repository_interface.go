package user

import (
	"context"
	"ams-sentuh/internal/entities"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user entities.User) (entities.User, error)
	FindByEmail(ctx context.Context, user entities.User) (*entities.User, error)
	GetList(ctx context.Context, roleId uint64) ([]entities.User, error)
	FindById(ctx context.Context, userId uint64) (entities.User, error)
	Update(ctx context.Context, id uint64, data entities.User) error
	// InsertOTP(ctx context.Context, data entities.OTP) error
	// GetOTPByEmail(ctx context.Context, email string) (*entities.OTP, error)
	// UpdateResetToken(ctx context.Context, email, token string) error
	DeleteUser(ctx context.Context, userId uint64) error

	// GetByResetToken(ctx context.Context, token string) (*entities.OTP, error)
	// UpdatePassword(ctx context.Context, email, hashedPassword string) error
	// DeleteResetToken(ctx context.Context, token string) error
	// DeleteOTPByEmail(ctx context.Context, email string) error
}
