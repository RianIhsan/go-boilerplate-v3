package access

import "github.com/gin-gonic/gin"

type AccessDeliveryInterface interface {
	CreateAccess() gin.HandlerFunc
	GetAllAccess() gin.HandlerFunc
	GetAccess() gin.HandlerFunc
	UpdateAccess() gin.HandlerFunc
	DeleteAccess() gin.HandlerFunc
}
