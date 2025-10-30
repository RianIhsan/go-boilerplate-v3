package service

import (
	"github.com/sirupsen/logrus"
	"ams-sentuh/config"
	"ams-sentuh/internal/features/access"
)

type ServiceConfig struct {
	AccessRepoInterface access.AccessRepositoryInterface
	Logger              *logrus.Logger
	Config              *config.Config
}
