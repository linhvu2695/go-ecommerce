package constants

const (
	OTP_EXPIRATION_TIME  = 60 // 1 minute
	MAX_OTP_ATTEMPTS     = 3
	PASSWORD_SALT_LENGTH = 16

	// User status
	USER_STATUS_ACTIVE      = "active"
	USER_STATUS_DEACTIVATED = "deactivated"
	USER_STATUS_PENDING     = "pending"

	SUBJECT_UUID_KEY = "subjectUUID"
)
