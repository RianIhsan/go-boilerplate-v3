package role

import "github.com/gin-gonic/gin"

type RoleDeliveryInterface interface {
	RegisterRole() gin.HandlerFunc
	GetAllRole() gin.HandlerFunc
	ModifyRolePermission() gin.HandlerFunc
	GetRoleByID() gin.HandlerFunc
	UpdateRole() gin.HandlerFunc
	DeleteRoleByID() gin.HandlerFunc
}
