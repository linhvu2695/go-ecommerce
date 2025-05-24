package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-ecommerce/global"
	"go-ecommerce/internal/constants"
	"go-ecommerce/internal/database"
	"go-ecommerce/internal/service/user/models"
	"go-ecommerce/pkg/response"
	"go-ecommerce/pkg/utils/auth"
	"go-ecommerce/pkg/utils/cache"
	"go-ecommerce/pkg/utils/encrypt"
	"go-ecommerce/pkg/utils/random"
	emailUtils "go-ecommerce/pkg/utils/sendto/email"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type UserLogin struct {
	r *database.Queries
}

func NewUserLogin(r *database.Queries) UserLogin {
	return UserLogin{r: r}
}

func (u UserLogin) Setup2FA(ctx context.Context, in *models.Setup2FAInput) (int, error) {
	// Check if 2FA is already enabled
	is2FAEnabled, err := u.r.Is2FAEnabled(ctx, in.UserId)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("error checking if 2FA is enabled for user ID %d", in.UserId), zap.Error(err))
		return response.Err2FASetupFailed, err
	}
	if is2FAEnabled > 0 {
		return response.ErrInvalidRequest, fmt.Errorf("two factor authentication is already enabled for user %d", in.UserId)
	}

	// Create new 2FA
	err = u.r.Enable2FAEmail(ctx, database.Enable2FAEmailParams{
		UserID: in.UserId,
		Type:   database.GoDbUser2faType(in.Type),
		Email:  sql.NullString{String: in.Email, Valid: true},
	})
	if err != nil {
		global.Logger.Error(fmt.Sprintf("error enable 2FA for user ID %d", in.UserId), zap.Error(err))
		return response.Err2FASetupFailed, err
	}

	// Send OTP to email
	key2FAOtp := cache.Get2FAOtpKey(int(in.UserId))
	otp := random.GenerateOtp()
	go global.Redis.Set(ctx, key2FAOtp, otp, constants.OTP_EXPIRATION_TIME*time.Second)

	err = emailUtils.SendEmailOtp([]string{in.Email}, global.Config.Smtp.Username, otp)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("error sending 2FA OTP email.\nUser ID: %d.\nEmail: %s", in.UserId, in.Email), zap.Error(err))
		return response.Err2FASetupFailed, err
	}

	return response.CodeSuccess, nil
}

func (u UserLogin) Verify2FA(ctx context.Context, in *models.Verify2FAInput) (int, error) {
	// Check if OTP is available
	enable, err := u.r.Is2FAEnabled(ctx, in.UserId)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("error checking if 2FA is enabled for user ID %d", in.UserId), zap.Error(err))
		return response.Err2FAVerifyFailed, err
	}
	if enable > 0 {
		return response.ErrInvalidRequest, fmt.Errorf("two factor authentication is already enabled for user %d", in.UserId)
	}

	// Get OTP from redis
	key2FAOtp := cache.Get2FAOtpKey(int(in.UserId))
	otp, err := global.Redis.Get(ctx, key2FAOtp).Result()
	if err == redis.Nil {
		return response.Err2FAVerifyFailed, fmt.Errorf("key %s does not exist", key2FAOtp)
	} else if err != nil {
		global.Logger.Error(fmt.Sprintf("error getting 2FA OTP for user ID %d", in.UserId), zap.Error(err))
		return response.Err2FAVerifyFailed, err
	}

	// Check OTP
	if otp != in.Code2FA {
		return response.Err2FAVerifyFailed, fmt.Errorf("OTP does not match")
	}

	// Update 2FA record status
	err = u.r.Update2FAStatus(ctx, database.Update2FAStatusParams{
		UserID: in.UserId,
		Type:   database.GoDbUser2faTypeEMAIL,
	})
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Failed to update 2FA status for user ID %d", in.UserId), zap.Error(err))
		return response.ErrInternalError, err
	}

	// Remove OTP from redis
	_, err = global.Redis.Del(ctx, key2FAOtp).Result()
	if err != nil {
		global.Logger.Error(fmt.Sprintf("error deleting 2FA OTP from cache for user ID %d", in.UserId), zap.Error(err))
		return response.ErrInternalError, err
	}

	return response.CodeSuccess, nil
}

func (u UserLogin) Login(ctx context.Context, in *models.LoginInput) (out models.LoginOutput, err error) {
	user, err := u.r.GetUserByEmail(ctx, in.Email)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Error get user by email %s", in.Email), zap.Error(err))
		out.Code = response.ErrAuthFailed
		return out, err
	}

	// Check password
	if user.PasswordHash != encrypt.HashPassword(in.Password, user.PasswordSalt) {
		out.Code = response.ErrAuthFailed
		err = fmt.Errorf("does not match password")
		return out, err
	}

	// Check two-factor authentication
	is2FAEnabled, err := u.r.Is2FAEnabled(ctx, user.ID)
	if err != nil {
		out.Code = response.ErrAuthFailed
		return out, err
	}
	if is2FAEnabled > 0 {
		global.Logger.Info(fmt.Sprintf("two factor authentication enabled for user ID %d", user.ID))

		// Setpu 2FA OTP in redis
		key2FAOtp := cache.Get2FAOtpKey(int(user.ID))
		otp := random.GenerateOtp()
		err = global.Redis.SetEx(ctx, key2FAOtp, otp, constants.OTP_EXPIRATION_TIME*time.Second).Err()
		if err != nil {
			out.Code = response.ErrAuthFailed
			return out, err
		}

		// Send email (the one setup for 2FA - not the one used for login)
		twoFactorInfo, err := u.r.Get2FAByUserAndType(ctx, database.Get2FAByUserAndTypeParams{
			UserID: user.ID,
			Type:   database.GoDbUser2faTypeEMAIL,
		})
		if err != nil {
			out.Code = response.ErrInternalError
			return out, err
		}

		err = emailUtils.SendEmailOtp([]string{twoFactorInfo.Email.String}, global.Config.Smtp.Username, otp)
		if err != nil {
			out.Code = response.ErrAuthFailed
			return out, err
		}

		out.Code = response.CodeSuccess
		out.Message = fmt.Sprintf("Two factor OTP sent to email %s", twoFactorInfo.Email.String)
		return out, nil
	}

	// Update last login date
	go u.r.UpdateUserLastLoginDate(ctx, user.ID)

	// Generate JWT token
	uuidString := random.GenerateToken()

	userInfo, err := json.Marshal(user)
	if err != nil {
		global.Logger.Error("Fail to marshal user information", zap.Error(err))
		out.Code = response.ErrInternalError
		return out, err
	}

	err = global.Redis.Set(ctx, uuidString, userInfo, constants.OTP_EXPIRATION_TIME*time.Minute).Err()
	if err != nil {
		global.Logger.Error("Error adding Token to Redis", zap.Error(err))
		out.Code = response.ErrInternalError
		return out, err
	}

	out.Token, err = auth.CreateToken(uuidString)
	if err != nil {
		global.Logger.Error("Error creating JWT token", zap.Error(err))
		out.Code = response.ErrInternalError
		return out, err
	}

	out.Code = response.CodeSuccess
	return out, err
}

func (u UserLogin) Logout(ctx context.Context, token string) error {
	// Implement logic to invalidate user session
	return nil
}
