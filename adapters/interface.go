package adapters

import (
	"context"
	"github.com/feelgood-inc/flgd-gommon/models"
)

type Adapter interface {
	CreateClient(ctx context.Context, options *models.DBOptions) interface{}
}
