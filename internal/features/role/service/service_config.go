package service

import (
	"github.com/sirupsen/logrus"
	"ams-sentuh/config"
	"ams-sentuh/internal/features/role"
	"ams-sentuh/internal/middleware/casbin"
)

type ServiceConfig struct {
	RoleRepoInterface role.RoleRepositoryInterface
	Logger            *logrus.Logger
	Config            *config.Config
	Casbin            casbin.CasbinService
}
