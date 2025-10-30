package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/sirupsen/logrus"
	"ams-sentuh/config"
)

type MiddlewareConfig struct {
	Logger   *logrus.Logger
	Config   *config.Config
	Enforcer *casbin.Enforcer
}

// MiddlewareManager defines methods middleware
type MiddlewareManager struct {
	logger   *logrus.Logger
	cfg      *config.Config
	enforcer *casbin.Enforcer
}

// NewMiddlewareManager is a factory function for instance MiddlewareManager
func NewMiddlewareManager(config *MiddlewareConfig) *MiddlewareManager {
	return &MiddlewareManager{
		logger:   config.Logger,
		cfg:      config.Config,
		enforcer: config.Enforcer,
	}
}
