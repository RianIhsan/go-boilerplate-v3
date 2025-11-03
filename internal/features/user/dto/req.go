package dto

import (
	"ams-sentuh/config"
	"ams-sentuh/internal/entities"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Username string `json:"username" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
	RoleId   uint   `json:"role_id"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   uint64 `json:"role_id"`
}

type SelfUpdateRequest struct {
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	NFCTag *string `json:"nfc_tag"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type GenerateOTPCode struct {
	Email string `json:"email" validate:"required,email,max=100"`
}

type VerifyOTPCode struct {
	Email string `json:"email" validate:"required,email,max=100"`
	OTP   int    `json:"otp" validate:"required"`
}

type ResetPasswordRequest struct {
	ResetToken      string `json:"reset_token"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func ConvertToEntityLoginRequest(request LoginUserRequest) entities.User {
	return entities.User{
		Email:    request.Email,
		Password: request.Password,
	}
}

func ConvertToEntityUserRequest(request RegisterUserRequest, cfg *config.Config) entities.User {
	return entities.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Avatar:   cfg.Minio.DefaultAvatar,
		Password: request.Password,
		RoleID:   uint64(request.RoleId),
	}
}

func (u *RegisterUserRequest) HashPassword() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "User.HashPassword.GenerateFromPassword")
	}
	u.Password = string(hashedPass)
	return nil
}

func (u *RegisterUserRequest) PrepareCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}
	return nil
}
func (u *LoginUserRequest) ComparePassword(hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password)); err != nil {
		return errors.Wrap(err, "User.ComparePassword.CompareHashAndPassword")
	}
	return nil
}

func (u *ResetPasswordRequest) HashPasswordReset() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "User.HashPassword.GenerateFromPassword")
	}
	u.NewPassword = string(hashed)
	return nil
}

func (u *ResetPasswordRequest) PrepareResetPassword() error {
	u.NewPassword = strings.TrimSpace(u.NewPassword)
	if err := u.HashPasswordReset(); err != nil {
		return err
	}
	return nil
}

func (u *UpdateUserRequest) HashPassword() error {
	if u.Password == "" {
		return nil // jika kosong, tidak perlu hash
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "User.Update.HashPassword.GenerateFromPassword")
	}
	u.Password = string(hashedPass)
	return nil
}

func (u *UpdateUserRequest) PrepareUpdate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)
	u.Name = strings.TrimSpace(u.Name)

	if err := u.HashPassword(); err != nil {
		return err
	}
	return nil
}
