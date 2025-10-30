package dto

type PermissionRegister struct {
	Name     string `json:"name"`
	AccessId uint64 `json:"access_id"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	Type     string `json:"type"`
}

type PermissionUpdate struct {
	Name     string `json:"name"`
	AccessId uint64 `json:"access_id"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	Type     string `json:"type"`
}
