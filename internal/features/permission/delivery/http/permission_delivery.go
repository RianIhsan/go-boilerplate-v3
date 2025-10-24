package http

import (
	"ams-sentuh/config"
	"ams-sentuh/internal/features/permission"
	"ams-sentuh/internal/features/permission/dto"
	"ams-sentuh/pkg/httpErrors/response"
	"ams-sentuh/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type permissionDelivery struct {
	cfg     *config.Config
	logger  *logrus.Logger
	service permission.PermissionServiceInterface
}

func (p permissionDelivery) CreatePermission() gin.HandlerFunc {
	return func(context *gin.Context) {
		request := new(dto.PermissionRegister)
		if err := context.ShouldBindJSON(request); err != nil {
			utils.LogErrorResponse(context, p.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "invalid request")
			return
		}

		createdPermission, err := p.service.AddPermission(context, *request)
		if err != nil {
			utils.LogErrorResponse(context, p.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, "internal server error")
			return
		}

		response.SendSuccesResponse(context, http.StatusCreated, "permission created successfully", createdPermission)
	}
}

func (p permissionDelivery) GetListPermission() gin.HandlerFunc {
	return func(context *gin.Context) {
		idString := context.Query("access_id")
		idUint, _ := utils.ConvertStringToUint(idString)
		permissions, err := p.service.GetListPermission(context, &idUint)
		if err != nil {
			utils.LogErrorResponse(context, p.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, "failed to get permissions")
			return
		}

		response.SendSuccesResponse(context, http.StatusOK, "permissions retrieved successfully", permissions)
	}
}

func (p permissionDelivery) GetPermission() gin.HandlerFunc {
	return func(context *gin.Context) {
		idParam := context.Param("id")
		id, err := StringToUint64(idParam)
		if err != nil {
			utils.LogErrorResponse(context, p.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "invalid permission id")
			return
		}

		permission, err := p.service.GetPermission(context, id)
		if err != nil {
			utils.LogErrorResponse(context, p.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, "failed to get permission")
			return
		}

		response.SendSuccesResponse(context, http.StatusOK, "permission retrieved successfully", permission)
	}
}

func (p permissionDelivery) UpdatePermission() gin.HandlerFunc {
	return func(context *gin.Context) {
		idParam := context.Param("id")
		idUint, err := StringToUint64(idParam)
		if err != nil {
			utils.LogErrorResponse(context, p.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "invalid permission id")
			return
		}

		request := new(dto.PermissionUpdate)
		if err := context.ShouldBindJSON(request); err != nil {
			utils.LogErrorResponse(context, p.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "invalid request")
			return
		}

		err = p.service.UpdatePermission(context, idUint, *request)
		if err != nil {
			utils.LogErrorResponse(context, p.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, "internal server error")
			return
		}

		response.SendSuccesResponse(context, http.StatusOK, "permission updated successfully", nil)
	}
}

func (p permissionDelivery) DeletePermission() gin.HandlerFunc {
	return func(context *gin.Context) {
		idParam := context.Param("id")
		id, err := StringToUint64(idParam)
		if err != nil {
			utils.LogErrorResponse(context, p.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "invalid permission id")
			return
		}
		err = p.service.DeletePermission(context, id)
		if err != nil {
			utils.LogErrorResponse(context, p.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, "internal server error")
			return
		}
		response.SendSuccesResponse(context, http.StatusOK, "permission deleted successfully", nil)
	}
}

func NewPermissionDelivery(cfg *DeliveryConfig) permission.PermissionDeliveryInterface {
	return &permissionDelivery{
		cfg:     cfg.Config,
		logger:  cfg.Logger,
		service: cfg.PermissionServiceInterface,
	}
}

func StringToUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}
