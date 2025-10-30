package dto

import (
	"ams-sentuh/internal/entities"
	"time"
)

type RegisterUserResponse struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleId   uint64 `json:"role_id"`
}

func ConvertToRegisterUserResponse(entities *entities.User) RegisterUserResponse {
	return RegisterUserResponse{
		Id:     entities.ID,
		Name:   entities.Name,
		Avatar: entities.Avatar,
		Email:  entities.Email,
		RoleId: entities.RoleID,
	}
}

func ToListUsers(users []entities.User) (response []RegisterUserResponse) {
	for _, user := range users {
		response = append(response, ConvertToRegisterUserResponse(&user))
	}
	return response

}

type JwtToken struct {
	Email        string  `json:"email"`
	AccessToken  string  `json:"access_Token"`
	RefreshToken string  `json:"refresh_token"`
	CompanyID    uint64  `json:"company_id"`
	BranchID     *uint64 `json:"branch_id"`
	RoleID       uint64  `json:"role_id"`
}

type VerifyOTPResponse struct {
	ResetToken string `json:"reset_token"`
}

type UserDTO struct {
	ID        uint64    `json:"id"`
	Avatar    string    `json:"avatar"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	NfCTag    string    `json:"nfc_tag"`
	RoleId    uint64    `json:"role_id"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToUserDTO(user entities.User) UserDTO {
	return UserDTO{
		ID:        user.ID,
		Avatar:    user.Avatar,
		Name:      user.Name,
		Email:     user.Email,
		NfCTag:    user.NFCTag,
		RoleId:    user.RoleID,
		RoleName:  user.Role.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToListUsersResponse(users []entities.User) (response []UserDTO) {
	for _, user := range users {
		response = append(response, ToUserDTO(user))
	}
	return response
}
