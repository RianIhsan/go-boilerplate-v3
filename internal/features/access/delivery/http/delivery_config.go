package http

import (
	"github.com/sirupsen/logrus"
	"ams-sentuh/config"
	"ams-sentuh/internal/features/access"
)

type DeliveryConfig struct {
	AccessServiceInterface access.AccessServiceInterface
	Config                 *config.Config
	Logger                 *logrus.Logger
}
