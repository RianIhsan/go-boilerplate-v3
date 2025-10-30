package permission

import "github.com/gin-gonic/gin"

type PermissionDeliveryInterface interface {
	CreatePermission() gin.HandlerFunc
	GetListPermission() gin.HandlerFunc
	GetPermission() gin.HandlerFunc
	UpdatePermission() gin.HandlerFunc
	DeletePermission() gin.HandlerFunc
}
