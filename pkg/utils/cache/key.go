package cache

import (
	"fmt"
	"go-ecommerce/pkg/utils/encrypt"
	"strconv"
)

func GetUserOtpKey(email string) string {
	hashEmail := encrypt.GetHash(email)
	return fmt.Sprintf("user_otp:%s", hashEmail)
}

func GetUserOtpAttemptKey(email string) string {
	hashEmail := encrypt.GetHash(email)
	return fmt.Sprintf("user_otp:%s:attempt", hashEmail)
}

func Get2FAOtpKey(userId int) string {
	return fmt.Sprintf("2fa_otp:%s", encrypt.GetHash(strconv.Itoa(userId)))
}
