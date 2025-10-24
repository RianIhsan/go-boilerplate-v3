package http

import (
	"github.com/gin-gonic/gin"
	"ams-sentuh/internal/features/access"
)

func MapAccessRoute(accessGroup *gin.RouterGroup, controller access.AccessDeliveryInterface) {
	accessGroup.POST("/access", controller.CreateAccess())
	accessGroup.GET("/access", controller.GetAllAccess())
	accessGroup.GET("/access/:id", controller.GetAccess())
	accessGroup.PUT("/access/:id", controller.UpdateAccess())
	accessGroup.DELETE("/access/:id", controller.DeleteAccess())
}
