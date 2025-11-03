package http

import (
	"ams-sentuh/internal/features/role"
	"ams-sentuh/internal/middleware"

	"github.com/gin-gonic/gin"
)

func MapRoleRoutes(roleGroup *gin.RouterGroup, delivery role.RoleDeliveryInterface, mw *middleware.MiddlewareManager) {
	protectedGroup := roleGroup.Group("")
	protectedGroup.Use(mw.AuthMiddleware())
	protectedGroup.Use(mw.CasbinMiddleware())

	protectedGroup.GET("/role", delivery.GetAllRole())
	roleGroup.PUT("/role/permission", delivery.ModifyRolePermission())
	protectedGroup.POST("/role", delivery.RegisterRole())
	protectedGroup.GET("/role/:id", delivery.GetRoleByID())
	protectedGroup.PUT("/role/:id", delivery.UpdateRole())
	protectedGroup.DELETE("/role/:id", delivery.DeleteRoleByID())
}
