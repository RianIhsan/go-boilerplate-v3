package dto

import "ams-sentuh/internal/entities"

type AccessRegisterRequest struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Priority int64  `json:"priority"`
}

type UpdateAccessRequest struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Priority int64  `json:"priority"`
}

func ConvertToAccessEntity(req AccessRegisterRequest) entities.Access {
	return entities.Access{
		Name:     req.Name,
		Link:     req.Link,
		Priority: req.Priority,
	}
}
