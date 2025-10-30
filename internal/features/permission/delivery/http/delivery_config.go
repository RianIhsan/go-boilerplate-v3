package http

import (
	"github.com/sirupsen/logrus"
	"ams-sentuh/config"
	"ams-sentuh/internal/features/permission"
)

type DeliveryConfig struct {
	PermissionServiceInterface permission.PermissionServiceInterface
	Config                     *config.Config
	Logger                     *logrus.Logger
}
