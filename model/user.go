package model

import "gorm.io/gorm"

// User if you add sensitive fields, don't forget to clean them in setupLogin function.
// Otherwise, the sensitive information will be saved on local storage in plain text!
type User struct {
	Id        int            `json:"id"`
	CreatedAt int64          `gorm:"type:int;column:created_at;not null"`
	UpdatedAt int64          `gorm:"type:int;column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `json:"username" gorm:"unique;index"`
	Password  string         `json:"password" gorm:"not null;"`
}
