package user

import (
	"fmt"
	"go-ecommerce/global"
	"go-ecommerce/internal/service/user"
	"go-ecommerce/internal/service/user/models"
	"go-ecommerce/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterController struct {
	registerService user.IUserRegister
}

func NewRegisterController(registerService user.IUserRegister) *RegisterController {
	return &RegisterController{registerService: registerService}
}

// User Registration godoc
// @Summary      Register user
// @Description  When user is registered, this API will send OTP to email
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        payload body models.RegisterInput true "payload"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /user/register [post]
func (c *RegisterController) Register(ctx *gin.Context) {
	var request models.RegisterInput
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidRequest, err.Error())
		return
	}

	global.Logger.Info("Register user", zap.Any("request", request))

	result, err := c.registerService.Register(ctx, &request)
	if err != nil {
		response.ErrorResponse(ctx, result.Code, err.Error())
		return
	}

	response.SuccessRepsonse(ctx, response.CodeSuccess, result)
}

// Verify OTP godoc
// @Summary      OTP Verification
// @Description  Verify OTP sent to user's registered email
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        payload body models.VerifyOTPInput true "payload"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /user/verify_otp [post]
func (c *RegisterController) VerifyOTP(ctx *gin.Context) {
	var request models.VerifyOTPInput
	fmt.Println(ctx)
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidRequest, err.Error())
		return
	}

	global.Logger.Info("Verify OTP", zap.Any("request", request))

	result, err := c.registerService.VerifyOTP(ctx, &request)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidOTP, err.Error())
		return
	}

	response.SuccessRepsonse(ctx, response.CodeSuccess, result)
}
