package service

import (
	"context"
	"ams-sentuh/config"
	"ams-sentuh/internal/entities"
	"ams-sentuh/internal/features/user"
	"ams-sentuh/internal/features/user/dto"
	"ams-sentuh/internal/middleware/casbin"
	// "ams-sentuh/pkg/broker"
	"ams-sentuh/pkg/utils"

	"github.com/pkg/errors"
)

type userService struct {
	cfg           *config.Config
	userRepo      user.UserRepositoryInterface
	casbinService casbin.CasbinService
}

func NewUserService(cfg *ServiceConfig) user.UserServiceInterface {
	return &userService{
		cfg:           cfg.Config,
		userRepo:      cfg.UserRepoInterface,
		casbinService: cfg.Casbin,
	}
}

func (uS *userService) AddUser(ctx context.Context, req dto.RegisterUserRequest) (dto.RegisterUserResponse, error) {
	result, err := uS.userRepo.FindByEmail(ctx, entities.User{Email: req.Email})
	if result != nil && err == nil {
		return dto.RegisterUserResponse{}, errors.New("Email already exist")
	}

	if err := req.PrepareCreate(); err != nil {
		return dto.RegisterUserResponse{}, errors.Wrap(err, "failed to prepare user data")
	}

	createdUser, err := uS.userRepo.Create(ctx, dto.ConvertToEntityUserRequest(req))
	if err != nil {
		return dto.RegisterUserResponse{}, errors.Wrap(err, "failed to create user")
	}

	err = uS.casbinService.RegisterUser(casbin.UserCasbin{
		ID:     uint(createdUser.ID),
		RoleId: uint(createdUser.RoleID),
	})
	if err != nil {
		return dto.RegisterUserResponse{}, errors.Wrap(err, "failed to register user in casbin")
	}

	return dto.ConvertToRegisterUserResponse(&createdUser), nil
}

func (uS *userService) LoginUser(ctx context.Context, request *dto.LoginUserRequest) (*dto.JwtToken, error) {
	foundUser, err := uS.userRepo.FindByEmail(ctx, entities.User{Email: request.Email})
	if err != nil {
		return nil, errors.Wrap(err, "email not found")
	}

	if err := request.ComparePassword(foundUser.Password); err != nil {
		return nil, errors.New("invalid password")
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(foundUser, uS.cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate tokens")
	}

	return &dto.JwtToken{
		Email:        foundUser.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		RoleID:       foundUser.RoleID,
	}, nil
}

func (uS *userService) GetList(ctx context.Context, roleId uint64) ([]dto.UserDTO, error) {
	data, err := uS.userRepo.GetList(ctx, roleId)
	if err != nil {
		return []dto.UserDTO{}, errors.Wrap(err, "failed to list users")
	}
	return dto.ToListUsersResponse(data), nil
}

func (uS *userService) GetById(ctx context.Context, userId uint64) (dto.UserDTO, error) {
	fetchUser, err := uS.userRepo.FindById(ctx, userId)
	if err != nil {
		return dto.UserDTO{}, errors.Wrap(err, "failed to find user")
	}

	return dto.ToUserDTO(fetchUser), nil
}

func (uS *userService) Update(ctx context.Context, id uint64, data dto.UpdateUserRequest) error {
	err := uS.userRepo.Update(ctx, id, entities.User{
		Name:     data.Name,
		Email:    data.Email,
		RoleID:   data.RoleId,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func (uS *userService) Delete(ctx context.Context, userId uint64) error {
	userEntity, err := uS.userRepo.FindById(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to find user")
	}

	if err := uS.userRepo.DeleteUser(ctx, userId); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	err = uS.casbinService.DeleteRole(casbin.UserCasbin{
		ID:       uint(userEntity.ID),
		LastRole: uint(userEntity.RoleID),
	})
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}
	return nil
}

// func (uS *userService) ForgotPassword(ctx context.Context, email string) error {
// 	getUser, err := uS.userRepo.FindByEmail(ctx, entities.User{Email: email})
// 	if err != nil {
// 		return errors.Wrap(err, "user not found")
// 	}

// 	otp := utils.GenerateOTP(6)
// 	otpInt, _ := strconv.Atoi(otp)
// 	expiredAt3Min := time.Now().Add(time.Minute * 3).Unix()

// 	err = uS.userRepo.InsertOTP(ctx, entities.OTP{
// 		Email:     email,
// 		OTP:       otpInt,
// 		ExpiredAt: int(expiredAt3Min),
// 	})

// 	if err != nil {
// 		return errors.Wrap(err, "failed to insert otp")
// 	}

// 	msg := map[string]any{
// 		"email":    getUser.Email,
// 		"username": getUser.Username,
// 		"otp":      otpInt,
// 	}

// 	payload, err := json.Marshal(msg)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to marshal payload")
// 	}

// 	// publish event ke NATS
// 	if err := broker.Publish("user.forgot_password", payload); err != nil {
// 		return errors.Wrap(err, "failed to publish forgot_password event")
// 	}

// 	return nil

// 	return nil
// }

// func (uS *userService) VerifyOTP(ctx context.Context, email string, otp int) (string, error) {
// 	storedOtp, err := uS.userRepo.GetOTPByEmail(ctx, email)
// 	if err != nil {
// 		return "", errors.Wrap(err, "failed to find otp")
// 	}

// 	if storedOtp.OTP != otp {
// 		return "", errors.New("invalid otp")
// 	}
// 	if time.Now().Unix() > int64(storedOtp.ExpiredAt) {
// 		return "", errors.New("otp expired")
// 	}

// 	resetToken := utils.GenerateResetToken(32)

// 	err = uS.userRepo.UpdateResetToken(ctx, email, resetToken)
// 	if err != nil {
// 		return "", errors.Wrap(err, "failed to update reset token")
// 	}
// 	return resetToken, nil
// }

// func (uS *userService) ResetPassword(ctx context.Context, request dto.ResetPasswordRequest) error {
// 	if request.NewPassword != request.ConfirmPassword {
// 		return errors.New("passwords do not match")
// 	}

// 	otpRecord, err := uS.userRepo.GetByResetToken(ctx, request.ResetToken)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to find otp")
// 	}

// 	if err := request.PrepareResetPassword(); err != nil {
// 		return errors.Wrap(err, "failed to prepare reset password")
// 	}

// 	err = uS.userRepo.UpdatePassword(ctx, otpRecord.Email, request.NewPassword)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to update password")
// 	}

// 	err = uS.userRepo.DeleteResetToken(ctx, request.ResetToken)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to delete reset token")
// 	}
// 	return nil
// }

// func (uS *userService) ResendOTP(ctx context.Context, email string) error {
// 	getUser, err := uS.userRepo.FindByEmail(ctx, entities.User{Email: email})
// 	if err != nil {
// 		return errors.Wrap(err, "failed to find user")
// 	}

// 	existingOTP, err := uS.userRepo.GetOTPByEmail(ctx, email)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to find otp")
// 	}
// 	if existingOTP.ID == 0 {
// 		return errors.New("OTP has not been generated yet, please request a new OTP first")
// 	}
// 	err = uS.userRepo.DeleteOTPByEmail(ctx, email)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to delete otp")
// 	}

// 	otp := utils.GenerateOTP(6)
// 	otpInt, _ := strconv.Atoi(otp)
// 	expiredAt3Min := time.Now().Add(time.Minute * 3).Unix()

// 	err = uS.userRepo.InsertOTP(ctx, entities.OTP{
// 		Email:     email,
// 		OTP:       otpInt,
// 		ExpiredAt: int(expiredAt3Min),
// 	})

// 	if err != nil {
// 		return errors.Wrap(err, "failed to find otp")
// 	}

// 	msg := map[string]any{
// 		"email":    getUser.Email,
// 		"username": getUser.Username,
// 		"otp":      otpInt,
// 	}

// 	payload, err := json.Marshal(msg)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to marshal payload")
// 	}

// 	if err := broker.Publish("user.resend_otp", payload); err != nil {
// 		return errors.Wrap(err, "failed to publish resend otp")
// 	}

// 	return nil
// }


