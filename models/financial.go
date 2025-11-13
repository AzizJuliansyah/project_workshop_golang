package models

import (
	"time"

	"gorm.io/gorm"
)

type FinancialBase struct{}

func (FinancialBase) TableName() string {
	return "financial"
}

type Financial struct {
	FinancialBase
	ID          int64          `json:"id"`
	Category    string         `json:"category"`
	Nominal     int64          `json:"nominal"`
	Information string         `json:"information"`
	CreatedAt   time.Time      `gorm:"autoCreateAt" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateAt" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type FinancialInput struct {
	Category    string `json:"category" validate:"required"`
	Nominal     int64  `json:"nominal" validate:"required"`
	Information string `json:"information" validate:"required,min=10"`
}
