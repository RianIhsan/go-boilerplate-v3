package http

import (
	"ams-sentuh/internal/features/access"
	"ams-sentuh/internal/middleware"

	"github.com/gin-gonic/gin"
)

func MapAccessRoute(accessGroup *gin.RouterGroup, controller access.AccessDeliveryInterface, mw *middleware.MiddlewareManager) {
	protectedRoutes := accessGroup.Group("")
	protectedRoutes.Use(mw.AuthMiddleware())
	protectedRoutes.Use(mw.CasbinMiddleware())

	protectedRoutes.POST("/access", controller.CreateAccess())
	protectedRoutes.GET("/access", controller.GetAllAccess())
	protectedRoutes.GET("/access/:id", controller.GetAccess())
	protectedRoutes.PUT("/access/:id", controller.UpdateAccess())
	protectedRoutes.DELETE("/access/:id", controller.DeleteAccess())
}
