package http

import (
	"ams-sentuh/internal/features/user"
	"ams-sentuh/internal/middleware"

	"github.com/gin-gonic/gin"
)

func MapUserRoutes(
	userGroup *gin.RouterGroup,
	delivery user.UserDeliveryInterface,
	mw *middleware.MiddlewareManager) {

	userGroup.POST("/register", delivery.RegisterUser())
	userGroup.POST("/login", delivery.LoginUser())
	userGroup.PUT("/users/:id", delivery.Update())
	userGroup.DELETE("users/:id", delivery.DeleteUser())

	casbinGroup := userGroup.Group("")
	casbinGroup.Use(mw.AuthMiddleware())
	casbinGroup.Use(mw.CasbinMiddleware())

	casbinGroup.GET("/users", delivery.GetList())
	casbinGroup.GET("/users/:id", delivery.GetById())
	casbinGroup.PUT("/users/protected", delivery.SelfUpdate())
	casbinGroup.PUT("/users/avatar", delivery.UpdateAvatar())

}
