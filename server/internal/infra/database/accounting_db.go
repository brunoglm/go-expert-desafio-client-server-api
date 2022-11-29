package database

import (
	"context"
	"server-api/internal/entity"

	"gorm.io/gorm"
)

type Accounting struct {
	DB *gorm.DB
}

func NewAccounting(db *gorm.DB) *Accounting {
	return &Accounting{
		DB: db,
	}
}

func (a *Accounting) Create(ctx context.Context, usdbrlAccounting *entity.UsdbrlAccounting) error {
	return a.DB.WithContext(ctx).Create(usdbrlAccounting).Error
}
