package service

import (
	"ams-sentuh/config"
	"ams-sentuh/internal/features/user"
	"ams-sentuh/internal/middleware/casbin"
	"ams-sentuh/pkg/uploader"

	"github.com/sirupsen/logrus"
)

type ServiceConfig struct {
	UserRepoInterface user.UserRepositoryInterface
	Logger            *logrus.Logger
	Config            *config.Config
	Casbin            casbin.CasbinService
	MinioClient       uploader.FileUploaderInterface
}
