package user

import (
	"go-ecommerce/internal/service/user"
	"go-ecommerce/internal/service/user/models"
	"go-ecommerce/pkg/response"
	"go-ecommerce/pkg/utils/context"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginService user.IUserLogin
}

func NewLoginController(loginService user.IUserLogin) *LoginController {
	return &LoginController{loginService: loginService}
}

// User Login godoc
// @Summary      Login user
// @Description  Login a user and return a token
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        payload body models.LoginInput true "payload"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /user/login [post]
func (c *LoginController) Login(ctx *gin.Context) {
	var request models.LoginInput
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidRequest, err.Error())
		return
	}

	result, err := c.loginService.Login(ctx, &request)
	if err != nil {
		response.ErrorResponse(ctx, result.Code, err.Error())
	}

	response.SuccessRepsonse(ctx, response.CodeSuccess, result)
}

// User Setup2FA godoc
// @Summary      Setup 2FA
// @Description  Setup 2FA token via email
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        authorization header string true "authorization token"
// @Param        payload body models.Setup2FARequest true "payload"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /user/2fa/setup [post]
func (c *LoginController) Setup2FA(ctx *gin.Context) {
	var request models.Setup2FARequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidRequest, err.Error())
		return
	}

	input := models.Setup2FAInput{
		Email: request.Email,
		Type:  request.Type,
	}

	// Get user ID from subjectUUID (populated to context via middleware)
	userID, err := context.GetUserID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidRequest, err.Error())
		return
	}
	input.UserId = uint32(userID)

	result, err := c.loginService.Setup2FA(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, result, err.Error())
	}

	response.SuccessRepsonse(ctx, response.CodeSuccess, result)
}

// User Verify2FA godoc
// @Summary      Verify 2FA
// @Description  Verify 2FA token via email
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        authorization header string true "authorization token"
// @Param        payload body models.Verify2FARequest true "payload"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /user/2fa/verify [post]
func (c *LoginController) Verify2FA(ctx *gin.Context) {
	var request models.Verify2FARequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidRequest, err.Error())
		return
	}

	input := models.Verify2FAInput{
		Code2FA: request.Code2FA,
	}

	// Get user ID from subjectUUID (populated to context via middleware)
	userID, err := context.GetUserID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidRequest, err.Error())
		return
	}
	input.UserId = uint32(userID)

	result, err := c.loginService.Verify2FA(ctx, &input)
	if err != nil {
		response.ErrorResponse(ctx, result, err.Error())
	}

	response.SuccessRepsonse(ctx, response.CodeSuccess, result)
}
