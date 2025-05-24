package response

const (
	CodeSuccess = 2001

	ErrInvalidToken = 3001
	ErrSendEmailOtp = 3002

	ErrInvalidRequest        = 4001
	ErrUserExists            = 4002
	ErrUserAlreadyRegistered = 4003
	ErrInvalidOTP            = 4004
	ErrMaxOTPAttemptsReached = 4005
	ErrAuthFailed            = 4006

	ErrInternalError = 5001

	Err2FASetupFailed  = 8001
	Err2FAVerifyFailed = 8002
)

var msg = map[int]string{
	CodeSuccess:              "success",
	ErrInvalidToken:          "token is invalid",
	ErrInvalidOTP:            "otp is invalid",
	ErrSendEmailOtp:          "failed to send email otp",
	ErrInvalidRequest:        "invalid request",
	ErrUserExists:            "user already exists",
	ErrUserAlreadyRegistered: "user already registered",
	ErrInternalError:         "internal server error",
	ErrAuthFailed:            "authentication failed",
	Err2FASetupFailed:        "two factor authentication setup failed",
	Err2FAVerifyFailed:       "two factor authentication verify failed",
}
