package main

import (
	"ams-sentuh/config"
	// "context"
	// "fmt"
	"log"
	// "ams-sentuh/internal/entities"
	// logRepository "ams-sentuh/internal/features/log"
	// initLogRepo "ams-sentuh/internal/features/log/repository"
	"ams-sentuh/internal/server"
	// "ams-sentuh/pkg/audit"
	// "ams-sentuh/pkg/broker"
	// "ams-sentuh/pkg/cache"
	pg "ams-sentuh/pkg/db"
	"ams-sentuh/pkg/logger"
	"os"
	// "strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

// func StartDeviceStatusChecker(db *gorm.DB, logRepo logRepository.LogRepositoryInterface) {
// 	ticker := time.NewTicker(30 * time.Second)
// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				fmt.Println("🔄 Checking device status...")
// 				checkDevicesStatus(db, logRepo)
// 			}
// 		}
// 	}()
// }

// func checkDevicesStatus(DB *gorm.DB, logRepo logRepository.LogRepositoryInterface) {
// 	ctx := context.Background()

// 	keys, err := cache.RedisClient.Keys(ctx, "device_status:*").Result()
// 	if err != nil {
// 		log.Println("❌ Redis error:", err)
// 		return
// 	}

// 	activeIDs := make(map[string]bool)
// 	for _, key := range keys {
// 		deviceID := strings.TrimPrefix(key, "device_status:")
// 		activeIDs[deviceID] = true
// 	}

// 	var devices []entities.Device
// 	if err := DB.Select("id", "status").Find(&devices).Error; err != nil {
// 		log.Println("❌ Error fetching devices:", err)
// 		return
// 	}

// 	for _, device := range devices {
// 		expectedStatus := "DOWN"
// 		if activeIDs[device.ID] {
// 			expectedStatus = "UP"
// 		}

// 		if device.Status != expectedStatus {
// 			changes := []audit.FieldChange{
// 				{
// 					Field:     "Status",
// 					OldValue:  device.Status,
// 					NewValue:  expectedStatus,
// 					ChangedBy: "system",
// 				},
// 			}
// 			logs := audit.LogDeviceChanges(device.ID, changes)
// 			_ = logRepo.SaveLogs(ctx, logs)

// 			err := DB.Model(&device).Update("status", expectedStatus).Error
// 			if err != nil {
// 				log.Printf("❌ Failed to update status for device %s: %v\n", device.ID, err)
// 			}
// 		}
// 	}
// }

func main() {
	//gin.SetMode(gin.ReleaseMode)

	// -----------------------------------------------------------------------------------------------------------
	// initializing app config
	cfg, err := config.NewAppConfig(os.Getenv("config"))
	if err != nil {
		log.Fatal(err)
	}

	// -----------------------------------------------------------------------------------------------------------
	// initializing Logger
	apiLogger := logger.NewLogrusLogger(cfg)

	// -----------------------------------------------------------------------------------------------------------
	// initializing postgreSQL connection
	psqlDB, err := pg.NewPostgresConnection(cfg)
	if err != nil {
		apiLogger.Fatalf("Postgresql initialize: %v", err)
	}
	apiLogger.Info("PostgreSQL connected")

	// _, err = cache.NewRedisConnection(cfg)
	// if err != nil {
	// 	apiLogger.Fatalf("Redis initialize: %v", err)
	// }
	// apiLogger.Info("Redis connected")

	err = pg.Migrate(psqlDB)
	if err != nil {
		apiLogger.Fatalf("Postgresql migration: %v", err)
	}
	apiLogger.Info("PostgreSQL migration completed")

	// err = broker.InitBroker(cfg)
	// if err != nil {
	// 	apiLogger.Fatalf("Broker initialize: %v", err)
	// }
	// apiLogger.Info("Broker initialized")

	// -----------------------------------------------------------------------------------------------------------
	// instance gin framework
	app := gin.New()

	err = app.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		apiLogger.Fatalf("SetTrustedProxies failed | from Gin Instance: %v", err)
	}
	apiLogger.Info("SetTrustedProxies completed")

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins, you can specify specific origins if needed
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	app.Use(gin.Recovery())

	// -----------------------------------------------------------------------------------------------------------
	// create a new instance server
	s := server.NewServer(&server.ServerConfig{
		App:    app,
		Cfg:    cfg,
		Logger: apiLogger,
		Db:     psqlDB,
	})

	// Checker Device
	// logRepo := initLogRepo.NewLogPostgresRepository(psqlDB)
	// StartDeviceStatusChecker(psqlDB, logRepo)

	if err := s.Run(); err != nil {
		apiLogger.Fatal(err)
	}
}
