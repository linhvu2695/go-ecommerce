package models

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Token   string `json:"token"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RegisterInput struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Purpose   string `json:"purpose"`
}

type RegisterOutput struct {
	UserId  string `json:"user_id"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type VerifyOTPInput struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type VerifyOTPOutput struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UpdatePasswordInput struct {
	Email       string `json:"email"`
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type Setup2FARequest struct {
	Type  string `json:"type"`
	Email string `json:"email"`
}

type Setup2FAInput struct {
	UserId uint32 `json:"user_id"`
	Type   string `json:"type"`
	Email  string `json:"email"`
}

type Verify2FARequest struct {
	Code2FA string `json:"code_2fa"`
}

type Verify2FAInput struct {
	UserId  uint32 `json:"user_id"`
	Code2FA string `json:"code_2fa"`
}
