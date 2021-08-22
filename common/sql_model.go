package common

import (
	"time"
)

type SQLModel struct {
	Id        int       `json:"id" gorm:"column:id;"`
	Status    int       `json:"status" gorm:"column:status;default:1;"` /// Set duoc default value
	CreatedAt time.Time `json:"create_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:updated_at;"`
}
