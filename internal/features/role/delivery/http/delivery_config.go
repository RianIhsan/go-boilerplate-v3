package http

import (
	"github.com/sirupsen/logrus"
	"ams-sentuh/config"
	"ams-sentuh/internal/features/role"
)

type DeliveryConfig struct {
	RoleServiceInterface role.RoleServiceInterface
	Config               *config.Config
	Logger               *logrus.Logger
}
