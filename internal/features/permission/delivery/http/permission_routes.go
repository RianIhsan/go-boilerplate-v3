package http

import (
	"github.com/gin-gonic/gin"
	"ams-sentuh/internal/features/permission"
)

func MapPermissionRoutes(pGroup *gin.RouterGroup, delivery permission.PermissionDeliveryInterface) {
	pGroup.POST("/permission", delivery.CreatePermission())
	pGroup.GET("/permission", delivery.GetListPermission())
	pGroup.GET("/permission/:id", delivery.GetPermission())
	pGroup.PUT("/permission/:id", delivery.UpdatePermission())
	pGroup.DELETE("/permission/:id", delivery.DeletePermission())
}
