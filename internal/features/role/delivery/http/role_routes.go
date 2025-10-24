package http

import (
	"github.com/gin-gonic/gin"
	"ams-sentuh/internal/features/role"
	"ams-sentuh/internal/middleware"
)

func MapRoleRoutes(roleGroup *gin.RouterGroup, delivery role.RoleDeliveryInterface, mw *middleware.MiddlewareManager) {
	casbinGroup := roleGroup.Group("")
	casbinGroup.Use(mw.AuthMiddleware())
	casbinGroup.Use(mw.CasbinMiddleware())
	casbinGroup.GET("/role", delivery.GetAllRole())
	casbinGroup.PUT("/role/permission", delivery.ModifyRolePermission())
	casbinGroup.POST("/role", delivery.RegisterRole())
	casbinGroup.GET("/role/:id", delivery.GetRoleByID())
	casbinGroup.PUT("/role/:id", delivery.UpdateRole())
	casbinGroup.DELETE("/role/:id", delivery.DeleteRoleByID())
}
