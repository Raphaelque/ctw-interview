package model

import "gorm.io/gorm"

type Task struct {
	Id        int64          `json:"id"`
	CreatedAt int64          `gorm:"type:int;column:created_at;not null"`
	UpdatedAt int64          `gorm:"type:int;column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
