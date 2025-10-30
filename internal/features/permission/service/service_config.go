package service

import (
	"github.com/sirupsen/logrus"
	"ams-sentuh/config"
	"ams-sentuh/internal/features/permission"
)

type ServiceConfig struct {
	PermissionRepoInterface permission.PermissionRepositoryInterface
	Logger                  *logrus.Logger
	Config                  *config.Config
}
