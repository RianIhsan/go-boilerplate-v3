package http

import (
	"ams-sentuh/config"
	"ams-sentuh/internal/features/user"
	"ams-sentuh/internal/features/user/dto"
	"ams-sentuh/pkg/httpErrors/response"
	"ams-sentuh/pkg/utils"
	"ams-sentuh/pkg/validation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type userDelivery struct {
	cfg     *config.Config
	logger  *logrus.Logger
	service user.UserServiceInterface
}

func NewUserDelivery(cfg *DeliveryConfig) user.UserDeliveryInterface {
	return &userDelivery{
		cfg:     cfg.Config,
		logger:  cfg.Logger,
		service: cfg.UserServiceInterface,
	}
}

func (u *userDelivery) RegisterUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		request := new(dto.RegisterUserRequest)
		if err := context.ShouldBindJSON(request); err != nil {
			utils.LogErrorResponse(context, u.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "Invalid request body")
			return
		}

		if err := validation.ValidateStruct(request); err != nil {
			utils.LogErrorResponse(context, u.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, err.Error())
			return
		}

		createdUser, err := u.service.AddUser(context, *request)
		if err != nil {
			utils.LogErrorResponse(context, u.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(context, http.StatusCreated, "User created successfully", createdUser)

	}
}

func (u *userDelivery) LoginUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		request := new(dto.LoginUserRequest)
		if err := context.ShouldBindJSON(request); err != nil {
			utils.LogErrorResponse(context, u.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "Invalid request body")
			return
		}

		if err := validation.ValidateStruct(request); err != nil {
			utils.LogErrorResponse(context, u.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, err.Error())
			return
		}

		token, err := u.service.LoginUser(context, request)
		if err != nil {
			utils.LogErrorResponse(context, u.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(context, http.StatusOK, "Login successful", token)
	}
}

func (u *userDelivery) GetList() gin.HandlerFunc {
	return func(context *gin.Context) {
		roleId := context.Query("roleId")
		roleIdUint, _ := strconv.ParseUint(roleId, 10, 64)

		data, err := u.service.GetList(context, roleIdUint)
		if err != nil {
			utils.LogErrorResponse(context, u.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(context, http.StatusOK, "success", data)

	}
}

func (u *userDelivery) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)
		user, err := u.service.GetById(c, idUint)
		if err != nil {
			utils.LogErrorResponse(c, u.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(c, http.StatusOK, "success", user)
	}
}

func (u *userDelivery) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)
		request := new(dto.UpdateUserRequest)
		if err := c.ShouldBindJSON(request); err != nil {
			utils.LogErrorResponse(c, u.logger, err)
			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
			return
		}

		err := u.service.Update(c, idUint, *request)
		if err != nil {
			utils.LogErrorResponse(c, u.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(c, http.StatusOK, "success", nil)
	}
}

func (u *userDelivery) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)
		err := u.service.Delete(c, idUint)
		if err != nil {
			utils.LogErrorResponse(c, u.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", nil)
	}
}

// func (u *userDelivery) GenerateOTP() gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		request := new(dto.GenerateOTPCode)
// 		if err := c.ShouldBindJSON(request); err != nil {
// 			utils.LogErrorResponse(c, u.logger, err)
// 			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
// 			return
// 		}

// 		err := u.service.ForgotPassword(c, request.Email)
// 		if err != nil {
// 			utils.LogErrorResponse(c, u.logger, err)
// 			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
// 			return
// 		}
// 		response.SendSuccesResponse(c, http.StatusOK, "success", nil)

// 	}
// }

// func (u *userDelivery) VerifyOTP() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		request := new(dto.VerifyOTPCode)
// 		if err := c.ShouldBindJSON(request); err != nil {
// 			utils.LogErrorResponse(c, u.logger, err)
// 			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
// 			return
// 		}

// 		token, err := u.service.VerifyOTP(c, request.Email, request.OTP)
// 		if err != nil {
// 			utils.LogErrorResponse(c, u.logger, err)
// 			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
// 			return
// 		}

// 		res := dto.VerifyOTPResponse{}
// 		res.ResetToken = token

// 		response.SendSuccesResponse(c, http.StatusOK, "success", res)
// 	}
// }

// func (u *userDelivery) ResetPassword() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		request := new(dto.ResetPasswordRequest)
// 		if err := c.ShouldBindJSON(request); err != nil {
// 			utils.LogErrorResponse(c, u.logger, err)
// 			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
// 			return
// 		}
// 		err := u.service.ResetPassword(c, *request)
// 		if err != nil {
// 			utils.LogErrorResponse(c, u.logger, err)
// 			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
// 			return
// 		}

// 		response.SendSuccesResponse(c, http.StatusOK, "success", nil)
// 	}
// }

// func (u *userDelivery) ResendOTP() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		request := new(dto.GenerateOTPCode)
// 		if err := c.ShouldBindJSON(request); err != nil {
// 			utils.LogErrorResponse(c, u.logger, err)
// 			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
// 			return
// 		}
// 		err := u.service.ResendOTP(c, request.Email)
// 		if err != nil {
// 			utils.LogErrorResponse(c, u.logger, err)
// 			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
// 			return
// 		}
// 		response.SendSuccesResponse(c, http.StatusOK, "success", nil)
// 	}
// }
