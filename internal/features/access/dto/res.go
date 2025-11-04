package dto

import (
	"ams-sentuh/internal/entities"
	"time"
)

type AccessRegisterResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Link     string `json:"link"`
	Priority int64  `json:"priority"`
}

type PermissionResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAllAccessResponse struct {
	ID          uint64               `json:"id"`
	Name        string               `json:"name"`
	Link        string               `json:"link"`
	Priority    int64                `json:"priority"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Permissions []PermissionResponse `json:"permissions"`
}

func ConvertToPermissionResponse(data entities.Permission) PermissionResponse {
	return PermissionResponse{
		ID:        data.ID,
		Name:      data.Name,
		Path:      data.Path,
		Method:    data.Method,
		Type:      data.Type,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

// Convert slice of permission entities to slice of PermissionResponse
func ConvertToPermissionResponseList(data []entities.Permission) []PermissionResponse {
	var result []PermissionResponse
	for _, permission := range data {
		result = append(result, ConvertToPermissionResponse(permission))
	}
	return result
}

func ConvertToGetAllAccessResponse(data entities.Access) GetAllAccessResponse {
	return GetAllAccessResponse{
		ID:          data.ID,
		Name:        data.Name,
		Link:        data.Link,
		Priority:    data.Priority,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		Permissions: ConvertToPermissionResponseList(data.Permissions),
	}
}

func ConvertToGetAllAccessResponseList(data []entities.Access) []GetAllAccessResponse {
	var result []GetAllAccessResponse
	for _, d := range data {
		result = append(result, ConvertToGetAllAccessResponse(d))
	}
	return result
}

func ConvertToAccessRegisterResponse(data entities.Access) *AccessRegisterResponse {
	return &AccessRegisterResponse{
		ID:       data.ID,
		Name:     data.Name,
		Link:     data.Link,
		Priority: data.Priority,
	}
}
