package server

import (
	"context"
	"fmt"
	"ams-sentuh/config"
	
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	ctxTimeout = 5
	certFile   = "./ssl/server.crt"
	keyFile    = "./ssl/server.key"
)

type ServerConfig struct {
	App    *gin.Engine
	Logger *logrus.Logger
	Cfg    *config.Config
	Db     *gorm.DB
}

type Server struct {
	app    *gin.Engine
	logger *logrus.Logger
	cfg    *config.Config
	db     *gorm.DB
}

func NewServer(config *ServerConfig) *Server {
	return &Server{
		app:    config.App,
		logger: config.Logger,
		cfg:    config.Cfg,
		db:     config.Db,
	}
}

func (s *Server) Run() error {
	if err := s.Bootstrap(); err != nil {
		return errors.Wrap(err, "Server.Run.Bootstrap")
	}

	// Setup CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// HTTP Server configuration
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port),
		ReadTimeout:  time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout: time.Second * s.cfg.Server.WriteTimeout,
		Handler:      corsMiddleware.Handler(s.app),
	}

	// Context yang otomatis cancel kalau ada signal interrupt/terminate
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	serverError := make(chan error, 1)

	go func() {
		if s.cfg.Server.SSL {
			s.logger.Infof("TLS server listening on %s", server.Addr)
			serverError <- server.ListenAndServeTLS(certFile, keyFile)
		} else {
			s.logger.Infof("Server listening on %s", server.Addr)
			serverError <- server.ListenAndServe()
		}
	}()

	select {
	case err := <-serverError:
		if err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("server error: %w", err)
		}
	case <-ctx.Done():
		s.logger.Info("Shutting down server gracefully...")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	if s.db != nil {
		if dbSQL, err := s.db.DB(); err == nil {
			if err := dbSQL.Close(); err != nil {
				s.logger.Errorf("Error closing DB: %v", err)
			} else {
				s.logger.Info("Database connection closed")
			}
		} else {
			s.logger.Errorf("Error getting DB from GORM: %v", err)
		}
	}

	// if cache.RedisClient != nil {
	// 	if err := cache.RedisClient.Close(); err != nil {
	// 		s.logger.Errorf("Error closing Redis: %v", err)
	// 	} else {
	// 		s.logger.Info("Redis connection closed")
	// 	}
	// }

	// if ws.GetMelodyInstance() != nil {
	// 	if err := ws.GetMelodyInstance().Close(); err != nil {
	// 		s.logger.Errorf("Error closing all websocket connections: %v", err)
	// 	} else {
	// 		s.logger.Info("All websocket connections closed")
	// 	}
	// }

	s.logger.Info("Server exited properly")
	return nil
}
