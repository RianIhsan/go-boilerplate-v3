package repository

import (
	"context"
	"fmt"
	"ams-sentuh/internal/entities"
	"ams-sentuh/internal/features/user"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) user.UserRepositoryInterface {
	return &userPostgresRepository{db: db}
}

func (u *userPostgresRepository) Create(ctx context.Context, entity entities.User) (entities.User, error) {
	if err := u.db.WithContext(ctx).Create(&entity).Error; err != nil {
		return entities.User{}, err
	}
	return entity, nil
}

func (u *userPostgresRepository) FindByEmail(ctx context.Context, entity entities.User) (*entities.User, error) {
	user := new(entities.User)
	DB := u.db.WithContext(ctx)

	if err := DB.Where(entities.User{Email: entity.Email}).Take(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userPostgresRepository) GetList(ctx context.Context, roleId uint64) ([]entities.User, error) {
	var users []entities.User
	query := u.db.WithContext(ctx).Model(&entities.User{}).
		Preload("Role").
		Preload("Branch").
		Preload("Company")

	if roleId != 0 {
		query = query.Where("role_id = ?", roleId)
	}

	if err := query.Order("created_at ASC").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userPostgresRepository) FindById(ctx context.Context, userId uint64) (entities.User, error) {
	var user entities.User
	DB := u.db.WithContext(ctx).
		Preload("Role").
		Preload("Branch").
		Preload("Company")

	if err := DB.Where("id = ?", userId).Take(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (u *userPostgresRepository) Update(ctx context.Context, id uint64, data entities.User) error {
	var existingUser entities.User
	if err := u.db.WithContext(ctx).First(&existingUser, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with id %d not found", id)
		}
		return fmt.Errorf("cannot update user with id %d: %w", id, err)
	}
	if err := u.db.WithContext(ctx).Model(&existingUser).Updates(data).Error; err != nil {
		return fmt.Errorf("cannot update user with id %d: %w", id, err)
	}
	return nil
}

func (u *userPostgresRepository) DeleteUser(ctx context.Context, userId uint64) error {
	if err := u.db.WithContext(ctx).Delete(&entities.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}

// func (u *userPostgresRepository) InsertOTP(ctx context.Context, data entities.OTP) error {
// 	if err := u.db.WithContext(ctx).Create(&data).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (u *userPostgresRepository) GetOTPByEmail(ctx context.Context, email string) (*entities.OTP, error) {
// 	var otp entities.OTP
// 	err := u.db.WithContext(ctx).
// 		Where("email = ?", email).
// 		Order("created_at DESC").
// 		First(&otp).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &otp, nil
// }

// func (u *userPostgresRepository) UpdateResetToken(ctx context.Context, email, token string) error {
// 	return u.db.WithContext(ctx).
// 		Model(&entities.OTP{}).
// 		Where("email = ?", email).
// 		Order("created_at DESC").
// 		Limit(1).
// 		Update("token", token).Error
// }

// func (u *userPostgresRepository) GetByResetToken(ctx context.Context, token string) (*entities.OTP, error) {
// 	var otp entities.OTP
// 	err := u.db.WithContext(ctx).
// 		Where("token = ?", token).
// 		First(&otp).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	//if time.Now().Unix() > int64(otp.ExpiredAt) {
// 	//	return nil, errors.New("reset token expired")
// 	//}

// 	return &otp, nil
// }

func (u *userPostgresRepository) UpdatePassword(ctx context.Context, email, hashedPassword string) error {
	return u.db.WithContext(ctx).
		Model(&entities.User{}).
		Where("email = ?", email).
		Update("password", hashedPassword).Error
}

// func (u *userPostgresRepository) DeleteResetToken(ctx context.Context, token string) error {
// 	return u.db.WithContext(ctx).
// 		Where("token = ?", token).
// 		Delete(&entities.OTP{}).Error
// }

// func (u *userPostgresRepository) DeleteOTPByEmail(ctx context.Context, email string) error {
// 	return u.db.WithContext(ctx).
// 		Where("email = ?", email).
// 		Delete(&entities.OTP{}).Error
// }
