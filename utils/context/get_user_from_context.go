package context

import (
	"context"
	"github.com/feelgood-inc/flgd-gommon/models"
)

func GetUserFromContext(ctx context.Context, key string) models.User {
	user := ctx.Value(key)
	if user == nil {
		return models.User{}
	}

	return user.(models.User)
}
