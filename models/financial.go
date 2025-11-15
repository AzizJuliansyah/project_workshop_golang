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
	ID          int            `json:"id"`
	Category    string         `json:"category"`
	Nominal     int            `json:"nominal"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateAt" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateAt" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type FinancialInput struct {
	Category    string `json:"category" validate:"required"`
	Nominal     int    `json:"nominal" validate:"required,min=2"`
	Description string `json:"description" validate:"required,min=10"`
}
