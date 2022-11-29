package database

import (
	"context"
	"server-api/internal/entity"
)

type AccountingInterface interface {
	Create(ctx context.Context, usdbrlAccounting *entity.UsdbrlAccounting) error
}
