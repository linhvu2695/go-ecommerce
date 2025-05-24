package context

import (
	"context"
	"errors"
	"go-ecommerce/internal/constants"
	"go-ecommerce/pkg/utils/cache"
)

type CtxUserInfo struct {
	ID       uint64
	Username string
	Email    string
}

func GetSubjectUUID(ctx context.Context) (string, error) {
	uuid, ok := ctx.Value(constants.SUBJECT_UUID_KEY).(string)
	if !ok {
		return "", errors.New("failed to get subject UUID")
	}

	return uuid, nil
}

func GetUserID(ctx context.Context) (uint64, error) {
	uuid, err := GetSubjectUUID(ctx)
	if err != nil {
		return 0, err
	}

	var userInfo CtxUserInfo
	if err := cache.GetCache(ctx, uuid, &userInfo); err != nil {
		return 0, err
	}

	return userInfo.ID, nil
}
