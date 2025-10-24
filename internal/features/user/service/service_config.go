package service

import (
	"github.com/sirupsen/logrus"
	"ams-sentuh/config"
	"ams-sentuh/internal/features/user"
	"ams-sentuh/internal/middleware/casbin"
)

type ServiceConfig struct {
	UserRepoInterface user.UserRepositoryInterface
	Logger            *logrus.Logger
	Config            *config.Config
	Casbin            casbin.CasbinService
}
