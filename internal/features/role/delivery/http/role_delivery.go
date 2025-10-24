package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ams-sentuh/config"
	"ams-sentuh/internal/features/role"
	"ams-sentuh/internal/features/role/dto"
	"ams-sentuh/pkg/httpErrors/response"
	"ams-sentuh/pkg/utils"
	"ams-sentuh/pkg/validation"
	"net/http"
	"strconv"
)

type roleDelivery struct {
	cfg     *config.Config
	logger  *logrus.Logger
	service role.RoleServiceInterface
}

func (d roleDelivery) RegisterRole() gin.HandlerFunc {
	return func(context *gin.Context) {
		request := new(dto.RegisterRoleRequest)
		if err := context.ShouldBindJSON(request); err != nil {
			utils.LogErrorResponse(context, d.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "Invalid request body")
			return
		}

		if err := validation.ValidateStruct(request); err != nil {
			utils.LogErrorResponse(context, d.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, err.Error())
			return
		}

		createdRole, err := d.service.AddRole(context, *request)
		if err != nil {
			utils.LogErrorResponse(context, d.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(context, http.StatusCreated, "Role created successfully", createdRole)
	}
}

func (d roleDelivery) GetAllRole() gin.HandlerFunc {
	return func(context *gin.Context) {
		roles, err := d.service.GetAll(context)
		if err != nil {
			utils.LogErrorResponse(context, d.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(context, http.StatusOK, "success", roles)
	}
}

func (d roleDelivery) ModifyRolePermission() gin.HandlerFunc {
	return func(context *gin.Context) {
		request := new(dto.RolePermissions)
		if err := context.ShouldBindJSON(request); err != nil {
			utils.LogErrorResponse(context, d.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "Invalid request body")
			return
		}
		datas, err := d.service.ModifyRolePermission(context, *request)
		if err != nil {
			utils.LogErrorResponse(context, d.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(context, http.StatusOK, "success", dto.ConvertToRoleResponse(datas))
	}
}

func (d roleDelivery) GetRoleByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		IdStr := c.Param("id")
		IdUint, err := strconv.ParseUint(IdStr, 10, 64)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
			return
		}
		data, err := d.service.GetByID(c, IdUint)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", data)
	}
}

func (d roleDelivery) UpdateRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(dto.UpdateRoleRequest)
		idStr := c.Param("id")
		idUint, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
			return
		}
		err = d.service.UpdateRole(c, idUint, *req)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", nil)
	}
}

func (d roleDelivery) DeleteRoleByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		IdStr := c.Param("id")
		IdUint, err := strconv.ParseUint(IdStr, 10, 64)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
			return
		}
		err = d.service.DeleteRole(c, IdUint)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", nil)
	}
}
func NewRoleDelivery(cfg *DeliveryConfig) role.RoleDeliveryInterface {
	return &roleDelivery{
		cfg:     cfg.Config,
		logger:  cfg.Logger,
		service: cfg.RoleServiceInterface,
	}
}
