package impl

import (
	"context"
	"fmt"
	"go-ecommerce/global"
	"go-ecommerce/internal/constants"
	"go-ecommerce/internal/database"
	"go-ecommerce/internal/service/user/models"
	"go-ecommerce/pkg/response"
	"go-ecommerce/pkg/utils/cache"
	"go-ecommerce/pkg/utils/encrypt"
	"go-ecommerce/pkg/utils/random"
	emailUtils "go-ecommerce/pkg/utils/sendto/email"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type UserRegister struct {
	r *database.Queries
}

func NewUserRegister(r *database.Queries) UserRegister {
	return UserRegister{r: r}
}

func (u UserRegister) Register(ctx context.Context, in *models.RegisterInput) (out models.RegisterOutput, err error) {
	// Check email exists in DB
	emailExists, err := u.r.CheckEmailExists(ctx, in.Email)
	if err != nil {
		global.Logger.Error("Error checking email existence", zap.Error(err))
		out.Code = response.ErrInternalError
		return out, err
	}
	if emailExists {
		out.Code = response.ErrUserExists
		return out, fmt.Errorf("user already registered with this email")
	}

	// Check if registration already existed for this email
	userOtpKey := cache.GetUserOtpKey(in.Email)
	userOtp, err := global.Redis.Get(ctx, userOtpKey).Result()

	if err == nil && userOtp != "" {
		out.Code = response.ErrUserAlreadyRegistered
		return out, nil
	} else if err != nil && err != redis.Nil {
		global.Logger.Error("Error getting user OTP from Redis", zap.Error(err))
		out.Code = response.ErrInternalError
		return out, err
	}

	// Detect spam

	// Encrypt password
	passwordSalt, err := encrypt.GenerateSalt(constants.PASSWORD_SALT_LENGTH)
	if err != nil {
		global.Logger.Error("Error generating password salt", zap.Error(err))
		out.Code = response.ErrInternalError
		return out, err
	}

	passwordHash := encrypt.HashPassword(in.Password, passwordSalt)

	// Create user in DB
	createUserResult, err := u.r.CreateUser(ctx, database.CreateUserParams{
		Firstname:    in.Firstname,
		Lastname:     in.Lastname,
		Username:     in.Username,
		Email:        in.Email,
		PasswordHash: passwordHash,
		PasswordSalt: passwordSalt,
		Status:       constants.USER_STATUS_PENDING,
	})
	if err != nil {
		global.Logger.Error("Error creating user in DB", zap.Error(err))
		out.Code = response.ErrInternalError
		return out, err
	}

	createUserId, err := createUserResult.LastInsertId()
	if err != nil {
		global.Logger.Error("Error getting user ID from DB", zap.Error(err))
		out.Code = response.ErrInternalError
		return out, err
	}

	// Create & Add OTP to Redis
	otp := random.GenerateOtp()
	if in.Purpose == "testing" {
		otp = 123456 // For testing purpose
	}

	err = global.Redis.Set(ctx, userOtpKey, otp, constants.OTP_EXPIRATION_TIME*time.Second).Err()
	if err != nil {
		global.Logger.Error("Error adding OTP to Redis", zap.Error(err))
		out.Code = response.ErrInternalError
		return out, err
	}

	// Send email OTP
	err = emailUtils.SendEmailOtp([]string{in.Email}, global.Config.Smtp.Username, otp)
	if err != nil {
		global.Logger.Error("Error sending email OTP", zap.Error(err))
		out.Code = response.ErrInternalError
		return out, err
	}

	// Handle output
	out.Code = response.CodeSuccess
	out.UserId = fmt.Sprintf("%d", createUserId)

	return out, nil
}

func (u UserRegister) VerifyOTP(ctx context.Context, in *models.VerifyOTPInput) (out models.VerifyOTPOutput, err error) {
	// Get OTP from Redis
	otp, err := global.Redis.Get(ctx, cache.GetUserOtpKey(in.Email)).Result()
	if err != nil && err != redis.Nil {
		global.Logger.Error("Error getting user OTP from Redis", zap.Error(err))
		out.Code = response.ErrInternalError
		return out, err
	}

	// Verify OTP
	if otp != in.OTP {
		// Spam detection
		attemptCount, err := global.Redis.Get(ctx, cache.GetUserOtpAttemptKey(in.Email)).Int()
		if err != nil {
			if err == redis.Nil {
				attemptCount = 0
			} else {
				global.Logger.Error("Error getting user OTP attempt count from Redis", zap.Error(err))
				out.Code = response.ErrInternalError
				return out, err
			}
		}

		attemptCount++
		if attemptCount >= constants.MAX_OTP_ATTEMPTS {
			global.Logger.Info("User exceeded maximum OTP attempts", zap.String("email", in.Email))
			out.Code = response.ErrMaxOTPAttemptsReached
			return out, fmt.Errorf("user exceeded maximum OTP attempts")
		}

		err = global.Redis.Set(ctx, cache.GetUserOtpAttemptKey(in.Email), attemptCount, constants.OTP_EXPIRATION_TIME*time.Second).Err()
		if err != nil {
			global.Logger.Error("Error setting user OTP attempt count to Redis", zap.Error(err))
			out.Code = response.ErrInternalError
			return out, err
		}

		out.Code = response.ErrInvalidOTP
		return out, fmt.Errorf("OTP does not match")
	}

	// Remove OTP data from Redis
	err = global.Redis.Del(ctx, cache.GetUserOtpKey(in.Email)).Err()
	if err != nil {
		global.Logger.Error("Error deleting user OTP from Redis", zap.Error(err))
		return out, err
	}

	err = global.Redis.Del(ctx, cache.GetUserOtpAttemptKey(in.Email)).Err()
	if err != nil {
		global.Logger.Error("Error deleting user OTP attempt count from Redis", zap.Error(err))
		return out, err
	}

	// Update user status
	err = u.r.UpdateUserStatusByEmail(ctx, database.UpdateUserStatusByEmailParams{
		Email:  in.Email,
		Status: constants.USER_STATUS_ACTIVE,
	})
	if err != nil {
		global.Logger.Error("Error updating user status in DB", zap.Error(err))
		out.Code = response.ErrInternalError
		return out, err
	}

	// Output
	out.Code = response.CodeSuccess

	return out, nil
}
