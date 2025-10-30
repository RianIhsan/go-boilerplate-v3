package server

import (
	permissionDelivery "ams-sentuh/internal/features/permission/delivery/http"
	userDelivery "ams-sentuh/internal/features/user/delivery/http"
	"ams-sentuh/pkg/uploader"

	userRepo "ams-sentuh/internal/features/user/repository"
	userService "ams-sentuh/internal/features/user/service"

	"github.com/gin-gonic/gin"

	roleDelivery "ams-sentuh/internal/features/role/delivery/http"
	roleRepo "ams-sentuh/internal/features/role/repository"
	roleService "ams-sentuh/internal/features/role/service"

	accessDelivery "ams-sentuh/internal/features/access/delivery/http"
	accessRepo "ams-sentuh/internal/features/access/repository"
	accessService "ams-sentuh/internal/features/access/service"

	permissionRepo "ams-sentuh/internal/features/permission/repository"
	permissionService "ams-sentuh/internal/features/permission/service"
	"ams-sentuh/internal/middleware"
	casbinSV "ams-sentuh/internal/middleware/casbin"
	"net/http"
)

func (s *Server) Bootstrap() error {
	// -----------------------------------------------------------------------------------------------------------
	// create instance casbin
	enforcer := casbinSV.InitCasbin(s.cfg.Casbin.Model, s.cfg.Casbin.Policy)

	// -----------------------------------------------------------------------------------------------------------
	// create a new instance repositories
	userPostgresRepo := userRepo.NewUserPostgresRepository(s.db)
	rolePostgresRepo := roleRepo.NewRolePostgresRepository(s.db)
	accessPostgresRepo := accessRepo.NewAccessPostgresRepository(s.db)
	permissionPostgresRepo := permissionRepo.NewPermissionPostgresRepository(s.db)

	// -----------------------------------------------------------------------------------------------------------
	// create a new instance services
	initCasbin := casbinSV.NewService(enforcer)
	initMinio := uploader.NewMinioUploader(s.cfg)

	// deviceSV := deviceService.NewDeviceService(&deviceService.ServiceConfig{
	// 	DeviceRepoInterface: devicePostgresRepo,
	// 	Logger:              s.logger,
	// 	Config:              s.cfg,
	// 	LogRepoInterface:    logPostgresRepo,
	// })

	userSV := userService.NewUserService(&userService.ServiceConfig{
		UserRepoInterface: userPostgresRepo,
		Logger:            s.logger,
		Config:            s.cfg,
		Casbin:            initCasbin,
		MinioClient:       initMinio,
	})

	roleSV := roleService.NewRoleService(&roleService.ServiceConfig{
		RoleRepoInterface: rolePostgresRepo,
		Logger:            s.logger,
		Config:            s.cfg,
		Casbin:            initCasbin,
	})

	accessSV := accessService.NewAccessService(&accessService.ServiceConfig{
		AccessRepoInterface: accessPostgresRepo,
		Logger:              s.logger,
		Config:              s.cfg,
	})

	permissionSV := permissionService.NewPermissionService(&permissionService.ServiceConfig{
		PermissionRepoInterface: permissionPostgresRepo,
		Config:                  s.cfg,
		Logger:                  s.logger,
	})

	// whiteAppSV := whiteAppService.NewApplicationService(&whiteAppService.ServiceConfig{
	// 	ApplicationRepoInterface: appPostgresRepo,
	// 	Logger:                   s.logger,
	// 	Config:                   s.cfg,
	// 	DeviceServiceInterface:   deviceSV,
	// 	MinioClient:              initMinio,
	// })

	// areaSv := areaService.NewAreaService(&areaService.ServiceConfig{
	// 	AreaRepoInterface: areaPostgresRepo,
	// 	Config:            s.cfg,
	// 	MinioClient:       initMinio,
	// })

	// branchSV := branchService.NewBranchService(&branchService.ServiceConfig{
	// 	BranchRepoInterface: branchPostgresRepo,
	// })

	// clientSV := clientService.NewClientService(&clientService.ServiceConfig{
	// 	ClientRepo: clientPostgresRepo,
	// 	Config:     s.cfg,
	// })

	// ticketSV := ticketService.NewTicketService(&ticketService.ServiceConfig{
	// 	TicketRepoInterface: ticketPostgresRepo,
	// 	UserRepoInterface:   userPostgresRepo,
	// 	DeviceRepoInterface: devicePostgresRepo,
	// 	MinioClient:         initMinio,
	// 	Config:              s.cfg,
	// })

	// companySV := companyService.NewService(&companyService.ServiceConfig{
	// 	CompanyPostgresRepository: companyPostgresRepo,
	// })

	// widgetSV := widgetService.NewWidgetService(&widgetService.ServiceConfig{
	// 	WidgetRepo: widgetPostgresRepo,
	// })

	// -----------------------------------------------------------------------------------------------------------
	// create a new instance delivery
	// deviceDel := deviceDelivery.NewDeviceDelivery(&deviceDelivery.DeliveryConfig{
	// 	DeviceServiceInterface: deviceSV,
	// 	Config:                 s.cfg,
	// 	Logger:                 s.logger,
	// })

	userDel := userDelivery.NewUserDelivery(&userDelivery.DeliveryConfig{
		UserServiceInterface: userSV,
		Config:               s.cfg,
		Logger:               s.logger,
	})

	roleDel := roleDelivery.NewRoleDelivery(&roleDelivery.DeliveryConfig{
		RoleServiceInterface: roleSV,
		Config:               s.cfg,
		Logger:               s.logger,
	})

	accessDel := accessDelivery.NewAccessDelivery(&accessDelivery.DeliveryConfig{
		AccessServiceInterface: accessSV,
		Config:                 s.cfg,
		Logger:                 s.logger,
	})

	permissionDel := permissionDelivery.NewPermissionDelivery(&permissionDelivery.DeliveryConfig{
		PermissionServiceInterface: permissionSV,
		Config:                     s.cfg,
		Logger:                     s.logger,
	})

	// whiteAppDel := whiteAppDelivery.NewApplicationDelivery(&whiteAppDelivery.DeliveryConfig{
	// 	ApplicationServiceInterface: whiteAppSV,
	// 	Config:                      s.cfg,
	// 	Logger:                      s.logger,
	// })

	// areDel := areaDelivery.NewAreaDelivery(&areaDelivery.DeliveryConfig{
	// 	AreaServiceInterface: areaSv,
	// 	Config:               s.cfg,
	// 	Logger:               s.logger,
	// })

	// branchDel := branchDelivery.NewBranchDelivery(&branchDelivery.DeliveryConfig{
	// 	BranchServiceInterface: branchSV,
	// 	Config:                 s.cfg,
	// 	Logger:                 s.logger,
	// })

	// clientDel := clientDelivery.NewClientDelivery(&clientDelivery.DeliveryConfig{
	// 	ClientSV: clientSV,
	// 	Config:   s.cfg,
	// 	Logger:   s.logger,
	// })

	// ticketDel := ticketDelivery.NewTicketDelivery(&ticketDelivery.DeliveryConfig{
	// 	TicketSV: ticketSV,
	// 	Logger:   s.logger,
	// 	Config:   s.cfg,
	// })

	// companyDel := companyHandler.NewCompanyDelivery(&companyHandler.DeliveryConfig{
	// 	CompanyService: companySV,
	// 	Config:         s.cfg,
	// 	Logger:         s.logger,
	// })

	// widgetDel := widgetHandler.NewWidgetDelivery(&widgetHandler.DeliveryConfig{
	// 	WidgetSV: widgetSV,
	// 	Config:   s.cfg,
	// 	Logger:   s.logger,
	// })

	// -----------------------------------------------------------------------------------------------------------
	// create a new instance middleware
	middlewareManager := middleware.NewMiddlewareManager(&middleware.MiddlewareConfig{
		Logger:   s.logger,
		Config:   s.cfg,
		Enforcer: enforcer,
	})

	// middlewareManager.LogRateLimiterStatus(s.cfg)
	// Global middlewares
	s.app.Use(middlewareManager.RequestIdMiddleware())
	s.app.Use(middlewareManager.RequestLoggerMiddleware())

	// Apply general API rate limiting to all API routes
	// s.app.Use(middlewareManager.APIRateLimitMiddleware(s.cfg))

	// -----------------------------------------------------------------------------------------------------------
	// setup routes
	apiV1 := s.app.Group("/api")
	{
		// group device routes with device-specific rate limiting
		// deviceGroup := apiV1.Group("/v1")
		// deviceGroup.Use(middlewareManager.DeviceRateLimitMiddleware(s.cfg))
		// {
		// 	deviceDelivery.MapDeviceRoutes(deviceGroup, deviceDel, middlewareManager)
		// }

		// group user routes with auth rate limiting for sensitive endpoints
		userGroup := apiV1.Group("/v1")
		// Apply auth rate limiting to user endpoints that involve authentication
		// userAuthGroup := userGroup.Group("")
		//userAuthGroup.Use(middlewareManager.AuthRateLimitMiddleware(s.cfg))
		{
			// Map routes that need stricter rate limiting
			// You'll need to modify MapUserRoutes to separate auth routes
			userDelivery.MapUserRoutes(userGroup, userDel, middlewareManager)
		}

		// group role routes with admin rate limiting
		roleGroup := apiV1.Group("/v1")
		// roleGroup.Use(middlewareManager.AdminRateLimitMiddleware(s.cfg))
		{
			roleDelivery.MapRoleRoutes(roleGroup, roleDel, middlewareManager)
		}

		// access routes with admin rate limiting
		accessGroup := apiV1.Group("/v1")
		// accessGroup.Use(middlewareManager.AdminRateLimitMiddleware(s.cfg))
		{
			accessDelivery.MapAccessRoute(accessGroup, accessDel)
		}

		// permission routes with admin rate limiting
		permissionGroup := apiV1.Group("/v1")
		// permissionGroup.Use(middlewareManager.AdminRateLimitMiddleware(s.cfg))
		{
			permissionDelivery.MapPermissionRoutes(permissionGroup, permissionDel)
		}

		// application routes with standard API rate limiting
		// whiteAppGroup := apiV1.Group("/v1")
		// {
		// 	whiteAppDelivery.MapApplicationRoute(whiteAppGroup, whiteAppDel)
		// }

		// // area
		// areaGroup := apiV1.Group("/v1")
		// {
		// 	areaDelivery.MapAreaRoute(areaGroup, areDel)
		// }

		// // branch routes with admin rate limiting
		// branchGroup := apiV1.Group("/v1")
		// branchGroup.Use(middlewareManager.AdminRateLimitMiddleware(s.cfg))
		// {
		// 	branchDelivery.MapBranchRoutes(branchGroup, branchDel)
		// }

		// clientGroup := apiV1.Group("/v1")
		// clientGroup.Use(middlewareManager.ClientRateLimitMiddleware(s.cfg))
		// {
		// 	clientDelivery.MapClientRoutes(clientGroup, clientDel)
		// }

		// ticketGroup := apiV1.Group("/v1")
		// {
		// 	ticketDelivery.MapTicketRoutes(ticketGroup, ticketDel)
		// }

		// companyGroup := apiV1.Group("/v1")
		// {
		// 	companyHandler.MapCompanyRoutes(companyGroup, companyDel)
		// }

		// widgetGroup := apiV1.Group("/v1")
		// {
		// 	widgetHandler.MapWidgetRoutes(widgetGroup, widgetDel)
		// }
	}

	// Ping endpoint without additional rate limiting (already covered by global)
	apiV1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Pong!!",
		})
	})

	// Initialize WebSocket with Melody
	// ws.InitMelody()

	// WebSocket endpoint - you might want to add specific rate limiting for this
	// s.app.GET("/ws", ws.WebSocketHandler)

	// Rate limit status endpoint for monitoring
	// apiV1.GET("/rate-limit/status", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "Rate limiter is active",
	// 		"ip":      ctx.ClientIP(),
	// 		"limits": gin.H{
	// 			"api":    "100 requests per minute",
	// 			"auth":   "10 requests per minute",
	// 			"device": "1000 requests per minute",
	// 			"admin":  "200 requests per minute",
	// 			"client": "50 requests per minute",
	// 		},
	// 	})
	// })

	// apiV1.GET("/health", func(c *gin.Context) {
	// 	status := "healthy"
	// 	rateLimiterStatus := "disabled"

	// 	if s.cfg.RateLimiter.Enabled {
	// 		rateLimiterStatus = "enabled"
	// 	}

	// 	c.JSON(200, gin.H{
	// 		"status":       status,
	// 		"rate_limiter": rateLimiterStatus,
	// 		"timestamp":    time.Now(),
	// 	})
	// })

	return nil
}
