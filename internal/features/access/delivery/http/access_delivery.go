package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ams-sentuh/config"
	"ams-sentuh/internal/features/access"
	"ams-sentuh/internal/features/access/dto"
	"ams-sentuh/pkg/httpErrors/response"
	"ams-sentuh/pkg/utils"
	"net/http"
	"strconv"
)

type accessDelivery struct {
	cfg     *config.Config
	logger  *logrus.Logger
	service access.AccessServiceInterface
}

func NewAccessDelivery(cfg *DeliveryConfig) access.AccessDeliveryInterface {
	return &accessDelivery{
		cfg:     cfg.Config,
		logger:  cfg.Logger,
		service: cfg.AccessServiceInterface,
	}
}

func (a accessDelivery) CreateAccess() gin.HandlerFunc {
	return func(context *gin.Context) {
		reqDTO := new(dto.AccessRegisterRequest)
		if err := context.ShouldBindJSON(reqDTO); err != nil {
			utils.LogErrorResponse(context, a.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "invalid payload")
			return
		}

		data, err := a.service.RegisterAccess(context, *reqDTO)
		if err != nil {
			utils.LogErrorResponse(context, a.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "failed register access")
			return
		}

		response.SendSuccesResponse(context, http.StatusCreated, "access created", data)
	}
}

func (a accessDelivery) GetAllAccess() gin.HandlerFunc {
	return func(context *gin.Context) {
		accessData, err := a.service.GetAllAccess(context)
		if err != nil {
			utils.LogErrorResponse(context, a.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, err.Error())
			return
		}
		response.SendSuccesResponse(context, http.StatusOK, "success", accessData)
	}
}

func (a accessDelivery) GetAccess() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		idUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			utils.LogErrorResponse(context, a.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "invalid id")
			return
		}
		accessData, err := a.service.GetAccess(context, idUint)
		if err != nil {
			utils.LogErrorResponse(context, a.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "failed get access")
			return
		}
		response.SendSuccesResponse(context, http.StatusOK, "success", accessData)
	}
}

func (a accessDelivery) UpdateAccess() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		idUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			utils.LogErrorResponse(context, a.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "invalid id")
			return

		}
		reqDTO := new(dto.UpdateAccessRequest)
		if err := context.ShouldBindJSON(reqDTO); err != nil {
			utils.LogErrorResponse(context, a.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "invalid payload")
			return
		}

		err = a.service.UpdateAccess(context, idUint, *reqDTO)
		if err != nil {
			utils.LogErrorResponse(context, a.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "failed update access")
			return
		}

		response.SendSuccesResponse(context, http.StatusOK, "success", nil)

	}
}

func (a accessDelivery) DeleteAccess() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		idUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			utils.LogErrorResponse(context, a.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "invalid id")
			return
		}
		err = a.service.DeleteAccess(context, idUint)
		if err != nil {
			utils.LogErrorResponse(context, a.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "failed delete access")
			return
		}
		response.SendSuccesResponse(context, http.StatusOK, "success", nil)
	}
}
