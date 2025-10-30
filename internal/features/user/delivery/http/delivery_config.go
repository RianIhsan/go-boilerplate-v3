package http

import (
	"github.com/sirupsen/logrus"
	"ams-sentuh/config"
	"ams-sentuh/internal/features/user"
)

type DeliveryConfig struct {
	UserServiceInterface user.UserServiceInterface
	Config               *config.Config
	Logger               *logrus.Logger
}
