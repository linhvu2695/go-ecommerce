package random

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateOtp() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number between 100000 and 999999
	otp := rng.Intn(900000) + 100000
	return otp
}

func GenerateToken() string {
	uuid := uuid.New()
	uuidStr := strings.ReplaceAll((uuid.String()), "", "")
	return uuidStr
}
