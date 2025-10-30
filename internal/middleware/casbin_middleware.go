package middleware

import (
	"fmt"
	"ams-sentuh/pkg/httpErrors/response"
	"ams-sentuh/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mw *MiddlewareManager) CasbinMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if mw.enforcer == nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Casbin enforcer is not initialized",
			})
			return
		}

		authRaw, exists := context.Get("auth")
		if !exists {
			fmt.Println("Error from HERE!")
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		auth := authRaw.(*Auth)

		sub := fmt.Sprintf("%d", auth.Id)
		obj := context.Request.URL.Path
		act := context.Request.Method

		ok, err := mw.enforcer.Enforce(sub, obj, act)
		if err != nil {
			fmt.Printf("⛔ Enforce ERROR: %v\n", err)
			utils.LogErrorResponse(context, mw.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, "Error checking permission")
			context.Abort()
			return
		}
		if !ok {
			utils.LogErrorResponse(context, mw.logger, err)
			response.SendErrorResponse(context, http.StatusForbidden, "Forbidden Casbin")
			context.Abort()
			return
		}

		context.Next()
	}
}
